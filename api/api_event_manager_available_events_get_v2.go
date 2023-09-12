/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

// EventManagerAvailableEventsGetV2ApiService EventManagerAvailableEventsGetV2Api service
type EventManagerAvailableEventsGetV2ApiService service

type ApiOpenApi2EventManagerAvailableEventsGetGetRequest struct {
	ctx          context.Context
	ApiService   *EventManagerAvailableEventsGetV2ApiService
	advertiserId *int64
	assetId      *int64
}

// 广告主ID
func (r *ApiOpenApi2EventManagerAvailableEventsGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2EventManagerAvailableEventsGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 资产ID
func (r *ApiOpenApi2EventManagerAvailableEventsGetGetRequest) AssetId(assetId int64) *ApiOpenApi2EventManagerAvailableEventsGetGetRequest {
	r.assetId = &assetId
	return r
}

func (r *ApiOpenApi2EventManagerAvailableEventsGetGetRequest) Execute() (*EventManagerAvailableEventsGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2EventManagerAvailableEventsGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2EventManagerAvailableEventsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2EventManagerAvailableEventsGetGetRequest) WithLog(enable bool) *ApiOpenApi2EventManagerAvailableEventsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2EventManagerAvailableEventsGetGet Method for OpenApi2EventManagerAvailableEventsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2EventManagerAvailableEventsGetGetRequest
*/
func (a *EventManagerAvailableEventsGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2EventManagerAvailableEventsGetGetRequest {
	return &ApiOpenApi2EventManagerAvailableEventsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EventManagerAvailableEventsGetV2Response
func (a *EventManagerAvailableEventsGetV2ApiService) getExecute(r *ApiOpenApi2EventManagerAvailableEventsGetGetRequest) (*EventManagerAvailableEventsGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *EventManagerAvailableEventsGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/event_manager/available_events/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if *r.advertiserId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than 9223372036854775807")
	}
	if r.assetId == nil {
		return localVarReturnValue, nil, ReportError("assetId is required and must be specified")
	}
	if *r.assetId < 1 {
		return localVarReturnValue, nil, ReportError("assetId must be greater than 1")
	}
	if *r.assetId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("assetId must be less than 9223372036854775807")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId)
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
