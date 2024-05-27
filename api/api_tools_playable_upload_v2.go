/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
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

// ToolsPlayableUploadV2ApiService ToolsPlayableUploadV2Api service
type ToolsPlayableUploadV2ApiService service

type ApiOpenApi2ToolsPlayableUploadPostRequest struct {
	ctx             context.Context
	ApiService      *ToolsPlayableUploadV2ApiService
	advertiserId    *int64
	playablePackage *FormFileInfo
}

func (r *ApiOpenApi2ToolsPlayableUploadPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsPlayableUploadPostRequest {
	r.advertiserId = &advertiserId
	return r
}

// 试玩素材文件 【格式说明】： 1.格式为zip，大小不超过3M 2.已接入穿山甲JS-SDK，并且调用方式为window.openAppStore(); 3.主html文件需命名为index，且位于一级目录中 4.试玩广告播放方向字段应存储于config.json文件中，位于一级目录中，取值为0,1,2 5.文件名称仅支持大小写字母、数字、减号和下划线 6.素材中不允许使用 mraid.js 格式 7.素材不允许通过外部网络加载动态素材 8.素材中不允许包含JS 重定向 9.素材不允许发出 HTTP 请求 10.素材中不允许存在内容为空的文件
func (r *ApiOpenApi2ToolsPlayableUploadPostRequest) PlayablePackage(playablePackage *FormFileInfo) *ApiOpenApi2ToolsPlayableUploadPostRequest {
	r.playablePackage = playablePackage
	return r
}

func (r *ApiOpenApi2ToolsPlayableUploadPostRequest) Execute() (*ToolsPlayableUploadV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsPlayableUploadPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPlayableUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPlayableUploadPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsPlayableUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPlayableUploadPost Method for OpenApi2ToolsPlayableUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPlayableUploadPostRequest
*/
func (a *ToolsPlayableUploadV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsPlayableUploadPostRequest {
	return &ApiOpenApi2ToolsPlayableUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPlayableUploadV2Response
func (a *ToolsPlayableUploadV2ApiService) postExecute(r *ApiOpenApi2ToolsPlayableUploadPostRequest) (*ToolsPlayableUploadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPlayableUploadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/playable/upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if *r.advertiserId < 1 {
		return localVarReturnValue, nil, ReportError("advertiserId must be greater than 1")
	}
	if *r.advertiserId > -9223372036854775616 {
		return localVarReturnValue, nil, ReportError("advertiserId must be less than -9223372036854775616")
	}
	if r.playablePackage == nil {
		return localVarReturnValue, nil, ReportError("playablePackage is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.playablePackage != nil {
		formFiles["playable_package"] = r.playablePackage
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
