package models

import (
	"net/url"
	"time"
)

// GetCharactersCharacterIdAttributesOk 200 ok object.
// GetCharactersCharacterIdAttributesOk 200 ok 对象.
type GetCharactersCharacterIdAttributesOk struct {
	// AccruedRemapCooldownDate Neural remapping cooldown after a character uses remap accrued over time.
	// AccruedRemapCooldownDate 角色使用随时间累积的重映射后的神经重映射冷却时间.
	AccruedRemapCooldownDate time.Time `json:"accrued_remap_cooldown_date"`
	// BonusRemaps Number of available bonus character neural remaps.
	// BonusRemaps 角色可用的奖励神经重映射次数.
	BonusRemaps int32 `json:"bonus_remaps"`
	// Charisma charisma integer.
	// Charisma 魅力整数.
	Charisma int32 `json:"charisma"`
	// Intelligence intelligence integer.
	// Intelligence 智力整数.
	Intelligence int32 `json:"intelligence"`
	// LastRemapDate Datetime of last neural remap, including usage of bonus remaps.
	// LastRemapDate 上次神经重映射的时间，包括使用奖励重映射.
	LastRemapDate time.Time `json:"last_remap_date"`
	// Memory memory integer.
	// Memory memory 整数.
	Memory int32 `json:"memory"`
	// Perception perception integer.
	// Perception perception 整数.
	Perception int32 `json:"perception"`
	// Willpower willpower integer.
	// Willpower 毅力（willpower）整数.
	Willpower int32 `json:"willpower"`
}

// GetCharactersCharacterIdAttributesParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdAttributesParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdAttributesParams struct {
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

func (p *GetCharactersCharacterIdAttributesParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdSkillqueue 200 ok object.
// GetCharactersCharacterIdSkillqueue 200 ok 对象.
type GetCharactersCharacterIdSkillqueue struct {
	// FinishDate Date on which training of the skill will complete. Omitted if the skill queue is paused.
	// FinishDate 该技能训练完成的日期。如果技能队列暂停，则省略此项。
	FinishDate time.Time `json:"finish_date"`
	// FinishedLevel finished_level integer.
	// FinishedLevel 完成等级整数.
	FinishedLevel int32 `json:"finished_level"`
	// LevelEndSp level_end_sp integer.
	// LevelEndSp level_end_sp 整数.
	LevelEndSp int32 `json:"level_end_sp"`
	// LevelStartSp Amount of SP that was in the skill when it started training it's current level. Used to calculate % of current level complete.
	// LevelStartSp 该技能开始训练当前等级时已有的技能点数量。用于计算当前等级完成的百分比。
	LevelStartSp int32 `json:"level_start_sp"`
	// QueuePosition queue_position integer.
	// QueuePosition 队列位置 integer.
	QueuePosition int32 `json:"queue_position"`
	// SkillId skill_id integer.
	// SkillId 技能ID integer.
	SkillId int32 `json:"skill_id"`
	// StartDate start_date string.
	// StartDate 开始日期 string.
	StartDate time.Time `json:"start_date"`
	// TrainingStartSp training_start_sp integer.
	// TrainingStartSp 训练起始技能点（SP）整数.
	TrainingStartSp int32 `json:"training_start_sp"`
}

// GetCharactersCharacterIdSkillqueueParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdSkillqueueParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdSkillqueueParams struct {
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

func (p *GetCharactersCharacterIdSkillqueueParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdSkillsOk 200 ok object.
// GetCharactersCharacterIdSkillsOk 200 ok 对象.
type GetCharactersCharacterIdSkillsOk struct {
	// Skills skills array.
	// Skills 技能列表 array.
	Skills []GetCharactersCharacterIdSkillsSkill `json:"skills"`
	// TotalSp total_sp integer.
	// TotalSp 总技能点数（SP）整数.
	TotalSp int64 `json:"total_sp"`
	// UnallocatedSp Skill points available to be assigned.
	// UnallocatedSp 可分配的技能点.
	UnallocatedSp int32 `json:"unallocated_sp"`
}

// GetCharactersCharacterIdSkillsSkill skill object.
// GetCharactersCharacterIdSkillsSkill 技能 object.
type GetCharactersCharacterIdSkillsSkill struct {
	// ActiveSkillLevel active_skill_level integer.
	// ActiveSkillLevel 当前技能等级整数.
	ActiveSkillLevel int32 `json:"active_skill_level"`
	// SkillId skill_id integer.
	// SkillId 技能ID integer.
	SkillId int32 `json:"skill_id"`
	// SkillpointsInSkill skillpoints_in_skill integer.
	// SkillpointsInSkill 该技能中的技能点 integer.
	SkillpointsInSkill int64 `json:"skillpoints_in_skill"`
	// TrainedSkillLevel trained_skill_level integer.
	// TrainedSkillLevel 已训练技能等级整数.
	TrainedSkillLevel int32 `json:"trained_skill_level"`
}

// GetCharactersCharacterIdSkillsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdSkillsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdSkillsParams struct {
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

func (p *GetCharactersCharacterIdSkillsParams) Values() (url.Values, map[string]string) {
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
