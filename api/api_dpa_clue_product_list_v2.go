/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

// DpaClueProductListV2ApiService DpaClueProductListV2Api service
type DpaClueProductListV2ApiService service

type ApiOpenApi2DpaClueProductListGetRequest struct {
	ctx          context.Context
	ApiService   *DpaClueProductListV2ApiService
	advertiserId *int64
	page         *int64
	pageSize     *int64
	filtering    *DpaClueProductListV2Filtering
}

func (r *ApiOpenApi2DpaClueProductListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2DpaClueProductListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 分页参数
func (r *ApiOpenApi2DpaClueProductListGetRequest) Page(page int64) *ApiOpenApi2DpaClueProductListGetRequest {
	r.page = &page
	return r
}

// 分页参数
func (r *ApiOpenApi2DpaClueProductListGetRequest) PageSize(pageSize int64) *ApiOpenApi2DpaClueProductListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2DpaClueProductListGetRequest) Filtering(filtering DpaClueProductListV2Filtering) *ApiOpenApi2DpaClueProductListGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApi2DpaClueProductListGetRequest) Execute() (*DpaClueProductListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2DpaClueProductListGetRequest) AccessToken(accessToken string) *ApiOpenApi2DpaClueProductListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2DpaClueProductListGetRequest) WithLog(enable bool) *ApiOpenApi2DpaClueProductListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2DpaClueProductListGet Method for OpenApi2DpaClueProductListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2DpaClueProductListGetRequest
*/
func (a *DpaClueProductListV2ApiService) Get(ctx context.Context) *ApiOpenApi2DpaClueProductListGetRequest {
	return &ApiOpenApi2DpaClueProductListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DpaClueProductListV2Response
func (a *DpaClueProductListV2ApiService) getExecute(r *ApiOpenApi2DpaClueProductListGetRequest) (*DpaClueProductListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *DpaClueProductListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/dpa/clue_product/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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