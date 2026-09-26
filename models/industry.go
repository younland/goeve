package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetCharactersCharacterIdIndustryJobs 200 ok object.
// GetCharactersCharacterIdIndustryJobs 200 ok 对象.
type GetCharactersCharacterIdIndustryJobs struct {
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

// GetCharactersCharacterIdIndustryJobsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdIndustryJobsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdIndustryJobsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// IncludeCompleted Whether to retrieve completed character industry jobs. Only includes jobs from the past 90 days.
	// IncludeCompleted 是否获取已完成的角色工业作业。仅包含过去 90 天内的作业.
	IncludeCompleted *bool
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdIndustryJobsParams) Values() (url.Values, map[string]string) {
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
	if p.IncludeCompleted != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("include_completed", strconv.FormatBool(*p.IncludeCompleted))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharactersCharacterIdMining 200 ok object.
// GetCharactersCharacterIdMining 200 ok 对象.
type GetCharactersCharacterIdMining struct {
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

// GetCharactersCharacterIdMiningParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdMiningParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdMiningParams struct {
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

func (p *GetCharactersCharacterIdMiningParams) Values() (url.Values, map[string]string) {
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

// GetCorporationCorporationIdMiningExtractions 200 ok object.
// GetCorporationCorporationIdMiningExtractions 200 ok 对象.
type GetCorporationCorporationIdMiningExtractions struct {
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

// GetCorporationCorporationIdMiningExtractionsParams holds the optional query and header parameters of the request.
// GetCorporationCorporationIdMiningExtractionsParams 保存请求的可选查询与头部参数。
type GetCorporationCorporationIdMiningExtractionsParams struct {
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

func (p *GetCorporationCorporationIdMiningExtractionsParams) Values() (url.Values, map[string]string) {
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

// GetCorporationCorporationIdMiningObservers 200 ok object.
// GetCorporationCorporationIdMiningObservers 200 ok 对象.
type GetCorporationCorporationIdMiningObservers struct {
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

// GetCorporationCorporationIdMiningObserversObserverId 200 ok object.
// GetCorporationCorporationIdMiningObserversObserverId 200 ok 对象.
type GetCorporationCorporationIdMiningObserversObserverId struct {
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

// GetCorporationCorporationIdMiningObserversObserverIdParams holds the optional query and header parameters of the request.
// GetCorporationCorporationIdMiningObserversObserverIdParams 保存请求的可选查询与头部参数。
type GetCorporationCorporationIdMiningObserversObserverIdParams struct {
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

func (p *GetCorporationCorporationIdMiningObserversObserverIdParams) Values() (url.Values, map[string]string) {
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

// GetCorporationCorporationIdMiningObserversParams holds the optional query and header parameters of the request.
// GetCorporationCorporationIdMiningObserversParams 保存请求的可选查询与头部参数。
type GetCorporationCorporationIdMiningObserversParams struct {
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

func (p *GetCorporationCorporationIdMiningObserversParams) Values() (url.Values, map[string]string) {
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

// GetCorporationsCorporationIdIndustryJobs 200 ok object.
// GetCorporationsCorporationIdIndustryJobs 200 ok 对象.
type GetCorporationsCorporationIdIndustryJobs struct {
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

// GetCorporationsCorporationIdIndustryJobsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdIndustryJobsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdIndustryJobsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// IncludeCompleted Whether to retrieve completed corporation industry jobs. Only includes jobs from the past 90 days.
	// IncludeCompleted 是否获取已完成的军团工业作业。仅包含过去 90 天内的作业.
	IncludeCompleted *bool
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCorporationsCorporationIdIndustryJobsParams) Values() (url.Values, map[string]string) {
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
	if p.IncludeCompleted != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("include_completed", strconv.FormatBool(*p.IncludeCompleted))
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

// GetFacilitiesParams holds the optional query and header parameters of the request.
// GetFacilitiesParams 保存请求的可选查询与头部参数。
type GetFacilitiesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetFacilitiesParams) Values() (url.Values, map[string]string) {
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

// GetIndustryFacilities 200 ok object.
// GetIndustryFacilities 200 ok 对象.
type GetIndustryFacilities struct {
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

// GetIndustrySystems 200 ok object.
// GetIndustrySystems 200 ok 对象.
type GetIndustrySystems struct {
	// CostIndices cost_indices array.
	// CostIndices cost_indices 数组.
	CostIndices []GetIndustrySystemsCostIndice `json:"cost_indices"`
	// SolarSystemId solar_system_id integer.
	// SolarSystemId 星系ID integer.
	SolarSystemId int32 `json:"solar_system_id"`
}

// GetIndustrySystemsCostIndice cost_indice object.
// GetIndustrySystemsCostIndice cost_indice 对象.
type GetIndustrySystemsCostIndice struct {
	// Activity activity string.
	// Activity 活动字符串.
	// Enum values: "copying", "duplicating", "invention", "manufacturing", "none", "reaction", "researching_material_efficiency", "researching_technology", "researching_time_efficiency", "reverse_engineering".
	Activity string `json:"activity"`
	// CostIndex cost_index number.
	// CostIndex 成本指数数值.
	CostIndex float64 `json:"cost_index"`
}

// GetSystemsParams holds the optional query and header parameters of the request.
// GetSystemsParams 保存请求的可选查询与头部参数。
type GetSystemsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetSystemsParams) Values() (url.Values, map[string]string) {
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
