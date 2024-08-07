/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
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

// AdvertiserAvatarUploadV2ApiService AdvertiserAvatarUploadV2Api service
type AdvertiserAvatarUploadV2ApiService service

type ApiOpenApi2AdvertiserAvatarUploadPostRequest struct {
	ctx          context.Context
	ApiService   *AdvertiserAvatarUploadV2ApiService
	advertiserId *int64
	imageFile    *FormFileInfo
}

func (r *ApiOpenApi2AdvertiserAvatarUploadPostRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2AdvertiserAvatarUploadPostRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2AdvertiserAvatarUploadPostRequest) ImageFile(imageFile *FormFileInfo) *ApiOpenApi2AdvertiserAvatarUploadPostRequest {
	r.imageFile = imageFile
	return r
}

func (r *ApiOpenApi2AdvertiserAvatarUploadPostRequest) Execute() (*AdvertiserAvatarUploadV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2AdvertiserAvatarUploadPostRequest) AccessToken(accessToken string) *ApiOpenApi2AdvertiserAvatarUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2AdvertiserAvatarUploadPostRequest) WithLog(enable bool) *ApiOpenApi2AdvertiserAvatarUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2AdvertiserAvatarUploadPost Method for OpenApi2AdvertiserAvatarUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2AdvertiserAvatarUploadPostRequest
*/
func (a *AdvertiserAvatarUploadV2ApiService) Post(ctx context.Context) *ApiOpenApi2AdvertiserAvatarUploadPostRequest {
	return &ApiOpenApi2AdvertiserAvatarUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserAvatarUploadV2Response
func (a *AdvertiserAvatarUploadV2ApiService) postExecute(r *ApiOpenApi2AdvertiserAvatarUploadPostRequest) (*AdvertiserAvatarUploadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserAvatarUploadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/advertiser/avatar/upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.imageFile == nil {
		return localVarReturnValue, nil, ReportError("imageFile is required and must be specified")
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
	if r.imageFile != nil {
		formFiles["image_file"] = r.imageFile
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
