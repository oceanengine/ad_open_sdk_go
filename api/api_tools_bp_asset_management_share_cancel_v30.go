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

// ToolsBpAssetManagementShareCancelV30ApiService ToolsBpAssetManagementShareCancelV30Api service
type ToolsBpAssetManagementShareCancelV30ApiService service

type ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest struct {
	ctx                                         context.Context
	ApiService                                  *ToolsBpAssetManagementShareCancelV30ApiService
	toolsBpAssetManagementShareCancelV30Request *ToolsBpAssetManagementShareCancelV30Request
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest) ToolsBpAssetManagementShareCancelV30Request(toolsBpAssetManagementShareCancelV30Request ToolsBpAssetManagementShareCancelV30Request) *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest {
	r.toolsBpAssetManagementShareCancelV30Request = &toolsBpAssetManagementShareCancelV30Request
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest) Execute() (*ToolsBpAssetManagementShareCancelV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest) WithLog(enable bool) *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsBpAssetManagementShareCancelPost Method for OpenApiV30ToolsBpAssetManagementShareCancelPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest
*/
func (a *ToolsBpAssetManagementShareCancelV30ApiService) Post(ctx context.Context) *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest {
	return &ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsBpAssetManagementShareCancelV30Response
func (a *ToolsBpAssetManagementShareCancelV30ApiService) postExecute(r *ApiOpenApiV30ToolsBpAssetManagementShareCancelPostRequest) (*ToolsBpAssetManagementShareCancelV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsBpAssetManagementShareCancelV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/bp_asset_management/share/cancel/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// body params
	localVarPostBody = r.toolsBpAssetManagementShareCancelV30Request
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
