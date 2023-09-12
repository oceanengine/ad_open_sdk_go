/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"code.byted.org/ad/ad_open_sdk_go/config"
	. "code.byted.org/ad/ad_open_sdk_go/models"
)

// ToolsAppManagementExtendPackageListV2ApiService ToolsAppManagementExtendPackageListV2Api service
type ToolsAppManagementExtendPackageListV2ApiService service

type ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsAppManagementExtendPackageListV2ApiService
	advertiserId *int64
	filtering    *ToolsAppManagementExtendPackageListV2Filtering
	packageId    *string
	page         *int64
	pageSize     *int64
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) Filtering(filtering ToolsAppManagementExtendPackageListV2Filtering) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) PackageId(packageId string) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	r.packageId = &packageId
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) Page(page int64) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) PageSize(pageSize int64) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) Execute() (*ToolsAppManagementExtendPackageListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAppManagementExtendPackageListGet Method for OpenApi2ToolsAppManagementExtendPackageListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest
*/
func (a *ToolsAppManagementExtendPackageListV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest {
	return &ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAppManagementExtendPackageListV2Response
func (a *ToolsAppManagementExtendPackageListV2ApiService) getExecute(r *ApiOpenApi2ToolsAppManagementExtendPackageListGetRequest) (*ToolsAppManagementExtendPackageListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAppManagementExtendPackageListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/app_management/extend_package/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	if r.packageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "package_id", r.packageId)
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
