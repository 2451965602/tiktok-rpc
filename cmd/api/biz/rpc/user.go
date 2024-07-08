package rpc

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"tiktokrpc/cmd/api/biz/model/user"
	"tiktokrpc/cmd/api/pkg/errmsg"
	rpcUser "tiktokrpc/kitex_gen/user"
)

func Register(req *user.RegisterRequest) (userResp *rpcUser.RegisterResponse, err error) {

	userReq := new(rpcUser.RegisterRequest)

	userReq.Username = req.Username
	userReq.Password = req.Password

	userResp, err = userClient.Register(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func Login(req *user.LoginRequest) (userResp *rpcUser.LoginResponse, err error) {

	userReq := new(rpcUser.LoginRequest)

	userReq.Username = req.Username
	userReq.Password = req.Password
	if req.Code != nil {
		userReq.Code = req.Code
	}

	userResp, err = userClient.Login(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		hlog.Info(userResp.Base.Msg)
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func Info(req *user.InfoRequest) (userResp *rpcUser.InfoResponse, err error) {

	userReq := new(rpcUser.InfoRequest)

	userReq.UserId = req.UserID

	userResp, err = userClient.Info(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func UploadAvater(avaterUrl string, userId int64) (userResp *rpcUser.UploadResponse, err error) {

	userReq := new(rpcUser.UploadRequest)

	userReq.UserId = userId
	userReq.AvatarUrl = avaterUrl

	userResp, err = userClient.Upload(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func MFAGet(userId int64) (userResp *rpcUser.MFAGetResponse, err error) {

	userReq := new(rpcUser.MFAGetRequest)

	userReq.UserId = userId

	userResp, err = userClient.MFAGet(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func MFA(req *user.MFABindRequest, userId int64) (userResp *rpcUser.MFABindResponse, err error) {

	userReq := new(rpcUser.MFABindRequest)

	userReq.Code = req.Code
	userReq.Secret = req.Secret
	userReq.UserId = userId

	userResp, err = userClient.MFA(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func MFAStatus(req *user.MFAStatusRequest, userId int64) (userResp *rpcUser.MFAStatusResponse, err error) {

	userReq := new(rpcUser.MFAStatusRequest)

	userReq.Code = req.Code
	userReq.ActionType = req.ActionType
	userReq.UserId = userId

	userResp, err = userClient.MFAStatus(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func AiUploadImages(url, path string) (userResp *rpcUser.UploadImagesResponse, err error) {

	userReq := new(rpcUser.UploadImagesRequest)

	userReq.ImgUrl = url
	userReq.ImgPath = "/media/fzuer/数据/golang/src/GoCode/Practice/tiktok-rpc/cmd/api/" + path
	userReq.CollectionName = "images"

	userResp, err = userClient.UploadImages(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func AiSearchImages(url, path string) (userResp *rpcUser.SearchImagesResponse, err error) {

	userReq := new(rpcUser.SearchImagesRequest)

	userReq.ImgPath = "/media/fzuer/数据/golang/src/GoCode/Practice/tiktok-rpc/cmd/api/" + path
	userReq.ImgUrl = url
	userReq.CollectionName = "images"

	userResp, err = userClient.SearchImages(context.Background(), userReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}
