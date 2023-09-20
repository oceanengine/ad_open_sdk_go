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

// DpaAssetsDetailReadV2ApiService DpaAssetsDetailReadV2Api service
type DpaAssetsDetailReadV2ApiService service

type ApiOpenApi2DpaAssetsDetailReadGetRequest struct {
	ctx          context.Context
	ApiService   *DpaAssetsDetailReadV2ApiService
	advertiserId *int64
	assetIds     *[]int64
}

func (r *ApiOpenApi2DpaAssetsDetailReadGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DpaAssetsDetailReadGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DpaAssetsDetailReadGetRequest) AssetIds(assetIds []int64) *ApiOpenApi2DpaAssetsDetailReadGetRequest {
	r.assetIds = &assetIds
	return r
}

func (r *ApiOpenApi2DpaAssetsDetailReadGetRequest) Execute() (*DpaAssetsDetailReadV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DpaAssetsDetailReadGetRequest) AccessToken(accessToken string) *ApiOpenApi2DpaAssetsDetailReadGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DpaAssetsDetailReadGetRequest) WithLog(enable bool) *ApiOpenApi2DpaAssetsDetailReadGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DpaAssetsDetailReadGet Method for OpenApi2DpaAssetsDetailReadGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DpaAssetsDetailReadGetRequest
*/
func (a *DpaAssetsDetailReadV2ApiService) Get(ctx context.Context) *ApiOpenApi2DpaAssetsDetailReadGetRequest {
	return &ApiOpenApi2DpaAssetsDetailReadGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DpaAssetsDetailReadV2Response
func (a *DpaAssetsDetailReadV2ApiService) getExecute(r *ApiOpenApi2DpaAssetsDetailReadGetRequest) (*DpaAssetsDetailReadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *DpaAssetsDetailReadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dpa/assets/detail/read/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.assetIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asset_ids", r.assetIds)
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
