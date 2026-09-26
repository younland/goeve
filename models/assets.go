package models

// CharacterAsset 200 ok object.
// CharacterAsset 200 ok 对象.
type CharacterAsset struct {
	// IsBlueprintCopy is_blueprint_copy boolean.
	// IsBlueprintCopy 是否为蓝图拷贝布尔值.
	IsBlueprintCopy bool `json:"is_blueprint_copy"`
	// IsSingleton is_singleton boolean.
	// IsSingleton 是否为单一物品布尔值.
	IsSingleton bool `json:"is_singleton"`
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// LocationFlag location_flag string.
	// LocationFlag location_flag 字符串.
	// Enum values: "AssetSafety", "AutoFit", "BoosterBay", "Cargo", "CorporationGoalDeliveries", "CorpseBay", "Deliveries", "DroneBay", "FighterBay", "FighterTube0", "FighterTube1", "FighterTube2", "FighterTube3", "FighterTube4", "FleetHangar", "FrigateEscapeBay", "Hangar", "HangarAll", "HiSlot0", "HiSlot1", "HiSlot2", "HiSlot3", "HiSlot4", "HiSlot5", "HiSlot6", "HiSlot7", "HiddenModifiers", "Implant", "LoSlot0", "LoSlot1", "LoSlot2", "LoSlot3", "LoSlot4", "LoSlot5", "LoSlot6", "LoSlot7", "Locked", "MedSlot0", "MedSlot1", "MedSlot2", "MedSlot3", "MedSlot4", "MedSlot5", "MedSlot6", "MedSlot7", "MobileDepotHold", "QuafeBay", "RigSlot0", "RigSlot1", "RigSlot2", "RigSlot3", "RigSlot4", "RigSlot5", "RigSlot6", "RigSlot7", "ShipHangar", "Skill", "SpecializedAmmoHold", "SpecializedAsteroidHold", "SpecializedCommandCenterHold", "SpecializedFuelBay", "SpecializedGasHold", "SpecializedIceHold", "SpecializedIndustrialShipHold", "SpecializedLargeShipHold", "SpecializedMaterialBay", "SpecializedMediumShipHold", "SpecializedMineralHold", "SpecializedOreHold", "SpecializedPlanetaryCommoditiesHold", "SpecializedSalvageHold", "SpecializedShipHold", "SpecializedSmallShipHold", "StructureDeedBay", "SubSystemBay", "SubSystemSlot0", "SubSystemSlot1", "SubSystemSlot2", "SubSystemSlot3", "SubSystemSlot4", "SubSystemSlot5", "SubSystemSlot6", "SubSystemSlot7", "Unlocked", "Wardrobe".
	LocationFlag string `json:"location_flag"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LocationType location_type string.
	// LocationType location_type 字符串.
	// Enum values: "station", "solar_system", "item", "other".
	LocationType string `json:"location_type"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// CorporationAsset 200 ok object.
// CorporationAsset 200 ok 对象.
type CorporationAsset struct {
	// IsBlueprintCopy is_blueprint_copy boolean.
	// IsBlueprintCopy 是否为蓝图拷贝布尔值.
	IsBlueprintCopy bool `json:"is_blueprint_copy"`
	// IsSingleton is_singleton boolean.
	// IsSingleton 是否为单一物品布尔值.
	IsSingleton bool `json:"is_singleton"`
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// LocationFlag location_flag string.
	// LocationFlag location_flag 字符串.
	// Enum values: "AssetSafety", "AutoFit", "Bonus", "Booster", "BoosterBay", "Capsule", "Cargo", "CorpDeliveries", "CorpSAG1", "CorpSAG2", "CorpSAG3", "CorpSAG4", "CorpSAG5", "CorpSAG6", "CorpSAG7", "CorporationGoalDeliveries", "CrateLoot", "Deliveries", "DroneBay", "DustBattle", "DustDatabank", "FighterBay", "FighterTube0", "FighterTube1", "FighterTube2", "FighterTube3", "FighterTube4", "FleetHangar", "FrigateEscapeBay", "Hangar", "HangarAll", "HiSlot0", "HiSlot1", "HiSlot2", "HiSlot3", "HiSlot4", "HiSlot5", "HiSlot6", "HiSlot7", "HiddenModifiers", "Implant", "Impounded", "JunkyardReprocessed", "JunkyardTrashed", "LoSlot0", "LoSlot1", "LoSlot2", "LoSlot3", "LoSlot4", "LoSlot5", "LoSlot6", "LoSlot7", "Locked", "MedSlot0", "MedSlot1", "MedSlot2", "MedSlot3", "MedSlot4", "MedSlot5", "MedSlot6", "MedSlot7", "MobileDepotHold", "OfficeFolder", "Pilot", "PlanetSurface", "QuafeBay", "QuantumCoreRoom", "Reward", "RigSlot0", "RigSlot1", "RigSlot2", "RigSlot3", "RigSlot4", "RigSlot5", "RigSlot6", "RigSlot7", "SecondaryStorage", "ServiceSlot0", "ServiceSlot1", "ServiceSlot2", "ServiceSlot3", "ServiceSlot4", "ServiceSlot5", "ServiceSlot6", "ServiceSlot7", "ShipHangar", "ShipOffline", "Skill", "SkillInTraining", "SpecializedAmmoHold", "SpecializedAsteroidHold", "SpecializedCommandCenterHold", "SpecializedFuelBay", "SpecializedGasHold", "SpecializedIceHold", "SpecializedIndustrialShipHold", "SpecializedLargeShipHold", "SpecializedMaterialBay", "SpecializedMediumShipHold", "SpecializedMineralHold", "SpecializedOreHold", "SpecializedPlanetaryCommoditiesHold", "SpecializedSalvageHold", "SpecializedShipHold", "SpecializedSmallShipHold", "StructureActive", "StructureFuel", "StructureInactive", "StructureOffline", "SubSystemBay", "SubSystemSlot0", "SubSystemSlot1", "SubSystemSlot2", "SubSystemSlot3", "SubSystemSlot4", "SubSystemSlot5", "SubSystemSlot6", "SubSystemSlot7", "Unlocked", "Wallet", "Wardrobe".
	LocationFlag string `json:"location_flag"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LocationType location_type string.
	// LocationType location_type 字符串.
	// Enum values: "station", "solar_system", "item", "other".
	LocationType string `json:"location_type"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// AssetLocation 200 ok object.
// AssetLocation 200 ok 对象.
type AssetLocation struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// Position position object.
	// Position position 对象.
	Position AssetPosition `json:"position"`
}

// AssetPosition position object.
// AssetPosition position 对象.
type AssetPosition struct {
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

// AssetName 200 ok object.
// AssetName 200 ok 对象.
type AssetName struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}
