/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
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

// ReportVideoFrameGetV2ApiService ReportVideoFrameGetV2Api service
type ReportVideoFrameGetV2ApiService service

type ApiOpenApi2ReportVideoFrameGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportVideoFrameGetV2ApiService
	advertiserId *int64
	endDate      **string
	filtering    *ReportVideoFrameGetV2Filtering
	metrics      *[]*ReportVideoFrameGetV2Metrics
	startDate    **string
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) EndDate(endDate *string) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) Filtering(filtering ReportVideoFrameGetV2Filtering) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) Metrics(metrics []*ReportVideoFrameGetV2Metrics) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	r.metrics = &metrics
	return r
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) StartDate(startDate *string) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) Execute() (*ReportVideoFrameGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportVideoFrameGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportVideoFrameGetGet Method for OpenApi2ReportVideoFrameGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportVideoFrameGetGetRequest
*/
func (a *ReportVideoFrameGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportVideoFrameGetGetRequest {
	return &ApiOpenApi2ReportVideoFrameGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportVideoFrameGetV2Response
func (a *ReportVideoFrameGetV2ApiService) getExecute(r *ApiOpenApi2ReportVideoFrameGetGetRequest) (*ReportVideoFrameGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ReportVideoFrameGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/video/frame/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.metrics != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics", r.metrics)
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
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
