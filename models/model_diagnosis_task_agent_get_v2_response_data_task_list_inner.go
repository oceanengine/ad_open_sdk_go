/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DiagnosisTaskAgentGetV2ResponseDataTaskListInner struct for DiagnosisTaskAgentGetV2ResponseDataTaskListInner
type DiagnosisTaskAgentGetV2ResponseDataTaskListInner struct {
	//
	AdvertiserId             *int64                                                       `json:"advertiser_id,omitempty"`
	IsAdHighQualityMaterial  *DiagnosisTaskAgentGetV2DataTaskListIsAdHighQualityMaterial  `json:"is_ad_high_quality_material,omitempty"`
	IsEcpHighQualityMaterial *DiagnosisTaskAgentGetV2DataTaskListIsEcpHighQualityMaterial `json:"is_ecp_high_quality_material,omitempty"`
	IsFirstPublishMaterial   *DiagnosisTaskAgentGetV2DataTaskListIsFirstPublishMaterial   `json:"is_first_publish_material,omitempty"`
	IsInefficientMaterial    *DiagnosisTaskAgentGetV2DataTaskListIsInefficientMaterial    `json:"is_inefficient_material,omitempty"`
	//
	MaterialId *int64                                    `json:"material_id,omitempty"`
	Status     DiagnosisTaskAgentGetV2DataTaskListStatus `json:"status"`
	//
	TaskId int64 `json:"task_id"`
	//
	VideoId string `json:"video_id"`
}
