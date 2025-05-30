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

// BrandCampaignAddV30ApiService BrandCampaignAddV30Api service
type BrandCampaignAddV30ApiService service

type ApiOpenApiV30BrandCampaignAddPostRequest struct {
	ctx                        context.Context
	ApiService                 *BrandCampaignAddV30ApiService
	brandCampaignAddV30Request *BrandCampaignAddV30Request
}

func (r *ApiOpenApiV30BrandCampaignAddPostRequest) BrandCampaignAddV30Request(brandCampaignAddV30Request BrandCampaignAddV30Request) *ApiOpenApiV30BrandCampaignAddPostRequest {
	r.brandCampaignAddV30Request = &brandCampaignAddV30Request
	return r
}

func (r *ApiOpenApiV30BrandCampaignAddPostRequest) Execute() (*BrandCampaignAddV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30BrandCampaignAddPostRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandCampaignAddPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandCampaignAddPostRequest) WithLog(enable bool) *ApiOpenApiV30BrandCampaignAddPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandCampaignAddPost Method for OpenApiV30BrandCampaignAddPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandCampaignAddPostRequest
*/
func (a *BrandCampaignAddV30ApiService) Post(ctx context.Context) *ApiOpenApiV30BrandCampaignAddPostRequest {
	return &ApiOpenApiV30BrandCampaignAddPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandCampaignAddV30Response
func (a *BrandCampaignAddV30ApiService) postExecute(r *ApiOpenApiV30BrandCampaignAddPostRequest) (*BrandCampaignAddV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandCampaignAddV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/campaign/add/"

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
	localVarPostBody = r.brandCampaignAddV30Request
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
