/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

// StarUserGetAwemeAuthorIdV2ApiService StarUserGetAwemeAuthorIdV2Api service
type StarUserGetAwemeAuthorIdV2ApiService service

type ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest struct {
	ctx          context.Context
	ApiService   *StarUserGetAwemeAuthorIdV2ApiService
	starId       *int64
	authorStarId *int64
}

func (r *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest) StarId(starId int64) *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest) AuthorStarId(authorStarId int64) *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest {
	r.authorStarId = &authorStarId
	return r
}

func (r *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest) Execute() (*StarUserGetAwemeAuthorIdV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest) WithLog(enable bool) *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarUserGetAwemeAuthorIdGet Method for OpenApi2StarUserGetAwemeAuthorIdGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest
*/
func (a *StarUserGetAwemeAuthorIdV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest {
	return &ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarUserGetAwemeAuthorIdV2Response
func (a *StarUserGetAwemeAuthorIdV2ApiService) getExecute(r *ApiOpenApi2StarUserGetAwemeAuthorIdGetRequest) (*StarUserGetAwemeAuthorIdV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarUserGetAwemeAuthorIdV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/user/get_aweme_author_id/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.authorStarId == nil {
		return localVarReturnValue, nil, ReportError("authorStarId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "author_star_id", r.authorStarId)
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
