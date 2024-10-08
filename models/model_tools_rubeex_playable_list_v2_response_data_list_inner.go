/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRubeexPlayableListV2ResponseDataListInner struct for ToolsRubeexPlayableListV2ResponseDataListInner
type ToolsRubeexPlayableListV2ResponseDataListInner struct {
	AdStatus *ToolsRubeexPlayableListV2DataListAdStatus `json:"ad_status,omitempty"`
	//
	DataMd5 *string `json:"data_md5,omitempty"`
	//
	PlayableId *int64 `json:"playable_id,omitempty"`
	//
	PlayableName        *string                                               `json:"playable_name,omitempty"`
	PlayableOrientation *ToolsRubeexPlayableListV2DataListPlayableOrientation `json:"playable_orientation,omitempty"`
	//
	PlayableUrl *string `json:"playable_url,omitempty"`
	//
	PreviewUrl *string                                  `json:"preview_url,omitempty"`
	Status     *ToolsRubeexPlayableListV2DataListStatus `json:"status,omitempty"`
}
