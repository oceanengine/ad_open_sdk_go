/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
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

// FileVideoMaterialClearTaskGetV2ApiService FileVideoMaterialClearTaskGetV2Api service
type FileVideoMaterialClearTaskGetV2ApiService service

type ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest struct {
	ctx          context.Context
	ApiService   *FileVideoMaterialClearTaskGetV2ApiService
	advertiserId *int64
	clearIds     *[]int64
	page         *int64
	pageSize     *int64
}

// 广告主id
func (r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 清理任务id列表，最多支持10个
func (r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) ClearIds(clearIds []int64) *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest {
	r.clearIds = &clearIds
	return r
}

// 页码，默认：1
func (r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) Page(page int64) *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest {
	r.page = &page
	return r
}

// 页面大小，默认10，范围：1-100
func (r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) Execute() (*FileVideoMaterialClearTaskGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) WithLog(enable bool) *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoMaterialClearTaskGetGet Method for OpenApi2FileVideoMaterialClearTaskGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest
*/
func (a *FileVideoMaterialClearTaskGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest {
	return &ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoMaterialClearTaskGetV2Response
func (a *FileVideoMaterialClearTaskGetV2ApiService) getExecute(r *ApiOpenApi2FileVideoMaterialClearTaskGetGetRequest) (*FileVideoMaterialClearTaskGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *FileVideoMaterialClearTaskGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/material/clear_task/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.clearIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clear_ids", r.clearIds)
	}
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
