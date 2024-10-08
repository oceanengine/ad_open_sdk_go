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

// StarVasGetCommonAuthorPackageListV2ApiService StarVasGetCommonAuthorPackageListV2Api service
type StarVasGetCommonAuthorPackageListV2ApiService service

type ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest struct {
	ctx        context.Context
	ApiService *StarVasGetCommonAuthorPackageListV2ApiService
	starId     *int64
}

// 客户ID
func (r *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest) StarId(starId int64) *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest) Execute() (*StarVasGetCommonAuthorPackageListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest) WithLog(enable bool) *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarVasGetCommonAuthorPackageListGet Method for OpenApi2StarVasGetCommonAuthorPackageListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest
*/
func (a *StarVasGetCommonAuthorPackageListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest {
	return &ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarVasGetCommonAuthorPackageListV2Response
func (a *StarVasGetCommonAuthorPackageListV2ApiService) getExecute(r *ApiOpenApi2StarVasGetCommonAuthorPackageListGetRequest) (*StarVasGetCommonAuthorPackageListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarVasGetCommonAuthorPackageListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/vas/get_common_author_package_list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
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
