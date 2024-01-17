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

// StarCampaignListV2ApiService StarCampaignListV2Api service
type StarCampaignListV2ApiService service

type ApiOpenApi2StarCampaignListGetRequest struct {
	ctx        context.Context
	ApiService *StarCampaignListV2ApiService
	starId     *int64
	searchName *string
	page       *int32
	limit      *int32
}

// 星图客户ID
func (r *ApiOpenApi2StarCampaignListGetRequest) StarId(starId int64) *ApiOpenApi2StarCampaignListGetRequest {
	r.starId = &starId
	return r
}

// 搜索名称 少于20字
func (r *ApiOpenApi2StarCampaignListGetRequest) SearchName(searchName string) *ApiOpenApi2StarCampaignListGetRequest {
	r.searchName = &searchName
	return r
}

// 分页页码
func (r *ApiOpenApi2StarCampaignListGetRequest) Page(page int32) *ApiOpenApi2StarCampaignListGetRequest {
	r.page = &page
	return r
}

// 页大小
func (r *ApiOpenApi2StarCampaignListGetRequest) Limit(limit int32) *ApiOpenApi2StarCampaignListGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApi2StarCampaignListGetRequest) Execute() (*StarCampaignListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarCampaignListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarCampaignListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarCampaignListGetRequest) WithLog(enable bool) *ApiOpenApi2StarCampaignListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarCampaignListGet Method for OpenApi2StarCampaignListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarCampaignListGetRequest
*/
func (a *StarCampaignListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarCampaignListGetRequest {
	return &ApiOpenApi2StarCampaignListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarCampaignListV2Response
func (a *StarCampaignListV2ApiService) getExecute(r *ApiOpenApi2StarCampaignListGetRequest) (*StarCampaignListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *StarCampaignListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/campaign/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	if r.searchName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_name", r.searchName)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
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
