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

// ReportCustomConfigGetV30ApiService ReportCustomConfigGetV30Api service
type ReportCustomConfigGetV30ApiService service

type ApiOpenApiV30ReportCustomConfigGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportCustomConfigGetV30ApiService
	advertiserId *int64
	dataTopics   *[]*ReportCustomConfigGetV30DataTopics
}

func (r *ApiOpenApiV30ReportCustomConfigGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportCustomConfigGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 数据主题查询列表
func (r *ApiOpenApiV30ReportCustomConfigGetGetRequest) DataTopics(dataTopics []*ReportCustomConfigGetV30DataTopics) *ApiOpenApiV30ReportCustomConfigGetGetRequest {
	r.dataTopics = &dataTopics
	return r
}

func (r *ApiOpenApiV30ReportCustomConfigGetGetRequest) Execute() (*ReportCustomConfigGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportCustomConfigGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportCustomConfigGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportCustomConfigGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportCustomConfigGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportCustomConfigGetGet Method for OpenApiV30ReportCustomConfigGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportCustomConfigGetGetRequest
*/
func (a *ReportCustomConfigGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportCustomConfigGetGetRequest {
	return &ApiOpenApiV30ReportCustomConfigGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportCustomConfigGetV30Response
func (a *ReportCustomConfigGetV30ApiService) getExecute(r *ApiOpenApiV30ReportCustomConfigGetGetRequest) (*ReportCustomConfigGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportCustomConfigGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/custom/config/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.dataTopics == nil {
		return localVarReturnValue, nil, ReportError("dataTopics is required and must be specified")
	}
	if len(*r.dataTopics) < 1 {
		return localVarReturnValue, nil, ReportError("dataTopics must have at least 1 elements")
	}
	if len(*r.dataTopics) > 10 {
		return localVarReturnValue, nil, ReportError("dataTopics must have less than 10 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "data_topics", r.dataTopics)
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
