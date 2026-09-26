package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdSearch Search on a string.
// GetCharactersCharacterIdSearch 对字符串进行搜索.
//
// Route: GET /characters/{character_id}/search/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/search/ — 该路由缓存长达 3600 秒
// Scopes: esi-search.search_structures.v1
// 权限: esi-search.search_structures.v1
func (c *Client) GetCharactersCharacterIdSearch(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdSearchParams) (*models.GetCharactersCharacterIdSearch, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdSearch
	err := c.get(ctx, "/characters/{character_id}/search/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
