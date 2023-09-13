package middleware

import (
	"context"
	"net/http"

	"github.com/oceanengine/ad_open_sdk_go/api"
	"github.com/oceanengine/ad_open_sdk_go/config"
)

func AuthMiddleware(next api.Endpoint) api.Endpoint {
	return func(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
		token := ctx.Value(config.ContextAccessToken)
		if token != nil {
			if tokenStr, ok := token.(string); ok {
				req.Header.Set("Access-Token", tokenStr)
			}
		}
		return next(ctx, req)
	}
}
