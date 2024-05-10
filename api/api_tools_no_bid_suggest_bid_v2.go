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

// ToolsNoBidSuggestBidV2ApiService ToolsNoBidSuggestBidV2Api service
type ToolsNoBidSuggestBidV2ApiService service

type ApiOpenApi2ToolsNoBidSuggestBidGetRequest struct {
	ctx            context.Context
	ApiService     *ToolsNoBidSuggestBidV2ApiService
	advertiserId   *int64
	budget         *int64
	budgetMode     *ToolsNoBidSuggestBidV2BudgetMode
	externalAction *ToolsNoBidSuggestBidV2ExternalAction
	filtering      *ToolsNoBidSuggestBidV2Filtering
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) Budget(budget int64) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	r.budget = &budget
	return r
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) BudgetMode(budgetMode ToolsNoBidSuggestBidV2BudgetMode) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	r.budgetMode = &budgetMode
	return r
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) ExternalAction(externalAction ToolsNoBidSuggestBidV2ExternalAction) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	r.externalAction = &externalAction
	return r
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) Filtering(filtering ToolsNoBidSuggestBidV2Filtering) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) Execute() (*ToolsNoBidSuggestBidV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsNoBidSuggestBidGet Method for OpenApi2ToolsNoBidSuggestBidGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsNoBidSuggestBidGetRequest
*/
func (a *ToolsNoBidSuggestBidV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsNoBidSuggestBidGetRequest {
	return &ApiOpenApi2ToolsNoBidSuggestBidGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsNoBidSuggestBidV2Response
func (a *ToolsNoBidSuggestBidV2ApiService) getExecute(r *ApiOpenApi2ToolsNoBidSuggestBidGetRequest) (*ToolsNoBidSuggestBidV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsNoBidSuggestBidV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/no_bid/suggest_bid/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.budget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "budget", r.budget)
	}
	if r.budgetMode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "budget_mode", r.budgetMode)
	}
	if r.externalAction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_action", r.externalAction)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
