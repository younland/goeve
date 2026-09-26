package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdLocation Get character location.
// GetCharactersCharacterIdLocation 获取角色位置.
//
// Route: GET /characters/{character_id}/location/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/location/ — 该路由缓存长达 5 秒
// Scopes: esi-location.read_location.v1
// 权限: esi-location.read_location.v1
func (c *Client) GetCharactersCharacterIdLocation(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdLocationParams) (*models.GetCharactersCharacterIdLocation, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdLocation
	err := c.get(ctx, "/characters/{character_id}/location/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdOnline Get character online.
// GetCharactersCharacterIdOnline 获取角色在线状态.
//
// Route: GET /characters/{character_id}/online/ — This route is cached for up to 60 seconds
// 路由: GET /characters/{character_id}/online/ — 该路由缓存长达 60 秒
// Scopes: esi-location.read_online.v1
// 权限: esi-location.read_online.v1
func (c *Client) GetCharactersCharacterIdOnline(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOnlineParams) (*models.GetCharactersCharacterIdOnline, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdOnline
	err := c.get(ctx, "/characters/{character_id}/online/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdShip Get current ship.
// GetCharactersCharacterIdShip 获取当前舰船.
//
// Route: GET /characters/{character_id}/ship/ — This route is cached for up to 5 seconds
// 路由: GET /characters/{character_id}/ship/ — 该路由缓存长达 5 秒
// Scopes: esi-location.read_ship_type.v1
// 权限: esi-location.read_ship_type.v1
func (c *Client) GetCharactersCharacterIdShip(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdShipParams) (*models.GetCharactersCharacterIdShip, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdShip
	err := c.get(ctx, "/characters/{character_id}/ship/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
