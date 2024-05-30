package rpc

import (
	"context"
	"strconv"
	"tiktokrpc/cmd/api/biz/model/video"
	"tiktokrpc/cmd/api/pkg/errmsg"
	rpcVideo "tiktokrpc/kitex_gen/video"
)

func Feed(req *video.FeedRequest) (videoResp *rpcVideo.FeedResponse, err error) {

	videoReq := new(rpcVideo.FeedRequest)

	videoReq.PageSize = req.PageSize
	videoReq.PageNum = req.PageNum
	if req.LatestTime != nil {
		videoReq.LatestTime = req.LatestTime
	}

	videoResp, err = videoClient.Feed(context.Background(), videoReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(videoResp.Base.Code, videoResp.Base.Msg)
	}

	return videoResp, nil
}

func UploadVideo(req *video.UploadRequest, videoUrl, coverUrl string, userId int64) (videoResp *rpcVideo.UploadResponse, err error) {

	videoReq := new(rpcVideo.UploadRequest)

	videoReq.Title = req.Title
	videoReq.Description = req.Description
	videoReq.UserId = userId
	videoReq.VideoUrl = videoUrl
	videoReq.CoverUrl = coverUrl

	videoResp, err = videoClient.Upload(context.Background(), videoReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(videoResp.Base.Code, videoResp.Base.Msg)
	}

	return videoResp, nil
}

func UploadList(req *video.UploadListRequest) (videoResp *rpcVideo.UploadListResponse, err error) {

	videoReq := new(rpcVideo.UploadListRequest)

	userid, _ := strconv.Atoi(req.UserID)

	videoReq.PageSize = req.PageSize
	videoReq.PageNum = req.PageNum
	videoReq.UserId = int64(userid)

	videoResp, err = videoClient.UploadList(context.Background(), videoReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(videoResp.Base.Code, videoResp.Base.Msg)
	}

	return videoResp, nil
}

func Rank(req *video.RankRequest) (videoResp *rpcVideo.RankResponse, err error) {

	videoReq := new(rpcVideo.RankRequest)

	videoReq.PageSize = req.PageSize
	videoReq.PageNum = req.PageNum

	videoResp, err = videoClient.Rank(context.Background(), videoReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(videoResp.Base.Code, videoResp.Base.Msg)
	}

	return videoResp, nil
}

func Query(req *video.QueryRequest) (videoResp *rpcVideo.QueryResponse, err error) {

	videoReq := new(rpcVideo.QueryRequest)

	videoReq.PageSize = req.PageSize
	videoReq.PageNum = req.PageNum

	if req.Username != nil {
		videoReq.Username = req.Username
	}
	if req.FromDate != nil {
		videoReq.FromDate = req.FromDate
	}
	if req.ToDate != nil {
		videoReq.ToDate = req.ToDate
	}
	if req.Keywords != nil {
		videoReq.Keywords = req.Keywords
	}

	videoResp, err = videoClient.Query(context.Background(), videoReq)

	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(videoResp.Base.Code, videoResp.Base.Msg)
	}

	return videoResp, nil
}
