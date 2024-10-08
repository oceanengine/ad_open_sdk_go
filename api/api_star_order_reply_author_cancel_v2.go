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

// StarOrderReplyAuthorCancelV2ApiService StarOrderReplyAuthorCancelV2Api service
type StarOrderReplyAuthorCancelV2ApiService service

type ApiOpenApi2StarOrderReplyAuthorCancelPostRequest struct {
	ctx                                 context.Context
	ApiService                          *StarOrderReplyAuthorCancelV2ApiService
	starOrderReplyAuthorCancelV2Request *StarOrderReplyAuthorCancelV2Request
}

func (r *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest) StarOrderReplyAuthorCancelV2Request(starOrderReplyAuthorCancelV2Request StarOrderReplyAuthorCancelV2Request) *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest {
	r.starOrderReplyAuthorCancelV2Request = &starOrderReplyAuthorCancelV2Request
	return r
}

func (r *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest) Execute() (*StarOrderReplyAuthorCancelV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest) AccessToken(accessToken string) *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest) WithLog(enable bool) *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarOrderReplyAuthorCancelPost Method for OpenApi2StarOrderReplyAuthorCancelPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarOrderReplyAuthorCancelPostRequest
*/
func (a *StarOrderReplyAuthorCancelV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest {
	return &ApiOpenApi2StarOrderReplyAuthorCancelPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarOrderReplyAuthorCancelV2Response
func (a *StarOrderReplyAuthorCancelV2ApiService) postExecute(r *ApiOpenApi2StarOrderReplyAuthorCancelPostRequest) (*StarOrderReplyAuthorCancelV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarOrderReplyAuthorCancelV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/order/reply_author_cancel/"

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
	localVarPostBody = r.starOrderReplyAuthorCancelV2Request
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
