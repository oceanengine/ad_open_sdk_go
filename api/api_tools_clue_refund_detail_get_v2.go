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

// ToolsClueRefundDetailGetV2ApiService ToolsClueRefundDetailGetV2Api service
type ToolsClueRefundDetailGetV2ApiService service

type ApiOpenApi2ToolsClueRefundDetailGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueRefundDetailGetV2ApiService
	advertiserId *int64
	month        *string
	page         *int64
	pageSize     *int64
}

// 广告主id
func (r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueRefundDetailGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 查询账单月份，格式YYYYMM，只支持查询202112月及之后的账单
func (r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) Month(month string) *ApiOpenApi2ToolsClueRefundDetailGetGetRequest {
	r.month = &month
	return r
}

// 页码
func (r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) Page(page int64) *ApiOpenApi2ToolsClueRefundDetailGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，即每页展示的数据量
func (r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsClueRefundDetailGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) Execute() (*ToolsClueRefundDetailGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueRefundDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueRefundDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueRefundDetailGetGet Method for OpenApi2ToolsClueRefundDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueRefundDetailGetGetRequest
*/
func (a *ToolsClueRefundDetailGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueRefundDetailGetGetRequest {
	return &ApiOpenApi2ToolsClueRefundDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueRefundDetailGetV2Response
func (a *ToolsClueRefundDetailGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueRefundDetailGetGetRequest) (*ToolsClueRefundDetailGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueRefundDetailGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/refund_detail/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if r.month == nil {
		return localVarReturnValue, nil, ReportError("month is required and must be specified")
	}
	if strlen(*r.month) < 6 {
		return localVarReturnValue, nil, ReportError("month must have at least 6 elements")
	}
	if strlen(*r.month) > 6 {
		return localVarReturnValue, nil, ReportError("month must have less than 6 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "month", r.month)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

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
