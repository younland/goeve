package models

import (
	"net/url"
	"strconv"
)

// GetCharactersCharacterIdAssets 200 ok object.
// GetCharactersCharacterIdAssets 200 ok 对象.
type GetCharactersCharacterIdAssets struct {
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

// GetCharactersCharacterIdAssetsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdAssetsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdAssetsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdAssetsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCorporationsCorporationIdAssets 200 ok object.
// GetCorporationsCorporationIdAssets 200 ok 对象.
type GetCorporationsCorporationIdAssets struct {
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

// GetCorporationsCorporationIdAssetsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdAssetsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdAssetsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCorporationsCorporationIdAssetsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostCharactersCharacterIdAssetsLocations 200 ok object.
// PostCharactersCharacterIdAssetsLocations 200 ok 对象.
type PostCharactersCharacterIdAssetsLocations struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// Position position object.
	// Position position 对象.
	Position PostCharactersCharacterIdAssetsLocationsPosition `json:"position"`
}

// PostCharactersCharacterIdAssetsLocationsParams holds the optional query and header parameters of the request.
// PostCharactersCharacterIdAssetsLocationsParams 保存请求的可选查询与头部参数。
type PostCharactersCharacterIdAssetsLocationsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCharactersCharacterIdAssetsLocationsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostCharactersCharacterIdAssetsLocationsPosition position object.
// PostCharactersCharacterIdAssetsLocationsPosition position 对象.
type PostCharactersCharacterIdAssetsLocationsPosition struct {
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

// PostCharactersCharacterIdAssetsNames 200 ok object.
// PostCharactersCharacterIdAssetsNames 200 ok 对象.
type PostCharactersCharacterIdAssetsNames struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostCharactersCharacterIdAssetsNamesParams holds the optional query and header parameters of the request.
// PostCharactersCharacterIdAssetsNamesParams 保存请求的可选查询与头部参数。
type PostCharactersCharacterIdAssetsNamesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCharactersCharacterIdAssetsNamesParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostCorporationsCorporationIdAssetsLocations 200 ok object.
// PostCorporationsCorporationIdAssetsLocations 200 ok 对象.
type PostCorporationsCorporationIdAssetsLocations struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// Position position object.
	// Position position 对象.
	Position PostCorporationsCorporationIdAssetsLocationsPosition `json:"position"`
}

// PostCorporationsCorporationIdAssetsLocationsParams holds the optional query and header parameters of the request.
// PostCorporationsCorporationIdAssetsLocationsParams 保存请求的可选查询与头部参数。
type PostCorporationsCorporationIdAssetsLocationsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCorporationsCorporationIdAssetsLocationsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostCorporationsCorporationIdAssetsLocationsPosition position object.
// PostCorporationsCorporationIdAssetsLocationsPosition position 对象.
type PostCorporationsCorporationIdAssetsLocationsPosition struct {
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

// PostCorporationsCorporationIdAssetsNames 200 ok object.
// PostCorporationsCorporationIdAssetsNames 200 ok 对象.
type PostCorporationsCorporationIdAssetsNames struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostCorporationsCorporationIdAssetsNamesParams holds the optional query and header parameters of the request.
// PostCorporationsCorporationIdAssetsNamesParams 保存请求的可选查询与头部参数。
type PostCorporationsCorporationIdAssetsNamesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCorporationsCorporationIdAssetsNamesParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}
