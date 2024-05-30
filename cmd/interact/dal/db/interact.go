package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"tiktokrpc/cmd/interact/pkg/constants"
	"tiktokrpc/cmd/interact/pkg/errmsg"
	"tiktokrpc/cmd/interact/rpc"
)

func CreateLike(ctx context.Context, userid, id int64, actiontype, sort string) (err error) {

	var LikeResp *Like

	if sort == "Video" {

		exist, err := rpc.IsVideoExist(id)
		if err != nil {
			return err
		}
		if !exist {
			return errmsg.VideoNotExistError
		}

		LikeResp = &Like{
			UserId:  userid,
			VideoId: id,
			RootId:  0,
		}

	} else if sort == "Comment" {

		exist, err := IsCommentExist(ctx, id)
		if err != nil {
			return err
		}
		if !exist {
			return errmsg.CommentNotExistError
		}

		LikeResp = &Like{
			UserId:  userid,
			VideoId: 0,
			RootId:  id,
		}

	} else {
		return errmsg.ServiceError.WithMessage("parametric error")
	}

	if actiontype == "1" {

		err = DB.
			WithContext(ctx).
			Table(constants.LikeTable).
			Where("user_id=?", userid).
			Where("root_id=?", id).
			Or("video_id=?", id).
			First(&LikeResp).
			Error

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errmsg.DuplicationError.WithMessage("Currently liked, please do not repeat the action")
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return errmsg.DatabaseError.WithMessage(err.Error())
		}

		err = DB.
			WithContext(ctx).
			Table(constants.LikeTable).
			Create(&LikeResp).
			Error

		if err != nil {
			return errmsg.DatabaseError.WithMessage(err.Error())
		}

	} else if actiontype == "2" {

		err = DB.
			WithContext(ctx).
			Table(constants.LikeTable).
			Where("user_id=?", userid).
			Where("root_id=?", id).
			Or("video_id=?", id).
			First(&LikeResp).
			Error

		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return errmsg.DuplicationError.WithMessage("Currently not liked, please do not repeat the action")
		}

		err = DB.
			WithContext(ctx).
			Table(constants.LikeTable).
			Where("user_id=?", userid).
			Where("root_id=?", id).
			Or("video_id=?", id).
			Delete(&LikeResp).
			Error

		if err != nil {
			return errmsg.DatabaseError.WithMessage(err.Error())
		}
	} else {
		return errmsg.ServiceError.WithMessage("parametric error")
	}

	return nil
}

func LikeList(ctx context.Context, userid string, pagenum, pagesize int64) ([]*Video, int64, error) {

	var videoid []*int64

	err := DB.
		WithContext(ctx).
		Table(constants.LikeTable).
		Where("user_id = ?", userid).
		Select("video_id").
		Find(&videoid).
		Error

	if err != nil {
		return nil, -1, errmsg.DatabaseError.WithMessage(err.Error())
	}

	LikeResp, count, err := rpc.GetVideoById(videoid, pagesize, pagenum)
	if err != nil {
		return nil, -1, err
	}

	videoList := GetVideoByIdRespToModel(LikeResp)

	return videoList, count, nil
}

func CreatComment(ctx context.Context, userid, id, content, sort string) error {

	var CommentResp *Comment

	if sort == "video" {
		video, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errmsg.ParseError
		}

		exist, err := rpc.IsVideoExist(video)
		if err != nil {
			return err
		}
		if !exist {
			return errmsg.VideoNotExistError
		}

		CommentResp = &Comment{
			UserId:  userid,
			VideoId: id,
			Content: content,
		}

	} else if sort == "comment" {

		commentid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return errmsg.ParseError
		}

		exist, err := IsCommentExist(ctx, commentid)
		if err != nil {
			return err
		}
		if !exist {
			return errmsg.CommentNotExistError
		}

		CommentResp = &Comment{
			UserId:  userid,
			RootId:  id,
			Content: content,
		}

	} else {
		return errmsg.ServiceError.WithMessage("parametric error")
	}

	err := DB.
		WithContext(ctx).
		Table(constants.CommentTable).
		Create(&CommentResp).
		Error

	if err != nil {
		return errmsg.DatabaseError.WithMessage(err.Error())
	}

	return nil
}

func CommentList(ctx context.Context, id string, pagenum, pagesize int64, sort string) (CommentResp []*Comment, count int64, err error) {

	if sort == "video" {
		err = DB.
			WithContext(ctx).
			Table(constants.CommentTable).
			Where("video_id=?", id).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Count(&count).
			Find(&CommentResp).
			Error

	} else {
		err = DB.
			WithContext(ctx).
			Table(constants.CommentTable).
			Where("root_id=?", id).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Count(&count).
			Find(&CommentResp).
			Error
	}

	if err != nil {
		return nil, -1, errmsg.DatabaseError.WithMessage(err.Error())
	}

	return CommentResp, count, nil
}

func DeleteComment(ctx context.Context, userid string, commentid int64) (string, error) {

	var commentInfo Comment

	exist, err := IsCommentExist(ctx, commentid)
	if err != nil {
		return "-1", err
	}
	if !exist {
		return "-1", errmsg.CommentNotExistError
	}

	err = DB.
		WithContext(ctx).
		Table(constants.CommentTable).
		Where("comment_id=?", commentid).
		Select("video_id").
		First(&commentInfo).
		Error

	if err != nil {
		return "-1", errmsg.DatabaseError.WithMessage(err.Error())
	}

	err = DB.
		WithContext(ctx).
		Table(constants.CommentTable).
		Where("comment_id = ?", commentid).
		Delete(&Comment{
			CommentId: commentid,
			UserId:    userid,
		}).
		Error

	if err != nil {
		return "-1", errmsg.DatabaseError.WithMessage(err.Error())
	}

	return commentInfo.VideoId, nil
}
