/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataRoomInfoInner struct for QianchuanAdDetailGetV10ResponseDataRoomInfoInner
type QianchuanAdDetailGetV10ResponseDataRoomInfoInner struct {
	//
	AnchorAvatar *string `json:"anchor_avatar,omitempty"`
	//
	AnchorId *int64 `json:"anchor_id,omitempty"`
	//
	AnchorName *string                                        `json:"anchor_name,omitempty"`
	RoomStatus *QianchuanAdDetailGetV10DataRoomInfoRoomStatus `json:"room_status,omitempty"`
	//
	RoomTitle *string `json:"room_title,omitempty"`
}
