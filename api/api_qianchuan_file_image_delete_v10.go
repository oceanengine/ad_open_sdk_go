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

// QianchuanFileImageDeleteV10ApiService QianchuanFileImageDeleteV10Api service
type QianchuanFileImageDeleteV10ApiService service

type ApiOpenApiV10QianchuanFileImageDeletePostRequest struct {
	ctx                                context.Context
	ApiService                         *QianchuanFileImageDeleteV10ApiService
	qianchuanFileImageDeleteV10Request *QianchuanFileImageDeleteV10Request
}

func (r *ApiOpenApiV10QianchuanFileImageDeletePostRequest) QianchuanFileImageDeleteV10Request(qianchuanFileImageDeleteV10Request QianchuanFileImageDeleteV10Request) *ApiOpenApiV10QianchuanFileImageDeletePostRequest {
	r.qianchuanFileImageDeleteV10Request = &qianchuanFileImageDeleteV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanFileImageDeletePostRequest) Execute() (*QianchuanFileImageDeleteV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanFileImageDeletePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanFileImageDeletePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanFileImageDeletePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanFileImageDeletePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanFileImageDeletePost Method for OpenApiV10QianchuanFileImageDeletePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanFileImageDeletePostRequest
*/
func (a *QianchuanFileImageDeleteV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanFileImageDeletePostRequest {
	return &ApiOpenApiV10QianchuanFileImageDeletePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanFileImageDeleteV10Response
func (a *QianchuanFileImageDeleteV10ApiService) postExecute(r *ApiOpenApiV10QianchuanFileImageDeletePostRequest) (*QianchuanFileImageDeleteV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanFileImageDeleteV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/file/image/delete/"

	localVarHeaderParams := make(map[string]string)
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
	localVarPostBody = r.qianchuanFileImageDeleteV10Request
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
