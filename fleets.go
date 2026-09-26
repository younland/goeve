package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharacterFleet Get character fleet info.
// GetCharacterFleet 获取角色所在舰队信息.
//
// Route: GET /characters/{character_id}/fleet/ — This route is cached for up to 60 seconds
// 路由: GET /characters/{character_id}/fleet/ — 该路由缓存长达 60 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetCharacterFleet(ctx context.Context, characterID int32, opts ...RequestOption) (*models.FleetMembership, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.FleetMembership
	err := c.get(ctx, "/characters/{character_id}/fleet/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFleet Get fleet information.
// GetFleet 获取舰队信息.
//
// Route: GET /fleets/{fleet_id}/ — This route is cached for up to 5 seconds
// 路由: GET /fleets/{fleet_id}/ — 该路由缓存长达 5 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetFleet(ctx context.Context, fleetID int64, opts ...RequestOption) (*models.Fleet, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10)}
	var result *models.Fleet
	err := c.get(ctx, "/fleets/{fleet_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateFleetSettings Update fleet.
// UpdateFleetSettings 更新舰队.
//
// Route: PUT /fleets/{fleet_id}/
// 路由: PUT /fleets/{fleet_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) UpdateFleetSettings(ctx context.Context, fleetID int64, body *models.FleetSettings, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10)}
	if body == nil {
		return errBodyRequired
	}
	if err := c.put(ctx, "/fleets/{fleet_id}/", pathParams, query, headers, body); err != nil {
		return err
	}
	return nil
}

// GetFleetMembers Get fleet members.
// GetFleetMembers 获取舰队成员.
//
// Route: GET /fleets/{fleet_id}/members/ — This route is cached for up to 5 seconds
// 路由: GET /fleets/{fleet_id}/members/ — 该路由缓存长达 5 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetFleetMembers(ctx context.Context, fleetID int64, opts ...RequestOption) ([]models.FleetMember, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10)}
	var result []models.FleetMember
	err := c.get(ctx, "/fleets/{fleet_id}/members/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateFleetInvitation Create fleet invitation.
// CreateFleetInvitation 创建舰队邀请.
//
// Route: POST /fleets/{fleet_id}/members/
// 路由: POST /fleets/{fleet_id}/members/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) CreateFleetInvitation(ctx context.Context, fleetID int64, body *models.FleetInvitation, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.post(ctx, "/fleets/{fleet_id}/members/", pathParams, query, headers, body, nil)
	return err
}

// KickFleetMember Kick fleet member.
// KickFleetMember 踢出舰队成员.
//
// Route: DELETE /fleets/{fleet_id}/members/{member_id}/
// 路由: DELETE /fleets/{fleet_id}/members/{member_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) KickFleetMember(ctx context.Context, fleetID int64, memberID int32, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10), "member_id": strconv.FormatInt(int64(memberID), 10)}
	if err := c.delete(ctx, "/fleets/{fleet_id}/members/{member_id}/", pathParams, query, headers); err != nil {
		return err
	}
	return nil
}

// MoveFleetMember Move fleet member.
// MoveFleetMember 移动舰队成员.
//
// Route: PUT /fleets/{fleet_id}/members/{member_id}/
// 路由: PUT /fleets/{fleet_id}/members/{member_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) MoveFleetMember(ctx context.Context, fleetID int64, memberID int32, body *models.FleetMemberMovement, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10), "member_id": strconv.FormatInt(int64(memberID), 10)}
	if body == nil {
		return errBodyRequired
	}
	if err := c.put(ctx, "/fleets/{fleet_id}/members/{member_id}/", pathParams, query, headers, body); err != nil {
		return err
	}
	return nil
}

// DeleteFleetSquad Delete fleet squad.
// DeleteFleetSquad 删除舰队小队.
//
// Route: DELETE /fleets/{fleet_id}/squads/{squad_id}/
// 路由: DELETE /fleets/{fleet_id}/squads/{squad_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) DeleteFleetSquad(ctx context.Context, fleetID int64, squadID int64, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10), "squad_id": strconv.FormatInt(int64(squadID), 10)}
	if err := c.delete(ctx, "/fleets/{fleet_id}/squads/{squad_id}/", pathParams, query, headers); err != nil {
		return err
	}
	return nil
}

