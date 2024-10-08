/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsRtaGetInfoTmpV2ResponseDataInterfaceInfo RTA配置数据
type ToolsRtaGetInfoTmpV2ResponseDataInterfaceInfo struct {
	DeliveryRange *ToolsRtaGetInfoTmpV2DataInterfaceInfoDeliveryRange `json:"delivery_range,omitempty"`
	// 站内QPS
	LocalQps *int64 `json:"local_qps,omitempty"`
	// 接口地址状态；1：生效 0：无效
	Status *int64 `json:"status,omitempty"`
	// 穿山甲QPS
	UnionQps *int64 `json:"union_qps,omitempty"`
	// 接口地址
	Url *string `json:"url,omitempty"`
}
