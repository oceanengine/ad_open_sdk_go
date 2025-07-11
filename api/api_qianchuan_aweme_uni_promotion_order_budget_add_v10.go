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

// QianchuanAwemeUniPromotionOrderBudgetAddV10ApiService QianchuanAwemeUniPromotionOrderBudgetAddV10Api service
type QianchuanAwemeUniPromotionOrderBudgetAddV10ApiService service

type ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest struct {
	ctx                                                context.Context
	ApiService                                         *QianchuanAwemeUniPromotionOrderBudgetAddV10ApiService
	qianchuanAwemeUniPromotionOrderBudgetAddV10Request *QianchuanAwemeUniPromotionOrderBudgetAddV10Request
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest) QianchuanAwemeUniPromotionOrderBudgetAddV10Request(qianchuanAwemeUniPromotionOrderBudgetAddV10Request QianchuanAwemeUniPromotionOrderBudgetAddV10Request) *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest {
	r.qianchuanAwemeUniPromotionOrderBudgetAddV10Request = &qianchuanAwemeUniPromotionOrderBudgetAddV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest) Execute() (*QianchuanAwemeUniPromotionOrderBudgetAddV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPost Method for OpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest
*/
func (a *QianchuanAwemeUniPromotionOrderBudgetAddV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest {
	return &ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeUniPromotionOrderBudgetAddV10Response
func (a *QianchuanAwemeUniPromotionOrderBudgetAddV10ApiService) postExecute(r *ApiOpenApiV10QianchuanAwemeUniPromotionOrderBudgetAddPostRequest) (*QianchuanAwemeUniPromotionOrderBudgetAddV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeUniPromotionOrderBudgetAddV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/uni_promotion/order/budget/add/"

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
	localVarPostBody = r.qianchuanAwemeUniPromotionOrderBudgetAddV10Request
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
