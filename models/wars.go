package models

import "time"

// WarAggressor The aggressor corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
// WarAggressor 宣战的进攻方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
type WarAggressor struct {
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

// WarAlly ally object.
// WarAlly 盟友对象.
type WarAlly struct {
	// AllianceId Alliance ID if and only if this ally is an alliance.
	// AllianceId 当且仅当该盟友为联盟时的联盟 ID.
	AllianceId int32 `json:"alliance_id"`
	// CorporationId Corporation ID if and only if this ally is a corporation.
	// CorporationId 当且仅当该盟友为军团时的军团 ID.
	CorporationId int32 `json:"corporation_id"`
}

// WarDefender The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
// WarDefender 宣战的防守方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
type WarDefender struct {
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

// War 200 ok object.
// War 200 ok 对象.
type War struct {
	// Aggressor The aggressor corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
	// Aggressor 宣战的进攻方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
	Aggressor WarAggressor `json:"aggressor"`
	// Allies allied corporations or alliances, each object contains either corporation_id or alliance_id.
	// Allies 盟军军团或联盟，每个对象包含 corporation_id 或 alliance_id 之一.
	Allies []WarAlly `json:"allies"`
	// Declared Time that the war was declared.
	// Declared 宣战的时间.
	Declared time.Time `json:"declared"`
	// Defender The defending corporation or alliance that declared this war, only contains either corporation_id or alliance_id.
	// Defender 宣战的防守方军团或联盟，仅包含 corporation_id 或 alliance_id 之一.
	Defender WarDefender `json:"defender"`
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
