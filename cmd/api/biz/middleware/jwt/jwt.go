package jwt

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"strconv"
	"tiktokrpc/cmd/api/biz/model/user"
	"tiktokrpc/cmd/api/biz/pack"
	"tiktokrpc/cmd/api/biz/rpc"
	"tiktokrpc/cmd/api/biz/service"
	"tiktokrpc/cmd/api/pkg/errmsg"
	"time"
)

var (
	IdentityKey               = "userid"
	AccessTokenJwtMiddleware  *jwt.HertzJWTMiddleware
	RefreshTokenJwtMiddleware *jwt.HertzJWTMiddleware
)

func AccessTokenJwt() {
	var err error
	AccessTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:                       "Video",
		Key:                         []byte("AccessToken_key"),
		Timeout:                     time.Hour,
		MaxRefresh:                  time.Hour,
		WithoutDefaultTokenHeadName: true,
		TokenLookup:                 "header: Access-Token",
		IdentityKey:                 IdentityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					AccessTokenJwtMiddleware.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[AccessTokenJwtMiddleware.IdentityKey]
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			pack.BuildFailResponse(c, errmsg.AuthError)
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("Access-Token", token)
		},

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct user.LoginRequest
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := rpc.Login(&loginStruct)
			if err != nil {
				return nil, err
			}
			uid, _ := strconv.Atoi(users.Data.Id)
			c.Set("userid", uid)
			return int64(uid), nil
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

func RefreshTokenJwt() {
	var err error
	RefreshTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:                       "Video",
		Key:                         []byte("refresh_token_key"),
		Timeout:                     time.Hour * 72,
		WithoutDefaultTokenHeadName: true,
		TokenLookup:                 "header: Refresh-Token",
		IdentityKey:                 IdentityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {

			return jwt.MapClaims{
				RefreshTokenJwtMiddleware.IdentityKey: data,
			}

		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)

			return claims[RefreshTokenJwtMiddleware.IdentityKey]
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			pack.BuildFailResponse(c, errmsg.AuthError)
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("Refresh-Token", token)
		},

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			uid, exist := c.Get("userid")
			if !exist {
				return nil, err
			}

			return uid, nil

		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

func GenerateAccessToken(c *app.RequestContext) {
	data := service.GetUidFormContext(c)
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(data)
	c.Header("New-Access-Token", tokenString)

}

func IsAccessTokenAvailable(ctx context.Context, c *app.RequestContext) bool {
	claims, err := AccessTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}
	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}
	c.Set("JWT_PAYLOAD", claims)
	identity := AccessTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(AccessTokenJwtMiddleware.IdentityKey, identity)
	}
	if !AccessTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}

	return true

}

func IsRefreshTokenAvailable(ctx context.Context, c *app.RequestContext) bool {

	claims, err := RefreshTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return false
	}

	switch v := claims["exp"].(type) {
	case nil:
		return false
	case float64:
		if int64(v) < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return false
		}
		if n < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return false
		}
	default:
		return false
	}

	c.Set("JWT_PAYLOAD", claims)
	identity := RefreshTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(RefreshTokenJwtMiddleware.IdentityKey, identity)
	}
	if !RefreshTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return false
	}

	return true
}

func Init() {
	AccessTokenJwt()
	RefreshTokenJwt()
	errInit := AccessTokenJwtMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("AccessTokenJwtMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	errInit = RefreshTokenJwtMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("RefreshTokenJwtMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
