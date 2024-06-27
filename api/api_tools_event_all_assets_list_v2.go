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

// ToolsEventAllAssetsListV2ApiService ToolsEventAllAssetsListV2Api service
type ToolsEventAllAssetsListV2ApiService service

type ApiOpenApi2ToolsEventAllAssetsListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsEventAllAssetsListV2ApiService
	advertiserId *int64
	filtering    *ToolsEventAllAssetsListV2Filtering
	page         *int32
	pageSize     *int32
}

// 广告账户id
func (r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsEventAllAssetsListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) Filtering(filtering ToolsEventAllAssetsListV2Filtering) *ApiOpenApi2ToolsEventAllAssetsListGetRequest {
	r.filtering = &filtering
	return r
}

// 页数，默认值&#x60;1&#x60;，最大值&#x60;999999&#x60;
func (r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) Page(page int32) *ApiOpenApi2ToolsEventAllAssetsListGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值&#x60;10&#x60;，最大&#x60;100&#x60;
func (r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsEventAllAssetsListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) Execute() (*ToolsEventAllAssetsListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsEventAllAssetsListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsEventAllAssetsListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsEventAllAssetsListGet Method for OpenApi2ToolsEventAllAssetsListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsEventAllAssetsListGetRequest
*/
func (a *ToolsEventAllAssetsListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsEventAllAssetsListGetRequest {
	return &ApiOpenApi2ToolsEventAllAssetsListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsEventAllAssetsListV2Response
func (a *ToolsEventAllAssetsListV2ApiService) getExecute(r *ApiOpenApi2ToolsEventAllAssetsListGetRequest) (*ToolsEventAllAssetsListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsEventAllAssetsListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/event/all_assets/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
