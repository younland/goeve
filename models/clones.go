package models

import (
	"net/url"
	"time"
)

// GetCharactersCharacterIdClonesHomeLocation home_location object.
// GetCharactersCharacterIdClonesHomeLocation 常驻地点对象.
type GetCharactersCharacterIdClonesHomeLocation struct {
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LocationType location_type string.
	// LocationType location_type 字符串.
	// Enum values: "station", "structure".
	LocationType string `json:"location_type"`
}

// GetCharactersCharacterIdClonesJumpClone jump_clone object.
// GetCharactersCharacterIdClonesJumpClone jump_clone 对象.
type GetCharactersCharacterIdClonesJumpClone struct {
	// Implants implants array.
	// Implants 植入体数组.
	Implants []int32 `json:"implants"`
	// JumpCloneId jump_clone_id integer.
	// JumpCloneId jump_clone_id 整数.
	JumpCloneId int32 `json:"jump_clone_id"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LocationType location_type string.
	// LocationType location_type 字符串.
	// Enum values: "station", "structure".
	LocationType string `json:"location_type"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// GetCharactersCharacterIdClones 200 ok object.
// GetCharactersCharacterIdClones 200 ok 对象.
type GetCharactersCharacterIdClones struct {
	// HomeLocation home_location object.
	// HomeLocation 常驻地点对象.
	HomeLocation GetCharactersCharacterIdClonesHomeLocation `json:"home_location"`
	// JumpClones jump_clones array.
	// JumpClones jump_clones 数组.
	JumpClones []GetCharactersCharacterIdClonesJumpClone `json:"jump_clones"`
	// LastCloneJumpDate last_clone_jump_date string.
	// LastCloneJumpDate last_clone_jump_date 字符串.
	LastCloneJumpDate time.Time `json:"last_clone_jump_date"`
	// LastStationChangeDate last_station_change_date string.
	// LastStationChangeDate last_station_change_date 字符串.
	LastStationChangeDate time.Time `json:"last_station_change_date"`
}

// GetCharactersCharacterIdClonesParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdClonesParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdClonesParams struct {
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

func (p *GetCharactersCharacterIdClonesParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdImplantsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdImplantsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdImplantsParams struct {
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

func (p *GetCharactersCharacterIdImplantsParams) Values() (url.Values, map[string]string) {
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
