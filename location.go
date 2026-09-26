package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterLocation Get character location.
// GetCharacterLocation 获取角色位置.
//
// Route: GET /characters/{character_id}/location/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/location/ — 该路由缓存长达 5 秒
// Scopes: esi-location.read_location.v1
// 权限: esi-location.read_location.v1
func (c *Client) GetCharacterLocation(ctx context.Context, characterID int32, token string, ifNoneMatch ...string) (*models.CharacterLocation, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.CharacterLocation
	err := c.get(ctx, "/characters/{character_id}/location/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterOnline Get character online.
// GetCharacterOnline 获取角色在线状态.
//
// Route: GET /characters/{character_id}/online/ — This route is cached for up to 60 seconds
// 路由: GET /characters/{character_id}/online/ — 该路由缓存长达 60 秒
// Scopes: esi-location.read_online.v1
// 权限: esi-location.read_online.v1
func (c *Client) GetCharacterOnline(ctx context.Context, characterID int32, token string, ifNoneMatch ...string) (*models.OnlineStatus, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.OnlineStatus
	err := c.get(ctx, "/characters/{character_id}/online/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterShip Get current ship.
// GetCharacterShip 获取当前舰船.
//
// Route: GET /characters/{character_id}/ship/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/ship/ — 该路由缓存长达 5 秒
// Scopes: esi-location.read_ship_type.v1
// 权限: esi-location.read_ship_type.v1
func (c *Client) GetCharacterShip(ctx context.Context, characterID int32, token string, ifNoneMatch ...string) (*models.CharacterShip, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.CharacterShip
	err := c.get(ctx, "/characters/{character_id}/ship/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
