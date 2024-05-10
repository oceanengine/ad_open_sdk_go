/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

// QianchuanToolsSmartBoostAdBoostSetV10ApiService QianchuanToolsSmartBoostAdBoostSetV10Api service
type QianchuanToolsSmartBoostAdBoostSetV10ApiService service

type ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest struct {
	ctx                                          context.Context
	ApiService                                   *QianchuanToolsSmartBoostAdBoostSetV10ApiService
	qianchuanToolsSmartBoostAdBoostSetV10Request *QianchuanToolsSmartBoostAdBoostSetV10Request
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest) QianchuanToolsSmartBoostAdBoostSetV10Request(qianchuanToolsSmartBoostAdBoostSetV10Request QianchuanToolsSmartBoostAdBoostSetV10Request) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest {
	r.qianchuanToolsSmartBoostAdBoostSetV10Request = &qianchuanToolsSmartBoostAdBoostSetV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest) Execute() (*QianchuanToolsSmartBoostAdBoostSetV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanToolsSmartBoostAdBoostSetPost Method for OpenApiV10QianchuanToolsSmartBoostAdBoostSetPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest
*/
func (a *QianchuanToolsSmartBoostAdBoostSetV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest {
	return &ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanToolsSmartBoostAdBoostSetV10Response
func (a *QianchuanToolsSmartBoostAdBoostSetV10ApiService) postExecute(r *ApiOpenApiV10QianchuanToolsSmartBoostAdBoostSetPostRequest) (*QianchuanToolsSmartBoostAdBoostSetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanToolsSmartBoostAdBoostSetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/tools/smart_boost/ad_boost/set/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.qianchuanToolsSmartBoostAdBoostSetV10Request
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
