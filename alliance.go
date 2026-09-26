package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetAllianceId Get alliance information.
// GetAllianceId 获取联盟信息.
//
// Route: GET /alliances/{alliance_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /alliances/{alliance_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAllianceId(ctx context.Context, allianceId int32, params *models.GetAllianceIdParams) (*models.GetAlliancesAllianceId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceId), 10)}
	var result *models.GetAlliancesAllianceId
	err := c.get(ctx, "/alliances/{alliance_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllianceIdCorporations List alliance's corporations.
// GetAllianceIdCorporations 列出联盟的军团.
//
// Route: GET /alliances/{alliance_id}/corporations/ — This route is cached for up to 3600 seconds
// 路由: GET /alliances/{alliance_id}/corporations/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAllianceIdCorporations(ctx context.Context, allianceId int32, params *models.GetAllianceIdCorporationsParams) ([]int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceId), 10)}
	var result []int32
	err := c.get(ctx, "/alliances/{alliance_id}/corporations/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllianceIdIcons Get alliance icon.
// GetAllianceIdIcons 获取联盟图标.
//
// Route: GET /alliances/{alliance_id}/icons/
// 路由: GET /alliances/{alliance_id}/icons/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAllianceIdIcons(ctx context.Context, allianceId int32, params *models.GetAllianceIdIconsParams) (*models.GetAlliancesAllianceIdIcons, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceId), 10)}
	var result *models.GetAlliancesAllianceIdIcons
	err := c.get(ctx, "/alliances/{alliance_id}/icons/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAlliances List all alliances.
// GetAlliances 列出所有联盟.
//
// Route: GET /alliances/ — This route is cached for up to 3600 seconds
// 路由: GET /alliances/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAlliances(ctx context.Context, params *models.GetAlliancesParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/alliances/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
