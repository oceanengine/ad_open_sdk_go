/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// StardeliveryTaskShareableListV30ApiService StardeliveryTaskShareableListV30Api service
type StardeliveryTaskShareableListV30ApiService service

type ApiOpenApiV30StardeliveryTaskShareableListGetRequest struct {
	ctx          context.Context
	ApiService   *StardeliveryTaskShareableListV30ApiService
	advertiserId *int64
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApiV30StardeliveryTaskShareableListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30StardeliveryTaskShareableListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 页码
func (r *ApiOpenApiV30StardeliveryTaskShareableListGetRequest) Page(page int64) *ApiOpenApiV30StardeliveryTaskShareableListGetRequest {
	r.page = &page
	return r
}

// 页面大小，上限100
func (r *ApiOpenApiV30StardeliveryTaskShareableListGetRequest) PageSize(pageSize int64) *ApiOpenApiV30StardeliveryTaskShareableListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskShareableListGetRequest) Execute() (*StardeliveryTaskShareableListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30StardeliveryTaskShareableListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30StardeliveryTaskShareableListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskShareableListGetRequest) WithLog(enable bool) *ApiOpenApiV30StardeliveryTaskShareableListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30StardeliveryTaskShareableListGet Method for OpenApiV30StardeliveryTaskShareableListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30StardeliveryTaskShareableListGetRequest
*/
func (a *StardeliveryTaskShareableListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30StardeliveryTaskShareableListGetRequest {
	return &ApiOpenApiV30StardeliveryTaskShareableListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StardeliveryTaskShareableListV30Response
func (a *StardeliveryTaskShareableListV30ApiService) getExecute(r *ApiOpenApiV30StardeliveryTaskShareableListGetRequest) (*StardeliveryTaskShareableListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StardeliveryTaskShareableListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/stardelivery/task/shareable/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
