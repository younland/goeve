package models

// Fitting 200 ok object.
// Fitting 200 ok 对象.
type Fitting struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// FittingId fitting_id integer.
	// FittingId 装配 ID 整数.
	FittingId int32 `json:"fitting_id"`
	// Items items array.
	// Items items 数组.
	Items []FittingItem `json:"items"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
}

// FittingItem item object.
// FittingItem 物品对象.
type FittingItem struct {
	// Flag flag string.
	// Flag 位置标记字符串.
	// Enum values: "Cargo", "DroneBay", "FighterBay", "HiSlot0", "HiSlot1", "HiSlot2", "HiSlot3", "HiSlot4", "HiSlot5", "HiSlot6", "HiSlot7", "Invalid", "LoSlot0", "LoSlot1", "LoSlot2", "LoSlot3", "LoSlot4", "LoSlot5", "LoSlot6", "LoSlot7", "MedSlot0", "MedSlot1", "MedSlot2", "MedSlot3", "MedSlot4", "MedSlot5", "MedSlot6", "MedSlot7", "RigSlot0", "RigSlot1", "RigSlot2", "ServiceSlot0", "ServiceSlot1", "ServiceSlot2", "ServiceSlot3", "ServiceSlot4", "ServiceSlot5", "ServiceSlot6", "ServiceSlot7", "SubSystemSlot0", "SubSystemSlot1", "SubSystemSlot2", "SubSystemSlot3".
	Flag string `json:"flag"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// NewFitting 201 created object.
// NewFitting 201 created 对象.
type NewFitting struct {
	// FittingId fitting_id integer.
	// FittingId 装配 ID 整数.
	FittingId int32 `json:"fitting_id"`
}

// FittingRequest fitting object.
// FittingRequest 装配对象.
type FittingRequest struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Items items array.
	// Items items 数组.
	Items []FittingItem `json:"items"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
}
