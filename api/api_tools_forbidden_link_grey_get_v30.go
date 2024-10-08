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

// ToolsForbiddenLinkGreyGetV30ApiService ToolsForbiddenLinkGreyGetV30Api service
type ToolsForbiddenLinkGreyGetV30ApiService service

type ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsForbiddenLinkGreyGetV30ApiService
	advertiserId *int64
}

func (r *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest) Execute() (*ToolsForbiddenLinkGreyGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsForbiddenLinkGreyGetGet Method for OpenApiV30ToolsForbiddenLinkGreyGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest
*/
func (a *ToolsForbiddenLinkGreyGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest {
	return &ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsForbiddenLinkGreyGetV30Response
func (a *ToolsForbiddenLinkGreyGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsForbiddenLinkGreyGetGetRequest) (*ToolsForbiddenLinkGreyGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsForbiddenLinkGreyGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/forbidden_link/grey/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
