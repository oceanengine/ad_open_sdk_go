/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeSyncItemToAdV2ResponseData
type StarChallengeSyncItemToAdV2ResponseData struct {
	// 是否存在成功的推送
	ExistSuccessSync bool `json:"exist_success_sync"`
	// 推送失败的原因
	FailSyncReason []*StarChallengeSyncItemToAdV2ResponseDataFailSyncReasonInner `json:"fail_sync_reason"`
	// 推送详情，格式等同于素材集市
	SyncDetails []*StarChallengeSyncItemToAdV2ResponseDataSyncDetailsInner `json:"sync_details,omitempty"`
}
