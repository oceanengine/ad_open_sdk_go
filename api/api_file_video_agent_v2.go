/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
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

// FileVideoAgentV2ApiService FileVideoAgentV2Api service
type FileVideoAgentV2ApiService service

type ApiOpenApi2FileVideoAgentPostRequest struct {
	ctx            context.Context
	ApiService     *FileVideoAgentV2ApiService
	agentId        *int64
	fileName       *string
	isNeedAuth     *bool
	isAigc         *bool
	uploadType     *FileVideoAgentV2UploadType
	videoFile      *FormFileInfo
	videoSignature *string
	videoUrl       *string
}

// 代理商id
func (r *ApiOpenApi2FileVideoAgentPostRequest) AgentId(agentId int64) *ApiOpenApi2FileVideoAgentPostRequest {
	r.agentId = &agentId
	return r
}

// 视频上传名称
func (r *ApiOpenApi2FileVideoAgentPostRequest) FileName(fileName string) *ApiOpenApi2FileVideoAgentPostRequest {
	r.fileName = &fileName
	return r
}

// 标注是否允许授权(素材搬运)，目前只支持true
func (r *ApiOpenApi2FileVideoAgentPostRequest) IsNeedAuth(isNeedAuth bool) *ApiOpenApi2FileVideoAgentPostRequest {
	r.isNeedAuth = &isNeedAuth
	return r
}

// 标注是否为AIGC生成,默认false
func (r *ApiOpenApi2FileVideoAgentPostRequest) IsAigc(isAigc bool) *ApiOpenApi2FileVideoAgentPostRequest {
	r.isAigc = &isAigc
	return r
}

func (r *ApiOpenApi2FileVideoAgentPostRequest) UploadType(uploadType FileVideoAgentV2UploadType) *ApiOpenApi2FileVideoAgentPostRequest {
	r.uploadType = &uploadType
	return r
}

// 文件数据
func (r *ApiOpenApi2FileVideoAgentPostRequest) VideoFile(videoFile *FormFileInfo) *ApiOpenApi2FileVideoAgentPostRequest {
	r.videoFile = videoFile
	return r
}

// 视频的md5值(用于服务端校验) upload_type为UPLOAD_BY_File必填
func (r *ApiOpenApi2FileVideoAgentPostRequest) VideoSignature(videoSignature string) *ApiOpenApi2FileVideoAgentPostRequest {
	r.videoSignature = &videoSignature
	return r
}

// 视频URL地址，upload_type为UPLOAD_BY_URL必传 只支持连山tos地址
func (r *ApiOpenApi2FileVideoAgentPostRequest) VideoUrl(videoUrl string) *ApiOpenApi2FileVideoAgentPostRequest {
	r.videoUrl = &videoUrl
	return r
}

func (r *ApiOpenApi2FileVideoAgentPostRequest) Execute() (*FileVideoAgentV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2FileVideoAgentPostRequest) AccessToken(accessToken string) *ApiOpenApi2FileVideoAgentPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileVideoAgentPostRequest) WithLog(enable bool) *ApiOpenApi2FileVideoAgentPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoAgentPost Method for OpenApi2FileVideoAgentPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoAgentPostRequest
*/
func (a *FileVideoAgentV2ApiService) Post(ctx context.Context) *ApiOpenApi2FileVideoAgentPostRequest {
	return &ApiOpenApi2FileVideoAgentPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoAgentV2Response
func (a *FileVideoAgentV2ApiService) postExecute(r *ApiOpenApi2FileVideoAgentPostRequest) (*FileVideoAgentV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileVideoAgentV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/agent/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.agentId == nil {
		return localVarReturnValue, nil, ReportError("agentId is required and must be specified")
	}
	if r.fileName == nil {
		return localVarReturnValue, nil, ReportError("fileName is required and must be specified")
	}
	if r.isNeedAuth == nil {
		return localVarReturnValue, nil, ReportError("isNeedAuth is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "agent_id", r.agentId)
	parameterAddToHeaderOrQuery(localVarFormParams, "file_name", r.fileName)
	if r.isAigc != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_aigc", r.isAigc)
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "is_need_auth", r.isNeedAuth)
	if r.uploadType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload_type", r.uploadType)
	}
	if r.videoFile != nil {
		formFiles["video_file"] = r.videoFile
	}
	if r.videoSignature != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "video_signature", r.videoSignature)
	}
	if r.videoUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "video_url", r.videoUrl)
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
