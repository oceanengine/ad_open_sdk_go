/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderUserDistributionGetV2ResponseData
type StarReportOrderUserDistributionGetV2ResponseData struct {
	//
	Activity []*StarReportOrderUserDistributionGetV2ResponseDataActivityInner `json:"activity,omitempty"`
	//
	Age []*StarReportOrderUserDistributionGetV2ResponseDataAgeInner `json:"age,omitempty"`
	//
	City []*StarReportOrderUserDistributionGetV2ResponseDataCityInner `json:"city,omitempty"`
	//
	Device []*StarReportOrderUserDistributionGetV2ResponseDataDeviceInner `json:"device,omitempty"`
	//
	Gender []*StarReportOrderUserDistributionGetV2ResponseDataGenderInner `json:"gender,omitempty"`
	//
	Interest []*StarReportOrderUserDistributionGetV2ResponseDataInterestInner `json:"interest,omitempty"`
	//
	Province []*StarReportOrderUserDistributionGetV2ResponseDataProvinceInner `json:"province,omitempty"`
	//
	UpdateTime *string `json:"update_time,omitempty"`
}
