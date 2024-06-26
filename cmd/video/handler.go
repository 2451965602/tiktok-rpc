package main

import (
	"context"
	"tiktokrpc/cmd/video/pkg/errmsg"
	"tiktokrpc/cmd/video/service"
	"tiktokrpc/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)

	videoResp, count, err := service.NewVideoService(ctx).Feed(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.VideoList(videoResp, count)
	return
}

// Upload implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Upload(ctx context.Context, req *video.UploadRequest) (resp *video.UploadResponse, err error) {
	resp = new(video.UploadResponse)

	err = service.NewVideoService(ctx).UploadVideo(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	return
}

// UploadList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UploadList(ctx context.Context, req *video.UploadListRequest) (resp *video.UploadListResponse, err error) {
	resp = new(video.UploadListResponse)

	videoResp, count, err := service.NewVideoService(ctx).UploadList(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.VideoList(videoResp, count)
	return
}

// Rank implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Rank(ctx context.Context, req *video.RankRequest) (resp *video.RankResponse, err error) {
	resp = new(video.RankResponse)

	videoResp, count, err := service.NewVideoService(ctx).Rank(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.VideoList(videoResp, count)

	return
}

// Query implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Query(ctx context.Context, req *video.QueryRequest) (resp *video.QueryResponse, err error) {
	resp = new(video.QueryResponse)

	videoResp, count, err := service.NewVideoService(ctx).Query(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.VideoList(videoResp, count)
	return
}

// IsExist implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) IsExist(ctx context.Context, req *video.IsExistRequest) (resp *video.IsExistResponse, err error) {
	resp = new(video.IsExistResponse)

	exist, err := service.NewVideoService(ctx).IsVideoExist(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = exist
	return
}

// GetVideoById implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoById(ctx context.Context, req *video.GetVideoByIdRequest) (resp *video.GetVideoByIdResponse, err error) {
	resp = new(video.GetVideoByIdResponse)

	videoResp, count, err := service.NewVideoService(ctx).GetVideoByid(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.VideoList(videoResp, count)
	return
}

// UpdataRank implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UpdataRank(ctx context.Context, req *video.UpdataRankRequest) (resp *video.UpdataRankResponse, err error) {
	resp = new(video.UpdataRankResponse)

	err = service.NewVideoService(ctx).UpdataRank(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)

	return
}
