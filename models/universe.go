package models

import (
	"net/url"
	"strconv"
)

// GetUniverseAncestries 200 ok object.
// GetUniverseAncestries 200 ok 对象.
type GetUniverseAncestries struct {
	// BloodlineId The bloodline associated with this ancestry.
	// BloodlineId 与该血统关联的血脉.
	BloodlineId int32 `json:"bloodline_id"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// IconId icon_id integer.
	// IconId 图标 ID 整数.
	IconId int32 `json:"icon_id"`
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// ShortDescription short_description string.
	// ShortDescription 简短描述 string.
	ShortDescription string `json:"short_description"`
}

// GetUniverseAsteroidBeltsAsteroidBeltIdOk 200 ok object.
// GetUniverseAsteroidBeltsAsteroidBeltIdOk 200 ok 对象.
type GetUniverseAsteroidBeltsAsteroidBeltIdOk struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position GetUniverseAsteroidBeltsAsteroidBeltIdPosition `json:"position"`
	// SystemId The solar system this asteroid belt is in.
	// SystemId 该小行星带所在的星系.
	SystemId int32 `json:"system_id"`
}

// GetUniverseAsteroidBeltsAsteroidBeltIdPosition position object.
// GetUniverseAsteroidBeltsAsteroidBeltIdPosition position 对象.
type GetUniverseAsteroidBeltsAsteroidBeltIdPosition struct {
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

// GetUniverseBloodlines 200 ok object.
// GetUniverseBloodlines 200 ok 对象.
type GetUniverseBloodlines struct {
	// BloodlineId bloodline_id integer.
	// BloodlineId 血统 ID 整数.
	BloodlineId int32 `json:"bloodline_id"`
	// Charisma charisma integer.
	// Charisma 魅力整数.
	Charisma int32 `json:"charisma"`
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Intelligence intelligence integer.
	// Intelligence 智力整数.
	Intelligence int32 `json:"intelligence"`
	// Memory memory integer.
	// Memory memory 整数.
	Memory int32 `json:"memory"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Perception perception integer.
	// Perception perception 整数.
	Perception int32 `json:"perception"`
	// RaceId race_id integer.
	// RaceId 种族ID integer.
	RaceId int32 `json:"race_id"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
	// Willpower willpower integer.
	// Willpower 毅力（willpower）整数.
	Willpower int32 `json:"willpower"`
}

// GetUniverseCategoriesCategoryIdOk 200 ok object.
// GetUniverseCategoriesCategoryIdOk 200 ok 对象.
type GetUniverseCategoriesCategoryIdOk struct {
	// CategoryId category_id integer.
	// CategoryId 分类 ID 整数.
	CategoryId int32 `json:"category_id"`
	// Groups groups array.
	// Groups 分组数组.
	Groups []int32 `json:"groups"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Published published boolean.
	// Published published 布尔.
	Published bool `json:"published"`
}

