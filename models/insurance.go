package models

// InsurancePrice 200 ok object.
// InsurancePrice 200 ok 对象.
type InsurancePrice struct {
	// Levels A list of a available insurance levels for this ship type.
	// Levels 该舰船类型可用的保险等级列表.
	Levels []InsurancePriceLevel `json:"levels"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// InsurancePriceLevel level object.
// InsurancePriceLevel level 对象.
type InsurancePriceLevel struct {
	// Cost cost number.
	// Cost 成本数值.
	Cost float64 `json:"cost"`
	// Name Localized insurance level.
	// Name 本地化的保险等级.
	Name string `json:"name"`
	// Payout payout number.
	// Payout payout 数值.
	Payout float64 `json:"payout"`
}
