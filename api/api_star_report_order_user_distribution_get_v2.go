/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
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

// StarReportOrderUserDistributionGetV2ApiService StarReportOrderUserDistributionGetV2Api service
type StarReportOrderUserDistributionGetV2ApiService service

type ApiOpenApi2StarReportOrderUserDistributionGetGetRequest struct {
	ctx             context.Context
	ApiService      *StarReportOrderUserDistributionGetV2ApiService
	fanType         *interface{}
	interactiveType *interface{}
	orderId         *int64
	starId          *int64
}

func (r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) FanType(fanType interface{}) *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest {
	r.fanType = &fanType
	return r
}

func (r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) InteractiveType(interactiveType interface{}) *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest {
	r.interactiveType = &interactiveType
	return r
}

func (r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) OrderId(orderId int64) *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest {
	r.orderId = &orderId
	return r
}

func (r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) StarId(starId int64) *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) Execute() (*StarReportOrderUserDistributionGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) WithLog(enable bool) *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarReportOrderUserDistributionGetGet Method for OpenApi2StarReportOrderUserDistributionGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarReportOrderUserDistributionGetGetRequest
*/
func (a *StarReportOrderUserDistributionGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest {
	return &ApiOpenApi2StarReportOrderUserDistributionGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarReportOrderUserDistributionGetV2Response
func (a *StarReportOrderUserDistributionGetV2ApiService) getExecute(r *ApiOpenApi2StarReportOrderUserDistributionGetGetRequest) (*StarReportOrderUserDistributionGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]FormFile
		localVarReturnValue *StarReportOrderUserDistributionGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/report/order_user_distribution/get/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fanType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fan_type", r.fanType)
	}
	if r.interactiveType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interactive_type", r.interactiveType)
	}
	if r.orderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_id", r.orderId)
	}
	if r.starId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
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
