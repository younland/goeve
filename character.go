package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharacterId Get character's public information.
// GetCharacterId 获取角色公开信息.
//
// Route: GET /characters/{character_id}/ — This route is cached for up to 604800 seconds
// 路由: GET /characters/{character_id}/ — 该路由缓存长达 604800 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCharacterId(ctx context.Context, characterId int32, params *models.GetCharacterIdParams) (*models.GetCharactersCharacterId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterId
	err := c.get(ctx, "/characters/{character_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdAgentsResearch Get agents research.
// GetCharacterIdAgentsResearch 获取代理人研究.
//
// Route: GET /characters/{character_id}/agents_research/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/agents_research/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_agents_research.v1
// 权限: esi-characters.read_agents_research.v1
func (c *Client) GetCharacterIdAgentsResearch(ctx context.Context, characterId int32, params *models.GetCharacterIdAgentsResearchParams) ([]models.GetCharactersCharacterIdAgentsResearch, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdAgentsResearch
	err := c.get(ctx, "/characters/{character_id}/agents_research/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdBlueprints Get blueprints.
// GetCharacterIdBlueprints 获取蓝图.
//
// Route: GET /characters/{character_id}/blueprints/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/blueprints/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_blueprints.v1
// 权限: esi-characters.read_blueprints.v1
func (c *Client) GetCharacterIdBlueprints(ctx context.Context, characterId int32, params *models.GetCharacterIdBlueprintsParams) ([]models.GetCharactersCharacterIdBlueprints, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdBlueprints
	err := c.get(ctx, "/characters/{character_id}/blueprints/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdCorporationhistory Get corporation history.
// GetCharacterIdCorporationhistory 获取军团历史.
//
// Route: GET /characters/{character_id}/corporationhistory/ — This route is cached for up to 86400 seconds
// 路由: GET /characters/{character_id}/corporationhistory/ — 该路由缓存长达 86400 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCharacterIdCorporationhistory(ctx context.Context, characterId int32, params *models.GetCharacterIdCorporationhistoryParams) ([]models.GetCharactersCharacterIdCorporationhistory, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdCorporationhistory
	err := c.get(ctx, "/characters/{character_id}/corporationhistory/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdFatigue Get jump fatigue.
// GetCharacterIdFatigue 获取跳跃疲劳.
//
// Route: GET /characters/{character_id}/fatigue/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/fatigue/ — 该路由缓存长达 300 秒
// Scopes: esi-characters.read_fatigue.v1
// 权限: esi-characters.read_fatigue.v1
func (c *Client) GetCharacterIdFatigue(ctx context.Context, characterId int32, params *models.GetCharacterIdFatigueParams) (*models.GetCharactersCharacterIdFatigue, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdFatigue
	err := c.get(ctx, "/characters/{character_id}/fatigue/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdMedals Get medals.
// GetCharacterIdMedals 获取勋章.
//
// Route: GET /characters/{character_id}/medals/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/medals/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_medals.v1
// 权限: esi-characters.read_medals.v1
func (c *Client) GetCharacterIdMedals(ctx context.Context, characterId int32, params *models.GetCharacterIdMedalsParams) ([]models.GetCharactersCharacterIdMedals, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdMedals
	err := c.get(ctx, "/characters/{character_id}/medals/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdNotifications Get character notifications.
// GetCharacterIdNotifications 获取角色通知.
//
// Route: GET /characters/{character_id}/notifications/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/notifications/ — 该路由缓存长达 600 秒
// Scopes: esi-characters.read_notifications.v1
// 权限: esi-characters.read_notifications.v1
func (c *Client) GetCharacterIdNotifications(ctx context.Context, characterId int32, params *models.GetCharacterIdNotificationsParams) ([]models.GetCharactersCharacterIdNotifications, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdNotifications
	err := c.get(ctx, "/characters/{character_id}/notifications/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdNotificationsContacts Get new contact notifications.
// GetCharacterIdNotificationsContacts 获取新的联系人通知.
//
// Route: GET /characters/{character_id}/notifications/contacts/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/notifications/contacts/ — 该路由缓存长达 600 秒
// Scopes: esi-characters.read_notifications.v1
// 权限: esi-characters.read_notifications.v1
func (c *Client) GetCharacterIdNotificationsContacts(ctx context.Context, characterId int32, params *models.GetCharacterIdNotificationsContactsParams) ([]models.GetCharactersCharacterIdNotificationsContacts, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdNotificationsContacts
	err := c.get(ctx, "/characters/{character_id}/notifications/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdPortrait Get character portraits.
// GetCharacterIdPortrait 获取角色头像.
//
// Route: GET /characters/{character_id}/portrait/
// 路由: GET /characters/{character_id}/portrait/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetCharacterIdPortrait(ctx context.Context, characterId int32, params *models.GetCharacterIdPortraitParams) (*models.GetCharactersCharacterIdPortrait, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdPortrait
	err := c.get(ctx, "/characters/{character_id}/portrait/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdRoles Get character corporation roles.
// GetCharacterIdRoles 获取角色在军团内的职务.
//
// Route: GET /characters/{character_id}/roles/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/roles/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_corporation_roles.v1
// 权限: esi-characters.read_corporation_roles.v1
func (c *Client) GetCharacterIdRoles(ctx context.Context, characterId int32, params *models.GetCharacterIdRolesParams) (*models.GetCharactersCharacterIdRoles, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdRoles
	err := c.get(ctx, "/characters/{character_id}/roles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdStandings Get standings.
// GetCharacterIdStandings 获取声望.
//
// Route: GET /characters/{character_id}/standings/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/standings/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_standings.v1
// 权限: esi-characters.read_standings.v1
func (c *Client) GetCharacterIdStandings(ctx context.Context, characterId int32, params *models.GetCharacterIdStandingsParams) ([]models.GetCharactersCharacterIdStandings, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdStandings
	err := c.get(ctx, "/characters/{character_id}/standings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterIdTitles Get character corporation titles.
// GetCharacterIdTitles 获取角色在军团内的头衔.
//
// Route: GET /characters/{character_id}/titles/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/titles/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_titles.v1
// 权限: esi-characters.read_titles.v1
func (c *Client) GetCharacterIdTitles(ctx context.Context, characterId int32, params *models.GetCharacterIdTitlesParams) ([]models.GetCharactersCharacterIdTitles, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdTitles
	err := c.get(ctx, "/characters/{character_id}/titles/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostAffiliation Character affiliation.
// PostAffiliation 角色所属阵营.
//
// Route: POST /characters/affiliation/ — This route is cached for up to 3600 seconds
// 路由: POST /characters/affiliation/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) PostAffiliation(ctx context.Context, body []int32, params *models.PostAffiliationParams) ([]models.PostCharactersAffiliation, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	if body == nil {
		return nil, errBodyRequired
	}
	var result []models.PostCharactersAffiliation
	err := c.post(ctx, "/characters/affiliation/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCharacterIdCspa Calculate a CSPA charge cost.
// PostCharacterIdCspa 计算 CSPA 收费成本.
//
// Route: POST /characters/{character_id}/cspa/
// 路由: POST /characters/{character_id}/cspa/
// Scopes: esi-characters.read_contacts.v1
// 权限: esi-characters.read_contacts.v1
func (c *Client) PostCharacterIdCspa(ctx context.Context, characterId int32, body []int32, params *models.PostCharacterIdCspaParams) (float64, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
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
