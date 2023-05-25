package middleware

import (
	"context"
	"net/http"
	"runtime"

	"github.com/oceanengine/ad_open_sdk_go/api"
	"github.com/oceanengine/ad_open_sdk_go/config"
)

func HeaderMiddleware(next api.Endpoint) api.Endpoint {
	return func(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
		req.Header.Set("X-Sdk-Language", "go")
		req.Header.Set("X-Sdk-Language-Version", runtime.Version())
		req.Header.Set("X-Sdk-Version", config.Version)
		return next(ctx, req)
	}
}
