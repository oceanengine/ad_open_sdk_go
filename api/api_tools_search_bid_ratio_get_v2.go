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

// ToolsSearchBidRatioGetV2ApiService ToolsSearchBidRatioGetV2Api service
type ToolsSearchBidRatioGetV2ApiService service

type ApiOpenApi2ToolsSearchBidRatioGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsSearchBidRatioGetV2ApiService
	advertiserId *int64
	adId         *int64
}

// 广告主ID
func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告ID，修改广告时需要传
func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) AdId(adId int64) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) Execute() (*ToolsSearchBidRatioGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsSearchBidRatioGetGet Method for OpenApi2ToolsSearchBidRatioGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsSearchBidRatioGetGetRequest
*/
func (a *ToolsSearchBidRatioGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsSearchBidRatioGetGetRequest {
	return &ApiOpenApi2ToolsSearchBidRatioGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsSearchBidRatioGetV2Response
func (a *ToolsSearchBidRatioGetV2ApiService) getExecute(r *ApiOpenApi2ToolsSearchBidRatioGetGetRequest) (*ToolsSearchBidRatioGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsSearchBidRatioGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/search_bid_ratio/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.adId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
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
