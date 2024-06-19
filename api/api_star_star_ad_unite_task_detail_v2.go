/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
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

// StarStarAdUniteTaskDetailV2ApiService StarStarAdUniteTaskDetailV2Api service
type StarStarAdUniteTaskDetailV2ApiService service

type ApiOpenApi2StarStarAdUniteTaskDetailGetRequest struct {
	ctx           context.Context
	ApiService    *StarStarAdUniteTaskDetailV2ApiService
	starId        *int64
	demandId      *int64
	statStartDate *string
	statEndDate   *string
}

// 客户的星图id
func (r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) StarId(starId int64) *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest {
	r.starId = &starId
	return r
}

// 任务id
func (r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) DemandId(demandId int64) *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest {
	r.demandId = &demandId
	return r
}

// 查询起始时间，格式：yyyy-mm-dd ，只和安卓/iOS消耗、转化数、深度转化数相关
func (r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) StatStartDate(statStartDate string) *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest {
	r.statStartDate = &statStartDate
	return r
}

// 查询截止时间，格式：yyyy-mm-dd ，只和安卓/iOS消耗、转化数、深度转化数相关
func (r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) StatEndDate(statEndDate string) *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest {
	r.statEndDate = &statEndDate
	return r
}

func (r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) Execute() (*StarStarAdUniteTaskDetailV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) WithLog(enable bool) *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarStarAdUniteTaskDetailGet Method for OpenApi2StarStarAdUniteTaskDetailGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarStarAdUniteTaskDetailGetRequest
*/
func (a *StarStarAdUniteTaskDetailV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest {
	return &ApiOpenApi2StarStarAdUniteTaskDetailGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarStarAdUniteTaskDetailV2Response
func (a *StarStarAdUniteTaskDetailV2ApiService) getExecute(r *ApiOpenApi2StarStarAdUniteTaskDetailGetRequest) (*StarStarAdUniteTaskDetailV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarStarAdUniteTaskDetailV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/star_ad_unite_task/detail/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.demandId == nil {
		return localVarReturnValue, nil, ReportError("demandId is required and must be specified")
	}
	if r.statStartDate == nil {
		return localVarReturnValue, nil, ReportError("statStartDate is required and must be specified")
	}
	if r.statEndDate == nil {
		return localVarReturnValue, nil, ReportError("statEndDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "stat_start_date", r.statStartDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "stat_end_date", r.statEndDate)
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
