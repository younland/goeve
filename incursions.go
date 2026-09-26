package goeve

import (
	"context"
	"github.com/younland/goeve/models"
)

// GetIncursions List incursions.
// GetIncursions 列出入侵活动.
//
// Route: GET /incursions/ — This route is cached for up to 300 seconds
// 路由: GET /incursions/ — 该路由缓存长达 300 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetIncursions(ctx context.Context, params *models.GetIncursionsParams) ([]models.GetIncursions, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetIncursions
	err := c.get(ctx, "/incursions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
