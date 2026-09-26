package models

import (
	"net/url"
	"time"
)

// GetAlliancesAllianceIdIcons 200 ok object.
// GetAlliancesAllianceIdIcons 200 ok 对象.
type GetAlliancesAllianceIdIcons struct {
	// Px128x128 px128x128 string.
	// Px128x128 px128x128 字符串.
	Px128x128 string `json:"px128x128"`
	// Px64x64 px64x64 string.
	// Px64x64 px64x64 字符串.
	Px64x64 string `json:"px64x64"`
}

// GetAlliancesAllianceId 200 ok object.
// GetAlliancesAllianceId 200 ok 对象.
type GetAlliancesAllianceId struct {
	// CreatorCorporationId ID of the corporation that created the alliance.
	// CreatorCorporationId 创建该联盟的军团 ID.
	CreatorCorporationId int32 `json:"creator_corporation_id"`
	// CreatorId ID of the character that created the alliance.
	// CreatorId 创建该联盟的角色 ID.
	CreatorId int32 `json:"creator_id"`
	// DateFounded date_founded string.
	// DateFounded 成立日期字符串.
	DateFounded time.Time `json:"date_founded"`
	// ExecutorCorporationId the executor corporation ID, if this alliance is not closed.
	// ExecutorCorporationId 执行军团 ID，若该联盟未解散.
	ExecutorCorporationId int32 `json:"executor_corporation_id"`
	// FactionId Faction ID this alliance is fighting for, if this alliance is enlisted in factional warfare.
	// FactionId 该联盟加入势力战争后所效忠的势力 ID.
	FactionId int32 `json:"faction_id"`
	// Name the full name of the alliance.
	// Name 联盟的完整名称.
	Name string `json:"name"`
	// Ticker the short name of the alliance.
	// Ticker 联盟的简称.
	Ticker string `json:"ticker"`
}

// GetAllianceIdCorporationsParams holds the optional query and header parameters of the request.
// GetAllianceIdCorporationsParams 保存请求的可选查询与头部参数。
type GetAllianceIdCorporationsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetAllianceIdCorporationsParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetAllianceIdIconsParams holds the optional query and header parameters of the request.
// GetAllianceIdIconsParams 保存请求的可选查询与头部参数。
type GetAllianceIdIconsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetAllianceIdIconsParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetAllianceIdParams holds the optional query and header parameters of the request.
// GetAllianceIdParams 保存请求的可选查询与头部参数。
type GetAllianceIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetAllianceIdParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetAlliancesParams holds the optional query and header parameters of the request.
// GetAlliancesParams 保存请求的可选查询与头部参数。
type GetAlliancesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetAlliancesParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}
