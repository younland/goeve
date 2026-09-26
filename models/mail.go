package models

import (
	"net/url"
	"strconv"
	"time"
)

// DeleteCharactersCharacterIdMailLabelsLabelIdParams holds the optional query and header parameters of the request.
// DeleteCharactersCharacterIdMailLabelsLabelIdParams 保存请求的可选查询与头部参数。
type DeleteCharactersCharacterIdMailLabelsLabelIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *DeleteCharactersCharacterIdMailLabelsLabelIdParams) Values() (url.Values, map[string]string) {
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

// DeleteCharactersCharacterIdMailMailIdParams holds the optional query and header parameters of the request.
// DeleteCharactersCharacterIdMailMailIdParams 保存请求的可选查询与头部参数。
type DeleteCharactersCharacterIdMailMailIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *DeleteCharactersCharacterIdMailMailIdParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdMail 200 ok object.
// GetCharactersCharacterIdMail 200 ok 对象.
type GetCharactersCharacterIdMail struct {
	// From From whom the mail was sent.
	// From 邮件的发送者.
	From int32 `json:"from"`
	// IsRead is_read boolean.
	// IsRead 是否已读布尔值.
	IsRead bool `json:"is_read"`
	// Labels labels array.
	// Labels labels 数组.
	Labels []int32 `json:"labels"`
	// MailId mail_id integer.
	// MailId mail_id 整数.
	MailId int32 `json:"mail_id"`
	// Recipients Recipients of the mail.
	// Recipients 邮件的收件人.
	Recipients []GetCharactersCharacterIdMailRecipient `json:"recipients"`
	// Subject Mail subject.
	// Subject 邮件主题.
	Subject string `json:"subject"`
	// Timestamp When the mail was sent.
	// Timestamp 邮件发送时间.
	Timestamp time.Time `json:"timestamp"`
}

// GetCharactersCharacterIdMailLabelsLabel label object.
// GetCharactersCharacterIdMailLabelsLabel label 对象.
type GetCharactersCharacterIdMailLabelsLabel struct {
	// Color color string.
	// Color 颜色字符串.
	// Enum values: "#0000fe", "#006634", "#0099ff", "#00ff33", "#01ffff", "#349800", "#660066", "#666666", "#999999", "#99ffff", "#9a0000", "#ccff9a", "#e6e6e6", "#fe0000", "#ff6600", "#ffff01", "#ffffcd", "#ffffff".
	Color string `json:"color"`
	// LabelId label_id integer.
	// LabelId label_id 整数.
	LabelId int32 `json:"label_id"`
	// Name name string.
	// Name 名称字符串.
	Name string `json:"name"`
	// UnreadCount unread_count integer.
	// UnreadCount 未读数整数.
	UnreadCount int32 `json:"unread_count"`
}

// GetCharactersCharacterIdMailLabels 200 ok object.
// GetCharactersCharacterIdMailLabels 200 ok 对象.
type GetCharactersCharacterIdMailLabels struct {
	// Labels labels array.
	// Labels labels 数组.
	Labels []GetCharactersCharacterIdMailLabelsLabel `json:"labels"`
	// TotalUnreadCount total_unread_count integer.
	// TotalUnreadCount 未读总数整数.
	TotalUnreadCount int32 `json:"total_unread_count"`
}

// GetCharactersCharacterIdMailLists 200 ok object.
// GetCharactersCharacterIdMailLists 200 ok 对象.
type GetCharactersCharacterIdMailLists struct {
	// MailingListId Mailing list ID.
	// MailingListId 邮件列表 ID.
	MailingListId int32 `json:"mailing_list_id"`
	// Name name string.
	// Name 名称字符串.
	Name string `json:"name"`
}

// GetCharactersCharacterIdMailMailId 200 ok object.
// GetCharactersCharacterIdMailMailId 200 ok 对象.
type GetCharactersCharacterIdMailMailId struct {
	// Body Mail's body.
	// Body 邮件正文.
	Body string `json:"body"`
	// From From whom the mail was sent.
	// From 邮件的发送者.
	From int32 `json:"from"`
	// Labels Labels attached to the mail.
	// Labels 附加到邮件的位标标签.
	Labels []int32 `json:"labels"`
	// Read Whether the mail is flagged as read.
	// Read 邮件是否被标记为已读.
	Read bool `json:"read"`
	// Recipients Recipients of the mail.
	// Recipients 邮件的收件人.
	Recipients []GetCharactersCharacterIdMailMailIdRecipient `json:"recipients"`
	// Subject Mail subject.
	// Subject 邮件主题.
	Subject string `json:"subject"`
	// Timestamp When the mail was sent.
	// Timestamp 邮件发送时间.
	Timestamp time.Time `json:"timestamp"`
}

// GetCharactersCharacterIdMailMailIdRecipient recipient object.
// GetCharactersCharacterIdMailMailIdRecipient 收件人 object.
type GetCharactersCharacterIdMailMailIdRecipient struct {
	// RecipientId recipient_id integer.
	// RecipientId 收件人ID integer.
	RecipientId int32 `json:"recipient_id"`
	// RecipientType recipient_type string.
	// RecipientType 收件人类型 string.
	// Enum values: "alliance", "character", "corporation", "mailing_list".
	RecipientType string `json:"recipient_type"`
}

// GetCharactersCharacterIdMailParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdMailParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdMailParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Labels Fetch only mails that match one or more of the given labels.
	// Labels 仅获取匹配一个或多个指定标签的邮件.
	Labels []int32
	// LastMailId List only mail with an ID lower than the given ID, if present.
	// LastMailId 仅列出 ID 小于给定 ID 的邮件（如果提供）
	LastMailId *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdMailParams) Values() (url.Values, map[string]string) {
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
	for _, v := range p.Labels {
		if query == nil {
			query = url.Values{}
		}
		query.Add("labels", strconv.Itoa(int(v)))
	}
	if p.LastMailId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("last_mail_id", strconv.FormatInt(int64(*p.LastMailId), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharactersCharacterIdMailLabelsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdMailLabelsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdMailLabelsParams struct {
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

func (p *GetCharactersCharacterIdMailLabelsParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdMailListsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdMailListsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdMailListsParams struct {
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

func (p *GetCharactersCharacterIdMailListsParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdMailMailIdParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdMailMailIdParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdMailMailIdParams struct {
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

func (p *GetCharactersCharacterIdMailMailIdParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdMailRecipient recipient object.
// GetCharactersCharacterIdMailRecipient 收件人 object.
type GetCharactersCharacterIdMailRecipient struct {
	// RecipientId recipient_id integer.
	// RecipientId 收件人ID integer.
	RecipientId int32 `json:"recipient_id"`
	// RecipientType recipient_type string.
	// RecipientType 收件人类型 string.
	// Enum values: "alliance", "character", "corporation", "mailing_list".
	RecipientType string `json:"recipient_type"`
}

// PostCharactersCharacterIdMailLabelsLabel label object.
// PostCharactersCharacterIdMailLabelsLabel label 对象.
type PostCharactersCharacterIdMailLabelsLabel struct {
	// Color Hexadecimal string representing label color, in RGB format.
	// Color 表示位标颜色的十六进制字符串，RGB 格式.
	// Enum values: "#0000fe", "#006634", "#0099ff", "#00ff33", "#01ffff", "#349800", "#660066", "#666666", "#999999", "#99ffff", "#9a0000", "#ccff9a", "#e6e6e6", "#fe0000", "#ff6600", "#ffff01", "#ffffcd", "#ffffff".
	Color string `json:"color"`
	// Name name string.
	// Name 名称字符串.
	Name string `json:"name"`
}

// PostCharactersCharacterIdMailLabelsParams holds the optional query and header parameters of the request.
// PostCharactersCharacterIdMailLabelsParams 保存请求的可选查询与头部参数。
type PostCharactersCharacterIdMailLabelsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCharactersCharacterIdMailLabelsParams) Values() (url.Values, map[string]string) {
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

// PostCharactersCharacterIdMailMail mail object.
// PostCharactersCharacterIdMailMail mail 对象.
type PostCharactersCharacterIdMailMail struct {
	// ApprovedCost approved_cost integer.
	// ApprovedCost 批准成本整数.
	ApprovedCost int64 `json:"approved_cost"`
	// Body body string.
	// Body 正文字符串.
	Body string `json:"body"`
	// Recipients recipients array.
	// Recipients 收件人列表 array.
	Recipients []PostCharactersCharacterIdMailRecipient `json:"recipients"`
	// Subject subject string.
	// Subject 主题 string.
	Subject string `json:"subject"`
}

// PostCharactersCharacterIdMailParams holds the optional query and header parameters of the request.
// PostCharactersCharacterIdMailParams 保存请求的可选查询与头部参数。
type PostCharactersCharacterIdMailParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCharactersCharacterIdMailParams) Values() (url.Values, map[string]string) {
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

// PostCharactersCharacterIdMailRecipient recipient object.
// PostCharactersCharacterIdMailRecipient 收件人 object.
type PostCharactersCharacterIdMailRecipient struct {
	// RecipientId recipient_id integer.
	// RecipientId 收件人ID integer.
	RecipientId int32 `json:"recipient_id"`
	// RecipientType recipient_type string.
	// RecipientType 收件人类型 string.
	// Enum values: "alliance", "character", "corporation", "mailing_list".
	RecipientType string `json:"recipient_type"`
}

// PutCharactersCharacterIdMailMailIdContents contents object.
// PutCharactersCharacterIdMailMailIdContents 内容对象.
type PutCharactersCharacterIdMailMailIdContents struct {
	// Labels Labels to assign to the mail. Pre-existing labels are unassigned.
	// Labels 要分配给邮件的标签。原先已存在的标签会被解除分配。
	Labels []int32 `json:"labels"`
	// Read Whether the mail is flagged as read.
	// Read 邮件是否被标记为已读.
	Read bool `json:"read"`
}

// PutCharactersCharacterIdMailMailIdParams holds the optional query and header parameters of the request.
// PutCharactersCharacterIdMailMailIdParams 保存请求的可选查询与头部参数。
type PutCharactersCharacterIdMailMailIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PutCharactersCharacterIdMailMailIdParams) Values() (url.Values, map[string]string) {
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
