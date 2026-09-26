package models

import (
	"time"
)

// AllianceIcons 200 ok object.
// AllianceIcons 200 ok 对象.
type AllianceIcons struct {
	// Px128x128 px128x128 string.
	// Px128x128 px128x128 字符串.
	Px128x128 string `json:"px128x128"`
	// Px64x64 px64x64 string.
	// Px64x64 px64x64 字符串.
	Px64x64 string `json:"px64x64"`
}

// Alliance 200 ok object.
// Alliance 200 ok 对象.
type Alliance struct {
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
