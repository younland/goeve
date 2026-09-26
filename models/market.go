package models

import "time"

// CharacterMarketOrder 200 ok object.
// CharacterMarketOrder 200 ok 对象.
type CharacterMarketOrder struct {
	// Duration Number of days for which order is valid (starting from the issued date). An order expires at time issued + duration.
	// Duration 订单有效的天数（自发布日期起）。订单在发布时间 + 持续时长时过期.
	Duration int32 `json:"duration"`
	// Escrow For buy orders, the amount of ISK in escrow.
	// Escrow 对于买单，托管中的 ISK 金额.
	Escrow float64 `json:"escrow"`
	// IsBuyOrder True if the order is a bid (buy) order.
	// IsBuyOrder 如果该订单是买单则为 true.
	IsBuyOrder bool `json:"is_buy_order"`
	// IsCorporation Signifies whether the buy/sell order was placed on behalf of a corporation.
	// IsCorporation 表示该买单/卖单是否以军团名义下单。
	IsCorporation bool `json:"is_corporation"`
	// Issued Date and time when this order was issued.
	// Issued 该订单发布的日期和时间.
	Issued time.Time `json:"issued"`
	// LocationId ID of the location where order was placed.
	// LocationId 订单所在地点的 ID.
	LocationId int64 `json:"location_id"`
	// MinVolume For buy orders, the minimum quantity that will be accepted in a matching sell order.
	// MinVolume 对于买单，匹配卖单时可接受的最小数量.
	MinVolume int32 `json:"min_volume"`
	// OrderId Unique order ID.
	// OrderId 订单的唯一 ID.
	OrderId int64 `json:"order_id"`
	// Price Cost per unit for this order.
	// Price 该订单的每单位成本.
	Price float64 `json:"price"`
	// RangeValue Valid order range, numbers are ranges in jumps.
	// RangeValue 有效订单范围，数字为以跳数计的范围.
	// Enum values: "1", "10", "2", "20", "3", "30", "4", "40", "5", "region", "solarsystem", "station".
	RangeValue string `json:"range"`
	// RegionId ID of the region where order was placed.
	// RegionId 订单所在星域的 ID.
	RegionId int32 `json:"region_id"`
	// TypeId The type ID of the item transacted in this order.
	// TypeId 此订单中交易物品的 type ID.
	TypeId int32 `json:"type_id"`
	// VolumeRemain Quantity of items still required or offered.
	// VolumeRemain 仍然需要或提供的物品数量.
	VolumeRemain int32 `json:"volume_remain"`
	// VolumeTotal Quantity of items required or offered at time order was placed.
	// VolumeTotal 下单时所需或提供的物品数量.
	VolumeTotal int32 `json:"volume_total"`
}

// CharacterMarketOrderHistory 200 ok object.
// CharacterMarketOrderHistory 200 ok 对象.
type CharacterMarketOrderHistory struct {
	// Duration Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration.
	// Duration 订单的有效天数（自发布日期起）。订单在发布时间 + 持续时长时过期.
	Duration int32 `json:"duration"`
	// Escrow For buy orders, the amount of ISK in escrow.
	// Escrow 对于买单，托管中的 ISK 金额.
	Escrow float64 `json:"escrow"`
	// IsBuyOrder True if the order is a bid (buy) order.
	// IsBuyOrder 如果该订单是买单则为 true.
	IsBuyOrder bool `json:"is_buy_order"`
	// IsCorporation Signifies whether the buy/sell order was placed on behalf of a corporation.
	// IsCorporation 表示该买单/卖单是否以军团名义下单。
	IsCorporation bool `json:"is_corporation"`
	// Issued Date and time when this order was issued.
	// Issued 该订单发布的日期和时间.
	Issued time.Time `json:"issued"`
	// LocationId ID of the location where order was placed.
	// LocationId 订单所在地点的 ID.
	LocationId int64 `json:"location_id"`
	// MinVolume For buy orders, the minimum quantity that will be accepted in a matching sell order.
	// MinVolume 对于买单，匹配卖单时可接受的最小数量.
	MinVolume int32 `json:"min_volume"`
	// OrderId Unique order ID.
	// OrderId 订单的唯一 ID.
	OrderId int64 `json:"order_id"`
	// Price Cost per unit for this order.
	// Price 该订单的每单位成本.
	Price float64 `json:"price"`
	// RangeValue Valid order range, numbers are ranges in jumps.
	// RangeValue 有效订单范围，数字为以跳数计的范围.
	// Enum values: "1", "10", "2", "20", "3", "30", "4", "40", "5", "region", "solarsystem", "station".
	RangeValue string `json:"range"`
	// RegionId ID of the region where order was placed.
	// RegionId 订单所在星域的 ID.
	RegionId int32 `json:"region_id"`
	// State Current order state.
	// State 当前订单状态.
	// Enum values: "cancelled", "expired".
	State string `json:"state"`
	// TypeId The type ID of the item transacted in this order.
	// TypeId 此订单中交易物品的 type ID.
	TypeId int32 `json:"type_id"`
	// VolumeRemain Quantity of items still required or offered.
	// VolumeRemain 仍然需要或提供的物品数量.
	VolumeRemain int32 `json:"volume_remain"`
	// VolumeTotal Quantity of items required or offered at time order was placed.
	// VolumeTotal 下单时所需或提供的物品数量.
	VolumeTotal int32 `json:"volume_total"`
}

