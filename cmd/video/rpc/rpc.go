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
	"tiktokrpc/cmd/video/pkg/constants"
	"tiktokrpc/cmd/video/pkg/errmsg"
	"tiktokrpc/kitex_gen/user"
	"tiktokrpc/kitex_gen/user/userservice"
)

var userClient userservice.Client

func InitJaeger(service string) (client.Suite, io.Closer) {
	config := jaegercfg.Configuration{
		ServiceName: service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}
	tracer, closer, err := config.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return internalOpentracing.NewDefaultClientSuite(), closer
}

func Init() {
	tracerSuite, closer := InitJaeger("tiktokrpc-video")
	defer closer.Close()
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	c, err := userservice.NewClient("tiktokrpc.user", client.WithResolver(r), client.WithSuite(tracerSuite))
	if err != nil {
		log.Fatal(err)
	}
	userClient = c
}

func GetUserInfoByName(username string) (*user.NameToInfoResponse, error) {
	userReq := user.NewNameToInfoRequest()
	userReq.UserName = username
	userResp, err := userClient.NameToInfo(context.Background(), userReq)
	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}

func GetUserInfoById(userid string) (*user.InfoResponse, error) {
	userReq := user.NewInfoRequest()
	userReq.UserId = userid
	userResp, err := userClient.Info(context.Background(), userReq)
	if err != nil {
		return nil, errmsg.RpcCommunicationError
	} else if userResp.Base.Code != errmsg.NoErrorCode {
		return nil, errmsg.NewErrorMessage(userResp.Base.Code, userResp.Base.Msg)
	}

	return userResp, nil
}
