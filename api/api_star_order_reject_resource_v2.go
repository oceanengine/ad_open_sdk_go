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

// StarOrderRejectResourceV2ApiService StarOrderRejectResourceV2Api service
type StarOrderRejectResourceV2ApiService service

type ApiOpenApi2StarOrderRejectResourcePostRequest struct {
	ctx                              context.Context
	ApiService                       *StarOrderRejectResourceV2ApiService
	starOrderRejectResourceV2Request *StarOrderRejectResourceV2Request
}

func (r *ApiOpenApi2StarOrderRejectResourcePostRequest) StarOrderRejectResourceV2Request(starOrderRejectResourceV2Request StarOrderRejectResourceV2Request) *ApiOpenApi2StarOrderRejectResourcePostRequest {
	r.starOrderRejectResourceV2Request = &starOrderRejectResourceV2Request
	return r
}

func (r *ApiOpenApi2StarOrderRejectResourcePostRequest) Execute() (*StarOrderRejectResourceV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarOrderRejectResourcePostRequest) AccessToken(accessToken string) *ApiOpenApi2StarOrderRejectResourcePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarOrderRejectResourcePostRequest) WithLog(enable bool) *ApiOpenApi2StarOrderRejectResourcePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarOrderRejectResourcePost Method for OpenApi2StarOrderRejectResourcePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarOrderRejectResourcePostRequest
*/
func (a *StarOrderRejectResourceV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarOrderRejectResourcePostRequest {
	return &ApiOpenApi2StarOrderRejectResourcePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarOrderRejectResourceV2Response
func (a *StarOrderRejectResourceV2ApiService) postExecute(r *ApiOpenApi2StarOrderRejectResourcePostRequest) (*StarOrderRejectResourceV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarOrderRejectResourceV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/order/reject_resource/"

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
	localVarPostBody = r.starOrderRejectResourceV2Request
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
