package pack

import (
	"tiktokrpc/cmd/api/biz/model/model"
	"tiktokrpc/cmd/api/biz/model/social"
	rpcModel "tiktokrpc/kitex_gen/model"
	rpcSocial "tiktokrpc/kitex_gen/social"
)

func ToUserInfoList(data []*rpcModel.UserInfo, total int64) *model.UserInfoList {
	resp := make([]*model.UserInfo, 0, len(data))

	for _, v := range data {
		resp = append(resp, ToUserInfo(v))
	}

	return &model.UserInfoList{
		Items: resp,
		Total: total,
	}
}

func Star(socialResp *rpcSocial.StarResponse) (resp *social.StarResponse) {
	resp = new(social.StarResponse)

	resp.Base = (*model.BaseResp)(socialResp.Base)

	return
}

func StarList(socialResp *rpcSocial.StarListResponse) (resp *social.StarListResponse) {
	resp = new(social.StarListResponse)

	resp.Base = (*model.BaseResp)(socialResp.Base)
	resp.Data = ToUserInfoList(socialResp.Data.Items, socialResp.Data.Total)
	return
}

func FanList(socialResp *rpcSocial.FanListResponse) (resp *social.FanListResponse) {
	resp = new(social.FanListResponse)

	resp.Base = (*model.BaseResp)(socialResp.Base)
	resp.Data = ToUserInfoList(socialResp.Data.Items, socialResp.Data.Total)
	return
}

func FriendList(socialResp *rpcSocial.FriendListResponse) (resp *social.FriendListResponse) {
	resp = new(social.FriendListResponse)

	resp.Base = (*model.BaseResp)(socialResp.Base)
	resp.Data = ToUserInfoList(socialResp.Data.Items, socialResp.Data.Total)
	return
}
