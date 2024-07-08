package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"mime/multipart"
	"strconv"
	"tiktokrpc/cmd/api/biz/model/user"
	"tiktokrpc/cmd/api/biz/pack"
	"tiktokrpc/cmd/api/biz/rpc"
	"tiktokrpc/cmd/api/pkg/oss"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

func (s *UserService) Register(req *user.RegisterRequest) (*user.RegisterResponse, error) {
	resp, err := rpc.Register(req)
	if err != nil {
		return nil, err
	}
	return pack.Register(resp), nil
}

func (s *UserService) Login(req *user.LoginRequest) (*user.LoginResponse, error) {
	resp, err := rpc.Login(req)
	if err != nil {
		return nil, err
	}
	return pack.Login(resp), nil
}

func (s *UserService) Info(req *user.InfoRequest) (*user.InfoResponse, error) {
	resp, err := rpc.Info(req)
	if err != nil {
		return nil, err
	}
	return pack.Info(resp), nil
}

func (s *UserService) Upload(avatar *multipart.FileHeader) (*user.UploadResponse, error) {
	userid := strconv.FormatInt(GetUidFormContext(s.c), 10)

	err := oss.IsImage(avatar)
	if err != nil {
		return nil, err
	}

	avatarUrl, err := UploadAvatarAndGetUrl(avatar, userid)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.UploadAvater(avatarUrl, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}
	return pack.UploadAvater(resp), nil
}

func (s *UserService) MFAGet(req *user.MFAGetRequest) (*user.MFAGetResponse, error) {
	resp, err := rpc.MFAGet(GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}
	return pack.MFAGet(resp), nil
}

func (s *UserService) MFA(req *user.MFABindRequest) (*user.MFABindResponse, error) {
	resp, err := rpc.MFA(req, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}
	return pack.MFA(resp), nil
}

func (s *UserService) MFAStatus(req *user.MFAStatusRequest) (*user.MFAStatusResponse, error) {
	resp, err := rpc.MFAStatus(req, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}
	return pack.MFAStatus(resp), nil
}

func (s *UserService) UploadImage(img *multipart.FileHeader) (*user.UploadImagesResponse, error) {

	err := oss.IsImage(img)
	if err != nil {
		return nil, err
	}

	imagesUrl, path, err := UploadImagesAndGetUrl(img)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.AiUploadImages(imagesUrl, path)
	if err != nil {
		return nil, err
	}

	return pack.AiUpload(resp), nil
}

func (s *UserService) SearchImages(img *multipart.FileHeader) (*user.SearchImagesResponse, error) {

	err := oss.IsImage(img)
	if err != nil {
		return nil, err
	}

	imagesUrl, path, err := UploadImagesAndGetUrl(img)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.AiSearchImages(imagesUrl, path)
	if err != nil {
		return nil, err
	}
	return pack.AiSearch(resp), nil
}
