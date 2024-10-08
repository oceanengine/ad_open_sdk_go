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

// AdvertiserQualificationSubmitV30ApiService AdvertiserQualificationSubmitV30Api service
type AdvertiserQualificationSubmitV30ApiService service

type ApiOpenApiV30AdvertiserQualificationSubmitPostRequest struct {
	ctx                                     context.Context
	ApiService                              *AdvertiserQualificationSubmitV30ApiService
	advertiserQualificationSubmitV30Request *AdvertiserQualificationSubmitV30Request
}

func (r *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest) AdvertiserQualificationSubmitV30Request(advertiserQualificationSubmitV30Request AdvertiserQualificationSubmitV30Request) *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest {
	r.advertiserQualificationSubmitV30Request = &advertiserQualificationSubmitV30Request
	return r
}

func (r *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest) Execute() (*AdvertiserQualificationSubmitV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest) AccessToken(accessToken string) *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest) WithLog(enable bool) *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdvertiserQualificationSubmitPost Method for OpenApiV30AdvertiserQualificationSubmitPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdvertiserQualificationSubmitPostRequest
*/
func (a *AdvertiserQualificationSubmitV30ApiService) Post(ctx context.Context) *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest {
	return &ApiOpenApiV30AdvertiserQualificationSubmitPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserQualificationSubmitV30Response
func (a *AdvertiserQualificationSubmitV30ApiService) postExecute(r *ApiOpenApiV30AdvertiserQualificationSubmitPostRequest) (*AdvertiserQualificationSubmitV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserQualificationSubmitV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/advertiser/qualification/submit/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// body params
	localVarPostBody = r.advertiserQualificationSubmitV30Request
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
