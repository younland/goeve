package models

import (
	"time"
)

// FleetMembership 200 ok object.
// FleetMembership 200 ok 对象.
type FleetMembership struct {
	// FleetId The character's current fleet ID.
	// FleetId 角色当前的舰队 ID.
	FleetId int64 `json:"fleet_id"`
	// Role Member’s role in fleet.
	// Role 成员在舰队中的职位.
	// Enum values: "fleet_commander", "squad_commander", "squad_member", "wing_commander".
	Role string `json:"role"`
	// SquadId ID of the squad the member is in. If not applicable, will be set to -1.
	// SquadId 成员所在小队的 ID。如不适用，则为 -1.
	SquadId int64 `json:"squad_id"`
	// WingId ID of the wing the member is in. If not applicable, will be set to -1.
	// WingId 成员所在联队的 ID。如不适用，则为 -1.
	WingId int64 `json:"wing_id"`
}

// FleetMember 200 ok object.
// FleetMember 200 ok 对象.
type FleetMember struct {
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// JoinTime join_time string.
	// JoinTime join_time 字符串.
	JoinTime time.Time `json:"join_time"`
	// Role Member’s role in fleet.
	// Role 成员在舰队中的职位.
	// Enum values: "fleet_commander", "wing_commander", "squad_commander", "squad_member".
	Role string `json:"role"`
	// RoleName Localized role names.
	// RoleName 本地化的职位名称.
	RoleName string `json:"role_name"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
	// SolarSystemId Solar system the member is located in.
	// SolarSystemId 成员所在的星系.
	SolarSystemId int32 `json:"solar_system_id"`
	// SquadId ID of the squad the member is in. If not applicable, will be set to -1.
	// SquadId 成员所在小队的 ID。如不适用，则为 -1.
	SquadId int64 `json:"squad_id"`
	// StationId Station in which the member is docked in, if applicable.
	// StationId 该成员停靠的空间站（如适用）
	StationId int64 `json:"station_id"`
	// TakesFleetWarp Whether the member take fleet warps.
	// TakesFleetWarp 该成员是否接受舰队跃迁.
	TakesFleetWarp bool `json:"takes_fleet_warp"`
	// WingId ID of the wing the member is in. If not applicable, will be set to -1.
	// WingId 成员所在联队的 ID。如不适用，则为 -1.
	WingId int64 `json:"wing_id"`
}

// Fleet 200 ok object.
// Fleet 200 ok 对象.
type Fleet struct {
	// IsFreeMove Is free-move enabled.
	// IsFreeMove 是否启用自由移动.
	IsFreeMove bool `json:"is_free_move"`
	// IsRegistered Does the fleet have an active fleet advertisement.
	// IsRegistered 该舰队是否有活跃的舰队招募广告.
	IsRegistered bool `json:"is_registered"`
	// IsVoiceEnabled Is EVE Voice enabled.
	// IsVoiceEnabled 是否启用 EVE Voice.
	IsVoiceEnabled bool `json:"is_voice_enabled"`
	// Motd Fleet MOTD in CCP flavoured HTML.
	// Motd 以 CCP 风格 HTML 表示的舰队每日公告（MOTD）
	Motd string `json:"motd"`
}

// FleetWing 200 ok object.
// FleetWing 200 ok 对象.
type FleetWing struct {
	// Id id integer.
	// Id ID 整数.
	Id int64 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Squads squads array.
	// Squads 小队列表 array.
	Squads []FleetSquad `json:"squads"`
}

// FleetSquad squad object.
// FleetSquad 小队 object.
type FleetSquad struct {
	// Id id integer.
	// Id ID 整数.
	Id int64 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// FleetInvitation invitation object.
// FleetInvitation 邀请对象.
type FleetInvitation struct {
	// CharacterId The character you want to invite.
	// CharacterId 你要邀请的角色.
	CharacterId int32 `json:"character_id"`
	// Role If a character is invited with the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified. If a character is invited with the `wing_commander` role, only `wing_id` should be specified. If a character is invited with the `squad_commander` role, both `wing_id` and `squad_id` should be specified. If a character is invited with the `squad_member` role, `wing_id` and `squad_id` should either both be specified or not specified at all. If they aren’t specified, the invited character will join any squad with available positions.
	// Role 若以 `fleet_commander`（舰队指挥官）角色邀请角色，则不应指定 `wing_id` 或 `squad_id`。若以 `wing_commander`（联队指挥官）角色邀请，则只应指定 `wing_id`。若以 `squad_commander`（小队指挥官）角色邀请，则应同时指定 `wing_id` 和 `squad_id`。若以 `squad_member`（小队成员）角色邀请，则 `wing_id` 和 `squad_id` 应同时指定或同时不指定。若未指定，被邀请的角色将加入任何有空位的小队。
	// Enum values: "fleet_commander", "wing_commander", "squad_commander", "squad_member".
	Role string `json:"role"`
	// SquadId squad_id integer.
	// SquadId 小队ID integer.
	SquadId int64 `json:"squad_id"`
	// WingId wing_id integer.
	// WingId wing_id 整数.
	WingId int64 `json:"wing_id"`
}

// NewFleetWing 201 created object.
// NewFleetWing 201 created 对象.
type NewFleetWing struct {
	// WingId The wing_id of the newly created wing.
	// WingId 新建小队（wing）的 wing_id.
	WingId int64 `json:"wing_id"`
}

// NewFleetSquad 201 created object.
// NewFleetSquad 201 created 对象.
type NewFleetSquad struct {
	// SquadId The squad_id of the newly created squad.
	// SquadId 新建小队的 squad_id.
	SquadId int64 `json:"squad_id"`
}

// FleetMemberMovement movement object.
// FleetMemberMovement movement 对象.
type FleetMemberMovement struct {
	// Role If a character is moved to the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified. If a character is moved to the `wing_commander` role, only `wing_id` should be specified. If a character is moved to the `squad_commander` role, both `wing_id` and `squad_id` should be specified. If a character is moved to the `squad_member` role, both `wing_id` and `squad_id` should be specified.
	// Role 若角色被调整为 `fleet_commander`（舰队指挥官）角色，则不应指定 `wing_id` 或 `squad_id`。若被调整为 `wing_commander`（联队指挥官）角色，则只应指定 `wing_id`。若被调整为 `squad_commander`（小队指挥官）角色，则应同时指定 `wing_id` 和 `squad_id`。若被调整为 `squad_member`（小队成员）角色，则应同时指定 `wing_id` 和 `squad_id`。
	// Enum values: "fleet_commander", "wing_commander", "squad_commander", "squad_member".
	Role string `json:"role"`
	// SquadId squad_id integer.
	// SquadId 小队ID integer.
	SquadId int64 `json:"squad_id"`
	// WingId wing_id integer.
	// WingId wing_id 整数.
	WingId int64 `json:"wing_id"`
}

// FleetSettings new_settings object.
// FleetSettings new_settings 对象.
type FleetSettings struct {
	// IsFreeMove Should free-move be enabled in the fleet.
	// IsFreeMove 舰队中是否启用自由移动.
	IsFreeMove bool `json:"is_free_move"`
	// Motd New fleet MOTD in CCP flavoured HTML.
	// Motd CCP 风格 HTML 格式的新舰队公告（MOTD）
	Motd string `json:"motd"`
}

// FleetNaming naming object.
// FleetNaming naming 对象.
type FleetNaming struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}
