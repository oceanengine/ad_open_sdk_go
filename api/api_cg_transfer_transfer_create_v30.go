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

// CgTransferTransferCreateV30ApiService CgTransferTransferCreateV30Api service
type CgTransferTransferCreateV30ApiService service

type ApiOpenApiV30CgTransferTransferCreatePostRequest struct {
	ctx                                context.Context
	ApiService                         *CgTransferTransferCreateV30ApiService
	cgTransferTransferCreateV30Request *CgTransferTransferCreateV30Request
}

func (r *ApiOpenApiV30CgTransferTransferCreatePostRequest) CgTransferTransferCreateV30Request(cgTransferTransferCreateV30Request CgTransferTransferCreateV30Request) *ApiOpenApiV30CgTransferTransferCreatePostRequest {
	r.cgTransferTransferCreateV30Request = &cgTransferTransferCreateV30Request
	return r
}

func (r *ApiOpenApiV30CgTransferTransferCreatePostRequest) Execute() (*CgTransferTransferCreateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30CgTransferTransferCreatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30CgTransferTransferCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30CgTransferTransferCreatePostRequest) WithLog(enable bool) *ApiOpenApiV30CgTransferTransferCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30CgTransferTransferCreatePost Method for OpenApiV30CgTransferTransferCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30CgTransferTransferCreatePostRequest
*/
func (a *CgTransferTransferCreateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30CgTransferTransferCreatePostRequest {
	return &ApiOpenApiV30CgTransferTransferCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CgTransferTransferCreateV30Response
func (a *CgTransferTransferCreateV30ApiService) postExecute(r *ApiOpenApiV30CgTransferTransferCreatePostRequest) (*CgTransferTransferCreateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *CgTransferTransferCreateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/cg_transfer/transfer/create/"

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
	localVarPostBody = r.cgTransferTransferCreateV30Request
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
