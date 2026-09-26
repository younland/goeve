package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetWarId Get war information.
// GetWarId 获取战争信息.
//
// Route: GET /wars/{war_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /wars/{war_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetWarId(ctx context.Context, warId int32, params *models.GetWarIdParams) (*models.GetWarsWarIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"war_id": strconv.FormatInt(int64(warId), 10)}
	var result *models.GetWarsWarIdOk
	err := c.get(ctx, "/wars/{war_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWarIdKillmails List kills for a war.
// GetWarIdKillmails 列出某场战争的击杀.
//
// Route: GET /wars/{war_id}/killmails/ — This route is cached for up to 3600 seconds
// 路由: GET /wars/{war_id}/killmails/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetWarIdKillmails(ctx context.Context, warId int32, params *models.GetWarIdKillmailsParams) ([]models.GetWarsWarIdKillmails, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"war_id": strconv.FormatInt(int64(warId), 10)}
	var result []models.GetWarsWarIdKillmails
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
func (c *Client) GetWars(ctx context.Context, params *models.GetWarsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/wars/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
