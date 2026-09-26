package models

import (
	"time"
)

// CalendarEventSummary event.
// CalendarEventSummary 事件.
type CalendarEventSummary struct {
	// EventDate event_date string.
	// EventDate 事件日期字符串.
	EventDate time.Time `json:"event_date"`
	// EventId event_id integer.
	// EventId 事件 ID 整数.
	EventId int32 `json:"event_id"`
	// EventResponse event_response string.
	// EventResponse 事件响应字符串.
	// Enum values: "declined", "not_responded", "accepted", "tentative".
	EventResponse string `json:"event_response"`
	// Importance importance integer.
	// Importance 重要度整数.
	Importance int32 `json:"importance"`
	// Title title string.
	// Title 标题字符串.
	Title string `json:"title"`
}

// CalendarEventAttendee character_id and response of an attendee.
// CalendarEventAttendee 角色 ID 及与会者的响应.
type CalendarEventAttendee struct {
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// EventResponse event_response string.
	// EventResponse 事件响应字符串.
	// Enum values: "declined", "not_responded", "accepted", "tentative".
	EventResponse string `json:"event_response"`
}

// CalendarEvent Full details of a specific event.
// CalendarEvent 某个特定事件的完整详情.
type CalendarEvent struct {
	// Date date string.
	// Date 日期字符串.
	Date time.Time `json:"date"`
	// Duration Length in minutes.
	// Duration 时长（分钟）
	Duration int32 `json:"duration"`
	// EventId event_id integer.
	// EventId 事件 ID 整数.
	EventId int32 `json:"event_id"`
	// Importance importance integer.
	// Importance 重要度整数.
	Importance int32 `json:"importance"`
	// OwnerId owner_id integer.
	// OwnerId owner_id 整数.
	OwnerId int32 `json:"owner_id"`
	// OwnerName owner_name string.
	// OwnerName owner_name 字符串.
	OwnerName string `json:"owner_name"`
	// OwnerType owner_type string.
	// OwnerType owner_type 字符串.
	// Enum values: "eve_server", "corporation", "faction", "character", "alliance".
	OwnerType string `json:"owner_type"`
	// Response response string.
	// Response 响应 string.
	Response string `json:"response"`
	// Text text string.
	// Text 文本字符串.
	Text string `json:"text"`
	// Title title string.
	// Title 标题字符串.
	Title string `json:"title"`
}

// CalendarEventResponse response object.
// CalendarEventResponse 响应 object.
type CalendarEventResponse struct {
	// Response response string.
	// Response 响应 string.
	// Enum values: "accepted", "declined", "tentative".
	Response string `json:"response"`
}
