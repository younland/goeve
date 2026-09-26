package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdClones Get clones.
// GetCharactersCharacterIdClones 获取克隆.
//
// Route: GET /characters/{character_id}/clones/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/clones/ — 该路由缓存长达 120 秒
// Scopes: esi-clones.read_clones.v1
// 权限: esi-clones.read_clones.v1
func (c *Client) GetCharactersCharacterIdClones(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdClonesParams) (*models.GetCharactersCharacterIdClonesOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdClonesOk
	err := c.get(ctx, "/characters/{character_id}/clones/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdImplants Get active implants.
// GetCharactersCharacterIdImplants 获取已植入的脑插.
//
// Route: GET /characters/{character_id}/implants/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/implants/ — 该路由缓存长达 120 秒
// Scopes: esi-clones.read_implants.v1
// 权限: esi-clones.read_implants.v1
func (c *Client) GetCharactersCharacterIdImplants(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdImplantsParams) ([]int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []int32
	err := c.get(ctx, "/characters/{character_id}/implants/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
