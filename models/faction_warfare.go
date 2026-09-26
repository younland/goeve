package models

import (
	"net/url"
	"time"
)

// GetCharactersCharacterIdFwStatsKills Summary of kills done by the given character against enemy factions.
// GetCharactersCharacterIdFwStatsKills 指定角色对敌方势力的击杀汇总.
type GetCharactersCharacterIdFwStatsKills struct {
	// LastWeek Last week's total number of kills by a given character against enemy factions.
	// LastWeek 上周指定角色对敌势力击杀的总数.
	LastWeek int32 `json:"last_week"`
	// Total Total number of kills by a given character against enemy factions since the character enlisted.
	// Total 自指定角色入伍以来其对敌方派系的总击杀数.
	Total int32 `json:"total"`
	// Yesterday Yesterday's total number of kills by a given character against enemy factions.
	// Yesterday 昨日指定角色对敌方势力的总击杀数.
	Yesterday int32 `json:"yesterday"`
}

// GetCharactersCharacterIdFwStats 200 ok object.
// GetCharactersCharacterIdFwStats 200 ok 对象.
type GetCharactersCharacterIdFwStats struct {
	// CurrentRank The given character's current faction rank.
	// CurrentRank 指定角色当前的势力军衔.
	CurrentRank int32 `json:"current_rank"`
	// EnlistedOn The enlistment date of the given character into faction warfare. Will not be included if character is not enlisted in faction warfare.
	// EnlistedOn 指定角色加入势力战争的日期。如果角色未加入势力战争则不包含此项.
	EnlistedOn time.Time `json:"enlisted_on"`
	// FactionId The faction the given character is enlisted to fight for. Will not be included if character is not enlisted in faction warfare.
	// FactionId 指定角色加入并为之作战的势力。如果角色未加入势力战争则不包含此项.
	FactionId int32 `json:"faction_id"`
	// HighestRank The given character's highest faction rank achieved.
	// HighestRank 指定角色达到过的最高势力军衔.
	HighestRank int32 `json:"highest_rank"`
	// Kills Summary of kills done by the given character against enemy factions.
	// Kills 指定角色对敌方势力的击杀汇总.
	Kills GetCharactersCharacterIdFwStatsKills `json:"kills"`
	// VictoryPoints Summary of victory points gained by the given character for the enlisted faction.
	// VictoryPoints 指定角色为所属势力获得的胜利点数汇总.
	VictoryPoints GetCharactersCharacterIdFwStatsVictoryPoints `json:"victory_points"`
}

// GetCharactersCharacterIdFwStatsVictoryPoints Summary of victory points gained by the given character for the enlisted faction.
// GetCharactersCharacterIdFwStatsVictoryPoints 指定角色为所属势力获得的胜利点数汇总.
type GetCharactersCharacterIdFwStatsVictoryPoints struct {
	// LastWeek Last week's victory points gained by the given character.
	// LastWeek 上周指定角色获得的胜利点数.
	LastWeek int32 `json:"last_week"`
	// Total Total victory points gained since the given character enlisted.
	// Total 自指定角色入伍以来获得的胜利点总数.
	Total int32 `json:"total"`
	// Yesterday Yesterday's victory points gained by the given character.
	// Yesterday 昨日指定角色获得的胜利点数.
	Yesterday int32 `json:"yesterday"`
}

// GetCharactersCharacterIdFwStatsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdFwStatsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdFwStatsParams struct {
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

