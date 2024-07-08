package pack

import (
	"tiktokrpc/cmd/api/biz/model/model"
	"tiktokrpc/cmd/api/biz/model/user"
	rpcModel "tiktokrpc/kitex_gen/model"
	rpcUser "tiktokrpc/kitex_gen/user"
)

func ToUser(data *rpcModel.User) *model.User {

	return &model.User{
		ID:        data.Id,
		Username:  data.Username,
		Password:  data.Password,
		AvatarURL: data.AvatarUrl,
		OptSecret: data.OptSecret,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToUserInfo(data *rpcModel.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		ID:        data.Id,
		Username:  data.Username,
		AvatarURL: data.AvatarUrl,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func ToMFA(data *rpcModel.MFA) *model.MFA {

	return &model.MFA{
		Secret: data.Secret,
		Qrcode: data.Qrcode,
	}
}

func Register(userResp *rpcUser.RegisterResponse) (resp *user.RegisterResponse) {
	resp = new(user.RegisterResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToUser(userResp.Data)
	return
}

func Login(userResp *rpcUser.LoginResponse) (resp *user.LoginResponse) {
	resp = new(user.LoginResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToUserInfo(userResp.Data)

	return
}

func Info(userResp *rpcUser.InfoResponse) (resp *user.InfoResponse) {
	resp = new(user.InfoResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToUserInfo(userResp.Data)

	return
}

func UploadAvater(userResp *rpcUser.UploadResponse) (resp *user.UploadResponse) {
	resp = new(user.UploadResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)

	resp.Data = ToUser(userResp.Data)

	return
}

func AiUpload(userResp *rpcUser.UploadImagesResponse) (resp *user.UploadImagesResponse) {
	resp = new(user.UploadImagesResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)

	return
}

func AiSearch(userResp *rpcUser.SearchImagesResponse) (resp *user.SearchImagesResponse) {
	resp = new(user.SearchImagesResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = userResp.Data

	return
}

func MFAGet(userResp *rpcUser.MFAGetResponse) (resp *user.MFAGetResponse) {
	resp = new(user.MFAGetResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToMFA(userResp.Data)

	return
}

func MFA(userResp *rpcUser.MFABindResponse) (resp *user.MFABindResponse) {
	resp = new(user.MFABindResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)

	return
}

func MFAStatus(userResp *rpcUser.MFAStatusResponse) (resp *user.MFAStatusResponse) {
	resp = new(user.MFAStatusResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)

	return
}
