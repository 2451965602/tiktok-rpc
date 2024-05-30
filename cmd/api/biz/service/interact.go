package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tiktokrpc/cmd/api/biz/model/interact"
	"tiktokrpc/cmd/api/biz/pack"
	"tiktokrpc/cmd/api/biz/rpc"
)

type InteractService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewInteractService(ctx context.Context, c *app.RequestContext) *InteractService {
	return &InteractService{ctx: ctx, c: c}
}

func (s *InteractService) Like(req *interact.LikeRequest) (*interact.LikeResponse, error) {

	resp, err := rpc.Like(req, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}

	return pack.Like(resp), nil
}

func (s *InteractService) LikeList(req *interact.LikeListRequest) (*interact.LikeListResponse, error) {

	resp, err := rpc.LikeList(req)
	if err != nil {
		return nil, err
	}

	return pack.LikeList(resp), nil
}

func (s *InteractService) Comment(req *interact.CommentRequest) (*interact.CommentResponse, error) {

	resp, err := rpc.Comment(req, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}

	return pack.Comment(resp), nil
}

func (s *InteractService) CommentList(req *interact.CommentListRequest) (*interact.CommentListResponse, error) {

	resp, err := rpc.CommentList(req)
	if err != nil {
		return nil, err
	}

	return pack.CommentList(resp), nil
}

func (s *InteractService) DeleteComment(req *interact.DeleteCommentRequest) (*interact.DeleteCommentResponse, error) {

	resp, err := rpc.DeleteComment(req, GetUidFormContext(s.c))
	if err != nil {
		return nil, err
	}

	return pack.DeleteComment(resp), nil
}
