package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterKillmails Get a character's recent kills and losses.
// GetCharacterKillmails 获取角色近期的击杀与损失.
//
// Route: GET /characters/{character_id}/killmails/recent/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/killmails/recent/ — 该路由缓存长达 300 秒
// Scopes: esi-killmails.read_killmails.v1
// 权限: esi-killmails.read_killmails.v1
func (c *Client) GetCharacterKillmails(ctx context.Context, characterID int32, token string, page int32, ifNoneMatch string) ([]models.KillmailRef, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.KillmailRef
	err := c.get(ctx, "/characters/{character_id}/killmails/recent/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationKillmails Get a corporation's recent kills and losses.
// GetCorporationKillmails 获取军团近期的击杀与损失.
//
// Route: GET /corporations/{corporation_id}/killmails/recent/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/killmails/recent/ — 该路由缓存长达 300 秒
// Scopes: esi-killmails.read_corporation_killmails.v1
// 权限: esi-killmails.read_corporation_killmails.v1
func (c *Client) GetCorporationKillmails(ctx context.Context, corporationID int32, token string, page int32, ifNoneMatch string) ([]models.KillmailRef, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.KillmailRef
	err := c.get(ctx, "/corporations/{corporation_id}/killmails/recent/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetKillmail Get a single killmail.
// GetKillmail 获取单条击杀报告.
//
// Route: GET /killmails/{killmail_id}/{killmail_hash}/ — This route is cached for up to 30758400 seconds
// 路由: GET /killmails/{killmail_id}/{killmail_hash}/ — 该路由缓存长达 30758400 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetKillmail(ctx context.Context, killmailHash string, killmailID int32, ifNoneMatch string) (*models.Killmail, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"killmail_hash": killmailHash, "killmail_id": strconv.FormatInt(int64(killmailID), 10)}
	var result *models.Killmail
	err := c.get(ctx, "/killmails/{killmail_id}/{killmail_hash}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
