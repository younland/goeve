package models

import (
	"net/url"
	"time"
)

// GetCharactersCharacterIdOpportunities 200 ok object.
// GetCharactersCharacterIdOpportunities 200 ok 对象.
type GetCharactersCharacterIdOpportunities struct {
	// CompletedAt completed_at string.
	// CompletedAt 完成时间字符串.
	CompletedAt time.Time `json:"completed_at"`
	// TaskId task_id integer.
	// TaskId task_id 整数.
	TaskId int32 `json:"task_id"`
}

// GetOpportunitiesGroupsGroupId 200 ok object.
// GetOpportunitiesGroupsGroupId 200 ok 对象.
type GetOpportunitiesGroupsGroupId struct {
	// ConnectedGroups The groups that are connected to this group on the opportunities map.
	// ConnectedGroups 机遇地图上与此组相连的其他组.
	ConnectedGroups []int32 `json:"connected_groups"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// GroupId group_id integer.
	// GroupId 分组 ID 整数.
	GroupId int32 `json:"group_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Notification notification string.
	// Notification notification 字符串.
	Notification string `json:"notification"`
	// RequiredTasks Tasks need to complete for this group.
	// RequiredTasks 完成该组所需的任务.
	RequiredTasks []int32 `json:"required_tasks"`
}

// GetOpportunitiesTasksTaskId 200 ok object.
// GetOpportunitiesTasksTaskId 200 ok 对象.
type GetOpportunitiesTasksTaskId struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Notification notification string.
	// Notification notification 字符串.
	Notification string `json:"notification"`
	// TaskId task_id integer.
	// TaskId task_id 整数.
	TaskId int32 `json:"task_id"`
}

// GetCharactersCharacterIdOpportunitiesParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdOpportunitiesParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdOpportunitiesParams struct {
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

func (p *GetCharactersCharacterIdOpportunitiesParams) Values() (url.Values, map[string]string) {
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

// GetGroupsGroupIdParams holds the optional query and header parameters of the request.
// GetGroupsGroupIdParams 保存请求的可选查询与头部参数。
type GetGroupsGroupIdParams struct {
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

func (p *GetGroupsGroupIdParams) Values() (url.Values, map[string]string) {
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

// GetGroupsParams holds the optional query and header parameters of the request.
// GetGroupsParams 保存请求的可选查询与头部参数。
type GetGroupsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetGroupsParams) Values() (url.Values, map[string]string) {
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

// GetTasksParams holds the optional query and header parameters of the request.
// GetTasksParams 保存请求的可选查询与头部参数。
type GetTasksParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetTasksParams) Values() (url.Values, map[string]string) {
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

// GetTasksTaskIdParams holds the optional query and header parameters of the request.
// GetTasksTaskIdParams 保存请求的可选查询与头部参数。
type GetTasksTaskIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetTasksTaskIdParams) Values() (url.Values, map[string]string) {
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
