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

// QianchuanAccountBudgetUpdateV10ApiService QianchuanAccountBudgetUpdateV10Api service
type QianchuanAccountBudgetUpdateV10ApiService service

type ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest struct {
	ctx                                    context.Context
	ApiService                             *QianchuanAccountBudgetUpdateV10ApiService
	qianchuanAccountBudgetUpdateV10Request *QianchuanAccountBudgetUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest) QianchuanAccountBudgetUpdateV10Request(qianchuanAccountBudgetUpdateV10Request QianchuanAccountBudgetUpdateV10Request) *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest {
	r.qianchuanAccountBudgetUpdateV10Request = &qianchuanAccountBudgetUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest) Execute() (*QianchuanAccountBudgetUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAccountBudgetUpdatePost Method for OpenApiV10QianchuanAccountBudgetUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest
*/
func (a *QianchuanAccountBudgetUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest {
	return &ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAccountBudgetUpdateV10Response
func (a *QianchuanAccountBudgetUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanAccountBudgetUpdatePostRequest) (*QianchuanAccountBudgetUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAccountBudgetUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/account/budget/update/"

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
	localVarPostBody = r.qianchuanAccountBudgetUpdateV10Request
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
