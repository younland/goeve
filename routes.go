package goeve

import (
	"context"
	"net/url"
	"strconv"
)

// GetRoute Get route.
// GetRoute 获取航线.
//
// Route: GET /route/{origin}/{destination}/ — This route is cached for up to 86400 seconds
// 路由: GET /route/{origin}/{destination}/ — 该路由缓存长达 86400 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetRoute(ctx context.Context, destination int32, origin int32, avoid []int32, connections [][]int32, flag string, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	for _, v := range avoid {
		query.Add("avoid", strconv.FormatInt(int64(v), 10))
	}
	for _, pair := range connections {
		if len(pair) == 2 {
			query.Add("connections", strconv.FormatInt(int64(pair[0]), 10)+"|"+strconv.FormatInt(int64(pair[1]), 10))
		}
	}
	if flag != "" {
		query.Set("flag", flag)
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"destination": strconv.FormatInt(int64(destination), 10), "origin": strconv.FormatInt(int64(origin), 10)}
	var result []int32
	err := c.get(ctx, "/route/{origin}/{destination}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
