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

// AdvertiserDeliveryQualificationListV30ApiService AdvertiserDeliveryQualificationListV30Api service
type AdvertiserDeliveryQualificationListV30ApiService service

type ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest struct {
	ctx               context.Context
	ApiService        *AdvertiserDeliveryQualificationListV30ApiService
	advertiserId      *int64
	page              *int32
	pageSize          *int32
	qualificationType *AdvertiserDeliveryQualificationListV30QualificationType
	status            *AdvertiserDeliveryQualificationListV30Status
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) AdvertiserId(advertiserId int64) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	r.advertiserId = &advertiserId
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) Page(page int32) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	r.page = &page
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) PageSize(pageSize int32) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) QualificationType(qualificationType AdvertiserDeliveryQualificationListV30QualificationType) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	r.qualificationType = &qualificationType
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) Status(status AdvertiserDeliveryQualificationListV30Status) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	r.status = &status
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) Execute() (*AdvertiserDeliveryQualificationListV30Response, *http.Response, error) {
	return r.ApiService.getExecute(r)
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) AccessToken(accessToken string) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) WithLog(enable bool) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApiV30AdvertiserDeliveryQualificationListGet Method for OpenApiV30AdvertiserDeliveryQualificationListGet

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest
*/
func (a *AdvertiserDeliveryQualificationListV30ApiService) Get(ctx context.Context) *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest {
	return &ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AdvertiserDeliveryQualificationListV30Response
func (a *AdvertiserDeliveryQualificationListV30ApiService) getExecute(r *ApiOpenApiV30AdvertiserDeliveryQualificationListGetRequest) (*AdvertiserDeliveryQualificationListV30Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *AdvertiserDeliveryQualificationListV30Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/v3.0/advertiser/delivery_qualification/list/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.advertiserId == nil {
		return localVarReturnValue, nil, ReportError("advertiserId is required and must be specified")
	}
	if r.page == nil {
		return localVarReturnValue, nil, ReportError("page is required and must be specified")
	}
	if *r.page < 1 {
		return localVarReturnValue, nil, ReportError("page must be greater than 1")
	}
	if *r.page > 10000 {
		return localVarReturnValue, nil, ReportError("page must be less than 10000")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, ReportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 1 {
		return localVarReturnValue, nil, ReportError("pageSize must be greater than 1")
	}
	if *r.pageSize > 100 {
		return localVarReturnValue, nil, ReportError("pageSize must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "advertiser_id", r.advertiserId)
	if r.qualificationType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "qualification_type", r.qualificationType)
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status)
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page)
	parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize)
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
