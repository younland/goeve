package goeve

import (
	"context"
	"github.com/younland/goeve/models"
)

// GetCampaigns List sovereignty campaigns.
// GetCampaigns 列出主权战役.
//
// Route: GET /sovereignty/campaigns/ — This route is cached for up to 5 seconds
// 路由: GET /sovereignty/campaigns/ — 该路由缓存长达 5 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCampaigns(ctx context.Context, params *models.GetCampaignsParams) ([]models.GetSovereigntyCampaigns, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetSovereigntyCampaigns
	err := c.get(ctx, "/sovereignty/campaigns/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMap List sovereignty of systems.
// GetMap 列出各星系的主权归属.
//
// Route: GET /sovereignty/map/ — This route is cached for up to 3600 seconds
// 路由: GET /sovereignty/map/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMap(ctx context.Context, params *models.GetMapParams) ([]models.GetSovereigntyMap, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetSovereigntyMap
	err := c.get(ctx, "/sovereignty/map/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSovereigntyStructures List sovereignty structures.
// GetSovereigntyStructures 列出主权建筑.
//
// Route: GET /sovereignty/structures/ — This route is cached for up to 120 seconds
// 路由: GET /sovereignty/structures/ — 该路由缓存长达 120 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSovereigntyStructures(ctx context.Context, params *models.GetStructuresParams) ([]models.GetSovereigntyStructures, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetSovereigntyStructures
	err := c.get(ctx, "/sovereignty/structures/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
