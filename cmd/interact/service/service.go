package service

import (
	"context"
	"strconv"
	"tiktokrpc/cmd/interact/dal/db"
	"tiktokrpc/cmd/interact/dal/redis"
	"tiktokrpc/cmd/interact/pkg/errmsg"
	"tiktokrpc/cmd/interact/rpc"
	"tiktokrpc/kitex_gen/interact"
)

type InteractService struct {
	ctx context.Context
}

func NewInteractService(ctx context.Context) *InteractService {
	return &InteractService{ctx: ctx}
}

func (s *InteractService) Like(req *interact.LikeRequest) (err error) {

	if req.VideoId != nil && req.CommentId == nil {

		VideoID, err := strconv.ParseInt(*req.VideoId, 10, 64)
		if err != nil {
			return errmsg.ParseError
		}

		err = db.CreateLike(s.ctx, req.UserId, VideoID, req.ActionType, "Video")
		if err != nil {
			return err
		}

	} else if req.VideoId == nil && req.CommentId != nil {

		CommentID, err := strconv.ParseInt(*req.CommentId, 10, 64)
		if err != nil {
			return errmsg.ParseError
		}

		err = db.CreateLike(s.ctx, req.UserId, CommentID, req.ActionType, "Comment")
		if err != nil {
			return err
		}

	} else {
		return errmsg.DuplicationError.WithMessage("No liking of videos and comments at the same time")
	}

	if req.VideoId != nil {
		if req.ActionType == "1" {
			err = redis.AddLikeCount(s.ctx, *req.VideoId)
			if err != nil {
				return err
			}
		} else if req.ActionType == "2" {
			err = redis.ReduceLikeCount(s.ctx, *req.VideoId)
			if err != nil {
				return err
			}
		} else {
			return errmsg.ServiceError.WithMessage("parametric error")
		}

		VideoId, _ := strconv.Atoi(*req.VideoId)
		err = rpc.UpdataRank(int64(VideoId))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *InteractService) LikeList(req *interact.LikeListRequest) ([]*db.Video, int64, error) {

	var resp []*db.Video

	temp, count, err := db.LikeList(s.ctx, req.UserId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, -1, err
	}

	for _, v := range temp {
		count, err := redis.GetCounts(s.ctx, strconv.FormatInt(v.VideoId, 10))
		if err != nil {
			return nil, -1, err
		}
		v.VisitCount = count.VisitCount
		v.LikeCount = count.LikeCount
		v.CommentCount = count.CommentCount
		resp = append(resp, v)
	}

	return resp, count, nil
}

func (s *InteractService) Comment(req *interact.CommentRequest) error {

	if req.CommentId == nil && req.VideoId != nil {

		err := db.CreatComment(s.ctx, strconv.FormatInt(req.UserId, 10), *req.VideoId, req.Content, "video")
		if err != nil {
			return err
		}

		VideoId, _ := strconv.Atoi(*req.VideoId)
		err = rpc.UpdataRank(int64(VideoId))
		if err != nil {
			return err
		}

		err = redis.AddCommentCount(s.ctx, *req.VideoId)
		if err != nil {
			return err
		}

	} else if req.CommentId != nil && req.VideoId == nil {

		err := db.CreatComment(s.ctx, strconv.FormatInt(req.UserId, 10), *req.CommentId, req.Content, "comment")
		if err != nil {
			return err
		}

	} else {
		return errmsg.DuplicationError.WithMessage("No liking of videos and comments at the same time")
	}

	return nil
}

func (s *InteractService) CommentList(req *interact.CommentListRequest) (resp []*db.Comment, count int64, err error) {

	if req.VideoId != nil && req.CommentId == nil {
		resp, count, err = db.CommentList(s.ctx, *req.VideoId, req.PageNum, req.PageSize, "video")
	} else if req.CommentId != nil && req.VideoId == nil {
		resp, count, err = db.CommentList(s.ctx, *req.CommentId, req.PageNum, req.PageSize, "comment")
	} else {
		return nil, -1, errmsg.IllegalParamError
	}

	if err != nil {
		return nil, -1, err
	}

	return resp, count, nil
}

func (s *InteractService) DeleteComment(req *interact.DeleteCommentRequest) error {

	commentid, err := strconv.ParseInt(req.CommentId, 10, 64)
	if err != nil {
		return errmsg.ParseError
	}

	videoid, err := db.DeleteComment(s.ctx, strconv.FormatInt(req.UserId, 10), commentid)
	if err != nil {
		return errmsg.DatabaseError
	}

	err = redis.ReduceCommentCount(s.ctx, videoid)
	if err != nil {
		return err
	}

	return nil
}
