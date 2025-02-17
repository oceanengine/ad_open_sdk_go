/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10DataCampaignScene
type QianchuanAdDetailGetV10DataCampaignScene string

// List of qianchuan_ad_detail_get_v1.0_data_campaign_scene
const (
	DAILY_SALE_QianchuanAdDetailGetV10DataCampaignScene                  QianchuanAdDetailGetV10DataCampaignScene = "DAILY_SALE"
	LIVE_HEAT_QianchuanAdDetailGetV10DataCampaignScene                   QianchuanAdDetailGetV10DataCampaignScene = "LIVE_HEAT"
	NEW_CUSTOMER_TRANSFORMATION_QianchuanAdDetailGetV10DataCampaignScene QianchuanAdDetailGetV10DataCampaignScene = "NEW_CUSTOMER_TRANSFORMATION"
	NEW_PRODUCT_BOOST_QianchuanAdDetailGetV10DataCampaignScene           QianchuanAdDetailGetV10DataCampaignScene = "NEW_PRODUCT_BOOST"
	PLANT_GRASS_QianchuanAdDetailGetV10DataCampaignScene                 QianchuanAdDetailGetV10DataCampaignScene = "PLANT_GRASS"
	PRODUCT_HEAT_QianchuanAdDetailGetV10DataCampaignScene                QianchuanAdDetailGetV10DataCampaignScene = "PRODUCT_HEAT"
)

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_campaign_scene value
func (v QianchuanAdDetailGetV10DataCampaignScene) Ptr() *QianchuanAdDetailGetV10DataCampaignScene {
	return &v
}
