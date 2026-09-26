package models

import (
	"net/url"
	"strconv"
)

// GetAlliancesAllianceIdContacts 200 ok object.
// GetAlliancesAllianceIdContacts 200 ok 对象.
type GetAlliancesAllianceIdContacts struct {
	// ContactId contact_id integer.
	// ContactId 联系人 ID 整数.
	ContactId int32 `json:"contact_id"`
	// ContactType contact_type string.
	// ContactType 联系人类型字符串.
	// Enum values: "character", "corporation", "alliance", "faction".
	ContactType string `json:"contact_type"`
	// LabelIds label_ids array.
	// LabelIds label_ids 数组.
	LabelIds []int64 `json:"label_ids"`
	// Standing Standing of the contact.
	// Standing 联系人的声望.
	Standing float64 `json:"standing"`
}

// GetAlliancesAllianceIdContactsLabels 200 ok object.
// GetAlliancesAllianceIdContactsLabels 200 ok 对象.
type GetAlliancesAllianceIdContactsLabels struct {
	// LabelId label_id integer.
	// LabelId label_id 整数.
	LabelId int64 `json:"label_id"`
	// LabelName label_name string.
	// LabelName label_name 字符串.
	LabelName string `json:"label_name"`
}

// GetCharactersCharacterIdContacts 200 ok object.
// GetCharactersCharacterIdContacts 200 ok 对象.
type GetCharactersCharacterIdContacts struct {
	// ContactId contact_id integer.
	// ContactId 联系人 ID 整数.
	ContactId int32 `json:"contact_id"`
	// ContactType contact_type string.
	// ContactType 联系人类型字符串.
	// Enum values: "character", "corporation", "alliance", "faction".
	ContactType string `json:"contact_type"`
	// IsBlocked Whether this contact is in the blocked list. Note a missing value denotes unknown, not true or false.
	// IsBlocked 此联系人是否在屏蔽列表中。注意缺失值表示未知，而非真或假.
	IsBlocked bool `json:"is_blocked"`
	// IsWatched Whether this contact is being watched.
	// IsWatched 是否正在关注此联系人.
	IsWatched bool `json:"is_watched"`
	// LabelIds label_ids array.
	// LabelIds label_ids 数组.
	LabelIds []int64 `json:"label_ids"`
	// Standing Standing of the contact.
	// Standing 联系人的声望.
	Standing float64 `json:"standing"`
}

// GetCharactersCharacterIdContactsLabels 200 ok object.
// GetCharactersCharacterIdContactsLabels 200 ok 对象.
type GetCharactersCharacterIdContactsLabels struct {
	// LabelId label_id integer.
	// LabelId label_id 整数.
	LabelId int64 `json:"label_id"`
	// LabelName label_name string.
	// LabelName label_name 字符串.
	LabelName string `json:"label_name"`
}

// GetCorporationsCorporationIdContacts 200 ok object.
// GetCorporationsCorporationIdContacts 200 ok 对象.
type GetCorporationsCorporationIdContacts struct {
	// ContactId contact_id integer.
	// ContactId 联系人 ID 整数.
	ContactId int32 `json:"contact_id"`
	// ContactType contact_type string.
	// ContactType 联系人类型字符串.
	// Enum values: "character", "corporation", "alliance", "faction".
	ContactType string `json:"contact_type"`
	// IsWatched Whether this contact is being watched.
	// IsWatched 是否正在关注此联系人.
	IsWatched bool `json:"is_watched"`
	// LabelIds label_ids array.
	// LabelIds label_ids 数组.
	LabelIds []int64 `json:"label_ids"`
	// Standing Standing of the contact.
	// Standing 联系人的声望.
	Standing float64 `json:"standing"`
}

// GetCorporationsCorporationIdContactsLabels 200 ok object.
// GetCorporationsCorporationIdContactsLabels 200 ok 对象.
type GetCorporationsCorporationIdContactsLabels struct {
	// LabelId label_id integer.
	// LabelId label_id 整数.
	LabelId int64 `json:"label_id"`
	// LabelName label_name string.
	// LabelName label_name 字符串.
	LabelName string `json:"label_name"`
}

