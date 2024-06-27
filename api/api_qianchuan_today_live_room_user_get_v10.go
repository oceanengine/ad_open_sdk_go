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

// QianchuanTodayLiveRoomUserGetV10ApiService QianchuanTodayLiveRoomUserGetV10Api service
type QianchuanTodayLiveRoomUserGetV10ApiService service

type ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest struct {
	ctx          context.Context
	ApiService   *QianchuanTodayLiveRoomUserGetV10ApiService
	advertiserId *int64
	roomId       *int64
	actionEvent  *QianchuanTodayLiveRoomUserGetV10ActionEvent
	dimension    *[]*QianchuanTodayLiveRoomUserGetV10Dimension
	flowSource   *QianchuanTodayLiveRoomUserGetV10FlowSource
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) RoomId(roomId int64) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	r.roomId = &roomId
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) ActionEvent(actionEvent QianchuanTodayLiveRoomUserGetV10ActionEvent) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	r.actionEvent = &actionEvent
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) Dimension(dimension []*QianchuanTodayLiveRoomUserGetV10Dimension) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	r.dimension = &dimension
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) FlowSource(flowSource QianchuanTodayLiveRoomUserGetV10FlowSource) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	r.flowSource = &flowSource
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) Execute() (*QianchuanTodayLiveRoomUserGetV10Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanTodayLiveRoomUserGetGet Method for OpenApiV10QianchuanTodayLiveRoomUserGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest
*/
func (a *QianchuanTodayLiveRoomUserGetV10ApiService) Get(ctx context.Context) *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest {
	return &ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanTodayLiveRoomUserGetV10Response
func (a *QianchuanTodayLiveRoomUserGetV10ApiService) getExecute(r *ApiOpenApiV10QianchuanTodayLiveRoomUserGetGetRequest) (*QianchuanTodayLiveRoomUserGetV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanTodayLiveRoomUserGetV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/today_live/room/user/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.roomId == nil {
		return localVarReturnValue, nil, ReportError("roomId is required and must be specified")
	}
	if r.actionEvent == nil {
		return localVarReturnValue, nil, ReportError("actionEvent is required and must be specified")
	}
	if r.dimension == nil {
		return localVarReturnValue, nil, ReportError("dimension is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "room_id", r.roomId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "action_event", r.actionEvent)
	parameterAddToHeaderOrQuery(localVarQueryParams, "dimension", r.dimension)
	if r.flowSource != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flow_source", r.flowSource)
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
