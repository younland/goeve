package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdFwStats Overview of a character involved in faction warfare.
// GetCharactersCharacterIdFwStats 参与势力战争的角色的概况.
//
// Route: GET /characters/{character_id}/fw/stats/
// 路由: GET /characters/{character_id}/fw/stats/
// Scopes: esi-characters.read_fw_stats.v1
// 权限: esi-characters.read_fw_stats.v1
func (c *Client) GetCharactersCharacterIdFwStats(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdFwStatsParams) (*models.GetCharactersCharacterIdFwStats, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdFwStats
	err := c.get(ctx, "/characters/{character_id}/fw/stats/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdFwStats Overview of a corporation involved in faction warfare.
// GetCorporationsCorporationIdFwStats 参与势力战争的军团的概况.
//
// Route: GET /corporations/{corporation_id}/fw/stats/
// 路由: GET /corporations/{corporation_id}/fw/stats/
// Scopes: esi-corporations.read_fw_stats.v1
// 权限: esi-corporations.read_fw_stats.v1
func (c *Client) GetCorporationsCorporationIdFwStats(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdFwStatsParams) (*models.GetCorporationsCorporationIdFwStats, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result *models.GetCorporationsCorporationIdFwStats
	err := c.get(ctx, "/corporations/{corporation_id}/fw/stats/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFwLeaderboards List of the top factions in faction warfare.
// GetFwLeaderboards 势力战争排名前列的势力列表.
//
// Route: GET /fw/leaderboards/
// 路由: GET /fw/leaderboards/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFwLeaderboards(ctx context.Context, params *models.GetFwLeaderboardsParams) (*models.GetFwLeaderboards, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result *models.GetFwLeaderboards
	err := c.get(ctx, "/fw/leaderboards/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFwLeaderboardsCharacters List of the top pilots in faction warfare.
// GetFwLeaderboardsCharacters 势力战争排名前列的飞行员列表.
//
// Route: GET /fw/leaderboards/characters/
// 路由: GET /fw/leaderboards/characters/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFwLeaderboardsCharacters(ctx context.Context, params *models.GetFwLeaderboardsCharactersParams) (*models.GetFwLeaderboardsCharacters, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result *models.GetFwLeaderboardsCharacters
	err := c.get(ctx, "/fw/leaderboards/characters/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFwLeaderboardsCorporations List of the top corporations in faction warfare.
// GetFwLeaderboardsCorporations 势力战争排名前列的军团列表.
//
// Route: GET /fw/leaderboards/corporations/
// 路由: GET /fw/leaderboards/corporations/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFwLeaderboardsCorporations(ctx context.Context, params *models.GetFwLeaderboardsCorporationsParams) (*models.GetFwLeaderboardsCorporations, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result *models.GetFwLeaderboardsCorporations
	err := c.get(ctx, "/fw/leaderboards/corporations/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFwStats An overview of statistics about factions involved in faction warfare.
// GetFwStats 参与势力战争的各势力统计信息概览.
//
// Route: GET /fw/stats/
// 路由: GET /fw/stats/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFwStats(ctx context.Context, params *models.GetFwStatsParams) ([]models.GetFwStats, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetFwStats
	err := c.get(ctx, "/fw/stats/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFwSystems Ownership of faction warfare systems.
// GetFwSystems 势力战争星系的归属权.
//
// Route: GET /fw/systems/ — This route is cached for up to 1800 seconds
// 路由: GET /fw/systems/ — 该路由缓存长达 1800 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFwSystems(ctx context.Context, params *models.GetFwSystemsParams) ([]models.GetFwSystems, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetFwSystems
	err := c.get(ctx, "/fw/systems/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFwWars Data about which NPC factions are at war.
// GetFwWars 关于哪些 NPC 势力处于交战状态的数据.
//
// Route: GET /fw/wars/
// 路由: GET /fw/wars/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFwWars(ctx context.Context, params *models.GetFwWarsParams) ([]models.GetFwWars, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetFwWars
	err := c.get(ctx, "/fw/wars/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
