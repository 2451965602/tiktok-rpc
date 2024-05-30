package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"tiktokrpc/cmd/interact/pkg/constants"
	"tiktokrpc/cmd/interact/pkg/errmsg"
	"tiktokrpc/kitex_gen/model"

	"time"
)

func IsCommentExist(ctx context.Context, commentid int64) (bool, error) {
	var comment Comment

	err := DB.
		WithContext(ctx).
		Table(constants.CommentTable).
		Where("comment_id=?", commentid).
		First(&comment).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, errmsg.DatabaseError.WithMessage(err.Error())
	}

	return true, nil
}

func TransVideo(data *model.Video) *Video {
	createdAtUnix, _ := strconv.ParseInt(data.CreatedAt, 10, 64)
	updatedAtUnix, _ := strconv.ParseInt(data.UpdatedAt, 10, 64)
	Id, _ := strconv.ParseInt(data.Id, 10, 64)

	return &Video{
		VideoId:      Id,
		UserId:       data.UserId,
		VideoUrl:     data.VideoUrl,
		CoverUrl:     data.CoverUrl,
		Title:        data.Title,
		Description:  data.Description,
		VisitCount:   data.VisitCount,
		LikeCount:    data.LikeCount,
		CommentCount: data.CommentCount,
		CreatedAt:    time.Unix(createdAtUnix, 0),
		UpdatedAt:    time.Unix(updatedAtUnix, 0),
	}
}

func GetVideoByIdRespToModel(data *model.VideoList) []*Video {
	resp := make([]*Video, 0, len(data.Items))

	for _, v := range data.Items {
		resp = append(resp, TransVideo(v))
	}

	return resp
}
