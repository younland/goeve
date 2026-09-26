package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacter Get character's public information.
// GetCharacter 获取角色公开信息.
//
// Route: GET /characters/{character_id}/ — This route is cached for up to 604800 seconds
// 路由: GET /characters/{character_id}/ — 该路由缓存长达 604800 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCharacter(ctx context.Context, characterID int32, ifNoneMatch ...string) (*models.Character, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.Character
	err := c.get(ctx, "/characters/{character_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterAgentsResearch Get agents research.
// GetCharacterAgentsResearch 获取代理人研究.
//
// Route: GET /characters/{character_id}/agents_research/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/agents_research/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_agents_research.v1
// 权限: esi-characters.read_agents_research.v1
func (c *Client) GetCharacterAgentsResearch(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.AgentResearch, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.AgentResearch
	err := c.get(ctx, "/characters/{character_id}/agents_research/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterBlueprints Get blueprints.
// GetCharacterBlueprints 获取蓝图.
//
// Route: GET /characters/{character_id}/blueprints/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/blueprints/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_blueprints.v1
// 权限: esi-characters.read_blueprints.v1
func (c *Client) GetCharacterBlueprints(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.Blueprint, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.Blueprint
	err := c.get(ctx, "/characters/{character_id}/blueprints/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterCorporationHistory Get corporation history.
// GetCharacterCorporationHistory 获取军团历史.
//
// Route: GET /characters/{character_id}/corporationhistory/ — This route is cached for up to 86400 seconds
// 路由: GET /characters/{character_id}/corporationhistory/ — 该路由缓存长达 86400 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCharacterCorporationHistory(ctx context.Context, characterID int32, ifNoneMatch ...string) ([]models.CorporationHistoryEntry, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CorporationHistoryEntry
	err := c.get(ctx, "/characters/{character_id}/corporationhistory/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterJumpFatigue Get jump fatigue.
// GetCharacterJumpFatigue 获取跳跃疲劳.
//
// Route: GET /characters/{character_id}/fatigue/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/fatigue/ — 该路由缓存长达 300 秒
// Scopes: esi-characters.read_fatigue.v1
// 权限: esi-characters.read_fatigue.v1
func (c *Client) GetCharacterJumpFatigue(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.JumpFatigue, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.JumpFatigue
	err := c.get(ctx, "/characters/{character_id}/fatigue/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterMedals Get medals.
// GetCharacterMedals 获取勋章.
//
// Route: GET /characters/{character_id}/medals/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/medals/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_medals.v1
// 权限: esi-characters.read_medals.v1
func (c *Client) GetCharacterMedals(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Medal, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.Medal
	err := c.get(ctx, "/characters/{character_id}/medals/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterNotifications Get character notifications.
// GetCharacterNotifications 获取角色通知.
//
// Route: GET /characters/{character_id}/notifications/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/notifications/ — 该路由缓存长达 600 秒
// Scopes: esi-characters.read_notifications.v1
// 权限: esi-characters.read_notifications.v1
func (c *Client) GetCharacterNotifications(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Notification, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.Notification
	err := c.get(ctx, "/characters/{character_id}/notifications/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterContactNotifications Get new contact notifications.
// GetCharacterContactNotifications 获取新的联系人通知.
//
// Route: GET /characters/{character_id}/notifications/contacts/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/notifications/contacts/ — 该路由缓存长达 600 秒
// Scopes: esi-characters.read_notifications.v1
// 权限: esi-characters.read_notifications.v1
func (c *Client) GetCharacterContactNotifications(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.ContactNotification, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.ContactNotification
	err := c.get(ctx, "/characters/{character_id}/notifications/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterPortrait Get character portraits.
// GetCharacterPortrait 获取角色头像.
//
// Route: GET /characters/{character_id}/portrait/
// 路由: GET /characters/{character_id}/portrait/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCharacterPortrait(ctx context.Context, characterID int32, ifNoneMatch ...string) (*models.CharacterPortraits, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.CharacterPortraits
	err := c.get(ctx, "/characters/{character_id}/portrait/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterCorporationRoles Get character corporation roles.
// GetCharacterCorporationRoles 获取角色在军团内的职务.
//
// Route: GET /characters/{character_id}/roles/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/roles/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_corporation_roles.v1
// 权限: esi-characters.read_corporation_roles.v1
func (c *Client) GetCharacterCorporationRoles(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterCorporationRoles, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.CharacterCorporationRoles
	err := c.get(ctx, "/characters/{character_id}/roles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterStandings Get standings.
// GetCharacterStandings 获取声望.
//
// Route: GET /characters/{character_id}/standings/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/standings/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_standings.v1
// 权限: esi-characters.read_standings.v1
func (c *Client) GetCharacterStandings(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Standing, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.Standing
	err := c.get(ctx, "/characters/{character_id}/standings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterCorporationTitles Get character corporation titles.
// GetCharacterCorporationTitles 获取角色在军团内的头衔.
//
// Route: GET /characters/{character_id}/titles/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/titles/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_titles.v1
// 权限: esi-characters.read_titles.v1
func (c *Client) GetCharacterCorporationTitles(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.CharacterTitle, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterTitle
	err := c.get(ctx, "/characters/{character_id}/titles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CharacterAffiliation Character affiliation.
// CharacterAffiliation 角色所属阵营.
//
// Route: POST /characters/affiliation/ — This route is cached for up to 3600 seconds
// 路由: POST /characters/affiliation/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) CharacterAffiliation(ctx context.Context, body []int32) ([]models.CharacterAffiliation, error) {
	query := url.Values{}
	headers := map[string]string{}
	var pathParams map[string]string
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.CharacterAffiliation
	err := c.post(ctx, "/characters/affiliation/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CalculateCharacterCspaCharge Calculate a CSPA charge cost.
// CalculateCharacterCspaCharge 计算 CSPA 收费成本.
//
// Route: POST /characters/{character_id}/cspa/
// 路由: POST /characters/{character_id}/cspa/
// Scopes: esi-characters.read_contacts.v1
// 权限: esi-characters.read_contacts.v1
func (c *Client) CalculateCharacterCspaCharge(ctx context.Context, token string, characterID int32, body []int32) (float64, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	if body == nil {
		return 0, errBodyRequired
	}
	var result float64
	err := c.post(ctx, "/characters/{character_id}/cspa/", pathParams, query, headers, body, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}
