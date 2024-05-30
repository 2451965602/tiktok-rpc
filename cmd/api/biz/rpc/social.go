package rpc

import (
	"context"
	"tiktokrpc/cmd/api/biz/model/social"
	"tiktokrpc/cmd/api/pkg/errmsg"
	rpcSocial "tiktokrpc/kitex_gen/social"
)

func Star(req *social.StarRequest, userId int64) (socialResp *rpcSocial.StarResponse, err error) {

	socialReq := new(rpcSocial.StarRequest)

	socialReq.ToUserId = req.ToUserID
	socialReq.UserId = userId
	socialReq.ActionType = req.ActionType

	socialResp, err = socialClient.Star(context.Background(), socialReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if socialResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(socialResp.Base.Code, socialResp.Base.Msg)
	}

	return socialResp, nil
}

func StarList(req *social.StarListRequest) (socialResp *rpcSocial.StarListResponse, err error) {

	socialReq := new(rpcSocial.StarListRequest)

	socialReq.UserId = req.UserID
	socialReq.PageSize = req.PageSize
	socialReq.PageNum = req.PageNum

	socialResp, err = socialClient.StarList(context.Background(), socialReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if socialResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(socialResp.Base.Code, socialResp.Base.Msg)
	}

	return socialResp, nil
}

func FanList(req *social.FanListRequest) (socialResp *rpcSocial.FanListResponse, err error) {

	socialReq := new(rpcSocial.FanListRequest)

	socialReq.UserId = req.UserID
	socialReq.PageSize = req.PageSize
	socialReq.PageNum = req.PageNum

	socialResp, err = socialClient.FanList(context.Background(), socialReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if socialResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(socialResp.Base.Code, socialResp.Base.Msg)
	}

	return socialResp, nil
}

func FriendList(req *social.FriendListRequest, userId int64) (socialResp *rpcSocial.FriendListResponse, err error) {

	socialReq := new(rpcSocial.FriendListRequest)

	socialReq.UserId = userId
	socialReq.PageSize = req.PageSize
	socialReq.PageNum = req.PageNum

	socialResp, err = socialClient.FriendList(context.Background(), socialReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if socialResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(socialResp.Base.Code, socialResp.Base.Msg)
	}

	return socialResp, nil
}
