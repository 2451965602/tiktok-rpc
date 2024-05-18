package service

import (
	"strconv"
	"tiktokrpc/cmd/user/dal/db"
	"tiktokrpc/kitex_gen/model"
	"tiktokrpc/pkg/constants"
	"tiktokrpc/pkg/errmsg"
)

func BuildBaseResp(err errmsg.ErrorMessage) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func User(data *db.User) *model.User {
	create := strconv.FormatInt(data.CreatedAt.Unix(), 10)
	update := strconv.FormatInt(data.UpdatedAt.Unix(), 10)

	return &model.User{
		Id:        strconv.FormatInt(data.UserId, 10),
		Username:  data.Username,
		Password:  &data.Password,
		AvatarUrl: data.AvatarUrl,
		OptSecret: &data.OptSecret,
		CreatedAt: &create,
		UpdatedAt: &update,
	}
}

func UserInfoDetail(data *db.UserInfoDetail) *model.UserInfo {
	createat := data.CreatedAt.Format(constants.TimeFormat)

	updateat := ""
	if !data.UpdatedAt.IsZero() {
		updateat = data.UpdatedAt.Format(constants.TimeFormat)
	} else {
		updateat = constants.DefaultTime
	}

	deleteat := ""
	if !data.DeletedAt.Time.IsZero() {
		deleteat = data.DeletedAt.Time.Format(constants.TimeFormat)
	} else {
		deleteat = constants.DefaultTime
	}

	return &model.UserInfo{
		Id:        strconv.FormatInt(data.UserId, 10),
		Username:  data.Username,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: &createat,
		UpdatedAt: &updateat,
		DeletedAt: &deleteat,
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

func MFA(data *db.MFA) *model.MFA {

	return &model.MFA{
		Secret: data.Secret,
		Qrcode: data.Qrcode,
	}
}
