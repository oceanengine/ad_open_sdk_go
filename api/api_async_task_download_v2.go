/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

// AsyncTaskDownloadV2ApiService AsyncTaskDownloadV2Api service
type AsyncTaskDownloadV2ApiService service

type ApiOpenApi2AsyncTaskDownloadGetRequest struct {
	ctx          context.Context
	ApiService   *AsyncTaskDownloadV2ApiService
	advertiserId *int64
	taskId       *int64
	rangeFrom    *int64
	rangeTo      *int64
}

func (r *ApiOpenApi2AsyncTaskDownloadGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AsyncTaskDownloadGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2AsyncTaskDownloadGetRequest) TaskId(taskId int64) *ApiOpenApi2AsyncTaskDownloadGetRequest {
	r.taskId = &taskId
	return r
}

func (r *ApiOpenApi2AsyncTaskDownloadGetRequest) RangeFrom(rangeFrom int64) *ApiOpenApi2AsyncTaskDownloadGetRequest {
	r.rangeFrom = &rangeFrom
	return r
}

func (r *ApiOpenApi2AsyncTaskDownloadGetRequest) RangeTo(rangeTo int64) *ApiOpenApi2AsyncTaskDownloadGetRequest {
	r.rangeTo = &rangeTo
	return r
}

func (r *ApiOpenApi2AsyncTaskDownloadGetRequest) Execute() (*AsyncTaskDownloadV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AsyncTaskDownloadGetRequest) AccessToken(accessToken string) *ApiOpenApi2AsyncTaskDownloadGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AsyncTaskDownloadGetRequest) WithLog(enable bool) *ApiOpenApi2AsyncTaskDownloadGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AsyncTaskDownloadGet Method for OpenApi2AsyncTaskDownloadGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AsyncTaskDownloadGetRequest
*/
func (a *AsyncTaskDownloadV2ApiService) Get(ctx context.Context) *ApiOpenApi2AsyncTaskDownloadGetRequest {
	return &ApiOpenApi2AsyncTaskDownloadGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AsyncTaskDownloadV2Response
func (a *AsyncTaskDownloadV2ApiService) getExecute(r *ApiOpenApi2AsyncTaskDownloadGetRequest) (*AsyncTaskDownloadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AsyncTaskDownloadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/async_task/download/"

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
	if r.taskId == nil {
		return localVarReturnValue, nil, ReportError("taskId is required and must be specified")
	}
	if *r.taskId < 1 {
		return localVarReturnValue, nil, ReportError("taskId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "task_id", r.taskId)
	if r.rangeFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range_from", r.rangeFrom)
	}
	if r.rangeTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "range_to", r.rangeTo)
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
