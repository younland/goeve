package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterAttributes Get character attributes.
// GetCharacterAttributes 获取角色属性.
//
// Route: GET /characters/{character_id}/attributes/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/attributes/ — 该路由缓存长达 120 秒
// Scopes: esi-skills.read_skills.v1
// 权限: esi-skills.read_skills.v1
func (c *Client) GetCharacterAttributes(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterAttributes, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.CharacterAttributes
	err := c.get(ctx, "/characters/{character_id}/attributes/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterSkillQueue Get character's skill queue.
// GetCharacterSkillQueue 获取角色技能队列.
//
// Route: GET /characters/{character_id}/skillqueue/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/skillqueue/ — 该路由缓存长达 120 秒
// Scopes: esi-skills.read_skillqueue.v1
// 权限: esi-skills.read_skillqueue.v1
func (c *Client) GetCharacterSkillQueue(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.SkillQueueEntry, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.SkillQueueEntry
	err := c.get(ctx, "/characters/{character_id}/skillqueue/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterSkills Get character skills.
// GetCharacterSkills 获取角色技能.
//
// Route: GET /characters/{character_id}/skills/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/skills/ — 该路由缓存长达 120 秒
// Scopes: esi-skills.read_skills.v1
// 权限: esi-skills.read_skills.v1
func (c *Client) GetCharacterSkills(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterSkills, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.CharacterSkills
	err := c.get(ctx, "/characters/{character_id}/skills/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
