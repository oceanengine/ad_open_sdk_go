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

// BrandCampaignRevokeModifyV30ApiService BrandCampaignRevokeModifyV30Api service
type BrandCampaignRevokeModifyV30ApiService service

type ApiOpenApiV30BrandCampaignRevokeModifyPostRequest struct {
	ctx                                 context.Context
	ApiService                          *BrandCampaignRevokeModifyV30ApiService
	brandCampaignRevokeModifyV30Request *BrandCampaignRevokeModifyV30Request
}

func (r *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest) BrandCampaignRevokeModifyV30Request(brandCampaignRevokeModifyV30Request BrandCampaignRevokeModifyV30Request) *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest {
	r.brandCampaignRevokeModifyV30Request = &brandCampaignRevokeModifyV30Request
	return r
}

func (r *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest) Execute() (*BrandCampaignRevokeModifyV30Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest) AccessToken(accessToken string) *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest) WithLog(enable bool) *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30BrandCampaignRevokeModifyPost Method for OpenApiV30BrandCampaignRevokeModifyPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30BrandCampaignRevokeModifyPostRequest
*/
func (a *BrandCampaignRevokeModifyV30ApiService) Post(ctx context.Context) *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest {
	return &ApiOpenApiV30BrandCampaignRevokeModifyPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BrandCampaignRevokeModifyV30Response
func (a *BrandCampaignRevokeModifyV30ApiService) postExecute(r *ApiOpenApiV30BrandCampaignRevokeModifyPostRequest) (*BrandCampaignRevokeModifyV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *BrandCampaignRevokeModifyV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/brand/campaign/revoke_modify/"

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
	localVarPostBody = r.brandCampaignRevokeModifyV30Request
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
