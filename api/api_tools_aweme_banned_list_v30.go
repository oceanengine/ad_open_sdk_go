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

// ToolsAwemeBannedListV30ApiService ToolsAwemeBannedListV30Api service
type ToolsAwemeBannedListV30ApiService service

type ApiOpenApiV30ToolsAwemeBannedListGetRequest struct {
	ctx             context.Context
	ApiService      *ToolsAwemeBannedListV30ApiService
	advertiserId    *int64
	bannedType      *ToolsAwemeBannedListV30BannedType
	awemeId         *string
	isApplyToAdv    *bool
	nicknameKeyword *string
	awemeUserId     *string
	page            *int64
	pageSize        *int64
}

// 广告主id
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 屏蔽类型， 允许值： CUSTOM_TYPE：自定义规则，根据昵称关键词屏蔽；AWEME_TYPE：根据抖音id屏蔽。
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) BannedType(bannedType ToolsAwemeBannedListV30BannedType) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.bannedType = &bannedType
	return r
}

// 抖音号，当is_apply_to_adv不传或为false时，aweme_id生效
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) AwemeId(awemeId string) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.awemeId = &awemeId
	return r
}

// 是否应用于当前账户，当is_apply_to_adv不传或为false时，aweme_id生效
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) IsApplyToAdv(isApplyToAdv bool) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.isApplyToAdv = &isApplyToAdv
	return r
}

// 昵称关键词，最大不超过20字
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) NicknameKeyword(nicknameKeyword string) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.nicknameKeyword = &nicknameKeyword
	return r
}

// 抖音用户id
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) AwemeUserId(awemeUserId string) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.awemeUserId = &awemeUserId
	return r
}

// 页码
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) Page(page int64) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.page = &page
	return r
}

// 页面大小
func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) PageSize(pageSize int64) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) Execute() (*ToolsAwemeBannedListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) WithLog(enable bool) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ToolsAwemeBannedListGet Method for OpenApiV30ToolsAwemeBannedListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ToolsAwemeBannedListGetRequest
*/
func (a *ToolsAwemeBannedListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ToolsAwemeBannedListGetRequest {
	return &ApiOpenApiV30ToolsAwemeBannedListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAwemeBannedListV30Response
func (a *ToolsAwemeBannedListV30ApiService) getExecute(r *ApiOpenApiV30ToolsAwemeBannedListGetRequest) (*ToolsAwemeBannedListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ToolsAwemeBannedListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/tools/aweme_banned/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.bannedType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "banned_type", r.bannedType)
	}
	if r.awemeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	}
	if r.isApplyToAdv != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_apply_to_adv", r.isApplyToAdv)
	}
	if r.nicknameKeyword != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nickname_keyword", r.nicknameKeyword)
	}
	if r.awemeUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_user_id", r.awemeUserId)
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
