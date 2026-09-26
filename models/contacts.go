package models

// AllianceContact 200 ok object.
// AllianceContact 200 ok 对象.
type AllianceContact struct {
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

// ContactLabel 200 ok object.
// ContactLabel 200 ok 对象.
type ContactLabel struct {
	// LabelId label_id integer.
	// LabelId label_id 整数.
	LabelId int64 `json:"label_id"`
	// LabelName label_name string.
	// LabelName label_name 字符串.
	LabelName string `json:"label_name"`
}

// CharacterContact 200 ok object.
// CharacterContact 200 ok 对象.
type CharacterContact struct {
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

// CorporationContact 200 ok object.
// CorporationContact 200 ok 对象.
type CorporationContact struct {
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
