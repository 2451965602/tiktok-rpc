package main

import (
	"context"
	social "tiktokrpc/kitex_gen/social"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// Star implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) Star(ctx context.Context, req *social.StarRequest) (resp *social.StarResponse, err error) {
	// TODO: Your code here...
	return
}

// StarList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) StarList(ctx context.Context, req *social.StarListRequest) (resp *social.StarListResponse, err error) {
	// TODO: Your code here...
	return
}

// FanList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FanList(ctx context.Context, req *social.FanListRequest) (resp *social.FanListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) FriendList(ctx context.Context, req *social.FriendListRequest) (resp *social.FriendListResponse, err error) {
	// TODO: Your code here...
	return
}