func (p *GetCharactersCharacterIdFwStatsParams) Values() (url.Values, map[string]string) {
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

// GetCorporationsCorporationIdFwStatsKills Summary of kills done by the given corporation against enemy factions.
// GetCorporationsCorporationIdFwStatsKills 指定军团对敌方势力的击杀汇总.
type GetCorporationsCorporationIdFwStatsKills struct {
	// LastWeek Last week's total number of kills by members of the given corporation against enemy factions.
	// LastWeek 上周指定军团成员对敌势力击杀的总数.
	LastWeek int32 `json:"last_week"`
	// Total Total number of kills by members of the given corporation against enemy factions since the corporation enlisted.
	// Total 自指定军团入伍以来其成员对敌方派系的总击杀数.
	Total int32 `json:"total"`
	// Yesterday Yesterday's total number of kills by members of the given corporation against enemy factions.
	// Yesterday 昨日指定军团成员对敌方势力的总击杀数.
	Yesterday int32 `json:"yesterday"`
}

// GetCorporationsCorporationIdFwStats 200 ok object.
// GetCorporationsCorporationIdFwStats 200 ok 对象.
type GetCorporationsCorporationIdFwStats struct {
	// EnlistedOn The enlistment date of the given corporation into faction warfare. Will not be included if corporation is not enlisted in faction warfare.
	// EnlistedOn 指定军团加入势力战争的日期。如果军团未加入势力战争则不包含此项.
	EnlistedOn time.Time `json:"enlisted_on"`
	// FactionId The faction the given corporation is enlisted to fight for. Will not be included if corporation is not enlisted in faction warfare.
	// FactionId 指定军团加入并为之作战的势力。如果军团未加入势力战争则不包含此项.
	FactionId int32 `json:"faction_id"`
	// Kills Summary of kills done by the given corporation against enemy factions.
	// Kills 指定军团对敌方势力的击杀汇总.
	Kills GetCorporationsCorporationIdFwStatsKills `json:"kills"`
	// Pilots How many pilots the enlisted corporation has. Will not be included if corporation is not enlisted in faction warfare.
	// Pilots 已参战的军团拥有的飞行员数量。如果军团未加入势力战争则不包含此字段.
	Pilots int32 `json:"pilots"`
	// VictoryPoints Summary of victory points gained by the given corporation for the enlisted faction.
	// VictoryPoints 指定军团为所属势力获得的胜利点数汇总.
	VictoryPoints GetCorporationsCorporationIdFwStatsVictoryPoints `json:"victory_points"`
}

// GetCorporationsCorporationIdFwStatsVictoryPoints Summary of victory points gained by the given corporation for the enlisted faction.
// GetCorporationsCorporationIdFwStatsVictoryPoints 指定军团为所属势力获得的胜利点数汇总.
type GetCorporationsCorporationIdFwStatsVictoryPoints struct {
	// LastWeek Last week's victory points gained by members of the given corporation.
	// LastWeek 上周指定军团成员获得的胜利点数.
	LastWeek int32 `json:"last_week"`
	// Total Total victory points gained since the given corporation enlisted.
	// Total 自指定军团入伍以来获得的胜利点总数.
	Total int32 `json:"total"`
	// Yesterday Yesterday's victory points gained by members of the given corporation.
	// Yesterday 昨日指定军团成员获得的胜利点数.
	Yesterday int32 `json:"yesterday"`
}

// GetCorporationsCorporationIdFwStatsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdFwStatsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdFwStatsParams struct {
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

func (p *GetCorporationsCorporationIdFwStatsParams) Values() (url.Values, map[string]string) {
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

// GetFwLeaderboardsActiveTotalActiveTotal active_total object.
// GetFwLeaderboardsActiveTotalActiveTotal 当前总计对象.
type GetFwLeaderboardsActiveTotalActiveTotal struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// GetFwLeaderboardsActiveTotalActiveTotal1 active_total object.
// GetFwLeaderboardsActiveTotalActiveTotal1 当前总计对象.
type GetFwLeaderboardsActiveTotalActiveTotal1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// GetFwLeaderboardsCharactersActiveTotalActiveTotal active_total object.
// GetFwLeaderboardsCharactersActiveTotalActiveTotal 当前总计对象.
type GetFwLeaderboardsCharactersActiveTotalActiveTotal struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
}

// GetFwLeaderboardsCharactersActiveTotalActiveTotal1 active_total object.
// GetFwLeaderboardsCharactersActiveTotalActiveTotal1 当前总计对象.
type GetFwLeaderboardsCharactersActiveTotalActiveTotal1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
}

