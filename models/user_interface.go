package models

// NewMailRequest new_mail object.
// NewMailRequest new_mail 对象.
type NewMailRequest struct {
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
