package middleware

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"

	"code.byted.org/ad/ad_open_sdk_go/api"
	"code.byted.org/ad/ad_open_sdk_go/config"
)

func LogMiddleware(next api.Endpoint) api.Endpoint {
	return func(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
		enableLog := ctx.Value(config.ContextEnableLog)
		if enableLog != nil {
			reqLog, _ := httputil.DumpRequest(req, true)
			log.Println(string(reqLog))
		}
		resp, err = next(ctx, req)
		if err != nil {
			return
		}
		if enableLog != nil {
			respLog, _ := httputil.DumpResponse(resp, true)
			log.Println(string(respLog))
		}
		return
	}
}
