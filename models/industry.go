package models

import (
	"time"
)

// CharacterIndustryJob 200 ok object.
// CharacterIndustryJob 200 ok 对象.
type CharacterIndustryJob struct {
	// ActivityId Job activity ID.
	// ActivityId 工作活动 ID.
	ActivityId int32 `json:"activity_id"`
	// BlueprintId blueprint_id integer.
	// BlueprintId 蓝图 ID 整数.
	BlueprintId int64 `json:"blueprint_id"`
	// BlueprintLocationId Location ID of the location from which the blueprint was installed. Normally a station ID, but can also be an asset (e.g. container) or corporation facility.
	// BlueprintLocationId 蓝图安装地点的 ID。通常是空间站 ID，但也可以是资产（如集装箱）或军团设施.
	BlueprintLocationId int64 `json:"blueprint_location_id"`
	// BlueprintTypeId blueprint_type_id integer.
	// BlueprintTypeId 蓝图类型 ID 整数.
	BlueprintTypeId int32 `json:"blueprint_type_id"`
	// CompletedCharacterId ID of the character which completed this job.
	// CompletedCharacterId 完成此工作的角色 ID.
	CompletedCharacterId int32 `json:"completed_character_id"`
	// CompletedDate Date and time when this job was completed.
	// CompletedDate 该作业完成的日期和时间.
	CompletedDate time.Time `json:"completed_date"`
	// Cost The sume of job installation fee and industry facility tax.
	// Cost 作业安装费与工业设施税之和.
	Cost float64 `json:"cost"`
	// Duration Job duration in seconds.
	// Duration 工作持续时间（秒）
	Duration int32 `json:"duration"`
	// EndDate Date and time when this job finished.
	// EndDate 该作业完成的日期和时间.
	EndDate time.Time `json:"end_date"`
	// FacilityId ID of the facility where this job is running.
	// FacilityId 此工作正在运行的设施 ID.
	FacilityId int64 `json:"facility_id"`
	// InstallerId ID of the character which installed this job.
	// InstallerId 安装此工作的角色 ID.
	InstallerId int32 `json:"installer_id"`
	// JobId Unique job ID.
	// JobId 作业的唯一 ID.
	JobId int32 `json:"job_id"`
	// LicensedRuns Number of runs blueprint is licensed for.
	// LicensedRuns 蓝图许可的运转次数.
	LicensedRuns int32 `json:"licensed_runs"`
	// OutputLocationId Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility.
	// OutputLocationId 作业产出交付地点的 ID。通常是空间站 ID，但也可以是军团设施.
	OutputLocationId int64 `json:"output_location_id"`
	// PauseDate Date and time when this job was paused (i.e. time when the facility where this job was installed went offline)
	// PauseDate 该作业暂停的日期和时间（即安装该作业的设施离线的时刻）
	PauseDate time.Time `json:"pause_date"`
	// Probability Chance of success for invention.
	// Probability 发明成功率.
	Probability float64 `json:"probability"`
	// ProductTypeId Type ID of product (manufactured, copied or invented)
	// ProductTypeId 产物（制造、拷贝或发明）的 type ID.
	ProductTypeId int32 `json:"product_type_id"`
	// Runs Number of runs for a manufacturing job, or number of copies to make for a blueprint copy.
	// Runs 制造作业的运转次数，或制造一份蓝图拷贝的拷贝数.
	Runs int32 `json:"runs"`
	// StartDate Date and time when this job started.
	// StartDate 该作业开始的日期和时间.
	StartDate time.Time `json:"start_date"`
	// StationId ID of the station where industry facility is located.
	// StationId 工业设施所在空间站的 ID.
	StationId int64 `json:"station_id"`
	// Status status string.
	// Status 状态 string.
	// Enum values: "active", "cancelled", "delivered", "paused", "ready", "reverted".
	Status string `json:"status"`
	// SuccessfulRuns Number of successful runs for this job. Equal to runs unless this is an invention job.
	// SuccessfulRuns 该作业成功运转的次数。除非这是发明作业，否则等于运转次数.
	SuccessfulRuns int32 `json:"successful_runs"`
}

