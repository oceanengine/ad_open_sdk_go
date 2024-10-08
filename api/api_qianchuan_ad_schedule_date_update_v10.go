/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// QianchuanAdScheduleDateUpdateV10ApiService QianchuanAdScheduleDateUpdateV10Api service
type QianchuanAdScheduleDateUpdateV10ApiService service

type ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest struct {
	ctx                                     context.Context
	ApiService                              *QianchuanAdScheduleDateUpdateV10ApiService
	qianchuanAdScheduleDateUpdateV10Request *QianchuanAdScheduleDateUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest) QianchuanAdScheduleDateUpdateV10Request(qianchuanAdScheduleDateUpdateV10Request QianchuanAdScheduleDateUpdateV10Request) *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest {
	r.qianchuanAdScheduleDateUpdateV10Request = &qianchuanAdScheduleDateUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest) Execute() (*QianchuanAdScheduleDateUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAdScheduleDateUpdatePost Method for OpenApiV10QianchuanAdScheduleDateUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest
*/
func (a *QianchuanAdScheduleDateUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest {
	return &ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAdScheduleDateUpdateV10Response
func (a *QianchuanAdScheduleDateUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanAdScheduleDateUpdatePostRequest) (*QianchuanAdScheduleDateUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAdScheduleDateUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/ad/schedule_date/update/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// body params
	localVarPostBody = r.qianchuanAdScheduleDateUpdateV10Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
