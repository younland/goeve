package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetAncestries Get ancestries.
// GetAncestries 获取血统.
//
// Route: GET /universe/ancestries/
// 路由: GET /universe/ancestries/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAncestries(ctx context.Context, params *models.GetAncestriesParams) ([]models.GetUniverseAncestries, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetUniverseAncestries
	err := c.get(ctx, "/universe/ancestries/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAsteroidBeltsAsteroidBeltId Get asteroid belt information.
// GetAsteroidBeltsAsteroidBeltId 获取小行星带信息.
//
// Route: GET /universe/asteroid_belts/{asteroid_belt_id}/
// 路由: GET /universe/asteroid_belts/{asteroid_belt_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAsteroidBeltsAsteroidBeltId(ctx context.Context, asteroidBeltId int32, params *models.GetAsteroidBeltsAsteroidBeltIdParams) (*models.GetUniverseAsteroidBeltsAsteroidBeltId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"asteroid_belt_id": strconv.FormatInt(int64(asteroidBeltId), 10)}
	var result *models.GetUniverseAsteroidBeltsAsteroidBeltId
	err := c.get(ctx, "/universe/asteroid_belts/{asteroid_belt_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetBloodlines Get bloodlines.
// GetBloodlines 获取种族血脉.
//
// Route: GET /universe/bloodlines/
// 路由: GET /universe/bloodlines/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetBloodlines(ctx context.Context, params *models.GetBloodlinesParams) ([]models.GetUniverseBloodlines, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetUniverseBloodlines
	err := c.get(ctx, "/universe/bloodlines/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCategories Get item categories.
// GetCategories 获取物品类别.
//
// Route: GET /universe/categories/
// 路由: GET /universe/categories/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCategories(ctx context.Context, params *models.GetCategoriesParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/categories/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCategoriesCategoryId Get item category information.
// GetCategoriesCategoryId 获取物品类别信息.
//
// Route: GET /universe/categories/{category_id}/
// 路由: GET /universe/categories/{category_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCategoriesCategoryId(ctx context.Context, categoryId int32, params *models.GetCategoriesCategoryIdParams) (*models.GetUniverseCategoriesCategoryId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"category_id": strconv.FormatInt(int64(categoryId), 10)}
	var result *models.GetUniverseCategoriesCategoryId
	err := c.get(ctx, "/universe/categories/{category_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConstellations Get constellations.
// GetConstellations 获取星座.
//
// Route: GET /universe/constellations/
// 路由: GET /universe/constellations/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetConstellations(ctx context.Context, params *models.GetConstellationsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/constellations/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConstellationsConstellationId Get constellation information.
// GetConstellationsConstellationId 获取星座信息.
//
// Route: GET /universe/constellations/{constellation_id}/
// 路由: GET /universe/constellations/{constellation_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetConstellationsConstellationId(ctx context.Context, constellationId int32, params *models.GetConstellationsConstellationIdParams) (*models.GetUniverseConstellationsConstellationId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"constellation_id": strconv.FormatInt(int64(constellationId), 10)}
	var result *models.GetUniverseConstellationsConstellationId
	err := c.get(ctx, "/universe/constellations/{constellation_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFactions Get factions.
// GetFactions 获取势力.
//
// Route: GET /universe/factions/
// 路由: GET /universe/factions/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFactions(ctx context.Context, params *models.GetFactionsParams) ([]models.GetUniverseFactions, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetUniverseFactions
	err := c.get(ctx, "/universe/factions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetGraphics Get graphics.
// GetGraphics 获取图形资源.
//
// Route: GET /universe/graphics/
// 路由: GET /universe/graphics/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetGraphics(ctx context.Context, params *models.GetGraphicsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/graphics/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetGraphicsGraphicId Get graphic information.
// GetGraphicsGraphicId 获取图形资源信息.
//
// Route: GET /universe/graphics/{graphic_id}/
// 路由: GET /universe/graphics/{graphic_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetGraphicsGraphicId(ctx context.Context, graphicId int32, params *models.GetGraphicsGraphicIdParams) (*models.GetUniverseGraphicsGraphicId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"graphic_id": strconv.FormatInt(int64(graphicId), 10)}
	var result *models.GetUniverseGraphicsGraphicId
	err := c.get(ctx, "/universe/graphics/{graphic_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetGroups Get item groups.
// GetGroups 获取物品分组.
//
// Route: GET /universe/groups/
// 路由: GET /universe/groups/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetGroups(ctx context.Context, params *models.GetGroupsParamsX) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/groups/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetGroupsGroupId Get item group information.
// GetGroupsGroupId 获取物品分组信息.
//
// Route: GET /universe/groups/{group_id}/
// 路由: GET /universe/groups/{group_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetGroupsGroupId(ctx context.Context, groupId int32, params *models.GetGroupsGroupIdParamsX) (*models.GetUniverseGroupsGroupId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"group_id": strconv.FormatInt(int64(groupId), 10)}
	var result *models.GetUniverseGroupsGroupId
	err := c.get(ctx, "/universe/groups/{group_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMoonsMoonId Get moon information.
// GetMoonsMoonId 获取卫星信息.
//
// Route: GET /universe/moons/{moon_id}/
// 路由: GET /universe/moons/{moon_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMoonsMoonId(ctx context.Context, moonId int32, params *models.GetMoonsMoonIdParams) (*models.GetUniverseMoonsMoonId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"moon_id": strconv.FormatInt(int64(moonId), 10)}
	var result *models.GetUniverseMoonsMoonId
	err := c.get(ctx, "/universe/moons/{moon_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPlanetsPlanetId Get planet information.
// GetPlanetsPlanetId 获取行星信息.
//
// Route: GET /universe/planets/{planet_id}/
// 路由: GET /universe/planets/{planet_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPlanetsPlanetId(ctx context.Context, planetId int32, params *models.GetPlanetsPlanetIdParams) (*models.GetUniversePlanetsPlanetId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"planet_id": strconv.FormatInt(int64(planetId), 10)}
	var result *models.GetUniversePlanetsPlanetId
	err := c.get(ctx, "/universe/planets/{planet_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRaces Get character races.
// GetRaces 获取角色种族.
//
// Route: GET /universe/races/
// 路由: GET /universe/races/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetRaces(ctx context.Context, params *models.GetRacesParams) ([]models.GetUniverseRaces, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetUniverseRaces
	err := c.get(ctx, "/universe/races/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRegions Get regions.
// GetRegions 获取星域列表.
//
// Route: GET /universe/regions/
// 路由: GET /universe/regions/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetRegions(ctx context.Context, params *models.GetRegionsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/regions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRegionsRegionId Get region information.
// GetRegionsRegionId 获取星域信息.
//
// Route: GET /universe/regions/{region_id}/
// 路由: GET /universe/regions/{region_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetRegionsRegionId(ctx context.Context, regionId int32, params *models.GetRegionsRegionIdParams) (*models.GetUniverseRegionsRegionId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionId), 10)}
	var result *models.GetUniverseRegionsRegionId
	err := c.get(ctx, "/universe/regions/{region_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStargatesStargateId Get stargate information.
// GetStargatesStargateId 获取星门信息.
//
// Route: GET /universe/stargates/{stargate_id}/
// 路由: GET /universe/stargates/{stargate_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetStargatesStargateId(ctx context.Context, stargateId int32, params *models.GetStargatesStargateIdParams) (*models.GetUniverseStargatesStargateId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"stargate_id": strconv.FormatInt(int64(stargateId), 10)}
	var result *models.GetUniverseStargatesStargateId
	err := c.get(ctx, "/universe/stargates/{stargate_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStarsStarId Get star information.
// GetStarsStarId 获取恒星信息.
//
// Route: GET /universe/stars/{star_id}/
// 路由: GET /universe/stars/{star_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetStarsStarId(ctx context.Context, starId int32, params *models.GetStarsStarIdParams) (*models.GetUniverseStarsStarId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"star_id": strconv.FormatInt(int64(starId), 10)}
	var result *models.GetUniverseStarsStarId
	err := c.get(ctx, "/universe/stars/{star_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStationsStationId Get station information.
// GetStationsStationId 获取空间站信息.
//
// Route: GET /universe/stations/{station_id}/
// 路由: GET /universe/stations/{station_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetStationsStationId(ctx context.Context, stationId int32, params *models.GetStationsStationIdParams) (*models.GetUniverseStationsStationId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"station_id": strconv.FormatInt(int64(stationId), 10)}
	var result *models.GetUniverseStationsStationId
	err := c.get(ctx, "/universe/stations/{station_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStructures List all public structures.
// GetStructures 列出所有公开建筑.
//
// Route: GET /universe/structures/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/structures/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetStructures(ctx context.Context, params *models.GetStructuresParamsX) ([]int64, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int64
	err := c.get(ctx, "/universe/structures/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStructuresStructureId Get structure information.
// GetStructuresStructureId 获取建筑信息.
//
// Route: GET /universe/structures/{structure_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/structures/{structure_id}/ — 该路由缓存长达 3600 秒
// Scopes: esi-universe.read_structures.v1
// 权限: esi-universe.read_structures.v1
func (c *Client) GetStructuresStructureId(ctx context.Context, structureId int64, params *models.GetStructuresStructureIdParams) (*models.GetUniverseStructuresStructureId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"structure_id": strconv.FormatInt(int64(structureId), 10)}
	var result *models.GetUniverseStructuresStructureId
	err := c.get(ctx, "/universe/structures/{structure_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSystemJumps Get system jumps.
// GetSystemJumps 获取星系跃迁数.
//
// Route: GET /universe/system_jumps/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/system_jumps/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSystemJumps(ctx context.Context, params *models.GetSystemJumpsParams) ([]models.GetUniverseSystemJumps, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetUniverseSystemJumps
	err := c.get(ctx, "/universe/system_jumps/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSystemKills Get system kills.
// GetSystemKills 获取星系击杀数.
//
// Route: GET /universe/system_kills/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/system_kills/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSystemKills(ctx context.Context, params *models.GetSystemKillsParams) ([]models.GetUniverseSystemKills, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetUniverseSystemKills
	err := c.get(ctx, "/universe/system_kills/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSystems Get solar systems.
// GetSystems 获取星系列表.
//
// Route: GET /universe/systems/
// 路由: GET /universe/systems/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSystems(ctx context.Context, params *models.GetSystemsParamsX) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/systems/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSystemsSystemId Get solar system information.
// GetSystemsSystemId 获取星系信息.
//
// Route: GET /universe/systems/{system_id}/
// 路由: GET /universe/systems/{system_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetSystemsSystemId(ctx context.Context, systemId int32, params *models.GetSystemsSystemIdParams) (*models.GetUniverseSystemsSystemId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"system_id": strconv.FormatInt(int64(systemId), 10)}
	var result *models.GetUniverseSystemsSystemId
	err := c.get(ctx, "/universe/systems/{system_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTypes Get types.
// GetTypes 获取类型列表.
//
// Route: GET /universe/types/
// 路由: GET /universe/types/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetTypes(ctx context.Context, params *models.GetTypesParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/types/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTypesTypeId Get type information.
// GetTypesTypeId 获取类型信息.
//
// Route: GET /universe/types/{type_id}/
// 路由: GET /universe/types/{type_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetTypesTypeId(ctx context.Context, typeId int32, params *models.GetTypesTypeIdParams) (*models.GetUniverseTypesTypeId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"type_id": strconv.FormatInt(int64(typeId), 10)}
	var result *models.GetUniverseTypesTypeId
	err := c.get(ctx, "/universe/types/{type_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostIds Bulk names to IDs.
// PostIds 批量将名称转换为 ID.
//
// Route: POST /universe/ids/
// 路由: POST /universe/ids/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) PostIds(ctx context.Context, body []string, params *models.PostIdsParams) (*models.PostUniverseIds, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	if body == nil {
		return nil, errBodyRequired
	}
	var result *models.PostUniverseIds
	err := c.post(ctx, "/universe/ids/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostNames Get names and categories for a set of IDs.
// PostNames 按一组 ID 获取名称与类别.
//
// Route: POST /universe/names/
// 路由: POST /universe/names/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) PostNames(ctx context.Context, body []int32, params *models.PostNamesParams) ([]models.PostUniverseNames, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.PostUniverseNames
	err := c.post(ctx, "/universe/names/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
