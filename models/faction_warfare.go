package models

import (
	"time"
)

// FactionWarfareKills Summary of kills done by the given character against enemy factions.
// FactionWarfareKills 指定角色对敌方势力的击杀汇总.
type FactionWarfareKills struct {
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

// CharacterFactionWarfareStats 200 ok object.
// CharacterFactionWarfareStats 200 ok 对象.
type CharacterFactionWarfareStats struct {
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
	Kills FactionWarfareKills `json:"kills"`
	// VictoryPoints Summary of victory points gained by the given character for the enlisted faction.
	// VictoryPoints 指定角色为所属势力获得的胜利点数汇总.
	VictoryPoints FactionWarfareVictoryPoints `json:"victory_points"`
}

// FactionWarfareVictoryPoints Summary of victory points gained by the given character for the enlisted faction.
// FactionWarfareVictoryPoints 指定角色为所属势力获得的胜利点数汇总.
type FactionWarfareVictoryPoints struct {
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

// CorporationFactionWarfareStats 200 ok object.
// CorporationFactionWarfareStats 200 ok 对象.
type CorporationFactionWarfareStats struct {
	// EnlistedOn The enlistment date of the given corporation into faction warfare. Will not be included if corporation is not enlisted in faction warfare.
	// EnlistedOn 指定军团加入势力战争的日期。如果军团未加入势力战争则不包含此项.
	EnlistedOn time.Time `json:"enlisted_on"`
	// FactionId The faction the given corporation is enlisted to fight for. Will not be included if corporation is not enlisted in faction warfare.
	// FactionId 指定军团加入并为之作战的势力。如果军团未加入势力战争则不包含此项.
	FactionId int32 `json:"faction_id"`
	// Kills Summary of kills done by the given corporation against enemy factions.
	// Kills 指定军团对敌方势力的击杀汇总.
	Kills FactionWarfareKills `json:"kills"`
	// Pilots How many pilots the enlisted corporation has. Will not be included if corporation is not enlisted in faction warfare.
	// Pilots 已参战的军团拥有的飞行员数量。如果军团未加入势力战争则不包含此字段.
	Pilots int32 `json:"pilots"`
	// VictoryPoints Summary of victory points gained by the given corporation for the enlisted faction.
	// VictoryPoints 指定军团为所属势力获得的胜利点数汇总.
	VictoryPoints FactionWarfareVictoryPoints `json:"victory_points"`
}

// FactionWarfareLeaderboardEntry active_total object.
// FactionWarfareLeaderboardEntry 当前总计对象.
type FactionWarfareLeaderboardEntry struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}

// FactionWarfareCharacterLeaderboardEntry active_total object.
// FactionWarfareCharacterLeaderboardEntry 当前总计对象.
type FactionWarfareCharacterLeaderboardEntry struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
}

// FactionWarfareCharacterLeaderboardKills Top 100 rankings of pilots by number of kills from yesterday, last week and in total.
// FactionWarfareCharacterLeaderboardKills 飞行员按昨天、上周及总计击杀数排名的前 100 名.
type FactionWarfareCharacterLeaderboardKills struct {
	// ActiveTotal Top 100 ranking of pilots active in faction warfare by total kills. A pilot is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总击杀数排名的势力战争活跃飞行员前 100 名。若飞行员在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []FactionWarfareCharacterLeaderboardEntry `json:"active_total"`
	// LastWeek Top 100 ranking of pilots by kills in the past week.
	// LastWeek 飞行员过去一周击杀数前 100 名.
	LastWeek []FactionWarfareCharacterLeaderboardEntry `json:"last_week"`
	// Yesterday Top 100 ranking of pilots by kills in the past day.
	// Yesterday 飞行员过去一天击杀数前 100 名.
	Yesterday []FactionWarfareCharacterLeaderboardEntry `json:"yesterday"`
}

// FactionWarfareCharacterLeaderboard 200 ok object.
// FactionWarfareCharacterLeaderboard 200 ok 对象.
type FactionWarfareCharacterLeaderboard struct {
	// Kills Top 100 rankings of pilots by number of kills from yesterday, last week and in total.
	// Kills 飞行员按昨天、上周及总计击杀数排名的前 100 名.
	Kills FactionWarfareCharacterLeaderboardKills `json:"kills"`
	// VictoryPoints Top 100 rankings of pilots by victory points from yesterday, last week and in total.
	// VictoryPoints 飞行员按昨天、上周及总计胜利点数排名的前 100 名.
	VictoryPoints FactionWarfareCharacterLeaderboardVictoryPoints `json:"victory_points"`
}

