/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectCreateV30LocalDeliveryScene
type LocalProjectCreateV30LocalDeliveryScene string

// List of local_project_create_v3.0_local_delivery_scene
const (
	CONTENT_HEAT_LocalProjectCreateV30LocalDeliveryScene  LocalProjectCreateV30LocalDeliveryScene = "CONTENT_HEAT"
	POI_RECOMMEND_LocalProjectCreateV30LocalDeliveryScene LocalProjectCreateV30LocalDeliveryScene = "POI_RECOMMEND"
	PRODUCT_PAY_LocalProjectCreateV30LocalDeliveryScene   LocalProjectCreateV30LocalDeliveryScene = "PRODUCT_PAY"
)

// Ptr returns reference to local_project_create_v3.0_local_delivery_scene value
func (v LocalProjectCreateV30LocalDeliveryScene) Ptr() *LocalProjectCreateV30LocalDeliveryScene {
	return &v
}