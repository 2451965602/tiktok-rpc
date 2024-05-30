package service

import (
	"strconv"
	"tiktokrpc/cmd/video/dal/db"
	"tiktokrpc/cmd/video/pkg/errmsg"
	"tiktokrpc/kitex_gen/model"
)

func BuildBaseResp(err errmsg.ErrorMessage) *model.BaseResp {
	return &model.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
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

func VideoList(data []*db.Video, total int64) *model.VideoList {
	resp := make([]*model.Video, 0, len(data))

	for _, v := range data {
		resp = append(resp, Video(v))
	}

	return &model.VideoList{
		Items: resp,
		Total: total,
	}
}
