/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
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
	ENTERTAINMENT_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "ENTERTAINMENT"
	CARS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "CARS"
	HEALTH_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "HEALTH"
	TECHNOLOGY_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "TECHNOLOGY"
	BUSINESS_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "BUSINESS"
	GRADUATES_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "GRADUATES"
	GAMES_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "GAMES"
	EMOTION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EMOTION"
	FINANCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FINANCE"
	INTERNATIONAL_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "INTERNATIONAL"
	AMUSEMENT_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "AMUSEMENT"
	COLLECTION_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "COLLECTION"
	EXPLORE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "EXPLORE"
	ANIMATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "ANIMATION"
	HISTORY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "HISTORY"
	RUMOR_CRACKER_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "RUMOR_CRACKER"
	STORIES_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "STORIES"
	MOVIE_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "MOVIE"
	DIGITAL_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "DIGITAL"
	FREAK_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "FREAK"
	CULTURE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "CULTURE"
	LAWS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "LAWS"
	DESIGN_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "DESIGN"
	MILITARY_ToolsEstimateAudienceV2ArticleCategory      ToolsEstimateAudienceV2ArticleCategory = "MILITARY"
	WEIGHT_LOSING_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "WEIGHT_LOSING"
	PARENTING_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "PARENTING"
	GOVERNMENT_ToolsEstimateAudienceV2ArticleCategory    ToolsEstimateAudienceV2ArticleCategory = "GOVERNMENT"
	SPORTS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "SPORTS"
	TRAVEL_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "TRAVEL"
	ESTATE_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "ESTATE"
	CONSTELLATION_ToolsEstimateAudienceV2ArticleCategory ToolsEstimateAudienceV2ArticleCategory = "CONSTELLATION"
	SOCIETY_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SOCIETY"
	SCIENCE_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "SCIENCE"
	LOCAL_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "LOCAL"
	COMICS_ToolsEstimateAudienceV2ArticleCategory        ToolsEstimateAudienceV2ArticleCategory = "COMICS"
	WORKPLACE_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "WORKPLACE"
	FASHION_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "FASHION"
	HOME_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "HOME"
	PETS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "PETS"
	TIPS_ToolsEstimateAudienceV2ArticleCategory          ToolsEstimateAudienceV2ArticleCategory = "TIPS"
	ESSAY_ToolsEstimateAudienceV2ArticleCategory         ToolsEstimateAudienceV2ArticleCategory = "ESSAY"
	PHOTOGRAPHY_ToolsEstimateAudienceV2ArticleCategory   ToolsEstimateAudienceV2ArticleCategory = "PHOTOGRAPHY"
	GOURMET_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "GOURMET"
	EDUCATION_ToolsEstimateAudienceV2ArticleCategory     ToolsEstimateAudienceV2ArticleCategory = "EDUCATION"
	REGIMEN_ToolsEstimateAudienceV2ArticleCategory       ToolsEstimateAudienceV2ArticleCategory = "REGIMEN"
)

// All allowed values of ToolsEstimateAudienceV2ArticleCategory enum
var AllowedToolsEstimateAudienceV2ArticleCategoryEnumValues = []ToolsEstimateAudienceV2ArticleCategory{
	"ENTERTAINMENT",
	"CARS",
	"HEALTH",
	"TECHNOLOGY",
	"BUSINESS",
	"GRADUATES",
	"GAMES",
	"EMOTION",
	"FINANCE",
	"INTERNATIONAL",
	"AMUSEMENT",
	"COLLECTION",
	"EXPLORE",
	"ANIMATION",
	"HISTORY",
	"RUMOR_CRACKER",
	"STORIES",
	"MOVIE",
	"DIGITAL",
	"FREAK",
	"CULTURE",
	"LAWS",
	"DESIGN",
	"MILITARY",
	"WEIGHT_LOSING",
	"PARENTING",
	"GOVERNMENT",
	"SPORTS",
	"TRAVEL",
	"ESTATE",
	"CONSTELLATION",
	"SOCIETY",
	"SCIENCE",
	"LOCAL",
	"COMICS",
	"WORKPLACE",
	"FASHION",
	"HOME",
	"PETS",
	"TIPS",
	"ESSAY",
	"PHOTOGRAPHY",
	"GOURMET",
	"EDUCATION",
	"REGIMEN",
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
