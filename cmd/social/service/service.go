package service

import (
	"context"
	"strconv"
	"tiktokrpc/cmd/social/dal/db"
	"tiktokrpc/cmd/social/pkg/errmsg"
	"tiktokrpc/kitex_gen/social"
)

type SocialService struct {
	ctx context.Context
}

func NewSocialService(ctx context.Context) *SocialService {
	return &SocialService{ctx: ctx}
}

func (s *SocialService) Star(req *social.StarRequest) error {

	userid := req.UserId
	touserid, err := strconv.ParseInt(req.ToUserId, 10, 64)
	if err != nil {
		return errmsg.ParseError
	}

	if userid > touserid {
		return db.StarUser(s.ctx, userid, touserid, req.ActionType, 1)
	} else {
		return db.StarUser(s.ctx, touserid, userid, req.ActionType, 0)
	}

}

func (s *SocialService) StarList(req *social.StarListRequest) ([]*db.UserInfoDetail, int64, error) {

	userId, _ := strconv.ParseInt(req.UserId, 10, 64)

	resp, count, err := db.StarUserList(s.ctx, userId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, -1, err
	}

	return resp, count, nil
}

func (s *SocialService) FanList(req *social.FanListRequest) ([]*db.UserInfoDetail, int64, error) {

	resp, count, err := db.FanUserList(s.ctx, req.UserId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, -1, err
	}

	return resp, count, nil
}

func (s *SocialService) FriendList(req *social.FriendListRequest) ([]*db.UserInfoDetail, int64, error) {

	resp, count, err := db.FriendUser(s.ctx, strconv.FormatInt(req.UserId, 10), req.PageNum, req.PageSize)
	if err != nil {
		return nil, -1, err
	}

	return resp, count, nil
}
