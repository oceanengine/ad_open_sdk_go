/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteTemplateCreateV2ResponseDataBricksInnerForm 表单组件描述
type ToolsSiteTemplateCreateV2ResponseDataBricksInnerForm struct {
	// 表单ID，当`form`不为空时，有值。用户可以通过[【获取表单列表】]（https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710634663948#item-link-%E5%BA%94%E7%AD%94%E7%A4%BA%E4%BE%8B）接口或[【青鸟线索通平台】]（https://clue.oceanengine.com/）获取表单ID
	InstanceId int64 `json:"instance_id"`
}
