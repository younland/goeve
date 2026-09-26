package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdAttributes Get character attributes.
// GetCharactersCharacterIdAttributes 获取角色属性.
//
// Route: GET /characters/{character_id}/attributes/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/attributes/ — 该路由缓存长达 120 秒
// Scopes: esi-skills.read_skills.v1
// 权限: esi-skills.read_skills.v1
func (c *Client) GetCharactersCharacterIdAttributes(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdAttributesParams) (*models.GetCharactersCharacterIdAttributesOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdAttributesOk
	err := c.get(ctx, "/characters/{character_id}/attributes/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdSkillqueue Get character's skill queue.
// GetCharactersCharacterIdSkillqueue 获取角色技能队列.
//
// Route: GET /characters/{character_id}/skillqueue/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/skillqueue/ — 该路由缓存长达 120 秒
// Scopes: esi-skills.read_skillqueue.v1
// 权限: esi-skills.read_skillqueue.v1
func (c *Client) GetCharactersCharacterIdSkillqueue(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdSkillqueueParams) ([]models.GetCharactersCharacterIdSkillqueue, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdSkillqueue
	err := c.get(ctx, "/characters/{character_id}/skillqueue/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdSkills Get character skills.
// GetCharactersCharacterIdSkills 获取角色技能.
//
// Route: GET /characters/{character_id}/skills/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/skills/ — 该路由缓存长达 120 秒
// Scopes: esi-skills.read_skills.v1
// 权限: esi-skills.read_skills.v1
func (c *Client) GetCharactersCharacterIdSkills(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdSkillsParams) (*models.GetCharactersCharacterIdSkillsOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdSkillsOk
	err := c.get(ctx, "/characters/{character_id}/skills/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
