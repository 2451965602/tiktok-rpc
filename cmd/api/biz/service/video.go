package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"mime/multipart"
	"strconv"
	"tiktokrpc/cmd/api/biz/model/video"
	"tiktokrpc/cmd/api/biz/pack"
	"tiktokrpc/cmd/api/biz/rpc"
	"tiktokrpc/cmd/api/pkg/oss"
)

type VideoService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewVideoService(ctx context.Context, c *app.RequestContext) *VideoService {
	return &VideoService{ctx: ctx, c: c}
}

func (s *VideoService) Feed(req *video.FeedRequest) (*video.FeedResponse, error) {
	resp, err := rpc.Feed(req)
	if err != nil {
		return nil, err
	}
	return pack.Feed(resp), nil

}

func (s *VideoService) Upload(videoData *multipart.FileHeader, req *video.UploadRequest) (*video.UploadResponse, error) {

	userid := strconv.FormatInt(GetUidFormContext(s.c), 10)

	err := oss.IsVideo(videoData)
	if err != nil {
		return nil, err
	}

	videoUrl, coverUrl, err := UploadVideoAndGetUrl(videoData, userid)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.UploadVideo(req, videoUrl, coverUrl, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}
	return pack.UploadVideo(resp), nil

}

func (s *VideoService) UploadList(req *video.UploadListRequest) (*video.UploadListResponse, error) {
	resp, err := rpc.UploadList(req)
	if err != nil {
		return nil, err
	}
	return pack.UploadList(resp), nil

}

func (s *VideoService) Rank(req *video.RankRequest) (*video.RankResponse, error) {
	resp, err := rpc.Rank(req)
	if err != nil {
		return nil, err
	}
	return pack.Rank(resp), nil

}

func (s *VideoService) Query(req *video.QueryRequest) (*video.QueryResponse, error) {
	resp, err := rpc.Query(req)
	if err != nil {
		return nil, err
	}
	return pack.Query(resp), nil

}