// GetFwLeaderboardsCharactersKills Top 100 rankings of pilots by number of kills from yesterday, last week and in total.
// GetFwLeaderboardsCharactersKills 飞行员按昨天、上周及总计击杀数排名的前 100 名.
type GetFwLeaderboardsCharactersKills struct {
	// ActiveTotal Top 100 ranking of pilots active in faction warfare by total kills. A pilot is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总击杀数排名的势力战争活跃飞行员前 100 名。若飞行员在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []GetFwLeaderboardsCharactersActiveTotalActiveTotal `json:"active_total"`
	// LastWeek Top 100 ranking of pilots by kills in the past week.
	// LastWeek 飞行员过去一周击杀数前 100 名.
	LastWeek []GetFwLeaderboardsCharactersLastWeekLastWeek `json:"last_week"`
	// Yesterday Top 100 ranking of pilots by kills in the past day.
	// Yesterday 飞行员过去一天击杀数前 100 名.
	Yesterday []GetFwLeaderboardsCharactersYesterdayYesterday `json:"yesterday"`
}

// GetFwLeaderboardsCharactersLastWeekLastWeek last_week object.
// GetFwLeaderboardsCharactersLastWeekLastWeek last_week 对象.
type GetFwLeaderboardsCharactersLastWeekLastWeek struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
}

// GetFwLeaderboardsCharactersLastWeekLastWeek1 last_week object.
// GetFwLeaderboardsCharactersLastWeekLastWeek1 last_week 对象.
type GetFwLeaderboardsCharactersLastWeekLastWeek1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
}

// GetFwLeaderboardsCharacters 200 ok object.
// GetFwLeaderboardsCharacters 200 ok 对象.
type GetFwLeaderboardsCharacters struct {
	// Kills Top 100 rankings of pilots by number of kills from yesterday, last week and in total.
	// Kills 飞行员按昨天、上周及总计击杀数排名的前 100 名.
	Kills GetFwLeaderboardsCharactersKills `json:"kills"`
	// VictoryPoints Top 100 rankings of pilots by victory points from yesterday, last week and in total.
	// VictoryPoints 飞行员按昨天、上周及总计胜利点数排名的前 100 名.
	VictoryPoints GetFwLeaderboardsCharactersVictoryPoints `json:"victory_points"`
}

// GetFwLeaderboardsCharactersVictoryPoints Top 100 rankings of pilots by victory points from yesterday, last week and in total.
// GetFwLeaderboardsCharactersVictoryPoints 飞行员按昨天、上周及总计胜利点数排名的前 100 名.
type GetFwLeaderboardsCharactersVictoryPoints struct {
	// ActiveTotal Top 100 ranking of pilots active in faction warfare by total victory points. A pilot is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总胜利点数排名的势力战争活跃飞行员前 100 名。若飞行员在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []GetFwLeaderboardsCharactersActiveTotalActiveTotal1 `json:"active_total"`
	// LastWeek Top 100 ranking of pilots by victory points in the past week.
	// LastWeek 飞行员过去一周胜利点数前 100 名.
	LastWeek []GetFwLeaderboardsCharactersLastWeekLastWeek1 `json:"last_week"`
	// Yesterday Top 100 ranking of pilots by victory points in the past day.
	// Yesterday 飞行员过去一天胜利点数前 100 名.
	Yesterday []GetFwLeaderboardsCharactersYesterdayYesterday1 `json:"yesterday"`
}

// GetFwLeaderboardsCharactersYesterdayYesterday yesterday object.
// GetFwLeaderboardsCharactersYesterdayYesterday yesterday 对象.
type GetFwLeaderboardsCharactersYesterdayYesterday struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
}

// GetFwLeaderboardsCharactersYesterdayYesterday1 yesterday object.
// GetFwLeaderboardsCharactersYesterdayYesterday1 yesterday 对象.
type GetFwLeaderboardsCharactersYesterdayYesterday1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
}

// GetFwLeaderboardsCharactersParams holds the optional query and header parameters of the request.
// GetFwLeaderboardsCharactersParams 保存请求的可选查询与头部参数。
type GetFwLeaderboardsCharactersParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetFwLeaderboardsCharactersParams) Values() (url.Values, map[string]string) {
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

// GetFwLeaderboardsCorporationsActiveTotalActiveTotal active_total object.
// GetFwLeaderboardsCorporationsActiveTotalActiveTotal 当前总计对象.
type GetFwLeaderboardsCorporationsActiveTotalActiveTotal struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
}

