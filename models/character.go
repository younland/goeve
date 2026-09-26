package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetCharacterIdAgentsResearchParams holds the optional query and header parameters of the request.
// GetCharacterIdAgentsResearchParams 保存请求的可选查询与头部参数。
type GetCharacterIdAgentsResearchParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdAgentsResearchParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharacterIdBlueprintsParams holds the optional query and header parameters of the request.
// GetCharacterIdBlueprintsParams 保存请求的可选查询与头部参数。
type GetCharacterIdBlueprintsParams struct {
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

func (p *GetCharacterIdBlueprintsParams) Values() (url.Values, map[string]string) {
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

// GetCharacterIdCorporationhistoryParams holds the optional query and header parameters of the request.
// GetCharacterIdCorporationhistoryParams 保存请求的可选查询与头部参数。
type GetCharacterIdCorporationhistoryParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetCharacterIdCorporationhistoryParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetCharacterIdFatigueParams holds the optional query and header parameters of the request.
// GetCharacterIdFatigueParams 保存请求的可选查询与头部参数。
type GetCharacterIdFatigueParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdFatigueParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharacterIdMedalsParams holds the optional query and header parameters of the request.
// GetCharacterIdMedalsParams 保存请求的可选查询与头部参数。
type GetCharacterIdMedalsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdMedalsParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharacterIdNotificationsContactsParams holds the optional query and header parameters of the request.
// GetCharacterIdNotificationsContactsParams 保存请求的可选查询与头部参数。
type GetCharacterIdNotificationsContactsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdNotificationsContactsParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharacterIdNotificationsParams holds the optional query and header parameters of the request.
// GetCharacterIdNotificationsParams 保存请求的可选查询与头部参数。
type GetCharacterIdNotificationsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdNotificationsParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharacterIdParams holds the optional query and header parameters of the request.
// GetCharacterIdParams 保存请求的可选查询与头部参数。
type GetCharacterIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetCharacterIdParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetCharacterIdPortraitParams holds the optional query and header parameters of the request.
// GetCharacterIdPortraitParams 保存请求的可选查询与头部参数。
type GetCharacterIdPortraitParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetCharacterIdPortraitParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetCharacterIdRolesParams holds the optional query and header parameters of the request.
// GetCharacterIdRolesParams 保存请求的可选查询与头部参数。
type GetCharacterIdRolesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdRolesParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharacterIdStandingsParams holds the optional query and header parameters of the request.
// GetCharacterIdStandingsParams 保存请求的可选查询与头部参数。
type GetCharacterIdStandingsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdStandingsParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharacterIdTitlesParams holds the optional query and header parameters of the request.
// GetCharacterIdTitlesParams 保存请求的可选查询与头部参数。
type GetCharacterIdTitlesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharacterIdTitlesParams) Values() (url.Values, map[string]string) {
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
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharactersCharacterIdAgentsResearch 200 ok object.
// GetCharactersCharacterIdAgentsResearch 200 ok 对象.
type GetCharactersCharacterIdAgentsResearch struct {
	// AgentId agent_id integer.
	// AgentId 代理人 ID 整数.
	AgentId int32 `json:"agent_id"`
	// PointsPerDay points_per_day number.
	// PointsPerDay points_per_day 数值.
	PointsPerDay float64 `json:"points_per_day"`
	// RemainderPoints remainder_points number.
	// RemainderPoints 剩余点数 number.
	RemainderPoints float64 `json:"remainder_points"`
	// SkillTypeId skill_type_id integer.
	// SkillTypeId 技能类型ID integer.
	SkillTypeId int32 `json:"skill_type_id"`
	// StartedAt started_at string.
	// StartedAt 开始时间 string.
	StartedAt time.Time `json:"started_at"`
}

// GetCharactersCharacterIdBlueprints 200 ok object.
// GetCharactersCharacterIdBlueprints 200 ok 对象.
type GetCharactersCharacterIdBlueprints struct {
	// ItemId Unique ID for this item.
	// ItemId 此物品的唯一 ID。
	ItemId int64 `json:"item_id"`
	// LocationFlag Type of the location_id.
	// LocationFlag location_id 的类型.
	// Enum values: "AutoFit", "Cargo", "CorpseBay", "DroneBay", "FleetHangar", "Deliveries", "HiddenModifiers", "Hangar", "HangarAll", "LoSlot0", "LoSlot1", "LoSlot2", "LoSlot3", "LoSlot4", "LoSlot5", "LoSlot6", "LoSlot7", "MedSlot0", "MedSlot1", "MedSlot2", "MedSlot3", "MedSlot4", "MedSlot5", "MedSlot6", "MedSlot7", "HiSlot0", "HiSlot1", "HiSlot2", "HiSlot3", "HiSlot4", "HiSlot5", "HiSlot6", "HiSlot7", "AssetSafety", "Locked", "Unlocked", "Implant", "QuafeBay", "RigSlot0", "RigSlot1", "RigSlot2", "RigSlot3", "RigSlot4", "RigSlot5", "RigSlot6", "RigSlot7", "ShipHangar", "SpecializedFuelBay", "SpecializedOreHold", "SpecializedGasHold", "SpecializedMineralHold", "SpecializedSalvageHold", "SpecializedShipHold", "SpecializedSmallShipHold", "SpecializedMediumShipHold", "SpecializedLargeShipHold", "SpecializedIndustrialShipHold", "SpecializedAmmoHold", "SpecializedCommandCenterHold", "SpecializedPlanetaryCommoditiesHold", "SpecializedMaterialBay", "SubSystemSlot0", "SubSystemSlot1", "SubSystemSlot2", "SubSystemSlot3", "SubSystemSlot4", "SubSystemSlot5", "SubSystemSlot6", "SubSystemSlot7", "FighterBay", "FighterTube0", "FighterTube1", "FighterTube2", "FighterTube3", "FighterTube4", "Module".
	LocationFlag string `json:"location_flag"`
	// LocationId References a station, a ship or an item_id if this blueprint is located within a container. If the return value is an item_id, then the Character AssetList API must be queried to find the container using the given item_id to determine the correct location of the Blueprint.
	// LocationId 若蓝图位于集装箱内，则引用一个空间站、舰船或 item_id。如果返回值是 item_id，则必须查询角色资产列表 API，用给定的 item_id 找到集装箱，以确定蓝图的确切位置。
	LocationId int64 `json:"location_id"`
	// MaterialEfficiency Material Efficiency Level of the blueprint.
	// MaterialEfficiency 蓝图的材料效率等级。
	MaterialEfficiency int32 `json:"material_efficiency"`
	// Quantity A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet).
	// Quantity 一个数值范围，最小值为 -2，无最大值，其中 -1 表示原版蓝图，-2 表示拷贝。如果是一叠刚从市场上买来的原版蓝图（例如尚未对其执行任何活动），则可以是正整数。
	Quantity int32 `json:"quantity"`
	// Runs Number of runs remaining if the blueprint is a copy, -1 if it is an original.
	// Runs 若蓝图为拷贝则为其剩余运转次数，若为原始蓝图则为 -1。
	Runs int32 `json:"runs"`
	// TimeEfficiency Time Efficiency Level of the blueprint.
	// TimeEfficiency 蓝图的时间效率等级。
	TimeEfficiency int32 `json:"time_efficiency"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetCharactersCharacterIdCorporationhistory 200 ok object.
// GetCharactersCharacterIdCorporationhistory 200 ok 对象.
type GetCharactersCharacterIdCorporationhistory struct {
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// IsDeleted True if the corporation has been deleted.
	// IsDeleted 如果军团已删除则为 true.
	IsDeleted bool `json:"is_deleted"`
	// RecordId An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous.
	// RecordId 一个递增的 ID，可用于在日期可能含糊不清的情况下规范地确定记录的顺序.
	RecordId int32 `json:"record_id"`
	// StartDate start_date string.
	// StartDate 开始日期 string.
	StartDate time.Time `json:"start_date"`
}

// GetCharactersCharacterIdFatigue 200 ok object.
// GetCharactersCharacterIdFatigue 200 ok 对象.
type GetCharactersCharacterIdFatigue struct {
	// JumpFatigueExpireDate Character's jump fatigue expiry.
	// JumpFatigueExpireDate 角色跳跃疲劳到期时间.
	JumpFatigueExpireDate time.Time `json:"jump_fatigue_expire_date"`
	// LastJumpDate Character's last jump activation.
	// LastJumpDate 角色上次跳跃激活时间.
	LastJumpDate time.Time `json:"last_jump_date"`
	// LastUpdateDate Character's last jump update.
	// LastUpdateDate 角色上次跳跃更新时间.
	LastUpdateDate time.Time `json:"last_update_date"`
}

// GetCharactersCharacterIdMedals 200 ok object.
// GetCharactersCharacterIdMedals 200 ok 对象.
type GetCharactersCharacterIdMedals struct {
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// Date date string.
	// Date 日期字符串.
	Date time.Time `json:"date"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Graphics graphics array.
	// Graphics 图形数组.
	Graphics []GetCharactersCharacterIdMedalsGraphic `json:"graphics"`
	// IssuerId issuer_id integer.
	// IssuerId 发布者 ID 整数.
	IssuerId int32 `json:"issuer_id"`
	// MedalId medal_id integer.
	// MedalId medal_id 整数.
	MedalId int32 `json:"medal_id"`
	// Reason reason string.
	// Reason 原因 string.
	Reason string `json:"reason"`
	// Status status string.
	// Status 状态 string.
	// Enum values: "public", "private".
	Status string `json:"status"`
	// Title title string.
	// Title 标题字符串.
	Title string `json:"title"`
}

// GetCharactersCharacterIdMedalsGraphic graphic object.
// GetCharactersCharacterIdMedalsGraphic 图形对象.
type GetCharactersCharacterIdMedalsGraphic struct {
	// Color color integer.
	// Color 颜色整数.
	Color int32 `json:"color"`
	// Graphic graphic string.
	// Graphic 图形字符串.
	Graphic string `json:"graphic"`
	// Layer layer integer.
	// Layer layer 整数.
	Layer int32 `json:"layer"`
	// Part part integer.
	// Part part 整数.
	Part int32 `json:"part"`
}

// GetCharactersCharacterIdNotifications 200 ok object.
// GetCharactersCharacterIdNotifications 200 ok 对象.
type GetCharactersCharacterIdNotifications struct {
	// IsRead is_read boolean.
	// IsRead 是否已读布尔值.
	IsRead bool `json:"is_read"`
	// NotificationId notification_id integer.
	// NotificationId notification_id 整数.
	NotificationId int64 `json:"notification_id"`
	// SenderId sender_id integer.
	// SenderId 发送者ID integer.
	SenderId int32 `json:"sender_id"`
	// SenderType sender_type string.
	// SenderType 发送者类型 string.
	// Enum values: "character", "corporation", "alliance", "faction", "other".
	SenderType string `json:"sender_type"`
	// Text text string.
	// Text 文本字符串.
	Text string `json:"text"`
	// Timestamp timestamp string.
	// Timestamp 时间戳字符串.
	Timestamp time.Time `json:"timestamp"`
	// TypeValue type string.
	// TypeValue 类型字符串.
	// Enum values: "AcceptedAlly", "AcceptedSurrender", "AgentRetiredTrigravian", "AllAnchoringMsg", "AllMaintenanceBillMsg", "AllStrucInvulnerableMsg", "AllStructVulnerableMsg", "AllWarCorpJoinedAllianceMsg", "AllWarDeclaredMsg", "AllWarInvalidatedMsg", "AllWarRetractedMsg", "AllWarSurrenderMsg", "AllianceCapitalChanged", "AllianceWarDeclaredV2", "AllyContractCancelled", "AllyJoinedWarAggressorMsg", "AllyJoinedWarAllyMsg", "AllyJoinedWarDefenderMsg", "BattlePunishFriendlyFire", "BillOutOfMoneyMsg", "BillPaidCorpAllMsg", "BountyClaimMsg", "BountyESSShared", "BountyESSTaken", "BountyPlacedAlliance", "BountyPlacedChar", "BountyPlacedCorp", "BountyYourBountyClaimed", "BuddyConnectContactAdd", "CharAppAcceptMsg", "CharAppRejectMsg", "CharAppWithdrawMsg", "CharLeftCorpMsg", "CharMedalMsg", "CharTerminationMsg", "CloneActivationMsg", "CloneActivationMsg2", "CloneMovedMsg", "CloneRevokedMsg1", "CloneRevokedMsg2", "CombatOperationFinished", "ContactAdd", "ContactEdit", "ContainerPasswordMsg", "ContractRegionChangedToPochven", "CorpAllBillMsg", "CorpAppAcceptMsg", "CorpAppInvitedMsg", "CorpAppNewMsg", "CorpAppRejectCustomMsg", "CorpAppRejectMsg", "CorpBecameWarEligible", "CorpDividendMsg", "CorpFriendlyFireDisableTimerCompleted", "CorpFriendlyFireDisableTimerStarted", "CorpFriendlyFireEnableTimerCompleted", "CorpFriendlyFireEnableTimerStarted", "CorpKicked", "CorpLiquidationMsg", "CorpNewCEOMsg", "CorpNewsMsg", "CorpNoLongerWarEligible", "CorpOfficeExpirationMsg", "CorpStructLostMsg", "CorpTaxChangeMsg", "CorpVoteCEORevokedMsg", "CorpVoteMsg", "CorpWarDeclaredMsg", "CorpWarDeclaredV2", "CorpWarFightingLegalMsg", "CorpWarInvalidatedMsg", "CorpWarRetractedMsg", "CorpWarSurrenderMsg", "CorporationGoalClosed", "CorporationGoalCompleted", "CorporationGoalCreated", "CustomsMsg", "DeclareWar", "DistrictAttacked", "DustAppAcceptedMsg", "ESSMainBankLink", "EntosisCaptureStarted", "ExpertSystemExpired", "ExpertSystemExpiryImminent", "FWAllianceKickMsg", "FWAllianceWarningMsg", "FWCharKickMsg", "FWCharRankGainMsg", "FWCharRankLossMsg", "FWCharWarningMsg", "FWCorpJoinMsg", "FWCorpKickMsg", "FWCorpLeaveMsg", "FWCorpWarningMsg", "FacWarCorpJoinRequestMsg", "FacWarCorpJoinWithdrawMsg", "FacWarCorpLeaveRequestMsg", "FacWarCorpLeaveWithdrawMsg", "FacWarLPDisqualifiedEvent", "FacWarLPDisqualifiedKill", "FacWarLPPayoutEvent", "FacWarLPPayoutKill", "GameTimeAdded", "GameTimeReceived", "GameTimeSent", "GiftReceived", "IHubDestroyedByBillFailure", "IncursionCompletedMsg", "IndustryOperationFinished", "IndustryTeamAuctionLost", "IndustryTeamAuctionWon", "InfrastructureHubBillAboutToExpire", "InsuranceExpirationMsg", "InsuranceFirstShipMsg", "InsuranceInvalidatedMsg", "InsuranceIssuedMsg", "InsurancePayoutMsg", "InvasionCompletedMsg", "InvasionSystemLogin", "InvasionSystemStart", "JumpCloneDeletedMsg1", "JumpCloneDeletedMsg2", "KillReportFinalBlow", "KillReportVictim", "KillRightAvailable", "KillRightAvailableOpen", "KillRightEarned", "KillRightUnavailable", "KillRightUnavailableOpen", "KillRightUsed", "LocateCharMsg", "MadeWarMutual", "MercOfferRetractedMsg", "MercOfferedNegotiationMsg", "MissionCanceledTriglavian", "MissionOfferExpirationMsg", "MissionTimeoutMsg", "MoonminingAutomaticFracture", "MoonminingExtractionCancelled", "MoonminingExtractionFinished", "MoonminingExtractionStarted", "MoonminingLaserFired", "MutualWarExpired", "MutualWarInviteAccepted", "MutualWarInviteRejected", "MutualWarInviteSent", "NPCStandingsGained", "NPCStandingsLost", "OfferToAllyRetracted", "OfferedSurrender", "OfferedToAlly", "OfficeLeaseCanceledInsufficientStandings", "OldLscMessages", "OperationFinished", "OrbitalAttacked", "OrbitalReinforced", "OwnershipTransferred", "RaffleCreated", "RaffleExpired", "RaffleFinished", "ReimbursementMsg", "ResearchMissionAvailableMsg", "RetractsWar", "SeasonalChallengeCompleted", "SovAllClaimAquiredMsg", "SovAllClaimLostMsg", "SovCommandNodeEventStarted", "SovCorpBillLateMsg", "SovCorpClaimFailMsg", "SovDisruptorMsg", "SovStationEnteredFreeport", "SovStructureDestroyed", "SovStructureReinforced", "SovStructureSelfDestructCancel", "SovStructureSelfDestructFinished", "SovStructureSelfDestructRequested", "SovereigntyIHDamageMsg", "SovereigntySBUDamageMsg", "SovereigntyTCUDamageMsg", "StationAggressionMsg1", "StationAggressionMsg2", "StationConquerMsg", "StationServiceDisabled", "StationServiceEnabled", "StationStateChangeMsg", "StoryLineMissionAvailableMsg", "StructureAnchoring", "StructureCourierContractChanged", "StructureDestroyed", "StructureFuelAlert", "StructureImpendingAbandonmentAssetsAtRisk", "StructureItemsDelivered", "StructureItemsMovedToSafety", "StructureLostArmor", "StructureLostShields", "StructureOnline", "StructurePaintPurchased", "StructureServicesOffline", "StructureUnanchoring", "StructureUnderAttack", "StructureWentHighPower", "StructureWentLowPower", "StructuresJobsCancelled", "StructuresJobsPaused", "StructuresReinforcementChanged", "TowerAlertMsg", "TowerResourceAlertMsg", "TransactionReversalMsg", "TutorialMsg", "WarAdopted ", "WarAllyInherited", "WarAllyOfferDeclinedMsg", "WarConcordInvalidates", "WarDeclared", "WarEndedHqSecurityDrop", "WarHQRemovedFromSpace", "WarInherited", "WarInvalid", "WarRetracted", "WarRetractedByConcord", "WarSurrenderDeclinedMsg", "WarSurrenderOfferMsg".
	TypeValue string `json:"type"`
}

// GetCharactersCharacterIdNotificationsContacts 200 ok object.
// GetCharactersCharacterIdNotificationsContacts 200 ok 对象.
type GetCharactersCharacterIdNotificationsContacts struct {
	// Message message string.
	// Message message 字符串.
	Message string `json:"message"`
	// NotificationId notification_id integer.
	// NotificationId notification_id 整数.
	NotificationId int32 `json:"notification_id"`
	// SendDate send_date string.
	// SendDate 发送日期 string.
	SendDate time.Time `json:"send_date"`
	// SenderCharacterId sender_character_id integer.
	// SenderCharacterId 发送者角色ID integer.
	SenderCharacterId int32 `json:"sender_character_id"`
	// StandingLevel A number representing the standing level the receiver has been added at by the sender. The standing levels are as follows: -10 -> Terrible | -5 -> Bad |  0 -> Neutral |  5 -> Good |  10 -> Excellent.
	// StandingLevel 一个数字，表示发送者将接收者添加到的声望等级。声望等级如下：-10 -> 极差 | -5 -> 糟糕 | 0 -> 中立 | 5 -> 良好 | 10 -> 极佳.
	StandingLevel float64 `json:"standing_level"`
}

// GetCharactersCharacterId 200 ok object.
// GetCharactersCharacterId 200 ok 对象.
type GetCharactersCharacterId struct {
	// AllianceId The character's alliance ID.
	// AllianceId 角色的联盟 ID.
	AllianceId int32 `json:"alliance_id"`
	// Birthday Creation date of the character.
	// Birthday 角色的创建日期.
	Birthday time.Time `json:"birthday"`
	// BloodlineId bloodline_id integer.
	// BloodlineId 血统 ID 整数.
	BloodlineId int32 `json:"bloodline_id"`
	// CorporationId The character's corporation ID.
	// CorporationId 角色的军团 ID.
	CorporationId int32 `json:"corporation_id"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// FactionId ID of the faction the character is fighting for, if the character is enlisted in Factional Warfare.
	// FactionId 角色在势力战争中为之作战的势力 ID（若角色已加入势力战争）
	FactionId int32 `json:"faction_id"`
	// Gender gender string.
	// Gender 性别字符串.
	// Enum values: "female", "male".
	Gender string `json:"gender"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// RaceId race_id integer.
	// RaceId 种族ID integer.
	RaceId int32 `json:"race_id"`
	// SecurityStatus security_status number.
	// SecurityStatus 安全状态 number.
	SecurityStatus float64 `json:"security_status"`
	// Title The individual title of the character.
	// Title 角色的个人头衔.
	Title string `json:"title"`
}

// GetCharactersCharacterIdPortrait 200 ok object.
// GetCharactersCharacterIdPortrait 200 ok 对象.
type GetCharactersCharacterIdPortrait struct {
	// Px128x128 px128x128 string.
	// Px128x128 px128x128 字符串.
	Px128x128 string `json:"px128x128"`
	// Px256x256 px256x256 string.
	// Px256x256 px256x256 字符串.
	Px256x256 string `json:"px256x256"`
	// Px512x512 px512x512 string.
	// Px512x512 px512x512 字符串.
	Px512x512 string `json:"px512x512"`
	// Px64x64 px64x64 string.
	// Px64x64 px64x64 字符串.
	Px64x64 string `json:"px64x64"`
}

// GetCharactersCharacterIdRoles 200 ok object.
// GetCharactersCharacterIdRoles 200 ok 对象.
type GetCharactersCharacterIdRoles struct {
	// Roles roles array.
	// Roles 角色列表 array.
	Roles []string `json:"roles"`
	// RolesAtBase roles_at_base array.
	// RolesAtBase 基地角色列表 array.
	RolesAtBase []string `json:"roles_at_base"`
	// RolesAtHq roles_at_hq array.
	// RolesAtHq 总部角色列表 array.
	RolesAtHq []string `json:"roles_at_hq"`
	// RolesAtOther roles_at_other array.
	// RolesAtOther 其他地点角色列表 array.
	RolesAtOther []string `json:"roles_at_other"`
}

// GetCharactersCharacterIdStandings 200 ok object.
// GetCharactersCharacterIdStandings 200 ok 对象.
type GetCharactersCharacterIdStandings struct {
	// FromId from_id integer.
	// FromId 发送者 ID 整数.
	FromId int32 `json:"from_id"`
	// FromType from_type string.
	// FromType 发送者类型字符串.
	// Enum values: "agent", "npc_corp", "faction".
	FromType string `json:"from_type"`
	// Standing standing number.
	// Standing 声望 number.
	Standing float64 `json:"standing"`
}

// GetCharactersCharacterIdTitles 200 ok object.
// GetCharactersCharacterIdTitles 200 ok 对象.
type GetCharactersCharacterIdTitles struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// TitleId title_id integer.
	// TitleId title_id 整数.
	TitleId int32 `json:"title_id"`
}

// PostAffiliationParams holds the optional query and header parameters of the request.
// PostAffiliationParams 保存请求的可选查询与头部参数。
type PostAffiliationParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
}

func (p *PostAffiliationParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// PostCharacterIdCspaParams holds the optional query and header parameters of the request.
// PostCharacterIdCspaParams 保存请求的可选查询与头部参数。
type PostCharacterIdCspaParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCharacterIdCspaParams) Values() (url.Values, map[string]string) {
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

// PostCharactersAffiliation 200 ok object.
// PostCharactersAffiliation 200 ok 对象.
type PostCharactersAffiliation struct {
	// AllianceId The character's alliance ID, if their corporation is in an alliance.
	// AllianceId 角色的联盟 ID，前提是其军团加入了联盟.
	AllianceId int32 `json:"alliance_id"`
	// CharacterId The character's ID.
	// CharacterId 角色 ID.
	CharacterId int32 `json:"character_id"`
	// CorporationId The character's corporation ID.
	// CorporationId 角色的军团 ID.
	CorporationId int32 `json:"corporation_id"`
	// FactionId The character's faction ID, if their corporation is in a faction.
	// FactionId 角色的势力 ID，前提是其军团属于某势力.
	FactionId int32 `json:"faction_id"`
}
