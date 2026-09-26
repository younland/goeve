package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetAlliance Get alliance information.
// GetAlliance 获取联盟信息.
//
// Route: GET /alliances/{alliance_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /alliances/{alliance_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAlliance(ctx context.Context, allianceID int32, opts ...RequestOption) (*models.Alliance, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceID), 10)}
	var result *models.Alliance
	err := c.get(ctx, "/alliances/{alliance_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllianceCorporations List alliance's corporations.
// GetAllianceCorporations 列出联盟的军团.
//
// Route: GET /alliances/{alliance_id}/corporations/ — This route is cached for up to 3600 seconds
// 路由: GET /alliances/{alliance_id}/corporations/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAllianceCorporations(ctx context.Context, allianceID int32, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceID), 10)}
	var result []int32
	err := c.get(ctx, "/alliances/{alliance_id}/corporations/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllianceIcons Get alliance icon.
// GetAllianceIcons 获取联盟图标.
//
// Route: GET /alliances/{alliance_id}/icons/
// 路由: GET /alliances/{alliance_id}/icons/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAllianceIcons(ctx context.Context, allianceID int32, opts ...RequestOption) (*models.AllianceIcons, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceID), 10)}
	var result *models.AllianceIcons
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
func (c *Client) GetAlliances(ctx context.Context, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/alliances/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
