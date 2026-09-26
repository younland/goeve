package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdPlanets Get colonies.
// GetCharactersCharacterIdPlanets 获取殖民地.
//
// Route: GET /characters/{character_id}/planets/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/planets/ — 该路由缓存长达 600 秒
// Scopes: esi-planets.manage_planets.v1
// 权限: esi-planets.manage_planets.v1
func (c *Client) GetCharactersCharacterIdPlanets(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdPlanetsParams) ([]models.GetCharactersCharacterIdPlanets, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdPlanets
	err := c.get(ctx, "/characters/{character_id}/planets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdPlanetsPlanetId Get colony layout.
// GetCharactersCharacterIdPlanetsPlanetId 获取殖民地布局.
//
// Route: GET /characters/{character_id}/planets/{planet_id}/
// 路由: GET /characters/{character_id}/planets/{planet_id}/
// Scopes: esi-planets.manage_planets.v1
// 权限: esi-planets.manage_planets.v1
func (c *Client) GetCharactersCharacterIdPlanetsPlanetId(ctx context.Context, characterId int32, planetId int32, params *models.GetCharactersCharacterIdPlanetsPlanetIdParams) (*models.GetCharactersCharacterIdPlanetsPlanetIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "planet_id": strconv.FormatInt(int64(planetId), 10)}
	var result *models.GetCharactersCharacterIdPlanetsPlanetIdOk
	err := c.get(ctx, "/characters/{character_id}/planets/{planet_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdCustomsOffices List corporation customs offices.
// GetCorporationsCorporationIdCustomsOffices 列出军团海关办公室.
//
// Route: GET /corporations/{corporation_id}/customs_offices/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/customs_offices/ — 该路由缓存长达 3600 秒
// Scopes: esi-planets.read_customs_offices.v1
// 权限: esi-planets.read_customs_offices.v1
func (c *Client) GetCorporationsCorporationIdCustomsOffices(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdCustomsOfficesParams) ([]models.GetCorporationsCorporationIdCustomsOffices, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdCustomsOffices
	err := c.get(ctx, "/corporations/{corporation_id}/customs_offices/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseSchematicsSchematicId Get schematic information.
// GetUniverseSchematicsSchematicId 获取示意图信息.
//
// Route: GET /universe/schematics/{schematic_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/schematics/{schematic_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseSchematicsSchematicId(ctx context.Context, schematicId int32, params *models.GetUniverseSchematicsSchematicIdParams) (*models.GetUniverseSchematicsSchematicIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"schematic_id": strconv.FormatInt(int64(schematicId), 10)}
	var result *models.GetUniverseSchematicsSchematicIdOk
	err := c.get(ctx, "/universe/schematics/{schematic_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
