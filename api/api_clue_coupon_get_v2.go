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

// ClueCouponGetV2ApiService ClueCouponGetV2Api service
type ClueCouponGetV2ApiService service

type ApiOpenApi2ClueCouponGetGetRequest struct {
	ctx           context.Context
	ApiService    *ClueCouponGetV2ApiService
	activityIds   *[]int64
	activityTypes *[]*ClueCouponGetV2ActivityTypes
	advertiserId  *int64
	endTime       **string
	isDel         *ClueCouponGetV2IsDel
	page          *int64
	pageSize      *int64
	startTime     **string
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) ActivityIds(activityIds []int64) *ApiOpenApi2ClueCouponGetGetRequest {
	r.activityIds = &activityIds
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) ActivityTypes(activityTypes []*ClueCouponGetV2ActivityTypes) *ApiOpenApi2ClueCouponGetGetRequest {
	r.activityTypes = &activityTypes
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ClueCouponGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) EndTime(endTime *string) *ApiOpenApi2ClueCouponGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) IsDel(isDel ClueCouponGetV2IsDel) *ApiOpenApi2ClueCouponGetGetRequest {
	r.isDel = &isDel
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) Page(page int64) *ApiOpenApi2ClueCouponGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ClueCouponGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) StartTime(startTime *string) *ApiOpenApi2ClueCouponGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) Execute() (*ClueCouponGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ClueCouponGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueCouponGetGetRequest) WithLog(enable bool) *ApiOpenApi2ClueCouponGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueCouponGetGet Method for OpenApi2ClueCouponGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueCouponGetGetRequest
*/
func (a *ClueCouponGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ClueCouponGetGetRequest {
	return &ApiOpenApi2ClueCouponGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueCouponGetV2Response
func (a *ClueCouponGetV2ApiService) getExecute(r *ApiOpenApi2ClueCouponGetGetRequest) (*ClueCouponGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *ClueCouponGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/coupon/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.activityIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activity_ids", r.activityIds)
	}
	if r.activityTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activity_types", r.activityTypes)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.isDel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_del", r.isDel)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
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
