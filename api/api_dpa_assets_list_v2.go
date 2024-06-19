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

// DpaAssetsListV2ApiService DpaAssetsListV2Api service
type DpaAssetsListV2ApiService service

type ApiOpenApi2DpaAssetsListGetRequest struct {
	ctx          context.Context
	ApiService   *DpaAssetsListV2ApiService
	advertiserId *int64
	filtering    *DpaAssetsListV2Filtering
	page         *int64
	pageSize     *int64
	platformId   *int64
	productIds   *[]int64
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DpaAssetsListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) Filtering(filtering DpaAssetsListV2Filtering) *ApiOpenApi2DpaAssetsListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) Page(page int64) *ApiOpenApi2DpaAssetsListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) PageSize(pageSize int64) *ApiOpenApi2DpaAssetsListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) PlatformId(platformId int64) *ApiOpenApi2DpaAssetsListGetRequest {
	r.platformId = &platformId
	return r
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) ProductIds(productIds []int64) *ApiOpenApi2DpaAssetsListGetRequest {
	r.productIds = &productIds
	return r
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) Execute() (*DpaAssetsListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) AccessToken(accessToken string) *ApiOpenApi2DpaAssetsListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DpaAssetsListGetRequest) WithLog(enable bool) *ApiOpenApi2DpaAssetsListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DpaAssetsListGet Method for OpenApi2DpaAssetsListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DpaAssetsListGetRequest
*/
func (a *DpaAssetsListV2ApiService) Get(ctx context.Context) *ApiOpenApi2DpaAssetsListGetRequest {
	return &ApiOpenApi2DpaAssetsListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DpaAssetsListV2Response
func (a *DpaAssetsListV2ApiService) getExecute(r *ApiOpenApi2DpaAssetsListGetRequest) (*DpaAssetsListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DpaAssetsListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dpa/assets/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
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
	if r.productIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_ids", r.productIds)
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
