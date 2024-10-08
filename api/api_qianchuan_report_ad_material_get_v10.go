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

// QianchuanReportAdMaterialGetV10ApiService QianchuanReportAdMaterialGetV10Api service
type QianchuanReportAdMaterialGetV10ApiService service

type ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanReportAdMaterialGetV10ApiService
	advertiserId *int64
	adId         *int64
	startDate    *string
	endDate      *string
	fields       *[]string
	filtering    *QianchuanReportAdMaterialGetV10Filtering
	orderType    *QianchuanReportAdMaterialGetV10OrderType
	orderField   *string
}

func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) AdId(adId int64) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.adId = &adId
	return r
}

// 查询起始日期 格式 yyyy-mm-dd
func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) StartDate(startDate string) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.startDate = &startDate
	return r
}

// 查询结束日期 格式 yyyy-mm-dd
func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) EndDate(endDate string) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.endDate = &endDate
	return r
}

// 必填 需要查询的消耗指标
func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) Filtering(filtering QianchuanReportAdMaterialGetV10Filtering) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.filtering = &filtering
	return r
}

// 排序方式
func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) OrderType(orderType QianchuanReportAdMaterialGetV10OrderType) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.orderType = &orderType
	return r
}

// 排序字段
func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) Execute() (*QianchuanReportAdMaterialGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanReportAdMaterialGetGet Method for OpenApiV10QianchuanReportAdMaterialGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest
*/
func (a *QianchuanReportAdMaterialGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest {
	return &ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanReportAdMaterialGetV10Response
func (a *QianchuanReportAdMaterialGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanReportAdMaterialGetGetRequest) (*QianchuanReportAdMaterialGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanReportAdMaterialGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/report/ad/material/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.adId == nil {
		return localVarReturnValue, nil, ReportError("adId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}
	if len(*r.fields) < 1 {
		return localVarReturnValue, nil, ReportError("fields must have at least 1 elements")
	}
	if len(*r.fields) > 150 {
		return localVarReturnValue, nil, ReportError("fields must have less than 150 elements")
	}
	if r.filtering == nil {
		return localVarReturnValue, nil, ReportError("filtering is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
