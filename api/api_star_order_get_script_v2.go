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

// StarOrderGetScriptV2ApiService StarOrderGetScriptV2Api service
type StarOrderGetScriptV2ApiService service

type ApiOpenApi2StarOrderGetScriptGetRequest struct {
	ctx        context.Context
	ApiService *StarOrderGetScriptV2ApiService
	starId     *int64
	orderIds   *[]int64
}

// 星图客户ID
func (r *ApiOpenApi2StarOrderGetScriptGetRequest) StarId(starId int64) *ApiOpenApi2StarOrderGetScriptGetRequest {
	r.starId = &starId
	return r
}

// 任务ID
func (r *ApiOpenApi2StarOrderGetScriptGetRequest) OrderIds(orderIds []int64) *ApiOpenApi2StarOrderGetScriptGetRequest {
	r.orderIds = &orderIds
	return r
}

func (r *ApiOpenApi2StarOrderGetScriptGetRequest) Execute() (*StarOrderGetScriptV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarOrderGetScriptGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarOrderGetScriptGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarOrderGetScriptGetRequest) WithLog(enable bool) *ApiOpenApi2StarOrderGetScriptGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarOrderGetScriptGet Method for OpenApi2StarOrderGetScriptGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarOrderGetScriptGetRequest
*/
func (a *StarOrderGetScriptV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarOrderGetScriptGetRequest {
	return &ApiOpenApi2StarOrderGetScriptGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarOrderGetScriptV2Response
func (a *StarOrderGetScriptV2ApiService) getExecute(r *ApiOpenApi2StarOrderGetScriptGetRequest) (*StarOrderGetScriptV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarOrderGetScriptV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/order/get_script/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.orderIds == nil {
		return localVarReturnValue, nil, ReportError("orderIds is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "order_ids", r.orderIds)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
