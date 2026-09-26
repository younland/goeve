package goeve

import (
	"context"
	"github.com/younland/goeve/models"
)

// GetPrices List insurance levels.
// GetPrices 列出保险等级.
//
// Route: GET /insurance/prices/ — This route is cached for up to 3600 seconds
// 路由: GET /insurance/prices/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPrices(ctx context.Context, opts ...RequestOption) ([]models.InsurancePrice, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []models.InsurancePrice
	err := c.get(ctx, "/insurance/prices/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
