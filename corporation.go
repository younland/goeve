package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCorporationId Get corporation information.
// GetCorporationId 获取军团信息.
//
// Route: GET /corporations/{corporation_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCorporationId(ctx context.Context, corporationId int32, params *models.GetCorporationIdParams) (*models.GetCorporationsCorporationIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result *models.GetCorporationsCorporationIdOk
	err := c.get(ctx, "/corporations/{corporation_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdAlliancehistory Get alliance history.
// GetCorporationIdAlliancehistory 获取联盟历史.
//
// Route: GET /corporations/{corporation_id}/alliancehistory/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/alliancehistory/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCorporationIdAlliancehistory(ctx context.Context, corporationId int32, params *models.GetCorporationIdAlliancehistoryParams) ([]models.GetCorporationsCorporationIdAlliancehistory, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdAlliancehistory
	err := c.get(ctx, "/corporations/{corporation_id}/alliancehistory/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdBlueprints Get corporation blueprints.
// GetCorporationIdBlueprints 获取军团蓝图.
//
// Route: GET /corporations/{corporation_id}/blueprints/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/blueprints/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_blueprints.v1
// 权限: esi-corporations.read_blueprints.v1
func (c *Client) GetCorporationIdBlueprints(ctx context.Context, corporationId int32, params *models.GetCorporationIdBlueprintsParams) ([]models.GetCorporationsCorporationIdBlueprints, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdBlueprints
	err := c.get(ctx, "/corporations/{corporation_id}/blueprints/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdContainersLogs Get all corporation ALSC logs.
// GetCorporationIdContainersLogs 获取军团全部 ALSC 日志.
//
// Route: GET /corporations/{corporation_id}/containers/logs/ — This route is cached for up to 600 seconds
// 路由: GET /corporations/{corporation_id}/containers/logs/ — 该路由缓存长达 600 秒
// Scopes: esi-corporations.read_container_logs.v1
// 权限: esi-corporations.read_container_logs.v1
func (c *Client) GetCorporationIdContainersLogs(ctx context.Context, corporationId int32, params *models.GetCorporationIdContainersLogsParams) ([]models.GetCorporationsCorporationIdContainersLogs, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdContainersLogs
	err := c.get(ctx, "/corporations/{corporation_id}/containers/logs/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdDivisions Get corporation divisions.
// GetCorporationIdDivisions 获取军团部门.
//
// Route: GET /corporations/{corporation_id}/divisions/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/divisions/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_divisions.v1
// 权限: esi-corporations.read_divisions.v1
func (c *Client) GetCorporationIdDivisions(ctx context.Context, corporationId int32, params *models.GetCorporationIdDivisionsParams) (*models.GetCorporationsCorporationIdDivisionsOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result *models.GetCorporationsCorporationIdDivisionsOk
	err := c.get(ctx, "/corporations/{corporation_id}/divisions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdFacilities Get corporation facilities.
// GetCorporationIdFacilities 获取军团设施.
//
// Route: GET /corporations/{corporation_id}/facilities/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/facilities/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_facilities.v1
// 权限: esi-corporations.read_facilities.v1
func (c *Client) GetCorporationIdFacilities(ctx context.Context, corporationId int32, params *models.GetCorporationIdFacilitiesParams) ([]models.GetCorporationsCorporationIdFacilities, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdFacilities
	err := c.get(ctx, "/corporations/{corporation_id}/facilities/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdIcons Get corporation icon.
// GetCorporationIdIcons 获取军团图标.
//
// Route: GET /corporations/{corporation_id}/icons/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/icons/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCorporationIdIcons(ctx context.Context, corporationId int32, params *models.GetCorporationIdIconsParams) (*models.GetCorporationsCorporationIdIconsOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result *models.GetCorporationsCorporationIdIconsOk
	err := c.get(ctx, "/corporations/{corporation_id}/icons/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdMedals Get corporation medals.
// GetCorporationIdMedals 获取军团勋章.
//
// Route: GET /corporations/{corporation_id}/medals/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/medals/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_medals.v1
// 权限: esi-corporations.read_medals.v1
func (c *Client) GetCorporationIdMedals(ctx context.Context, corporationId int32, params *models.GetCorporationIdMedalsParams) ([]models.GetCorporationsCorporationIdMedals, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdMedals
	err := c.get(ctx, "/corporations/{corporation_id}/medals/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdMedalsIssued Get corporation issued medals.
// GetCorporationIdMedalsIssued 获取军团颁发的勋章.
//
// Route: GET /corporations/{corporation_id}/medals/issued/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/medals/issued/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_medals.v1
// 权限: esi-corporations.read_medals.v1
func (c *Client) GetCorporationIdMedalsIssued(ctx context.Context, corporationId int32, params *models.GetCorporationIdMedalsIssuedParams) ([]models.GetCorporationsCorporationIdMedalsIssued, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdMedalsIssued
	err := c.get(ctx, "/corporations/{corporation_id}/medals/issued/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdMembers Get corporation members.
// GetCorporationIdMembers 获取军团成员.
//
// Route: GET /corporations/{corporation_id}/members/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/members/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_corporation_membership.v1
// 权限: esi-corporations.read_corporation_membership.v1
func (c *Client) GetCorporationIdMembers(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembersParams) ([]int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []int32
	err := c.get(ctx, "/corporations/{corporation_id}/members/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdMembersLimit Get corporation member limit.
// GetCorporationIdMembersLimit 获取军团成员上限.
//
// Route: GET /corporations/{corporation_id}/members/limit/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/members/limit/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.track_members.v1
// 权限: esi-corporations.track_members.v1
func (c *Client) GetCorporationIdMembersLimit(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembersLimitParams) (int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result int32
	err := c.get(ctx, "/corporations/{corporation_id}/members/limit/", pathParams, query, headers, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GetCorporationIdMembersTitles Get corporation's members' titles.
// GetCorporationIdMembersTitles 获取军团成员的头衔.
//
// Route: GET /corporations/{corporation_id}/members/titles/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/members/titles/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_titles.v1
// 权限: esi-corporations.read_titles.v1
func (c *Client) GetCorporationIdMembersTitles(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembersTitlesParams) ([]models.GetCorporationsCorporationIdMembersTitles, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdMembersTitles
	err := c.get(ctx, "/corporations/{corporation_id}/members/titles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdMembertracking Track corporation members.
// GetCorporationIdMembertracking 追踪军团成员.
//
// Route: GET /corporations/{corporation_id}/membertracking/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/membertracking/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.track_members.v1
// 权限: esi-corporations.track_members.v1
func (c *Client) GetCorporationIdMembertracking(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembertrackingParams) ([]models.GetCorporationsCorporationIdMembertracking, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdMembertracking
	err := c.get(ctx, "/corporations/{corporation_id}/membertracking/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdRoles Get corporation member roles.
// GetCorporationIdRoles 获取军团成员职务.
//
// Route: GET /corporations/{corporation_id}/roles/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/roles/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_corporation_membership.v1
// 权限: esi-corporations.read_corporation_membership.v1
func (c *Client) GetCorporationIdRoles(ctx context.Context, corporationId int32, params *models.GetCorporationIdRolesParams) ([]models.GetCorporationsCorporationIdRoles, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdRoles
	err := c.get(ctx, "/corporations/{corporation_id}/roles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdRolesHistory Get corporation member roles history.
// GetCorporationIdRolesHistory 获取军团成员职务历史.
//
// Route: GET /corporations/{corporation_id}/roles/history/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/roles/history/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_corporation_membership.v1
// 权限: esi-corporations.read_corporation_membership.v1
func (c *Client) GetCorporationIdRolesHistory(ctx context.Context, corporationId int32, params *models.GetCorporationIdRolesHistoryParams) ([]models.GetCorporationsCorporationIdRolesHistory, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdRolesHistory
	err := c.get(ctx, "/corporations/{corporation_id}/roles/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdShareholders Get corporation shareholders.
// GetCorporationIdShareholders 获取军团股东.
//
// Route: GET /corporations/{corporation_id}/shareholders/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/shareholders/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationIdShareholders(ctx context.Context, corporationId int32, params *models.GetCorporationIdShareholdersParams) ([]models.GetCorporationsCorporationIdShareholders, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdShareholders
	err := c.get(ctx, "/corporations/{corporation_id}/shareholders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdStandings Get corporation standings.
// GetCorporationIdStandings 获取军团声望.
//
// Route: GET /corporations/{corporation_id}/standings/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/standings/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_standings.v1
// 权限: esi-corporations.read_standings.v1
func (c *Client) GetCorporationIdStandings(ctx context.Context, corporationId int32, params *models.GetCorporationIdStandingsParams) ([]models.GetCorporationsCorporationIdStandings, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdStandings
	err := c.get(ctx, "/corporations/{corporation_id}/standings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdStarbases Get corporation starbases (POSes)
// GetCorporationIdStarbases 获取军团母星基地（POS）
//
// Route: GET /corporations/{corporation_id}/starbases/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/starbases/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_starbases.v1
// 权限: esi-corporations.read_starbases.v1
func (c *Client) GetCorporationIdStarbases(ctx context.Context, corporationId int32, params *models.GetCorporationIdStarbasesParams) ([]models.GetCorporationsCorporationIdStarbases, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdStarbases
	err := c.get(ctx, "/corporations/{corporation_id}/starbases/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdStarbasesStarbaseId Get starbase (POS) detail.
// GetCorporationIdStarbasesStarbaseId 获取母星基地（POS）详情.
//
// Route: GET /corporations/{corporation_id}/starbases/{starbase_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/starbases/{starbase_id}/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_starbases.v1
// 权限: esi-corporations.read_starbases.v1
func (c *Client) GetCorporationIdStarbasesStarbaseId(ctx context.Context, corporationId int32, starbaseId int64, params *models.GetCorporationIdStarbasesStarbaseIdParams) (*models.GetCorporationsCorporationIdStarbasesStarbaseIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10), "starbase_id": strconv.FormatInt(int64(starbaseId), 10)}
	var result *models.GetCorporationsCorporationIdStarbasesStarbaseIdOk
	err := c.get(ctx, "/corporations/{corporation_id}/starbases/{starbase_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdStructures Get corporation structures.
// GetCorporationIdStructures 获取军团建筑.
//
// Route: GET /corporations/{corporation_id}/structures/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/structures/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_structures.v1
// 权限: esi-corporations.read_structures.v1
func (c *Client) GetCorporationIdStructures(ctx context.Context, corporationId int32, params *models.GetCorporationIdStructuresParams) ([]models.GetCorporationsCorporationIdStructures, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdStructures
	err := c.get(ctx, "/corporations/{corporation_id}/structures/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIdTitles Get corporation titles.
// GetCorporationIdTitles 获取军团头衔.
//
// Route: GET /corporations/{corporation_id}/titles/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/titles/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_titles.v1
// 权限: esi-corporations.read_titles.v1
func (c *Client) GetCorporationIdTitles(ctx context.Context, corporationId int32, params *models.GetCorporationIdTitlesParams) ([]models.GetCorporationsCorporationIdTitles, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdTitles
	err := c.get(ctx, "/corporations/{corporation_id}/titles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNpccorps Get npc corporations.
// GetNpccorps 获取 NPC 军团.
//
// Route: GET /corporations/npccorps/
// 路由: GET /corporations/npccorps/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetNpccorps(ctx context.Context, params *models.GetNpccorpsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/corporations/npccorps/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
