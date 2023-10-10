/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestDpaProductTargetInner struct for ProjectCreateV30RequestDpaProductTargetInner
type ProjectCreateV30RequestDpaProductTargetInner struct {
	// 定向规则，允许值：=、 !=、 >、 <、 contain、 exclude、 notEmpty
	Rule *string `json:"rule,omitempty"`
	// 筛选字段
	Title *string `json:"title,omitempty"`
	// 字段类型，允许值：int、 double、 long、 string
	Type *string `json:"type,omitempty"`
	// 规则值
	Value *string `json:"value,omitempty"`
}
