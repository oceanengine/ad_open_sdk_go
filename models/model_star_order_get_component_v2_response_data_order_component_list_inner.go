/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetComponentV2ResponseDataOrderComponentListInner struct for StarOrderGetComponentV2ResponseDataOrderComponentListInner
type StarOrderGetComponentV2ResponseDataOrderComponentListInner struct {
	// 常规组件信息
	LinkComponentList []*StarOrderGetComponentV2ResponseDataOrderComponentListInnerLinkComponentListInner `json:"link_component_list,omitempty"`
	// 任务ID
	OrderId *int64 `json:"order_id,omitempty"`
}
