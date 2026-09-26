package goeve

import (
	"context"
	"strconv"

	"github.com/younland/goeve/models"
)

// GetCharacterClones Get clones.
// GetCharacterClones 获取克隆.
//
// Route: GET /characters/{character_id}/clones/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/clones/ — 该路由缓存长达 120 秒
// Scopes: esi-clones.read_clones.v1
// 权限: esi-clones.read_clones.v1
func (c *Client) GetCharacterClones(ctx context.Context, characterID int32, opts ...RequestOption) (*models.Clones, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.Clones
	err := c.get(ctx, "/characters/{character_id}/clones/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterImplants Get active implants.
// GetCharacterImplants 获取已植入的脑插.
//
// Route: GET /characters/{character_id}/implants/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/implants/ — 该路由缓存长达 120 秒
// Scopes: esi-clones.read_implants.v1
// 权限: esi-clones.read_implants.v1
func (c *Client) GetCharacterImplants(ctx context.Context, characterID int32, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []int32
	err := c.get(ctx, "/characters/{character_id}/implants/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
