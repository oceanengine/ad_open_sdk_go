/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
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

// ReportRtaCusExpGetV2ApiService ReportRtaCusExpGetV2Api service
type ReportRtaCusExpGetV2ApiService service

type ApiOpenApi2ReportRtaCusExpGetGetRequest struct {
	ctx            context.Context
	ApiService     *ReportRtaCusExpGetV2ApiService
	advertiserId   *int64
	rtaInterfaceId *int64
	rtaId          *int64
	rtaVid         *string
	startTime      *string
	endTime        *string
}

// 广告主id
func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// rta客户接口id
func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) RtaInterfaceId(rtaInterfaceId int64) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	r.rtaInterfaceId = &rtaInterfaceId
	return r
}

// rta_id
func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) RtaId(rtaId int64) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	r.rtaId = &rtaId
	return r
}

// rta反向联合实验id
func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) RtaVid(rtaVid string) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	r.rtaVid = &rtaVid
	return r
}

// 开始时间，天级别（格式20240101），月份和日期必须为两位数
func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) StartTime(startTime string) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	r.startTime = &startTime
	return r
}

// 结束时间，天级别（格式20240101），月份和日期必须为两位数
func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) EndTime(endTime string) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	r.endTime = &endTime
	return r
}

func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) Execute() (*ReportRtaCusExpGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ReportRtaCusExpGetGetRequest) WithLog(enable bool) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ReportRtaCusExpGetGet Method for OpenApi2ReportRtaCusExpGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ReportRtaCusExpGetGetRequest
*/
func (a *ReportRtaCusExpGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2ReportRtaCusExpGetGetRequest {
	return &ApiOpenApi2ReportRtaCusExpGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportRtaCusExpGetV2Response
func (a *ReportRtaCusExpGetV2ApiService) getExecute(r *ApiOpenApi2ReportRtaCusExpGetGetRequest) (*ReportRtaCusExpGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportRtaCusExpGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/report/rta_cus_exp/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.rtaInterfaceId == nil {
		return localVarReturnValue, nil, ReportError("rtaInterfaceId is required and must be specified")
	}
	if r.rtaId == nil {
		return localVarReturnValue, nil, ReportError("rtaId is required and must be specified")
	}
	if r.rtaVid == nil {
		return localVarReturnValue, nil, ReportError("rtaVid is required and must be specified")
	}
	if r.startTime == nil {
		return localVarReturnValue, nil, ReportError("startTime is required and must be specified")
	}
	if r.endTime == nil {
		return localVarReturnValue, nil, ReportError("endTime is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "rta_interface_id", r.rtaInterfaceId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "rta_id", r.rtaId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "rta_vid", r.rtaVid)
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_time", r.startTime)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_time", r.endTime)
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