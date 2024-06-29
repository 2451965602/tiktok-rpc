package service

import (
	"context"
	"strconv"
	"tiktokrpc/cmd/video/dal/db"
	"tiktokrpc/cmd/video/dal/redis"
	"tiktokrpc/cmd/video/pkg/errmsg"
	"tiktokrpc/kitex_gen/video"
)

type VideoService struct {
	ctx context.Context
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (s *VideoService) Feed(req *video.FeedRequest) ([]*db.Video, int64, error) {

	var resp []*db.Video

	resp, num, err := db.Feed(s.ctx, req)
	if err != nil {
		return nil, -1, err
	}

	return resp, num, err
}

func (s *VideoService) UploadVideo(req *video.UploadRequest) error {

	videoId, err := db.UploadVideo(s.ctx, strconv.FormatInt(req.UserId, 10), req.VideoUrl, req.CoverUrl, req.Title, req.Description)
	if err != nil {
		return err
	}

	err = redis.AddIdToRank(s.ctx, strconv.FormatInt(videoId, 10))
	if err != nil {
		return err
	}

	return nil
}

func (s *VideoService) UploadList(req *video.UploadListRequest) ([]*db.Video, int64, error) {

	var resp []*db.Video

	resp, num, err := db.UploadList(s.ctx, req.PageNum, req.PageSize, strconv.FormatInt(req.UserId, 10))
	if err != nil {
		return nil, -1, err
	}

	return resp, num, err
}

func (s *VideoService) Rank(req *video.RankRequest) ([]*db.Video, int64, error) {

	var resp []*db.Video
	resp, err := redis.RankList(s.ctx)
	if err != nil {
		return nil, -1, err
	}

	errChan := make(chan error, 1)

	if resp == nil {

		rank, err := redis.IdRankList(s.ctx)
		if err != nil {
			return nil, -1, err
		}

		resp, err = db.Rank(s.ctx, rank)
		if err != nil {
			return nil, -1, err
		}

		go func() {
			err := redis.AddToRank(s.ctx, resp)
			errChan <- err
		}()
	} else {
		close(errChan) // 直接关闭通道，表示不再发送错误
	}

	startIndex := (req.PageNum - 1) * req.PageSize
	endIndex := startIndex + req.PageSize

	if startIndex >= int64(len(resp)) {
		return []*db.Video{}, 0, nil
	}

	if endIndex > int64(len(resp)) {
		endIndex = int64(len(resp))
	}

	select {
	case err := <-errChan:
		if err != nil {
			return nil, -1, err
		}
	default:

	}
	return resp[startIndex:endIndex], int64(len(resp)), nil

}

func (s *VideoService) Query(req *video.QueryRequest) ([]*db.Video, int64, error) {

	if req.Keywords == nil && req.Username == nil {

		return nil, -1, errmsg.IllegalParamError
	}

	var resp []*db.Video

	resp, num, err := db.Query(s.ctx, req)
	if err != nil {
		return nil, -1, err
	}

	return resp, num, err
}

func (s *VideoService) IsVideoExist(req *video.IsExistRequest) (bool, error) {

	exist, err := db.IsVideoExist(s.ctx, req.VideoId)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (s *VideoService) GetVideoByid(req *video.GetVideoByIdRequest) ([]*db.Video, int64, error) {

	likeResp, count, err := db.GetVideoById(s.ctx, req.VideoId)
	if err != nil {
		return nil, -1, err
	}

	return likeResp, count, err
}

func (s *VideoService) UpdataRank(req *video.UpdataRankRequest) error {

	err := redis.UpdateIdInRank(s.ctx, strconv.FormatInt(req.VideoId, 10))
	if err != nil {
		return err
	}

	return err
}
