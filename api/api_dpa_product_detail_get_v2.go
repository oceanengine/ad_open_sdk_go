/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
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

// DpaProductDetailGetV2ApiService DpaProductDetailGetV2Api service
type DpaProductDetailGetV2ApiService service

type ApiOpenApi2DpaProductDetailGetGetRequest struct {
	ctx          context.Context
	ApiService   *DpaProductDetailGetV2ApiService
	advertiserId *int64
	filtering    *DpaProductDetailGetV2Filtering
	page         *int64
	pageSize     *int64
	platformId   *int64
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DpaProductDetailGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) Filtering(filtering DpaProductDetailGetV2Filtering) *ApiOpenApi2DpaProductDetailGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) Page(page int64) *ApiOpenApi2DpaProductDetailGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2DpaProductDetailGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) PlatformId(platformId int64) *ApiOpenApi2DpaProductDetailGetGetRequest {
	r.platformId = &platformId
	return r
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) Execute() (*DpaProductDetailGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2DpaProductDetailGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DpaProductDetailGetGetRequest) WithLog(enable bool) *ApiOpenApi2DpaProductDetailGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DpaProductDetailGetGet Method for OpenApi2DpaProductDetailGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DpaProductDetailGetGetRequest
*/
func (a *DpaProductDetailGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2DpaProductDetailGetGetRequest {
	return &ApiOpenApi2DpaProductDetailGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DpaProductDetailGetV2Response
func (a *DpaProductDetailGetV2ApiService) getExecute(r *ApiOpenApi2DpaProductDetailGetGetRequest) (*DpaProductDetailGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *DpaProductDetailGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dpa/product/detail/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
	if r.platformId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "platform_id", r.platformId)
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