package goeve

import (
	"context"
	"github.com/younland/goeve/models"
)

// GetSovereigntyCampaigns List sovereignty campaigns.
// GetSovereigntyCampaigns 列出主权战役.
//
// Route: GET /sovereignty/campaigns/ — This route is cached for up to 5 seconds
// 路由: GET /sovereignty/campaigns/ — 该路由缓存长达 5 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSovereigntyCampaigns(ctx context.Context, opts ...RequestOption) ([]models.SovereigntyCampaign, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []models.SovereigntyCampaign
	err := c.get(ctx, "/sovereignty/campaigns/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSovereigntyMap List sovereignty of systems.
// GetSovereigntyMap 列出各星系的主权归属.
//
// Route: GET /sovereignty/map/ — This route is cached for up to 3600 seconds
// 路由: GET /sovereignty/map/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSovereigntyMap(ctx context.Context, opts ...RequestOption) ([]models.SovereigntySystem, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []models.SovereigntySystem
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
func (c *Client) GetSovereigntyStructures(ctx context.Context, opts ...RequestOption) ([]models.SovereigntyStructure, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []models.SovereigntyStructure
	err := c.get(ctx, "/sovereignty/structures/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
