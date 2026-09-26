package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetUniverseAncestries Get ancestries.
// GetUniverseAncestries 获取血统.
//
// Route: GET /universe/ancestries/
// 路由: GET /universe/ancestries/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseAncestries(ctx context.Context, ifNoneMatch string) ([]models.UniverseAncestry, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.UniverseAncestry
	err := c.get(ctx, "/universe/ancestries/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseAsteroidBelt Get asteroid belt information.
// GetUniverseAsteroidBelt 获取小行星带信息.
//
// Route: GET /universe/asteroid_belts/{asteroid_belt_id}/
// 路由: GET /universe/asteroid_belts/{asteroid_belt_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseAsteroidBelt(ctx context.Context, asteroidBeltID int32, ifNoneMatch string) (*models.UniverseAsteroidBelt, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"asteroid_belt_id": strconv.FormatInt(int64(asteroidBeltID), 10)}
	var result *models.UniverseAsteroidBelt
	err := c.get(ctx, "/universe/asteroid_belts/{asteroid_belt_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseBloodlines Get bloodlines.
// GetUniverseBloodlines 获取种族血脉.
//
// Route: GET /universe/bloodlines/
// 路由: GET /universe/bloodlines/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseBloodlines(ctx context.Context, ifNoneMatch string) ([]models.UniverseBloodline, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.UniverseBloodline
	err := c.get(ctx, "/universe/bloodlines/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseCategories Get item categories.
// GetUniverseCategories 获取物品类别.
//
// Route: GET /universe/categories/
// 路由: GET /universe/categories/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseCategories(ctx context.Context, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/categories/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseCategory Get item category information.
// GetUniverseCategory 获取物品类别信息.
//
// Route: GET /universe/categories/{category_id}/
// 路由: GET /universe/categories/{category_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseCategory(ctx context.Context, categoryID int32, ifNoneMatch string) (*models.UniverseCategory, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"category_id": strconv.FormatInt(int64(categoryID), 10)}
	var result *models.UniverseCategory
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
func (c *Client) GetConstellations(ctx context.Context, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/constellations/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetConstellationInformation Get constellation information.
// GetConstellationInformation 获取星座信息.
//
// Route: GET /universe/constellations/{constellation_id}/
// 路由: GET /universe/constellations/{constellation_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetConstellationInformation(ctx context.Context, constellationID int32, ifNoneMatch string) (*models.UniverseConstellation, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"constellation_id": strconv.FormatInt(int64(constellationID), 10)}
	var result *models.UniverseConstellation
	err := c.get(ctx, "/universe/constellations/{constellation_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseFactions Get factions.
// GetUniverseFactions 获取势力.
//
// Route: GET /universe/factions/
// 路由: GET /universe/factions/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseFactions(ctx context.Context, ifNoneMatch string) ([]models.UniverseFaction, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.UniverseFaction
	err := c.get(ctx, "/universe/factions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseGraphics Get graphics.
// GetUniverseGraphics 获取图形资源.
//
// Route: GET /universe/graphics/
// 路由: GET /universe/graphics/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseGraphics(ctx context.Context, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/graphics/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseGraphic Get graphic information.
// GetUniverseGraphic 获取图形资源信息.
//
// Route: GET /universe/graphics/{graphic_id}/
// 路由: GET /universe/graphics/{graphic_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseGraphic(ctx context.Context, graphicID int32, ifNoneMatch string) (*models.UniverseGraphic, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"graphic_id": strconv.FormatInt(int64(graphicID), 10)}
	var result *models.UniverseGraphic
	err := c.get(ctx, "/universe/graphics/{graphic_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseGroups Get item groups.
// GetUniverseGroups 获取物品分组.
//
// Route: GET /universe/groups/
// 路由: GET /universe/groups/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseGroups(ctx context.Context, page int32, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/groups/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseGroup Get item group information.
// GetUniverseGroup 获取物品分组信息.
//
// Route: GET /universe/groups/{group_id}/
// 路由: GET /universe/groups/{group_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseGroup(ctx context.Context, groupID int32, ifNoneMatch string) (*models.UniverseGroup, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"group_id": strconv.FormatInt(int64(groupID), 10)}
	var result *models.UniverseGroup
	err := c.get(ctx, "/universe/groups/{group_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseMoon Get moon information.
// GetUniverseMoon 获取卫星信息.
//
// Route: GET /universe/moons/{moon_id}/
// 路由: GET /universe/moons/{moon_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseMoon(ctx context.Context, moonID int32, ifNoneMatch string) (*models.UniverseMoon, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"moon_id": strconv.FormatInt(int64(moonID), 10)}
	var result *models.UniverseMoon
	err := c.get(ctx, "/universe/moons/{moon_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniversePlanet Get planet information.
// GetUniversePlanet 获取行星信息.
//
// Route: GET /universe/planets/{planet_id}/
// 路由: GET /universe/planets/{planet_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniversePlanet(ctx context.Context, planetID int32, ifNoneMatch string) (*models.UniversePlanet, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"planet_id": strconv.FormatInt(int64(planetID), 10)}
	var result *models.UniversePlanet
	err := c.get(ctx, "/universe/planets/{planet_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseRaces Get character races.
// GetUniverseRaces 获取角色种族.
//
// Route: GET /universe/races/
// 路由: GET /universe/races/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseRaces(ctx context.Context, ifNoneMatch string) ([]models.UniverseRace, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.UniverseRace
	err := c.get(ctx, "/universe/races/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseRegions Get regions.
// GetUniverseRegions 获取星域列表.
//
// Route: GET /universe/regions/
// 路由: GET /universe/regions/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseRegions(ctx context.Context, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/regions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseRegion Get region information.
// GetUniverseRegion 获取星域信息.
//
// Route: GET /universe/regions/{region_id}/
// 路由: GET /universe/regions/{region_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseRegion(ctx context.Context, regionID int32, ifNoneMatch string) (*models.UniverseRegion, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionID), 10)}
	var result *models.UniverseRegion
	err := c.get(ctx, "/universe/regions/{region_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseStargate Get stargate information.
// GetUniverseStargate 获取星门信息.
//
// Route: GET /universe/stargates/{stargate_id}/
// 路由: GET /universe/stargates/{stargate_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseStargate(ctx context.Context, stargateID int32, ifNoneMatch string) (*models.UniverseStargate, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"stargate_id": strconv.FormatInt(int64(stargateID), 10)}
	var result *models.UniverseStargate
	err := c.get(ctx, "/universe/stargates/{stargate_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseStar Get star information.
// GetUniverseStar 获取恒星信息.
//
// Route: GET /universe/stars/{star_id}/
// 路由: GET /universe/stars/{star_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseStar(ctx context.Context, starID int32, ifNoneMatch string) (*models.UniverseStar, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"star_id": strconv.FormatInt(int64(starID), 10)}
	var result *models.UniverseStar
	err := c.get(ctx, "/universe/stars/{star_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseStation Get station information.
// GetUniverseStation 获取空间站信息.
//
// Route: GET /universe/stations/{station_id}/
// 路由: GET /universe/stations/{station_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseStation(ctx context.Context, stationID int32, ifNoneMatch string) (*models.UniverseStation, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"station_id": strconv.FormatInt(int64(stationID), 10)}
	var result *models.UniverseStation
	err := c.get(ctx, "/universe/stations/{station_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicStructures List all public structures.
// GetPublicStructures 列出所有公开建筑.
//
// Route: GET /universe/structures/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/structures/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPublicStructures(ctx context.Context, filter string, ifNoneMatch string) ([]int64, error) {
	query := url.Values{}
	headers := map[string]string{}
	if filter != "" {
		query.Set("filter", filter)
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int64
	err := c.get(ctx, "/universe/structures/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseStructure Get structure information.
// GetUniverseStructure 获取建筑信息.
//
// Route: GET /universe/structures/{structure_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/structures/{structure_id}/ — 该路由缓存长达 3600 秒
// Scopes: esi-universe.read_structures.v1
// 权限: esi-universe.read_structures.v1
func (c *Client) GetUniverseStructure(ctx context.Context, structureID int64, token string, ifNoneMatch string) (*models.UniverseStructure, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"structure_id": strconv.FormatInt(int64(structureID), 10)}
	var result *models.UniverseStructure
	err := c.get(ctx, "/universe/structures/{structure_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseSystemJumps Get system jumps.
// GetUniverseSystemJumps 获取星系跃迁数.
//
// Route: GET /universe/system_jumps/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/system_jumps/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseSystemJumps(ctx context.Context, ifNoneMatch string) ([]models.UniverseSystemJump, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.UniverseSystemJump
	err := c.get(ctx, "/universe/system_jumps/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseSystemKills Get system kills.
// GetUniverseSystemKills 获取星系击杀数.
//
// Route: GET /universe/system_kills/ — This route is cached for up to 3600 seconds
// 路由: GET /universe/system_kills/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseSystemKills(ctx context.Context, ifNoneMatch string) ([]models.UniverseSystemKills, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.UniverseSystemKills
	err := c.get(ctx, "/universe/system_kills/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseSystems Get solar systems.
// GetUniverseSystems 获取星系列表.
//
// Route: GET /universe/systems/
// 路由: GET /universe/systems/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseSystems(ctx context.Context, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/systems/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseSystem Get solar system information.
// GetUniverseSystem 获取星系信息.
//
// Route: GET /universe/systems/{system_id}/
// 路由: GET /universe/systems/{system_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseSystem(ctx context.Context, systemID int32, ifNoneMatch string) (*models.UniverseSolarSystem, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"system_id": strconv.FormatInt(int64(systemID), 10)}
	var result *models.UniverseSolarSystem
	err := c.get(ctx, "/universe/systems/{system_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseTypes Get types.
// GetUniverseTypes 获取类型列表.
//
// Route: GET /universe/types/
// 路由: GET /universe/types/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseTypes(ctx context.Context, page int32, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/universe/types/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUniverseType Get type information.
// GetUniverseType 获取类型信息.
//
// Route: GET /universe/types/{type_id}/
// 路由: GET /universe/types/{type_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetUniverseType(ctx context.Context, typeID int32, ifNoneMatch string) (*models.UniverseType, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"type_id": strconv.FormatInt(int64(typeID), 10)}
	var result *models.UniverseType
	err := c.get(ctx, "/universe/types/{type_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResolveNamesToIDs Bulk names to IDs.
// ResolveNamesToIDs 批量将名称转换为 ID.
//
// Route: POST /universe/ids/
// 路由: POST /universe/ids/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) ResolveNamesToIDs(ctx context.Context, body []string) (*models.ResolvedIds, error) {
	query := url.Values{}
	headers := map[string]string{}
	var pathParams map[string]string
	if body == nil {
		return nil, errBodyRequired
	}
	var result *models.ResolvedIds
	err := c.post(ctx, "/universe/ids/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResolveIDsToNames Get names and categories for a set of IDs.
// ResolveIDsToNames 按一组 ID 获取名称与类别.
//
// Route: POST /universe/names/
// 路由: POST /universe/names/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) ResolveIDsToNames(ctx context.Context, body []int32) ([]models.UniverseName, error) {
	query := url.Values{}
	headers := map[string]string{}
	var pathParams map[string]string
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.UniverseName
	err := c.post(ctx, "/universe/names/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
