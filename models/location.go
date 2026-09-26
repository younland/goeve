package models

import (
	"time"
)

// CharacterLocation 200 ok object.
// CharacterLocation 200 ok 对象.
type CharacterLocation struct {
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// StationId station_id integer.
	// StationId 空间站ID integer.
	StationId int32 `json:"station_id"`
	// StructureId structure_id integer.
	// StructureId 建筑ID integer.
	StructureId int64 `json:"structure_id"`
}

// OnlineStatus 200 ok object.
// OnlineStatus 200 ok 对象.
type OnlineStatus struct {
	// LastLogin Timestamp of the last login.
	// LastLogin 最后一次登录的时间戳.
	LastLogin time.Time `json:"last_login"`
	// LastLogout Timestamp of the last logout.
	// LastLogout 最后一次登出的时间戳.
	LastLogout time.Time `json:"last_logout"`
	// Logins Total number of times the character has logged in.
	// Logins 该角色的累计登录次数.
	Logins int32 `json:"logins"`
	// Online If the character is online.
	// Online 角色是否在线.
	Online bool `json:"online"`
}

// CharacterShip 200 ok object.
// CharacterShip 200 ok 对象.
type CharacterShip struct {
	// ShipItemId Item id's are unique to a ship and persist until it is repackaged. This value can be used to track repeated uses of a ship, or detect when a pilot changes into a different instance of the same ship type.
	// ShipItemId 物品 ID 对一艘舰船唯一，并在其重新打包之前保持不变。该值可用于追踪舰船的重复使用，或检测飞行员何时更换为同型舰船的另一实例。
	ShipItemId int64 `json:"ship_item_id"`
	// ShipName ship_name string.
	// ShipName 舰船名称 string.
	ShipName string `json:"ship_name"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
}
