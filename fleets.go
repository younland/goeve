package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// DeleteFleetIdMembersMemberId Kick fleet member.
// DeleteFleetIdMembersMemberId 踢出舰队成员.
//
// Route: DELETE /fleets/{fleet_id}/members/{member_id}/
// 路由: DELETE /fleets/{fleet_id}/members/{member_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) DeleteFleetIdMembersMemberId(ctx context.Context, fleetId int64, memberId int32, params *models.DeleteFleetIdMembersMemberIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10), "member_id": strconv.FormatInt(int64(memberId), 10)}
	err := c.delete(ctx, "/fleets/{fleet_id}/members/{member_id}/", pathParams, query, headers)
	return err
}

// DeleteFleetIdSquadsSquadId Delete fleet squad.
// DeleteFleetIdSquadsSquadId 删除舰队小队.
//
// Route: DELETE /fleets/{fleet_id}/squads/{squad_id}/
// 路由: DELETE /fleets/{fleet_id}/squads/{squad_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) DeleteFleetIdSquadsSquadId(ctx context.Context, fleetId int64, squadId int64, params *models.DeleteFleetIdSquadsSquadIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10), "squad_id": strconv.FormatInt(int64(squadId), 10)}
	err := c.delete(ctx, "/fleets/{fleet_id}/squads/{squad_id}/", pathParams, query, headers)
	return err
}

// DeleteFleetIdWingsWingId Delete fleet wing.
// DeleteFleetIdWingsWingId 删除舰队联队.
//
// Route: DELETE /fleets/{fleet_id}/wings/{wing_id}/
// 路由: DELETE /fleets/{fleet_id}/wings/{wing_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) DeleteFleetIdWingsWingId(ctx context.Context, fleetId int64, wingId int64, params *models.DeleteFleetIdWingsWingIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10), "wing_id": strconv.FormatInt(int64(wingId), 10)}
	err := c.delete(ctx, "/fleets/{fleet_id}/wings/{wing_id}/", pathParams, query, headers)
	return err
}

