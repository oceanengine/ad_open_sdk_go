/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

// StardeliveryTaskSharingListV30ApiService StardeliveryTaskSharingListV30Api service
type StardeliveryTaskSharingListV30ApiService service

type ApiOpenApiV30StardeliveryTaskSharingListGetRequest struct {
	ctx          context.Context
	ApiService   *StardeliveryTaskSharingListV30ApiService
	advertiserId *int64
	starTaskId   *int64
	page         *int32
	pageSize     *int32
}

// 广告账户id
func (r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30StardeliveryTaskSharingListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 星广联投任务id
func (r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) StarTaskId(starTaskId int64) *ApiOpenApiV30StardeliveryTaskSharingListGetRequest {
	r.starTaskId = &starTaskId
	return r
}

// 页数，默认1
func (r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) Page(page int32) *ApiOpenApiV30StardeliveryTaskSharingListGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认10
func (r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) PageSize(pageSize int32) *ApiOpenApiV30StardeliveryTaskSharingListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) Execute() (*StardeliveryTaskSharingListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30StardeliveryTaskSharingListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) WithLog(enable bool) *ApiOpenApiV30StardeliveryTaskSharingListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30StardeliveryTaskSharingListGet Method for OpenApiV30StardeliveryTaskSharingListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30StardeliveryTaskSharingListGetRequest
*/
func (a *StardeliveryTaskSharingListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30StardeliveryTaskSharingListGetRequest {
	return &ApiOpenApiV30StardeliveryTaskSharingListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StardeliveryTaskSharingListV30Response
func (a *StardeliveryTaskSharingListV30ApiService) getExecute(r *ApiOpenApiV30StardeliveryTaskSharingListGetRequest) (*StardeliveryTaskSharingListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StardeliveryTaskSharingListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/stardelivery/task/sharing/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.starTaskId == nil {
		return localVarReturnValue, nil, ReportError("starTaskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "star_task_id", r.starTaskId)
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