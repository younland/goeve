package goeve

import (
	"context"
	"github.com/younland/goeve/models"
)

// GetStatus Retrieve the uptime and player counts.
// GetStatus 获取服务器运行时间和玩家数量.
//
// Route: GET /status/ — This route is cached for up to 30 seconds
// 路由: GET /status/ — 该路由缓存长达 30 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetStatus(ctx context.Context, params *models.GetStatusParams) (*models.GetStatus, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result *models.GetStatus
	err := c.get(ctx, "/status/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
