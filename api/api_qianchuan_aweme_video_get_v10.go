/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// QianchuanAwemeVideoGetV10ApiService QianchuanAwemeVideoGetV10Api service
type QianchuanAwemeVideoGetV10ApiService service

type ApiOpenApiV10QianchuanAwemeVideoGetGetRequest struct {
	ctx           context.Context
	ApiService    *QianchuanAwemeVideoGetV10ApiService
	advertiserId  *int64
	awemeId       *int64
	marketingGoal *QianchuanAwemeVideoGetV10MarketingGoal
	cursor        *int64
	count         *int64
}

func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 需要拉取视频的视频作者uid
func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) AwemeId(awemeId int64) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	r.awemeId = &awemeId
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) MarketingGoal(marketingGoal QianchuanAwemeVideoGetV10MarketingGoal) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) Cursor(cursor int64) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	r.cursor = &cursor
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) Count(count int64) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	r.count = &count
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) Execute() (*QianchuanAwemeVideoGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanAwemeVideoGetGet Method for OpenApiV10QianchuanAwemeVideoGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanAwemeVideoGetGetRequest
*/
func (a *QianchuanAwemeVideoGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest {
	return &ApiOpenApiV10QianchuanAwemeVideoGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanAwemeVideoGetV10Response
func (a *QianchuanAwemeVideoGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanAwemeVideoGetGetRequest) (*QianchuanAwemeVideoGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanAwemeVideoGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/aweme/video/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.awemeId == nil {
		return localVarReturnValue, nil, ReportError("awemeId is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "aweme_id", r.awemeId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor)
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count)
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
