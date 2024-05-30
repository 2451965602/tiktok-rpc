package db

import (
	"strconv"
	"tiktokrpc/kitex_gen/user"
)

func InfoRespToModel(userReq *user.InfoResponse) *UserInfoDetail {
	userid, _ := strconv.Atoi(userReq.Data.Id)

	return &UserInfoDetail{
		Username:  userReq.Data.Username,
		UserId:    int64(userid),
		AvatarUrl: userReq.Data.AvatarUrl,
	}
}
