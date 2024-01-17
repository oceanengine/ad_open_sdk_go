/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
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

// SpiTaskGetV2ApiService SpiTaskGetV2Api service
type SpiTaskGetV2ApiService service

type ApiOpenApi2SpiTaskGetGetRequest struct {
	ctx          context.Context
	ApiService   *SpiTaskGetV2ApiService
	appId        *int64
	endDate      *string
	page         *int64
	pageSize     *int64
	serviceLabel *string
	startDate    *string
	status       *SpiTaskGetV2Status
	subscribeId  *int64
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) AppId(appId int64) *ApiOpenApi2SpiTaskGetGetRequest {
	r.appId = &appId
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) EndDate(endDate string) *ApiOpenApi2SpiTaskGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) Page(page int64) *ApiOpenApi2SpiTaskGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2SpiTaskGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) ServiceLabel(serviceLabel string) *ApiOpenApi2SpiTaskGetGetRequest {
	r.serviceLabel = &serviceLabel
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) StartDate(startDate string) *ApiOpenApi2SpiTaskGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) Status(status SpiTaskGetV2Status) *ApiOpenApi2SpiTaskGetGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) SubscribeId(subscribeId int64) *ApiOpenApi2SpiTaskGetGetRequest {
	r.subscribeId = &subscribeId
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) Execute() (*SpiTaskGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2SpiTaskGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2SpiTaskGetGetRequest) WithLog(enable bool) *ApiOpenApi2SpiTaskGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2SpiTaskGetGet Method for OpenApi2SpiTaskGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2SpiTaskGetGetRequest
*/
func (a *SpiTaskGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2SpiTaskGetGetRequest {
	return &ApiOpenApi2SpiTaskGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SpiTaskGetV2Response
func (a *SpiTaskGetV2ApiService) getExecute(r *ApiOpenApi2SpiTaskGetGetRequest) (*SpiTaskGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *SpiTaskGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/spi_task/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.appId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_id", r.appId)
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.serviceLabel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "service_label", r.serviceLabel)
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	if r.subscribeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subscribe_id", r.subscribeId)
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
