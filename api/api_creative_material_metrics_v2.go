/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

// CreativeMaterialMetricsV2ApiService CreativeMaterialMetricsV2Api service
type CreativeMaterialMetricsV2ApiService service

type ApiOpenApi2CreativeMaterialMetricsGetRequest struct {
	ctx          context.Context
	ApiService   *CreativeMaterialMetricsV2ApiService
	advertiserId *int64
	materialIds  *[]int64
	startDate    *string
	endDate      *string
}

// 广告主 ID
func (r *ApiOpenApi2CreativeMaterialMetricsGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2CreativeMaterialMetricsGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 素材 ID
func (r *ApiOpenApi2CreativeMaterialMetricsGetRequest) MaterialIds(materialIds []int64) *ApiOpenApi2CreativeMaterialMetricsGetRequest {
	r.materialIds = &materialIds
	return r
}

// 统计开始日期，格式 yyyy-mm-dd
func (r *ApiOpenApi2CreativeMaterialMetricsGetRequest) StartDate(startDate string) *ApiOpenApi2CreativeMaterialMetricsGetRequest {
	r.startDate = &startDate
	return r
}

// 统计结束时间，格式 yyyy-mm-dd
func (r *ApiOpenApi2CreativeMaterialMetricsGetRequest) EndDate(endDate string) *ApiOpenApi2CreativeMaterialMetricsGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2CreativeMaterialMetricsGetRequest) Execute() (*CreativeMaterialMetricsV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2CreativeMaterialMetricsGetRequest) AccessToken(accessToken string) *ApiOpenApi2CreativeMaterialMetricsGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2CreativeMaterialMetricsGetRequest) WithLog(enable bool) *ApiOpenApi2CreativeMaterialMetricsGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2CreativeMaterialMetricsGet Method for OpenApi2CreativeMaterialMetricsGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2CreativeMaterialMetricsGetRequest
*/
func (a *CreativeMaterialMetricsV2ApiService) Get(ctx context.Context) *ApiOpenApi2CreativeMaterialMetricsGetRequest {
	return &ApiOpenApi2CreativeMaterialMetricsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreativeMaterialMetricsV2Response
func (a *CreativeMaterialMetricsV2ApiService) getExecute(r *ApiOpenApi2CreativeMaterialMetricsGetRequest) (*CreativeMaterialMetricsV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CreativeMaterialMetricsV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/creative/material/metrics/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.materialIds == nil {
		return localVarReturnValue, nil, ReportError("materialIds is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_ids", r.materialIds)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
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
