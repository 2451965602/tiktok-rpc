package service

import (
	"github.com/cloudwego/hertz/pkg/app"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
	"tiktokrpc/cmd/api/pkg/constants"
	"tiktokrpc/cmd/api/pkg/errmsg"
	"tiktokrpc/cmd/api/pkg/oss"
)

func GetUidFormContext(c *app.RequestContext) int64 {
	uid, _ := c.Get(constants.ContextUid)
	userid, err := convertToInt64(uid)
	if err != nil {
		panic(err)
	}

	return userid
}

func convertToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	default:
		return 0, errmsg.ParseError
	}
}

func UploadAvatarAndGetUrl(data *multipart.FileHeader, userid string) (string, error) {
	ext := strings.ToLower(path.Ext(data.Filename))

	fileName := uuid.NewV4().String() + ext
	storePath := filepath.Join("static", userid, "avatar")

	if err := oss.SaveFile(data, storePath, fileName); err != nil {
		return "", err
	}

	url, err := oss.Upload(filepath.Join(storePath, fileName), "", fileName, userid, "avatar")

	if err != nil {
		return "", err
	}

	return url, nil
}

func UploadVideoAndGetUrl(data *multipart.FileHeader, userid string) (string, string, error) {
	ext := strings.ToLower(path.Ext(data.Filename))

	fileName := uuid.NewV4().String()
	storePath := filepath.Join("static", userid, "video")

	if err := oss.SaveFile(data, filepath.Join(storePath, fileName), fileName+ext); err != nil {
		return "", "", err
	}
	if err := oss.GetVideoCover(filepath.Join(storePath, fileName), fileName, ext); err != nil {
		return "", "", err
	}
	videoUrl, err := oss.Upload(filepath.Join(storePath, fileName, fileName+ext), fileName, fileName+ext, userid, "video")
	if err != nil {
		return "", "", err
	}
	coverUrl, err := oss.Upload(filepath.Join(storePath, fileName, fileName+".png"), fileName, fileName+".png", userid, "video")
	if err != nil {
		return "", "", err
	}

	return videoUrl, coverUrl, nil
}
