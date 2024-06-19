/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
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

// StarDemandOmCreateChallengeV2ApiService StarDemandOmCreateChallengeV2Api service
type StarDemandOmCreateChallengeV2ApiService service

type ApiOpenApi2StarDemandOmCreateChallengePostRequest struct {
	ctx                                  context.Context
	ApiService                           *StarDemandOmCreateChallengeV2ApiService
	starDemandOmCreateChallengeV2Request *StarDemandOmCreateChallengeV2Request
}

func (r *ApiOpenApi2StarDemandOmCreateChallengePostRequest) StarDemandOmCreateChallengeV2Request(starDemandOmCreateChallengeV2Request StarDemandOmCreateChallengeV2Request) *ApiOpenApi2StarDemandOmCreateChallengePostRequest {
	r.starDemandOmCreateChallengeV2Request = &starDemandOmCreateChallengeV2Request
	return r
}

func (r *ApiOpenApi2StarDemandOmCreateChallengePostRequest) Execute() (*StarDemandOmCreateChallengeV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2StarDemandOmCreateChallengePostRequest) AccessToken(accessToken string) *ApiOpenApi2StarDemandOmCreateChallengePostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarDemandOmCreateChallengePostRequest) WithLog(enable bool) *ApiOpenApi2StarDemandOmCreateChallengePostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarDemandOmCreateChallengePost Method for OpenApi2StarDemandOmCreateChallengePost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarDemandOmCreateChallengePostRequest
*/
func (a *StarDemandOmCreateChallengeV2ApiService) Post(ctx context.Context) *ApiOpenApi2StarDemandOmCreateChallengePostRequest {
	return &ApiOpenApi2StarDemandOmCreateChallengePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarDemandOmCreateChallengeV2Response
func (a *StarDemandOmCreateChallengeV2ApiService) postExecute(r *ApiOpenApi2StarDemandOmCreateChallengePostRequest) (*StarDemandOmCreateChallengeV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarDemandOmCreateChallengeV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/demand/om_create_challenge/"

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

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.starDemandOmCreateChallengeV2Request
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
