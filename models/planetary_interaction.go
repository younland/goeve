package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetCharactersCharacterIdPlanets 200 ok object.
// GetCharactersCharacterIdPlanets 200 ok 对象.
type GetCharactersCharacterIdPlanets struct {
	// LastUpdate last_update string.
	// LastUpdate last_update 字符串.
	LastUpdate time.Time `json:"last_update"`
	// NumPins num_pins integer.
	// NumPins num_pins 整数.
	NumPins int32 `json:"num_pins"`
	// OwnerId owner_id integer.
	// OwnerId owner_id 整数.
	OwnerId int32 `json:"owner_id"`
	// PlanetId planet_id integer.
	// PlanetId planet_id 整数.
	PlanetId int32 `json:"planet_id"`
	// PlanetType planet_type string.
	// PlanetType planet_type 字符串.
	// Enum values: "temperate", "barren", "oceanic", "ice", "gas", "lava", "storm", "plasma".
	PlanetType string `json:"planet_type"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// UpgradeLevel upgrade_level integer.
	// UpgradeLevel 升级等级整数.
	UpgradeLevel int32 `json:"upgrade_level"`
}

// GetCharactersCharacterIdPlanetsPlanetIdContent content object.
// GetCharactersCharacterIdPlanetsPlanetIdContent 内容对象.
type GetCharactersCharacterIdPlanetsPlanetIdContent struct {
	// Amount amount integer.
	// Amount 金额整数.
	Amount int64 `json:"amount"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails extractor_details object.
// GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails 采集器详情对象.
type GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails struct {
	// CycleTime in seconds.
	// CycleTime 单位：秒.
	CycleTime int32 `json:"cycle_time"`
	// HeadRadius head_radius number.
	// HeadRadius 头部半径数值.
	HeadRadius float64 `json:"head_radius"`
	// Heads heads array.
	// Heads 头部数组.
	Heads []GetCharactersCharacterIdPlanetsPlanetIdHead `json:"heads"`
	// ProductTypeId product_type_id integer.
	// ProductTypeId product_type_id 整数.
	ProductTypeId int32 `json:"product_type_id"`
	// QtyPerCycle qty_per_cycle integer.
	// QtyPerCycle 每周期数量 integer.
	QtyPerCycle int32 `json:"qty_per_cycle"`
}

// GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails factory_details object.
// GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails 工厂详情对象.
type GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails struct {
	// SchematicId schematic_id integer.
	// SchematicId 示意图ID integer.
	SchematicId int32 `json:"schematic_id"`
}

// GetCharactersCharacterIdPlanetsPlanetIdHead head object.
// GetCharactersCharacterIdPlanetsPlanetIdHead 头部对象.
type GetCharactersCharacterIdPlanetsPlanetIdHead struct {
	// HeadId head_id integer.
	// HeadId 头部 ID 整数.
	HeadId int32 `json:"head_id"`
	// Latitude latitude number.
	// Latitude latitude 数值.
	Latitude float64 `json:"latitude"`
	// Longitude longitude number.
	// Longitude longitude 数值.
	Longitude float64 `json:"longitude"`
}

// GetCharactersCharacterIdPlanetsPlanetIdLink link object.
// GetCharactersCharacterIdPlanetsPlanetIdLink link 对象.
type GetCharactersCharacterIdPlanetsPlanetIdLink struct {
	// DestinationPinId destination_pin_id integer.
	// DestinationPinId 目标开采阵列 ID 整数.
	DestinationPinId int64 `json:"destination_pin_id"`
	// LinkLevel link_level integer.
	// LinkLevel link_level 整数.
	LinkLevel int32 `json:"link_level"`
	// SourcePinId source_pin_id integer.
	// SourcePinId 源采集器ID integer.
	SourcePinId int64 `json:"source_pin_id"`
}

// GetCharactersCharacterIdPlanetsPlanetIdOk 200 ok object.
// GetCharactersCharacterIdPlanetsPlanetIdOk 200 ok 对象.
type GetCharactersCharacterIdPlanetsPlanetIdOk struct {
	// Links links array.
	// Links links 数组.
	Links []GetCharactersCharacterIdPlanetsPlanetIdLink `json:"links"`
	// Pins pins array.
	// Pins pins 数组.
	Pins []GetCharactersCharacterIdPlanetsPlanetIdPin `json:"pins"`
	// Routes routes array.
	// Routes 路线列表 array.
	Routes []GetCharactersCharacterIdPlanetsPlanetIdRoute `json:"routes"`
}

// GetCharactersCharacterIdPlanetsPlanetIdPin pin object.
// GetCharactersCharacterIdPlanetsPlanetIdPin pin 对象.
type GetCharactersCharacterIdPlanetsPlanetIdPin struct {
	// Contents contents array.
	// Contents 内容数组.
	Contents []GetCharactersCharacterIdPlanetsPlanetIdContent `json:"contents"`
	// ExpiryTime expiry_time string.
	// ExpiryTime 过期时间字符串.
	ExpiryTime time.Time `json:"expiry_time"`
	// ExtractorDetails extractor_details object.
	// ExtractorDetails 采集器详情对象.
	ExtractorDetails GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails `json:"extractor_details"`
	// FactoryDetails factory_details object.
	// FactoryDetails 工厂详情对象.
	FactoryDetails GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails `json:"factory_details"`
	// InstallTime install_time string.
	// InstallTime 安装时间字符串.
	InstallTime time.Time `json:"install_time"`
	// LastCycleStart last_cycle_start string.
	// LastCycleStart last_cycle_start 字符串.
	LastCycleStart time.Time `json:"last_cycle_start"`
	// Latitude latitude number.
	// Latitude latitude 数值.
	Latitude float64 `json:"latitude"`
	// Longitude longitude number.
	// Longitude longitude 数值.
	Longitude float64 `json:"longitude"`
	// PinId pin_id integer.
	// PinId pin_id 整数.
	PinId int64 `json:"pin_id"`
	// SchematicId schematic_id integer.
	// SchematicId 示意图ID integer.
	SchematicId int32 `json:"schematic_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetCharactersCharacterIdPlanetsPlanetIdRoute route object.
// GetCharactersCharacterIdPlanetsPlanetIdRoute 路线 object.
type GetCharactersCharacterIdPlanetsPlanetIdRoute struct {
	// ContentTypeId content_type_id integer.
	// ContentTypeId 内容类型 ID 整数.
	ContentTypeId int32 `json:"content_type_id"`
	// DestinationPinId destination_pin_id integer.
	// DestinationPinId 目标开采阵列 ID 整数.
	DestinationPinId int64 `json:"destination_pin_id"`
	// Quantity quantity number.
	// Quantity 数量 number.
	Quantity float64 `json:"quantity"`
	// RouteId route_id integer.
	// RouteId 路线ID integer.
	RouteId int64 `json:"route_id"`
	// SourcePinId source_pin_id integer.
	// SourcePinId 源采集器ID integer.
	SourcePinId int64 `json:"source_pin_id"`
	// Waypoints list of pin ID waypoints.
	// Waypoints 行星设施 ID 路径点列表.
	Waypoints []int64 `json:"waypoints"`
}

// GetCorporationsCorporationIdCustomsOffices 200 ok object.
// GetCorporationsCorporationIdCustomsOffices 200 ok 对象.
type GetCorporationsCorporationIdCustomsOffices struct {
	// AllianceTaxRate Only present if alliance access is allowed.
	// AllianceTaxRate 仅当允许联盟访问时才会出现.
	AllianceTaxRate float64 `json:"alliance_tax_rate"`
	// AllowAccessWithStandings standing_level and any standing related tax rate only present when this is true.
	// AllowAccessWithStandings 仅当此项为 true 时才会包含声望等级及任何与声望相关的税率.
	AllowAccessWithStandings bool `json:"allow_access_with_standings"`
	// AllowAllianceAccess allow_alliance_access boolean.
	// AllowAllianceAccess 是否允许联盟访问布尔值.
	AllowAllianceAccess bool `json:"allow_alliance_access"`
	// BadStandingTaxRate bad_standing_tax_rate number.
	// BadStandingTaxRate 糟糕声望税率数值.
	BadStandingTaxRate float64 `json:"bad_standing_tax_rate"`
	// CorporationTaxRate corporation_tax_rate number.
	// CorporationTaxRate 军团税率数值.
	CorporationTaxRate float64 `json:"corporation_tax_rate"`
	// ExcellentStandingTaxRate Tax rate for entities with excellent level of standing, only present if this level is allowed, same for all other standing related tax rates.
	// ExcellentStandingTaxRate 声望等级为极佳（excellent）的实体适用的税率，仅当允许该等级时存在，其他声望相关税率同理.
	ExcellentStandingTaxRate float64 `json:"excellent_standing_tax_rate"`
	// GoodStandingTaxRate good_standing_tax_rate number.
	// GoodStandingTaxRate 良好声望税率数值.
	GoodStandingTaxRate float64 `json:"good_standing_tax_rate"`
	// NeutralStandingTaxRate neutral_standing_tax_rate number.
	// NeutralStandingTaxRate neutral_standing_tax_rate 数值.
	NeutralStandingTaxRate float64 `json:"neutral_standing_tax_rate"`
	// OfficeId unique ID of this customs office.
	// OfficeId 该海关办公室的唯一 ID.
	OfficeId int64 `json:"office_id"`
	// ReinforceExitEnd reinforce_exit_end integer.
	// ReinforceExitEnd 增强解除结束时间 integer.
	ReinforceExitEnd int32 `json:"reinforce_exit_end"`
	// ReinforceExitStart Together with reinforce_exit_end, marks a 2-hour period where this customs office could exit reinforcement mode during the day after initial attack.
	// ReinforceExitStart 与 reinforce_exit_end 一起，标记初始攻击后第二天该海关办公室可以退出增强模式的两小时时段.
	ReinforceExitStart int32 `json:"reinforce_exit_start"`
	// StandingLevel Access is allowed only for entities with this level of standing or better.
	// StandingLevel 仅允许声望达到或超过此等级的实体访问.
	// Enum values: "bad", "excellent", "good", "neutral", "terrible".
	StandingLevel string `json:"standing_level"`
	// SystemId ID of the solar system this customs office is located in.
	// SystemId 此海关办公室所在星系的 ID.
	SystemId int32 `json:"system_id"`
	// TerribleStandingTaxRate terrible_standing_tax_rate number.
	// TerribleStandingTaxRate 极差声望税率数字.
	TerribleStandingTaxRate float64 `json:"terrible_standing_tax_rate"`
}

// GetUniverseSchematicsSchematicIdOk 200 ok object.
// GetUniverseSchematicsSchematicIdOk 200 ok 对象.
type GetUniverseSchematicsSchematicIdOk struct {
	// CycleTime Time in seconds to process a run.
	// CycleTime 处理一轮作业所需的时间（秒）
	CycleTime int32 `json:"cycle_time"`
	// SchematicName schematic_name string.
	// SchematicName 示意图名称 string.
	SchematicName string `json:"schematic_name"`
}

// GetCharactersCharacterIdPlanetsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdPlanetsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdPlanetsParams struct {
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

func (p *GetCharactersCharacterIdPlanetsParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdPlanetsPlanetIdParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdPlanetsPlanetIdParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdPlanetsPlanetIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdPlanetsPlanetIdParams) Values() (url.Values, map[string]string) {
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

// GetCorporationsCorporationIdCustomsOfficesParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdCustomsOfficesParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdCustomsOfficesParams struct {
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

func (p *GetCorporationsCorporationIdCustomsOfficesParams) Values() (url.Values, map[string]string) {
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

// GetUniverseSchematicsSchematicIdParams holds the optional query and header parameters of the request.
// GetUniverseSchematicsSchematicIdParams 保存请求的可选查询与头部参数。
type GetUniverseSchematicsSchematicIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetUniverseSchematicsSchematicIdParams) Values() (url.Values, map[string]string) {
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
