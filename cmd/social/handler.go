package main

import (
	"context"
	"tiktokrpc/cmd/social/pkg/errmsg"
	"tiktokrpc/cmd/social/service"
	"tiktokrpc/kitex_gen/social"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// Star implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) Star(ctx context.Context, req *social.StarRequest) (resp *social.StarResponse, err error) {
	resp = new(social.StarResponse)

	err = service.NewSocialService(ctx).Star(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	return
}

// StarList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) StarList(ctx context.Context, req *social.StarListRequest) (resp *social.StarListResponse, err error) {
	resp = new(social.StarListResponse)

	SocialResp, count, err := service.NewSocialService(ctx).StarList(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.UserInfoList(SocialResp, count)
	return
}

// FanList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FanList(ctx context.Context, req *social.FanListRequest) (resp *social.FanListResponse, err error) {
	resp = new(social.FanListResponse)

	SocialResp, count, err := service.NewSocialService(ctx).FanList(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.UserInfoList(SocialResp, count)
	return
}

// FriendList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FriendList(ctx context.Context, req *social.FriendListRequest) (resp *social.FriendListResponse, err error) {
	resp = new(social.FriendListResponse)

	SocialResp, count, err := service.NewSocialService(ctx).FriendList(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.UserInfoList(SocialResp, count)
	return
}
