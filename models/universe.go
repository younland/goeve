package models

// UniverseAncestry 200 ok object.
// UniverseAncestry 200 ok 对象.
type UniverseAncestry struct {
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

// UniverseAsteroidBelt 200 ok object.
// UniverseAsteroidBelt 200 ok 对象.
type UniverseAsteroidBelt struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position AsteroidBeltPosition `json:"position"`
	// SystemId The solar system this asteroid belt is in.
	// SystemId 该小行星带所在的星系.
	SystemId int32 `json:"system_id"`
}

// AsteroidBeltPosition position object.
// AsteroidBeltPosition position 对象.
type AsteroidBeltPosition struct {
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

// UniverseBloodline 200 ok object.
// UniverseBloodline 200 ok 对象.
type UniverseBloodline struct {
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

// UniverseCategory 200 ok object.
// UniverseCategory 200 ok 对象.
type UniverseCategory struct {
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

// UniverseConstellation 200 ok object.
// UniverseConstellation 200 ok 对象.
type UniverseConstellation struct {
	// ConstellationId constellation_id integer.
	// ConstellationId 星座 ID 整数.
	ConstellationId int32 `json:"constellation_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position ConstellationPosition `json:"position"`
	// RegionId The region this constellation is in.
	// RegionId 该星座所在的星域.
	RegionId int32 `json:"region_id"`
	// Systems systems array.
	// Systems systems 数组.
	Systems []int32 `json:"systems"`
}

// ConstellationPosition position object.
// ConstellationPosition position 对象.
type ConstellationPosition struct {
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

// UniverseFaction 200 ok object.
// UniverseFaction 200 ok 对象.
type UniverseFaction struct {
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

// UniverseGraphic 200 ok object.
// UniverseGraphic 200 ok 对象.
type UniverseGraphic struct {
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

// UniverseGroup 200 ok object.
// UniverseGroup 200 ok 对象.
type UniverseGroup struct {
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

// UniverseMoon 200 ok object.
// UniverseMoon 200 ok 对象.
type UniverseMoon struct {
	// MoonId moon_id integer.
	// MoonId moon_id 整数.
	MoonId int32 `json:"moon_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position MoonPosition `json:"position"`
	// SystemId The solar system this moon is in.
	// SystemId 该卫星所在的星系.
	SystemId int32 `json:"system_id"`
}

// MoonPosition position object.
// MoonPosition position 对象.
type MoonPosition struct {
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

// UniversePlanet 200 ok object.
// UniversePlanet 200 ok 对象.
type UniversePlanet struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// PlanetId planet_id integer.
	// PlanetId planet_id 整数.
	PlanetId int32 `json:"planet_id"`
	// Position position object.
	// Position position 对象.
	Position PlanetPosition `json:"position"`
	// SystemId The solar system this planet is in.
	// SystemId 该行星所在的星系.
	SystemId int32 `json:"system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// PlanetPosition position object.
// PlanetPosition position 对象.
type PlanetPosition struct {
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

// UniverseRace 200 ok object.
// UniverseRace 200 ok 对象.
type UniverseRace struct {
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

// UniverseRegion 200 ok object.
// UniverseRegion 200 ok 对象.
type UniverseRegion struct {
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

// StargateDestination destination object.
// StargateDestination 目的地对象.
type StargateDestination struct {
	// StargateId The stargate this stargate connects to.
	// StargateId 此星门连接的星门.
	StargateId int32 `json:"stargate_id"`
	// SystemId The solar system this stargate connects to.
	// SystemId 此星门连接的星系.
	SystemId int32 `json:"system_id"`
}

// UniverseStargate 200 ok object.
// UniverseStargate 200 ok 对象.
type UniverseStargate struct {
	// Destination destination object.
	// Destination 目的地对象.
	Destination StargateDestination `json:"destination"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Position position object.
	// Position position 对象.
	Position StargatePosition `json:"position"`
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

// StargatePosition position object.
// StargatePosition position 对象.
type StargatePosition struct {
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

// UniverseStar 200 ok object.
// UniverseStar 200 ok 对象.
type UniverseStar struct {
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

// UniverseStation 200 ok object.
// UniverseStation 200 ok 对象.
type UniverseStation struct {
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
	Position StationPosition `json:"position"`
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

// StationPosition position object.
// StationPosition position 对象.
type StationPosition struct {
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

// UniverseStructure 200 ok object.
// UniverseStructure 200 ok 对象.
type UniverseStructure struct {
	// Name The full name of the structure.
	// Name 建筑的全名.
	Name string `json:"name"`
	// OwnerId The ID of the corporation who owns this particular structure.
	// OwnerId 拥有此建筑的公司 ID.
	OwnerId int32 `json:"owner_id"`
	// Position Coordinates of the structure in Cartesian space relative to the Sun, in metres.
	// Position 建筑相对于太阳的笛卡尔空间坐标，单位为米。
	Position StructurePosition `json:"position"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// StructurePosition Coordinates of the structure in Cartesian space relative to the Sun, in metres.
// StructurePosition 建筑相对于太阳的笛卡尔空间坐标，单位为米。
type StructurePosition struct {
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

// UniverseSystemJump 200 ok object.
// UniverseSystemJump 200 ok 对象.
type UniverseSystemJump struct {
	// ShipJumps ship_jumps integer.
	// ShipJumps 舰船跳跃次数 integer.
	ShipJumps int32 `json:"ship_jumps"`
	// SystemId system_id integer.
	// SystemId system_id 整数.
	SystemId int32 `json:"system_id"`
}

// UniverseSystemKills 200 ok object.
// UniverseSystemKills 200 ok 对象.
type UniverseSystemKills struct {
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

// UniverseSolarSystem 200 ok object.
// UniverseSolarSystem 200 ok 对象.
type UniverseSolarSystem struct {
	// ConstellationId The constellation this solar system is in.
	// ConstellationId 该星系所在的星座.
	ConstellationId int32 `json:"constellation_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Planets planets array.
	// Planets planets 数组.
	Planets []SolarSystemPlanet `json:"planets"`
	// Position position object.
	// Position position 对象.
	Position SolarSystemPosition `json:"position"`
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

// SolarSystemPlanet planet object.
// SolarSystemPlanet planet 对象.
type SolarSystemPlanet struct {
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

// SolarSystemPosition position object.
// SolarSystemPosition position 对象.
type SolarSystemPosition struct {
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

// UniverseType 200 ok object.
// UniverseType 200 ok 对象.
type UniverseType struct {
	// Capacity capacity number.
	// Capacity 容量数值.
	Capacity float64 `json:"capacity"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// DogmaAttributes dogma_attributes array.
	// DogmaAttributes 教条属性数组.
	DogmaAttributes []DogmaAttributeValue `json:"dogma_attributes"`
	// DogmaEffects dogma_effects array.
	// DogmaEffects 教条效果数组.
	DogmaEffects []DogmaEffectValue `json:"dogma_effects"`
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

// UniverseIdName agent object.
// UniverseIdName 代理人对象.
type UniverseIdName struct {
	// Id id integer.
	// Id ID 整数.
	Id int32 `json:"id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// ResolvedIds 200 ok object.
// ResolvedIds 200 ok 对象.
type ResolvedIds struct {
	// Agents agents array.
	// Agents 代理人数组.
	Agents []UniverseIdName `json:"agents"`
	// Alliances alliances array.
	// Alliances 联盟数组.
	Alliances []UniverseIdName `json:"alliances"`
	// Characters characters array.
	// Characters 角色数组.
	Characters []UniverseIdName `json:"characters"`
	// Constellations constellations array.
	// Constellations 星座数组.
	Constellations []UniverseIdName `json:"constellations"`
	// Corporations corporations array.
	// Corporations 军团数组.
	Corporations []UniverseIdName `json:"corporations"`
	// Factions factions array.
	// Factions 势力数组.
	Factions []UniverseIdName `json:"factions"`
	// InventoryTypes inventory_types array.
	// InventoryTypes 物品类型数组.
	InventoryTypes []UniverseIdName `json:"inventory_types"`
	// Regions regions array.
	// Regions 星域列表 array.
	Regions []UniverseIdName `json:"regions"`
	// Stations stations array.
	// Stations 空间站列表 array.
	Stations []UniverseIdName `json:"stations"`
	// Systems systems array.
	// Systems systems 数组.
	Systems []UniverseIdName `json:"systems"`
}

// UniverseName 200 ok object.
// UniverseName 200 ok 对象.
type UniverseName struct {
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
