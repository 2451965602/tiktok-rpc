package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tiktokrpc/cmd/api/biz/model/social"
	"tiktokrpc/cmd/api/biz/pack"
	"tiktokrpc/cmd/api/biz/rpc"
)

type SocialService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewSocialService(ctx context.Context, c *app.RequestContext) *SocialService {
	return &SocialService{ctx: ctx, c: c}
}

func (s *SocialService) Star(req *social.StarRequest) (*social.StarResponse, error) {
	resp, err := rpc.Star(req, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}

	return pack.Star(resp), nil
}

func (s *SocialService) StarList(req *social.StarListRequest) (*social.StarListResponse, error) {
	resp, err := rpc.StarList(req)
	if err != nil {
		return nil, err
	}

	return pack.StarList(resp), nil
}

func (s *SocialService) FanList(req *social.FanListRequest) (*social.FanListResponse, error) {
	resp, err := rpc.FanList(req)
	if err != nil {
		return nil, err
	}

	return pack.FanList(resp), nil
}

func (s *SocialService) FriendList(req *social.FriendListRequest) (*social.FriendListResponse, error) {
	resp, err := rpc.FriendList(req, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}

	return pack.FriendList(resp), nil
}
