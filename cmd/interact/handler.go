package main

import (
	"context"
	"tiktokrpc/cmd/interact/pkg/errmsg"
	"tiktokrpc/cmd/interact/service"
	"tiktokrpc/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// Like implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Like(ctx context.Context, req *interact.LikeRequest) (resp *interact.LikeResponse, err error) {
	resp = new(interact.LikeResponse)

	err = service.NewInteractService(ctx).Like(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	return
}

// LikeList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) LikeList(ctx context.Context, req *interact.LikeListRequest) (resp *interact.LikeListResponse, err error) {
	resp = new(interact.LikeListResponse)

	LikeListResp, count, err := service.NewInteractService(ctx).LikeList(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.LikeList(LikeListResp, count)
	return
}

// Comment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) Comment(ctx context.Context, req *interact.CommentRequest) (resp *interact.CommentResponse, err error) {
	resp = new(interact.CommentResponse)

	err = service.NewInteractService(ctx).Comment(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.CommentListRequest) (resp *interact.CommentListResponse, err error) {
	resp = new(interact.CommentListResponse)

	CommentListResp, count, err := service.NewInteractService(ctx).CommentList(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)
	resp.Data = service.CommentList(CommentListResp, count)
	return
}

// DeleteComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) DeleteComment(ctx context.Context, req *interact.DeleteCommentRequest) (resp *interact.DeleteCommentResponse, err error) {
	resp = new(interact.DeleteCommentResponse)

	err = service.NewInteractService(ctx).DeleteComment(req)
	if err != nil {
		return nil, err
	}

	resp.Base = service.BuildBaseResp(errmsg.NoError)

	return
}
