package models

import (
	"time"
)

// CloneHomeLocation home_location object.
// CloneHomeLocation 常驻地点对象.
type CloneHomeLocation struct {
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LocationType location_type string.
	// LocationType location_type 字符串.
	// Enum values: "station", "structure".
	LocationType string `json:"location_type"`
}

// JumpClone jump_clone object.
// JumpClone jump_clone 对象.
type JumpClone struct {
	// Implants implants array.
	// Implants 植入体数组.
	Implants []int32 `json:"implants"`
	// JumpCloneId jump_clone_id integer.
	// JumpCloneId jump_clone_id 整数.
	JumpCloneId int32 `json:"jump_clone_id"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LocationType location_type string.
	// LocationType location_type 字符串.
	// Enum values: "station", "structure".
	LocationType string `json:"location_type"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// Clones 200 ok object.
// Clones 200 ok 对象.
type Clones struct {
	// HomeLocation home_location object.
	// HomeLocation 常驻地点对象.
	HomeLocation CloneHomeLocation `json:"home_location"`
	// JumpClones jump_clones array.
	// JumpClones jump_clones 数组.
	JumpClones []JumpClone `json:"jump_clones"`
	// LastCloneJumpDate last_clone_jump_date string.
	// LastCloneJumpDate last_clone_jump_date 字符串.
	LastCloneJumpDate time.Time `json:"last_clone_jump_date"`
	// LastStationChangeDate last_station_change_date string.
	// LastStationChangeDate last_station_change_date 字符串.
	LastStationChangeDate time.Time `json:"last_station_change_date"`
}
