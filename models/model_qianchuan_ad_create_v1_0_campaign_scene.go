/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10CampaignScene
type QianchuanAdCreateV10CampaignScene string

// List of qianchuan_ad_create_v1.0_campaign_scene
const (
	DAILY_SALE_QianchuanAdCreateV10CampaignScene                  QianchuanAdCreateV10CampaignScene = "DAILY_SALE"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanAdCreateV10CampaignScene QianchuanAdCreateV10CampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanAdCreateV10CampaignScene           QianchuanAdCreateV10CampaignScene = "NEW_PRODUCT_BOOST"
)

// Ptr returns reference to qianchuan_ad_create_v1.0_campaign_scene value
func (v QianchuanAdCreateV10CampaignScene) Ptr() *QianchuanAdCreateV10CampaignScene {
	return &v
}
