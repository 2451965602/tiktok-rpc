package main

import (
	"context"
	video "tiktokrpc/kitex_gen/video"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// Upload implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Upload(ctx context.Context, req *video.UploadRequest) (resp *video.UploadResponse, err error) {
	// TODO: Your code here...
	return
}

// UploadList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UploadList(ctx context.Context, req *video.UploadListRequest) (resp *video.UploadListResponse, err error) {
	// TODO: Your code here...
	return
}

// Rank implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Rank(ctx context.Context, req *video.RankRequest) (resp *video.RankResponse, err error) {
	// TODO: Your code here...
	return
}

// Query implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Query(ctx context.Context, req *video.QueryRequest) (resp *video.QueryResponse, err error) {
	// TODO: Your code here...
	return
}