// GetUniverseConstellationsConstellationIdOk 200 ok object.
// GetUniverseConstellationsConstellationIdOk 200 ok 对象.
type GetUniverseConstellationsConstellationIdOk struct {
	// ConstellationId constellation_id integer.
	// ConstellationId 星座 ID 整数.
	ConstellationId int32 `json:"constellation_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position GetUniverseConstellationsConstellationIdPosition `json:"position"`
	// RegionId The region this constellation is in.
	// RegionId 该星座所在的星域.
	RegionId int32 `json:"region_id"`
	// Systems systems array.
	// Systems systems 数组.
	Systems []int32 `json:"systems"`
}

// GetUniverseConstellationsConstellationIdPosition position object.
// GetUniverseConstellationsConstellationIdPosition position 对象.
type GetUniverseConstellationsConstellationIdPosition struct {
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

// GetUniverseFactions 200 ok object.
// GetUniverseFactions 200 ok 对象.
type GetUniverseFactions struct {
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// IsUnique is_unique boolean.
	// IsUnique 是否唯一布尔值.
	IsUnique bool `json:"is_unique"`
	// MilitiaCorporationId militia_corporation_id integer.
	// MilitiaCorporationId militia_corporation_id 整数（民兵军团 ID）
	MilitiaCorporationId int32 `json:"militia_corporation_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// SizeFactor size_factor number.
	// SizeFactor 尺寸系数 number.
	SizeFactor float64 `json:"size_factor"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// StationCount station_count integer.
	// StationCount 空间站数量 integer.
	StationCount int32 `json:"station_count"`
	// StationSystemCount station_system_count integer.
	// StationSystemCount 含空间站星系数量 integer.
	StationSystemCount int32 `json:"station_system_count"`
}

// GetUniverseGraphicsGraphicIdOk 200 ok object.
// GetUniverseGraphicsGraphicIdOk 200 ok 对象.
type GetUniverseGraphicsGraphicIdOk struct {
	// CollisionFile collision_file string.
	// CollisionFile 碰撞文件字符串.
	CollisionFile string `json:"collision_file"`
	// GraphicFile graphic_file string.
	// GraphicFile 图形文件字符串.
	GraphicFile string `json:"graphic_file"`
	// GraphicId graphic_id integer.
	// GraphicId 图形 ID 整数.
	GraphicId int32 `json:"graphic_id"`
	// IconFolder icon_folder string.
	// IconFolder 图标文件夹字符串.
	IconFolder string `json:"icon_folder"`
	// SofDna sof_dna string.
	// SofDna SOF DNA string.
	SofDna string `json:"sof_dna"`
	// SofFationName sof_fation_name string.
	// SofFationName SOF 派系名称 string.
	SofFationName string `json:"sof_fation_name"`
	// SofHullName sof_hull_name string.
	// SofHullName SOF 船体名称 string.
	SofHullName string `json:"sof_hull_name"`
	// SofRaceName sof_race_name string.
	// SofRaceName SOF 种族名称 string.
	SofRaceName string `json:"sof_race_name"`
}

// GetUniverseGroupsGroupIdOk 200 ok object.
// GetUniverseGroupsGroupIdOk 200 ok 对象.
type GetUniverseGroupsGroupIdOk struct {
	// CategoryId category_id integer.
	// CategoryId 分类 ID 整数.
	CategoryId int32 `json:"category_id"`
	// GroupId group_id integer.
	// GroupId 分组 ID 整数.
	GroupId int32 `json:"group_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Published published boolean.
	// Published published 布尔.
	Published bool `json:"published"`
	// Types types array.
	// Types types 数组.
	Types []int32 `json:"types"`
}

