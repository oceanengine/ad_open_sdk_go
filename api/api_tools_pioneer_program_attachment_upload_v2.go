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

// ToolsPioneerProgramAttachmentUploadV2ApiService ToolsPioneerProgramAttachmentUploadV2Api service
type ToolsPioneerProgramAttachmentUploadV2ApiService service

type ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest struct {
	ctx          context.Context
	ApiService   *ToolsPioneerProgramAttachmentUploadV2ApiService
	appId        *int64
	dataFileType *ToolsPioneerProgramAttachmentUploadV2DataFileType
	fileData     *FormFileInfo
	fileIndex    *int64
	fileTotalNum *int64
	pDate        *string
	platform     *ToolsPioneerProgramAttachmentUploadV2Platform
	debugMode    *bool
}

// 即应用ID，开发者创建应用的唯一标识，可通过应用管理查看
func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) AppId(appId int64) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.appId = &appId
	return r
}

func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) DataFileType(dataFileType ToolsPioneerProgramAttachmentUploadV2DataFileType) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.dataFileType = &dataFileType
	return r
}

// 文件数据，binary格式
func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) FileData(fileData *FormFileInfo) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.fileData = fileData
	return r
}

// 当前文件序号，从1开始，最大值为file_total_num
func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) FileIndex(fileIndex int64) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.fileIndex = &fileIndex
	return r
}

// p_date - platform - data_file_type维度下总文件数，如“2022-06-12”日“巨量引擎”平台的“后端投放数据”共计3份，则file_total_num为3
func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) FileTotalNum(fileTotalNum int64) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.fileTotalNum = &fileTotalNum
	return r
}

// 数据日期，格式为yyyy-MM-dd
func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) PDate(pDate string) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.pDate = &pDate
	return r
}

func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) Platform(platform ToolsPioneerProgramAttachmentUploadV2Platform) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.platform = &platform
	return r
}

// 是否为调试模式，调试模式下数据不会被最终记录，默认为False
func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) DebugMode(debugMode bool) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.debugMode = &debugMode
	return r
}

func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) Execute() (*ToolsPioneerProgramAttachmentUploadV2Response, *http.Response, error) {
	return r.ApiService.postExecute(r)
}

func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) AccessToken(accessToken string) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	r.ctx = context.WithValue(r.ctx, config.ContextAccessToken, accessToken)
	return r
}

func (r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) WithLog(enable bool) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, config.ContextEnableLog, true)
	}
	return r
}

/*
OpenApi2ToolsPioneerProgramAttachmentUploadPost Method for OpenApi2ToolsPioneerProgramAttachmentUploadPost

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest
*/
func (a *ToolsPioneerProgramAttachmentUploadV2ApiService) Post(ctx context.Context) *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest {
	return &ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ToolsPioneerProgramAttachmentUploadV2Response
func (a *ToolsPioneerProgramAttachmentUploadV2ApiService) postExecute(r *ApiOpenApi2ToolsPioneerProgramAttachmentUploadPostRequest) (*ToolsPioneerProgramAttachmentUploadV2Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           map[string]*FormFileInfo
		localVarReturnValue *ToolsPioneerProgramAttachmentUploadV2Response
	)

	r.ctx = a.client.prepareCtx(r.ctx)

	localBasePath := a.client.Cfg.GetBasePath()

	localVarPath := localBasePath + "/open_api/2/tools/pioneer_program/attachment/upload/"

	localVarHeaderParams := make(map[string]string)
	formFiles = make(map[string]*FormFileInfo)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appId == nil {
		return localVarReturnValue, nil, ReportError("appId is required and must be specified")
	}
	if r.dataFileType == nil {
		return localVarReturnValue, nil, ReportError("dataFileType is required and must be specified")
	}
	if r.fileData == nil {
		return localVarReturnValue, nil, ReportError("fileData is required and must be specified")
	}
	if r.fileIndex == nil {
		return localVarReturnValue, nil, ReportError("fileIndex is required and must be specified")
	}
	if r.fileTotalNum == nil {
		return localVarReturnValue, nil, ReportError("fileTotalNum is required and must be specified")
	}
	if r.pDate == nil {
		return localVarReturnValue, nil, ReportError("pDate is required and must be specified")
	}
	if r.platform == nil {
		return localVarReturnValue, nil, ReportError("platform is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "app_id", r.appId)
	parameterAddToHeaderOrQuery(localVarFormParams, "data_file_type", r.dataFileType)
	if r.debugMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "debug_mode", r.debugMode)
	}
	if r.fileData != nil {
		formFiles["file_data"] = r.fileData
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "file_index", r.fileIndex)
	parameterAddToHeaderOrQuery(localVarFormParams, "file_total_num", r.fileTotalNum)
	parameterAddToHeaderOrQuery(localVarFormParams, "p_date", r.pDate)
	parameterAddToHeaderOrQuery(localVarFormParams, "platform", r.platform)
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
