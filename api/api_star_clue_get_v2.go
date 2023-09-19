/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
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

// StarClueGetV2ApiService StarClueGetV2Api service
type StarClueGetV2ApiService service

type ApiOpenApi2StarClueGetGetRequest struct {
	ctx        context.Context
	ApiService *StarClueGetV2ApiService
	demandId   *int64
	orderId    *int64
	page       *int64
	pageSize   *int64
	starId     *int64
}

func (r *ApiOpenApi2StarClueGetGetRequest) DemandId(demandId int64) *ApiOpenApi2StarClueGetGetRequest {
	r.demandId = &demandId
	return r
}

func (r *ApiOpenApi2StarClueGetGetRequest) OrderId(orderId int64) *ApiOpenApi2StarClueGetGetRequest {
	r.orderId = &orderId
	return r
}

func (r *ApiOpenApi2StarClueGetGetRequest) Page(page int64) *ApiOpenApi2StarClueGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2StarClueGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2StarClueGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2StarClueGetGetRequest) StarId(starId int64) *ApiOpenApi2StarClueGetGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarClueGetGetRequest) Execute() (*StarClueGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarClueGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarClueGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarClueGetGetRequest) WithLog(enable bool) *ApiOpenApi2StarClueGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarClueGetGet Method for OpenApi2StarClueGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarClueGetGetRequest
*/
func (a *StarClueGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarClueGetGetRequest {
	return &ApiOpenApi2StarClueGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarClueGetV2Response
func (a *StarClueGetV2ApiService) getExecute(r *ApiOpenApi2StarClueGetGetRequest) (*StarClueGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *StarClueGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/clue/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.demandId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	}
	if r.orderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.starId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
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
