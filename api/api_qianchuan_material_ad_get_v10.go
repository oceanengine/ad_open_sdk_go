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

// QianchuanMaterialAdGetV10ApiService QianchuanMaterialAdGetV10Api service
type QianchuanMaterialAdGetV10ApiService service

type ApiOpenApiV10QianchuanMaterialAdGetGetRequest struct {
	ctx            context.Context
	ApiService     *QianchuanMaterialAdGetV10ApiService
	advertiserId   *int64
	materialId     *int64
	materialType   *QianchuanMaterialAdGetV10MaterialType
	marketingScene *QianchuanMaterialAdGetV10MarketingScene
	marketingGoal  *QianchuanMaterialAdGetV10MarketingGoal
	startTime      *string
	endTime        *string
	fields         *[]string
	orderField     *string
	orderType      *QianchuanMaterialAdGetV10OrderType
	page           *int64
	pageSize       *QianchuanMaterialAdGetV10PageSize
}

// 广告主id
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 素材id
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) MaterialId(materialId int64) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.materialId = &materialId
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) MaterialType(materialType QianchuanMaterialAdGetV10MaterialType) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.materialType = &materialType
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) MarketingScene(marketingScene QianchuanMaterialAdGetV10MarketingScene) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.marketingScene = &marketingScene
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) MarketingGoal(marketingGoal QianchuanMaterialAdGetV10MarketingGoal) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

// 查询起始日期
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.startTime = &startTime
	return r
}

// 查询结束日期
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.endTime = &endTime
	return r
}

// 查询指标
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) Fields(fields []string) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.fields = &fields
	return r
}

// 排序字段
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) OrderField(orderField string) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.orderField = &orderField
	return r
}

// 排序方式
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) OrderType(orderType QianchuanMaterialAdGetV10OrderType) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.orderType = &orderType
	return r
}

// 页码
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) Page(page int64) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) PageSize(pageSize QianchuanMaterialAdGetV10PageSize) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) Execute() (*QianchuanMaterialAdGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanMaterialAdGetGet Method for OpenApiV10QianchuanMaterialAdGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanMaterialAdGetGetRequest
*/
func (a *QianchuanMaterialAdGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanMaterialAdGetGetRequest {
	return &ApiOpenApiV10QianchuanMaterialAdGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanMaterialAdGetV10Response
func (a *QianchuanMaterialAdGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanMaterialAdGetGetRequest) (*QianchuanMaterialAdGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanMaterialAdGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/material/ad/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.materialId == nil {
		return localVarReturnValue, nil, ReportError("materialId is required and must be specified")
	}
	if r.materialType == nil {
		return localVarReturnValue, nil, ReportError("materialType is required and must be specified")
	}
	if r.marketingScene == nil {
		return localVarReturnValue, nil, ReportError("marketingScene is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}
	if len(*r.fields) < 1 {
		return localVarReturnValue, nil, ReportError("fields must have at least 1 elements")
	}
	if len(*r.fields) > 20 {
		return localVarReturnValue, nil, ReportError("fields must have less than 20 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_id", r.materialId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_type", r.materialType)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_scene", r.marketingScene)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
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