// CorporationMarketOrder 200 ok object.
// CorporationMarketOrder 200 ok 对象.
type CorporationMarketOrder struct {
	// Duration Number of days for which order is valid (starting from the issued date). An order expires at time issued + duration.
	// Duration 订单有效的天数（自发布日期起）。订单在发布时间 + 持续时长时过期.
	Duration int32 `json:"duration"`
	// Escrow For buy orders, the amount of ISK in escrow.
	// Escrow 对于买单，托管中的 ISK 金额.
	Escrow float64 `json:"escrow"`
	// IsBuyOrder True if the order is a bid (buy) order.
	// IsBuyOrder 如果该订单是买单则为 true.
	IsBuyOrder bool `json:"is_buy_order"`
	// Issued Date and time when this order was issued.
	// Issued 该订单发布的日期和时间.
	Issued time.Time `json:"issued"`
	// IssuedBy The character who issued this order.
	// IssuedBy 下达此订单的角色.
	IssuedBy int32 `json:"issued_by"`
	// LocationId ID of the location where order was placed.
	// LocationId 订单所在地点的 ID.
	LocationId int64 `json:"location_id"`
	// MinVolume For buy orders, the minimum quantity that will be accepted in a matching sell order.
	// MinVolume 对于买单，匹配卖单时可接受的最小数量.
	MinVolume int32 `json:"min_volume"`
	// OrderId Unique order ID.
	// OrderId 订单的唯一 ID.
	OrderId int64 `json:"order_id"`
	// Price Cost per unit for this order.
	// Price 该订单的每单位成本.
	Price float64 `json:"price"`
	// RangeValue Valid order range, numbers are ranges in jumps.
	// RangeValue 有效订单范围，数字为以跳数计的范围.
	// Enum values: "1", "10", "2", "20", "3", "30", "4", "40", "5", "region", "solarsystem", "station".
	RangeValue string `json:"range"`
	// RegionId ID of the region where order was placed.
	// RegionId 订单所在星域的 ID.
	RegionId int32 `json:"region_id"`
	// TypeId The type ID of the item transacted in this order.
	// TypeId 此订单中交易物品的 type ID.
	TypeId int32 `json:"type_id"`
	// VolumeRemain Quantity of items still required or offered.
	// VolumeRemain 仍然需要或提供的物品数量.
	VolumeRemain int32 `json:"volume_remain"`
	// VolumeTotal Quantity of items required or offered at time order was placed.
	// VolumeTotal 下单时所需或提供的物品数量.
	VolumeTotal int32 `json:"volume_total"`
	// WalletDivision The corporation wallet division used for this order.
	// WalletDivision 此订单使用的军团钱包分部。
	WalletDivision int32 `json:"wallet_division"`
}

// CorporationMarketOrderHistory 200 ok object.
// CorporationMarketOrderHistory 200 ok 对象.
type CorporationMarketOrderHistory struct {
	// Duration Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration.
	// Duration 订单的有效天数（自发布日期起）。订单在发布时间 + 持续时长时过期.
	Duration int32 `json:"duration"`
	// Escrow For buy orders, the amount of ISK in escrow.
	// Escrow 对于买单，托管中的 ISK 金额.
	Escrow float64 `json:"escrow"`
	// IsBuyOrder True if the order is a bid (buy) order.
	// IsBuyOrder 如果该订单是买单则为 true.
	IsBuyOrder bool `json:"is_buy_order"`
	// Issued Date and time when this order was issued.
	// Issued 该订单发布的日期和时间.
	Issued time.Time `json:"issued"`
	// IssuedBy The character who issued this order.
	// IssuedBy 下达此订单的角色.
	IssuedBy int32 `json:"issued_by"`
	// LocationId ID of the location where order was placed.
	// LocationId 订单所在地点的 ID.
	LocationId int64 `json:"location_id"`
	// MinVolume For buy orders, the minimum quantity that will be accepted in a matching sell order.
	// MinVolume 对于买单，匹配卖单时可接受的最小数量.
	MinVolume int32 `json:"min_volume"`
	// OrderId Unique order ID.
	// OrderId 订单的唯一 ID.
	OrderId int64 `json:"order_id"`
	// Price Cost per unit for this order.
	// Price 该订单的每单位成本.
	Price float64 `json:"price"`
	// RangeValue Valid order range, numbers are ranges in jumps.
	// RangeValue 有效订单范围，数字为以跳数计的范围.
	// Enum values: "1", "10", "2", "20", "3", "30", "4", "40", "5", "region", "solarsystem", "station".
	RangeValue string `json:"range"`
	// RegionId ID of the region where order was placed.
	// RegionId 订单所在星域的 ID.
	RegionId int32 `json:"region_id"`
	// State Current order state.
	// State 当前订单状态.
	// Enum values: "cancelled", "expired".
	State string `json:"state"`
	// TypeId The type ID of the item transacted in this order.
	// TypeId 此订单中交易物品的 type ID.
	TypeId int32 `json:"type_id"`
	// VolumeRemain Quantity of items still required or offered.
	// VolumeRemain 仍然需要或提供的物品数量.
	VolumeRemain int32 `json:"volume_remain"`
	// VolumeTotal Quantity of items required or offered at time order was placed.
	// VolumeTotal 下单时所需或提供的物品数量.
	VolumeTotal int32 `json:"volume_total"`
	// WalletDivision The corporation wallet division used for this order.
	// WalletDivision 此订单使用的军团钱包分部.
	WalletDivision int32 `json:"wallet_division"`
}

