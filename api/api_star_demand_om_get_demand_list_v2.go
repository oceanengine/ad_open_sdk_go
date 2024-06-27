/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

// StarDemandOmGetDemandListV2ApiService StarDemandOmGetDemandListV2Api service
type StarDemandOmGetDemandListV2ApiService service

type ApiOpenApi2StarDemandOmGetDemandListGetRequest struct {
	ctx                     context.Context
	ApiService              *StarDemandOmGetDemandListV2ApiService
	starId                  *int64
	pageNo                  *int32
	pageSize                *int32
	universalSettlementType *StarDemandOmGetDemandListV2UniversalSettlementType
	microAppId              *string
	createStartTime         *int64
	createEndTime           *int64
}

// 客户星图ID
func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) StarId(starId int64) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.starId = &starId
	return r
}

// page
func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) PageNo(pageNo int32) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.pageNo = &pageNo
	return r
}

// limit
func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) PageSize(pageSize int32) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.pageSize = &pageSize
	return r
}

// 结算方式 枚举，付费分佣、广告分成
func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) UniversalSettlementType(universalSettlementType StarDemandOmGetDemandListV2UniversalSettlementType) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.universalSettlementType = &universalSettlementType
	return r
}

// 小程序ID
func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) MicroAppId(microAppId string) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.microAppId = &microAppId
	return r
}

// 任务开始时间
func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) CreateStartTime(createStartTime int64) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.createStartTime = &createStartTime
	return r
}

// 任务结束时间
func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) CreateEndTime(createEndTime int64) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.createEndTime = &createEndTime
	return r
}

func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) Execute() (*StarDemandOmGetDemandListV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) WithLog(enable bool) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandOmGetDemandListGet Method for OpenApi2StarDemandOmGetDemandListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandOmGetDemandListGetRequest
*/
func (a *StarDemandOmGetDemandListV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarDemandOmGetDemandListGetRequest {
	return &ApiOpenApi2StarDemandOmGetDemandListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandOmGetDemandListV2Response
func (a *StarDemandOmGetDemandListV2ApiService) getExecute(r *ApiOpenApi2StarDemandOmGetDemandListGetRequest) (*StarDemandOmGetDemandListV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandOmGetDemandListV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/om_get_demand_list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.pageNo == nil {
		return localVarReturnValue, nil, ReportError("pageNo is required and must be specified")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	if r.universalSettlementType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "universal_settlement_type", r.universalSettlementType)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_no", r.pageNo)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
	if r.microAppId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "micro_app_id", r.microAppId)
	}
	if r.createStartTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_start_time", r.createStartTime)
	}
	if r.createEndTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "create_end_time", r.createEndTime)
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
