/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
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

// AdvertiserBudgetGetV2ApiService AdvertiserBudgetGetV2Api service
type AdvertiserBudgetGetV2ApiService service

type ApiOpenApi2AdvertiserBudgetGetGetRequest struct {
	ctx           context.Context
	ApiService    *AdvertiserBudgetGetV2ApiService
	advertiserIds *[]int64
}

func (r *ApiOpenApi2AdvertiserBudgetGetGetRequest) AdvertiserIds(advertiserIds []int64) *ApiOpenApi2AdvertiserBudgetGetGetRequest {
	r.advertiserIds = &advertiserIds
	return r
}

func (r *ApiOpenApi2AdvertiserBudgetGetGetRequest) Execute() (*AdvertiserBudgetGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2AdvertiserBudgetGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2AdvertiserBudgetGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdvertiserBudgetGetGetRequest) WithLog(enable bool) *ApiOpenApi2AdvertiserBudgetGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdvertiserBudgetGetGet Method for OpenApi2AdvertiserBudgetGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdvertiserBudgetGetGetRequest
*/
func (a *AdvertiserBudgetGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2AdvertiserBudgetGetGetRequest {
	return &ApiOpenApi2AdvertiserBudgetGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserBudgetGetV2Response
func (a *AdvertiserBudgetGetV2ApiService) getExecute(r *ApiOpenApi2AdvertiserBudgetGetGetRequest) (*AdvertiserBudgetGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserBudgetGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/advertiser/budget/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserIds == nil {
		return localVarReturnValue, nil, ReportError("advertiserIds is required and must be specified")
	}
	if len(*r.advertiserIds) < 1 {
		return localVarReturnValue, nil, ReportError("advertiserIds must have at least 1 elements")
	}
	if len(*r.advertiserIds) > 100 {
		return localVarReturnValue, nil, ReportError("advertiserIds must have less than 100 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_ids", r.advertiserIds)
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
