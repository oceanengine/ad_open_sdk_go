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

// LocalChargeListV30ApiService LocalChargeListV30Api service
type LocalChargeListV30ApiService service

type ApiOpenApiV30LocalChargeListGetRequest struct {
	ctx            context.Context
	ApiService     *LocalChargeListV30ApiService
	localAccountId *int64
	startTime      *string
	endTime        *string
	chargeMethods  *[]*LocalChargeListV30ChargeMethods
	page           *int32
	pageSize       *int32
}

// 本地推广告账户ID
func (r *ApiOpenApiV30LocalChargeListGetRequest) LocalAccountId(localAccountId int64) *ApiOpenApiV30LocalChargeListGetRequest {
	r.localAccountId = &localAccountId
	return r
}

// 充值起始时间，闭区间，格式：yyyy-MM-dd HH:mm:ss
func (r *ApiOpenApiV30LocalChargeListGetRequest) StartTime(startTime string) *ApiOpenApiV30LocalChargeListGetRequest {
	r.startTime = &startTime
	return r
}

// 充值起始时间，开区间，格式：yyyy-MM-dd HH:mm:ss
func (r *ApiOpenApiV30LocalChargeListGetRequest) EndTime(endTime string) *ApiOpenApiV30LocalChargeListGetRequest {
	r.endTime = &endTime
	return r
}

// 充值方式，多选
func (r *ApiOpenApiV30LocalChargeListGetRequest) ChargeMethods(chargeMethods []*LocalChargeListV30ChargeMethods) *ApiOpenApiV30LocalChargeListGetRequest {
	r.chargeMethods = &chargeMethods
	return r
}

// 页码
func (r *ApiOpenApiV30LocalChargeListGetRequest) Page(page int32) *ApiOpenApiV30LocalChargeListGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30LocalChargeListGetRequest) PageSize(pageSize int32) *ApiOpenApiV30LocalChargeListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30LocalChargeListGetRequest) Execute() (*LocalChargeListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30LocalChargeListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalChargeListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalChargeListGetRequest) WithLog(enable bool) *ApiOpenApiV30LocalChargeListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalChargeListGet Method for OpenApiV30LocalChargeListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalChargeListGetRequest
*/
func (a *LocalChargeListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30LocalChargeListGetRequest {
	return &ApiOpenApiV30LocalChargeListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalChargeListV30Response
func (a *LocalChargeListV30ApiService) getExecute(r *ApiOpenApiV30LocalChargeListGetRequest) (*LocalChargeListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalChargeListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/charge/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.localAccountId == nil {
		return localVarReturnValue, nil, ReportError("localAccountId is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "local_account_id", r.localAccountId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	if r.chargeMethods != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "charge_methods", r.chargeMethods)
	}
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
