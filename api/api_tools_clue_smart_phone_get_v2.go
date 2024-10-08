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

// ToolsClueSmartPhoneGetV2ApiService ToolsClueSmartPhoneGetV2Api service
type ToolsClueSmartPhoneGetV2ApiService service

type ApiOpenApi2ToolsClueSmartPhoneGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueSmartPhoneGetV2ApiService
	advertiserId *int64
	page         *int32
	pageSize     *int32
}

func (r *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest) Page(page int32) *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest) Execute() (*ToolsClueSmartPhoneGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueSmartPhoneGetGet Method for OpenApi2ToolsClueSmartPhoneGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueSmartPhoneGetGetRequest
*/
func (a *ToolsClueSmartPhoneGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest {
	return &ApiOpenApi2ToolsClueSmartPhoneGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueSmartPhoneGetV2Response
func (a *ToolsClueSmartPhoneGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueSmartPhoneGetGetRequest) (*ToolsClueSmartPhoneGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueSmartPhoneGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/smart_phone/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
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
