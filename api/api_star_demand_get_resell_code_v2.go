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

// StarDemandGetResellCodeV2ApiService StarDemandGetResellCodeV2Api service
type StarDemandGetResellCodeV2ApiService service

type ApiOpenApi2StarDemandGetResellCodeGetRequest struct {
	ctx        context.Context
	ApiService *StarDemandGetResellCodeV2ApiService
	starId     *int64
	taskId     *int64
}

func (r *ApiOpenApi2StarDemandGetResellCodeGetRequest) StarId(starId int64) *ApiOpenApi2StarDemandGetResellCodeGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarDemandGetResellCodeGetRequest) TaskId(taskId int64) *ApiOpenApi2StarDemandGetResellCodeGetRequest {
	r.taskId = &taskId
	return r
}

func (r *ApiOpenApi2StarDemandGetResellCodeGetRequest) Execute() (*StarDemandGetResellCodeV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarDemandGetResellCodeGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandGetResellCodeGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandGetResellCodeGetRequest) WithLog(enable bool) *ApiOpenApi2StarDemandGetResellCodeGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandGetResellCodeGet Method for OpenApi2StarDemandGetResellCodeGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandGetResellCodeGetRequest
*/
func (a *StarDemandGetResellCodeV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarDemandGetResellCodeGetRequest {
	return &ApiOpenApi2StarDemandGetResellCodeGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandGetResellCodeV2Response
func (a *StarDemandGetResellCodeV2ApiService) getExecute(r *ApiOpenApi2StarDemandGetResellCodeGetRequest) (*StarDemandGetResellCodeV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandGetResellCodeV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/get_resell_code/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.taskId == nil {
		return localVarReturnValue, nil, ReportError("taskId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "task_id", r.taskId)
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

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
