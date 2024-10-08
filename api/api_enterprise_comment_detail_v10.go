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

// EnterpriseCommentDetailV10ApiService EnterpriseCommentDetailV10Api service
type EnterpriseCommentDetailV10ApiService service

type ApiOpenApiV10EnterpriseCommentDetailGetRequest struct {
	ctx         context.Context
	ApiService  *EnterpriseCommentDetailV10ApiService
	ccAccountId *int64
	commentId   *int64
	eDouyinId   *string
}

func (r *ApiOpenApiV10EnterpriseCommentDetailGetRequest) CcAccountId(ccAccountId int64) *ApiOpenApiV10EnterpriseCommentDetailGetRequest {
	r.ccAccountId = &ccAccountId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentDetailGetRequest) CommentId(commentId int64) *ApiOpenApiV10EnterpriseCommentDetailGetRequest {
	r.commentId = &commentId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentDetailGetRequest) EDouyinId(eDouyinId string) *ApiOpenApiV10EnterpriseCommentDetailGetRequest {
	r.eDouyinId = &eDouyinId
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentDetailGetRequest) Execute() (*EnterpriseCommentDetailV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10EnterpriseCommentDetailGetRequest) AccessToken(accessToken string) *ApiOpenApiV10EnterpriseCommentDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10EnterpriseCommentDetailGetRequest) WithLog(enable bool) *ApiOpenApiV10EnterpriseCommentDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10EnterpriseCommentDetailGet Method for OpenApiV10EnterpriseCommentDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10EnterpriseCommentDetailGetRequest
*/
func (a *EnterpriseCommentDetailV10ApiService) Get(ctx context.Context) *ApiOpenApiV10EnterpriseCommentDetailGetRequest {
	return &ApiOpenApiV10EnterpriseCommentDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnterpriseCommentDetailV10Response
func (a *EnterpriseCommentDetailV10ApiService) getExecute(r *ApiOpenApiV10EnterpriseCommentDetailGetRequest) (*EnterpriseCommentDetailV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EnterpriseCommentDetailV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/enterprise/comment/detail/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ccAccountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cc_account_id", r.ccAccountId)
	}
	if r.commentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comment_id", r.commentId)
	}
	if r.eDouyinId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "e_douyin_id", r.eDouyinId)
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
