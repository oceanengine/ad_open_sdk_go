/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

// ReportBrandCampaignGetV30ApiService ReportBrandCampaignGetV30Api service
type ReportBrandCampaignGetV30ApiService service

type ApiOpenApiV30ReportBrandCampaignGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportBrandCampaignGetV30ApiService
	advertiserId *int64
	page         *int64
	size         *int64
	landingType  *ReportBrandCampaignGetV30LandingType
	pricingType  *ReportBrandCampaignGetV30PricingType
	campaignIds  *[]string
	startTime    *string
	endTime      *string
}

// 广告主ID
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 分页参数 默认从1开始
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) Page(page int64) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.page = &page
	return r
}

// 分页参数 每页最大值为100
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) Size(size int64) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.size = &size
	return r
}

// 推广目的
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) LandingType(landingType ReportBrandCampaignGetV30LandingType) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.landingType = &landingType
	return r
}

// 计费类型
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) PricingType(pricingType ReportBrandCampaignGetV30PricingType) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.pricingType = &pricingType
	return r
}

// 广告组ID 每次请求上线为100
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) CampaignIds(campaignIds []string) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.campaignIds = &campaignIds
	return r
}

// 开始时间 格式2023-01-01
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) StartTime(startTime string) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.startTime = &startTime
	return r
}

// 结束时间 格式2023-01-01
func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) EndTime(endTime string) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) Execute() (*ReportBrandCampaignGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportBrandCampaignGetGet Method for OpenApiV30ReportBrandCampaignGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportBrandCampaignGetGetRequest
*/
func (a *ReportBrandCampaignGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportBrandCampaignGetGetRequest {
	return &ApiOpenApiV30ReportBrandCampaignGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportBrandCampaignGetV30Response
func (a *ReportBrandCampaignGetV30ApiService) getExecute(r *ApiOpenApiV30ReportBrandCampaignGetGetRequest) (*ReportBrandCampaignGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportBrandCampaignGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/brand/campaign/get/"

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
	if r.landingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	}
	if r.pricingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pricing_type", r.pricingType)
	}
	if r.campaignIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_ids", r.campaignIds)
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
