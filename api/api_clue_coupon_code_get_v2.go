/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

// ClueCouponCodeGetV2ApiService ClueCouponCodeGetV2Api service
type ClueCouponCodeGetV2ApiService service

type ApiOpenApi2ClueCouponCodeGetGetRequest struct {
	ctx          context.Context
	ApiService   *ClueCouponCodeGetV2ApiService
	activityId   *int64
	advertiserId *int64
	couponId     *int64
	endTime      **string
	page         *int64
	pageSize     *int64
	resourceId   *int64
	startTime    **string
	status       *ClueCouponCodeGetV2Status
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) ActivityId(activityId int64) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.activityId = &activityId
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) CouponId(couponId int64) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.couponId = &couponId
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) EndTime(endTime *string) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) Page(page int64) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) ResourceId(resourceId int64) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.resourceId = &resourceId
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) StartTime(startTime *string) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) Status(status ClueCouponCodeGetV2Status) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) Execute() (*ClueCouponCodeGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ClueCouponCodeGetGetRequest) WithLog(enable bool) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ClueCouponCodeGetGet Method for OpenApi2ClueCouponCodeGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ClueCouponCodeGetGetRequest
*/
func (a *ClueCouponCodeGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ClueCouponCodeGetGetRequest {
	return &ApiOpenApi2ClueCouponCodeGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ClueCouponCodeGetV2Response
func (a *ClueCouponCodeGetV2ApiService) getExecute(r *ApiOpenApi2ClueCouponCodeGetGetRequest) (*ClueCouponCodeGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ClueCouponCodeGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/clue/coupon/code/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.activityId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "activity_id", r.activityId)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.couponId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "coupon_id", r.couponId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	}
	if r.resourceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resource_id", r.resourceId)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
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
