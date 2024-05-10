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

// ToolsCommentMid2itemIdV30ApiService ToolsCommentMid2itemIdV30Api service
type ToolsCommentMid2itemIdV30ApiService service

type ApiOpenApiV30ToolsCommentMid2itemIdGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsCommentMid2itemIdV30ApiService
	advertiserId *int64
	startTime    *string
	endTime      *string
	materialId   *int64
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) StartTime(startTime string) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) EndTime(endTime string) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) MaterialId(materialId int64) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	r.materialId = &materialId
	return r
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) Page(page int64) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) Execute() (*ToolsCommentMid2itemIdV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsCommentMid2itemIdGet Method for OpenApiV30ToolsCommentMid2itemIdGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsCommentMid2itemIdGetRequest
*/
func (a *ToolsCommentMid2itemIdV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest {
	return &ApiOpenApiV30ToolsCommentMid2itemIdGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsCommentMid2itemIdV30Response
func (a *ToolsCommentMid2itemIdV30ApiService) getExecute(r *ApiOpenApiV30ToolsCommentMid2itemIdGetRequest) (*ToolsCommentMid2itemIdV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsCommentMid2itemIdV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/comment/mid2item_id/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.materialId == nil {
		return localVarReturnValue, nil, ReportError("materialId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_id", r.materialId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
