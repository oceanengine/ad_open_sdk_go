/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryQualificationSubmitV30RequestQualificationsInner struct for AdvertiserDeliveryQualificationSubmitV30RequestQualificationsInner
type AdvertiserDeliveryQualificationSubmitV30RequestQualificationsInner struct {
	// 图片附件ids，通过【上传资质图片】接口获取
	AttachmentIds []int64 `json:"attachment_ids"`
	// 资质id ① 针对新增的场景传0即可  ② 提投放资质（单资质）时系统会返回生成的qualification_id，针对编辑的场景，再此传入对应需要编辑的qualification_id即可
	QualificationId   *int64                                                                  `json:"qualification_id,omitempty"`
	QualificationType AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType `json:"qualification_type"`
}
