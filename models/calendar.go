package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetCharactersCharacterIdCalendar event.
// GetCharactersCharacterIdCalendar 事件.
type GetCharactersCharacterIdCalendar struct {
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

// GetCharactersCharacterIdCalendarEventIdAttendees character_id and response of an attendee.
// GetCharactersCharacterIdCalendarEventIdAttendees 角色 ID 及与会者的响应.
type GetCharactersCharacterIdCalendarEventIdAttendees struct {
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// EventResponse event_response string.
	// EventResponse 事件响应字符串.
	// Enum values: "declined", "not_responded", "accepted", "tentative".
	EventResponse string `json:"event_response"`
}

// GetCharactersCharacterIdCalendarEventId Full details of a specific event.
// GetCharactersCharacterIdCalendarEventId 某个特定事件的完整详情.
type GetCharactersCharacterIdCalendarEventId struct {
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

// PutCharactersCharacterIdCalendarEventIdResponse response object.
// PutCharactersCharacterIdCalendarEventIdResponse 响应 object.
type PutCharactersCharacterIdCalendarEventIdResponse struct {
	// Response response string.
	// Response 响应 string.
	// Enum values: "accepted", "declined", "tentative".
	Response string `json:"response"`
}

// GetCharactersCharacterIdCalendarEventIdAttendeesParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdCalendarEventIdAttendeesParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdCalendarEventIdAttendeesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdCalendarEventIdAttendeesParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharactersCharacterIdCalendarEventIdParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdCalendarEventIdParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdCalendarEventIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdCalendarEventIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharactersCharacterIdCalendarParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdCalendarParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdCalendarParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// FromEvent The event ID to retrieve events from.
	// FromEvent 用于检索事件的事件 ID.
	FromEvent *int32
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdCalendarParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.FromEvent != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("from_event", strconv.FormatInt(int64(*p.FromEvent), 10))
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PutCharactersCharacterIdCalendarEventIdParams holds the optional query and header parameters of the request.
// PutCharactersCharacterIdCalendarEventIdParams 保存请求的可选查询与头部参数。
type PutCharactersCharacterIdCalendarEventIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PutCharactersCharacterIdCalendarEventIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}
