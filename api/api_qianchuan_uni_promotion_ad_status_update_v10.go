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

// QianchuanUniPromotionAdStatusUpdateV10ApiService QianchuanUniPromotionAdStatusUpdateV10Api service
type QianchuanUniPromotionAdStatusUpdateV10ApiService service

type ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest struct {
	ctx                                           context.Context
	ApiService                                    *QianchuanUniPromotionAdStatusUpdateV10ApiService
	qianchuanUniPromotionAdStatusUpdateV10Request *QianchuanUniPromotionAdStatusUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest) QianchuanUniPromotionAdStatusUpdateV10Request(qianchuanUniPromotionAdStatusUpdateV10Request QianchuanUniPromotionAdStatusUpdateV10Request) *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest {
	r.qianchuanUniPromotionAdStatusUpdateV10Request = &qianchuanUniPromotionAdStatusUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest) Execute() (*QianchuanUniPromotionAdStatusUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanUniPromotionAdStatusUpdatePost Method for OpenApiV10QianchuanUniPromotionAdStatusUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest
*/
func (a *QianchuanUniPromotionAdStatusUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest {
	return &ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanUniPromotionAdStatusUpdateV10Response
func (a *QianchuanUniPromotionAdStatusUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanUniPromotionAdStatusUpdatePostRequest) (*QianchuanUniPromotionAdStatusUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanUniPromotionAdStatusUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/uni_promotion/ad/status/update/"

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
	localVarPostBody = r.qianchuanUniPromotionAdStatusUpdateV10Request
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
