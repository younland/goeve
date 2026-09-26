package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetWarsWarIdAggressor The aggressor corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
// GetWarsWarIdAggressor 宣战的进攻方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
type GetWarsWarIdAggressor struct {
	// AllianceId Alliance ID if and only if the aggressor is an alliance.
	// AllianceId 当且仅当攻击方为联盟时的联盟 ID.
	AllianceId int32 `json:"alliance_id"`
	// CorporationId Corporation ID if and only if the aggressor is a corporation.
	// CorporationId 当且仅当攻击方为军团时的军团 ID.
	CorporationId int32 `json:"corporation_id"`
	// IskDestroyed ISK value of ships the aggressor has destroyed.
	// IskDestroyed 进攻方摧毁舰船的 ISK 价值.
	IskDestroyed float64 `json:"isk_destroyed"`
	// ShipsKilled The number of ships the aggressor has killed.
	// ShipsKilled 进攻方击毁的舰船数量.
	ShipsKilled int32 `json:"ships_killed"`
}

// GetWarsWarIdAlly ally object.
// GetWarsWarIdAlly 盟友对象.
type GetWarsWarIdAlly struct {
	// AllianceId Alliance ID if and only if this ally is an alliance.
	// AllianceId 当且仅当该盟友为联盟时的联盟 ID.
	AllianceId int32 `json:"alliance_id"`
	// CorporationId Corporation ID if and only if this ally is a corporation.
	// CorporationId 当且仅当该盟友为军团时的军团 ID.
	CorporationId int32 `json:"corporation_id"`
}

// GetWarsWarIdDefender The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
// GetWarsWarIdDefender 宣战的防守方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
type GetWarsWarIdDefender struct {
	// AllianceId Alliance ID if and only if the defender is an alliance.
	// AllianceId 当且仅当防守方为联盟时的联盟 ID.
	AllianceId int32 `json:"alliance_id"`
	// CorporationId Corporation ID if and only if the defender is a corporation.
	// CorporationId 当且仅当防守方为军团时的军团 ID.
	CorporationId int32 `json:"corporation_id"`
	// IskDestroyed ISK value of ships the defender has killed.
	// IskDestroyed 防守方击杀舰船的 ISK 价值.
	IskDestroyed float64 `json:"isk_destroyed"`
	// ShipsKilled The number of ships the defender has killed.
	// ShipsKilled 防守方击毁的舰船数量.
	ShipsKilled int32 `json:"ships_killed"`
}

// GetWarsWarIdKillmails 200 ok object.
// GetWarsWarIdKillmails 200 ok 对象.
type GetWarsWarIdKillmails struct {
	// KillmailHash A hash of this killmail.
	// KillmailHash 该击杀报告的哈希值.
	KillmailHash string `json:"killmail_hash"`
	// KillmailId ID of this killmail.
	// KillmailId 此击杀报告的 ID.
	KillmailId int32 `json:"killmail_id"`
}

// GetWarsWarIdOk 200 ok object.
// GetWarsWarIdOk 200 ok 对象.
type GetWarsWarIdOk struct {
	// Aggressor The aggressor corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
	// Aggressor 宣战的进攻方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
	Aggressor GetWarsWarIdAggressor `json:"aggressor"`
	// Allies allied corporations or alliances, each object contains either corporation_id or alliance_id.
	// Allies 盟军军团或联盟，每个对象包含 corporation_id 或 alliance_id 之一.
	Allies []GetWarsWarIdAlly `json:"allies"`
	// Declared Time that the war was declared.
	// Declared 宣战的时间.
	Declared time.Time `json:"declared"`
	// Defender The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
	// Defender 宣战的防守方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
	Defender GetWarsWarIdDefender `json:"defender"`
	// Finished Time the war ended and shooting was no longer allowed.
	// Finished 战争结束、不再允许交火的时间.
	Finished time.Time `json:"finished"`
	// Id ID of the specified war.
	// Id 指定战争的 ID.
	Id int32 `json:"id"`
	// Mutual Was the war declared mutual by both parties.
	// Mutual 战争是否由双方共同宣战.
	Mutual bool `json:"mutual"`
	// OpenForAllies Is the war currently open for allies or not.
	// OpenForAllies 此战争当前是否对盟友开放.
	OpenForAllies bool `json:"open_for_allies"`
	// Retracted Time the war was retracted but both sides could still shoot each other.
	// Retracted 战争被撤回但双方仍可相互交火的时间.
	Retracted time.Time `json:"retracted"`
	// Started Time when the war started and both sides could shoot each other.
	// Started 战争开始、双方可以相互交火的时间.
	Started time.Time `json:"started"`
}

// GetWarIdKillmailsParams holds the optional query and header parameters of the request.
// GetWarIdKillmailsParams 保存请求的可选查询与头部参数。
type GetWarIdKillmailsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
}

func (p *GetWarIdKillmailsParams) Values() (url.Values, map[string]string) {
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
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	return query, headers
}

// GetWarIdParams holds the optional query and header parameters of the request.
// GetWarIdParams 保存请求的可选查询与头部参数。
type GetWarIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetWarIdParams) Values() (url.Values, map[string]string) {
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

// GetWarsParams holds the optional query and header parameters of the request.
// GetWarsParams 保存请求的可选查询与头部参数。
type GetWarsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// MaxWarId Only return wars with ID smaller than this.
	// MaxWarId 仅返回 ID 小于此值的战争.
	MaxWarId *int32
}

func (p *GetWarsParams) Values() (url.Values, map[string]string) {
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
	if p.MaxWarId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("max_war_id", strconv.FormatInt(int64(*p.MaxWarId), 10))
	}
	return query, headers
}
