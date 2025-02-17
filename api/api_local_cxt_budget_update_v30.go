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

// LocalCxtBudgetUpdateV30ApiService LocalCxtBudgetUpdateV30Api service
type LocalCxtBudgetUpdateV30ApiService service

type ApiOpenApiV30LocalCxtBudgetUpdatePostRequest struct {
	ctx                            context.Context
	ApiService                     *LocalCxtBudgetUpdateV30ApiService
	localCxtBudgetUpdateV30Request *LocalCxtBudgetUpdateV30Request
}

func (r *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest) LocalCxtBudgetUpdateV30Request(localCxtBudgetUpdateV30Request LocalCxtBudgetUpdateV30Request) *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest {
	r.localCxtBudgetUpdateV30Request = &localCxtBudgetUpdateV30Request
	return r
}

func (r *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest) Execute() (*LocalCxtBudgetUpdateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30LocalCxtBudgetUpdatePost Method for OpenApiV30LocalCxtBudgetUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30LocalCxtBudgetUpdatePostRequest
*/
func (a *LocalCxtBudgetUpdateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest {
	return &ApiOpenApiV30LocalCxtBudgetUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return LocalCxtBudgetUpdateV30Response
func (a *LocalCxtBudgetUpdateV30ApiService) postExecute(r *ApiOpenApiV30LocalCxtBudgetUpdatePostRequest) (*LocalCxtBudgetUpdateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *LocalCxtBudgetUpdateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/local/cxt/budget/update/"

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
	localVarPostBody = r.localCxtBudgetUpdateV30Request
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
