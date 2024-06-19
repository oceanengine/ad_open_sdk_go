/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
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

// ToolsAppManagementBookingGetV2ApiService ToolsAppManagementBookingGetV2Api service
type ToolsAppManagementBookingGetV2ApiService service

type ApiOpenApi2ToolsAppManagementBookingGetGetRequest struct {
	ctx                  context.Context
	ApiService           *ToolsAppManagementBookingGetV2ApiService
	advertiserId         *int64
	page                 *int32
	pageSize             *int32
	searchKey            *string
	searchType           *ToolsAppManagementBookingGetV2SearchType
	status               *ToolsAppManagementBookingGetV2Status
	scheduledPublishTime *ToolsAppManagementBookingGetV2ScheduledPublishTime
	createTime           *ToolsAppManagementBookingGetV2CreateTime
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) Page(page int32) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) SearchKey(searchKey string) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.searchKey = &searchKey
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) SearchType(searchType ToolsAppManagementBookingGetV2SearchType) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.searchType = &searchType
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) Status(status ToolsAppManagementBookingGetV2Status) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) ScheduledPublishTime(scheduledPublishTime ToolsAppManagementBookingGetV2ScheduledPublishTime) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.scheduledPublishTime = &scheduledPublishTime
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) CreateTime(createTime ToolsAppManagementBookingGetV2CreateTime) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.createTime = &createTime
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) Execute() (*ToolsAppManagementBookingGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementBookingGetGet Method for OpenApi2ToolsAppManagementBookingGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementBookingGetGetRequest
*/
func (a *ToolsAppManagementBookingGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementBookingGetGetRequest {
	return &ApiOpenApi2ToolsAppManagementBookingGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementBookingGetV2Response
func (a *ToolsAppManagementBookingGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementBookingGetGetRequest) (*ToolsAppManagementBookingGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppManagementBookingGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/booking/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.searchKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_key", r.searchKey)
	}
	if r.searchType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_type", r.searchType)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	if r.scheduledPublishTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scheduled_publish_time", r.scheduledPublishTime)
	}
	if r.createTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_time", r.createTime)
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
