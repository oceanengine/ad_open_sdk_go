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

// FileImageAdvertiserV2ApiService FileImageAdvertiserV2Api service
type FileImageAdvertiserV2ApiService service

type ApiOpenApi2FileImageAdvertiserPostRequest struct {
	ctx            context.Context
	ApiService     *FileImageAdvertiserV2ApiService
	advertiserId   *int64
	uploadTo       *FileImageAdvertiserV2UploadTo
	uploadType     *FileImageAdvertiserV2UploadType
	imageFile      *FormFileInfo
	imageSignature *string
	imageUrl       *string
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2FileImageAdvertiserPostRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) UploadTo(uploadTo FileImageAdvertiserV2UploadTo) *ApiOpenApi2FileImageAdvertiserPostRequest {
	r.uploadTo = &uploadTo
	return r
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) UploadType(uploadType FileImageAdvertiserV2UploadType) *ApiOpenApi2FileImageAdvertiserPostRequest {
	r.uploadType = &uploadType
	return r
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) ImageFile(imageFile *FormFileInfo) *ApiOpenApi2FileImageAdvertiserPostRequest {
	r.imageFile = imageFile
	return r
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) ImageSignature(imageSignature string) *ApiOpenApi2FileImageAdvertiserPostRequest {
	r.imageSignature = &imageSignature
	return r
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) ImageUrl(imageUrl string) *ApiOpenApi2FileImageAdvertiserPostRequest {
	r.imageUrl = &imageUrl
	return r
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) Execute() (*FileImageAdvertiserV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) AccessToken(accessToken string) *ApiOpenApi2FileImageAdvertiserPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2FileImageAdvertiserPostRequest) WithLog(enable bool) *ApiOpenApi2FileImageAdvertiserPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2FileImageAdvertiserPost Method for OpenApi2FileImageAdvertiserPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2FileImageAdvertiserPostRequest
*/
func (a *FileImageAdvertiserV2ApiService) Post(ctx context.Context) *ApiOpenApi2FileImageAdvertiserPostRequest {
	return &ApiOpenApi2FileImageAdvertiserPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FileImageAdvertiserV2Response
func (a *FileImageAdvertiserV2ApiService) postExecute(r *ApiOpenApi2FileImageAdvertiserPostRequest) (*FileImageAdvertiserV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *FileImageAdvertiserV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/file/image/advertiser/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.uploadTo == nil {
		return localVarReturnValue, nil, ReportError("uploadTo is required and must be specified")
	}
	if r.uploadType == nil {
		return localVarReturnValue, nil, ReportError("uploadType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	parameterAddToHeaderOrQuery(localVarFormParams, "advertiser_id", r.advertiserId)
	if r.imageFile != nil {
		formFiles["image_file"] = r.imageFile
	}
	if r.imageSignature != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "image_signature", r.imageSignature)
	}
	if r.imageUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "image_url", r.imageUrl)
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "upload_to", r.uploadTo)
	parameterAddToHeaderOrQuery(localVarFormParams, "upload_type", r.uploadType)
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
