/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataAudienceArticleCategory
type AdGetV2DataAudienceArticleCategory string

// List of ad_get_v2_data_audience_article_category
const (
	ESSAY_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "ESSAY"
	SOCIETY_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "SOCIETY"
	AMUSEMENT_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "AMUSEMENT"
	COLLECTION_AdGetV2DataAudienceArticleCategory    AdGetV2DataAudienceArticleCategory = "COLLECTION"
	PHOTOGRAPHY_AdGetV2DataAudienceArticleCategory   AdGetV2DataAudienceArticleCategory = "PHOTOGRAPHY"
	INTERNATIONAL_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "INTERNATIONAL"
	TRAVEL_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "TRAVEL"
	DIGITAL_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "DIGITAL"
	ESTATE_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "ESTATE"
	HISTORY_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "HISTORY"
	MOVIE_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "MOVIE"
	FINANCE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "FINANCE"
	GAMES_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "GAMES"
	BUSINESS_AdGetV2DataAudienceArticleCategory      AdGetV2DataAudienceArticleCategory = "BUSINESS"
	CONSTELLATION_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "CONSTELLATION"
	WORKPLACE_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "WORKPLACE"
	GOVERNMENT_AdGetV2DataAudienceArticleCategory    AdGetV2DataAudienceArticleCategory = "GOVERNMENT"
	FASHION_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "FASHION"
	ENTERTAINMENT_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "ENTERTAINMENT"
	EDUCATION_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "EDUCATION"
	DESIGN_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "DESIGN"
	RUMOR_CRACKER_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "RUMOR_CRACKER"
	CULTURE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "CULTURE"
	PARENTING_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "PARENTING"
	CARS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "CARS"
	WEIGHT_LOSING_AdGetV2DataAudienceArticleCategory AdGetV2DataAudienceArticleCategory = "WEIGHT_LOSING"
	LOCAL_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "LOCAL"
	SPORTS_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "SPORTS"
	SCIENCE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "SCIENCE"
	HOME_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "HOME"
	TECHNOLOGY_AdGetV2DataAudienceArticleCategory    AdGetV2DataAudienceArticleCategory = "TECHNOLOGY"
	MILITARY_AdGetV2DataAudienceArticleCategory      AdGetV2DataAudienceArticleCategory = "MILITARY"
	STORIES_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "STORIES"
	EXPLORE_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "EXPLORE"
	ANIMATION_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "ANIMATION"
	PETS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "PETS"
	COMICS_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "COMICS"
	HEALTH_AdGetV2DataAudienceArticleCategory        AdGetV2DataAudienceArticleCategory = "HEALTH"
	FREAK_AdGetV2DataAudienceArticleCategory         AdGetV2DataAudienceArticleCategory = "FREAK"
	GOURMET_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "GOURMET"
	GRADUATES_AdGetV2DataAudienceArticleCategory     AdGetV2DataAudienceArticleCategory = "GRADUATES"
	TIPS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "TIPS"
	REGIMEN_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "REGIMEN"
	EMOTION_AdGetV2DataAudienceArticleCategory       AdGetV2DataAudienceArticleCategory = "EMOTION"
	LAWS_AdGetV2DataAudienceArticleCategory          AdGetV2DataAudienceArticleCategory = "LAWS"
)

// All allowed values of AdGetV2DataAudienceArticleCategory enum
var AllowedAdGetV2DataAudienceArticleCategoryEnumValues = []AdGetV2DataAudienceArticleCategory{
	"ESSAY",
	"SOCIETY",
	"AMUSEMENT",
	"COLLECTION",
	"PHOTOGRAPHY",
	"INTERNATIONAL",
	"TRAVEL",
	"DIGITAL",
	"ESTATE",
	"HISTORY",
	"MOVIE",
	"FINANCE",
	"GAMES",
	"BUSINESS",
	"CONSTELLATION",
	"WORKPLACE",
	"GOVERNMENT",
	"FASHION",
	"ENTERTAINMENT",
	"EDUCATION",
	"DESIGN",
	"RUMOR_CRACKER",
	"CULTURE",
	"PARENTING",
	"CARS",
	"WEIGHT_LOSING",
	"LOCAL",
	"SPORTS",
	"SCIENCE",
	"HOME",
	"TECHNOLOGY",
	"MILITARY",
	"STORIES",
	"EXPLORE",
	"ANIMATION",
	"PETS",
	"COMICS",
	"HEALTH",
	"FREAK",
	"GOURMET",
	"GRADUATES",
	"TIPS",
	"REGIMEN",
	"EMOTION",
	"LAWS",
}

// NewAdGetV2DataAudienceArticleCategoryFromValue returns a pointer to a valid AdGetV2DataAudienceArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceArticleCategoryFromValue(v string) (*AdGetV2DataAudienceArticleCategory, error) {
	ev := AdGetV2DataAudienceArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceArticleCategory: valid values are %v", v, AllowedAdGetV2DataAudienceArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceArticleCategory) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_article_category value
func (v AdGetV2DataAudienceArticleCategory) Ptr() *AdGetV2DataAudienceArticleCategory {
	return &v
}