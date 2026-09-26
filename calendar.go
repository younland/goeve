package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharacterCalendarEvents List calendar event summaries.
// GetCharacterCalendarEvents 列出日历事件摘要.
//
// Route: GET /characters/{character_id}/calendar/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/calendar/ — 该路由缓存长达 5 秒
// Scopes: esi-calendar.read_calendar_events.v1
// 权限: esi-calendar.read_calendar_events.v1
func (c *Client) GetCharacterCalendarEvents(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CalendarEventSummary, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CalendarEventSummary
	err := c.get(ctx, "/characters/{character_id}/calendar/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCalendarEvent Get an event.
// GetCalendarEvent 获取一个事件.
//
// Route: GET /characters/{character_id}/calendar/{event_id}/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/calendar/{event_id}/ — 该路由缓存长达 5 秒
// Scopes: esi-calendar.read_calendar_events.v1
// 权限: esi-calendar.read_calendar_events.v1
func (c *Client) GetCalendarEvent(ctx context.Context, characterID int32, eventID int32, opts ...RequestOption) (*models.CalendarEvent, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "event_id": strconv.FormatInt(int64(eventID), 10)}
	var result *models.CalendarEvent
	err := c.get(ctx, "/characters/{character_id}/calendar/{event_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCalendarEventAttendees Get attendees.
// GetCalendarEventAttendees 获取参与者.
//
// Route: GET /characters/{character_id}/calendar/{event_id}/attendees/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/calendar/{event_id}/attendees/ — 该路由缓存长达 600 秒
// Scopes: esi-calendar.read_calendar_events.v1
// 权限: esi-calendar.read_calendar_events.v1
func (c *Client) GetCalendarEventAttendees(ctx context.Context, characterID int32, eventID int32, opts ...RequestOption) ([]models.CalendarEventAttendee, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "event_id": strconv.FormatInt(int64(eventID), 10)}
	var result []models.CalendarEventAttendee
	err := c.get(ctx, "/characters/{character_id}/calendar/{event_id}/attendees/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RespondToCalendarEvent Respond to an event.
// RespondToCalendarEvent 响应日历事件.
//
// Route: PUT /characters/{character_id}/calendar/{event_id}/ — This route is cached for up to 5 seconds
// 路由: PUT /characters/{character_id}/calendar/{event_id}/ — 该路由缓存长达 5 秒
// Scopes: esi-calendar.respond_calendar_events.v1
// 权限: esi-calendar.respond_calendar_events.v1
func (c *Client) RespondToCalendarEvent(ctx context.Context, characterID int32, eventID int32, body *models.CalendarEventResponse, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "event_id": strconv.FormatInt(int64(eventID), 10)}
	if body == nil {
		return errBodyRequired
	}
	if err := c.put(ctx, "/characters/{character_id}/calendar/{event_id}/", pathParams, query, headers, body); err != nil {
		return err
	}
	return nil
}
