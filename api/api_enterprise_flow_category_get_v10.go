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

// EnterpriseFlowCategoryGetV10ApiService EnterpriseFlowCategoryGetV10Api service
type EnterpriseFlowCategoryGetV10ApiService service

type ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest struct {
	ctx           context.Context
	ApiService    *EnterpriseFlowCategoryGetV10ApiService
	advertiserId  *int64
	endTime       **string
	fields        *[]string
	filter        *EnterpriseFlowCategoryGetV10Filter
	lastEndTime   **string
	lastStartTime **string
	limit         *int64
	offset        *int64
	orderField    *string
	orderType     *EnterpriseFlowCategoryGetV10OrderType
	ratioFields   *[]string
	startTime     **string
	timeDimension *EnterpriseFlowCategoryGetV10TimeDimension
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) EndTime(endTime *string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) Fields(fields []string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) Filter(filter EnterpriseFlowCategoryGetV10Filter) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) LastEndTime(lastEndTime *string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.lastEndTime = &lastEndTime
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) LastStartTime(lastStartTime *string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.lastStartTime = &lastStartTime
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) Limit(limit int64) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) Offset(offset int64) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.offset = &offset
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) OrderField(orderField string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) OrderType(orderType EnterpriseFlowCategoryGetV10OrderType) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) RatioFields(ratioFields []string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.ratioFields = &ratioFields
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) StartTime(startTime *string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) TimeDimension(timeDimension EnterpriseFlowCategoryGetV10TimeDimension) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.timeDimension = &timeDimension
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) Execute() (*EnterpriseFlowCategoryGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) WithLog(enable bool) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10EnterpriseFlowCategoryGetGet Method for OpenApiV10EnterpriseFlowCategoryGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest
*/
func (a *EnterpriseFlowCategoryGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest {
	return &ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnterpriseFlowCategoryGetV10Response
func (a *EnterpriseFlowCategoryGetV10ApiService) getExecute(r *ApiOpenApiV10EnterpriseFlowCategoryGetGetRequest) (*EnterpriseFlowCategoryGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EnterpriseFlowCategoryGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/enterprise/flow/category/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields)
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter)
	}
	if r.lastEndTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_end_time", r.lastEndTime)
	}
	if r.lastStartTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "last_start_time", r.lastStartTime)
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit)
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset)
	}
	if r.orderField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_field", r.orderField)
	}
	if r.orderType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_type", r.orderType)
	}
	if r.ratioFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ratio_fields", r.ratioFields)
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	}
	if r.timeDimension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time_dimension", r.timeDimension)
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
