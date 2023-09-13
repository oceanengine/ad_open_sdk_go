/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
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

// ReportLiveRoomAttributeGetV2ApiService ReportLiveRoomAttributeGetV2Api service
type ReportLiveRoomAttributeGetV2ApiService service

type ApiOpenApi2ReportLiveRoomAttributeGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportLiveRoomAttributeGetV2ApiService
	advertiserId *int64
	endTime      **string
	filtering    *ReportLiveRoomAttributeGetV2Filtering
	page         *int64
	pageSize     *int64
	startTime    **string
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) EndTime(endTime *string) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) Filtering(filtering ReportLiveRoomAttributeGetV2Filtering) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) Page(page int64) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) StartTime(startTime *string) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) Execute() (*ReportLiveRoomAttributeGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportLiveRoomAttributeGetGet Method for OpenApi2ReportLiveRoomAttributeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportLiveRoomAttributeGetGetRequest
*/
func (a *ReportLiveRoomAttributeGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest {
	return &ApiOpenApi2ReportLiveRoomAttributeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportLiveRoomAttributeGetV2Response
func (a *ReportLiveRoomAttributeGetV2ApiService) getExecute(r *ApiOpenApi2ReportLiveRoomAttributeGetGetRequest) (*ReportLiveRoomAttributeGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ReportLiveRoomAttributeGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/live_room/attribute/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
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
