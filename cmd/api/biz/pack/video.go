package pack

import (
	"tiktokrpc/cmd/api/biz/model/model"
	"tiktokrpc/cmd/api/biz/model/video"
	rpcModel "tiktokrpc/kitex_gen/model"
	rpcVideo "tiktokrpc/kitex_gen/video"
)

func ToVideo(data *rpcModel.Video) *model.Video {
	return &model.Video{
		ID:           data.Id,
		UserID:       data.UserId,
		VideoURL:     data.VideoUrl,
		CoverURL:     data.CoverUrl,
		Title:        data.Title,
		Description:  data.Description,
		VisitCount:   data.VisitCount,
		LikeCount:    data.LikeCount,
		CommentCount: data.CommentCount,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToVideoList(data []*rpcModel.Video, total int64) *model.VideoList {
	resp := make([]*model.Video, 0, len(data))

	for _, v := range data {
		resp = append(resp, ToVideo(v))
	}

	return &model.VideoList{
		Items: resp,
		Total: total,
	}
}

func Feed(userResp *rpcVideo.FeedResponse) (resp *video.FeedResponse) {
	resp = new(video.FeedResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToVideoList(userResp.Data.Items, userResp.Data.Total)
	return
}

func UploadVideo(userResp *rpcVideo.UploadResponse) (resp *video.UploadResponse) {
	resp = new(video.UploadResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	return
}

func UploadList(userResp *rpcVideo.UploadListResponse) (resp *video.UploadListResponse) {
	resp = new(video.UploadListResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToVideoList(userResp.Data.Items, userResp.Data.Total)
	return
}

func Rank(userResp *rpcVideo.RankResponse) (resp *video.RankResponse) {
	resp = new(video.RankResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToVideoList(userResp.Data.Items, userResp.Data.Total)
	return
}

func Query(userResp *rpcVideo.QueryResponse) (resp *video.QueryResponse) {
	resp = new(video.QueryResponse)

	resp.Base = (*model.BaseResp)(userResp.Base)
	resp.Data = ToVideoList(userResp.Data.Items, userResp.Data.Total)
	return
}
