package pack

import (
	"tiktokrpc/cmd/api/biz/model/interact"
	"tiktokrpc/cmd/api/biz/model/model"
	rpcInteract "tiktokrpc/kitex_gen/interact"
	rpcModel "tiktokrpc/kitex_gen/model"
)

func ToComment(data *rpcModel.Comment) *model.Comment {
	return &model.Comment{
		ID:        data.Id,
		UserID:    data.UserId,
		VideoID:   data.VideoId,
		RootID:    data.RootId,
		Content:   data.Content,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToLikeList(data []*rpcModel.Video, total int64) *model.LikeList {
	resp := make([]*model.Video, 0, len(data))

	for _, v := range data {
		resp = append(resp, ToVideo(v))
	}

	return &model.LikeList{
		Items: resp,
		Total: total,
	}
}

func ToCommentList(data []*rpcModel.Comment, total int64) *model.CommentList {
	resp := make([]*model.Comment, 0, len(data))

	for _, v := range data {
		resp = append(resp, ToComment(v))
	}

	return &model.CommentList{
		Items: resp,
		Total: total,
	}
}

func Like(interactResp *rpcInteract.LikeResponse) (resp *interact.LikeResponse) {
	resp = new(interact.LikeResponse)

	resp.Base = (*model.BaseResp)(interactResp.Base)

	return
}

func LikeList(interactResp *rpcInteract.LikeListResponse) (resp *interact.LikeListResponse) {
	resp = new(interact.LikeListResponse)

	resp.Base = (*model.BaseResp)(interactResp.Base)
	resp.Data = ToLikeList(interactResp.Data.Items, interactResp.Data.Total)
	return
}

func Comment(interactResp *rpcInteract.CommentResponse) (resp *interact.CommentResponse) {
	resp = new(interact.CommentResponse)

	resp.Base = (*model.BaseResp)(interactResp.Base)

	return
}

func CommentList(interactResp *rpcInteract.CommentListResponse) (resp *interact.CommentListResponse) {
	resp = new(interact.CommentListResponse)

	resp.Base = (*model.BaseResp)(interactResp.Base)
	resp.Data = ToCommentList(interactResp.Data.Items, interactResp.Data.Total)
	return
}

func DeleteComment(interactResp *rpcInteract.DeleteCommentResponse) (resp *interact.DeleteCommentResponse) {
	resp = new(interact.DeleteCommentResponse)

	resp.Base = (*model.BaseResp)(interactResp.Base)

	return
}
