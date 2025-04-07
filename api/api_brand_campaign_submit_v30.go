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

// BrandCampaignSubmitV30ApiService BrandCampaignSubmitV30Api service
type BrandCampaignSubmitV30ApiService service

type ApiOpenApiV30BrandCampaignSubmitPostRequest struct {
	ctx                           context.Context
	ApiService                    *BrandCampaignSubmitV30ApiService
	brandCampaignSubmitV30Request *BrandCampaignSubmitV30Request
}

func (r *ApiOpenApiV30BrandCampaignSubmitPostRequest) BrandCampaignSubmitV30Request(brandCampaignSubmitV30Request BrandCampaignSubmitV30Request) *ApiOpenApiV30BrandCampaignSubmitPostRequest {
	r.brandCampaignSubmitV30Request = &brandCampaignSubmitV30Request
	return r
}

func (r *ApiOpenApiV30BrandCampaignSubmitPostRequest) Execute() (*BrandCampaignSubmitV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30BrandCampaignSubmitPostRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandCampaignSubmitPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandCampaignSubmitPostRequest) WithLog(enable bool) *ApiOpenApiV30BrandCampaignSubmitPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandCampaignSubmitPost Method for OpenApiV30BrandCampaignSubmitPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandCampaignSubmitPostRequest
*/
func (a *BrandCampaignSubmitV30ApiService) Post(ctx context.Context) *ApiOpenApiV30BrandCampaignSubmitPostRequest {
	return &ApiOpenApiV30BrandCampaignSubmitPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandCampaignSubmitV30Response
func (a *BrandCampaignSubmitV30ApiService) postExecute(r *ApiOpenApiV30BrandCampaignSubmitPostRequest) (*BrandCampaignSubmitV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandCampaignSubmitV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/campaign/submit/"

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
	localVarPostBody = r.brandCampaignSubmitV30Request
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
