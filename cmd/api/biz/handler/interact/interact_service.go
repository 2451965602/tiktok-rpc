// Code generated by hertz generator.

package interact

import (
	"context"
	"tiktokrpc/cmd/api/biz/pack"
	"tiktokrpc/cmd/api/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
	"tiktokrpc/cmd/api/biz/model/interact"
)

// Like .
// @router /like/action [POST]
func Like(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.LikeRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	resp := new(interact.LikeResponse)

	resp, err = service.NewInteractService(ctx, c).Like(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}

// LikeList .
// @router /like/list [GET]
func LikeList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.LikeListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	resp := new(interact.LikeListResponse)

	resp, err = service.NewInteractService(ctx, c).LikeList(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}

// Comment .
// @router /comment/publish [POST]
func Comment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.CommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	resp := new(interact.CommentResponse)

	resp, err = service.NewInteractService(ctx, c).Comment(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}

// CommentList .
// @router /comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.CommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	resp := new(interact.CommentListResponse)

	resp, err = service.NewInteractService(ctx, c).CommentList(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}

// DeleteComment .
// @router /comment/delete [DELETE]
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.DeleteCommentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	resp := new(interact.DeleteCommentResponse)

	resp, err = service.NewInteractService(ctx, c).DeleteComment(&req)
	if err != nil {
		pack.BuildFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}