// GetCharactersCharacterIdFleet Get character fleet info.
// GetCharactersCharacterIdFleet 获取角色所在舰队信息.
//
// Route: GET /characters/{character_id}/fleet/ — This route is cached for up to 60 seconds
// 路由: GET /characters/{character_id}/fleet/ — 该路由缓存长达 60 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetCharactersCharacterIdFleet(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdFleetParams) (*models.GetCharactersCharacterIdFleet, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdFleet
	err := c.get(ctx, "/characters/{character_id}/fleet/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFleetId Get fleet information.
// GetFleetId 获取舰队信息.
//
// Route: GET /fleets/{fleet_id}/ — This route is cached for up to 5 seconds
// 路由: GET /fleets/{fleet_id}/ — 该路由缓存长达 5 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetFleetId(ctx context.Context, fleetId int64, params *models.GetFleetIdParams) (*models.GetFleetsFleetId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10)}
	var result *models.GetFleetsFleetId
	err := c.get(ctx, "/fleets/{fleet_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFleetIdMembers Get fleet members.
// GetFleetIdMembers 获取舰队成员.
//
// Route: GET /fleets/{fleet_id}/members/ — This route is cached for up to 5 seconds
// 路由: GET /fleets/{fleet_id}/members/ — 该路由缓存长达 5 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetFleetIdMembers(ctx context.Context, fleetId int64, params *models.GetFleetIdMembersParams) ([]models.GetFleetsFleetIdMembers, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10)}
	var result []models.GetFleetsFleetIdMembers
	err := c.get(ctx, "/fleets/{fleet_id}/members/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFleetIdWings Get fleet wings.
// GetFleetIdWings 获取舰队联队.
//
// Route: GET /fleets/{fleet_id}/wings/ — This route is cached for up to 5 seconds
// 路由: GET /fleets/{fleet_id}/wings/ — 该路由缓存长达 5 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetFleetIdWings(ctx context.Context, fleetId int64, params *models.GetFleetIdWingsParams) ([]models.GetFleetsFleetIdWings, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10)}
	var result []models.GetFleetsFleetIdWings
	err := c.get(ctx, "/fleets/{fleet_id}/wings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostFleetIdMembers Create fleet invitation.
// PostFleetIdMembers 创建舰队邀请.
//
// Route: POST /fleets/{fleet_id}/members/
// 路由: POST /fleets/{fleet_id}/members/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) PostFleetIdMembers(ctx context.Context, fleetId int64, body *models.PostFleetsFleetIdMembersInvitation, params *models.PostFleetIdMembersParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.post(ctx, "/fleets/{fleet_id}/members/", pathParams, query, headers, body, nil)
	return err
}

// PostFleetIdWings Create fleet wing.
// PostFleetIdWings 创建舰队联队.
//
// Route: POST /fleets/{fleet_id}/wings/
// 路由: POST /fleets/{fleet_id}/wings/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) PostFleetIdWings(ctx context.Context, fleetId int64, params *models.PostFleetIdWingsParams) (*models.PostFleetsFleetIdWings, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10)}
	var result *models.PostFleetsFleetIdWings
	err := c.post(ctx, "/fleets/{fleet_id}/wings/", pathParams, query, headers, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostFleetIdWingsWingIdSquads Create fleet squad.
// PostFleetIdWingsWingIdSquads 创建舰队小队.
//
// Route: POST /fleets/{fleet_id}/wings/{wing_id}/squads/
// 路由: POST /fleets/{fleet_id}/wings/{wing_id}/squads/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) PostFleetIdWingsWingIdSquads(ctx context.Context, fleetId int64, wingId int64, params *models.PostFleetIdWingsWingIdSquadsParams) (*models.PostFleetsFleetIdWingsWingIdSquads, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10), "wing_id": strconv.FormatInt(int64(wingId), 10)}
	var result *models.PostFleetsFleetIdWingsWingIdSquads
	err := c.post(ctx, "/fleets/{fleet_id}/wings/{wing_id}/squads/", pathParams, query, headers, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutFleetId Update fleet.
// PutFleetId 更新舰队.
//
// Route: PUT /fleets/{fleet_id}/
// 路由: PUT /fleets/{fleet_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) PutFleetId(ctx context.Context, fleetId int64, body *models.PutFleetsFleetIdNewSettings, params *models.PutFleetIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/fleets/{fleet_id}/", pathParams, query, headers, body)
	return err
}

// PutFleetIdMembersMemberId Move fleet member.
// PutFleetIdMembersMemberId 移动舰队成员.
//
// Route: PUT /fleets/{fleet_id}/members/{member_id}/
// 路由: PUT /fleets/{fleet_id}/members/{member_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) PutFleetIdMembersMemberId(ctx context.Context, fleetId int64, memberId int32, body *models.PutFleetsFleetIdMembersMemberIdMovement, params *models.PutFleetIdMembersMemberIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10), "member_id": strconv.FormatInt(int64(memberId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/fleets/{fleet_id}/members/{member_id}/", pathParams, query, headers, body)
	return err
}

// PutFleetIdSquadsSquadId Rename fleet squad.
// PutFleetIdSquadsSquadId 重命名舰队小队.
//
// Route: PUT /fleets/{fleet_id}/squads/{squad_id}/
// 路由: PUT /fleets/{fleet_id}/squads/{squad_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) PutFleetIdSquadsSquadId(ctx context.Context, fleetId int64, squadId int64, body *models.PutFleetsFleetIdSquadsSquadIdNaming, params *models.PutFleetIdSquadsSquadIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10), "squad_id": strconv.FormatInt(int64(squadId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/fleets/{fleet_id}/squads/{squad_id}/", pathParams, query, headers, body)
	return err
}

// PutFleetIdWingsWingId Rename fleet wing.
// PutFleetIdWingsWingId 重命名舰队联队.
//
// Route: PUT /fleets/{fleet_id}/wings/{wing_id}/
// 路由: PUT /fleets/{fleet_id}/wings/{wing_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) PutFleetIdWingsWingId(ctx context.Context, fleetId int64, wingId int64, body *models.PutFleetsFleetIdWingsWingIdNaming, params *models.PutFleetIdWingsWingIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetId), 10), "wing_id": strconv.FormatInt(int64(wingId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/fleets/{fleet_id}/wings/{wing_id}/", pathParams, query, headers, body)
	return err
}
