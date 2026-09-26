package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterColonies Get colonies.
// GetCharacterColonies 获取殖民地.
//
// Route: GET /characters/{character_id}/planets/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/planets/ — 该路由缓存长达 600 秒
// Scopes: esi-planets.manage_planets.v1
// 权限: esi-planets.manage_planets.v1
func (c *Client) GetCharacterColonies(ctx context.Context, characterID int32, token string, ifNoneMatch string) ([]models.Colony, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.Colony
	err := c.get(ctx, "/characters/{character_id}/planets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterColonyLayout Get colony layout.
// GetCharacterColonyLayout 获取殖民地布局.
//
// Route: GET /characters/{character_id}/planets/{planet_id}/
// 路由: GET /characters/{character_id}/planets/{planet_id}/
// Scopes: esi-planets.manage_planets.v1
// 权限: esi-planets.manage_planets.v1
func (c *Client) GetCharacterColonyLayout(ctx context.Context, characterID int32, planetID int32, token string, ifNoneMatch string) (*models.ColonyLayout, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "planet_id": strconv.FormatInt(int64(planetID), 10)}
	var result *models.ColonyLayout
	err := c.get(ctx, "/characters/{character_id}/planets/{planet_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationCustomsOffices List corporation customs offices.
// GetCorporationCustomsOffices 列出军团海关办公室.
//
// Route: GET /corporations/{corporation_id}/customs_offices/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/customs_offices/ — 该路由缓存长达 3600 秒
// Scopes: esi-planets.read_customs_offices.v1
// 权限: esi-planets.read_customs_offices.v1
func (c *Client) GetCorporationCustomsOffices(ctx context.Context, corporationID int32, token string, page int32, ifNoneMatch string) ([]models.CustomsOffice, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CustomsOffice
	err := c.get(ctx, "/corporations/{corporation_id}/customs_offices/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSchematicInformation Get schematic information.
// GetSchematicInformation 获取示意图信息.
//
// Route: GET /universe/schematics/{schematic_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/schematics/{schematic_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSchematicInformation(ctx context.Context, schematicID int32, ifNoneMatch string) (*models.Schematic, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"schematic_id": strconv.FormatInt(int64(schematicID), 10)}
	var result *models.Schematic
	err := c.get(ctx, "/universe/schematics/{schematic_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