// FactionWarfareCharacterLeaderboardVictoryPoints Top 100 rankings of pilots by victory points from yesterday, last week and in total.
// FactionWarfareCharacterLeaderboardVictoryPoints 飞行员按昨天、上周及总计胜利点数排名的前 100 名.
type FactionWarfareCharacterLeaderboardVictoryPoints struct {
	// ActiveTotal Top 100 ranking of pilots active in faction warfare by total victory points. A pilot is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总胜利点数排名的势力战争活跃飞行员前 100 名。若飞行员在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []FactionWarfareCharacterLeaderboardEntry `json:"active_total"`
	// LastWeek Top 100 ranking of pilots by victory points in the past week.
	// LastWeek 飞行员过去一周胜利点数前 100 名.
	LastWeek []FactionWarfareCharacterLeaderboardEntry `json:"last_week"`
	// Yesterday Top 100 ranking of pilots by victory points in the past day.
	// Yesterday 飞行员过去一天胜利点数前 100 名.
	Yesterday []FactionWarfareCharacterLeaderboardEntry `json:"yesterday"`
}

// FactionWarfareCorporationLeaderboardEntry active_total object.
// FactionWarfareCorporationLeaderboardEntry 当前总计对象.
type FactionWarfareCorporationLeaderboardEntry struct {
	// Amount Amount of kills.
	// Amount 击杀数量.
	Amount int32 `json:"amount"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
}

// FactionWarfareCorporationLeaderboardKills Top 10 rankings of corporations by number of kills from yesterday, last week and in total.
// FactionWarfareCorporationLeaderboardKills 军团按昨天、上周及总计击杀数排名的前 10 名.
type FactionWarfareCorporationLeaderboardKills struct {
	// ActiveTotal Top 10 ranking of corporations active in faction warfare by total kills. A corporation is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总击杀数排名的势力战争活跃军团前 10 名。若军团在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []FactionWarfareCorporationLeaderboardEntry `json:"active_total"`
	// LastWeek Top 10 ranking of corporations by kills in the past week.
	// LastWeek 军团过去一周击杀数前 10 名.
	LastWeek []FactionWarfareCorporationLeaderboardEntry `json:"last_week"`
	// Yesterday Top 10 ranking of corporations by kills in the past day.
	// Yesterday 军团过去一天击杀数前 10 名.
	Yesterday []FactionWarfareCorporationLeaderboardEntry `json:"yesterday"`
}

// FactionWarfareCorporationLeaderboard 200 ok object.
// FactionWarfareCorporationLeaderboard 200 ok 对象.
type FactionWarfareCorporationLeaderboard struct {
	// Kills Top 10 rankings of corporations by number of kills from yesterday, last week and in total.
	// Kills 军团按昨天、上周及总计击杀数排名的前 10 名.
	Kills FactionWarfareCorporationLeaderboardKills `json:"kills"`
	// VictoryPoints Top 10 rankings of corporations by victory points from yesterday, last week and in total.
	// VictoryPoints 军团按昨天、上周及总计胜利点数排名的前 10 名.
	VictoryPoints FactionWarfareCorporationLeaderboardVictoryPoints `json:"victory_points"`
}

// FactionWarfareCorporationLeaderboardVictoryPoints Top 10 rankings of corporations by victory points from yesterday, last week and in total.
// FactionWarfareCorporationLeaderboardVictoryPoints 军团按昨天、上周及总计胜利点数排名的前 10 名.
type FactionWarfareCorporationLeaderboardVictoryPoints struct {
	// ActiveTotal Top 10 ranking of corporations active in faction warfare by total victory points. A corporation is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总胜利点数排名的势力战争活跃军团前 10 名。若军团在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []FactionWarfareCorporationLeaderboardEntry `json:"active_total"`
	// LastWeek Top 10 ranking of corporations by victory points in the past week.
	// LastWeek 军团过去一周胜利点数前 10 名.
	LastWeek []FactionWarfareCorporationLeaderboardEntry `json:"last_week"`
	// Yesterday Top 10 ranking of corporations by victory points in the past day.
	// Yesterday 军团过去一天胜利点数前 10 名.
	Yesterday []FactionWarfareCorporationLeaderboardEntry `json:"yesterday"`
}

// FactionWarfareLeaderboardKills Top 4 rankings of factions by number of kills from yesterday, last week and in total.
// FactionWarfareLeaderboardKills 派系按昨天、上周及总计击杀数排名的前 4 名.
type FactionWarfareLeaderboardKills struct {
	// ActiveTotal Top 4 ranking of factions active in faction warfare by total kills. A faction is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总击杀数排名的势力战争活跃派系前 4 名。若派系在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []FactionWarfareLeaderboardEntry `json:"active_total"`
	// LastWeek Top 4 ranking of factions by kills in the past week.
	// LastWeek 派系过去一周击杀数前 4 名.
	LastWeek []FactionWarfareLeaderboardEntry `json:"last_week"`
	// Yesterday Top 4 ranking of factions by kills in the past day.
	// Yesterday 派系过去一天击杀数前 4 名.
	Yesterday []FactionWarfareLeaderboardEntry `json:"yesterday"`
}

// FactionWarfareLeaderboard 200 ok object.
// FactionWarfareLeaderboard 200 ok 对象.
type FactionWarfareLeaderboard struct {
	// Kills Top 4 rankings of factions by number of kills from yesterday, last week and in total.
	// Kills 派系按昨天、上周及总计击杀数排名的前 4 名.
	Kills FactionWarfareLeaderboardKills `json:"kills"`
	// VictoryPoints Top 4 rankings of factions by victory points from yesterday, last week and in total.
	// VictoryPoints 派系按昨天、上周及总计胜利点数排名的前 4 名.
	VictoryPoints FactionWarfareLeaderboardVictoryPoints `json:"victory_points"`
}

// FactionWarfareLeaderboardVictoryPoints Top 4 rankings of factions by victory points from yesterday, last week and in total.
// FactionWarfareLeaderboardVictoryPoints 派系按昨天、上周及总计胜利点数排名的前 4 名.
type FactionWarfareLeaderboardVictoryPoints struct {
	// ActiveTotal Top 4 ranking of factions active in faction warfare by total victory points. A faction is considered "active" if they have participated in faction warfare in the past 14 days.
	// ActiveTotal 按总胜利点数排名的势力战争活跃派系前 4 名。若派系在过去 14 天内参与过势力战争，则视为“活跃”.
	ActiveTotal []FactionWarfareLeaderboardEntry `json:"active_total"`
	// LastWeek Top 4 ranking of factions by victory points in the past week.
	// LastWeek 派系过去一周胜利点数前 4 名.
	LastWeek []FactionWarfareLeaderboardEntry `json:"last_week"`
	// Yesterday Top 4 ranking of factions by victory points in the past day.
	// Yesterday 派系过去一天胜利点数前 4 名.
	Yesterday []FactionWarfareLeaderboardEntry `json:"yesterday"`
}

// FactionWarfareStats 200 ok object.
// FactionWarfareStats 200 ok 对象.
type FactionWarfareStats struct {
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// Kills Summary of kills against an enemy faction for the given faction.
	// Kills 指定势力对敌方势力的击杀汇总.
	Kills FactionWarfareKills `json:"kills"`
	// Pilots How many pilots fight for the given faction.
	// Pilots 为该势力作战的玩家飞行员数量.
	Pilots int32 `json:"pilots"`
	// SystemsControlled The number of solar systems controlled by the given faction.
	// SystemsControlled 指定势力控制的星系数量.
	SystemsControlled int32 `json:"systems_controlled"`
	// VictoryPoints Summary of victory points gained for the given faction.
	// VictoryPoints 为指定势力获得的胜利点数汇总.
	VictoryPoints FactionWarfareVictoryPoints `json:"victory_points"`
}

// FactionWarfareSystem 200 ok object.
// FactionWarfareSystem 200 ok 对象.
type FactionWarfareSystem struct {
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

// FactionWarfareWar 200 ok object.
// FactionWarfareWar 200 ok 对象.
type FactionWarfareWar struct {
	// AgainstId The faction ID of the enemy faction.
	// AgainstId 敌方势力的势力 ID。
	AgainstId int32 `json:"against_id"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
}
