package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdKillmailsRecent Get a character's recent kills and losses.
// GetCharactersCharacterIdKillmailsRecent 获取角色近期的击杀与损失.
//
// Route: GET /characters/{character_id}/killmails/recent/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/killmails/recent/ — 该路由缓存长达 300 秒
// Scopes: esi-killmails.read_killmails.v1
// 权限: esi-killmails.read_killmails.v1
func (c *Client) GetCharactersCharacterIdKillmailsRecent(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdKillmailsRecentParams) ([]models.GetCharactersCharacterIdKillmailsRecent, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdKillmailsRecent
	err := c.get(ctx, "/characters/{character_id}/killmails/recent/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdKillmailsRecent Get a corporation's recent kills and losses.
// GetCorporationsCorporationIdKillmailsRecent 获取军团近期的击杀与损失.
//
// Route: GET /corporations/{corporation_id}/killmails/recent/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/killmails/recent/ — 该路由缓存长达 300 秒
// Scopes: esi-killmails.read_corporation_killmails.v1
// 权限: esi-killmails.read_corporation_killmails.v1
func (c *Client) GetCorporationsCorporationIdKillmailsRecent(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdKillmailsRecentParams) ([]models.GetCorporationsCorporationIdKillmailsRecent, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdKillmailsRecent
	err := c.get(ctx, "/corporations/{corporation_id}/killmails/recent/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetKillmailIdKillmailHash Get a single killmail.
// GetKillmailIdKillmailHash 获取单条击杀报告.
//
// Route: GET /killmails/{killmail_id}/{killmail_hash}/ — This route is cached for up to 30758400 seconds
// 路由: GET /killmails/{killmail_id}/{killmail_hash}/ — 该路由缓存长达 30758400 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetKillmailIdKillmailHash(ctx context.Context, killmailHash string, killmailId int32, params *models.GetKillmailIdKillmailHashParams) (*models.GetKillmailsKillmailIdKillmailHash, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"killmail_hash": killmailHash, "killmail_id": strconv.FormatInt(int64(killmailId), 10)}
	var result *models.GetKillmailsKillmailIdKillmailHash
	err := c.get(ctx, "/killmails/{killmail_id}/{killmail_hash}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
