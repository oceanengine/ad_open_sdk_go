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
	"strings"

	"github.com/oceanengine/ad_open_sdk_go/config"
	. "github.com/oceanengine/ad_open_sdk_go/models"
)

// EventManagerAssetsCreateV2ApiService EventManagerAssetsCreateV2Api service
type EventManagerAssetsCreateV2ApiService service

type ApiOpenApi2EventManagerAssetsCreatePostRequest struct {
	ctx                               context.Context
	ApiService                        *EventManagerAssetsCreateV2ApiService
	version                           string
	eventManagerAssetsCreateV2Request *EventManagerAssetsCreateV2Request
}

func (r *ApiOpenApi2EventManagerAssetsCreatePostRequest) EventManagerAssetsCreateV2Request(eventManagerAssetsCreateV2Request EventManagerAssetsCreateV2Request) *ApiOpenApi2EventManagerAssetsCreatePostRequest {
	r.eventManagerAssetsCreateV2Request = &eventManagerAssetsCreateV2Request
	return r
}

func (r *ApiOpenApi2EventManagerAssetsCreatePostRequest) Execute() (*EventManagerAssetsCreateV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2EventManagerAssetsCreatePostRequest) AccessToken(accessToken string) *ApiOpenApi2EventManagerAssetsCreatePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2EventManagerAssetsCreatePostRequest) WithLog(enable bool) *ApiOpenApi2EventManagerAssetsCreatePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2EventManagerAssetsCreatePost Method for OpenApi2EventManagerAssetsCreatePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param version request version
	@return ApiOpenApi2EventManagerAssetsCreatePostRequest
*/
func (a *EventManagerAssetsCreateV2ApiService) Post(ctx context.Context, version string) *ApiOpenApi2EventManagerAssetsCreatePostRequest {
	return &ApiOpenApi2EventManagerAssetsCreatePostRequest{
		ApiService: a,
		ctx:        ctx,
		version:    version,
	}
}

// Execute executes the request
//
//	@return EventManagerAssetsCreateV2Response
func (a *EventManagerAssetsCreateV2ApiService) postExecute(r *ApiOpenApi2EventManagerAssetsCreatePostRequest) (*EventManagerAssetsCreateV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *EventManagerAssetsCreateV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/event_manager/assets/create/"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterValueToString(r.version, "version")), -1)

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
	localVarPostBody = r.eventManagerAssetsCreateV2Request
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
