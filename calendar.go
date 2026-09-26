package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdCalendar List calendar event summaries.
// GetCharactersCharacterIdCalendar 列出日历事件摘要.
//
// Route: GET /characters/{character_id}/calendar/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/calendar/ — 该路由缓存长达 5 秒
// Scopes: esi-calendar.read_calendar_events.v1
// 权限: esi-calendar.read_calendar_events.v1
func (c *Client) GetCharactersCharacterIdCalendar(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdCalendarParams) ([]models.GetCharactersCharacterIdCalendar, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdCalendar
	err := c.get(ctx, "/characters/{character_id}/calendar/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdCalendarEventId Get an event.
// GetCharactersCharacterIdCalendarEventId 获取一个事件.
//
// Route: GET /characters/{character_id}/calendar/{event_id}/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/calendar/{event_id}/ — 该路由缓存长达 5 秒
// Scopes: esi-calendar.read_calendar_events.v1
// 权限: esi-calendar.read_calendar_events.v1
func (c *Client) GetCharactersCharacterIdCalendarEventId(ctx context.Context, characterId int32, eventId int32, params *models.GetCharactersCharacterIdCalendarEventIdParams) (*models.GetCharactersCharacterIdCalendarEventId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "event_id": strconv.FormatInt(int64(eventId), 10)}
	var result *models.GetCharactersCharacterIdCalendarEventId
	err := c.get(ctx, "/characters/{character_id}/calendar/{event_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdCalendarEventIdAttendees Get attendees.
// GetCharactersCharacterIdCalendarEventIdAttendees 获取参与者.
//
// Route: GET /characters/{character_id}/calendar/{event_id}/attendees/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/calendar/{event_id}/attendees/ — 该路由缓存长达 600 秒
// Scopes: esi-calendar.read_calendar_events.v1
// 权限: esi-calendar.read_calendar_events.v1
func (c *Client) GetCharactersCharacterIdCalendarEventIdAttendees(ctx context.Context, characterId int32, eventId int32, params *models.GetCharactersCharacterIdCalendarEventIdAttendeesParams) ([]models.GetCharactersCharacterIdCalendarEventIdAttendees, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "event_id": strconv.FormatInt(int64(eventId), 10)}
	var result []models.GetCharactersCharacterIdCalendarEventIdAttendees
	err := c.get(ctx, "/characters/{character_id}/calendar/{event_id}/attendees/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutCharactersCharacterIdCalendarEventId Respond to an event.
// PutCharactersCharacterIdCalendarEventId 响应日历事件.
//
// Route: PUT /characters/{character_id}/calendar/{event_id}/ — This route is cached for up to 5 seconds
// 路由: PUT /characters/{character_id}/calendar/{event_id}/ — 该路由缓存长达 5 秒
// Scopes: esi-calendar.respond_calendar_events.v1
// 权限: esi-calendar.respond_calendar_events.v1
func (c *Client) PutCharactersCharacterIdCalendarEventId(ctx context.Context, characterId int32, eventId int32, body *models.PutCharactersCharacterIdCalendarEventIdResponse, params *models.PutCharactersCharacterIdCalendarEventIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "event_id": strconv.FormatInt(int64(eventId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/characters/{character_id}/calendar/{event_id}/", pathParams, query, headers, body)
	return err
}
