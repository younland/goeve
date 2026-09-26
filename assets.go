package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterAssets Get character assets.
// GetCharacterAssets 获取角色资产.
//
// Route: GET /characters/{character_id}/assets/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/assets/ — 该路由缓存长达 3600 秒
// Scopes: esi-assets.read_assets.v1
// 权限: esi-assets.read_assets.v1
func (c *Client) GetCharacterAssets(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.CharacterAsset, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterAsset
	err := c.get(ctx, "/characters/{character_id}/assets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationAssets Get corporation assets.
// GetCorporationAssets 获取军团资产.
//
// Route: GET /corporations/{corporation_id}/assets/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/assets/ — 该路由缓存长达 3600 秒
// Scopes: esi-assets.read_corporation_assets.v1
// 权限: esi-assets.read_corporation_assets.v1
func (c *Client) GetCorporationAssets(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationAsset, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationAsset
	err := c.get(ctx, "/corporations/{corporation_id}/assets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterAssetLocations Get character asset locations.
// GetCharacterAssetLocations 获取角色资产位置.
//
// Route: POST /characters/{character_id}/assets/locations/
// 路由: POST /characters/{character_id}/assets/locations/
// Scopes: esi-assets.read_assets.v1
// 权限: esi-assets.read_assets.v1
func (c *Client) GetCharacterAssetLocations(ctx context.Context, token string, characterID int32, body []int64) ([]models.AssetLocation, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.AssetLocation
	err := c.post(ctx, "/characters/{character_id}/assets/locations/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterAssetNames Get character asset names.
// GetCharacterAssetNames 获取角色资产名称.
//
// Route: POST /characters/{character_id}/assets/names/
// 路由: POST /characters/{character_id}/assets/names/
// Scopes: esi-assets.read_assets.v1
// 权限: esi-assets.read_assets.v1
func (c *Client) GetCharacterAssetNames(ctx context.Context, token string, characterID int32, body []int64) ([]models.AssetName, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.AssetName
	err := c.post(ctx, "/characters/{character_id}/assets/names/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationAssetLocations Get corporation asset locations.
// GetCorporationAssetLocations 获取军团资产位置.
//
// Route: POST /corporations/{corporation_id}/assets/locations/
// 路由: POST /corporations/{corporation_id}/assets/locations/
// Scopes: esi-assets.read_corporation_assets.v1
// 权限: esi-assets.read_corporation_assets.v1
func (c *Client) GetCorporationAssetLocations(ctx context.Context, token string, corporationID int32, body []int64) ([]models.AssetLocation, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.AssetLocation
	err := c.post(ctx, "/corporations/{corporation_id}/assets/locations/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationAssetNames Get corporation asset names.
// GetCorporationAssetNames 获取军团资产名称.
//
// Route: POST /corporations/{corporation_id}/assets/names/
// 路由: POST /corporations/{corporation_id}/assets/names/
// Scopes: esi-assets.read_corporation_assets.v1
// 权限: esi-assets.read_corporation_assets.v1
func (c *Client) GetCorporationAssetNames(ctx context.Context, token string, corporationID int32, body []int64) ([]models.AssetName, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.AssetName
	err := c.post(ctx, "/corporations/{corporation_id}/assets/names/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
