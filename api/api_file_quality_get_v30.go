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

// FileQualityGetV30ApiService FileQualityGetV30Api service
type FileQualityGetV30ApiService service

type ApiOpenApiV30FileQualityGetGetRequest struct {
	ctx          context.Context
	ApiService   *FileQualityGetV30ApiService
	advertiserId *int64
	materialIds  *[]int64
}

func (r *ApiOpenApiV30FileQualityGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30FileQualityGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30FileQualityGetGetRequest) MaterialIds(materialIds []int64) *ApiOpenApiV30FileQualityGetGetRequest {
	r.materialIds = &materialIds
	return r
}

func (r *ApiOpenApiV30FileQualityGetGetRequest) Execute() (*FileQualityGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30FileQualityGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30FileQualityGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30FileQualityGetGetRequest) WithLog(enable bool) *ApiOpenApiV30FileQualityGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30FileQualityGetGet Method for OpenApiV30FileQualityGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30FileQualityGetGetRequest
*/
func (a *FileQualityGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30FileQualityGetGetRequest {
	return &ApiOpenApiV30FileQualityGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileQualityGetV30Response
func (a *FileQualityGetV30ApiService) getExecute(r *ApiOpenApiV30FileQualityGetGetRequest) (*FileQualityGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileQualityGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/file/quality/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.materialIds == nil {
		return localVarReturnValue, nil, ReportError("materialIds is required and must be specified")
	}
	if len(*r.materialIds) < 1 {
		return localVarReturnValue, nil, ReportError("materialIds must have at least 1 elements")
	}
	if len(*r.materialIds) > 50 {
		return localVarReturnValue, nil, ReportError("materialIds must have less than 50 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "material_ids", r.materialIds)
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