// GetFwLeaderboardsCorporationsActiveTotalActiveTotal1 active_total object.
// GetFwLeaderboardsCorporationsActiveTotalActiveTotal1 当前总计对象.
type GetFwLeaderboardsCorporationsActiveTotalActiveTotal1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
}

// GetFwLeaderboardsCorporationsKills Top 10 rankings of corporations by number of kills from yesterday, last week and in total.
// GetFwLeaderboardsCorporationsKills 军团按昨天、上周及总计击杀数排名的前 10 名.
type GetFwLeaderboardsCorporationsKills struct {
	// ActiveTotal Top 10 ranking of corporations active in faction warfare by total kills. A corporation is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总击杀数排名的势力战争活跃军团前 10 名。若军团在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []GetFwLeaderboardsCorporationsActiveTotalActiveTotal `json:"active_total"`
	// LastWeek Top 10 ranking of corporations by kills in the past week.
	// LastWeek 军团过去一周击杀数前 10 名.
	LastWeek []GetFwLeaderboardsCorporationsLastWeekLastWeek `json:"last_week"`
	// Yesterday Top 10 ranking of corporations by kills in the past day.
	// Yesterday 军团过去一天击杀数前 10 名.
	Yesterday []GetFwLeaderboardsCorporationsYesterdayYesterday `json:"yesterday"`
}

// GetFwLeaderboardsCorporationsLastWeekLastWeek last_week object.
// GetFwLeaderboardsCorporationsLastWeekLastWeek last_week 对象.
type GetFwLeaderboardsCorporationsLastWeekLastWeek struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
}

// GetFwLeaderboardsCorporationsLastWeekLastWeek1 last_week object.
// GetFwLeaderboardsCorporationsLastWeekLastWeek1 last_week 对象.
type GetFwLeaderboardsCorporationsLastWeekLastWeek1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
}

// GetFwLeaderboardsCorporations 200 ok object.
// GetFwLeaderboardsCorporations 200 ok 对象.
type GetFwLeaderboardsCorporations struct {
	// Kills Top 10 rankings of corporations by number of kills from yesterday, last week and in total.
	// Kills 军团按昨天、上周及总计击杀数排名的前 10 名.
	Kills GetFwLeaderboardsCorporationsKills `json:"kills"`
	// VictoryPoints Top 10 rankings of corporations by victory points from yesterday, last week and in total.
	// VictoryPoints 军团按昨天、上周及总计胜利点数排名的前 10 名.
	VictoryPoints GetFwLeaderboardsCorporationsVictoryPoints `json:"victory_points"`
}

// GetFwLeaderboardsCorporationsVictoryPoints Top 10 rankings of corporations by victory points from yesterday, last week and in total.
// GetFwLeaderboardsCorporationsVictoryPoints 军团按昨天、上周及总计胜利点数排名的前 10 名.
type GetFwLeaderboardsCorporationsVictoryPoints struct {
	// ActiveTotal Top 10 ranking of corporations active in faction warfare by total victory points. A corporation is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总胜利点数排名的势力战争活跃军团前 10 名。若军团在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []GetFwLeaderboardsCorporationsActiveTotalActiveTotal1 `json:"active_total"`
	// LastWeek Top 10 ranking of corporations by victory points in the past week.
	// LastWeek 军团过去一周胜利点数前 10 名.
	LastWeek []GetFwLeaderboardsCorporationsLastWeekLastWeek1 `json:"last_week"`
	// Yesterday Top 10 ranking of corporations by victory points in the past day.
	// Yesterday 军团过去一天胜利点数前 10 名.
	Yesterday []GetFwLeaderboardsCorporationsYesterdayYesterday1 `json:"yesterday"`
}

// GetFwLeaderboardsCorporationsYesterdayYesterday yesterday object.
// GetFwLeaderboardsCorporationsYesterdayYesterday yesterday 对象.
type GetFwLeaderboardsCorporationsYesterdayYesterday struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
}

// GetFwLeaderboardsCorporationsYesterdayYesterday1 yesterday object.
// GetFwLeaderboardsCorporationsYesterdayYesterday1 yesterday 对象.
type GetFwLeaderboardsCorporationsYesterdayYesterday1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
}

