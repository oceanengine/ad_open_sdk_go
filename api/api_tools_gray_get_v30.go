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

// ToolsGrayGetV30ApiService ToolsGrayGetV30Api service
type ToolsGrayGetV30ApiService service

type ApiOpenApiV30ToolsGrayGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsGrayGetV30ApiService
	advertiserId *int64
	grayKeys     *[]string
	version      *ToolsGrayGetV30Version
}

func (r *ApiOpenApiV30ToolsGrayGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsGrayGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsGrayGetGetRequest) GrayKeys(grayKeys []string) *ApiOpenApiV30ToolsGrayGetGetRequest {
	r.grayKeys = &grayKeys
	return r
}

func (r *ApiOpenApiV30ToolsGrayGetGetRequest) Version(version ToolsGrayGetV30Version) *ApiOpenApiV30ToolsGrayGetGetRequest {
	r.version = &version
	return r
}

func (r *ApiOpenApiV30ToolsGrayGetGetRequest) Execute() (*ToolsGrayGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsGrayGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsGrayGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsGrayGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsGrayGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsGrayGetGet Method for OpenApiV30ToolsGrayGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsGrayGetGetRequest
*/
func (a *ToolsGrayGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsGrayGetGetRequest {
	return &ApiOpenApiV30ToolsGrayGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsGrayGetV30Response
func (a *ToolsGrayGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsGrayGetGetRequest) (*ToolsGrayGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsGrayGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/gray/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.grayKeys == nil {
		return localVarReturnValue, nil, ReportError("grayKeys is required and must be specified")
	}
	if len(*r.grayKeys) < 1 {
		return localVarReturnValue, nil, ReportError("grayKeys must have at least 1 elements")
	}
	if len(*r.grayKeys) > 1 {
		return localVarReturnValue, nil, ReportError("grayKeys must have less than 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "gray_keys", r.grayKeys)
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version)
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
