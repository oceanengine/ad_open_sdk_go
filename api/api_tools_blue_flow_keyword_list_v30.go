/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

// ToolsBlueFlowKeywordListV30ApiService ToolsBlueFlowKeywordListV30Api service
type ToolsBlueFlowKeywordListV30ApiService service

type ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsBlueFlowKeywordListV30ApiService
	advertiserId *int64
	projectId    *int64
	filtering    *ToolsBlueFlowKeywordListV30Filtering
}

func (r *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest) ProjectId(projectId int64) *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest {
	r.projectId = &projectId
	return r
}

// 过滤器
func (r *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest) Filtering(filtering ToolsBlueFlowKeywordListV30Filtering) *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest) Execute() (*ToolsBlueFlowKeywordListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsBlueFlowKeywordListGet Method for OpenApiV30ToolsBlueFlowKeywordListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest
*/
func (a *ToolsBlueFlowKeywordListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest {
	return &ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsBlueFlowKeywordListV30Response
func (a *ToolsBlueFlowKeywordListV30ApiService) getExecute(r *ApiOpenApiV30ToolsBlueFlowKeywordListGetRequest) (*ToolsBlueFlowKeywordListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsBlueFlowKeywordListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/blue_flow_keyword/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.projectId == nil {
		return localVarReturnValue, nil, ReportError("projectId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId)
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
