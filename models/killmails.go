package models

import (
	"time"
)

// KillmailRef 200 ok object.
// KillmailRef 200 ok 对象.
type KillmailRef struct {
	// KillmailHash A hash of this killmail.
	// KillmailHash 该击杀报告的哈希值.
	KillmailHash string `json:"killmail_hash"`
	// KillmailId ID of this killmail.
	// KillmailId 此击杀报告的 ID.
	KillmailId int32 `json:"killmail_id"`
}

// KillmailAttacker attacker object.
// KillmailAttacker 攻击者对象.
type KillmailAttacker struct {
	// AllianceId alliance_id integer.
	// AllianceId 联盟 ID 整数.
	AllianceId int32 `json:"alliance_id"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// DamageDone damage_done integer.
	// DamageDone 造成的伤害整数.
	DamageDone int32 `json:"damage_done"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// FinalBlow Was the attacker the one to achieve the final blow.
	// FinalBlow 攻击者是否完成了最后一击.
	FinalBlow bool `json:"final_blow"`
	// SecurityStatus Security status for the attacker.
	// SecurityStatus 攻击者的安全等级.
	SecurityStatus float64 `json:"security_status"`
	// ShipTypeId What ship was the attacker flying.
	// ShipTypeId 攻击者驾驶的舰船.
	ShipTypeId int32 `json:"ship_type_id"`
	// WeaponTypeId What weapon was used by the attacker for the kill.
	// WeaponTypeId 攻击者用于击杀的武器.
	WeaponTypeId int32 `json:"weapon_type_id"`
}

// KillmailItem item object.
// KillmailItem 物品对象.
type KillmailItem struct {
	// Flag Flag for the location of the item.
	// Flag 表示物品位置的标志.
	Flag int32 `json:"flag"`
	// ItemTypeId item_type_id integer.
	// ItemTypeId item_type_id 整数.
	ItemTypeId int32 `json:"item_type_id"`
	// Items items array.
	// Items items 数组.
	Items []KillmailContainedItem `json:"items"`
	// QuantityDestroyed How many of the item were destroyed if any.
	// QuantityDestroyed 该物品被摧毁的数量（如有）
	QuantityDestroyed int64 `json:"quantity_destroyed"`
	// QuantityDropped How many of the item were dropped if any.
	// QuantityDropped 该物品掉落的数量（如有）
	QuantityDropped int64 `json:"quantity_dropped"`
	// Singleton singleton integer.
	// Singleton singleton（单件标记） integer.
	Singleton int32 `json:"singleton"`
}

// KillmailContainedItem item object.
// KillmailContainedItem 物品对象.
type KillmailContainedItem struct {
	// Flag flag integer.
	// Flag 位置标记整数.
	Flag int32 `json:"flag"`
	// ItemTypeId item_type_id integer.
	// ItemTypeId item_type_id 整数.
	ItemTypeId int32 `json:"item_type_id"`
	// QuantityDestroyed quantity_destroyed integer.
	// QuantityDestroyed 被摧毁数量 integer.
	QuantityDestroyed int64 `json:"quantity_destroyed"`
	// QuantityDropped quantity_dropped integer.
	// QuantityDropped 掉落数量 integer.
	QuantityDropped int64 `json:"quantity_dropped"`
	// Singleton singleton integer.
	// Singleton singleton（单件标记） integer.
	Singleton int32 `json:"singleton"`
}

// Killmail 200 ok object.
// Killmail 200 ok 对象.
type Killmail struct {
	// Attackers attackers array.
	// Attackers 攻击者数组.
	Attackers []KillmailAttacker `json:"attackers"`
	// KillmailId ID of the killmail.
	// KillmailId 击杀报告 ID.
	KillmailId int32 `json:"killmail_id"`
	// KillmailTime Time that the victim was killed and the killmail generated.
	// KillmailTime 受害者被击杀并生成击杀报告的时间.
	KillmailTime time.Time `json:"killmail_time"`
	// MoonId Moon if the kill took place at one.
	// MoonId 若击杀发生在卫星处，则为该卫星.
	MoonId int32 `json:"moon_id"`
	// SolarSystemId Solar system that the kill took place in.
	// SolarSystemId 击杀发生的星系.
	SolarSystemId int32 `json:"solar_system_id"`
	// Supporters supporters array.
	// Supporters 支持者列表 array.
	Supporters []KillmailSupporter `json:"supporters"`
	// Victim victim object.
	// Victim victim 对象.
	Victim KillmailVictim `json:"victim"`
	// WarId War if the killmail is generated in relation to an official war.
	// WarId 如果击杀报告因正式战争而产生，则对应的战争.
	WarId int32 `json:"war_id"`
}

// KillmailPosition Coordinates of the victim in Cartesian space relative to the Sun.
// KillmailPosition 受害者相对于太阳的笛卡尔空间坐标.
type KillmailPosition struct {
	// X x number.
	// X x 数字.
	X float64 `json:"x"`
	// Y y number.
	// Y y 数字.
	Y float64 `json:"y"`
	// Z z number.
	// Z z 数字.
	Z float64 `json:"z"`
}

// KillmailSupporter supporter object.
// KillmailSupporter 支持者 object.
type KillmailSupporter struct {
	// AllianceId alliance_id integer.
	// AllianceId 联盟 ID 整数.
	AllianceId int32 `json:"alliance_id"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// IsRegen Obsolete attribute.
	// IsRegen 已废弃的属性.
	IsRegen bool `json:"is_regen"`
	// RepairDone The amount the supporter actually supporter toward the victim.
	// RepairDone 支援者实际支援受害者的金额.
	RepairDone int32 `json:"repair_done"`
	// RepairerTypeId The type the supporter used toward the victim.
	// RepairerTypeId 支援方对受害者使用的类型.
	RepairerTypeId int32 `json:"repairer_type_id"`
	// SecurityStatus Security status for the supporter.
	// SecurityStatus 支援者的安全等级.
	SecurityStatus float64 `json:"security_status"`
	// ShipTypeId What ship was the supporter flying.
	// ShipTypeId 支援方驾驶的舰船.
	ShipTypeId int32 `json:"ship_type_id"`
}

// KillmailVictim victim object.
// KillmailVictim victim 对象.
type KillmailVictim struct {
	// AllianceId alliance_id integer.
	// AllianceId 联盟 ID 整数.
	AllianceId int32 `json:"alliance_id"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// DamageTaken How much total damage was taken by the victim.
	// DamageTaken 受害者受到的总伤害.
	DamageTaken int32 `json:"damage_taken"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// Items items array.
	// Items items 数组.
	Items []KillmailItem `json:"items"`
	// Position Coordinates of the victim in Cartesian space relative to the Sun.
	// Position 受害者相对于太阳的笛卡尔空间坐标.
	Position KillmailPosition `json:"position"`
	// ShipTypeId The ship that the victim was piloting and was destroyed.
	// ShipTypeId 受害者驾驶的、已被击毁的舰船.
	ShipTypeId int32 `json:"ship_type_id"`
}
