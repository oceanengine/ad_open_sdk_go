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

// ClueSmartphoneRecordV2ApiService ClueSmartphoneRecordV2Api service
type ClueSmartphoneRecordV2ApiService service

type ApiOpenApi2ClueSmartphoneRecordGetRequest struct {
	ctx          context.Context
	ApiService   *ClueSmartphoneRecordV2ApiService
	advertiserId *int64
	instanceIds  *[]int64
	clueIds      *[]int64
	siteIds      *[]int64
	adIds        *[]int64
	startTime    *string
	endTime      *string
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) InstanceIds(instanceIds []int64) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.instanceIds = &instanceIds
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) ClueIds(clueIds []int64) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.clueIds = &clueIds
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) SiteIds(siteIds []int64) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.siteIds = &siteIds
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) AdIds(adIds []int64) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.adIds = &adIds
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) StartTime(startTime string) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) EndTime(endTime string) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) Page(page int32) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) PageSize(pageSize int32) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) Execute() (*ClueSmartphoneRecordV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) AccessToken(accessToken string) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueSmartphoneRecordGetRequest) WithLog(enable bool) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueSmartphoneRecordGet Method for OpenApi2ClueSmartphoneRecordGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueSmartphoneRecordGetRequest
*/
func (a *ClueSmartphoneRecordV2ApiService) Get(ctx context.Context) *ApiOpenApi2ClueSmartphoneRecordGetRequest {
	return &ApiOpenApi2ClueSmartphoneRecordGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueSmartphoneRecordV2Response
func (a *ClueSmartphoneRecordV2ApiService) getExecute(r *ApiOpenApi2ClueSmartphoneRecordGetRequest) (*ClueSmartphoneRecordV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ClueSmartphoneRecordV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/smartphone/record/"

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
	if *r.advertiserId > -9223372036854775616 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than -9223372036854775616")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.instanceIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "instance_ids", r.instanceIds)
	}
	if r.clueIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clue_ids", r.clueIds)
	}
	if r.siteIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "site_ids", r.siteIds)
	}
	if r.adIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_ids", r.adIds)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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
