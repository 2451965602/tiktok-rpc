package service

import (
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"
	"tiktokrpc/pkg/oss"
)

func UploadAndGetUrl(data *multipart.FileHeader, userid, sort string) (string, error) {
	ext := strings.ToLower(path.Ext(data.Filename))

	fileName := uuid.NewV4().String() + ext
	storePath := filepath.Join("static", userid, sort)

	if err := oss.SaveFile(data, storePath, fileName); err != nil {
		return "", err
	}

	url, err := oss.Upload(filepath.Join(storePath, fileName), fileName, userid, sort)

	if err != nil {
		return "", err
	}

	return url, nil
}
