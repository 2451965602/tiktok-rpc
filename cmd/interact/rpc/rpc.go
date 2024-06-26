package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	internalOpentracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"tiktokrpc/cmd/interact/pkg/constants"
	"tiktokrpc/cmd/interact/pkg/errmsg"
	"tiktokrpc/kitex_gen/model"
	"tiktokrpc/kitex_gen/video"
	"tiktokrpc/kitex_gen/video/videoservice"
)

var videoClient videoservice.Client

func InitJaeger(service string) (client.Suite, io.Closer) {
	config := jaegercfg.Configuration{
		ServiceName: service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: constants.JaegerAddr,
		},
	}
	tracer, closer, err := config.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return internalOpentracing.NewDefaultClientSuite(), closer
}

func Init() io.Closer {
	tracerSuite, closer := InitJaeger("tiktokrpc-interact")

	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := videoservice.NewClient("tiktokrpc.video", client.WithResolver(r), client.WithSuite(tracerSuite))
	if err != nil {
		log.Fatal(err)
	}
	videoClient = c

	return closer
}

func IsVideoExist(videoid int64) (bool, error) {
	videoReq := video.NewIsExistRequest()
	videoReq.VideoId = videoid
	videoResp, err := videoClient.IsExist(context.Background(), videoReq)
	if err != nil {
		return false, errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return false, errmsg.NewErrorMessage(videoResp.Base.Code, videoResp.Base.Msg)
	}

	return videoResp.Data, nil
}

func GetVideoById(videoid []*int64) (*model.VideoList, error) {
	var videoIdList []int64

	videoReq := video.NewGetVideoByIdRequest()

	for _, id := range videoid {
		if id != nil {
			videoIdList = append(videoIdList, *id)
		}
	}

	videoReq.VideoId = videoIdList

	videoResp, err := videoClient.GetVideoById(context.Background(), videoReq)
	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(videoResp.Base.Code, videoResp.Base.Msg)
	}

	return videoResp.Data, nil
}

func UpdataRank(videoid int64) error {
	videoReq := video.NewUpdataRankRequest()
	videoReq.VideoId = videoid
	videoResp, err := videoClient.UpdataRank(context.Background(), videoReq)
	if err != nil {
		return errmsg.RpcCommunicationError
	} else if videoResp.Base.Code != errmsg.NoErrorCode {
		return errmsg.RedisError
	}

	return nil
}
