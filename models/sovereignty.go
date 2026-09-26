package models

import (
	"time"
)

// SovereigntyCampaign 200 ok object.
// SovereigntyCampaign 200 ok 对象.
type SovereigntyCampaign struct {
	// AttackersScore Score for all attacking parties, only present in Defense Events.
	// AttackersScore 所有进攻方的得分，仅出现在防御事件中。
	AttackersScore float64 `json:"attackers_score"`
	// CampaignId Unique ID for this campaign.
	// CampaignId 此战役的唯一 ID。
	CampaignId int32 `json:"campaign_id"`
	// ConstellationId The constellation in which the campaign will take place.
	// ConstellationId 战役将发生的星座。
	ConstellationId int32 `json:"constellation_id"`
	// DefenderId Defending alliance, only present in Defense Events.
	// DefenderId 防守联盟，仅存在于防御事件中.
	DefenderId int32 `json:"defender_id"`
	// DefenderScore Score for the defending alliance, only present in Defense Events.
	// DefenderScore 防守联盟的得分，仅出现在防御事件中。
	DefenderScore float64 `json:"defender_score"`
	// EventType Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as "Defense Events", station_freeport as "Freeport Events".
	// EventType 此战役对应的事件类型。tcu_defense、ihub_defense 和 station_defense 称为“防御事件”，station_freeport 称为“自由港事件”。
	// Enum values: "tcu_defense", "ihub_defense", "station_defense", "station_freeport".
	EventType string `json:"event_type"`
	// Participants Alliance participating and their respective scores, only present in Freeport Events.
	// Participants 参与联盟及其各自得分，仅存在于自由港事件中。
	Participants []SovereigntyCampaignParticipant `json:"participants"`
	// SolarSystemId The solar system the structure is located in.
	// SolarSystemId 建筑所在的星系。
	SolarSystemId int32 `json:"solar_system_id"`
	// StartTime Time the event is scheduled to start.
	// StartTime 事件计划开始的时间。
	StartTime time.Time `json:"start_time"`
	// StructureId The structure item ID that is related to this campaign.
	// StructureId 与此战役相关的建筑物品 ID。
	StructureId int64 `json:"structure_id"`
}

// SovereigntyCampaignParticipant participant object.
// SovereigntyCampaignParticipant participant 对象.
type SovereigntyCampaignParticipant struct {
	// AllianceId alliance_id integer.
	// AllianceId 联盟 ID 整数.
	AllianceId int32 `json:"alliance_id"`
	// Score score number.
	// Score 分数 number.
	Score float64 `json:"score"`
}

// SovereigntySystem 200 ok object.
// SovereigntySystem 200 ok 对象.
type SovereigntySystem struct {
	// AllianceId alliance_id integer.
	// AllianceId 联盟 ID 整数.
	AllianceId int32 `json:"alliance_id"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// SystemId system_id integer.
	// SystemId system_id 整数.
	SystemId int32 `json:"system_id"`
}

// SovereigntyStructure 200 ok object.
// SovereigntyStructure 200 ok 对象.
type SovereigntyStructure struct {
	// AllianceId The alliance that owns the structure.
	// AllianceId 拥有该建筑的联盟。
	AllianceId int32 `json:"alliance_id"`
	// SolarSystemId Solar system in which the structure is located.
	// SolarSystemId 建筑所在的星系。
	SolarSystemId int32 `json:"solar_system_id"`
	// StructureId Unique item ID for this structure.
	// StructureId 此建筑的唯一物品 ID。
	StructureId int64 `json:"structure_id"`
	// StructureTypeId A reference to the type of structure this is.
	// StructureTypeId 对该建筑类型的引用。
	StructureTypeId int32 `json:"structure_type_id"`
	// VulnerabilityOccupancyLevel The occupancy level for the next or current vulnerability window. This takes into account all development indexes and capital system bonuses. Also known as Activity Defense Multiplier from in the client. It increases the time that attackers must spend using their entosis links on the structure.
	// VulnerabilityOccupancyLevel 下一个或当前可攻击窗口的占领等级。它会综合所有发展指数和首都星系加成计算。客户端中亦称 Activity Defense Multiplier。它会增加进攻方必须使用侵入链接作用于建筑的时间。
	VulnerabilityOccupancyLevel float64 `json:"vulnerability_occupancy_level"`
	// VulnerableEndTime The time at which the next or current vulnerability window ends. At the end of a vulnerability window the next window is recalculated and locked in along with the vulnerabilityOccupancyLevel. If the structure is not in 100% entosis control of the defender, it will go in to 'overtime' and stay vulnerable for as long as that situation persists. Only once the defenders have 100% entosis control and has the vulnerableEndTime passed does the vulnerability interval expire and a new one is calculated.
	// VulnerableEndTime 下一个或当前易受攻击窗口结束的时间。在易受攻击窗口结束时，下一个窗口会与 vulnerabilityOccupancyLevel 一起重新计算并锁定。如果建筑未被防守方 100% 恩索控制，则进入“加时”状态，并在该情况持续期间一直保持易受攻击。只有当防守方取得 100% 恩索控制且 vulnerableEndTime 已过，易受攻击周期才会结束并计算新的周期。
	VulnerableEndTime time.Time `json:"vulnerable_end_time"`
	// VulnerableStartTime The next time at which the structure will become vulnerable. Or the start time of the current window if current time is between this and vulnerableEndTime.
	// VulnerableStartTime 建筑下一次进入可攻击状态的时间。如果当前时间介于该时间与 vulnerableEndTime 之间，则为当前窗口的开始时间。
	VulnerableStartTime time.Time `json:"vulnerable_start_time"`
}
