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

// BudgetGroupListV30ApiService BudgetGroupListV30Api service
type BudgetGroupListV30ApiService service

type ApiOpenApiV30BudgetGroupListGetRequest struct {
	ctx          context.Context
	ApiService   *BudgetGroupListV30ApiService
	advertiserId *int64
	page         *int64
	pageSize     *int64
	filtering    *BudgetGroupListV30Filtering
}

// 广告主id
func (r *ApiOpenApiV30BudgetGroupListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BudgetGroupListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30BudgetGroupListGetRequest) Page(page int64) *ApiOpenApiV30BudgetGroupListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30BudgetGroupListGetRequest) PageSize(pageSize int64) *ApiOpenApiV30BudgetGroupListGetRequest {
	r.pageSize = &pageSize
	return r
}

// 过滤条件
func (r *ApiOpenApiV30BudgetGroupListGetRequest) Filtering(filtering BudgetGroupListV30Filtering) *ApiOpenApiV30BudgetGroupListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30BudgetGroupListGetRequest) Execute() (*BudgetGroupListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30BudgetGroupListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30BudgetGroupListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BudgetGroupListGetRequest) WithLog(enable bool) *ApiOpenApiV30BudgetGroupListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BudgetGroupListGet Method for OpenApiV30BudgetGroupListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BudgetGroupListGetRequest
*/
func (a *BudgetGroupListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30BudgetGroupListGetRequest {
	return &ApiOpenApiV30BudgetGroupListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BudgetGroupListV30Response
func (a *BudgetGroupListV30ApiService) getExecute(r *ApiOpenApiV30BudgetGroupListGetRequest) (*BudgetGroupListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *BudgetGroupListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/budget_group/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if *r.page > 1000 {
		return localVarReturnValue, nil, ReportError("page must be less than 1000")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, ReportError("pageSize must be greater than 1")
	}
	if *r.pageSize > 1000 {
		return localVarReturnValue, nil, ReportError("pageSize must be less than 1000")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
