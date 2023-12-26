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

// QianchuanUniPromotionListV10ApiService QianchuanUniPromotionListV10Api service
type QianchuanUniPromotionListV10ApiService service

type ApiOpenApiV10QianchuanUniPromotionListGetRequest struct {
	ctx           context.Context
	ApiService    *QianchuanUniPromotionListV10ApiService
	advertiserId  *int64
	startTime     *string
	endTime       *string
	marketingGoal *QianchuanUniPromotionListV10MarketingGoal
	fields        *[]*QianchuanUniPromotionListV10Fields
	orderType     *QianchuanUniPromotionListV10OrderType
	orderField    *QianchuanUniPromotionListV10OrderField
	page          *int32
	pageSize      *QianchuanUniPromotionListV10PageSize
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) StartTime(startTime string) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) EndTime(endTime string) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) MarketingGoal(marketingGoal QianchuanUniPromotionListV10MarketingGoal) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.marketingGoal = &marketingGoal
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) Fields(fields []*QianchuanUniPromotionListV10Fields) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) OrderType(orderType QianchuanUniPromotionListV10OrderType) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) OrderField(orderField QianchuanUniPromotionListV10OrderField) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) Page(page int32) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) PageSize(pageSize QianchuanUniPromotionListV10PageSize) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) Execute() (*QianchuanUniPromotionListV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanUniPromotionListGet Method for OpenApiV10QianchuanUniPromotionListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanUniPromotionListGetRequest
*/
func (a *QianchuanUniPromotionListV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanUniPromotionListGetRequest {
	return &ApiOpenApiV10QianchuanUniPromotionListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanUniPromotionListV10Response
func (a *QianchuanUniPromotionListV10ApiService) getExecute(r *ApiOpenApiV10QianchuanUniPromotionListGetRequest) (*QianchuanUniPromotionListV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *QianchuanUniPromotionListV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/uni_promotion/list/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}
	if r.marketingGoal == nil {
		return localVarReturnValue, nil, ReportError("marketingGoal is required and must be specified")
	}
	if r.fields == nil {
		return localVarReturnValue, nil, ReportError("fields is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_goal", r.marketingGoal)
	parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
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
