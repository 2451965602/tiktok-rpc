package main

import (
	"context"
	interact "tiktokrpc/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// Like implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Like(ctx context.Context, req *interact.LikeRequest) (resp *interact.LikeResponse, err error) {
	// TODO: Your code here...
	return
}

// LikeList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) LikeList(ctx context.Context, req *interact.LikeListRequest) (resp *interact.LikeListResponse, err error) {
	// TODO: Your code here...
	return
}

// Comment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Comment(ctx context.Context, req *interact.CommentRequest) (resp *interact.CommentResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) DeleteComment(ctx context.Context, req *interact.DeleteCommentRequest) (resp *interact.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	return
}
