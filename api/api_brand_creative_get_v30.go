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

// BrandCreativeGetV30ApiService BrandCreativeGetV30Api service
type BrandCreativeGetV30ApiService service

type ApiOpenApiV30BrandCreativeGetGetRequest struct {
	ctx             context.Context
	ApiService      *BrandCreativeGetV30ApiService
	advertiserId    *int64
	page            *int64
	size            *int64
	adIds           *[]string
	campaignIds     *[]string
	creativeIds     *[]string
	creativeStatus  *BrandCreativeGetV30CreativeStatus
	createStartTime *string
	createEndTime   *string
	startTime       *string
	endTime         *string
}

// 广告主ID
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 分页参数 默认从1开始
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) Page(page int64) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.page = &page
	return r
}

// 分页参数 每页最大100个结果
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) Size(size int64) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.size = &size
	return r
}

// 计划ID  每次查询最多100个
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) AdIds(adIds []string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.adIds = &adIds
	return r
}

// 广告组ID 每次查询最多100个
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) CampaignIds(campaignIds []string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.campaignIds = &campaignIds
	return r
}

// 广告创意ID
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) CreativeIds(creativeIds []string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.creativeIds = &creativeIds
	return r
}

// 广告创意状态
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) CreativeStatus(creativeStatus BrandCreativeGetV30CreativeStatus) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.creativeStatus = &creativeStatus
	return r
}

// 创建起始时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) CreateStartTime(createStartTime string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.createStartTime = &createStartTime
	return r
}

// 创建截止时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) CreateEndTime(createEndTime string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.createEndTime = &createEndTime
	return r
}

// 起始投放时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) StartTime(startTime string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.startTime = &startTime
	return r
}

// 截止投放时间 格式2023-01-01 00:00:00
func (r *ApiOpenApiV30BrandCreativeGetGetRequest) EndTime(endTime string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV30BrandCreativeGetGetRequest) Execute() (*BrandCreativeGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30BrandCreativeGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandCreativeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandCreativeGetGetRequest) WithLog(enable bool) *ApiOpenApiV30BrandCreativeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandCreativeGetGet Method for OpenApiV30BrandCreativeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandCreativeGetGetRequest
*/
func (a *BrandCreativeGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30BrandCreativeGetGetRequest {
	return &ApiOpenApiV30BrandCreativeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandCreativeGetV30Response
func (a *BrandCreativeGetV30ApiService) getExecute(r *ApiOpenApiV30BrandCreativeGetGetRequest) (*BrandCreativeGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandCreativeGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/creative/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if r.size == nil {
		return localVarReturnValue, nil, ReportError("size is required and must be specified")
	}
	if *r.size > 100 {
		return localVarReturnValue, nil, ReportError("size must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.adIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	}
	if r.campaignIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_ids", r.campaignIds)
	}
	if r.creativeIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creative_ids", r.creativeIds)
	}
	if r.creativeStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creative_status", r.creativeStatus)
	}
	if r.createStartTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_start_time", r.createStartTime)
	}
	if r.createEndTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_end_time", r.createEndTime)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size)
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
