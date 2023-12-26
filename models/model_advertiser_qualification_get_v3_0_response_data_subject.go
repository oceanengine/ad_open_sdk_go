/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationGetV30ResponseDataSubject 主体资质，相关字段见下:
type AdvertiserQualificationGetV30ResponseDataSubject struct {
	// 详细地址
	Address *string `json:"address,omitempty"`
	// 资质图片附件id
	AttachmentId *string                                            `json:"attachment_id,omitempty"`
	CheckType    *AdvertiserQualificationGetV30DataSubjectCheckType `json:"check_type,omitempty"`
	// 公司名称
	CompanyName *string                                              `json:"company_name,omitempty"`
	CompanyType *AdvertiserQualificationGetV30DataSubjectCompanyType `json:"company_type,omitempty"`
	// 过期时间，格式yyyy-mm-dd
	EffectiveDate *string `json:"effective_date,omitempty"`
	// 是否有有效日期
	HasEffectiveDate *bool `json:"has_effective_date,omitempty"`
	// 资质图片地址
	PictureUrl *string `json:"picture_url,omitempty"`
	// 法人
	ProprietorName *string `json:"proprietor_name,omitempty"`
	// 资质编号
	QualificationCode *string `json:"qualification_code,omitempty"`
	// 资质id
	QualificationId   *int64                                                     `json:"qualification_id,omitempty"`
	QualificationType *AdvertiserQualificationGetV30DataSubjectQualificationType `json:"qualification_type,omitempty"`
	// 注册城市
	RegisteredCityName *string `json:"registered_city_name,omitempty"`
	// 注册国家
	RegisteredNationName *string `json:"registered_nation_name,omitempty"`
	// 注册省份
	RegisteredProvinceName *string `json:"registered_province_name,omitempty"`
	// 拒绝理由
	RejectReason *string                                         `json:"reject_reason,omitempty"`
	Status       *AdvertiserQualificationGetV30DataSubjectStatus `json:"status,omitempty"`
}