// GetUniverseMoonsMoonIdOk 200 ok object.
// GetUniverseMoonsMoonIdOk 200 ok 对象.
type GetUniverseMoonsMoonIdOk struct {
	// MoonId moon_id integer.
	// MoonId moon_id 整数.
	MoonId int32 `json:"moon_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position GetUniverseMoonsMoonIdPosition `json:"position"`
	// SystemId The solar system this moon is in.
	// SystemId 该卫星所在的星系.
	SystemId int32 `json:"system_id"`
}

// GetUniverseMoonsMoonIdPosition position object.
// GetUniverseMoonsMoonIdPosition position 对象.
type GetUniverseMoonsMoonIdPosition struct {
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

// GetUniversePlanetsPlanetIdOk 200 ok object.
// GetUniversePlanetsPlanetIdOk 200 ok 对象.
type GetUniversePlanetsPlanetIdOk struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// PlanetId planet_id integer.
	// PlanetId planet_id 整数.
	PlanetId int32 `json:"planet_id"`
	// Position position object.
	// Position position 对象.
	Position GetUniversePlanetsPlanetIdPosition `json:"position"`
	// SystemId The solar system this planet is in.
	// SystemId 该行星所在的星系.
	SystemId int32 `json:"system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetUniversePlanetsPlanetIdPosition position object.
// GetUniversePlanetsPlanetIdPosition position 对象.
type GetUniversePlanetsPlanetIdPosition struct {
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

// GetUniverseRaces 200 ok object.
// GetUniverseRaces 200 ok 对象.
type GetUniverseRaces struct {
	// AllianceId The alliance generally associated with this race.
	// AllianceId 通常与该种族关联的联盟.
	AllianceId int32 `json:"alliance_id"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// RaceId race_id integer.
	// RaceId 种族ID integer.
	RaceId int32 `json:"race_id"`
}

// GetUniverseRegionsRegionIdOk 200 ok object.
// GetUniverseRegionsRegionIdOk 200 ok 对象.
type GetUniverseRegionsRegionIdOk struct {
	// Constellations constellations array.
	// Constellations 星座数组.
	Constellations []int32 `json:"constellations"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// RegionId region_id integer.
	// RegionId 星域ID integer.
	RegionId int32 `json:"region_id"`
}

// GetUniverseStargatesStargateIdDestination destination object.
// GetUniverseStargatesStargateIdDestination 目的地对象.
type GetUniverseStargatesStargateIdDestination struct {
	// StargateId The stargate this stargate connects to.
	// StargateId 此星门连接的星门.
	StargateId int32 `json:"stargate_id"`
	// SystemId The solar system this stargate connects to.
	// SystemId 此星门连接的星系.
	SystemId int32 `json:"system_id"`
}

// GetUniverseStargatesStargateIdOk 200 ok object.
// GetUniverseStargatesStargateIdOk 200 ok 对象.
type GetUniverseStargatesStargateIdOk struct {
	// Destination destination object.
	// Destination 目的地对象.
	Destination GetUniverseStargatesStargateIdDestination `json:"destination"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position GetUniverseStargatesStargateIdPosition `json:"position"`
	// StargateId stargate_id integer.
	// StargateId 星门ID integer.
	StargateId int32 `json:"stargate_id"`
	// SystemId The solar system this stargate is in.
	// SystemId 此星门所在的星系.
	SystemId int32 `json:"system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetUniverseStargatesStargateIdPosition position object.
// GetUniverseStargatesStargateIdPosition position 对象.
type GetUniverseStargatesStargateIdPosition struct {
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

// GetUniverseStarsStarIdOk 200 ok object.
// GetUniverseStarsStarIdOk 200 ok 对象.
type GetUniverseStarsStarIdOk struct {
	// Age Age of star in years.
	// Age 恒星的年龄（年）
	Age int64 `json:"age"`
	// Luminosity luminosity number.
	// Luminosity luminosity 数值.
	Luminosity float64 `json:"luminosity"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Radius radius integer.
	// Radius 半径 integer.
	Radius int64 `json:"radius"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// SpectralClass spectral_class string.
	// SpectralClass 光谱类型 string.
	// Enum values: "K2 V", "K4 V", "G2 V", "G8 V", "M7 V", "K7 V", "M2 V", "K5 V", "M3 V", "G0 V", "G7 V", "G3 V", "F9 V", "G5 V", "F6 V", "K8 V", "K9 V", "K6 V", "G9 V", "G6 V", "G4 VI", "G4 V", "F8 V", "F2 V", "F1 V", "K3 V", "F0 VI", "G1 VI", "G0 VI", "K1 V", "M4 V", "M1 V", "M6 V", "M0 V", "K2 IV", "G2 VI", "K0 V", "K5 IV", "F5 VI", "G6 VI", "F6 VI", "F2 IV", "G3 VI", "M8 V", "F1 VI", "K1 IV", "F7 V", "G5 VI", "M5 V", "G7 VI", "F5 V", "F4 VI", "F8 VI", "K3 IV", "F4 IV", "F0 V", "G7 IV", "G8 VI", "F2 VI", "F4 V", "F7 VI", "F3 V", "G1 V", "G9 VI", "F3 IV", "F9 VI", "M9 V", "K0 IV", "F1 IV", "G4 IV", "F3 VI", "K4 IV", "G5 IV", "G3 IV", "G1 IV", "K7 IV", "G0 IV", "K6 IV", "K9 IV", "G2 IV", "F9 IV", "F0 IV", "K8 IV", "G8 IV", "F6 IV", "F5 IV", "A0", "A0IV", "A0IV2".
	SpectralClass string `json:"spectral_class"`
	// Temperature temperature integer.
	// Temperature 温度整数.
	Temperature int32 `json:"temperature"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetUniverseStationsStationIdOk 200 ok object.
// GetUniverseStationsStationIdOk 200 ok 对象.
type GetUniverseStationsStationIdOk struct {
	// MaxDockableShipVolume max_dockable_ship_volume number.
	// MaxDockableShipVolume max_dockable_ship_volume 数值.
	MaxDockableShipVolume float64 `json:"max_dockable_ship_volume"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// OfficeRentalCost office_rental_cost number.
	// OfficeRentalCost office_rental_cost 数值.
	OfficeRentalCost float64 `json:"office_rental_cost"`
	// Owner ID of the corporation that controls this station.
	// Owner 控制此空间站的军团 ID.
	Owner int32 `json:"owner"`
	// Position position object.
	// Position position 对象.
	Position GetUniverseStationsStationIdPosition `json:"position"`
	// RaceId race_id integer.
	// RaceId 种族ID integer.
	RaceId int32 `json:"race_id"`
	// ReprocessingEfficiency reprocessing_efficiency number.
	// ReprocessingEfficiency 再处理效率 number.
	ReprocessingEfficiency float64 `json:"reprocessing_efficiency"`
	// ReprocessingStationsTake reprocessing_stations_take number.
	// ReprocessingStationsTake 空间站再处理抽成 number.
	ReprocessingStationsTake float64 `json:"reprocessing_stations_take"`
	// Services services array.
	// Services 服务列表 array.
	Services []string `json:"services"`
	// StationId station_id integer.
	// StationId 空间站ID integer.
	StationId int32 `json:"station_id"`
	// SystemId The solar system this station is in.
	// SystemId 此空间站所在的星系.
	SystemId int32 `json:"system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetUniverseStationsStationIdPosition position object.
// GetUniverseStationsStationIdPosition position 对象.
type GetUniverseStationsStationIdPosition struct {
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

// GetUniverseStructuresStructureIdOk 200 ok object.
// GetUniverseStructuresStructureIdOk 200 ok 对象.
type GetUniverseStructuresStructureIdOk struct {
	// Name The full name of the structure.
	// Name 建筑的全名.
	Name string `json:"name"`
	// OwnerId The ID of the corporation who owns this particular structure.
	// OwnerId 拥有此建筑的公司 ID.
	OwnerId int32 `json:"owner_id"`
	// Position Coordinates of the structure in Cartesian space relative to the Sun, in metres.
	// Position 建筑相对于太阳的笛卡尔空间坐标，单位为米。
	Position GetUniverseStructuresStructureIdPosition `json:"position"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetUniverseStructuresStructureIdPosition Coordinates of the structure in Cartesian space relative to the Sun, in metres.
// GetUniverseStructuresStructureIdPosition 建筑相对于太阳的笛卡尔空间坐标，单位为米。
type GetUniverseStructuresStructureIdPosition struct {
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

// GetUniverseSystemJumps 200 ok object.
// GetUniverseSystemJumps 200 ok 对象.
type GetUniverseSystemJumps struct {
	// ShipJumps ship_jumps integer.
	// ShipJumps 舰船跳跃次数 integer.
	ShipJumps int32 `json:"ship_jumps"`
	// SystemId system_id integer.
	// SystemId system_id 整数.
	SystemId int32 `json:"system_id"`
}

// GetUniverseSystemKills 200 ok object.
// GetUniverseSystemKills 200 ok 对象.
type GetUniverseSystemKills struct {
	// NpcKills Number of NPC ships killed in this system.
	// NpcKills 在该星系被击毁的 NPC 舰船数量.
	NpcKills int32 `json:"npc_kills"`
	// PodKills Number of pods killed in this system.
	// PodKills 在该星系被击毁的逃生舱数量.
	PodKills int32 `json:"pod_kills"`
	// ShipKills Number of player ships killed in this system.
	// ShipKills 在该星系被击毁的玩家舰船数量.
	ShipKills int32 `json:"ship_kills"`
	// SystemId system_id integer.
	// SystemId system_id 整数.
	SystemId int32 `json:"system_id"`
}

// GetUniverseSystemsSystemIdOk 200 ok object.
// GetUniverseSystemsSystemIdOk 200 ok 对象.
type GetUniverseSystemsSystemIdOk struct {
	// ConstellationId The constellation this solar system is in.
	// ConstellationId 该星系所在的星座.
	ConstellationId int32 `json:"constellation_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Planets planets array.
	// Planets planets 数组.
	Planets []GetUniverseSystemsSystemIdPlanet `json:"planets"`
	// Position position object.
	// Position position 对象.
	Position GetUniverseSystemsSystemIdPosition `json:"position"`
	// SecurityClass security_class string.
	// SecurityClass 安全等级 string.
	SecurityClass string `json:"security_class"`
	// SecurityStatus security_status number.
	// SecurityStatus 安全状态 number.
	SecurityStatus float64 `json:"security_status"`
	// StarId star_id integer.
	// StarId 恒星ID integer.
	StarId int32 `json:"star_id"`
	// Stargates stargates array.
	// Stargates 星门列表 array.
	Stargates []int32 `json:"stargates"`
	// Stations stations array.
	// Stations 空间站列表 array.
	Stations []int32 `json:"stations"`
	// SystemId system_id integer.
	// SystemId system_id 整数.
	SystemId int32 `json:"system_id"`
}

// GetUniverseSystemsSystemIdPlanet planet object.
// GetUniverseSystemsSystemIdPlanet planet 对象.
type GetUniverseSystemsSystemIdPlanet struct {
	// AsteroidBelts asteroid_belts array.
	// AsteroidBelts 小行星带数组.
	AsteroidBelts []int32 `json:"asteroid_belts"`
	// Moons moons array.
	// Moons moons 数组.
	Moons []int32 `json:"moons"`
	// PlanetId planet_id integer.
	// PlanetId planet_id 整数.
	PlanetId int32 `json:"planet_id"`
}

// GetUniverseSystemsSystemIdPosition position object.
// GetUniverseSystemsSystemIdPosition position 对象.
type GetUniverseSystemsSystemIdPosition struct {
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

// GetUniverseTypesTypeIdDogmaAttribute dogma_attribute object.
// GetUniverseTypesTypeIdDogmaAttribute 教条属性对象.
type GetUniverseTypesTypeIdDogmaAttribute struct {
	// AttributeId attribute_id integer.
	// AttributeId 属性 ID 整数.
	AttributeId int32 `json:"attribute_id"`
	// Value value number.
	// Value 价值数字.
	Value float64 `json:"value"`
}

// GetUniverseTypesTypeIdDogmaEffect dogma_effect object.
// GetUniverseTypesTypeIdDogmaEffect 教条效果对象.
type GetUniverseTypesTypeIdDogmaEffect struct {
	// EffectId effect_id integer.
	// EffectId 效果 ID 整数.
	EffectId int32 `json:"effect_id"`
	// IsDefault is_default boolean.
	// IsDefault 是否为默认布尔值.
	IsDefault bool `json:"is_default"`
}

// GetUniverseTypesTypeIdOk 200 ok object.
// GetUniverseTypesTypeIdOk 200 ok 对象.
type GetUniverseTypesTypeIdOk struct {
	// Capacity capacity number.
	// Capacity 容量数值.
	Capacity float64 `json:"capacity"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// DogmaAttributes dogma_attributes array.
	// DogmaAttributes 教条属性数组.
	DogmaAttributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes"`
	// DogmaEffects dogma_effects array.
	// DogmaEffects 教条效果数组.
	DogmaEffects []GetUniverseTypesTypeIdDogmaEffect `json:"dogma_effects"`
	// GraphicId graphic_id integer.
	// GraphicId 图形 ID 整数.
	GraphicId int32 `json:"graphic_id"`
	// GroupId group_id integer.
	// GroupId 分组 ID 整数.
	GroupId int32 `json:"group_id"`
	// IconId icon_id integer.
	// IconId 图标 ID 整数.
	IconId int32 `json:"icon_id"`
	// MarketGroupId This only exists for types that can be put on the market.
	// MarketGroupId 仅对可上架市场的类型存在此项.
	MarketGroupId int32 `json:"market_group_id"`
	// Mass mass number.
	// Mass mass 数值.
	Mass float64 `json:"mass"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// PackagedVolume packaged_volume number.
	// PackagedVolume packaged_volume 数值.
	PackagedVolume float64 `json:"packaged_volume"`
	// PortionSize portion_size integer.
	// PortionSize portion_size 整数.
	PortionSize int32 `json:"portion_size"`
	// Published published boolean.
	// Published published 布尔.
	Published bool `json:"published"`
	// Radius radius number.
	// Radius 半径 number.
	Radius float64 `json:"radius"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
	// Volume volume number.
	// Volume 体积数字.
	Volume float64 `json:"volume"`
}

// PostUniverseIdsAgent agent object.
// PostUniverseIdsAgent 代理人对象.
type PostUniverseIdsAgent struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsAlliance alliance object.
// PostUniverseIdsAlliance 联盟对象.
type PostUniverseIdsAlliance struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsCharacter character object.
// PostUniverseIdsCharacter 角色对象.
type PostUniverseIdsCharacter struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsConstellation constellation object.
// PostUniverseIdsConstellation 星座对象.
type PostUniverseIdsConstellation struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsCorporation corporation object.
// PostUniverseIdsCorporation 军团对象.
type PostUniverseIdsCorporation struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsFaction faction object.
// PostUniverseIdsFaction 势力对象.
type PostUniverseIdsFaction struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsInventoryType inventory_type object.
// PostUniverseIdsInventoryType 物品类型对象.
type PostUniverseIdsInventoryType struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsOk 200 ok object.
// PostUniverseIdsOk 200 ok 对象.
type PostUniverseIdsOk struct {
	// Agents agents array.
	// Agents 代理人数组.
	Agents []PostUniverseIdsAgent `json:"agents"`
	// Alliances alliances array.
	// Alliances 联盟数组.
	Alliances []PostUniverseIdsAlliance `json:"alliances"`
	// Characters characters array.
	// Characters 角色数组.
	Characters []PostUniverseIdsCharacter `json:"characters"`
	// Constellations constellations array.
	// Constellations 星座数组.
	Constellations []PostUniverseIdsConstellation `json:"constellations"`
	// Corporations corporations array.
	// Corporations 军团数组.
	Corporations []PostUniverseIdsCorporation `json:"corporations"`
	// Factions factions array.
	// Factions 势力数组.
	Factions []PostUniverseIdsFaction `json:"factions"`
	// InventoryTypes inventory_types array.
	// InventoryTypes 物品类型数组.
	InventoryTypes []PostUniverseIdsInventoryType `json:"inventory_types"`
	// Regions regions array.
	// Regions 星域列表 array.
	Regions []PostUniverseIdsRegion `json:"regions"`
	// Stations stations array.
	// Stations 空间站列表 array.
	Stations []PostUniverseIdsStation `json:"stations"`
	// Systems systems array.
	// Systems systems 数组.
	Systems []PostUniverseIdsSystem `json:"systems"`
}

// PostUniverseIdsRegion region object.
// PostUniverseIdsRegion 星域 object.
type PostUniverseIdsRegion struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsStation station object.
// PostUniverseIdsStation 空间站 object.
type PostUniverseIdsStation struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseIdsSystem system object.
// PostUniverseIdsSystem 星系 object.
type PostUniverseIdsSystem struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// PostUniverseNames 200 ok object.
// PostUniverseNames 200 ok 对象.
type PostUniverseNames struct {
	// Category category string.
	// Category 分类字符串.
	// Enum values: "alliance", "character", "constellation", "corporation", "inventory_type", "region", "solar_system", "station", "faction".
	Category string `json:"category"`
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// GetAncestriesParams holds the optional query and header parameters of the request.
// GetAncestriesParams 保存请求的可选查询与头部参数。
type GetAncestriesParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetAncestriesParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetAsteroidBeltsAsteroidBeltIdParams holds the optional query and header parameters of the request.
// GetAsteroidBeltsAsteroidBeltIdParams 保存请求的可选查询与头部参数。
type GetAsteroidBeltsAsteroidBeltIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetAsteroidBeltsAsteroidBeltIdParams) Values() (url.Values, map[string]string) {
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

// GetBloodlinesParams holds the optional query and header parameters of the request.
// GetBloodlinesParams 保存请求的可选查询与头部参数。
type GetBloodlinesParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetBloodlinesParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetCategoriesCategoryIdParams holds the optional query and header parameters of the request.
// GetCategoriesCategoryIdParams 保存请求的可选查询与头部参数。
type GetCategoriesCategoryIdParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetCategoriesCategoryIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetCategoriesParams holds the optional query and header parameters of the request.
// GetCategoriesParams 保存请求的可选查询与头部参数。
type GetCategoriesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetCategoriesParams) Values() (url.Values, map[string]string) {
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

// GetConstellationsConstellationIdParams holds the optional query and header parameters of the request.
// GetConstellationsConstellationIdParams 保存请求的可选查询与头部参数。
type GetConstellationsConstellationIdParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetConstellationsConstellationIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetConstellationsParams holds the optional query and header parameters of the request.
// GetConstellationsParams 保存请求的可选查询与头部参数。
type GetConstellationsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetConstellationsParams) Values() (url.Values, map[string]string) {
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

// GetFactionsParams holds the optional query and header parameters of the request.
// GetFactionsParams 保存请求的可选查询与头部参数。
type GetFactionsParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetFactionsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetGraphicsGraphicIdParams holds the optional query and header parameters of the request.
// GetGraphicsGraphicIdParams 保存请求的可选查询与头部参数。
type GetGraphicsGraphicIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetGraphicsGraphicIdParams) Values() (url.Values, map[string]string) {
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

// GetGraphicsParams holds the optional query and header parameters of the request.
// GetGraphicsParams 保存请求的可选查询与头部参数。
type GetGraphicsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetGraphicsParams) Values() (url.Values, map[string]string) {
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

// GetGroupsGroupIdParamsX holds the optional query and header parameters of the request.
// GetGroupsGroupIdParamsX 保存请求的可选查询与头部参数。
type GetGroupsGroupIdParamsX struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetGroupsGroupIdParamsX) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetGroupsParamsX holds the optional query and header parameters of the request.
// GetGroupsParamsX 保存请求的可选查询与头部参数。
type GetGroupsParamsX struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
}

func (p *GetGroupsParamsX) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetMoonsMoonIdParams holds the optional query and header parameters of the request.
// GetMoonsMoonIdParams 保存请求的可选查询与头部参数。
type GetMoonsMoonIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetMoonsMoonIdParams) Values() (url.Values, map[string]string) {
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

// GetPlanetsPlanetIdParams holds the optional query and header parameters of the request.
// GetPlanetsPlanetIdParams 保存请求的可选查询与头部参数。
type GetPlanetsPlanetIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetPlanetsPlanetIdParams) Values() (url.Values, map[string]string) {
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

// GetRacesParams holds the optional query and header parameters of the request.
// GetRacesParams 保存请求的可选查询与头部参数。
type GetRacesParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetRacesParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetRegionsParams holds the optional query and header parameters of the request.
// GetRegionsParams 保存请求的可选查询与头部参数。
type GetRegionsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetRegionsParams) Values() (url.Values, map[string]string) {
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

// GetRegionsRegionIdParams holds the optional query and header parameters of the request.
// GetRegionsRegionIdParams 保存请求的可选查询与头部参数。
type GetRegionsRegionIdParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetRegionsRegionIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetStargatesStargateIdParams holds the optional query and header parameters of the request.
// GetStargatesStargateIdParams 保存请求的可选查询与头部参数。
type GetStargatesStargateIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetStargatesStargateIdParams) Values() (url.Values, map[string]string) {
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

// GetStarsStarIdParams holds the optional query and header parameters of the request.
// GetStarsStarIdParams 保存请求的可选查询与头部参数。
type GetStarsStarIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetStarsStarIdParams) Values() (url.Values, map[string]string) {
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

// GetStationsStationIdParams holds the optional query and header parameters of the request.
// GetStationsStationIdParams 保存请求的可选查询与头部参数。
type GetStationsStationIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetStationsStationIdParams) Values() (url.Values, map[string]string) {
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

// GetStructuresParamsX holds the optional query and header parameters of the request.
// GetStructuresParamsX 保存请求的可选查询与头部参数。
type GetStructuresParamsX struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Filter Only list public structures that have this service online.
	// Filter 仅列出该服务在线的公共建筑.
	Filter *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetStructuresParamsX) Values() (url.Values, map[string]string) {
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
	if p.Filter != nil && *p.Filter != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("filter", *p.Filter)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	return query, headers
}

// GetStructuresStructureIdParams holds the optional query and header parameters of the request.
// GetStructuresStructureIdParams 保存请求的可选查询与头部参数。
type GetStructuresStructureIdParams struct {
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

func (p *GetStructuresStructureIdParams) Values() (url.Values, map[string]string) {
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

// GetSystemJumpsParams holds the optional query and header parameters of the request.
// GetSystemJumpsParams 保存请求的可选查询与头部参数。
type GetSystemJumpsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetSystemJumpsParams) Values() (url.Values, map[string]string) {
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

// GetSystemKillsParams holds the optional query and header parameters of the request.
// GetSystemKillsParams 保存请求的可选查询与头部参数。
type GetSystemKillsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetSystemKillsParams) Values() (url.Values, map[string]string) {
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

// GetSystemsParamsX holds the optional query and header parameters of the request.
// GetSystemsParamsX 保存请求的可选查询与头部参数。
type GetSystemsParamsX struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetSystemsParamsX) Values() (url.Values, map[string]string) {
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

// GetSystemsSystemIdParams holds the optional query and header parameters of the request.
// GetSystemsSystemIdParams 保存请求的可选查询与头部参数。
type GetSystemsSystemIdParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetSystemsSystemIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// GetTypesParams holds the optional query and header parameters of the request.
// GetTypesParams 保存请求的可选查询与头部参数。
type GetTypesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
}

func (p *GetTypesParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}

// GetTypesTypeIdParams holds the optional query and header parameters of the request.
// GetTypesTypeIdParams 保存请求的可选查询与头部参数。
type GetTypesTypeIdParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetTypesTypeIdParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// PostIdsParams holds the optional query and header parameters of the request.
// PostIdsParams 保存请求的可选查询与头部参数。
type PostIdsParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *PostIdsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}

// PostNamesParams holds the optional query and header parameters of the request.
// PostNamesParams 保存请求的可选查询与头部参数。
type PostNamesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
}

func (p *PostNamesParams) Values() (url.Values, map[string]string) {
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
