package models

import (
	"net/url"
)

// GetCharactersCharacterIdLoyaltyPoints 200 ok object.
// GetCharactersCharacterIdLoyaltyPoints 200 ok 对象.
type GetCharactersCharacterIdLoyaltyPoints struct {
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// LoyaltyPoints loyalty_points integer.
	// LoyaltyPoints loyalty_points 整数（忠诚点）
	LoyaltyPoints int32 `json:"loyalty_points"`
}

// GetCharactersCharacterIdLoyaltyPointsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdLoyaltyPointsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdLoyaltyPointsParams struct {
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

func (p *GetCharactersCharacterIdLoyaltyPointsParams) Values() (url.Values, map[string]string) {
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

// GetLoyaltyStoresCorporationIdOffers 200 ok object.
// GetLoyaltyStoresCorporationIdOffers 200 ok 对象.
type GetLoyaltyStoresCorporationIdOffers struct {
	// AkCost Analysis kredit cost.
	// AkCost 分析成本（ISK）
	AkCost int32 `json:"ak_cost"`
	// IskCost isk_cost integer.
	// IskCost ISK 费用整数.
	IskCost int64 `json:"isk_cost"`
	// LpCost lp_cost integer.
	// LpCost lp_cost 整数（LP 成本）
	LpCost int32 `json:"lp_cost"`
	// OfferId offer_id integer.
	// OfferId offer_id 整数.
	OfferId int32 `json:"offer_id"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// RequiredItems required_items array.
	// RequiredItems 必需物品列表 array.
	RequiredItems []GetLoyaltyStoresCorporationIdOffersRequiredItem `json:"required_items"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetLoyaltyStoresCorporationIdOffersRequiredItem required_item object.
// GetLoyaltyStoresCorporationIdOffersRequiredItem 必需物品 object.
type GetLoyaltyStoresCorporationIdOffersRequiredItem struct {
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetStoresCorporationIdOffersParams holds the optional query and header parameters of the request.
// GetStoresCorporationIdOffersParams 保存请求的可选查询与头部参数。
type GetStoresCorporationIdOffersParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetStoresCorporationIdOffersParams) Values() (url.Values, map[string]string) {
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
