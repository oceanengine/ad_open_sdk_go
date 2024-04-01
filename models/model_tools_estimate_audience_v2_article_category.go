/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEstimateAudienceV2ArticleCategory
type ToolsEstimateAudienceV2ArticleCategory string

// List of tools_estimate_audience_v2_article_category
const (
	MILITARY_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "MILITARY"
	CARS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "CARS"
	TECHNOLOGY_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "TECHNOLOGY"
	INTERNATIONAL_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "INTERNATIONAL"
	AMUSEMENT_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "AMUSEMENT"
	PETS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "PETS"
	HISTORY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "HISTORY"
	ENTERTAINMENT_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "ENTERTAINMENT"
	CONSTELLATION_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "CONSTELLATION"
	SOCIETY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SOCIETY"
	STORIES_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "STORIES"
	ANIMATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "ANIMATION"
	PHOTOGRAPHY_ToolsEstimateAudienceV2ArticleCategory   ToolsEstimateAudienceV2ArticleCategory = "PHOTOGRAPHY"
	SPORTS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "SPORTS"
	FASHION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FASHION"
	COLLECTION_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "COLLECTION"
	HOME_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "HOME"
	LAWS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "LAWS"
	WORKPLACE_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "WORKPLACE"
	REGIMEN_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "REGIMEN"
	EDUCATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "EDUCATION"
	TIPS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "TIPS"
	ESSAY_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "ESSAY"
	SCIENCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SCIENCE"
	COMICS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "COMICS"
	GOURMET_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "GOURMET"
	BUSINESS_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "BUSINESS"
	GAMES_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "GAMES"
	WEIGHT_LOSING_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "WEIGHT_LOSING"
	LOCAL_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "LOCAL"
	DIGITAL_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "DIGITAL"
	CULTURE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "CULTURE"
	FREAK_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "FREAK"
	DESIGN_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "DESIGN"
	EXPLORE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EXPLORE"
	MOVIE_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "MOVIE"
	FINANCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FINANCE"
	RUMOR_CRACKER_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "RUMOR_CRACKER"
	HEALTH_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "HEALTH"
	PARENTING_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "PARENTING"
	GOVERNMENT_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "GOVERNMENT"
	GRADUATES_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "GRADUATES"
	ESTATE_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "ESTATE"
	EMOTION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EMOTION"
	TRAVEL_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "TRAVEL"
)

// All allowed values of ToolsEstimateAudienceV2ArticleCategory enum
var AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues = []ToolsEstimateAudienceV2ArticleCategory{
	"MILITARY",
	"CARS",
	"TECHNOLOGY",
	"INTERNATIONAL",
	"AMUSEMENT",
	"PETS",
	"HISTORY",
	"ENTERTAINMENT",
	"CONSTELLATION",
	"SOCIETY",
	"STORIES",
	"ANIMATION",
	"PHOTOGRAPHY",
	"SPORTS",
	"FASHION",
	"COLLECTION",
	"HOME",
	"LAWS",
	"WORKPLACE",
	"REGIMEN",
	"EDUCATION",
	"TIPS",
	"ESSAY",
	"SCIENCE",
	"COMICS",
	"GOURMET",
	"BUSINESS",
	"GAMES",
	"WEIGHT_LOSING",
	"LOCAL",
	"DIGITAL",
	"CULTURE",
	"FREAK",
	"DESIGN",
	"EXPLORE",
	"MOVIE",
	"FINANCE",
	"RUMOR_CRACKER",
	"HEALTH",
	"PARENTING",
	"GOVERNMENT",
	"GRADUATES",
	"ESTATE",
	"EMOTION",
	"TRAVEL",
}

// NewToolsEstimateAudienceV2ArticleCategoryFromValue returns a pointer to a valid ToolsEstimateAudienceV2ArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEstimateAudienceV2ArticleCategoryFromValue(v string) (*ToolsEstimateAudienceV2ArticleCategory, error) {
	ev := ToolsEstimateAudienceV2ArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEstimateAudienceV2ArticleCategory: valid values are %v", v, AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEstimateAudienceV2ArticleCategory) IsValid() bool {
	for _, existing := range AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_estimate_audience_v2_article_category value
func (v ToolsEstimateAudienceV2ArticleCategory) Ptr() *ToolsEstimateAudienceV2ArticleCategory {
	return &v
}
