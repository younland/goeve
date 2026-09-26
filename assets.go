package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdAssets Get character assets.
// GetCharactersCharacterIdAssets 获取角色资产.
//
// Route: GET /characters/{character_id}/assets/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/assets/ — 该路由缓存长达 3600 秒
// Scopes: esi-assets.read_assets.v1
// 权限: esi-assets.read_assets.v1
func (c *Client) GetCharactersCharacterIdAssets(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdAssetsParams) ([]models.GetCharactersCharacterIdAssets, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdAssets
	err := c.get(ctx, "/characters/{character_id}/assets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdAssets Get corporation assets.
// GetCorporationsCorporationIdAssets 获取军团资产.
//
// Route: GET /corporations/{corporation_id}/assets/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/assets/ — 该路由缓存长达 3600 秒
// Scopes: esi-assets.read_corporation_assets.v1
// 权限: esi-assets.read_corporation_assets.v1
func (c *Client) GetCorporationsCorporationIdAssets(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdAssetsParams) ([]models.GetCorporationsCorporationIdAssets, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdAssets
	err := c.get(ctx, "/corporations/{corporation_id}/assets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCharactersCharacterIdAssetsLocations Get character asset locations.
// PostCharactersCharacterIdAssetsLocations 获取角色资产位置.
//
// Route: POST /characters/{character_id}/assets/locations/
// 路由: POST /characters/{character_id}/assets/locations/
// Scopes: esi-assets.read_assets.v1
// 权限: esi-assets.read_assets.v1
func (c *Client) PostCharactersCharacterIdAssetsLocations(ctx context.Context, characterId int32, body []int64, params *models.PostCharactersCharacterIdAssetsLocationsParams) ([]models.PostCharactersCharacterIdAssetsLocations, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.PostCharactersCharacterIdAssetsLocations
	err := c.post(ctx, "/characters/{character_id}/assets/locations/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCharactersCharacterIdAssetsNames Get character asset names.
// PostCharactersCharacterIdAssetsNames 获取角色资产名称.
//
// Route: POST /characters/{character_id}/assets/names/
// 路由: POST /characters/{character_id}/assets/names/
// Scopes: esi-assets.read_assets.v1
// 权限: esi-assets.read_assets.v1
func (c *Client) PostCharactersCharacterIdAssetsNames(ctx context.Context, characterId int32, body []int64, params *models.PostCharactersCharacterIdAssetsNamesParams) ([]models.PostCharactersCharacterIdAssetsNames, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.PostCharactersCharacterIdAssetsNames
	err := c.post(ctx, "/characters/{character_id}/assets/names/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCorporationsCorporationIdAssetsLocations Get corporation asset locations.
// PostCorporationsCorporationIdAssetsLocations 获取军团资产位置.
//
// Route: POST /corporations/{corporation_id}/assets/locations/
// 路由: POST /corporations/{corporation_id}/assets/locations/
// Scopes: esi-assets.read_corporation_assets.v1
// 权限: esi-assets.read_corporation_assets.v1
func (c *Client) PostCorporationsCorporationIdAssetsLocations(ctx context.Context, corporationId int32, body []int64, params *models.PostCorporationsCorporationIdAssetsLocationsParams) ([]models.PostCorporationsCorporationIdAssetsLocations, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.PostCorporationsCorporationIdAssetsLocations
	err := c.post(ctx, "/corporations/{corporation_id}/assets/locations/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCorporationsCorporationIdAssetsNames Get corporation asset names.
// PostCorporationsCorporationIdAssetsNames 获取军团资产名称.
//
// Route: POST /corporations/{corporation_id}/assets/names/
// 路由: POST /corporations/{corporation_id}/assets/names/
// Scopes: esi-assets.read_corporation_assets.v1
// 权限: esi-assets.read_corporation_assets.v1
func (c *Client) PostCorporationsCorporationIdAssetsNames(ctx context.Context, corporationId int32, body []int64, params *models.PostCorporationsCorporationIdAssetsNamesParams) ([]models.PostCorporationsCorporationIdAssetsNames, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.PostCorporationsCorporationIdAssetsNames
	err := c.post(ctx, "/corporations/{corporation_id}/assets/names/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
