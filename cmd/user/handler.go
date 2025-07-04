package main

import (
	"context"
	"tiktokrpc/cmd/user/pkg/errmsg"
	"tiktokrpc/cmd/user/service"
	"tiktokrpc/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	userResp, err := service.NewUserService(ctx).Register(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.User(userResp)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	userResp, err := service.NewUserService(ctx).Login(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.UserInfoDetail(userResp)

	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	resp = new(user.InfoResponse)

	userResp, err := service.NewUserService(ctx).GetInfo(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.UserInfoDetail(userResp)

	return
}

// Upload implements the UserServiceImpl interface.
func (s *UserServiceImpl) Upload(ctx context.Context, req *user.UploadRequest) (resp *user.UploadResponse, err error) {
	resp = new(user.UploadResponse)

	userResp, err := service.NewUserService(ctx).UploadAvatar(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.User(userResp)

	return
}

// MFAGet implements the UserServiceImpl interface.
func (s *UserServiceImpl) MFAGet(ctx context.Context, req *user.MFAGetRequest) (resp *user.MFAGetResponse, err error) {
	resp = new(user.MFAGetResponse)

	userResp, err := service.NewUserService(ctx).MFAGet(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.MFA(userResp)

	return
}

// MFA implements the UserServiceImpl interface.
func (s *UserServiceImpl) MFA(ctx context.Context, req *user.MFABindRequest) (resp *user.MFABindResponse, err error) {
	resp = new(user.MFABindResponse)

	err = service.NewUserService(ctx).MFABind(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	return
}

// MFAStatus implements the UserServiceImpl interface.
func (s *UserServiceImpl) MFAStatus(ctx context.Context, req *user.MFAStatusRequest) (resp *user.MFAStatusResponse, err error) {
	resp = new(user.MFAStatusResponse)

	err = service.NewUserService(ctx).MFAStatus(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)

	return
}

// NameToInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) NameToInfo(ctx context.Context, req *user.NameToInfoRequest) (resp *user.NameToInfoResponse, err error) {
	resp = new(user.NameToInfoResponse)

	userResp, err := service.NewUserService(ctx).GetInfoByName(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.UserInfoDetail(userResp)

	return
}

// UploadImages implements the UserServiceImpl interface.
func (s *UserServiceImpl) UploadImages(ctx context.Context, req *user.UploadImagesRequest) (resp *user.UploadImagesResponse, err error) {
	resp = new(user.UploadImagesResponse)

	err = service.NewUserService(ctx).AiUploadImages(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)

	return
}

// SearchImages implements the UserServiceImpl interface.
func (s *UserServiceImpl) SearchImages(ctx context.Context, req *user.SearchImagesRequest) (resp *user.SearchImagesResponse, err error) {
	resp = new(user.SearchImagesResponse)

	userResp, err := service.NewUserService(ctx).AiSearchImages(req)

	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = userResp

	return
}
