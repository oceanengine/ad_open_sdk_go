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

// FileVideoAdV2ApiService FileVideoAdV2Api service
type FileVideoAdV2ApiService service

type ApiOpenApi2FileVideoAdPostRequest struct {
	ctx            context.Context
	ApiService     *FileVideoAdV2ApiService
	advertiserId   *int64
	filename       *string
	isAigc         *bool
	isGuideVideo   *bool
	labels         *[]string
	uploadType     *FileVideoAdV2UploadType
	videoFile      *FormFileInfo
	videoSignature *string
	videoUrl       *string
}

// 广告主ID
func (r *ApiOpenApi2FileVideoAdPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileVideoAdPostRequest {
	r.advertiserId = &advertiserId
	return r
}

// 素材的文件名，可自定义素材名，不传择默认取文件名，最长255个字符 注：若同一素材已进行上传，重新上传不会改名
func (r *ApiOpenApi2FileVideoAdPostRequest) Filename(filename string) *ApiOpenApi2FileVideoAdPostRequest {
	r.filename = &filename
	return r
}

// 视频素材是否为AIGC生成
func (r *ApiOpenApi2FileVideoAdPostRequest) IsAigc(isAigc bool) *ApiOpenApi2FileVideoAdPostRequest {
	r.isAigc = &isAigc
	return r
}

// 上传视频场景是否是引导视频
func (r *ApiOpenApi2FileVideoAdPostRequest) IsGuideVideo(isGuideVideo bool) *ApiOpenApi2FileVideoAdPostRequest {
	r.isGuideVideo = &isGuideVideo
	return r
}

func (r *ApiOpenApi2FileVideoAdPostRequest) Labels(labels []string) *ApiOpenApi2FileVideoAdPostRequest {
	r.labels = &labels
	return r
}

func (r *ApiOpenApi2FileVideoAdPostRequest) UploadType(uploadType FileVideoAdV2UploadType) *ApiOpenApi2FileVideoAdPostRequest {
	r.uploadType = &uploadType
	return r
}

// 视频文件 允许格式：mp4、mpeg、3gp、avi（10s超时限制） upload_type为UPLOAD_BY_File必填
func (r *ApiOpenApi2FileVideoAdPostRequest) VideoFile(videoFile *FormFileInfo) *ApiOpenApi2FileVideoAdPostRequest {
	r.videoFile = videoFile
	return r
}

// 视频的md5值(用于服务端校验) upload_type为UPLOAD_BY_File必填
func (r *ApiOpenApi2FileVideoAdPostRequest) VideoSignature(videoSignature string) *ApiOpenApi2FileVideoAdPostRequest {
	r.videoSignature = &videoSignature
	return r
}

// 视频url地址，如http://xxx.xxx，upload_type为UPLOAD_BY_URL必填
func (r *ApiOpenApi2FileVideoAdPostRequest) VideoUrl(videoUrl string) *ApiOpenApi2FileVideoAdPostRequest {
	r.videoUrl = &videoUrl
	return r
}

func (r *ApiOpenApi2FileVideoAdPostRequest) Execute() (*FileVideoAdV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2FileVideoAdPostRequest) AccessToken(accessToken string) *ApiOpenApi2FileVideoAdPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileVideoAdPostRequest) WithLog(enable bool) *ApiOpenApi2FileVideoAdPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileVideoAdPost Method for OpenApi2FileVideoAdPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileVideoAdPostRequest
*/
func (a *FileVideoAdV2ApiService) Post(ctx context.Context) *ApiOpenApi2FileVideoAdPostRequest {
	return &ApiOpenApi2FileVideoAdPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileVideoAdV2Response
func (a *FileVideoAdV2ApiService) postExecute(r *ApiOpenApi2FileVideoAdPostRequest) (*FileVideoAdV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileVideoAdV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/video/ad/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.filename != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "filename", r.filename)
	}
	if r.isAigc != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_aigc", r.isAigc)
	}
	if r.isGuideVideo != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "is_guide_video", r.isGuideVideo)
	}
	if r.labels != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "labels", r.labels)
	}
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

	localVarHTTPResponse, err := a.client.call(r.ctx, req, &localVarReturnValue)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}
	return localVarReturnValue, localVarHTTPResponse, nil
}
