package models

import (
	"net/url"
)

// GetCharactersCharacterIdFittings 200 ok object.
// GetCharactersCharacterIdFittings 200 ok 对象.
type GetCharactersCharacterIdFittings struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// FittingId fitting_id integer.
	// FittingId 装配 ID 整数.
	FittingId int32 `json:"fitting_id"`
	// Items items array.
	// Items items 数组.
	Items []GetCharactersCharacterIdFittingsItem `json:"items"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
}

// GetCharactersCharacterIdFittingsItem item object.
// GetCharactersCharacterIdFittingsItem 物品对象.
type GetCharactersCharacterIdFittingsItem struct {
	// Flag flag string.
	// Flag 位置标记字符串.
	// Enum values: "Cargo", "DroneBay", "FighterBay", "HiSlot0", "HiSlot1", "HiSlot2", "HiSlot3", "HiSlot4", "HiSlot5", "HiSlot6", "HiSlot7", "Invalid", "LoSlot0", "LoSlot1", "LoSlot2", "LoSlot3", "LoSlot4", "LoSlot5", "LoSlot6", "LoSlot7", "MedSlot0", "MedSlot1", "MedSlot2", "MedSlot3", "MedSlot4", "MedSlot5", "MedSlot6", "MedSlot7", "RigSlot0", "RigSlot1", "RigSlot2", "ServiceSlot0", "ServiceSlot1", "ServiceSlot2", "ServiceSlot3", "ServiceSlot4", "ServiceSlot5", "ServiceSlot6", "ServiceSlot7", "SubSystemSlot0", "SubSystemSlot1", "SubSystemSlot2", "SubSystemSlot3".
	Flag string `json:"flag"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// PostCharactersCharacterIdFittingsCreated 201 created object.
// PostCharactersCharacterIdFittingsCreated 201 created 对象.
type PostCharactersCharacterIdFittingsCreated struct {
	// FittingId fitting_id integer.
	// FittingId 装配 ID 整数.
	FittingId int32 `json:"fitting_id"`
}

// PostCharactersCharacterIdFittingsFitting fitting object.
// PostCharactersCharacterIdFittingsFitting 装配对象.
type PostCharactersCharacterIdFittingsFitting struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Items items array.
	// Items items 数组.
	Items []PostCharactersCharacterIdFittingsItem `json:"items"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// ShipTypeId ship_type_id integer.
	// ShipTypeId 舰船类型ID integer.
	ShipTypeId int32 `json:"ship_type_id"`
}

// PostCharactersCharacterIdFittingsItem item object.
// PostCharactersCharacterIdFittingsItem 物品对象.
type PostCharactersCharacterIdFittingsItem struct {
	// Flag Fitting location for the item. Entries placed in 'Invalid' will be discarded. If this leaves the fitting with nothing, it will cause an error.
	// Flag 物品在装配中的位置。放入 'Invalid' 的条目将被丢弃。如果这导致装配为空，则会引发错误。
	// Enum values: "Cargo", "DroneBay", "FighterBay", "HiSlot0", "HiSlot1", "HiSlot2", "HiSlot3", "HiSlot4", "HiSlot5", "HiSlot6", "HiSlot7", "Invalid", "LoSlot0", "LoSlot1", "LoSlot2", "LoSlot3", "LoSlot4", "LoSlot5", "LoSlot6", "LoSlot7", "MedSlot0", "MedSlot1", "MedSlot2", "MedSlot3", "MedSlot4", "MedSlot5", "MedSlot6", "MedSlot7", "RigSlot0", "RigSlot1", "RigSlot2", "ServiceSlot0", "ServiceSlot1", "ServiceSlot2", "ServiceSlot3", "ServiceSlot4", "ServiceSlot5", "ServiceSlot6", "ServiceSlot7", "SubSystemSlot0", "SubSystemSlot1", "SubSystemSlot2", "SubSystemSlot3".
	Flag string `json:"flag"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// DeleteCharactersCharacterIdFittingsFittingIdParams holds the optional query and header parameters of the request.
// DeleteCharactersCharacterIdFittingsFittingIdParams 保存请求的可选查询与头部参数。
type DeleteCharactersCharacterIdFittingsFittingIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *DeleteCharactersCharacterIdFittingsFittingIdParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdFittingsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdFittingsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdFittingsParams struct {
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

func (p *GetCharactersCharacterIdFittingsParams) Values() (url.Values, map[string]string) {
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

// PostCharactersCharacterIdFittingsParams holds the optional query and header parameters of the request.
// PostCharactersCharacterIdFittingsParams 保存请求的可选查询与头部参数。
type PostCharactersCharacterIdFittingsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *PostCharactersCharacterIdFittingsParams) Values() (url.Values, map[string]string) {
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
