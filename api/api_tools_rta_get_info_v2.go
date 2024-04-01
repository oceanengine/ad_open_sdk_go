/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
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

// ToolsRtaGetInfoV2ApiService ToolsRtaGetInfoV2Api service
type ToolsRtaGetInfoV2ApiService service

type ApiOpenApi2ToolsRtaGetInfoGetRequest struct {
	ctx          context.Context
	ApiService   *ToolsRtaGetInfoV2ApiService
	advertiserId *int64
	campaignId   *int64
	accountType  *ToolsRtaGetInfoV2AccountType
}

// 广告账户id
func (r *ApiOpenApi2ToolsRtaGetInfoGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsRtaGetInfoGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 广告组id，若传入，则拉取的是组维度的RTA策略
func (r *ApiOpenApi2ToolsRtaGetInfoGetRequest) CampaignId(campaignId int64) *ApiOpenApi2ToolsRtaGetInfoGetRequest {
	r.campaignId = &campaignId
	return r
}

// 账户类型，默认值为AD，允许值： AD 表示入参advertiser_id为巨量广告账户ID STAR 表示入参advertiser_id为星图账户ID
func (r *ApiOpenApi2ToolsRtaGetInfoGetRequest) AccountType(accountType ToolsRtaGetInfoV2AccountType) *ApiOpenApi2ToolsRtaGetInfoGetRequest {
	r.accountType = &accountType
	return r
}

func (r *ApiOpenApi2ToolsRtaGetInfoGetRequest) Execute() (*ToolsRtaGetInfoV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsRtaGetInfoGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsRtaGetInfoGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsRtaGetInfoGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsRtaGetInfoGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsRtaGetInfoGet Method for OpenApi2ToolsRtaGetInfoGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsRtaGetInfoGetRequest
*/
func (a *ToolsRtaGetInfoV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsRtaGetInfoGetRequest {
	return &ApiOpenApi2ToolsRtaGetInfoGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsRtaGetInfoV2Response
func (a *ToolsRtaGetInfoV2ApiService) getExecute(r *ApiOpenApi2ToolsRtaGetInfoGetRequest) (*ToolsRtaGetInfoV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsRtaGetInfoV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/rta/get_info/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.campaignId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "campaign_id", r.campaignId)
	}
	if r.accountType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType)
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