// MiningLedgerEntry 200 ok object.
// MiningLedgerEntry 200 ok 对象.
type MiningLedgerEntry struct {
	// Date date string.
	// Date 日期字符串.
	Date string `json:"date"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int64 `json:"quantity"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// MoonExtraction 200 ok object.
// MoonExtraction 200 ok 对象.
type MoonExtraction struct {
	// ChunkArrivalTime The time at which the chunk being extracted will arrive and can be fractured by the moon mining drill.
	// ChunkArrivalTime 正在开采的矿层抵达并可用月球钻探机碎裂的时间。
	ChunkArrivalTime time.Time `json:"chunk_arrival_time"`
	// ExtractionStartTime The time at which the current extraction was initiated.
	// ExtractionStartTime 当前开采发起的时间。
	ExtractionStartTime time.Time `json:"extraction_start_time"`
	// MoonId moon_id integer.
	// MoonId moon_id 整数.
	MoonId int32 `json:"moon_id"`
	// NaturalDecayTime The time at which the chunk being extracted will naturally fracture if it is not first fractured by the moon mining drill.
	// NaturalDecayTime 如果未被月球钻探机提前碎裂，正在开采的矿层自然碎裂的时间。
	NaturalDecayTime time.Time `json:"natural_decay_time"`
	// StructureId structure_id integer.
	// StructureId 建筑ID integer.
	StructureId int64 `json:"structure_id"`
}

// MiningObserver 200 ok object.
// MiningObserver 200 ok 对象.
type MiningObserver struct {
	// LastUpdated last_updated string.
	// LastUpdated last_updated 字符串.
	LastUpdated string `json:"last_updated"`
	// ObserverId The entity that was observing the asteroid field when it was mined.
	// ObserverId 采掘时观察该小行星带的实体。
	ObserverId int64 `json:"observer_id"`
	// ObserverType The category of the observing entity.
	// ObserverType 观察实体的类别.
	// Enum values: "structure".
	ObserverType string `json:"observer_type"`
}

// MiningObserverEntry 200 ok object.
// MiningObserverEntry 200 ok 对象.
type MiningObserverEntry struct {
	// CharacterId The character that did the mining.
	// CharacterId 进行采矿的角色.
	CharacterId int32 `json:"character_id"`
	// LastUpdated last_updated string.
	// LastUpdated last_updated 字符串.
	LastUpdated string `json:"last_updated"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int64 `json:"quantity"`
	// RecordedCorporationId The corporation id of the character at the time data was recorded.
	// RecordedCorporationId 数据记录时该角色的军团 ID。
	RecordedCorporationId int32 `json:"recorded_corporation_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// CorporationIndustryJob 200 ok object.
// CorporationIndustryJob 200 ok 对象.
type CorporationIndustryJob struct {
	// ActivityId Job activity ID.
	// ActivityId 工作活动 ID.
	ActivityId int32 `json:"activity_id"`
	// BlueprintId blueprint_id integer.
	// BlueprintId 蓝图 ID 整数.
	BlueprintId int64 `json:"blueprint_id"`
	// BlueprintLocationId Location ID of the location from which the blueprint was installed. Normally a station ID, but can also be an asset (e.g. container) or corporation facility.
	// BlueprintLocationId 蓝图安装地点的 ID。通常是空间站 ID，但也可以是资产（如集装箱）或军团设施.
	BlueprintLocationId int64 `json:"blueprint_location_id"`
	// BlueprintTypeId blueprint_type_id integer.
	// BlueprintTypeId 蓝图类型 ID 整数.
	BlueprintTypeId int32 `json:"blueprint_type_id"`
	// CompletedCharacterId ID of the character which completed this job.
	// CompletedCharacterId 完成此工作的角色 ID.
	CompletedCharacterId int32 `json:"completed_character_id"`
	// CompletedDate Date and time when this job was completed.
	// CompletedDate 该作业完成的日期和时间.
	CompletedDate time.Time `json:"completed_date"`
	// Cost The sume of job installation fee and industry facility tax.
	// Cost 作业安装费与工业设施税之和.
	Cost float64 `json:"cost"`
	// Duration Job duration in seconds.
	// Duration 工作持续时间（秒）
	Duration int32 `json:"duration"`
	// EndDate Date and time when this job finished.
	// EndDate 该作业完成的日期和时间.
	EndDate time.Time `json:"end_date"`
	// FacilityId ID of the facility where this job is running.
	// FacilityId 此工作正在运行的设施 ID.
	FacilityId int64 `json:"facility_id"`
	// InstallerId ID of the character which installed this job.
	// InstallerId 安装此工作的角色 ID.
	InstallerId int32 `json:"installer_id"`
	// JobId Unique job ID.
	// JobId 作业的唯一 ID.
	JobId int32 `json:"job_id"`
	// LicensedRuns Number of runs blueprint is licensed for.
	// LicensedRuns 蓝图许可的运转次数.
	LicensedRuns int32 `json:"licensed_runs"`
	// LocationId ID of the location for the industry facility.
	// LocationId 工业设施所在地点的 ID.
	LocationId int64 `json:"location_id"`
	// OutputLocationId Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility.
	// OutputLocationId 作业产出交付地点的 ID。通常是空间站 ID，但也可以是军团设施.
	OutputLocationId int64 `json:"output_location_id"`
	// PauseDate Date and time when this job was paused (i.e. time when the facility where this job was installed went offline)
	// PauseDate 该作业暂停的日期和时间（即安装该作业的设施离线的时刻）
	PauseDate time.Time `json:"pause_date"`
	// Probability Chance of success for invention.
	// Probability 发明成功率.
	Probability float64 `json:"probability"`
	// ProductTypeId Type ID of product (manufactured, copied or invented)
	// ProductTypeId 产物（制造、拷贝或发明）的 type ID.
	ProductTypeId int32 `json:"product_type_id"`
	// Runs Number of runs for a manufacturing job, or number of copies to make for a blueprint copy.
	// Runs 制造作业的运转次数，或制造一份蓝图拷贝的拷贝数.
	Runs int32 `json:"runs"`
	// StartDate Date and time when this job started.
	// StartDate 该作业开始的日期和时间.
	StartDate time.Time `json:"start_date"`
	// Status status string.
	// Status 状态 string.
	// Enum values: "active", "cancelled", "delivered", "paused", "ready", "reverted".
	Status string `json:"status"`
	// SuccessfulRuns Number of successful runs for this job. Equal to runs unless this is an invention job.
	// SuccessfulRuns 该作业成功运转的次数。除非这是发明作业，否则等于运转次数.
	SuccessfulRuns int32 `json:"successful_runs"`
}

// IndustryFacility 200 ok object.
// IndustryFacility 200 ok 对象.
type IndustryFacility struct {
	// FacilityId ID of the facility.
	// FacilityId 设施 ID.
	FacilityId int64 `json:"facility_id"`
	// OwnerId Owner of the facility.
	// OwnerId 设施的拥有者.
	OwnerId int32 `json:"owner_id"`
	// RegionId Region ID where the facility is.
	// RegionId 设施所在星域的 ID.
	RegionId int32 `json:"region_id"`
	// SolarSystemId Solar system ID where the facility is.
	// SolarSystemId 设施所在星系的 ID.
	SolarSystemId int32 `json:"solar_system_id"`
	// Tax Tax imposed by the facility.
	// Tax 设施征收的税.
	Tax float64 `json:"tax"`
	// TypeId Type ID of the facility.
	// TypeId 设施的 type ID.
	TypeId int32 `json:"type_id"`
}

// IndustrySystemCostIndices 200 ok object.
// IndustrySystemCostIndices 200 ok 对象.
type IndustrySystemCostIndices struct {
	// CostIndices cost_indices array.
	// CostIndices cost_indices 数组.
	CostIndices []CostIndex `json:"cost_indices"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
}

// CostIndex cost_indice object.
// CostIndex cost_indice 对象.
type CostIndex struct {
	// Activity activity string.
	// Activity 活动字符串.
	// Enum values: "copying", "duplicating", "invention", "manufacturing", "none", "reaction", "researching_material_efficiency", "researching_technology", "researching_time_efficiency", "reverse_engineering".
	Activity string `json:"activity"`
	// CostIndex cost_index number.
	// CostIndex 成本指数数值.
	CostIndex float64 `json:"cost_index"`
}
