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

// StardeliveryTaskAuthorVideoAuditV30ApiService StardeliveryTaskAuthorVideoAuditV30Api service
type StardeliveryTaskAuthorVideoAuditV30ApiService service

type ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest struct {
	ctx                                        context.Context
	ApiService                                 *StardeliveryTaskAuthorVideoAuditV30ApiService
	stardeliveryTaskAuthorVideoAuditV30Request *StardeliveryTaskAuthorVideoAuditV30Request
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest) StardeliveryTaskAuthorVideoAuditV30Request(stardeliveryTaskAuthorVideoAuditV30Request StardeliveryTaskAuthorVideoAuditV30Request) *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest {
	r.stardeliveryTaskAuthorVideoAuditV30Request = &stardeliveryTaskAuthorVideoAuditV30Request
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest) Execute() (*StardeliveryTaskAuthorVideoAuditV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest) AccessToken(accessToken string) *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest) WithLog(enable bool) *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30StardeliveryTaskAuthorVideoAuditPost Method for OpenApiV30StardeliveryTaskAuthorVideoAuditPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest
*/
func (a *StardeliveryTaskAuthorVideoAuditV30ApiService) Post(ctx context.Context) *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest {
	return &ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StardeliveryTaskAuthorVideoAuditV30Response
func (a *StardeliveryTaskAuthorVideoAuditV30ApiService) postExecute(r *ApiOpenApiV30StardeliveryTaskAuthorVideoAuditPostRequest) (*StardeliveryTaskAuthorVideoAuditV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StardeliveryTaskAuthorVideoAuditV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/stardelivery/task_author_video/audit/"

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
	localVarPostBody = r.stardeliveryTaskAuthorVideoAuditV30Request
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
