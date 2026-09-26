package models

// Incursion 200 ok object.
// Incursion 200 ok 对象.
type Incursion struct {
	// ConstellationId The constellation id in which this incursion takes place.
	// ConstellationId 此次入侵活动发生的星座 ID.
	ConstellationId int32 `json:"constellation_id"`
	// FactionId The attacking faction's id.
	// FactionId 进攻方的势力 ID.
	FactionId int32 `json:"faction_id"`
	// HasBoss Whether the final encounter has boss or not.
	// HasBoss 最终遭遇战是否包含 boss.
	HasBoss bool `json:"has_boss"`
	// InfestedSolarSystems A list of infested solar system ids that are a part of this incursion.
	// InfestedSolarSystems 作为本次入侵活动一部分的被感染星系 ID 列表.
	InfestedSolarSystems []int32 `json:"infested_solar_systems"`
	// Influence Influence of this incursion as a float from 0 to 1.
	// Influence 此入侵活动的影响度，为 0 到 1 之间的浮点数.
	Influence float64 `json:"influence"`
	// StagingSolarSystemId Staging solar system for this incursion.
	// StagingSolarSystemId 本次入侵活动的集结星系.
	StagingSolarSystemId int32 `json:"staging_solar_system_id"`
	// State The state of this incursion.
	// State 此次入侵活动的状态.
	// Enum values: "withdrawing", "mobilizing", "established".
	State string `json:"state"`
	// TypeValue The type of this incursion.
	// TypeValue 此次入侵活动的类型.
	TypeValue string `json:"type"`
}
