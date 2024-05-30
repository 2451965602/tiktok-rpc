package service

import (
	"strconv"
	"tiktokrpc/cmd/interact/dal/db"
	"tiktokrpc/cmd/interact/pkg/errmsg"
	"tiktokrpc/kitex_gen/model"
)

func BuildBaseResp(err errmsg.ErrorMessage) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

func Comment(data *db.Comment) *model.Comment {
	return &model.Comment{
		Id:        strconv.FormatInt(data.CommentId, 10),
		UserId:    data.UserId,
		VideoId:   data.VideoId,
		RootId:    data.RootId,
		Content:   data.Content,
		CreatedAt: strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdatedAt: strconv.FormatInt(data.UpdatedAt.Unix(), 10),
	}
}

func Video(data *db.Video) *model.Video {
	return &model.Video{
		Id:           strconv.FormatInt(data.VideoId, 10),
		UserId:       data.UserId,
		VideoUrl:     data.VideoUrl,
		CoverUrl:     data.CoverUrl,
		Title:        data.Title,
		Description:  data.Description,
		VisitCount:   data.VisitCount,
		LikeCount:    data.LikeCount,
		CommentCount: data.CommentCount,
		CreatedAt:    strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdatedAt:    strconv.FormatInt(data.UpdatedAt.Unix(), 10),
	}
}

func CommentList(data []*db.Comment, total int64) *model.CommentList {
	resp := make([]*model.Comment, 0, len(data))

	for _, v := range data {
		resp = append(resp, Comment(v))
	}

	return &model.CommentList{
		Items: resp,
		Total: total,
	}
}

func LikeList(data []*db.Video, total int64) *model.LikeList {
	resp := make([]*model.Video, 0, len(data))

	for _, v := range data {
		resp = append(resp, Video(v))
	}

	return &model.LikeList{
		Items: resp,
		Total: total,
	}
}
