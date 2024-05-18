package main

import (
	"context"
	"tiktokrpc/cmd/user/service"
	user "tiktokrpc/kitex_gen/user"
	"tiktokrpc/pkg/errmsg"
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
