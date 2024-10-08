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

// ToolsAppManagementBookingRecordsGetV2ApiService ToolsAppManagementBookingRecordsGetV2Api service
type ToolsAppManagementBookingRecordsGetV2ApiService service

type ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAppManagementBookingRecordsGetV2ApiService
	advertiserId *int64
	packageId    *string
	hostType     *ToolsAppManagementBookingRecordsGetV2HostType
	page         *int64
	pageSize     *int64
	createTime   *ToolsAppManagementBookingRecordsGetV2CreateTime
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) PackageId(packageId string) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	r.packageId = &packageId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) HostType(hostType ToolsAppManagementBookingRecordsGetV2HostType) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	r.hostType = &hostType
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) Page(page int64) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) CreateTime(createTime ToolsAppManagementBookingRecordsGetV2CreateTime) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	r.createTime = &createTime
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) Execute() (*ToolsAppManagementBookingRecordsGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementBookingRecordsGetGet Method for OpenApi2ToolsAppManagementBookingRecordsGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest
*/
func (a *ToolsAppManagementBookingRecordsGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest {
	return &ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementBookingRecordsGetV2Response
func (a *ToolsAppManagementBookingRecordsGetV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementBookingRecordsGetGetRequest) (*ToolsAppManagementBookingRecordsGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAppManagementBookingRecordsGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/booking_records/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.packageId == nil {
		return localVarReturnValue, nil, ReportError("packageId is required and must be specified")
	}
	if r.hostType == nil {
		return localVarReturnValue, nil, ReportError("hostType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "package_id", r.packageId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "host_type", r.hostType)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
