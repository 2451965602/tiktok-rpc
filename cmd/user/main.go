package main

import (
	"fmt"
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
	"tiktokrpc/cmd/user/dal"
	"tiktokrpc/cmd/user/pkg/cfg"
	"tiktokrpc/cmd/user/pkg/constants"
	"tiktokrpc/cmd/user/pkg/pprof"
	user "tiktokrpc/kitex_gen/user/userservice"
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

func Init() {
	err := cfg.Init()
	if err != nil {
		os.Exit(1)
		return
	}
	dal.Init()

}

func main() {

	pprof.Load()

	Init()

	tracerSuite, closer := InitJaeger("tiktokrpc-user")
	defer closer.Close()

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddr})
	if err != nil {
		log.Fatal(err)
	}

	serviceAddr, err := netpoll.ResolveTCPAddr("tcp", constants.ServiceAddr)
	userServiceImpl := new(UserServiceImpl)

	svr := user.NewServer(userServiceImpl,
		server.WithServiceAddr(serviceAddr),
		server.WithRegistry(r),
		server.WithSuite(tracerSuite),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "tiktokrpc.user",
			}))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
