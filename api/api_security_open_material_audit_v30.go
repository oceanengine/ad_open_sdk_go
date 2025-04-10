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

// SecurityOpenMaterialAuditV30ApiService SecurityOpenMaterialAuditV30Api service
type SecurityOpenMaterialAuditV30ApiService service

type ApiOpenApiV30SecurityOpenMaterialAuditPostRequest struct {
	ctx                                 context.Context
	ApiService                          *SecurityOpenMaterialAuditV30ApiService
	securityOpenMaterialAuditV30Request *SecurityOpenMaterialAuditV30Request
}

func (r *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest) SecurityOpenMaterialAuditV30Request(securityOpenMaterialAuditV30Request SecurityOpenMaterialAuditV30Request) *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest {
	r.securityOpenMaterialAuditV30Request = &securityOpenMaterialAuditV30Request
	return r
}

func (r *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest) Execute() (*SecurityOpenMaterialAuditV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest) AccessToken(accessToken string) *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest) WithLog(enable bool) *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30SecurityOpenMaterialAuditPost Method for OpenApiV30SecurityOpenMaterialAuditPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30SecurityOpenMaterialAuditPostRequest
*/
func (a *SecurityOpenMaterialAuditV30ApiService) Post(ctx context.Context) *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest {
	return &ApiOpenApiV30SecurityOpenMaterialAuditPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SecurityOpenMaterialAuditV30Response
func (a *SecurityOpenMaterialAuditV30ApiService) postExecute(r *ApiOpenApiV30SecurityOpenMaterialAuditPostRequest) (*SecurityOpenMaterialAuditV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *SecurityOpenMaterialAuditV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/security/open_material_audit/"

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
	localVarPostBody = r.securityOpenMaterialAuditV30Request
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
