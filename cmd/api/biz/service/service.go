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

func UploadImagesAndGetUrl(data *multipart.FileHeader) (string, string, error) {
	ext := strings.ToLower(path.Ext(data.Filename))

	fileName := uuid.NewV4().String() + ext
	storePath := filepath.Join("static", "search", "images")

	if err := oss.SaveFile(data, storePath, fileName); err != nil {
		return "", "", err
	}

	url, err := oss.Upload(filepath.Join(storePath, fileName), "", fileName, "", "images")

	if err != nil {
		return "", "", err
	}

	return url, filepath.Join(storePath, fileName), nil
}

func UploadVideoAndGetUrl(data *multipart.FileHeader, userid string) (string, string, error) {
	ext := strings.ToLower(path.Ext(data.Filename))

	fileName := uuid.NewV4().String()
	storePath := filepath.Join("static", userid, "video")

	err := oss.SaveFile(data, filepath.Join(storePath, fileName), fileName+ext)
	if err != nil {
		return "", "", err
	}

	videoUrlChan := make(chan string)
	coverUrlChan := make(chan string)
	errChan := make(chan error)

	go func() {
		videoUrl, err := oss.Upload(filepath.Join(storePath, fileName, fileName+ext), fileName, fileName+ext, userid, "video")
		if err != nil {
			errChan <- err
			return
		}
		videoUrlChan <- videoUrl
	}()

	go func() {
		// 获取视频封面
		err = oss.GetVideoCover(filepath.Join(storePath, fileName), fileName, ext)
		if err != nil {
			errChan <- err
			return
		}

		// 上传封面文件
		coverUrl, err := oss.Upload(filepath.Join(storePath, fileName, fileName+".png"), fileName, fileName+".png", userid, "video")
		if err != nil {
			errChan <- err
			return
		}
		coverUrlChan <- coverUrl
	}()

	var videoUrl, coverUrl string
	for i := 0; i < 2; i++ {
		select {
		case url := <-videoUrlChan:
			videoUrl = url
		case url := <-coverUrlChan:
			coverUrl = url
		case err := <-errChan:
			return "", "", errmsg.ServiceError.WithMessage(err.Error())
		}
	}

	return videoUrl, coverUrl, nil
}
