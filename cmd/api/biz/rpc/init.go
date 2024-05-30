package rpc

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	internalOpentracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"tiktokrpc/kitex_gen/interact/interactservice"
	"tiktokrpc/kitex_gen/social/socialservice"
	"tiktokrpc/kitex_gen/user/userservice"
	"tiktokrpc/kitex_gen/video/videoservice"
)

var userClient userservice.Client
var interactClient interactservice.Client
var videoClient videoservice.Client
var socialClient socialservice.Client

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

	tracerSuite, closer := InitJaeger("tiktokrpc-api")
	defer closer.Close()

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	userClient, err = userservice.NewClient("tiktokrpc.user", client.WithResolver(r), client.WithSuite(tracerSuite))
	if err != nil {
		log.Fatal(err)
	}

	interactClient, err = interactservice.NewClient("tiktokrpc.interact", client.WithResolver(r), client.WithSuite(tracerSuite))
	if err != nil {
		log.Fatal(err)
	}

	videoClient, err = videoservice.NewClient("tiktokrpc.video", client.WithResolver(r), client.WithSuite(tracerSuite))
	if err != nil {
		log.Fatal(err)
	}

	socialClient, err = socialservice.NewClient("tiktokrpc.social", client.WithResolver(r), client.WithSuite(tracerSuite))
	if err != nil {
		log.Fatal(err)
	}

}
