/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
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

// ToolsAdRaiseEstimateGetV2ApiService ToolsAdRaiseEstimateGetV2Api service
type ToolsAdRaiseEstimateGetV2ApiService service

type ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAdRaiseEstimateGetV2ApiService
	advertiserId *int64
	adId         *int64
	modifyValue  *int64
	budgetValue  *int64
}

func (r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) AdId(adId int64) *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest {
	r.adId = &adId
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) ModifyValue(modifyValue int64) *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest {
	r.modifyValue = &modifyValue
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) BudgetValue(budgetValue int64) *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest {
	r.budgetValue = &budgetValue
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) Execute() (*ToolsAdRaiseEstimateGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdRaiseEstimateGetGet Method for OpenApi2ToolsAdRaiseEstimateGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest
*/
func (a *ToolsAdRaiseEstimateGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest {
	return &ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdRaiseEstimateGetV2Response
func (a *ToolsAdRaiseEstimateGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAdRaiseEstimateGetGetRequest) (*ToolsAdRaiseEstimateGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAdRaiseEstimateGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ad_raise_estimate/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 0 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 0")
	}
	if *r.advertiserId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than 9223372036854775807")
	}
	if r.adId == nil {
		return localVarReturnValue, nil, ReportError("adId is required and must be specified")
	}
	if *r.adId < 0 {
		return localVarReturnValue, nil, ReportError("adId must be greater than 0")
	}
	if *r.adId > 9223372036854775807 {
		return localVarReturnValue, nil, ReportError("adId must be less than 9223372036854775807")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "ad_id", r.adId)
	if r.modifyValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "modify_value", r.modifyValue)
	}
	if r.budgetValue != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "budget_value", r.budgetValue)
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
