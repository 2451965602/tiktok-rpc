package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/netpoll"
	etcd "github.com/kitex-contrib/registry-etcd"
	internalopentracing "github.com/kitex-contrib/tracer-opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	_ "net/http/pprof"
	"os"
	"tiktokrpc/cmd/social/dal"
	"tiktokrpc/cmd/social/pkg/cfg"
	"tiktokrpc/cmd/social/pkg/constants"
	"tiktokrpc/cmd/social/pkg/pprof"
	"tiktokrpc/cmd/social/rpc"
	social "tiktokrpc/kitex_gen/social/socialservice"
)

func InitJaeger(service string) (server.Suite, io.Closer) {
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
	opentracing.SetGlobalTracer(tracer)
	return internalopentracing.NewDefaultServerSuite(), closer
}

func Init() io.Closer {
	err := cfg.Init()
	if err != nil {
		hlog.Info(err)
		os.Exit(2)
		return nil
	}
	dal.Init()
	closer := rpc.Init()
	return closer
}

func main() {

	pprof.Load()

	rpcCloser := Init()
	defer rpcCloser.Close()

	tracerSuite, closer := InitJaeger("tiktokrpc-social")
	defer closer.Close()

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	serviceAddr, err := netpoll.ResolveTCPAddr("tcp", constants.ServiceAddr)
	socialServiceImpl := new(SocialServiceImpl)

	svr := social.NewServer(socialServiceImpl,
		server.WithServiceAddr(serviceAddr),
		server.WithRegistry(r),
		server.WithSuite(tracerSuite),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "tiktokrpc.social",
			}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
