package models

// SearchResult 200 ok object.
// SearchResult 200 ok 对象.
type SearchResult struct {
	// Agent agent array.
	// Agent 代理人数组.
	Agent []int32 `json:"agent"`
	// Alliance alliance array.
	// Alliance 联盟数组.
	Alliance []int32 `json:"alliance"`
	// Character character array.
	// Character 角色数组.
	Character []int32 `json:"character"`
	// Constellation constellation array.
	// Constellation 星座数组.
	Constellation []int32 `json:"constellation"`
	// Corporation corporation array.
	// Corporation 军团数组.
	Corporation []int32 `json:"corporation"`
	// Faction faction array.
	// Faction 势力数组.
	Faction []int32 `json:"faction"`
	// InventoryType inventory_type array.
	// InventoryType 物品类型数组.
	InventoryType []int32 `json:"inventory_type"`
	// Region region array.
	// Region 星域 array.
	Region []int32 `json:"region"`
	// SolarSystem solar_system array.
	// SolarSystem 星系 array.
	SolarSystem []int32 `json:"solar_system"`
	// Station station array.
	// Station 空间站 array.
	Station []int32 `json:"station"`
	// Structure structure array.
	// Structure 建筑 array.
	Structure []int64 `json:"structure"`
}
