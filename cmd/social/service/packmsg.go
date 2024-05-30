package service

import (
	"strconv"
	"tiktokrpc/cmd/social/dal/db"
	"tiktokrpc/cmd/social/pkg/constants"
	"tiktokrpc/cmd/social/pkg/errmsg"
	"tiktokrpc/kitex_gen/model"
)

func BuildBaseResp(err errmsg.ErrorMessage) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func UserInfo(data *db.UserInfoDetail) *model.UserInfo {
	createat := data.CreatedAt.Format(constants.TimeFormat)

	updateat := ""
	if !data.UpdatedAt.IsZero() {
		updateat = data.UpdatedAt.Format(constants.TimeFormat)
	} else {
		updateat = constants.DefaultTime
	}

	return &model.UserInfo{
		Id:        strconv.FormatInt(data.UserId, 10),
		Username:  data.Username,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: &createat,
		UpdatedAt: &updateat,
	}
}

func UserInfoList(data []*db.UserInfoDetail, total int64) *model.UserInfoList {
	resp := make([]*model.UserInfo, 0, len(data))

	for _, v := range data {
		resp = append(resp, UserInfo(v))
	}

	return &model.UserInfoList{
		Items: resp,
		Total: total,
	}
}
