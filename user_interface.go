package goeve

import (
	"context"
	"strconv"

	"github.com/younland/goeve/models"
)

// SetAutopilotWaypoint Set Autopilot Waypoint.
// SetAutopilotWaypoint 设置自动导航路径点.
//
// Route: POST /ui/autopilot/waypoint/
// 路由: POST /ui/autopilot/waypoint/
// Scopes: esi-ui.write_waypoint.v1
// 权限: esi-ui.write_waypoint.v1
func (c *Client) SetAutopilotWaypoint(ctx context.Context, addToBeginning bool, clearOtherWaypoints bool, destinationID int64, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	query.Set("add_to_beginning", strconv.FormatBool(addToBeginning))
	query.Set("clear_other_waypoints", strconv.FormatBool(clearOtherWaypoints))
	query.Set("destination_id", strconv.FormatInt(destinationID, 10))
	var pathParams map[string]string
	err := c.post(ctx, "/ui/autopilot/waypoint/", pathParams, query, headers, nil, nil)
	return err
}

// OpenContractWindow Open Contract Window.
// OpenContractWindow 打开合同窗口.
//
// Route: POST /ui/openwindow/contract/
// 路由: POST /ui/openwindow/contract/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) OpenContractWindow(ctx context.Context, contractID string, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	query.Set("contract_id", contractID)
	var pathParams map[string]string
	err := c.post(ctx, "/ui/openwindow/contract/", pathParams, query, headers, nil, nil)
	return err
}

// OpenInformationWindow Open Information Window.
// OpenInformationWindow 打开信息窗口.
//
// Route: POST /ui/openwindow/information/
// 路由: POST /ui/openwindow/information/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) OpenInformationWindow(ctx context.Context, targetID string, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	query.Set("target_id", targetID)
	var pathParams map[string]string
	err := c.post(ctx, "/ui/openwindow/information/", pathParams, query, headers, nil, nil)
	return err
}

// OpenMarketDetails Open Market Details.
// OpenMarketDetails 打开市场详情.
//
// Route: POST /ui/openwindow/marketdetails/
// 路由: POST /ui/openwindow/marketdetails/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) OpenMarketDetails(ctx context.Context, typeID string, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	query.Set("type_id", typeID)
	var pathParams map[string]string
	err := c.post(ctx, "/ui/openwindow/marketdetails/", pathParams, query, headers, nil, nil)
	return err
}

// OpenNewMailWindow Open New Mail Window.
// OpenNewMailWindow 打开新邮件窗口.
//
// Route: POST /ui/openwindow/newmail/
// 路由: POST /ui/openwindow/newmail/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) OpenNewMailWindow(ctx context.Context, body *models.NewMailRequest, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	if body == nil {
		return errBodyRequired
	}
	err := c.post(ctx, "/ui/openwindow/newmail/", pathParams, query, headers, body, nil)
	return err
}
