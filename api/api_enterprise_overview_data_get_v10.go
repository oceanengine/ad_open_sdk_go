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

// EnterpriseOverviewDataGetV10ApiService EnterpriseOverviewDataGetV10Api service
type EnterpriseOverviewDataGetV10ApiService service

type ApiOpenApiV10EnterpriseOverviewDataGetGetRequest struct {
	ctx           context.Context
	ApiService    *EnterpriseOverviewDataGetV10ApiService
	advertiserId  *int64
	endTime       **string
	fields        *[]string
	filter        *EnterpriseOverviewDataGetV10Filter
	lastEndTime   **string
	lastStartTime **string
	limit         *int64
	offset        *int64
	orderField    *string
	orderType     *EnterpriseOverviewDataGetV10OrderType
	ratioFields   *[]string
	startTime     **string
	timeDimension *EnterpriseOverviewDataGetV10TimeDimension
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) EndTime(endTime *string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) Fields(fields []string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.fields = &fields
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) Filter(filter EnterpriseOverviewDataGetV10Filter) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.filter = &filter
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) LastEndTime(lastEndTime *string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.lastEndTime = &lastEndTime
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) LastStartTime(lastStartTime *string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.lastStartTime = &lastStartTime
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) Limit(limit int64) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.limit = &limit
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) Offset(offset int64) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.offset = &offset
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) OrderField(orderField string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.orderField = &orderField
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) OrderType(orderType EnterpriseOverviewDataGetV10OrderType) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.orderType = &orderType
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) RatioFields(ratioFields []string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.ratioFields = &ratioFields
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) StartTime(startTime *string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.startTime = &startTime
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) TimeDimension(timeDimension EnterpriseOverviewDataGetV10TimeDimension) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.timeDimension = &timeDimension
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) Execute() (*EnterpriseOverviewDataGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) WithLog(enable bool) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10EnterpriseOverviewDataGetGet Method for OpenApiV10EnterpriseOverviewDataGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10EnterpriseOverviewDataGetGetRequest
*/
func (a *EnterpriseOverviewDataGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest {
	return &ApiOpenApiV10EnterpriseOverviewDataGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EnterpriseOverviewDataGetV10Response
func (a *EnterpriseOverviewDataGetV10ApiService) getExecute(r *ApiOpenApiV10EnterpriseOverviewDataGetGetRequest) (*EnterpriseOverviewDataGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EnterpriseOverviewDataGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/enterprise/overview/data/get/"

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