// DeleteCharactersCharacterIdContactsParams holds the optional query and header parameters of the request.
// DeleteCharactersCharacterIdContactsParams 保存请求的可选查询与头部参数。
type DeleteCharactersCharacterIdContactsParams struct {
	// ContactIds A list of contacts to delete.
	// ContactIds 要删除的联系人列表.
	ContactIds []int32
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *DeleteCharactersCharacterIdContactsParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	for _, v := range p.ContactIds {
		if query == nil {
			query = url.Values{}
		}
		query.Add("contact_ids", strconv.Itoa(int(v)))
	}
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

// GetAlliancesAllianceIdContactsLabelsParams holds the optional query and header parameters of the request.
// GetAlliancesAllianceIdContactsLabelsParams 保存请求的可选查询与头部参数。
type GetAlliancesAllianceIdContactsLabelsParams struct {
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

func (p *GetAlliancesAllianceIdContactsLabelsParams) Values() (url.Values, map[string]string) {
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

// GetAlliancesAllianceIdContactsParams holds the optional query and header parameters of the request.
// GetAlliancesAllianceIdContactsParams 保存请求的可选查询与头部参数。
type GetAlliancesAllianceIdContactsParams struct {
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

func (p *GetAlliancesAllianceIdContactsParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdContactsLabelsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdContactsLabelsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdContactsLabelsParams struct {
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

func (p *GetCharactersCharacterIdContactsLabelsParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdContactsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdContactsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdContactsParams struct {
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

func (p *GetCharactersCharacterIdContactsParams) Values() (url.Values, map[string]string) {
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

// GetCorporationsCorporationIdContactsLabelsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdContactsLabelsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdContactsLabelsParams struct {
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

func (p *GetCorporationsCorporationIdContactsLabelsParams) Values() (url.Values, map[string]string) {
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

// GetCorporationsCorporationIdContactsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdContactsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdContactsParams struct {
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

func (p *GetCorporationsCorporationIdContactsParams) Values() (url.Values, map[string]string) {
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

// PostCharactersCharacterIdContactsParams holds the optional query and header parameters of the request.
// PostCharactersCharacterIdContactsParams 保存请求的可选查询与头部参数。
type PostCharactersCharacterIdContactsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// LabelIds Add custom labels to the new contact.
	// LabelIds 为新联系人添加自定义标签.
	LabelIds []int64
	// Standing Standing for the contact.
	// Standing 该联系人的声望.
	Standing *float64
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
	// Watched Whether the contact should be watched, note this is only effective on characters.
	// Watched 是否关注该联系人，注意此设置仅对角色有效.
	Watched *bool
}

func (p *PostCharactersCharacterIdContactsParams) Values() (url.Values, map[string]string) {
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
	for _, v := range p.LabelIds {
		if query == nil {
			query = url.Values{}
		}
		query.Add("label_ids", strconv.Itoa(int(v)))
	}
	if p.Standing != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("standing", strconv.FormatFloat(*p.Standing, 'f', -1, 64))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	if p.Watched != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("watched", strconv.FormatBool(*p.Watched))
	}
	return query, headers
}

// PutCharactersCharacterIdContactsParams holds the optional query and header parameters of the request.
// PutCharactersCharacterIdContactsParams 保存请求的可选查询与头部参数。
type PutCharactersCharacterIdContactsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// LabelIds Add custom labels to the contact.
	// LabelIds 为该联系人添加自定义标签.
	LabelIds []int64
	// Standing Standing for the contact.
	// Standing 该联系人的声望.
	Standing *float64
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
	// Watched Whether the contact should be watched, note this is only effective on characters.
	// Watched 是否关注该联系人，注意此设置仅对角色有效.
	Watched *bool
}

func (p *PutCharactersCharacterIdContactsParams) Values() (url.Values, map[string]string) {
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
	for _, v := range p.LabelIds {
		if query == nil {
			query = url.Values{}
		}
		query.Add("label_ids", strconv.Itoa(int(v)))
	}
	if p.Standing != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("standing", strconv.FormatFloat(*p.Standing, 'f', -1, 64))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	if p.Watched != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("watched", strconv.FormatBool(*p.Watched))
	}
	return query, headers
}
