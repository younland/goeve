package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCorporationInformation Get corporation information.
// GetCorporationInformation 获取军团信息.
//
// Route: GET /corporations/{corporation_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCorporationInformation(ctx context.Context, corporationID int32, opts ...RequestOption) (*models.Corporation, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result *models.Corporation
	err := c.get(ctx, "/corporations/{corporation_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationAllianceHistory Get alliance history.
// GetCorporationAllianceHistory 获取联盟历史.
//
// Route: GET /corporations/{corporation_id}/alliancehistory/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/alliancehistory/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCorporationAllianceHistory(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.AllianceHistoryEntry, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.AllianceHistoryEntry
	err := c.get(ctx, "/corporations/{corporation_id}/alliancehistory/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationBlueprints Get corporation blueprints.
// GetCorporationBlueprints 获取军团蓝图.
//
// Route: GET /corporations/{corporation_id}/blueprints/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/blueprints/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_blueprints.v1
// 权限: esi-corporations.read_blueprints.v1
func (c *Client) GetCorporationBlueprints(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Blueprint, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.Blueprint
	err := c.get(ctx, "/corporations/{corporation_id}/blueprints/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationContainerLogs Get all corporation ALSC logs.
// GetCorporationContainerLogs 获取军团全部 ALSC 日志.
//
// Route: GET /corporations/{corporation_id}/containers/logs/ — This route is cached for up to 600 seconds
// 路由: GET /corporations/{corporation_id}/containers/logs/ — 该路由缓存长达 600 秒
// Scopes: esi-corporations.read_container_logs.v1
// 权限: esi-corporations.read_container_logs.v1
func (c *Client) GetCorporationContainerLogs(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.ContainerLog, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.ContainerLog
	err := c.get(ctx, "/corporations/{corporation_id}/containers/logs/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationDivisions Get corporation divisions.
// GetCorporationDivisions 获取军团部门.
//
// Route: GET /corporations/{corporation_id}/divisions/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/divisions/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_divisions.v1
// 权限: esi-corporations.read_divisions.v1
func (c *Client) GetCorporationDivisions(ctx context.Context, corporationID int32, opts ...RequestOption) (*models.CorporationDivisions, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result *models.CorporationDivisions
	err := c.get(ctx, "/corporations/{corporation_id}/divisions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationFacilities Get corporation facilities.
// GetCorporationFacilities 获取军团设施.
//
// Route: GET /corporations/{corporation_id}/facilities/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/facilities/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_facilities.v1
// 权限: esi-corporations.read_facilities.v1
func (c *Client) GetCorporationFacilities(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationFacility, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationFacility
	err := c.get(ctx, "/corporations/{corporation_id}/facilities/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIcon Get corporation icon.
// GetCorporationIcon 获取军团图标.
//
// Route: GET /corporations/{corporation_id}/icons/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/icons/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCorporationIcon(ctx context.Context, corporationID int32, opts ...RequestOption) (*models.CorporationIcons, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result *models.CorporationIcons
	err := c.get(ctx, "/corporations/{corporation_id}/icons/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMedals Get corporation medals.
// GetCorporationMedals 获取军团勋章.
//
// Route: GET /corporations/{corporation_id}/medals/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/medals/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_medals.v1
// 权限: esi-corporations.read_medals.v1
func (c *Client) GetCorporationMedals(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationMedal, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationMedal
	err := c.get(ctx, "/corporations/{corporation_id}/medals/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIssuedMedals Get corporation issued medals.
// GetCorporationIssuedMedals 获取军团颁发的勋章.
//
// Route: GET /corporations/{corporation_id}/medals/issued/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/medals/issued/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_medals.v1
// 权限: esi-corporations.read_medals.v1
func (c *Client) GetCorporationIssuedMedals(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.IssuedMedal, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.IssuedMedal
	err := c.get(ctx, "/corporations/{corporation_id}/medals/issued/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMembers Get corporation members.
// GetCorporationMembers 获取军团成员.
//
// Route: GET /corporations/{corporation_id}/members/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/members/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_corporation_membership.v1
// 权限: esi-corporations.read_corporation_membership.v1
func (c *Client) GetCorporationMembers(ctx context.Context, corporationID int32, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []int32
	err := c.get(ctx, "/corporations/{corporation_id}/members/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMemberLimit Get corporation member limit.
// GetCorporationMemberLimit 获取军团成员上限.
//
// Route: GET /corporations/{corporation_id}/members/limit/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/members/limit/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.track_members.v1
// 权限: esi-corporations.track_members.v1
func (c *Client) GetCorporationMemberLimit(ctx context.Context, corporationID int32, opts ...RequestOption) (int32, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result int32
	err := c.get(ctx, "/corporations/{corporation_id}/members/limit/", pathParams, query, headers, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GetCorporationMemberTitles Get corporation's members' titles.
// GetCorporationMemberTitles 获取军团成员的头衔.
//
// Route: GET /corporations/{corporation_id}/members/titles/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/members/titles/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_titles.v1
// 权限: esi-corporations.read_titles.v1
func (c *Client) GetCorporationMemberTitles(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.MemberTitles, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.MemberTitles
	err := c.get(ctx, "/corporations/{corporation_id}/members/titles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMemberTracking Track corporation members.
// GetCorporationMemberTracking 追踪军团成员.
//
// Route: GET /corporations/{corporation_id}/membertracking/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/membertracking/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.track_members.v1
// 权限: esi-corporations.track_members.v1
func (c *Client) GetCorporationMemberTracking(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.MemberTrackingEntry, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.MemberTrackingEntry
	err := c.get(ctx, "/corporations/{corporation_id}/membertracking/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMemberRoles Get corporation member roles.
// GetCorporationMemberRoles 获取军团成员职务.
//
// Route: GET /corporations/{corporation_id}/roles/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/roles/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_corporation_membership.v1
// 权限: esi-corporations.read_corporation_membership.v1
func (c *Client) GetCorporationMemberRoles(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationMemberRoles, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationMemberRoles
	err := c.get(ctx, "/corporations/{corporation_id}/roles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMemberRolesHistory Get corporation member roles history.
// GetCorporationMemberRolesHistory 获取军团成员职务历史.
//
// Route: GET /corporations/{corporation_id}/roles/history/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/roles/history/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_corporation_membership.v1
// 权限: esi-corporations.read_corporation_membership.v1
func (c *Client) GetCorporationMemberRolesHistory(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationRoleHistory, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationRoleHistory
	err := c.get(ctx, "/corporations/{corporation_id}/roles/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationShareholders Get corporation shareholders.
// GetCorporationShareholders 获取军团股东.
//
// Route: GET /corporations/{corporation_id}/shareholders/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/shareholders/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationShareholders(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Shareholder, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.Shareholder
	err := c.get(ctx, "/corporations/{corporation_id}/shareholders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationStandings Get corporation standings.
// GetCorporationStandings 获取军团声望.
//
// Route: GET /corporations/{corporation_id}/standings/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/standings/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_standings.v1
// 权限: esi-corporations.read_standings.v1
func (c *Client) GetCorporationStandings(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Standing, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.Standing
	err := c.get(ctx, "/corporations/{corporation_id}/standings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationStarbases Get corporation starbases (POSes)
// GetCorporationStarbases 获取军团母星基地（POS）
//
// Route: GET /corporations/{corporation_id}/starbases/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/starbases/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_starbases.v1
// 权限: esi-corporations.read_starbases.v1
func (c *Client) GetCorporationStarbases(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Starbase, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.Starbase
	err := c.get(ctx, "/corporations/{corporation_id}/starbases/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationStarbase Get starbase (POS) detail.
// GetCorporationStarbase 获取母星基地（POS）详情.
//
// Route: GET /corporations/{corporation_id}/starbases/{starbase_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/starbases/{starbase_id}/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_starbases.v1
// 权限: esi-corporations.read_starbases.v1
func (c *Client) GetCorporationStarbase(ctx context.Context, corporationID int32, starbaseID int64, systemID string, opts ...RequestOption) (*models.StarbaseDetail, error) {
	query, headers := newRequestOptions(opts...)
	query.Set("system_id", systemID)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10), "starbase_id": strconv.FormatInt(int64(starbaseID), 10)}
	var result *models.StarbaseDetail
	err := c.get(ctx, "/corporations/{corporation_id}/starbases/{starbase_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationStructures Get corporation structures.
// GetCorporationStructures 获取军团建筑.
//
// Route: GET /corporations/{corporation_id}/structures/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/structures/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_structures.v1
// 权限: esi-corporations.read_structures.v1
func (c *Client) GetCorporationStructures(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationStructure, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationStructure
	err := c.get(ctx, "/corporations/{corporation_id}/structures/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationTitles Get corporation titles.
// GetCorporationTitles 获取军团头衔.
//
// Route: GET /corporations/{corporation_id}/titles/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/titles/ — 该路由缓存长达 3600 秒
// Scopes: esi-corporations.read_titles.v1
// 权限: esi-corporations.read_titles.v1
func (c *Client) GetCorporationTitles(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationTitle, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationTitle
	err := c.get(ctx, "/corporations/{corporation_id}/titles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNpcCorporations Get npc corporations.
// GetNpcCorporations 获取 NPC 军团.
//
// Route: GET /corporations/npccorps/
// 路由: GET /corporations/npccorps/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetNpcCorporations(ctx context.Context, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/corporations/npccorps/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
