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

// ToolsClueFormGetV2ApiService ToolsClueFormGetV2Api service
type ToolsClueFormGetV2ApiService service

type ApiOpenApi2ToolsClueFormGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsClueFormGetV2ApiService
	advertiserId *int64
	page         *int32
	pageSize     *int32
	startTime    *string
	endTime      *string
	instanceId   *int64
	name         *string
	isDel        *int64
	formType     *ToolsClueFormGetV2FormType
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) Page(page int32) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) PageSize(pageSize int32) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) StartTime(startTime string) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) EndTime(endTime string) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) InstanceId(instanceId int64) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.instanceId = &instanceId
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) Name(name string) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.name = &name
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) IsDel(isDel int64) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.isDel = &isDel
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) FormType(formType ToolsClueFormGetV2FormType) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.formType = &formType
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) Execute() (*ToolsClueFormGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsClueFormGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsClueFormGetGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsClueFormGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsClueFormGetGet Method for OpenApi2ToolsClueFormGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsClueFormGetGetRequest
*/
func (a *ToolsClueFormGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsClueFormGetGetRequest {
	return &ApiOpenApi2ToolsClueFormGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsClueFormGetV2Response
func (a *ToolsClueFormGetV2ApiService) getExecute(r *ApiOpenApi2ToolsClueFormGetGetRequest) (*ToolsClueFormGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsClueFormGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/clue/form/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.instanceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "instance_id", r.instanceId)
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name)
	}
	if r.isDel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_del", r.isDel)
	}
	if r.formType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "form_type", r.formType)
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
