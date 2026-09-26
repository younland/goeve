package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// DeleteCharactersCharacterIdFittingsFittingId Delete fitting.
// DeleteCharactersCharacterIdFittingsFittingId 删除装配.
//
// Route: DELETE /characters/{character_id}/fittings/{fitting_id}/
// 路由: DELETE /characters/{character_id}/fittings/{fitting_id}/
// Scopes: esi-fittings.write_fittings.v1
// 权限: esi-fittings.write_fittings.v1
func (c *Client) DeleteCharactersCharacterIdFittingsFittingId(ctx context.Context, characterId int32, fittingId int32, params *models.DeleteCharactersCharacterIdFittingsFittingIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "fitting_id": strconv.FormatInt(int64(fittingId), 10)}
	err := c.delete(ctx, "/characters/{character_id}/fittings/{fitting_id}/", pathParams, query, headers)
	return err
}

// GetCharactersCharacterIdFittings Get fittings.
// GetCharactersCharacterIdFittings 获取装配.
//
// Route: GET /characters/{character_id}/fittings/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/fittings/ — 该路由缓存长达 300 秒
// Scopes: esi-fittings.read_fittings.v1
// 权限: esi-fittings.read_fittings.v1
func (c *Client) GetCharactersCharacterIdFittings(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdFittingsParams) ([]models.GetCharactersCharacterIdFittings, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdFittings
	err := c.get(ctx, "/characters/{character_id}/fittings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCharactersCharacterIdFittings Create fitting.
// PostCharactersCharacterIdFittings 创建装配.
//
// Route: POST /characters/{character_id}/fittings/
// 路由: POST /characters/{character_id}/fittings/
// Scopes: esi-fittings.write_fittings.v1
// 权限: esi-fittings.write_fittings.v1
func (c *Client) PostCharactersCharacterIdFittings(ctx context.Context, characterId int32, body *models.PostCharactersCharacterIdFittingsFitting, params *models.PostCharactersCharacterIdFittingsParams) (*models.PostCharactersCharacterIdFittingsCreated, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result *models.PostCharactersCharacterIdFittingsCreated
	err := c.post(ctx, "/characters/{character_id}/fittings/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
