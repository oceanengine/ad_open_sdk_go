package api

import (
	"context"
	"net/http"
)

// Endpoint represent one method for calling from remote.
type Endpoint func(ctx context.Context, req *http.Request) (resp *http.Response, err error)

// Middleware deal with input Endpoint and output Endpoint.
type Middleware func(Endpoint) Endpoint

// chain connect middlewares into one middleware.
func chain(mws ...Middleware) Middleware {
	return func(next Endpoint) Endpoint {
		for i := len(mws) - 1; i >= 0; i-- {
			next = mws[i](next)
		}
		return next
	}
}
