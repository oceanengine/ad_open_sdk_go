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

// BudgetGroupCreateV30ApiService BudgetGroupCreateV30Api service
type BudgetGroupCreateV30ApiService service

type ApiOpenApiV30BudgetGroupCreatePostRequest struct {
	ctx                         context.Context
	ApiService                  *BudgetGroupCreateV30ApiService
	budgetGroupCreateV30Request *BudgetGroupCreateV30Request
}

func (r *ApiOpenApiV30BudgetGroupCreatePostRequest) BudgetGroupCreateV30Request(budgetGroupCreateV30Request BudgetGroupCreateV30Request) *ApiOpenApiV30BudgetGroupCreatePostRequest {
	r.budgetGroupCreateV30Request = &budgetGroupCreateV30Request
	return r
}

func (r *ApiOpenApiV30BudgetGroupCreatePostRequest) Execute() (*BudgetGroupCreateV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30BudgetGroupCreatePostRequest) AccessToken(accessToken string) *ApiOpenApiV30BudgetGroupCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BudgetGroupCreatePostRequest) WithLog(enable bool) *ApiOpenApiV30BudgetGroupCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BudgetGroupCreatePost Method for OpenApiV30BudgetGroupCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BudgetGroupCreatePostRequest
*/
func (a *BudgetGroupCreateV30ApiService) Post(ctx context.Context) *ApiOpenApiV30BudgetGroupCreatePostRequest {
	return &ApiOpenApiV30BudgetGroupCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BudgetGroupCreateV30Response
func (a *BudgetGroupCreateV30ApiService) postExecute(r *ApiOpenApiV30BudgetGroupCreatePostRequest) (*BudgetGroupCreateV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BudgetGroupCreateV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/budget_group/create/"

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
	localVarPostBody = r.budgetGroupCreateV30Request
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
