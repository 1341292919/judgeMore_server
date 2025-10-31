package auth

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"judgeMore_server/app/gateway/mw/jwt"
	"judgeMore_server/app/gateway/pack"
	"judgeMore_server/pkg/errno"
)

func Auth() []app.HandlerFunc {
	//为了有扩展性
	return append(make([]app.HandlerFunc, 0),
		DoubleTokenAuthFunc(),
	)
}

func DoubleTokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		Aerr := jwt.IsAccessTokenAvailable(ctx, c)
		if Aerr != nil {
			if errors.Is(Aerr, errno.AuthAccessExpired) {
				Rerr := jwt.IsRefreshTokenAvailable(ctx, c)
				if Rerr != nil {
					pack.SendFailResponse(c, errno.ConvertErr(Rerr))
					c.Abort()
					return
				}
				jwt.GenerateAccessToken(c)
			} else {
				pack.SendFailResponse(c, errno.ConvertErr(Aerr))
				c.Abort()
				return
			}
		}
		c.Next(ctx)
	}
}
