/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
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

// ServeMarketOrderGetV10ApiService ServeMarketOrderGetV10Api service
type ServeMarketOrderGetV10ApiService service

type ApiOpenApiV10ServeMarketOrderGetGetRequest struct {
	ctx        context.Context
	ApiService *ServeMarketOrderGetV10ApiService
	appId      *int64
	filtering  *ServeMarketOrderGetV10Filtering
	page       *int64
	pageSize   *int64
}

func (r *ApiOpenApiV10ServeMarketOrderGetGetRequest) AppId(appId int64) *ApiOpenApiV10ServeMarketOrderGetGetRequest {
	r.appId = &appId
	return r
}

func (r *ApiOpenApiV10ServeMarketOrderGetGetRequest) Filtering(filtering ServeMarketOrderGetV10Filtering) *ApiOpenApiV10ServeMarketOrderGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV10ServeMarketOrderGetGetRequest) Page(page int64) *ApiOpenApiV10ServeMarketOrderGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10ServeMarketOrderGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV10ServeMarketOrderGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10ServeMarketOrderGetGetRequest) Execute() (*ServeMarketOrderGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10ServeMarketOrderGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10ServeMarketOrderGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10ServeMarketOrderGetGetRequest) WithLog(enable bool) *ApiOpenApiV10ServeMarketOrderGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10ServeMarketOrderGetGet Method for OpenApiV10ServeMarketOrderGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10ServeMarketOrderGetGetRequest
*/
func (a *ServeMarketOrderGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10ServeMarketOrderGetGetRequest {
	return &ApiOpenApiV10ServeMarketOrderGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ServeMarketOrderGetV10Response
func (a *ServeMarketOrderGetV10ApiService) getExecute(r *ApiOpenApiV10ServeMarketOrderGetGetRequest) (*ServeMarketOrderGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ServeMarketOrderGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/serve_market/order/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.appId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_id", r.appId)
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
