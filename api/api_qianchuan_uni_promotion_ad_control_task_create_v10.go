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

// QianchuanUniPromotionAdControlTaskCreateV10ApiService QianchuanUniPromotionAdControlTaskCreateV10Api service
type QianchuanUniPromotionAdControlTaskCreateV10ApiService service

type ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest struct {
	ctx                                                context.Context
	ApiService                                         *QianchuanUniPromotionAdControlTaskCreateV10ApiService
	qianchuanUniPromotionAdControlTaskCreateV10Request *QianchuanUniPromotionAdControlTaskCreateV10Request
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest) QianchuanUniPromotionAdControlTaskCreateV10Request(qianchuanUniPromotionAdControlTaskCreateV10Request QianchuanUniPromotionAdControlTaskCreateV10Request) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest {
	r.qianchuanUniPromotionAdControlTaskCreateV10Request = &qianchuanUniPromotionAdControlTaskCreateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest) Execute() (*QianchuanUniPromotionAdControlTaskCreateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanUniPromotionAdControlTaskCreatePost Method for OpenApiV10QianchuanUniPromotionAdControlTaskCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest
*/
func (a *QianchuanUniPromotionAdControlTaskCreateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest {
	return &ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanUniPromotionAdControlTaskCreateV10Response
func (a *QianchuanUniPromotionAdControlTaskCreateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskCreatePostRequest) (*QianchuanUniPromotionAdControlTaskCreateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanUniPromotionAdControlTaskCreateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/uni_promotion/ad/control_task/create/"

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
	localVarPostBody = r.qianchuanUniPromotionAdControlTaskCreateV10Request
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
