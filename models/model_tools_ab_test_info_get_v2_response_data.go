/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestInfoGetV2ResponseData
type ToolsAbTestInfoGetV2ResponseData struct {
	Conclusion       *ToolsAbTestInfoGetV2DataConclusion               `json:"conclusion,omitempty"`
	ConclusionDetail *ToolsAbTestInfoGetV2ResponseDataConclusionDetail `json:"conclusion_detail,omitempty"`
	// 实验对象ID列表
	ObjectList []*ToolsAbTestInfoGetV2ResponseDataObjectListInner `json:"object_list,omitempty"`
	// 报告日期，格式：\"20201225\"
	ReportDate *string `json:"report_date,omitempty"`
	// 报告时间，格式：\"2020-12-25 07:12:08\"
	ReportTime *string                         `json:"report_time,omitempty"`
	Status     *ToolsAbTestInfoGetV2DataStatus `json:"status,omitempty"`
	// 实验结束时间，格式：\"2020-12-25 07:12:08\"
	TestEndTime *string `json:"test_end_time,omitempty"`
	// 实验ID
	TestId *int64 `json:"test_id,omitempty"`
	// 实验名称
	TestName *string `json:"test_name,omitempty"`
	// 实验对象详细报告
	TestReports []*ToolsAbTestInfoGetV2ResponseDataTestReportsInner `json:"test_reports,omitempty"`
	// 实验开始时间，格式：\"2020-12-25 07:12:08\"
	TestStartTime *string                              `json:"test_start_time,omitempty"`
	TestType      *ToolsAbTestInfoGetV2DataTestType    `json:"test_type,omitempty"`
	TestVersion   *ToolsAbTestInfoGetV2DataTestVersion `json:"test_version,omitempty"`
}