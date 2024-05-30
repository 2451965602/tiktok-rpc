package auth

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tiktokrpc/cmd/api/biz/middleware/jwt"
	"tiktokrpc/cmd/api/biz/pack"
	"tiktokrpc/cmd/api/pkg/errmsg"
)

func Auth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		DoubleTokenAuthFunc(),
	)
}

func DoubleTokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if !jwt.IsAccessTokenAvailable(ctx, c) {
			if !jwt.IsRefreshTokenAvailable(ctx, c) {
				pack.BuildFailResponse(c, errmsg.AuthError)
				c.Abort()

				return
			}
			jwt.GenerateAccessToken(c)
		}

		c.Next(ctx)
	}
}
