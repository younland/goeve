package models

import (
	"net/url"
	"strconv"
)

// PostUiOpenwindowNewmailNewMail new_mail object.
// PostUiOpenwindowNewmailNewMail new_mail 对象.
type PostUiOpenwindowNewmailNewMail struct {
	// Body body string.
	// Body 正文字符串.
	Body string `json:"body"`
	// Recipients recipients array.
	// Recipients 收件人列表 array.
	Recipients []int32 `json:"recipients"`
	// Subject subject string.
	// Subject 主题 string.
	Subject string `json:"subject"`
	// ToCorpOrAllianceId to_corp_or_alliance_id integer.
	// ToCorpOrAllianceId to_corp_or_alliance_id 整数.
	ToCorpOrAllianceId int32 `json:"to_corp_or_alliance_id"`
	// ToMailingListId Corporations, alliances and mailing lists are all types of mailing groups. You may only send to one mailing group, at a time, so you may fill out either this field or the to_corp_or_alliance_ids field.
	// ToMailingListId 军团、联盟和邮件列表都属于邮件组。一次只能向一个邮件组发送邮件，因此你可以填写此字段或 to_corp_or_alliance_ids 字段.
	ToMailingListId int32 `json:"to_mailing_list_id"`
}

// PostUiAutopilotWaypointParams holds the optional query and header parameters of the request.
// PostUiAutopilotWaypointParams 保存请求的可选查询与头部参数。
type PostUiAutopilotWaypointParams struct {
	// AddToBeginning Whether this solar system should be added to the beginning of all waypoints.
	// AddToBeginning 是否将该星系添加到所有路径点的最前面.
	AddToBeginning *bool
	// ClearOtherWaypoints Whether clean other waypoints beforing adding this one.
	// ClearOtherWaypoints 是否在添加此路径点前清除其他路径点.
	ClearOtherWaypoints *bool
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// DestinationId The destination to travel to, can be solar system, station or structure's id.
	// DestinationId 要前往的目的地，可为星系、空间站或建筑的 ID.
	DestinationId *int64
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostUiAutopilotWaypointParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AddToBeginning != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("add_to_beginning", strconv.FormatBool(*p.AddToBeginning))
	}
	if p.ClearOtherWaypoints != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("clear_other_waypoints", strconv.FormatBool(*p.ClearOtherWaypoints))
	}
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.DestinationId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("destination_id", strconv.FormatInt(int64(*p.DestinationId), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostUiOpenwindowContractParams holds the optional query and header parameters of the request.
// PostUiOpenwindowContractParams 保存请求的可选查询与头部参数。
type PostUiOpenwindowContractParams struct {
	// ContractId The contract to open.
	// ContractId 要打开的合同.
	ContractId *int32
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostUiOpenwindowContractParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.ContractId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("contract_id", strconv.FormatInt(int64(*p.ContractId), 10))
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

// PostUiOpenwindowInformationParams holds the optional query and header parameters of the request.
// PostUiOpenwindowInformationParams 保存请求的可选查询与头部参数。
type PostUiOpenwindowInformationParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// TargetId The target to open.
	// TargetId 要打开的目标.
	TargetId *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostUiOpenwindowInformationParams) Values() (url.Values, map[string]string) {
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
	if p.TargetId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("target_id", strconv.FormatInt(int64(*p.TargetId), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// PostUiOpenwindowMarketdetailsParams holds the optional query and header parameters of the request.
// PostUiOpenwindowMarketdetailsParams 保存请求的可选查询与头部参数。
type PostUiOpenwindowMarketdetailsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
	// TypeId The item type to open in market window.
	// TypeId 在市场窗口中打开的物品类型.
	TypeId *int32
}

func (p *PostUiOpenwindowMarketdetailsParams) Values() (url.Values, map[string]string) {
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
	if p.TypeId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("type_id", strconv.FormatInt(int64(*p.TypeId), 10))
	}
	return query, headers
}

// PostUiOpenwindowNewmailParams holds the optional query and header parameters of the request.
// PostUiOpenwindowNewmailParams 保存请求的可选查询与头部参数。
type PostUiOpenwindowNewmailParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostUiOpenwindowNewmailParams) Values() (url.Values, map[string]string) {
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
