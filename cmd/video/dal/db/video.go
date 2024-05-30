package db

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"tiktokrpc/cmd/video/pkg/constants"
	"tiktokrpc/cmd/video/pkg/errmsg"
	"tiktokrpc/cmd/video/rpc"
	"tiktokrpc/kitex_gen/video"
	"time"
)

func Feed(ctx context.Context, req *video.FeedRequest) ([]*Video, int64, error) {

	var videoResp []*Video
	var err error
	var count int64

	if req.LatestTime != nil && *req.LatestTime != "" {
		toTime, err := strconv.ParseInt(*req.LatestTime, 10, 64)
		if err != nil {
			return nil, -1, errmsg.ParseError
		}

		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Where("created_at > ? ", time.Unix(toTime, 0)).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error

		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}
	} else {
		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error

		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}
	}

	return videoResp, count, nil
}

func UploadVideo(ctx context.Context, userid, videourl, coverurl, title, description string) (int64, error) {

	videoInfo := &Video{
		UserId:      userid,
		VideoUrl:    videourl,
		CoverUrl:    coverurl,
		Title:       title,
		Description: description,
	}

	err := DB.
		WithContext(ctx).
		Table(constants.VideoTable).
		Where("user_id=?", userid).
		Create(&videoInfo).
		Error

	if err != nil {
		return -1, errmsg.DatabaseError
	}

	return videoInfo.VideoId, nil
}

func UploadList(ctx context.Context, pagenum, pagesize int64, userid string) ([]*Video, int64, error) {

	var videoResp []*Video
	var err error
	var count int64

	_, err = rpc.GetUserInfoById(userid)
	if err != nil {
		return nil, -1, err
	}

	err = DB.
		WithContext(ctx).
		Table(constants.VideoTable).
		Where("user_id=?", userid).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&videoResp).
		Error

	if err != nil {
		return nil, -1, errmsg.DatabaseError
	}

	return videoResp, count, nil
}

func Rank(ctx context.Context, rank []string) ([]*Video, error) {

	var videoResp []*Video
	var err error
	var count int64

	err = DB.
		WithContext(ctx).
		Table("video").
		Where("video_id IN (?)", rank).
		Count(&count).
		Find(&videoResp).
		Error

	if err != nil {
		return nil, errmsg.DatabaseError
	}

	return videoResp, nil
}

func Query(ctx context.Context, req *video.QueryRequest) ([]*Video, int64, error) {

	var videoResp []*Video
	var err error
	var count int64

	if req.Keywords != nil && req.Username != nil && req.FromDate != nil && req.ToDate != nil {

		userInfoResp, err := rpc.GetUserInfoByName(*req.Username)
		if err != nil {
			return nil, -1, errmsg.UserNotExistError
		}
		userInfo := NameToInfoRespToModel(userInfoResp)

		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Where("id=?", userInfo.UserId).
			Where("created_at > ? and created_at < ?", time.Unix(*req.FromDate, 0), time.Unix(*req.ToDate, 0)).
			Where("title LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Or("description LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}

	} else if req.Keywords == nil && req.Username != nil && req.FromDate != nil && req.ToDate != nil {

		userInfoResp, err := rpc.GetUserInfoByName(*req.Username)
		if err != nil {
			return nil, -1, errmsg.UserNotExistError
		}
		userInfo := NameToInfoRespToModel(userInfoResp)

		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Where("id=?", userInfo.UserId).
			Where("created_at > ? and created_at < ?", time.Unix(*req.FromDate, 0), time.Unix(*req.ToDate, 0)).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}

	} else if req.Keywords != nil && req.Username == nil && req.FromDate != nil && req.ToDate != nil {

		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Where("created_at > ? and created_at < ?", time.Unix(*req.FromDate, 0), time.Unix(*req.ToDate, 0)).
			Where("title LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Or("description LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}
	}

	if req.Keywords != nil && req.Username != nil && req.FromDate == nil && req.ToDate == nil {

		userInfoResp, err := rpc.GetUserInfoByName(*req.Username)
		if err != nil {
			return nil, -1, errmsg.UserNotExistError
		}
		userInfo := NameToInfoRespToModel(userInfoResp)

		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Where("id=?", userInfo.UserId).
			Where("title LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Or("description LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}

	} else if req.Keywords == nil && req.Username != nil && req.FromDate == nil && req.ToDate == nil {

		userInfoResp, err := rpc.GetUserInfoByName(*req.Username)
		if err != nil {
			return nil, -1, errmsg.UserNotExistError
		}
		userInfo := NameToInfoRespToModel(userInfoResp)

		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Where("id=?", userInfo.UserId).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}

	} else if req.Keywords != nil && req.Username == nil && req.FromDate == nil && req.ToDate == nil {

		err = DB.
			WithContext(ctx).
			Table(constants.VideoTable).
			Where("title LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Or("description LIKE ?", fmt.Sprintf("%%%s%%", *req.Keywords)).
			Limit(int(req.PageSize)).
			Offset(int((req.PageNum - 1) * req.PageSize)).
			Count(&count).
			Find(&videoResp).
			Error
		if err != nil {
			return nil, -1, errmsg.DatabaseError
		}

	}

	return videoResp, count, nil
}

func UpdataVideoCounts(ctx context.Context, videoid string, counts Counts) error {
	err := DB.
		WithContext(ctx).
		Table(constants.VideoTable).
		Where("video_id=?", videoid).
		Update("like_count", counts.LikeCount).
		Update("comment_count", counts.CommentCount).
		Update("visit_count", counts.VisitCount).
		Error

	if err != nil {
		return errmsg.DatabaseError
	}

	return nil
}

func IsVideoExist(ctx context.Context, videoid int64) (bool, error) {
	var videoInfo Video

	err := DB.
		WithContext(ctx).
		Table(constants.VideoTable).
		Where("video_id=?", videoid).
		First(&videoInfo).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, errmsg.DatabaseError.WithMessage(err.Error())
	}

	return true, nil
}

func GetVideoById(ctx context.Context, videoid []int64, pagesize, pagenum int64) ([]*Video, int64, error) {
	var LikeResp []*Video
	var count int64

	err := DB.
		Table(constants.VideoTable).
		Where("`video_id` IN (?)", videoid).
		Limit(int(pagesize)).
		Offset(int((pagenum - 1) * pagesize)).
		Count(&count).
		Find(&LikeResp).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, -1, errmsg.VideoNotExistError
	} else if err != nil {
		return nil, -1, errmsg.DatabaseError.WithMessage(err.Error())
	}

	return LikeResp, count, nil
}
