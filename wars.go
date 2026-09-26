package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetWar Get war information.
// GetWar 获取战争信息.
//
// Route: GET /wars/{war_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /wars/{war_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetWar(ctx context.Context, warID int32, opts ...RequestOption) (*models.War, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"war_id": strconv.FormatInt(int64(warID), 10)}
	var result *models.War
	err := c.get(ctx, "/wars/{war_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWarKillmails List kills for a war.
// GetWarKillmails 列出某场战争的击杀.
//
// Route: GET /wars/{war_id}/killmails/ — This route is cached for up to 3600 seconds
// 路由: GET /wars/{war_id}/killmails/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetWarKillmails(ctx context.Context, warID int32, opts ...RequestOption) ([]models.KillmailRef, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"war_id": strconv.FormatInt(int64(warID), 10)}
	var result []models.KillmailRef
	err := c.get(ctx, "/wars/{war_id}/killmails/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWars List wars.
// GetWars 列出战争.
//
// Route: GET /wars/ — This route is cached for up to 3600 seconds
// 路由: GET /wars/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetWars(ctx context.Context, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/wars/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
