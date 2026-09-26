package models

import (
	"time"
)

// AllianceHistoryEntry 200 ok object.
// AllianceHistoryEntry 200 ok 对象.
type AllianceHistoryEntry struct {
	// AllianceId alliance_id integer.
	// AllianceId 联盟 ID 整数.
	AllianceId int32 `json:"alliance_id"`
	// IsDeleted True if the alliance has been closed.
	// IsDeleted 如果联盟已关闭则为 true.
	IsDeleted bool `json:"is_deleted"`
	// RecordId An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous.
	// RecordId 一个递增的 ID，可用于在日期可能含糊不清的情况下规范地确定记录的顺序.
	RecordId int32 `json:"record_id"`
	// StartDate start_date string.
	// StartDate 开始日期 string.
	StartDate time.Time `json:"start_date"`
}

// ContainerLog 200 ok object.
// ContainerLog 200 ok 对象.
type ContainerLog struct {
	// Action action string.
	// Action 操作字符串.
	// Enum values: "add", "assemble", "configure", "enter_password", "lock", "move", "repackage", "set_name", "set_password", "unlock".
	Action string `json:"action"`
	// CharacterId ID of the character who performed the action.
	// CharacterId 执行此操作的角色 ID。
	CharacterId int32 `json:"character_id"`
	// ContainerId ID of the container.
	// ContainerId 容器 ID.
	ContainerId int64 `json:"container_id"`
	// ContainerTypeId Type ID of the container.
	// ContainerTypeId 容器的 type ID.
	ContainerTypeId int32 `json:"container_type_id"`
	// LocationFlag location_flag string.
	// LocationFlag location_flag 字符串.
	// Enum values: "AssetSafety", "AutoFit", "Bonus", "Booster", "BoosterBay", "Capsule", "Cargo", "CorpDeliveries", "CorpSAG1", "CorpSAG2", "CorpSAG3", "CorpSAG4", "CorpSAG5", "CorpSAG6", "CorpSAG7", "CrateLoot", "Deliveries", "DroneBay", "DustBattle", "DustDatabank", "FighterBay", "FighterTube0", "FighterTube1", "FighterTube2", "FighterTube3", "FighterTube4", "FleetHangar", "FrigateEscapeBay", "Hangar", "HangarAll", "HiSlot0", "HiSlot1", "HiSlot2", "HiSlot3", "HiSlot4", "HiSlot5", "HiSlot6", "HiSlot7", "HiddenModifiers", "Implant", "Impounded", "JunkyardReprocessed", "JunkyardTrashed", "LoSlot0", "LoSlot1", "LoSlot2", "LoSlot3", "LoSlot4", "LoSlot5", "LoSlot6", "LoSlot7", "Locked", "MedSlot0", "MedSlot1", "MedSlot2", "MedSlot3", "MedSlot4", "MedSlot5", "MedSlot6", "MedSlot7", "OfficeFolder", "Pilot", "PlanetSurface", "QuafeBay", "QuantumCoreRoom", "Reward", "RigSlot0", "RigSlot1", "RigSlot2", "RigSlot3", "RigSlot4", "RigSlot5", "RigSlot6", "RigSlot7", "SecondaryStorage", "ServiceSlot0", "ServiceSlot1", "ServiceSlot2", "ServiceSlot3", "ServiceSlot4", "ServiceSlot5", "ServiceSlot6", "ServiceSlot7", "ShipHangar", "ShipOffline", "Skill", "SkillInTraining", "SpecializedAmmoHold", "SpecializedCommandCenterHold", "SpecializedFuelBay", "SpecializedGasHold", "SpecializedIndustrialShipHold", "SpecializedLargeShipHold", "SpecializedMaterialBay", "SpecializedMediumShipHold", "SpecializedMineralHold", "SpecializedOreHold", "SpecializedPlanetaryCommoditiesHold", "SpecializedSalvageHold", "SpecializedShipHold", "SpecializedSmallShipHold", "StructureActive", "StructureFuel", "StructureInactive", "StructureOffline", "SubSystemBay", "SubSystemSlot0", "SubSystemSlot1", "SubSystemSlot2", "SubSystemSlot3", "SubSystemSlot4", "SubSystemSlot5", "SubSystemSlot6", "SubSystemSlot7", "Unlocked", "Wallet", "Wardrobe".
	LocationFlag string `json:"location_flag"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LoggedAt Timestamp when this log was created.
	// LoggedAt 此日志创建的时间戳.
	LoggedAt time.Time `json:"logged_at"`
	// NewConfigBitmask new_config_bitmask integer.
	// NewConfigBitmask new_config_bitmask 整数.
	NewConfigBitmask int32 `json:"new_config_bitmask"`
	// OldConfigBitmask old_config_bitmask integer.
	// OldConfigBitmask old_config_bitmask 整数.
	OldConfigBitmask int32 `json:"old_config_bitmask"`
	// PasswordType Type of password set if action is of type SetPassword or EnterPassword.
	// PasswordType 当动作为 SetPassword 或 EnterPassword 类型时所设置的密码类型.
	// Enum values: "config", "general".
	PasswordType string `json:"password_type"`
	// Quantity Quantity of the item being acted upon.
	// Quantity 正在处理的物品的数量.
	Quantity int32 `json:"quantity"`
	// TypeId Type ID of the item being acted upon.
	// TypeId 被操作物品的 type ID.
	TypeId int32 `json:"type_id"`
}

// Division hangar object.
// Division 机库对象.
type Division struct {
	// Division division integer.
	// Division 部门整数.
	Division int32 `json:"division"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// CorporationDivisions 200 ok object.
// CorporationDivisions 200 ok 对象.
type CorporationDivisions struct {
	// Hangar hangar array.
	// Hangar 机库数组.
	Hangar []Division `json:"hangar"`
	// Wallet wallet array.
	// Wallet 钱包数组.
	Wallet []Division `json:"wallet"`
}

// CorporationFacility 200 ok object.
// CorporationFacility 200 ok 对象.
type CorporationFacility struct {
	// FacilityId facility_id integer.
	// FacilityId 设施 ID 整数.
	FacilityId int64 `json:"facility_id"`
	// SystemId system_id integer.
	// SystemId system_id 整数.
	SystemId int32 `json:"system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// CorporationIcons 200 ok object.
// CorporationIcons 200 ok 对象.
type CorporationIcons struct {
	// Px128x128 px128x128 string.
	// Px128x128 px128x128 字符串.
	Px128x128 string `json:"px128x128"`
	// Px256x256 px256x256 string.
	// Px256x256 px256x256 字符串.
	Px256x256 string `json:"px256x256"`
	// Px64x64 px64x64 string.
	// Px64x64 px64x64 字符串.
	Px64x64 string `json:"px64x64"`
}

// CorporationMedal 200 ok object.
// CorporationMedal 200 ok 对象.
type CorporationMedal struct {
	// CreatedAt created_at string.
	// CreatedAt 创建时间字符串.
	CreatedAt time.Time `json:"created_at"`
	// CreatorId ID of the character who created this medal.
	// CreatorId 创建此勋章的角色 ID.
	CreatorId int32 `json:"creator_id"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// MedalId medal_id integer.
	// MedalId medal_id 整数.
	MedalId int32 `json:"medal_id"`
	// Title title string.
	// Title 标题字符串.
	Title string `json:"title"`
}

// IssuedMedal 200 ok object.
// IssuedMedal 200 ok 对象.
type IssuedMedal struct {
	// CharacterId ID of the character who was rewarded this medal.
	// CharacterId 被授予此勋章的角色 ID.
	CharacterId int32 `json:"character_id"`
	// IssuedAt issued_at string.
	// IssuedAt 发布时间字符串.
	IssuedAt time.Time `json:"issued_at"`
	// IssuerId ID of the character who issued the medal.
	// IssuerId 颁发此勋章的角色 ID.
	IssuerId int32 `json:"issuer_id"`
	// MedalId medal_id integer.
	// MedalId medal_id 整数.
	MedalId int32 `json:"medal_id"`
	// Reason reason string.
	// Reason 原因 string.
	Reason string `json:"reason"`
	// Status status string.
	// Status 状态 string.
	// Enum values: "private", "public".
	Status string `json:"status"`
}

// MemberTitles 200 ok object.
// MemberTitles 200 ok 对象.
type MemberTitles struct {
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// Titles A list of title_id.
	// Titles title_id 列表.
	Titles []int32 `json:"titles"`
}

// MemberTrackingEntry 200 ok object.
// MemberTrackingEntry 200 ok 对象.
type MemberTrackingEntry struct {
	// BaseId base_id integer.
	// BaseId 基地 ID 整数.
	BaseId int32 `json:"base_id"`
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// LogoffDate logoff_date string.
	// LogoffDate logoff_date 字符串.
	LogoffDate time.Time `json:"logoff_date"`
	// LogonDate logon_date string.
	// LogonDate logon_date 字符串.
	LogonDate time.Time `json:"logon_date"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
	// StartDate start_date string.
	// StartDate 开始日期 string.
	StartDate time.Time `json:"start_date"`
}

// Corporation 200 ok object.
// Corporation 200 ok 对象.
type Corporation struct {
	// AllianceId ID of the alliance that corporation is a member of, if any.
	// AllianceId 该军团所属联盟的 ID（如有）
	AllianceId int32 `json:"alliance_id"`
	// CeoId ceo_id integer.
	// CeoId CEO ID 整数.
	CeoId int32 `json:"ceo_id"`
	// CreatorId creator_id integer.
	// CreatorId 创建者 ID 整数.
	CreatorId int32 `json:"creator_id"`
	// DateFounded date_founded string.
	// DateFounded 成立日期字符串.
	DateFounded time.Time `json:"date_founded"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// FactionId faction_id integer.
	// FactionId 势力 ID 整数.
	FactionId int32 `json:"faction_id"`
	// HomeStationId home_station_id integer.
	// HomeStationId 常驻空间站 ID 整数.
	HomeStationId int32 `json:"home_station_id"`
	// MemberCount member_count integer.
	// MemberCount member_count 整数.
	MemberCount int32 `json:"member_count"`
	// Name the full name of the corporation.
	// Name 军团的完整名称.
	Name string `json:"name"`
	// Shares shares integer.
	// Shares 股份 integer.
	Shares int64 `json:"shares"`
	// TaxRate tax_rate number.
	// TaxRate 税率数字.
	TaxRate float64 `json:"tax_rate"`
	// Ticker the short name of the corporation.
	// Ticker 军团的简称.
	Ticker string `json:"ticker"`
	// Url url string.
	// Url URL 字符串.
	Url string `json:"url"`
	// WarEligible war_eligible boolean.
	// WarEligible 可参战（war eligible）布尔值.
	WarEligible bool `json:"war_eligible"`
}

// CorporationMemberRoles 200 ok object.
// CorporationMemberRoles 200 ok 对象.
type CorporationMemberRoles struct {
	// CharacterId character_id integer.
	// CharacterId 角色 ID 整数.
	CharacterId int32 `json:"character_id"`
	// GrantableRoles grantable_roles array.
	// GrantableRoles 可授予角色数组.
	GrantableRoles []string `json:"grantable_roles"`
	// GrantableRolesAtBase grantable_roles_at_base array.
	// GrantableRolesAtBase 基地可授予角色数组.
	GrantableRolesAtBase []string `json:"grantable_roles_at_base"`
	// GrantableRolesAtHq grantable_roles_at_hq array.
	// GrantableRolesAtHq 总部可授予角色数组.
	GrantableRolesAtHq []string `json:"grantable_roles_at_hq"`
	// GrantableRolesAtOther grantable_roles_at_other array.
	// GrantableRolesAtOther 其他地点可授予角色数组.
	GrantableRolesAtOther []string `json:"grantable_roles_at_other"`
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

// CorporationRoleHistory 200 ok object.
// CorporationRoleHistory 200 ok 对象.
type CorporationRoleHistory struct {
	// ChangedAt changed_at string.
	// ChangedAt 变更时间字符串.
	ChangedAt time.Time `json:"changed_at"`
	// CharacterId The character whose roles are changed.
	// CharacterId 角色被更改的角色.
	CharacterId int32 `json:"character_id"`
	// IssuerId ID of the character who issued this change.
	// IssuerId 发起此变更的角色 ID.
	IssuerId int32 `json:"issuer_id"`
	// NewRoles new_roles array.
	// NewRoles new_roles 数组.
	NewRoles []string `json:"new_roles"`
	// OldRoles old_roles array.
	// OldRoles old_roles 数组.
	OldRoles []string `json:"old_roles"`
	// RoleType role_type string.
	// RoleType 角色类型 string.
	// Enum values: "grantable_roles", "grantable_roles_at_base", "grantable_roles_at_hq", "grantable_roles_at_other", "roles", "roles_at_base", "roles_at_hq", "roles_at_other".
	RoleType string `json:"role_type"`
}

// Shareholder 200 ok object.
// Shareholder 200 ok 对象.
type Shareholder struct {
	// ShareCount share_count integer.
	// ShareCount 股份数量 integer.
	ShareCount int64 `json:"share_count"`
	// ShareholderId shareholder_id integer.
	// ShareholderId 股东ID integer.
	ShareholderId int32 `json:"shareholder_id"`
	// ShareholderType shareholder_type string.
	// ShareholderType 股东类型 string.
	// Enum values: "character", "corporation".
	ShareholderType string `json:"shareholder_type"`
}

// Starbase 200 ok object.
// Starbase 200 ok 对象.
type Starbase struct {
	// MoonId The moon this starbase (POS) is anchored on, unanchored POSes do not have this information.
	// MoonId 此母星基地（POS）锚定的卫星，未锚定的 POS 无此信息.
	MoonId int32 `json:"moon_id"`
	// OnlinedSince When the POS onlined, for starbases (POSes) in online state.
	// OnlinedSince POS 上线的时间，适用于处于上线状态的母星基地（POS）
	OnlinedSince time.Time `json:"onlined_since"`
	// ReinforcedUntil When the POS will be out of reinforcement, for starbases (POSes) in reinforced state.
	// ReinforcedUntil 处于增强状态的母星（POS）何时脱离增强.
	ReinforcedUntil time.Time `json:"reinforced_until"`
	// StarbaseId Unique ID for this starbase (POS)
	// StarbaseId 此母星基地（POS）的唯一 ID.
	StarbaseId int64 `json:"starbase_id"`
	// State state string.
	// State 状态 string.
	// Enum values: "offline", "online", "onlining", "reinforced", "unanchoring".
	State string `json:"state"`
	// SystemId The solar system this starbase (POS) is in, unanchored POSes have this information.
	// SystemId 此母星基地（POS）所在的星系，未锚定的 POS 有此信息.
	SystemId int32 `json:"system_id"`
	// TypeId Starbase (POS) type.
	// TypeId 母星基地（POS）类型.
	TypeId int32 `json:"type_id"`
	// UnanchorAt When the POS started unanchoring, for starbases (POSes) in unanchoring state.
	// UnanchorAt POS 开始解锚的时间，适用于处于解锚状态的母星基地（POS）
	UnanchorAt time.Time `json:"unanchor_at"`
}

// StarbaseFuel fuel object.
// StarbaseFuel 燃料对象.
type StarbaseFuel struct {
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// StarbaseDetail 200 ok object.
// StarbaseDetail 200 ok 对象.
type StarbaseDetail struct {
	// AllowAllianceMembers allow_alliance_members boolean.
	// AllowAllianceMembers 是否允许联盟成员布尔值.
	AllowAllianceMembers bool `json:"allow_alliance_members"`
	// AllowCorporationMembers allow_corporation_members boolean.
	// AllowCorporationMembers 是否允许军团成员布尔值.
	AllowCorporationMembers bool `json:"allow_corporation_members"`
	// Anchor Who can anchor starbase (POS) and its structures.
	// Anchor 谁可以锚定母星（POS）及其建筑.
	// Enum values: "alliance_member", "config_starbase_equipment_role", "corporation_member", "starbase_fuel_technician_role".
	Anchor string `json:"anchor"`
	// AttackIfAtWar attack_if_at_war boolean.
	// AttackIfAtWar 处于战争状态时是否攻击布尔值.
	AttackIfAtWar bool `json:"attack_if_at_war"`
	// AttackIfOtherSecurityStatusDropping attack_if_other_security_status_dropping boolean.
	// AttackIfOtherSecurityStatusDropping 对方安全等级下降时是否攻击布尔值.
	AttackIfOtherSecurityStatusDropping bool `json:"attack_if_other_security_status_dropping"`
	// AttackSecurityStatusThreshold Starbase (POS) will attack if target's security standing is lower than this value.
	// AttackSecurityStatusThreshold 如果目标的安全声望低于此值，母星基地（POS）将发动攻击.
	AttackSecurityStatusThreshold float64 `json:"attack_security_status_threshold"`
	// AttackStandingThreshold Starbase (POS) will attack if target's standing is lower than this value.
	// AttackStandingThreshold 如果目标的声望低于此值，母星基地（POS）将发动攻击.
	AttackStandingThreshold float64 `json:"attack_standing_threshold"`
	// FuelBayTake Who can take fuel blocks out of the starbase (POS)'s fuel bay.
	// FuelBayTake 谁可以从母星（POS）的燃料舱中取走燃料块.
	// Enum values: "alliance_member", "config_starbase_equipment_role", "corporation_member", "starbase_fuel_technician_role".
	FuelBayTake string `json:"fuel_bay_take"`
	// FuelBayView Who can view the starbase (POS)'s fule bay. Characters either need to have required role or belong to the starbase (POS) owner's corporation or alliance, as described by the enum, all other access settings follows the same scheme.
	// FuelBayView 谁可以查看母星（POS）的燃料舱。角色需要拥有所需角色，或属于母星（POS）所有者的军团或联盟，如枚举所述，所有其他访问设置遵循同一规则.
	// Enum values: "alliance_member", "config_starbase_equipment_role", "corporation_member", "starbase_fuel_technician_role".
	FuelBayView string `json:"fuel_bay_view"`
	// Fuels Fuel blocks and other things that will be consumed when operating a starbase (POS)
	// Fuels 运营母星基地（POS）时将消耗的燃料块及其他物资.
	Fuels []StarbaseFuel `json:"fuels"`
	// Offline Who can offline starbase (POS) and its structures.
	// Offline 谁可以下线母星（POS）及其建筑.
	// Enum values: "alliance_member", "config_starbase_equipment_role", "corporation_member", "starbase_fuel_technician_role".
	Offline string `json:"offline"`
	// Online Who can online starbase (POS) and its structures.
	// Online 谁可以上线母星（POS）及其建筑.
	// Enum values: "alliance_member", "config_starbase_equipment_role", "corporation_member", "starbase_fuel_technician_role".
	Online string `json:"online"`
	// Unanchor Who can unanchor starbase (POS) and its structures.
	// Unanchor 谁可以解除锚定母星（POS）及其建筑.
	// Enum values: "alliance_member", "config_starbase_equipment_role", "corporation_member", "starbase_fuel_technician_role".
	Unanchor string `json:"unanchor"`
	// UseAllianceStandings True if the starbase (POS) is using alliance standings, otherwise using corporation's.
	// UseAllianceStandings 如果母星基地（POS）使用联盟声望则为 true，否则使用军团声望.
	UseAllianceStandings bool `json:"use_alliance_standings"`
}

// CorporationStructure 200 ok object.
// CorporationStructure 200 ok 对象.
type CorporationStructure struct {
	// CorporationId ID of the corporation that owns the structure.
	// CorporationId 拥有此建筑的军团 ID.
	CorporationId int32 `json:"corporation_id"`
	// FuelExpires Date on which the structure will run out of fuel.
	// FuelExpires 建筑燃料耗尽的日期.
	FuelExpires time.Time `json:"fuel_expires"`
	// Name The structure name.
	// Name 建筑名称.
	Name string `json:"name"`
	// NextReinforceApply The date and time when the structure's newly requested reinforcement times (e.g. next_reinforce_hour and next_reinforce_day) will take effect.
	// NextReinforceApply 建筑新申请的可增强时间（如 next_reinforce_hour 和 next_reinforce_day）生效的日期和时间.
	NextReinforceApply time.Time `json:"next_reinforce_apply"`
	// NextReinforceHour The requested change to reinforce_hour that will take effect at the time shown by next_reinforce_apply.
	// NextReinforceHour 对 reinforce_hour 的请求更改，将于 next_reinforce_apply 所示时间生效.
	NextReinforceHour int32 `json:"next_reinforce_hour"`
	// ProfileId The id of the ACL profile for this citadel.
	// ProfileId 此堡垒的 ACL 配置文件 ID.
	ProfileId int32 `json:"profile_id"`
	// ReinforceHour The hour of day that determines the four hour window when the structure will randomly exit its reinforcement periods and become vulnerable to attack against its armor and/or hull. The structure will become vulnerable at a random time that is +/- 2 hours centered on the value of this property.
	// ReinforceHour 决定建筑随机退出增强时段、进入装甲和/或结构可被攻击状态的四小时窗口的当日时刻。建筑将以该属性值为中心、±2 小时内的随机时刻进入可攻击状态.
	ReinforceHour int32 `json:"reinforce_hour"`
	// Services Contains a list of service upgrades, and their state.
	// Services 包含服务升级列表及其状态.
	Services []StructureService `json:"services"`
	// State state string.
	// State 状态 string.
	// Enum values: "anchor_vulnerable", "anchoring", "armor_reinforce", "armor_vulnerable", "deploy_vulnerable", "fitting_invulnerable", "hull_reinforce", "hull_vulnerable", "online_deprecated", "onlining_vulnerable", "shield_vulnerable", "unanchored", "unknown".
	State string `json:"state"`
	// StateTimerEnd Date at which the structure will move to it's next state.
	// StateTimerEnd 建筑将进入下一状态的日期.
	StateTimerEnd time.Time `json:"state_timer_end"`
	// StateTimerStart Date at which the structure entered it's current state.
	// StateTimerStart 建筑进入当前状态的日期.
	StateTimerStart time.Time `json:"state_timer_start"`
	// StructureId The Item ID of the structure.
	// StructureId 建筑的物品 ID.
	StructureId int64 `json:"structure_id"`
	// SystemId The solar system the structure is in.
	// SystemId 建筑所在的星系.
	SystemId int32 `json:"system_id"`
	// TypeId The type id of the structure.
	// TypeId 建筑的 type id.
	TypeId int32 `json:"type_id"`
	// UnanchorsAt Date at which the structure will unanchor.
	// UnanchorsAt 建筑解除锚定的日期.
	UnanchorsAt time.Time `json:"unanchors_at"`
}

// StructureService service object.
// StructureService 服务 object.
type StructureService struct {
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// State state string.
	// State 状态 string.
	// Enum values: "online", "offline", "cleanup".
	State string `json:"state"`
}

// CorporationTitle 200 ok object.
// CorporationTitle 200 ok 对象.
type CorporationTitle struct {
	// GrantableRoles grantable_roles array.
	// GrantableRoles 可授予角色数组.
	GrantableRoles []string `json:"grantable_roles"`
	// GrantableRolesAtBase grantable_roles_at_base array.
	// GrantableRolesAtBase 基地可授予角色数组.
	GrantableRolesAtBase []string `json:"grantable_roles_at_base"`
	// GrantableRolesAtHq grantable_roles_at_hq array.
	// GrantableRolesAtHq 总部可授予角色数组.
	GrantableRolesAtHq []string `json:"grantable_roles_at_hq"`
	// GrantableRolesAtOther grantable_roles_at_other array.
	// GrantableRolesAtOther 其他地点可授予角色数组.
	GrantableRolesAtOther []string `json:"grantable_roles_at_other"`
	// Name name string.
	// Name 名称 string.
	Name string `json:"name"`
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
	// TitleId title_id integer.
	// TitleId title_id 整数.
	TitleId int32 `json:"title_id"`
}