// MarketGroup 200 ok object.
// MarketGroup 200 ok 对象.
type MarketGroup struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// MarketGroupId market_group_id integer.
	// MarketGroupId market_group_id 整数.
	MarketGroupId int32 `json:"market_group_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// ParentGroupId parent_group_id integer.
	// ParentGroupId parent_group_id 整数.
	ParentGroupId int32 `json:"parent_group_id"`
	// Types types array.
	// Types types 数组.
	Types []int32 `json:"types"`
}

// MarketPrice 200 ok object.
// MarketPrice 200 ok 对象.
type MarketPrice struct {
	// AdjustedPrice adjusted_price number.
	// AdjustedPrice 调整价格数值.
	AdjustedPrice float64 `json:"adjusted_price"`
	// AveragePrice average_price number.
	// AveragePrice 平均价格数值.
	AveragePrice float64 `json:"average_price"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// MarketHistoryEntry 200 ok object.
// MarketHistoryEntry 200 ok 对象.
type MarketHistoryEntry struct {
	// Average average number.
	// Average 平均值数值.
	Average float64 `json:"average"`
	// Date The date of this historical statistic entry.
	// Date 此历史统计条目的日期.
	Date string `json:"date"`
	// Highest highest number.
	// Highest 最高值数值.
	Highest float64 `json:"highest"`
	// Lowest lowest number.
	// Lowest lowest 数值.
	Lowest float64 `json:"lowest"`
	// OrderCount Total number of orders happened that day.
	// OrderCount 当日发生的订单总数.
	OrderCount int64 `json:"order_count"`
	// Volume Total.
	// Volume 总计.
	Volume int64 `json:"volume"`
}

// MarketOrder 200 ok object.
// MarketOrder 200 ok 对象.
type MarketOrder struct {
	// Duration duration integer.
	// Duration 持续时间整数.
	Duration int32 `json:"duration"`
	// IsBuyOrder is_buy_order boolean.
	// IsBuyOrder 是否为买单布尔值.
	IsBuyOrder bool `json:"is_buy_order"`
	// Issued issued string.
	// Issued 发布时间字符串.
	Issued time.Time `json:"issued"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// MinVolume min_volume integer.
	// MinVolume min_volume 整数.
	MinVolume int32 `json:"min_volume"`
	// OrderId order_id integer.
	// OrderId order_id 整数.
	OrderId int64 `json:"order_id"`
	// Price price number.
	// Price price 数值.
	Price float64 `json:"price"`
	// RangeValue range string.
	// RangeValue 射程 string.
	// Enum values: "station", "region", "solarsystem", "1", "2", "3", "4", "5", "10", "20", "30", "40".
	RangeValue string `json:"range"`
	// SystemId The solar system this order was placed.
	// SystemId 此订单所在的星系.
	SystemId int32 `json:"system_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
	// VolumeRemain volume_remain integer.
	// VolumeRemain 剩余体积整数.
	VolumeRemain int32 `json:"volume_remain"`
	// VolumeTotal volume_total integer.
	// VolumeTotal 总体积整数.
	VolumeTotal int32 `json:"volume_total"`
}

// StructureMarketOrder 200 ok object.
// StructureMarketOrder 200 ok 对象.
type StructureMarketOrder struct {
	// Duration duration integer.
	// Duration 持续时间整数.
	Duration int32 `json:"duration"`
	// IsBuyOrder is_buy_order boolean.
	// IsBuyOrder 是否为买单布尔值.
	IsBuyOrder bool `json:"is_buy_order"`
	// Issued issued string.
	// Issued 发布时间字符串.
	Issued time.Time `json:"issued"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// MinVolume min_volume integer.
	// MinVolume min_volume 整数.
	MinVolume int32 `json:"min_volume"`
	// OrderId order_id integer.
	// OrderId order_id 整数.
	OrderId int64 `json:"order_id"`
	// Price price number.
	// Price price 数值.
	Price float64 `json:"price"`
	// RangeValue range string.
	// RangeValue 射程 string.
	// Enum values: "station", "region", "solarsystem", "1", "2", "3", "4", "5", "10", "20", "30", "40".
	RangeValue string `json:"range"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
	// VolumeRemain volume_remain integer.
	// VolumeRemain 剩余体积整数.
	VolumeRemain int32 `json:"volume_remain"`
	// VolumeTotal volume_total integer.
	// VolumeTotal 总体积整数.
	VolumeTotal int32 `json:"volume_total"`
}