// GetFwLeaderboardsCorporationsParams holds the optional query and header parameters of the request.
// GetFwLeaderboardsCorporationsParams 保存请求的可选查询与头部参数。
type GetFwLeaderboardsCorporationsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetFwLeaderboardsCorporationsParams) Values() (url.Values, map[string]string) {
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

// GetFwLeaderboardsKills Top 4 rankings of factions by number of kills from yesterday, last week and in total.
// GetFwLeaderboardsKills 派系按昨天、上周及总计击杀数排名的前 4 名.
type GetFwLeaderboardsKills struct {
	// ActiveTotal Top 4 ranking of factions active in faction warfare by total kills. A faction is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总击杀数排名的势力战争活跃派系前 4 名。若派系在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []GetFwLeaderboardsActiveTotalActiveTotal `json:"active_total"`
	// LastWeek Top 4 ranking of factions by kills in the past week.
	// LastWeek 派系过去一周击杀数前 4 名.
	LastWeek []GetFwLeaderboardsLastWeekLastWeek `json:"last_week"`
	// Yesterday Top 4 ranking of factions by kills in the past day.
	// Yesterday 派系过去一天击杀数前 4 名.
	Yesterday []GetFwLeaderboardsYesterdayYesterday `json:"yesterday"`
}

// GetFwLeaderboardsLastWeekLastWeek last_week object.
// GetFwLeaderboardsLastWeekLastWeek last_week 对象.
type GetFwLeaderboardsLastWeekLastWeek struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// GetFwLeaderboardsLastWeekLastWeek1 last_week object.
// GetFwLeaderboardsLastWeekLastWeek1 last_week 对象.
type GetFwLeaderboardsLastWeekLastWeek1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// GetFwLeaderboards 200 ok object.
// GetFwLeaderboards 200 ok 对象.
type GetFwLeaderboards struct {
	// Kills Top 4 rankings of factions by number of kills from yesterday, last week and in total.
	// Kills 派系按昨天、上周及总计击杀数排名的前 4 名.
	Kills GetFwLeaderboardsKills `json:"kills"`
	// VictoryPoints Top 4 rankings of factions by victory points from yesterday, last week and in total.
	// VictoryPoints 派系按昨天、上周及总计胜利点数排名的前 4 名.
	VictoryPoints GetFwLeaderboardsVictoryPoints `json:"victory_points"`
}

// GetFwLeaderboardsVictoryPoints Top 4 rankings of factions by victory points from yesterday, last week and in total.
// GetFwLeaderboardsVictoryPoints 派系按昨天、上周及总计胜利点数排名的前 4 名.
type GetFwLeaderboardsVictoryPoints struct {
	// ActiveTotal Top 4 ranking of factions active in faction warfare by total victory points. A faction is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总胜利点数排名的势力战争活跃派系前 4 名。若派系在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []GetFwLeaderboardsActiveTotalActiveTotal1 `json:"active_total"`
	// LastWeek Top 4 ranking of factions by victory points in the past week.
	// LastWeek 派系过去一周胜利点数前 4 名.
	LastWeek []GetFwLeaderboardsLastWeekLastWeek1 `json:"last_week"`
	// Yesterday Top 4 ranking of factions by victory points in the past day.
	// Yesterday 派系过去一天胜利点数前 4 名.
	Yesterday []GetFwLeaderboardsYesterdayYesterday1 `json:"yesterday"`
}

// GetFwLeaderboardsYesterdayYesterday yesterday object.
// GetFwLeaderboardsYesterdayYesterday yesterday 对象.
type GetFwLeaderboardsYesterdayYesterday struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// GetFwLeaderboardsYesterdayYesterday1 yesterday object.
// GetFwLeaderboardsYesterdayYesterday1 yesterday 对象.
type GetFwLeaderboardsYesterdayYesterday1 struct {
	// Amount Amount of victory points.
	// Amount 胜利点数.
	Amount int32 `json:"amount"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// GetFwLeaderboardsParams holds the optional query and header parameters of the request.
// GetFwLeaderboardsParams 保存请求的可选查询与头部参数。
type GetFwLeaderboardsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetFwLeaderboardsParams) Values() (url.Values, map[string]string) {
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

// GetFwStats 200 ok object.
// GetFwStats 200 ok 对象.
type GetFwStats struct {
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// Kills Summary of kills against an enemy faction for the given faction.
	// Kills 指定势力对敌方势力的击杀汇总.
	Kills GetFwStatsKills `json:"kills"`
	// Pilots How many pilots fight for the given faction.
	// Pilots 为该势力作战的玩家飞行员数量.
	Pilots int32 `json:"pilots"`
	// SystemsControlled The number of solar systems controlled by the given faction.
	// SystemsControlled 指定势力控制的星系数量.
	SystemsControlled int32 `json:"systems_controlled"`
	// VictoryPoints Summary of victory points gained for the given faction.
	// VictoryPoints 为指定势力获得的胜利点数汇总.
	VictoryPoints GetFwStatsVictoryPoints `json:"victory_points"`
}

// GetFwStatsKills Summary of kills against an enemy faction for the given faction.
// GetFwStatsKills 指定势力对敌方势力的击杀汇总.
type GetFwStatsKills struct {
	// LastWeek Last week's total number of kills against enemy factions.
	// LastWeek 上周对敌势力击杀的总数.
	LastWeek int32 `json:"last_week"`
	// Total Total number of kills against enemy factions since faction warfare began.
	// Total 自势力战争开始以来对敌方派系的总击杀数.
	Total int32 `json:"total"`
	// Yesterday Yesterday's total number of kills against enemy factions.
	// Yesterday 昨日对敌方势力的总击杀数.
	Yesterday int32 `json:"yesterday"`
}

// GetFwStatsVictoryPoints Summary of victory points gained for the given faction.
// GetFwStatsVictoryPoints 为指定势力获得的胜利点数汇总.
type GetFwStatsVictoryPoints struct {
	// LastWeek Last week's victory points gained.
	// LastWeek 上周获得的胜利点数.
	LastWeek int32 `json:"last_week"`
	// Total Total victory points gained since faction warfare began.
	// Total 自势力战争开始以来获得的胜利点总数.
	Total int32 `json:"total"`
	// Yesterday Yesterday's victory points gained.
	// Yesterday 昨日获得的胜利点数.
	Yesterday int32 `json:"yesterday"`
}

// GetFwStatsParams holds the optional query and header parameters of the request.
// GetFwStatsParams 保存请求的可选查询与头部参数。
type GetFwStatsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetFwStatsParams) Values() (url.Values, map[string]string) {
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

// GetFwSystems 200 ok object.
// GetFwSystems 200 ok 对象.
type GetFwSystems struct {
	// Contested contested string.
	// Contested 争夺状态字符串.
	// Enum values: "captured", "contested", "uncontested", "vulnerable".
	Contested string `json:"contested"`
	// OccupierFactionId occupier_faction_id integer.
	// OccupierFactionId occupier_faction_id 整数.
	OccupierFactionId int32 `json:"occupier_faction_id"`
	// OwnerFactionId owner_faction_id integer.
	// OwnerFactionId owner_faction_id 整数.
	OwnerFactionId int32 `json:"owner_faction_id"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// VictoryPoints victory_points integer.
	// VictoryPoints 胜利点数整数.
	VictoryPoints int32 `json:"victory_points"`
	// VictoryPointsThreshold victory_points_threshold integer.
	// VictoryPointsThreshold 胜利点数阈值整数.
	VictoryPointsThreshold int32 `json:"victory_points_threshold"`
}

// GetFwSystemsParams holds the optional query and header parameters of the request.
// GetFwSystemsParams 保存请求的可选查询与头部参数。
type GetFwSystemsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetFwSystemsParams) Values() (url.Values, map[string]string) {
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

// GetFwWars 200 ok object.
// GetFwWars 200 ok 对象.
type GetFwWars struct {
	// AgainstId The faction ID of the enemy faction.
	// AgainstId 敌方势力的势力 ID。
	AgainstId int32 `json:"against_id"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// GetFwWarsParams holds the optional query and header parameters of the request.
// GetFwWarsParams 保存请求的可选查询与头部参数。
type GetFwWarsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetFwWarsParams) Values() (url.Values, map[string]string) {
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
