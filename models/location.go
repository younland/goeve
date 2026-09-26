package models

import (
	"net/url"
	"time"
)

// GetCharactersCharacterIdLocationOk 200 ok object.
// GetCharactersCharacterIdLocationOk 200 ok 对象.
type GetCharactersCharacterIdLocationOk struct {
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// StationId station_id integer.
	// StationId 空间站ID integer.
	StationId int32 `json:"station_id"`
	// StructureId structure_id integer.
	// StructureId 建筑ID integer.
	StructureId int64 `json:"structure_id"`
}

// GetCharactersCharacterIdOnlineOk 200 ok object.
// GetCharactersCharacterIdOnlineOk 200 ok 对象.
type GetCharactersCharacterIdOnlineOk struct {
	// LastLogin Timestamp of the last login.
	// LastLogin 最后一次登录的时间戳.
	LastLogin time.Time `json:"last_login"`
	// LastLogout Timestamp of the last logout.
	// LastLogout 最后一次登出的时间戳.
	LastLogout time.Time `json:"last_logout"`
	// Logins Total number of times the character has logged in.
	// Logins 该角色的累计登录次数.
	Logins int32 `json:"logins"`
	// Online If the character is online.
	// Online 角色是否在线.
	Online bool `json:"online"`
}

// GetCharactersCharacterIdShipOk 200 ok object.
// GetCharactersCharacterIdShipOk 200 ok 对象.
type GetCharactersCharacterIdShipOk struct {
	// ShipItemId Item id's are unique to a ship and persist until it is repackaged. This value can be used to track repeated uses of a ship, or detect when a pilot changes into a different instance of the same ship type.
	// ShipItemId 物品 ID 对一艘舰船唯一，并在其重新打包之前保持不变。该值可用于追踪舰船的重复使用，或检测飞行员何时更换为同型舰船的另一实例。
	ShipItemId int64 `json:"ship_item_id"`
	// ShipName ship_name string.
	// ShipName 舰船名称 string.
	ShipName string `json:"ship_name"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
}

// GetCharactersCharacterIdLocationParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdLocationParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdLocationParams struct {
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

func (p *GetCharactersCharacterIdLocationParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdOnlineParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdOnlineParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdOnlineParams struct {
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

func (p *GetCharactersCharacterIdOnlineParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdShipParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdShipParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdShipParams struct {
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

func (p *GetCharactersCharacterIdShipParams) Values() (url.Values, map[string]string) {
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
