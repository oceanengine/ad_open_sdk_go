/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

// ToolsEventAssetsGetV2ApiService ToolsEventAssetsGetV2Api service
type ToolsEventAssetsGetV2ApiService service

type ApiOpenApi2ToolsEventAssetsGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsEventAssetsGetV2ApiService
	advertiserId *int64
	assetType    *ToolsEventAssetsGetV2AssetType
	filtering    *ToolsEventAssetsGetV2Filtering
	sortType     *ToolsEventAssetsGetV2SortType
	page         *int64
	pageSize     *int64
}

// 广告主 id
func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 资产类型，允许值:&#x60;THIRD_EXTERNAL&#x60;:三方落地页,&#x60;APP&#x60;:应用,&#x60;QUICK_APP&#x60;:快应用,&#x60;MINI_PROGRAME&#x60;:小程序
func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) AssetType(assetType ToolsEventAssetsGetV2AssetType) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	r.assetType = &assetType
	return r
}

// 过滤条件
func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) Filtering(filtering ToolsEventAssetsGetV2Filtering) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	r.filtering = &filtering
	return r
}

// 排序方式，允许值：&#x60;ASC&#x60;：升序  &#x60;DESC&#x60;：降序&lt;br&gt;默认值&#x60;ASC&#x60;
func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) SortType(sortType ToolsEventAssetsGetV2SortType) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	r.sortType = &sortType
	return r
}

// 页码，默认值&#x60;1&#x60;
func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) Page(page int64) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认值&#x60;10&#x60;，最大&#x60;30&#x60;
func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) Execute() (*ToolsEventAssetsGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsEventAssetsGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsEventAssetsGetGet Method for OpenApi2ToolsEventAssetsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsEventAssetsGetGetRequest
*/
func (a *ToolsEventAssetsGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsEventAssetsGetGetRequest {
	return &ApiOpenApi2ToolsEventAssetsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsEventAssetsGetV2Response
func (a *ToolsEventAssetsGetV2ApiService) getExecute(r *ApiOpenApi2ToolsEventAssetsGetGetRequest) (*ToolsEventAssetsGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsEventAssetsGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/event/assets/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than 9223372036854775807")
	}
	if r.assetType == nil {
		return localVarReturnValue, nil, ReportError("assetType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_type", r.assetType)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_type", r.sortType)
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
