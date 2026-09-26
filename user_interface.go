package goeve

import (
	"context"
	"github.com/younland/goeve/models"
)

// PostUiAutopilotWaypoint Set Autopilot Waypoint.
// PostUiAutopilotWaypoint 设置自动导航路径点.
//
// Route: POST /ui/autopilot/waypoint/
// 路由: POST /ui/autopilot/waypoint/
// Scopes: esi-ui.write_waypoint.v1
// 权限: esi-ui.write_waypoint.v1
func (c *Client) PostUiAutopilotWaypoint(ctx context.Context, params *models.PostUiAutopilotWaypointParams) error {
	query, headers := params.Values()
	var pathParams map[string]string
	err := c.post(ctx, "/ui/autopilot/waypoint/", pathParams, query, headers, nil, nil)
	return err
}

// PostUiOpenwindowContract Open Contract Window.
// PostUiOpenwindowContract 打开合同窗口.
//
// Route: POST /ui/openwindow/contract/
// 路由: POST /ui/openwindow/contract/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) PostUiOpenwindowContract(ctx context.Context, params *models.PostUiOpenwindowContractParams) error {
	query, headers := params.Values()
	var pathParams map[string]string
	err := c.post(ctx, "/ui/openwindow/contract/", pathParams, query, headers, nil, nil)
	return err
}

// PostUiOpenwindowInformation Open Information Window.
// PostUiOpenwindowInformation 打开信息窗口.
//
// Route: POST /ui/openwindow/information/
// 路由: POST /ui/openwindow/information/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) PostUiOpenwindowInformation(ctx context.Context, params *models.PostUiOpenwindowInformationParams) error {
	query, headers := params.Values()
	var pathParams map[string]string
	err := c.post(ctx, "/ui/openwindow/information/", pathParams, query, headers, nil, nil)
	return err
}

// PostUiOpenwindowMarketdetails Open Market Details.
// PostUiOpenwindowMarketdetails 打开市场详情.
//
// Route: POST /ui/openwindow/marketdetails/
// 路由: POST /ui/openwindow/marketdetails/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) PostUiOpenwindowMarketdetails(ctx context.Context, params *models.PostUiOpenwindowMarketdetailsParams) error {
	query, headers := params.Values()
	var pathParams map[string]string
	err := c.post(ctx, "/ui/openwindow/marketdetails/", pathParams, query, headers, nil, nil)
	return err
}

// PostUiOpenwindowNewmail Open New Mail Window.
// PostUiOpenwindowNewmail 打开新邮件窗口.
//
// Route: POST /ui/openwindow/newmail/
// 路由: POST /ui/openwindow/newmail/
// Scopes: esi-ui.open_window.v1
// 权限: esi-ui.open_window.v1
func (c *Client) PostUiOpenwindowNewmail(ctx context.Context, body *models.PostUiOpenwindowNewmailNewMail, params *models.PostUiOpenwindowNewmailParams) error {
	query, headers := params.Values()
	var pathParams map[string]string
	if body == nil {
		return errBodyRequired
	}
	err := c.post(ctx, "/ui/openwindow/newmail/", pathParams, query, headers, body, nil)
	return err
}
