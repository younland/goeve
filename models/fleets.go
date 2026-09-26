package models

import (
	"net/url"
	"time"
)

// GetCharactersCharacterIdFleet 200 ok object.
// GetCharactersCharacterIdFleet 200 ok 对象.
type GetCharactersCharacterIdFleet struct {
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

// GetFleetsFleetIdMembers 200 ok object.
// GetFleetsFleetIdMembers 200 ok 对象.
type GetFleetsFleetIdMembers struct {
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

// GetFleetsFleetId 200 ok object.
// GetFleetsFleetId 200 ok 对象.
type GetFleetsFleetId struct {
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

// GetFleetsFleetIdWings 200 ok object.
// GetFleetsFleetIdWings 200 ok 对象.
type GetFleetsFleetIdWings struct {
	// Id id integer.
	// Id ID 整数.
	Id int64 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Squads squads array.
	// Squads 小队列表 array.
	Squads []GetFleetsFleetIdWingsSquad `json:"squads"`
}

// GetFleetsFleetIdWingsSquad squad object.
// GetFleetsFleetIdWingsSquad 小队 object.
type GetFleetsFleetIdWingsSquad struct {
	// Id id integer.
	// Id ID 整数.
	Id int64 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostFleetsFleetIdMembersInvitation invitation object.
// PostFleetsFleetIdMembersInvitation 邀请对象.
type PostFleetsFleetIdMembersInvitation struct {
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

// PostFleetsFleetIdWings 201 created object.
// PostFleetsFleetIdWings 201 created 对象.
type PostFleetsFleetIdWings struct {
	// WingId The wing_id of the newly created wing.
	// WingId 新建小队（wing）的 wing_id.
	WingId int64 `json:"wing_id"`
}

// PostFleetsFleetIdWingsWingIdSquads 201 created object.
// PostFleetsFleetIdWingsWingIdSquads 201 created 对象.
type PostFleetsFleetIdWingsWingIdSquads struct {
	// SquadId The squad_id of the newly created squad.
	// SquadId 新建小队的 squad_id.
	SquadId int64 `json:"squad_id"`
}

// PutFleetsFleetIdMembersMemberIdMovement movement object.
// PutFleetsFleetIdMembersMemberIdMovement movement 对象.
type PutFleetsFleetIdMembersMemberIdMovement struct {
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

// PutFleetsFleetIdNewSettings new_settings object.
// PutFleetsFleetIdNewSettings new_settings 对象.
type PutFleetsFleetIdNewSettings struct {
	// IsFreeMove Should free-move be enabled in the fleet.
	// IsFreeMove 舰队中是否启用自由移动.
	IsFreeMove bool `json:"is_free_move"`
	// Motd New fleet MOTD in CCP flavoured HTML.
	// Motd CCP 风格 HTML 格式的新舰队公告（MOTD）
	Motd string `json:"motd"`
}

// PutFleetsFleetIdSquadsSquadIdNaming naming object.
// PutFleetsFleetIdSquadsSquadIdNaming naming 对象.
type PutFleetsFleetIdSquadsSquadIdNaming struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PutFleetsFleetIdWingsWingIdNaming naming object.
// PutFleetsFleetIdWingsWingIdNaming naming 对象.
type PutFleetsFleetIdWingsWingIdNaming struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// DeleteFleetIdMembersMemberIdParams holds the optional query and header parameters of the request.
// DeleteFleetIdMembersMemberIdParams 保存请求的可选查询与头部参数。
type DeleteFleetIdMembersMemberIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *DeleteFleetIdMembersMemberIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// DeleteFleetIdSquadsSquadIdParams holds the optional query and header parameters of the request.
// DeleteFleetIdSquadsSquadIdParams 保存请求的可选查询与头部参数。
type DeleteFleetIdSquadsSquadIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *DeleteFleetIdSquadsSquadIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// DeleteFleetIdWingsWingIdParams holds the optional query and header parameters of the request.
// DeleteFleetIdWingsWingIdParams 保存请求的可选查询与头部参数。
type DeleteFleetIdWingsWingIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *DeleteFleetIdWingsWingIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharactersCharacterIdFleetParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdFleetParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdFleetParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdFleetParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetFleetIdMembersParams holds the optional query and header parameters of the request.
// GetFleetIdMembersParams 保存请求的可选查询与头部参数。
type GetFleetIdMembersParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetFleetIdMembersParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetFleetIdParams holds the optional query and header parameters of the request.
// GetFleetIdParams 保存请求的可选查询与头部参数。
type GetFleetIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetFleetIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetFleetIdWingsParams holds the optional query and header parameters of the request.
// GetFleetIdWingsParams 保存请求的可选查询与头部参数。
type GetFleetIdWingsParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetFleetIdWingsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostFleetIdMembersParams holds the optional query and header parameters of the request.
// PostFleetIdMembersParams 保存请求的可选查询与头部参数。
type PostFleetIdMembersParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostFleetIdMembersParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostFleetIdWingsParams holds the optional query and header parameters of the request.
// PostFleetIdWingsParams 保存请求的可选查询与头部参数。
type PostFleetIdWingsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostFleetIdWingsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostFleetIdWingsWingIdSquadsParams holds the optional query and header parameters of the request.
// PostFleetIdWingsWingIdSquadsParams 保存请求的可选查询与头部参数。
type PostFleetIdWingsWingIdSquadsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostFleetIdWingsWingIdSquadsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PutFleetIdMembersMemberIdParams holds the optional query and header parameters of the request.
// PutFleetIdMembersMemberIdParams 保存请求的可选查询与头部参数。
type PutFleetIdMembersMemberIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PutFleetIdMembersMemberIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PutFleetIdParams holds the optional query and header parameters of the request.
// PutFleetIdParams 保存请求的可选查询与头部参数。
type PutFleetIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PutFleetIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PutFleetIdSquadsSquadIdParams holds the optional query and header parameters of the request.
// PutFleetIdSquadsSquadIdParams 保存请求的可选查询与头部参数。
type PutFleetIdSquadsSquadIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PutFleetIdSquadsSquadIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PutFleetIdWingsWingIdParams holds the optional query and header parameters of the request.
// PutFleetIdWingsWingIdParams 保存请求的可选查询与头部参数。
type PutFleetIdWingsWingIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PutFleetIdWingsWingIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}
