package rpc

import (
	"context"
	"tiktokrpc/cmd/api/biz/model/interact"
	"tiktokrpc/cmd/api/pkg/errmsg"
	rpcInteract "tiktokrpc/kitex_gen/interact"
)

func Like(req *interact.LikeRequest, userId int64) (interactResp *rpcInteract.LikeResponse, err error) {

	interactReq := new(rpcInteract.LikeRequest)

	interactReq.UserId = userId
	interactReq.ActionType = req.ActionType
	if req.VideoID != nil {
		interactReq.VideoId = req.VideoID
	}
	if req.CommentID != nil {
		interactReq.CommentId = req.CommentID
	}

	interactResp, err = interactClient.Like(context.Background(), interactReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if interactResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(interactResp.Base.Code, interactResp.Base.Msg)
	}

	return interactResp, nil
}

func LikeList(req *interact.LikeListRequest) (interactResp *rpcInteract.LikeListResponse, err error) {

	interactReq := new(rpcInteract.LikeListRequest)
	interactReq.UserId = req.UserID
	interactReq.PageSize = req.PageSize
	interactReq.PageNum = req.PageNum

	interactResp, err = interactClient.LikeList(context.Background(), interactReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if interactResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(interactResp.Base.Code, interactResp.Base.Msg)
	}

	return interactResp, nil
}

func Comment(req *interact.CommentRequest, userId int64) (interactResp *rpcInteract.CommentResponse, err error) {

	interactReq := new(rpcInteract.CommentRequest)

	interactReq.UserId = userId
	interactReq.Content = req.Content
	if req.VideoID != nil {
		interactReq.VideoId = req.VideoID
	}
	if req.CommentID != nil {
		interactReq.CommentId = req.CommentID
	}

	interactResp, err = interactClient.Comment(context.Background(), interactReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if interactResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(interactResp.Base.Code, interactResp.Base.Msg)
	}

	return interactResp, nil
}

func CommentList(req *interact.CommentListRequest) (interactResp *rpcInteract.CommentListResponse, err error) {

	interactReq := new(rpcInteract.CommentListRequest)

	interactReq.PageSize = req.PageSize
	interactReq.PageNum = req.PageNum
	if req.VideoID != nil {
		interactReq.VideoId = req.VideoID
	}
	if req.CommentID != nil {
		interactReq.CommentId = req.CommentID
	}

	interactResp, err = interactClient.CommentList(context.Background(), interactReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if interactResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(interactResp.Base.Code, interactResp.Base.Msg)
	}

	return interactResp, nil
}

func DeleteComment(req *interact.DeleteCommentRequest, userId int64) (interactResp *rpcInteract.DeleteCommentResponse, err error) {

	interactReq := new(rpcInteract.DeleteCommentRequest)

	interactReq.CommentId = req.CommentID
	interactReq.UserId = userId

	interactResp, err = interactClient.DeleteComment(context.Background(), interactReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError.WithMessage(err.Error())
	} else if interactResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(interactResp.Base.Code, interactResp.Base.Msg)
	}

	return interactResp, nil
}
