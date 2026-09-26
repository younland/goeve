package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetOriginDestination Get route.
// GetOriginDestination 获取航线.
//
// Route: GET /route/{origin}/{destination}/ — This route is cached for up to 86400 seconds
// 路由: GET /route/{origin}/{destination}/ — 该路由缓存长达 86400 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetOriginDestination(ctx context.Context, destination int32, origin int32, params *models.GetOriginDestinationParams) ([]int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"destination": strconv.FormatInt(int64(destination), 10), "origin": strconv.FormatInt(int64(origin), 10)}
	var result []int32
	err := c.get(ctx, "/route/{origin}/{destination}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
