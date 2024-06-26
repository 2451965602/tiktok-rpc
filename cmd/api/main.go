// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/opensergo/sentinel/adapter"
	"io"
	"os"
	"tiktokrpc/cmd/api/biz/middleware/jwt"
	"tiktokrpc/cmd/api/biz/router/websock"
	"tiktokrpc/cmd/api/biz/rpc"
	"tiktokrpc/cmd/api/pkg/cfg"
	"tiktokrpc/cmd/api/pkg/constants"
	"tiktokrpc/cmd/api/pkg/errmsg"
)

func Init() io.Closer {
	err := cfg.Init()
	if err != nil {
		hlog.Info(err.Error())
		os.Exit(1)
		return nil
	}
	closer := rpc.Init()
	jwt.Init()
	return closer
}

func main() {

	rpcCloser := Init()
	defer rpcCloser.Close()

	h := server.Default(
		server.WithHostPorts(constants.ServiceAddr),
		server.WithMaxRequestBodySize(1024*1024*1024),
	)

	h.Use(adapter.SentinelServerMiddleware(
		adapter.WithServerResourceExtractor(func(c context.Context, ctx *app.RequestContext) string {
			return "default"
		}),
		adapter.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
			ctx.AbortWithStatusJSON(429, utils.H{
				"base": utils.H{
					"code": errmsg.SentinelBlockCode,
					"msg":  errmsg.SentinelBlockMsg,
				},
			})
		}),
	))

	ws := server.Default(server.WithHostPorts(constants.WebsocketAddr))
	ws.NoHijackConnPool = true

	register(h)

	websock.WebsocketRegister(ws)
	go ws.Spin()
	h.Spin()
}
