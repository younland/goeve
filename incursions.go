package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
)

// GetIncursions List incursions.
// GetIncursions 列出入侵活动.
//
// Route: GET /incursions/ — This route is cached for up to 300 seconds
// 路由: GET /incursions/ — 该路由缓存长达 300 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetIncursions(ctx context.Context, ifNoneMatch ...string) ([]models.Incursion, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result []models.Incursion
	err := c.get(ctx, "/incursions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
