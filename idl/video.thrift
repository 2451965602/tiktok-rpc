namespace go video

include "model.thrift"


//视频流
struct FeedRequest{
    1:optional string latest_time,
    2:required i64 page_num ,
    3:required i64 page_size 
}

struct FeedResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}

//投稿
struct UploadRequest{
    1:required string title ,
    2:required string cover_url,
    3:required string video_url ,
    4:required string description ,
    5:required i64 user_id ,
}

struct UploadResponse{
    1:model.BaseResp base,
}
//发布列表
struct UploadListRequest{
    1:required i64 user_id ,
    2:required i64 page_num ,
    3:required i64 page_size ,
}
struct UploadListResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}

//热门排行榜
struct RankRequest{
    1:required i64 page_num ,
    2:required i64 page_size ,
}
struct RankResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}
//搜索视频
struct QueryRequest{
    1:optional string keywords ,
    2:required i64 page_size ,
    3:required i64 page_num ,
    4:optional i64 from_date ,
    5:optional i64 to_date ,
    6:optional string username ,
}
struct QueryResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}

//是否存在视频

struct IsExistRequest{
    1:required i64 video_id,
}

struct IsExistResponse{
    1:model.BaseResp base,
    2:bool data,
}

//根据id获取视频
struct GetVideoByIdRequest{
    1:required list<i64>  video_id,
}

struct GetVideoByIdResponse{
    1:model.BaseResp base,
    2:model.VideoList data,
}

//排行榜
struct UpdataRankRequest{
    1:required i64 video_id,
}

struct UpdataRankResponse{
    1:model.BaseResp base,
}

service VideoService{
    FeedResponse Feed(1:FeedRequest req)(api.get="/video/feed"),
    UploadResponse Upload(1:UploadRequest req)(api.post="/video/publish"),
    UploadListResponse UploadList(1:UploadListRequest req)(api.get="/video/list"),
    RankResponse Rank(1:RankRequest req)(api.get="/video/popular"),
    QueryResponse Query(1:QueryRequest req)(api.post="/video/search")
    IsExistResponse IsExist(1:IsExistRequest req)(api.post="/video/exist")
    GetVideoByIdResponse GetVideoById(1:GetVideoByIdRequest req)(api.post="/video/getbyid")
    UpdataRankResponse UpdataRank(1:UpdataRankRequest req)(api.post="/video/updatarank")
}