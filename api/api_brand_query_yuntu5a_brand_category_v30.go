/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
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

// BrandQueryYuntu5aBrandCategoryV30ApiService BrandQueryYuntu5aBrandCategoryV30Api service
type BrandQueryYuntu5aBrandCategoryV30ApiService service

type ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest struct {
	ctx          context.Context
	ApiService   *BrandQueryYuntu5aBrandCategoryV30ApiService
	advertiserId *int64
}

func (r *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest) Execute() (*BrandQueryYuntu5aBrandCategoryV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest) WithLog(enable bool) *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandQueryYuntu5aBrandCategoryGet Method for OpenApiV30BrandQueryYuntu5aBrandCategoryGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest
*/
func (a *BrandQueryYuntu5aBrandCategoryV30ApiService) Get(ctx context.Context) *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest {
	return &ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandQueryYuntu5aBrandCategoryV30Response
func (a *BrandQueryYuntu5aBrandCategoryV30ApiService) getExecute(r *ApiOpenApiV30BrandQueryYuntu5aBrandCategoryGetRequest) (*BrandQueryYuntu5aBrandCategoryV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *BrandQueryYuntu5aBrandCategoryV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/query_yuntu_5a_brand_category/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
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
