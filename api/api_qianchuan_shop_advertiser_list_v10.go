/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
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

// QianchuanShopAdvertiserListV10ApiService QianchuanShopAdvertiserListV10Api service
type QianchuanShopAdvertiserListV10ApiService service

type ApiOpenApiV10QianchuanShopAdvertiserListGetRequest struct {
	ctx        context.Context
	ApiService *QianchuanShopAdvertiserListV10ApiService
	shopId     *int64
	permission *[]*QianchuanShopAdvertiserListV10Permission
	page       *int64
	pageSize   *int64
}

func (r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) ShopId(shopId int64) *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest {
	r.shopId = &shopId
	return r
}

func (r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) Permission(permission []*QianchuanShopAdvertiserListV10Permission) *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest {
	r.permission = &permission
	return r
}

func (r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) Page(page int64) *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) PageSize(pageSize int64) *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) Execute() (*QianchuanShopAdvertiserListV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanShopAdvertiserListGet Method for OpenApiV10QianchuanShopAdvertiserListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanShopAdvertiserListGetRequest
*/
func (a *QianchuanShopAdvertiserListV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest {
	return &ApiOpenApiV10QianchuanShopAdvertiserListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanShopAdvertiserListV10Response
func (a *QianchuanShopAdvertiserListV10ApiService) getExecute(r *ApiOpenApiV10QianchuanShopAdvertiserListGetRequest) (*QianchuanShopAdvertiserListV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanShopAdvertiserListV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/shop/advertiser/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shopId == nil {
		return localVarReturnValue, nil, ReportError("shopId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "shop_id", r.shopId)
	if r.permission != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "permission", r.permission)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
