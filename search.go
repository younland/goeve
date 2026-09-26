package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// SearchEntities Search on a string.
// SearchEntities 对字符串进行搜索.
//
// Route: GET /characters/{character_id}/search/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/search/ — 该路由缓存长达 3600 秒
// Scopes: esi-search.search_structures.v1
// 权限: esi-search.search_structures.v1
func (c *Client) SearchEntities(ctx context.Context, characterID int32, categories []string, search string, token string, strict bool, ifNoneMatch ...string) (*models.SearchResult, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if strict {
		query.Set("strict", strconv.FormatBool(strict))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	for _, v := range categories {
		query.Add("categories", v)
	}
	query.Set("search", search)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.SearchResult
	err := c.get(ctx, "/characters/{character_id}/search/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
