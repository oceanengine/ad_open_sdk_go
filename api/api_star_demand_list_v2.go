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

// StarDemandListV2ApiService StarDemandListV2Api service
type StarDemandListV2ApiService service

type ApiOpenApi2StarDemandListGetRequest struct {
	ctx        context.Context
	ApiService *StarDemandListV2ApiService
	starId     *int64
	filtering  *StarDemandListV2Filtering
	page       *int64
	pageSize   *int64
}

func (r *ApiOpenApi2StarDemandListGetRequest) StarId(starId int64) *ApiOpenApi2StarDemandListGetRequest {
	r.starId = &starId
	return r
}

// 过滤器
func (r *ApiOpenApi2StarDemandListGetRequest) Filtering(filtering StarDemandListV2Filtering) *ApiOpenApi2StarDemandListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2StarDemandListGetRequest) Page(page int64) *ApiOpenApi2StarDemandListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2StarDemandListGetRequest) PageSize(pageSize int64) *ApiOpenApi2StarDemandListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2StarDemandListGetRequest) Execute() (*StarDemandListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarDemandListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandListGetRequest) WithLog(enable bool) *ApiOpenApi2StarDemandListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandListGet Method for OpenApi2StarDemandListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandListGetRequest
*/
func (a *StarDemandListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarDemandListGetRequest {
	return &ApiOpenApi2StarDemandListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandListV2Response
func (a *StarDemandListV2ApiService) getExecute(r *ApiOpenApi2StarDemandListGetRequest) (*StarDemandListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
