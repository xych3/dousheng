package mw

import (
	"context"
	"gitee.com/derrickball/douyin/app/api/biz/model/user"
	"gitee.com/derrickball/douyin/app/api/biz/rpc"
	"gitee.com/derrickball/douyin/app/user/kitex_gen/user_rpc"
	"gitee.com/derrickball/douyin/pkg/constants"
	"gitee.com/derrickball/douyin/pkg/errno"
	"gitee.com/derrickball/douyin/pkg/util"
	"github.com/cloudwego/hertz/pkg/app"
	jwt "github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

var AuthMiddleware *jwt.HertzJWTMiddleware

func InitAuth() {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(constants.SecretKey),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, form: token",
		TokenHeadName: "Bearer",
		IdentityKey:   constants.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			t, _ := AuthMiddleware.ParseTokenString(token)
			userID := util.ParseToken(t)
			c.JSON(http.StatusOK, &user.UserLoginResp{
				StatusCode: &errno.Success.ErrCode,
				StatusMsg:  &errno.Success.ErrMsg,
				UserId:     &userID,
				Token:      &token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"status_code": code,
				"status_msg":  message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginReq user.UserLoginReq
			if err := c.BindAndValidate(&loginReq); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(*loginReq.Username) == 0 || len(*loginReq.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.CheckUser(context.Background(), &user_rpc.RPCCheckUserReq{Username: *loginReq.Username, Password: *loginReq.Password})
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return int64(claims[constants.IdentityKey].(float64))
		},
	})
	if err != nil {
		panic(err)
	}
}
