/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
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

// ToolsAdConvertQueryV2ApiService ToolsAdConvertQueryV2Api service
type ToolsAdConvertQueryV2ApiService service

type ApiOpenApi2ToolsAdConvertQueryGetRequest struct {
	ctx                  context.Context
	ApiService           *ToolsAdConvertQueryV2ApiService
	advancedCreativeType *ToolsAdConvertQueryV2AdvancedCreativeType
	advertiserId         *int64
	appSchema            *string
	appType              *ToolsAdConvertQueryV2AppType
	dedicateType         *ToolsAdConvertQueryV2DedicateType
	deliveryRange        *ToolsAdConvertQueryV2DeliveryRange
	externalUrl          *string
	itunesUrl            *string
	landingType          *ToolsAdConvertQueryV2LandingType
	marketingScene       *ToolsAdConvertQueryV2MarketingScene
	packageName          *string
	promotionContent     *ToolsAdConvertQueryV2PromotionContent
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) AdvancedCreativeType(advancedCreativeType ToolsAdConvertQueryV2AdvancedCreativeType) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.advancedCreativeType = &advancedCreativeType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) AppSchema(appSchema string) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.appSchema = &appSchema
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) AppType(appType ToolsAdConvertQueryV2AppType) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.appType = &appType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) DedicateType(dedicateType ToolsAdConvertQueryV2DedicateType) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.dedicateType = &dedicateType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) DeliveryRange(deliveryRange ToolsAdConvertQueryV2DeliveryRange) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.deliveryRange = &deliveryRange
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) ExternalUrl(externalUrl string) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.externalUrl = &externalUrl
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) ItunesUrl(itunesUrl string) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.itunesUrl = &itunesUrl
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) LandingType(landingType ToolsAdConvertQueryV2LandingType) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.landingType = &landingType
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) MarketingScene(marketingScene ToolsAdConvertQueryV2MarketingScene) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.marketingScene = &marketingScene
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) PackageName(packageName string) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.packageName = &packageName
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) PromotionContent(promotionContent ToolsAdConvertQueryV2PromotionContent) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.promotionContent = &promotionContent
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) Execute() (*ToolsAdConvertQueryV2Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsAdConvertQueryGetRequest) WithLog(enable bool) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsAdConvertQueryGet Method for OpenApi2ToolsAdConvertQueryGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsAdConvertQueryGetRequest
*/
func (a *ToolsAdConvertQueryV2ApiService) Get(ctx context.Context) *ApiOpenApi2ToolsAdConvertQueryGetRequest {
	return &ApiOpenApi2ToolsAdConvertQueryGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsAdConvertQueryV2Response
func (a *ToolsAdConvertQueryV2ApiService) getExecute(r *ApiOpenApi2ToolsAdConvertQueryGetRequest) (*ToolsAdConvertQueryV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsAdConvertQueryV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/ad_convert/query/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.advancedCreativeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advanced_creative_type", r.advancedCreativeType)
	}
	if r.advertiserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	}
	if r.appSchema != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_schema", r.appSchema)
	}
	if r.appType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app_type", r.appType)
	}
	if r.dedicateType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dedicate_type", r.dedicateType)
	}
	if r.deliveryRange != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delivery_range", r.deliveryRange)
	}
	if r.externalUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_url", r.externalUrl)
	}
	if r.itunesUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itunes_url", r.itunesUrl)
	}
	if r.landingType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "landing_type", r.landingType)
	}
	if r.marketingScene != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "marketing_scene", r.marketingScene)
	}
	if r.packageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "package_name", r.packageName)
	}
	if r.promotionContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotion_content", r.promotionContent)
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
