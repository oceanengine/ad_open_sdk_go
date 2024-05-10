/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
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

// StarMcnGetContractedChallengeUrlV2ApiService StarMcnGetContractedChallengeUrlV2Api service
type StarMcnGetContractedChallengeUrlV2ApiService service

type ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest struct {
	ctx         context.Context
	ApiService  *StarMcnGetContractedChallengeUrlV2ApiService
	starId      *int64
	demandId    *int64
	channelId   *string
	developerId *int64
}

func (r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) StarId(starId int64) *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest {
	r.starId = &starId
	return r
}

func (r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) DemandId(demandId int64) *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest {
	r.demandId = &demandId
	return r
}

func (r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) ChannelId(channelId string) *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest {
	r.channelId = &channelId
	return r
}

func (r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) DeveloperId(developerId int64) *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest {
	r.developerId = &developerId
	return r
}

func (r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) Execute() (*StarMcnGetContractedChallengeUrlV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) AccessToken(accessToken string) *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) WithLog(enable bool) *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2StarMcnGetContractedChallengeUrlGet Method for OpenApi2StarMcnGetContractedChallengeUrlGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest
*/
func (a *StarMcnGetContractedChallengeUrlV2ApiService) Get(ctx context.Context) *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest {
	return &ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StarMcnGetContractedChallengeUrlV2Response
func (a *StarMcnGetContractedChallengeUrlV2ApiService) getExecute(r *ApiOpenApi2StarMcnGetContractedChallengeUrlGetRequest) (*StarMcnGetContractedChallengeUrlV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *StarMcnGetContractedChallengeUrlV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/star/mcn/get_contracted_challenge_url/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.starId == nil {
		return localVarReturnValue, nil, ReportError("starId is required and must be specified")
	}
	if r.demandId == nil {
		return localVarReturnValue, nil, ReportError("demandId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "star_id", r.starId)
	parameterAddToHeaderOrQuery(localVarQueryParams, "demand_id", r.demandId)
	if r.channelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "channel_id", r.channelId)
	}
	if r.developerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "developer_id", r.developerId)
	}
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
