package models

// LoyaltyPoints 200 ok object.
// LoyaltyPoints 200 ok 对象.
type LoyaltyPoints struct {
	// CorporationId corporation_id integer.
	// CorporationId 军团 ID 整数.
	CorporationId int32 `json:"corporation_id"`
	// LoyaltyPoints loyalty_points integer.
	// LoyaltyPoints loyalty_points 整数（忠诚点）
	LoyaltyPoints int32 `json:"loyalty_points"`
}

// LoyaltyStoreOffer 200 ok object.
// LoyaltyStoreOffer 200 ok 对象.
type LoyaltyStoreOffer struct {
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
	RequiredItems []LoyaltyStoreOfferRequiredItem `json:"required_items"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// LoyaltyStoreOfferRequiredItem required_item object.
// LoyaltyStoreOfferRequiredItem 必需物品 object.
type LoyaltyStoreOfferRequiredItem struct {
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}
