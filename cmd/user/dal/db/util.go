package db

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"image/png"
	"net/url"
	"tiktokrpc/cmd/user/pkg/constants"
	"tiktokrpc/cmd/user/pkg/errmsg"
)

func IsUserNameExist(ctx context.Context, username string) (bool, error) {
	var user User

	err := DB.
		WithContext(ctx).
		Table(constants.UserTable).
		Where("username=?", username).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, errmsg.DatabaseError.WithMessage(err.Error())
	}

	return true, nil
}

func GetUserInfoByName(ctx context.Context, username string) (*User, error) {
	var user *User

	err := DB.
		WithContext(ctx).
		Table(constants.UserTable).
		Where("username = ?", username).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserInfo(ctx context.Context, userid string) (*User, error) {
	var user *User

	err := DB.
		WithContext(ctx).
		Table(constants.UserTable).
		Where("user_id=?", userid).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func OptSecret(user *User) (*MFA, error) {
	var MFAResp = &MFA{}
	var buf bytes.Buffer

	if user.OptSecret == "" {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "westonline",
			AccountName: user.Username,
		})
		if err != nil {
			return nil, errmsg.MfaGenareteError
		}

		user.OptSecret = key.String()
	}

	key, err := otp.NewKeyFromURL(user.OptSecret)
	if err != nil {
		return nil, errmsg.MfaGenareteError.WithMessage(err.Error())
	}

	img, err := key.Image(200, 200)
	if err != nil {
		return nil, errmsg.MfaGenareteError.WithMessage(err.Error())
	}

	err = png.Encode(&buf, img)
	if err != nil {
		return nil, errmsg.MfaGenareteError.WithMessage(err.Error())
	}

	qrcode := base64.StdEncoding.EncodeToString(buf.Bytes())

	secret, err := ExtractSecretFromTOTPURL(user.OptSecret)
	if err != nil {
		return nil, err
	}

	MFAResp.Secret = secret
	MFAResp.Qrcode = qrcode

	return MFAResp, nil
}

func ExtractSecretFromTOTPURL(totpURL string) (string, error) {
	parsedURL, err := url.Parse(totpURL)
	if err != nil {
		return "", errmsg.MfaGenareteError.WithMessage(err.Error())
	}

	// 获取查询参数
	queryParams := parsedURL.Query()

	// 从查询参数中提取 "secret"
	secret := queryParams.Get("secret")
	if secret == "" {
		return "", errmsg.MfaGenareteError.WithMessage("secret not found in TOTP URL")
	}

	return secret, nil
}

func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false
	} else {
		return true
	}
}
