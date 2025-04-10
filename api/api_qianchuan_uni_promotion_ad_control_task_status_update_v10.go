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

// QianchuanUniPromotionAdControlTaskStatusUpdateV10ApiService QianchuanUniPromotionAdControlTaskStatusUpdateV10Api service
type QianchuanUniPromotionAdControlTaskStatusUpdateV10ApiService service

type ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest struct {
	ctx                                                      context.Context
	ApiService                                               *QianchuanUniPromotionAdControlTaskStatusUpdateV10ApiService
	qianchuanUniPromotionAdControlTaskStatusUpdateV10Request *QianchuanUniPromotionAdControlTaskStatusUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest) QianchuanUniPromotionAdControlTaskStatusUpdateV10Request(qianchuanUniPromotionAdControlTaskStatusUpdateV10Request QianchuanUniPromotionAdControlTaskStatusUpdateV10Request) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest {
	r.qianchuanUniPromotionAdControlTaskStatusUpdateV10Request = &qianchuanUniPromotionAdControlTaskStatusUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest) Execute() (*QianchuanUniPromotionAdControlTaskStatusUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePost Method for OpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest
*/
func (a *QianchuanUniPromotionAdControlTaskStatusUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest {
	return &ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanUniPromotionAdControlTaskStatusUpdateV10Response
func (a *QianchuanUniPromotionAdControlTaskStatusUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanUniPromotionAdControlTaskStatusUpdatePostRequest) (*QianchuanUniPromotionAdControlTaskStatusUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanUniPromotionAdControlTaskStatusUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/uni_promotion/ad/control_task/status/update/"

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
	localVarPostBody = r.qianchuanUniPromotionAdControlTaskStatusUpdateV10Request
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