// RenameFleetSquad Rename fleet squad.
// RenameFleetSquad 重命名舰队小队.
//
// Route: PUT /fleets/{fleet_id}/squads/{squad_id}/
// 路由: PUT /fleets/{fleet_id}/squads/{squad_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) RenameFleetSquad(ctx context.Context, fleetID int64, squadID int64, body *models.FleetNaming, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10), "squad_id": strconv.FormatInt(int64(squadID), 10)}
	if body == nil {
		return errBodyRequired
	}
	if err := c.put(ctx, "/fleets/{fleet_id}/squads/{squad_id}/", pathParams, query, headers, body); err != nil {
		return err
	}
	return nil
}

// GetFleetWings Get fleet wings.
// GetFleetWings 获取舰队联队.
//
// Route: GET /fleets/{fleet_id}/wings/ — This route is cached for up to 5 seconds
// 路由: GET /fleets/{fleet_id}/wings/ — 该路由缓存长达 5 秒
// Scopes: esi-fleets.read_fleet.v1
// 权限: esi-fleets.read_fleet.v1
func (c *Client) GetFleetWings(ctx context.Context, fleetID int64, opts ...RequestOption) ([]models.FleetWing, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10)}
	var result []models.FleetWing
	err := c.get(ctx, "/fleets/{fleet_id}/wings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateFleetWing Create fleet wing.
// CreateFleetWing 创建舰队联队.
//
// Route: POST /fleets/{fleet_id}/wings/
// 路由: POST /fleets/{fleet_id}/wings/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) CreateFleetWing(ctx context.Context, fleetID int64, opts ...RequestOption) (*models.NewFleetWing, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10)}
	var result *models.NewFleetWing
	err := c.post(ctx, "/fleets/{fleet_id}/wings/", pathParams, query, headers, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteFleetWing Delete fleet wing.
// DeleteFleetWing 删除舰队联队.
//
// Route: DELETE /fleets/{fleet_id}/wings/{wing_id}/
// 路由: DELETE /fleets/{fleet_id}/wings/{wing_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) DeleteFleetWing(ctx context.Context, fleetID int64, wingID int64, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10), "wing_id": strconv.FormatInt(int64(wingID), 10)}
	if err := c.delete(ctx, "/fleets/{fleet_id}/wings/{wing_id}/", pathParams, query, headers); err != nil {
		return err
	}
	return nil
}

// RenameFleetWing Rename fleet wing.
// RenameFleetWing 重命名舰队联队.
//
// Route: PUT /fleets/{fleet_id}/wings/{wing_id}/
// 路由: PUT /fleets/{fleet_id}/wings/{wing_id}/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) RenameFleetWing(ctx context.Context, fleetID int64, wingID int64, body *models.FleetNaming, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10), "wing_id": strconv.FormatInt(int64(wingID), 10)}
	if body == nil {
		return errBodyRequired
	}
	if err := c.put(ctx, "/fleets/{fleet_id}/wings/{wing_id}/", pathParams, query, headers, body); err != nil {
		return err
	}
	return nil
}

// CreateFleetSquad Create fleet squad.
// CreateFleetSquad 创建舰队小队.
//
// Route: POST /fleets/{fleet_id}/wings/{wing_id}/squads/
// 路由: POST /fleets/{fleet_id}/wings/{wing_id}/squads/
// Scopes: esi-fleets.write_fleet.v1
// 权限: esi-fleets.write_fleet.v1
func (c *Client) CreateFleetSquad(ctx context.Context, fleetID int64, wingID int64, opts ...RequestOption) (*models.NewFleetSquad, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"fleet_id": strconv.FormatInt(int64(fleetID), 10), "wing_id": strconv.FormatInt(int64(wingID), 10)}
	var result *models.NewFleetSquad
	err := c.post(ctx, "/fleets/{fleet_id}/wings/{wing_id}/squads/", pathParams, query, headers, nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
