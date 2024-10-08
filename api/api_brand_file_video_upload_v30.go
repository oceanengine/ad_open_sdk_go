/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// BrandFileVideoUploadV30ApiService BrandFileVideoUploadV30Api service
type BrandFileVideoUploadV30ApiService service

type ApiOpenApiV30BrandFileVideoUploadPostRequest struct {
	ctx          context.Context
	ApiService   *BrandFileVideoUploadV30ApiService
	advertiserId *int64
	videoFile    *FormFileInfo
}

// 广告主ID
func (r *ApiOpenApiV30BrandFileVideoUploadPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30BrandFileVideoUploadPostRequest {
	r.advertiserId = &advertiserId
	return r
}

// 视频文件
func (r *ApiOpenApiV30BrandFileVideoUploadPostRequest) VideoFile(videoFile *FormFileInfo) *ApiOpenApiV30BrandFileVideoUploadPostRequest {
	r.videoFile = videoFile
	return r
}

func (r *ApiOpenApiV30BrandFileVideoUploadPostRequest) Execute() (*BrandFileVideoUploadV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30BrandFileVideoUploadPostRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandFileVideoUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandFileVideoUploadPostRequest) WithLog(enable bool) *ApiOpenApiV30BrandFileVideoUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandFileVideoUploadPost Method for OpenApiV30BrandFileVideoUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandFileVideoUploadPostRequest
*/
func (a *BrandFileVideoUploadV30ApiService) Post(ctx context.Context) *ApiOpenApiV30BrandFileVideoUploadPostRequest {
	return &ApiOpenApiV30BrandFileVideoUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandFileVideoUploadV30Response
func (a *BrandFileVideoUploadV30ApiService) postExecute(r *ApiOpenApiV30BrandFileVideoUploadPostRequest) (*BrandFileVideoUploadV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandFileVideoUploadV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/file/video/upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.videoFile == nil {
		return localVarReturnValue, nil, ReportError("videoFile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.videoFile != nil {
		formFiles["video_file"] = r.videoFile
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
