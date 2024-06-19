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

// EventManagerTrackUrlGetV2ApiService EventManagerTrackUrlGetV2Api service
type EventManagerTrackUrlGetV2ApiService service

type ApiOpenApi2EventManagerTrackUrlGetGetRequest struct {
	ctx               context.Context
	ApiService        *EventManagerTrackUrlGetV2ApiService
	advertiserId      *int64
	assetId           *int64
	downloadUrl       *string
	trackUrlGroupName *string
	trackUrlGroupId   *int64
	page              *int64
	pageSize          *int64
}

// 广告主ID
func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.advertiserId = &advertiserId
	return r
}

// 资产ID
func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) AssetId(assetId int64) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.assetId = &assetId
	return r
}

// 应用下载链接
func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) DownloadUrl(downloadUrl string) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.downloadUrl = &downloadUrl
	return r
}

// 监测链接组名称
func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) TrackUrlGroupName(trackUrlGroupName string) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.trackUrlGroupName = &trackUrlGroupName
	return r
}

// 监测链接组ID
func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) TrackUrlGroupId(trackUrlGroupId int64) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.trackUrlGroupId = &trackUrlGroupId
	return r
}

// 页码
func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) Page(page int64) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.page = &page
	return r
}

// 分页个数
func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) PageSize(pageSize int64) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) Execute() (*EventManagerTrackUrlGetV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) AccessToken(accessToken string) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) WithLog(enable bool) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2EventManagerTrackUrlGetGet Method for OpenApi2EventManagerTrackUrlGetGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2EventManagerTrackUrlGetGetRequest
*/
func (a *EventManagerTrackUrlGetV2ApiService) Get(ctx context.Context) *ApiOpenApi2EventManagerTrackUrlGetGetRequest {
	return &ApiOpenApi2EventManagerTrackUrlGetGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return EventManagerTrackUrlGetV2Response
func (a *EventManagerTrackUrlGetV2ApiService) getExecute(r *ApiOpenApi2EventManagerTrackUrlGetGetRequest) (*EventManagerTrackUrlGetV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EventManagerTrackUrlGetV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/event_manager/track_url/get/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.assetId == nil {
		return localVarReturnValue, nil, ReportError("assetId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId)
	if r.downloadUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "download_url", r.downloadUrl)
	}
	if r.trackUrlGroupName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "track_url_group_name", r.trackUrlGroupName)
	}
	if r.trackUrlGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "track_url_group_id", r.trackUrlGroupId)
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
