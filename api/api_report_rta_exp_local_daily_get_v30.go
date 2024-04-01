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

// ReportRtaExpLocalDailyGetV30ApiService ReportRtaExpLocalDailyGetV30Api service
type ReportRtaExpLocalDailyGetV30ApiService service

type ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest struct {
	ctx          context.Context
	ApiService   *ReportRtaExpLocalDailyGetV30ApiService
	rtaId        *int64
	advertiserId *int64
	startDate    *string
	endDate      *string
	vid          *int64
	cusVid       *int64
	filtering    *ReportRtaExpLocalDailyGetV30Filtering
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) RtaId(rtaId int64) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.rtaId = &rtaId
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) StartDate(startDate string) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.startDate = &startDate
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) EndDate(endDate string) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.endDate = &endDate
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) Vid(vid int64) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.vid = &vid
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) CusVid(cusVid int64) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.cusVid = &cusVid
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) Filtering(filtering ReportRtaExpLocalDailyGetV30Filtering) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.filtering = &filtering
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) Execute() (*ReportRtaExpLocalDailyGetV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) WithLog(enable bool) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30ReportRtaExpLocalDailyGetGet Method for OpenApiV30ReportRtaExpLocalDailyGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest
*/
func (a *ReportRtaExpLocalDailyGetV30ApiService) Get(ctx context.Context) *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest {
	return &ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ReportRtaExpLocalDailyGetV30Response
func (a *ReportRtaExpLocalDailyGetV30ApiService) getExecute(r *ApiOpenApiV30ReportRtaExpLocalDailyGetGetRequest) (*ReportRtaExpLocalDailyGetV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ReportRtaExpLocalDailyGetV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/report/rta_exp_local_daily/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.rtaId == nil {
		return localVarReturnValue, nil, ReportError("rtaId is required and must be specified")
	}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.startDate == nil {
		return localVarReturnValue, nil, ReportError("startDate is required and must be specified")
	}
	if r.endDate == nil {
		return localVarReturnValue, nil, ReportError("endDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "rta_id", r.rtaId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.vid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vid", r.vid)
	}
	if r.cusVid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cus_vid", r.cusVid)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate)
	parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate)
	if r.filtering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filtering", r.filtering)
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
