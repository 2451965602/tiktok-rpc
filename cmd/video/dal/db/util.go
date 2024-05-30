package db

import (
	"strconv"
	"tiktokrpc/kitex_gen/user"
)

func NameToInfoRespToModel(userReq *user.NameToInfoResponse) *User {
	userid, _ := strconv.Atoi(userReq.Data.Id)

	return &User{
		Username:  userReq.Data.Username,
		UserId:    int64(userid),
		AvatarUrl: userReq.Data.AvatarUrl,
	}
}
