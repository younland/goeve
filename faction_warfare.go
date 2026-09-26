package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterFactionWarfareStats Overview of a character involved in faction warfare.
// GetCharacterFactionWarfareStats 参与势力战争的角色的概况.
//
// Route: GET /characters/{character_id}/fw/stats/
// 路由: GET /characters/{character_id}/fw/stats/
// Scopes: esi-characters.read_fw_stats.v1
// 权限: esi-characters.read_fw_stats.v1
func (c *Client) GetCharacterFactionWarfareStats(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterFactionWarfareStats, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.CharacterFactionWarfareStats
	err := c.get(ctx, "/characters/{character_id}/fw/stats/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationFactionWarfareStats Overview of a corporation involved in faction warfare.
// GetCorporationFactionWarfareStats 参与势力战争的军团的概况.
//
// Route: GET /corporations/{corporation_id}/fw/stats/
// 路由: GET /corporations/{corporation_id}/fw/stats/
// Scopes: esi-corporations.read_fw_stats.v1
// 权限: esi-corporations.read_fw_stats.v1
func (c *Client) GetCorporationFactionWarfareStats(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) (*models.CorporationFactionWarfareStats, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result *models.CorporationFactionWarfareStats
	err := c.get(ctx, "/corporations/{corporation_id}/fw/stats/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFactionWarfareLeaderboard List of the top factions in faction warfare.
// GetFactionWarfareLeaderboard 势力战争排名前列的势力列表.
//
// Route: GET /fw/leaderboards/
// 路由: GET /fw/leaderboards/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFactionWarfareLeaderboard(ctx context.Context, ifNoneMatch ...string) (*models.FactionWarfareLeaderboard, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result *models.FactionWarfareLeaderboard
	err := c.get(ctx, "/fw/leaderboards/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFactionWarfareCharacterLeaderboard List of the top pilots in faction warfare.
// GetFactionWarfareCharacterLeaderboard 势力战争排名前列的飞行员列表.
//
// Route: GET /fw/leaderboards/characters/
// 路由: GET /fw/leaderboards/characters/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFactionWarfareCharacterLeaderboard(ctx context.Context, ifNoneMatch ...string) (*models.FactionWarfareCharacterLeaderboard, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result *models.FactionWarfareCharacterLeaderboard
	err := c.get(ctx, "/fw/leaderboards/characters/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFactionWarfareCorporationLeaderboard List of the top corporations in faction warfare.
// GetFactionWarfareCorporationLeaderboard 势力战争排名前列的军团列表.
//
// Route: GET /fw/leaderboards/corporations/
// 路由: GET /fw/leaderboards/corporations/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFactionWarfareCorporationLeaderboard(ctx context.Context, ifNoneMatch ...string) (*models.FactionWarfareCorporationLeaderboard, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result *models.FactionWarfareCorporationLeaderboard
	err := c.get(ctx, "/fw/leaderboards/corporations/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFactionWarfareStats An overview of statistics about factions involved in faction warfare.
// GetFactionWarfareStats 参与势力战争的各势力统计信息概览.
//
// Route: GET /fw/stats/
// 路由: GET /fw/stats/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFactionWarfareStats(ctx context.Context, ifNoneMatch ...string) ([]models.FactionWarfareStats, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result []models.FactionWarfareStats
	err := c.get(ctx, "/fw/stats/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFactionWarfareSystems Ownership of faction warfare systems.
// GetFactionWarfareSystems 势力战争星系的归属权.
//
// Route: GET /fw/systems/ — This route is cached for up to 1800 seconds
// 路由: GET /fw/systems/ — 该路由缓存长达 1800 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFactionWarfareSystems(ctx context.Context, ifNoneMatch ...string) ([]models.FactionWarfareSystem, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result []models.FactionWarfareSystem
	err := c.get(ctx, "/fw/systems/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFactionWarfareWars Data about which NPC factions are at war.
// GetFactionWarfareWars 关于哪些 NPC 势力处于交战状态的数据.
//
// Route: GET /fw/wars/
// 路由: GET /fw/wars/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFactionWarfareWars(ctx context.Context, ifNoneMatch ...string) ([]models.FactionWarfareWar, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result []models.FactionWarfareWar
	err := c.get(ctx, "/fw/wars/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
