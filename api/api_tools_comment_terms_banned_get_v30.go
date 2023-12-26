/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
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

// ToolsCommentTermsBannedGetV30ApiService ToolsCommentTermsBannedGetV30Api service
type ToolsCommentTermsBannedGetV30ApiService service

type ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsCommentTermsBannedGetV30ApiService
	advertiserId *int64
	page         *int64
	pageSize     *int64
	awemeId      *string
	isApplyToAdv *bool
}

// 广告主id
func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 页码
func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) Page(page int64) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	r.pageSize = &pageSize
	return r
}

// 抖音号，当is_apply_to_adv不传或为false时，aweme_id生效
func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) AwemeId(awemeId string) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	r.awemeId = &awemeId
	return r
}

// 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) IsApplyToAdv(isApplyToAdv bool) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	r.isApplyToAdv = &isApplyToAdv
	return r
}

func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) Execute() (*ToolsCommentTermsBannedGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsCommentTermsBannedGetGet Method for OpenApiV30ToolsCommentTermsBannedGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest
*/
func (a *ToolsCommentTermsBannedGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest {
	return &ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsCommentTermsBannedGetV30Response
func (a *ToolsCommentTermsBannedGetV30ApiService) getExecute(r *ApiOpenApiV30ToolsCommentTermsBannedGetGetRequest) (*ToolsCommentTermsBannedGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsCommentTermsBannedGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/comment/terms_banned/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.awemeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	}
	if r.isApplyToAdv != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_apply_to_adv", r.isApplyToAdv)
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
