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

// StarAttachmentUploadV2ApiService StarAttachmentUploadV2Api service
type StarAttachmentUploadV2ApiService service

type ApiOpenApi2StarAttachmentUploadPostRequest struct {
	ctx        context.Context
	ApiService *StarAttachmentUploadV2ApiService
	file       *FormFileInfo
	fileName   *string
	starId     *int64
}

// 上传的文件
func (r *ApiOpenApi2StarAttachmentUploadPostRequest) File(file *FormFileInfo) *ApiOpenApi2StarAttachmentUploadPostRequest {
	r.file = file
	return r
}

// 文件名
func (r *ApiOpenApi2StarAttachmentUploadPostRequest) FileName(fileName string) *ApiOpenApi2StarAttachmentUploadPostRequest {
	r.fileName = &fileName
	return r
}

// 星图客户ID
func (r *ApiOpenApi2StarAttachmentUploadPostRequest) StarId(starId int64) *ApiOpenApi2StarAttachmentUploadPostRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarAttachmentUploadPostRequest) Execute() (*StarAttachmentUploadV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarAttachmentUploadPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarAttachmentUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarAttachmentUploadPostRequest) WithLog(enable bool) *ApiOpenApi2StarAttachmentUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarAttachmentUploadPost Method for OpenApi2StarAttachmentUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarAttachmentUploadPostRequest
*/
func (a *StarAttachmentUploadV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarAttachmentUploadPostRequest {
	return &ApiOpenApi2StarAttachmentUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarAttachmentUploadV2Response
func (a *StarAttachmentUploadV2ApiService) postExecute(r *ApiOpenApi2StarAttachmentUploadPostRequest) (*StarAttachmentUploadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarAttachmentUploadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/attachment/upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, ReportError("file is required and must be specified")
	}
	if r.fileName == nil {
		return localVarReturnValue, nil, ReportError("fileName is required and must be specified")
	}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	if r.file != nil {
		formFiles["file"] = r.file
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "file_name", r.fileName)
	parameterAddToHeaderOrQuery(localVarFormParams, "star_id", r.starId)
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
