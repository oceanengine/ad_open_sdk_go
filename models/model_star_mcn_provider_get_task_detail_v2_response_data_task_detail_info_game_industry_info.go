/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnProviderGetTaskDetailV2ResponseDataTaskDetailInfoGameIndustryInfo 游戏行业版信息
type StarMcnProviderGetTaskDetailV2ResponseDataTaskDetailInfoGameIndustryInfo struct {
	// 游戏节点时间
	GameNodeTime *int64 `json:"game_node_time,omitempty"`
	// 游戏节点
	GameNodeType *int64 `json:"game_node_type,omitempty"`
	// 游戏卖点
	GamePoints []*StarMcnProviderGetTaskDetailV2ResponseDataTaskDetailInfoGameIndustryInfoGamePointsInner `json:"game_points,omitempty"`
	// 游戏类型 id
	GameType *int64 `json:"game_type,omitempty"`
	// 游戏类型 名称
	GameTypeName *string `json:"game_type_name,omitempty"`
	// 相似游戏h
	SameGameNames []string `json:"same_game_names,omitempty"`
}
