/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectListV30DataListInternalAdvertiserInfoClassify
type ProjectListV30DataListInternalAdvertiserInfoClassify string

// List of project_list_v3.0_data_list_internal_advertiser_info_classify
const (
	CLASSIFY_APPORTION_ProjectListV30DataListInternalAdvertiserInfoClassify  ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_APPORTION"
	CLASSIFY_EXCHANGE_ProjectListV30DataListInternalAdvertiserInfoClassify   ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_EXCHANGE"
	CLASSIFY_INTERNAL_ProjectListV30DataListInternalAdvertiserInfoClassify   ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_INTERNAL"
	CLASSIFY_SALE_ProjectListV30DataListInternalAdvertiserInfoClassify       ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_SALE"
	CLASSIFY_SUPPLEMENT_ProjectListV30DataListInternalAdvertiserInfoClassify ProjectListV30DataListInternalAdvertiserInfoClassify = "CLASSIFY_SUPPLEMENT"
)

// Ptr returns reference to project_list_v3.0_data_list_internal_advertiser_info_classify value
func (v ProjectListV30DataListInternalAdvertiserInfoClassify) Ptr() *ProjectListV30DataListInternalAdvertiserInfoClassify {
	return &v
}
