/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInner struct for AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInner
type AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInner struct {
	// 附属商品
	AuxiliaryProducts []*AdlabGroupDetailV30ResponseDataDataAdInfoProductInfoInnerSupplementsInnerAuxiliaryProductsInner `json:"auxiliary_products,omitempty"`
	SupplementType    *AdlabGroupDetailV30DataDataAdInfoProductInfoSupplementsSupplementType                             `json:"supplement_type,omitempty"`
}
