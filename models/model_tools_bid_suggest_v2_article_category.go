/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2ArticleCategory
type ToolsBidSuggestV2ArticleCategory string

// List of tools_bid_suggest_v2_article_category
const (
	COLLECTION_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "COLLECTION"
	BUSINESS_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "BUSINESS"
	HEALTH_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "HEALTH"
	LOCAL_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "LOCAL"
	STORIES_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "STORIES"
	PARENTING_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "PARENTING"
	HISTORY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "HISTORY"
	FREAK_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "FREAK"
	PHOTOGRAPHY_ToolsBidSuggestV2ArticleCategory   ToolsBidSuggestV2ArticleCategory = "PHOTOGRAPHY"
	WEIGHT_LOSING_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "WEIGHT_LOSING"
	LAWS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "LAWS"
	GOVERNMENT_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "GOVERNMENT"
	PETS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "PETS"
	INTERNATIONAL_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "INTERNATIONAL"
	FASHION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FASHION"
	EDUCATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "EDUCATION"
	TIPS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "TIPS"
	ANIMATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "ANIMATION"
	MILITARY_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "MILITARY"
	ESSAY_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "ESSAY"
	ENTERTAINMENT_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "ENTERTAINMENT"
	TRAVEL_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "TRAVEL"
	TECHNOLOGY_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "TECHNOLOGY"
	REGIMEN_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "REGIMEN"
	SPORTS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "SPORTS"
	DIGITAL_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "DIGITAL"
	MOVIE_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "MOVIE"
	AMUSEMENT_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "AMUSEMENT"
	EXPLORE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EXPLORE"
	CULTURE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "CULTURE"
	SCIENCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SCIENCE"
	CONSTELLATION_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "CONSTELLATION"
	CARS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "CARS"
	FINANCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FINANCE"
	EMOTION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EMOTION"
	SOCIETY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SOCIETY"
	COMICS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "COMICS"
	HOME_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "HOME"
	DESIGN_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "DESIGN"
	GRADUATES_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "GRADUATES"
	GAMES_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "GAMES"
	ESTATE_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "ESTATE"
	WORKPLACE_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "WORKPLACE"
	RUMOR_CRACKER_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "RUMOR_CRACKER"
	GOURMET_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "GOURMET"
)

// All allowed values of ToolsBidSuggestV2ArticleCategory enum
var AllowedToolsBidSuggestV2ArticleCategoryEnumValues = []ToolsBidSuggestV2ArticleCategory{
	"COLLECTION",
	"BUSINESS",
	"HEALTH",
	"LOCAL",
	"STORIES",
	"PARENTING",
	"HISTORY",
	"FREAK",
	"PHOTOGRAPHY",
	"WEIGHT_LOSING",
	"LAWS",
	"GOVERNMENT",
	"PETS",
	"INTERNATIONAL",
	"FASHION",
	"EDUCATION",
	"TIPS",
	"ANIMATION",
	"MILITARY",
	"ESSAY",
	"ENTERTAINMENT",
	"TRAVEL",
	"TECHNOLOGY",
	"REGIMEN",
	"SPORTS",
	"DIGITAL",
	"MOVIE",
	"AMUSEMENT",
	"EXPLORE",
	"CULTURE",
	"SCIENCE",
	"CONSTELLATION",
	"CARS",
	"FINANCE",
	"EMOTION",
	"SOCIETY",
	"COMICS",
	"HOME",
	"DESIGN",
	"GRADUATES",
	"GAMES",
	"ESTATE",
	"WORKPLACE",
	"RUMOR_CRACKER",
	"GOURMET",
}

// NewToolsBidSuggestV2ArticleCategoryFromValue returns a pointer to a valid ToolsBidSuggestV2ArticleCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2ArticleCategoryFromValue(v string) (*ToolsBidSuggestV2ArticleCategory, error) {
	ev := ToolsBidSuggestV2ArticleCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2ArticleCategory: valid values are %v", v, AllowedToolsBidSuggestV2ArticleCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2ArticleCategory) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2ArticleCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_article_category value
func (v ToolsBidSuggestV2ArticleCategory) Ptr() *ToolsBidSuggestV2ArticleCategory {
	return &v
}
