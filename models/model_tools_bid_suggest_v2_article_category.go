/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
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
	COMICS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "COMICS"
	FASHION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FASHION"
	GOURMET_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "GOURMET"
	CARS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "CARS"
	TRAVEL_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "TRAVEL"
	RUMOR_CRACKER_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "RUMOR_CRACKER"
	AMUSEMENT_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "AMUSEMENT"
	ENTERTAINMENT_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "ENTERTAINMENT"
	SCIENCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SCIENCE"
	GAMES_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "GAMES"
	SPORTS_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "SPORTS"
	SOCIETY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "SOCIETY"
	MILITARY_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "MILITARY"
	ESTATE_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "ESTATE"
	WEIGHT_LOSING_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "WEIGHT_LOSING"
	EDUCATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "EDUCATION"
	HEALTH_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "HEALTH"
	TECHNOLOGY_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "TECHNOLOGY"
	WORKPLACE_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "WORKPLACE"
	DESIGN_ToolsBidSuggestV2ArticleCategory        ToolsBidSuggestV2ArticleCategory = "DESIGN"
	HOME_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "HOME"
	FINANCE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "FINANCE"
	PHOTOGRAPHY_ToolsBidSuggestV2ArticleCategory   ToolsBidSuggestV2ArticleCategory = "PHOTOGRAPHY"
	ESSAY_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "ESSAY"
	TIPS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "TIPS"
	DIGITAL_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "DIGITAL"
	CONSTELLATION_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "CONSTELLATION"
	GRADUATES_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "GRADUATES"
	LOCAL_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "LOCAL"
	COLLECTION_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "COLLECTION"
	REGIMEN_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "REGIMEN"
	ANIMATION_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "ANIMATION"
	GOVERNMENT_ToolsBidSuggestV2ArticleCategory    ToolsBidSuggestV2ArticleCategory = "GOVERNMENT"
	INTERNATIONAL_ToolsBidSuggestV2ArticleCategory ToolsBidSuggestV2ArticleCategory = "INTERNATIONAL"
	MOVIE_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "MOVIE"
	LAWS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "LAWS"
	BUSINESS_ToolsBidSuggestV2ArticleCategory      ToolsBidSuggestV2ArticleCategory = "BUSINESS"
	EXPLORE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EXPLORE"
	FREAK_ToolsBidSuggestV2ArticleCategory         ToolsBidSuggestV2ArticleCategory = "FREAK"
	CULTURE_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "CULTURE"
	PETS_ToolsBidSuggestV2ArticleCategory          ToolsBidSuggestV2ArticleCategory = "PETS"
	PARENTING_ToolsBidSuggestV2ArticleCategory     ToolsBidSuggestV2ArticleCategory = "PARENTING"
	STORIES_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "STORIES"
	EMOTION_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "EMOTION"
	HISTORY_ToolsBidSuggestV2ArticleCategory       ToolsBidSuggestV2ArticleCategory = "HISTORY"
)

// All allowed values of ToolsBidSuggestV2ArticleCategory enum
var AllowedToolsBidSuggestV2ArticleCategoryEnumValues = []ToolsBidSuggestV2ArticleCategory{
	"COMICS",
	"FASHION",
	"GOURMET",
	"CARS",
	"TRAVEL",
	"RUMOR_CRACKER",
	"AMUSEMENT",
	"ENTERTAINMENT",
	"SCIENCE",
	"GAMES",
	"SPORTS",
	"SOCIETY",
	"MILITARY",
	"ESTATE",
	"WEIGHT_LOSING",
	"EDUCATION",
	"HEALTH",
	"TECHNOLOGY",
	"WORKPLACE",
	"DESIGN",
	"HOME",
	"FINANCE",
	"PHOTOGRAPHY",
	"ESSAY",
	"TIPS",
	"DIGITAL",
	"CONSTELLATION",
	"GRADUATES",
	"LOCAL",
	"COLLECTION",
	"REGIMEN",
	"ANIMATION",
	"GOVERNMENT",
	"INTERNATIONAL",
	"MOVIE",
	"LAWS",
	"BUSINESS",
	"EXPLORE",
	"FREAK",
	"CULTURE",
	"PETS",
	"PARENTING",
	"STORIES",
	"EMOTION",
	"HISTORY",
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
