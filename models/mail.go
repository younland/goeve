package models

import (
	"time"
)

// MailHeader 200 ok object.
// MailHeader 200 ok 对象.
type MailHeader struct {
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
	Recipients []MailRecipient `json:"recipients"`
	// Subject Mail subject.
	// Subject 邮件主题.
	Subject string `json:"subject"`
	// Timestamp When the mail was sent.
	// Timestamp 邮件发送时间.
	Timestamp time.Time `json:"timestamp"`
}

// MailLabel label object.
// MailLabel label 对象.
type MailLabel struct {
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

// MailLabels 200 ok object.
// MailLabels 200 ok 对象.
type MailLabels struct {
	// Labels labels array.
	// Labels labels 数组.
	Labels []MailLabel `json:"labels"`
	// TotalUnreadCount total_unread_count integer.
	// TotalUnreadCount 未读总数整数.
	TotalUnreadCount int32 `json:"total_unread_count"`
}

// MailingList 200 ok object.
// MailingList 200 ok 对象.
type MailingList struct {
	// MailingListId Mailing list ID.
	// MailingListId 邮件列表 ID.
	MailingListId int32 `json:"mailing_list_id"`
	// Name name string.
	// Name 名称字符串.
	Name string `json:"name"`
}

// Mail 200 ok object.
// Mail 200 ok 对象.
type Mail struct {
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
	Recipients []MailRecipient `json:"recipients"`
	// Subject Mail subject.
	// Subject 邮件主题.
	Subject string `json:"subject"`
	// Timestamp When the mail was sent.
	// Timestamp 邮件发送时间.
	Timestamp time.Time `json:"timestamp"`
}

// MailRecipient recipient object.
// MailRecipient 收件人 object.
type MailRecipient struct {
	// RecipientId recipient_id integer.
	// RecipientId 收件人ID integer.
	RecipientId int32 `json:"recipient_id"`
	// RecipientType recipient_type string.
	// RecipientType 收件人类型 string.
	// Enum values: "alliance", "character", "corporation", "mailing_list".
	RecipientType string `json:"recipient_type"`
}

// MailLabelRequest label object.
// MailLabelRequest label 对象.
type MailLabelRequest struct {
	// Color Hexadecimal string representing label color, in RGB format.
	// Color 表示位标颜色的十六进制字符串，RGB 格式.
	// Enum values: "#0000fe", "#006634", "#0099ff", "#00ff33", "#01ffff", "#349800", "#660066", "#666666", "#999999", "#99ffff", "#9a0000", "#ccff9a", "#e6e6e6", "#fe0000", "#ff6600", "#ffff01", "#ffffcd", "#ffffff".
	Color string `json:"color"`
	// Name name string.
	// Name 名称字符串.
	Name string `json:"name"`
}

// MailRequest mail object.
// MailRequest mail 对象.
type MailRequest struct {
	// ApprovedCost approved_cost integer.
	// ApprovedCost 批准成本整数.
	ApprovedCost int64 `json:"approved_cost"`
	// Body body string.
	// Body 正文字符串.
	Body string `json:"body"`
	// Recipients recipients array.
	// Recipients 收件人列表 array.
	Recipients []MailRecipient `json:"recipients"`
	// Subject subject string.
	// Subject 主题 string.
	Subject string `json:"subject"`
}

// MailMetadata contents object.
// MailMetadata 内容对象.
type MailMetadata struct {
	// Labels Labels to assign to the mail. Pre-existing labels are unassigned.
	// Labels 要分配给邮件的标签。原先已存在的标签会被解除分配.
	Labels []int32 `json:"labels"`
	// Read Whether the mail is flagged as read.
	// Read 邮件是否被标记为已读.
	Read bool `json:"read"`
}
