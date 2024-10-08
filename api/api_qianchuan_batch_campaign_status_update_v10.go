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

// QianchuanBatchCampaignStatusUpdateV10ApiService QianchuanBatchCampaignStatusUpdateV10Api service
type QianchuanBatchCampaignStatusUpdateV10ApiService service

type ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest struct {
	ctx                                          context.Context
	ApiService                                   *QianchuanBatchCampaignStatusUpdateV10ApiService
	qianchuanBatchCampaignStatusUpdateV10Request *QianchuanBatchCampaignStatusUpdateV10Request
}

func (r *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest) QianchuanBatchCampaignStatusUpdateV10Request(qianchuanBatchCampaignStatusUpdateV10Request QianchuanBatchCampaignStatusUpdateV10Request) *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest {
	r.qianchuanBatchCampaignStatusUpdateV10Request = &qianchuanBatchCampaignStatusUpdateV10Request
	return r
}

func (r *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest) Execute() (*QianchuanBatchCampaignStatusUpdateV10Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest) AccessToken(accessToken string) *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest) WithLog(enable bool) *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV10QianchuanBatchCampaignStatusUpdatePost Method for OpenApiV10QianchuanBatchCampaignStatusUpdatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest
*/
func (a *QianchuanBatchCampaignStatusUpdateV10ApiService) Post(ctx context.Context) *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest {
	return &ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return QianchuanBatchCampaignStatusUpdateV10Response
func (a *QianchuanBatchCampaignStatusUpdateV10ApiService) postExecute(r *ApiOpenApiV10QianchuanBatchCampaignStatusUpdatePostRequest) (*QianchuanBatchCampaignStatusUpdateV10Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *QianchuanBatchCampaignStatusUpdateV10Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v1.0/qianchuan/batch_campaign_status/update/"

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
	localVarPostBody = r.qianchuanBatchCampaignStatusUpdateV10Request
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
