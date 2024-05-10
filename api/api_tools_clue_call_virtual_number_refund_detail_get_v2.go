/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// ToolsClueCallVirtualNumberRefundDetailGetV2ApiService ToolsClueCallVirtualNumberRefundDetailGetV2Api service
type ToolsClueCallVirtualNumberRefundDetailGetV2ApiService service

type ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueCallVirtualNumberRefundDetailGetV2ApiService
	advertiserId *int64
	month        *string
	page         *int64
	pageSize     *int64
}

// 广告主id
func (r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 查询账单月份，格式YYYYMM，只支持查询202112月及之后的账单
func (r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) Month(month string) *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest {
	r.month = &month
	return r
}

// 页码
func (r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) Page(page int64) *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，即每页展示的数据量
func (r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) Execute() (*ToolsClueCallVirtualNumberRefundDetailGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueCallVirtualNumberRefundDetailGetGet Method for OpenApi2ToolsClueCallVirtualNumberRefundDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest
*/
func (a *ToolsClueCallVirtualNumberRefundDetailGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest {
	return &ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueCallVirtualNumberRefundDetailGetV2Response
func (a *ToolsClueCallVirtualNumberRefundDetailGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueCallVirtualNumberRefundDetailGetGetRequest) (*ToolsClueCallVirtualNumberRefundDetailGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueCallVirtualNumberRefundDetailGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/call_virtual_number/refund_detail/get/"

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
