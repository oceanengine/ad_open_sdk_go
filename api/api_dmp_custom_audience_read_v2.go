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

// DmpCustomAudienceReadV2ApiService DmpCustomAudienceReadV2Api service
type DmpCustomAudienceReadV2ApiService service

type ApiOpenApi2DmpCustomAudienceReadGetRequest struct {
	ctx               context.Context
	ApiService        *DmpCustomAudienceReadV2ApiService
	advertiserId      *int64
	customAudienceIds *[]int64
}

func (r *ApiOpenApi2DmpCustomAudienceReadGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DmpCustomAudienceReadGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceReadGetRequest) CustomAudienceIds(customAudienceIds []int64) *ApiOpenApi2DmpCustomAudienceReadGetRequest {
	r.customAudienceIds = &customAudienceIds
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceReadGetRequest) Execute() (*DmpCustomAudienceReadV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DmpCustomAudienceReadGetRequest) AccessToken(accessToken string) *ApiOpenApi2DmpCustomAudienceReadGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DmpCustomAudienceReadGetRequest) WithLog(enable bool) *ApiOpenApi2DmpCustomAudienceReadGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DmpCustomAudienceReadGet Method for OpenApi2DmpCustomAudienceReadGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DmpCustomAudienceReadGetRequest
*/
func (a *DmpCustomAudienceReadV2ApiService) Get(ctx context.Context) *ApiOpenApi2DmpCustomAudienceReadGetRequest {
	return &ApiOpenApi2DmpCustomAudienceReadGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DmpCustomAudienceReadV2Response
func (a *DmpCustomAudienceReadV2ApiService) getExecute(r *ApiOpenApi2DmpCustomAudienceReadGetRequest) (*DmpCustomAudienceReadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DmpCustomAudienceReadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dmp/custom_audience/read/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.customAudienceIds == nil {
		return localVarReturnValue, nil, ReportError("customAudienceIds is required and must be specified")
	}
	if len(*r.customAudienceIds) > 100 {
		return localVarReturnValue, nil, ReportError("customAudienceIds must have less than 100 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "custom_audience_ids", r.customAudienceIds)
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
