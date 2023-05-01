package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func CompatibleTokenMiddleWare() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.PostForm("token")
		if len(token) > 0 {
			AuthMiddleware.MiddlewareFunc()(ctx, c)
		} else {
			token := c.Query("token")
			if len(token) > 0 {
				AuthMiddleware.MiddlewareFunc()(ctx, c)
			}
		}
	}
}
