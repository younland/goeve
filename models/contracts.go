package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetCharactersCharacterIdContracts 200 ok object.
// GetCharactersCharacterIdContracts 200 ok 对象.
type GetCharactersCharacterIdContracts struct {
	// AcceptorId Who will accept the contract.
	// AcceptorId 谁将接受该合同.
	AcceptorId int32 `json:"acceptor_id"`
	// AssigneeId ID to whom the contract is assigned, can be alliance, corporation or character ID.
	// AssigneeId 合同被分配对象的 ID，可以是联盟、军团或角色 ID.
	AssigneeId int32 `json:"assignee_id"`
	// Availability To whom the contract is available.
	// Availability 合同对谁开放.
	// Enum values: "public", "personal", "corporation", "alliance".
	Availability string `json:"availability"`
	// Buyout Buyout price (for Auctions only)
	// Buyout 一口价（仅拍卖）
	Buyout float64 `json:"buyout"`
	// Collateral Collateral price (for Couriers only)
	// Collateral 抵押金金额（仅快递合同）
	Collateral float64 `json:"collateral"`
	// ContractId contract_id integer.
	// ContractId 合同 ID 整数.
	ContractId int32 `json:"contract_id"`
	// DateAccepted Date of confirmation of contract.
	// DateAccepted 合同确认的日期.
	DateAccepted time.Time `json:"date_accepted"`
	// DateCompleted Date of completed of contract.
	// DateCompleted 合同完成的日期.
	DateCompleted time.Time `json:"date_completed"`
	// DateExpired Expiration date of the contract.
	// DateExpired 合同到期日期.
	DateExpired time.Time `json:"date_expired"`
	// DateIssued Сreation date of the contract.
	// DateIssued 合同创建日期.
	DateIssued time.Time `json:"date_issued"`
	// DaysToComplete Number of days to perform the contract.
	// DaysToComplete 完成合同所需的天数.
	DaysToComplete int32 `json:"days_to_complete"`
	// EndLocationId End location ID (for Couriers contract)
	// EndLocationId 终点位置 ID（用于快递合同）
	EndLocationId int64 `json:"end_location_id"`
	// ForCorporation true if the contract was issued on behalf of the issuer's corporation.
	// ForCorporation 若合同是以发布者军团的名义发布则为 true.
	ForCorporation bool `json:"for_corporation"`
	// IssuerCorporationId Character's corporation ID for the issuer.
	// IssuerCorporationId 发布者所在军团的军团 ID.
	IssuerCorporationId int32 `json:"issuer_corporation_id"`
	// IssuerId Character ID for the issuer.
	// IssuerId 发布者的角色 ID.
	IssuerId int32 `json:"issuer_id"`
	// Price Price of contract (for ItemsExchange and Auctions)
	// Price 合同价格（仅用于物品交换和拍卖）
	Price float64 `json:"price"`
	// Reward Remuneration for contract (for Couriers only)
	// Reward 合同报酬（仅用于快递合同）
	Reward float64 `json:"reward"`
	// StartLocationId Start location ID (for Couriers contract)
	// StartLocationId 起始地点 ID（用于快递合同）
	StartLocationId int64 `json:"start_location_id"`
	// Status Status of the the contract.
	// Status 合同状态.
	// Enum values: "outstanding", "in_progress", "finished_issuer", "finished_contractor", "finished", "cancelled", "rejected", "failed", "deleted", "reversed".
	Status string `json:"status"`
	// Title Title of the contract.
	// Title 合同的标题.
	Title string `json:"title"`
	// TypeValue Type of the contract.
	// TypeValue 合同的类型.
	// Enum values: "unknown", "item_exchange", "auction", "courier", "loan".
	TypeValue string `json:"type"`
	// Volume Volume of items in the contract.
	// Volume 合同中物品的体积.
	Volume float64 `json:"volume"`
}

// GetCharactersCharacterIdContractsContractIdBids 200 ok object.
// GetCharactersCharacterIdContractsContractIdBids 200 ok 对象.
type GetCharactersCharacterIdContractsContractIdBids struct {
	// Amount The amount bid, in ISK.
	// Amount 出价金额（ISK）
	Amount float64 `json:"amount"`
	// BidId Unique ID for the bid.
	// BidId 竞价的唯一 ID.
	BidId int32 `json:"bid_id"`
	// BidderId Character ID of the bidder.
	// BidderId 出价者的角色 ID.
	BidderId int32 `json:"bidder_id"`
	// DateBid Datetime when the bid was placed.
	// DateBid 出价的时间.
	DateBid time.Time `json:"date_bid"`
}

// GetCharactersCharacterIdContractsContractIdItems 200 ok object.
// GetCharactersCharacterIdContractsContractIdItems 200 ok 对象.
type GetCharactersCharacterIdContractsContractIdItems struct {
	// IsIncluded true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract.
	// IsIncluded 若合同发布者随合同提交了该物品则为 true，若发布者在合同中求购该物品则为 false.
	IsIncluded bool `json:"is_included"`
	// IsSingleton is_singleton boolean.
	// IsSingleton 是否为单一物品布尔值.
	IsSingleton bool `json:"is_singleton"`
	// Quantity Number of items in the stack.
	// Quantity 该堆叠中的物品数量.
	Quantity int32 `json:"quantity"`
	// RawQuantity -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy.
	// RawQuantity -1 表示该物品是单件物品（不可堆叠）。如果该物品恰好是蓝图，-1 表示原版蓝图，-2 表示蓝图拷贝.
	RawQuantity int32 `json:"raw_quantity"`
	// RecordId Unique ID for the item.
	// RecordId 物品的唯一 ID.
	RecordId int64 `json:"record_id"`
	// TypeId Type ID for item.
	// TypeId 物品的 type ID.
	TypeId int32 `json:"type_id"`
}

// GetContractsPublicBidsContractId 200 ok object.
// GetContractsPublicBidsContractId 200 ok 对象.
type GetContractsPublicBidsContractId struct {
	// Amount The amount bid, in ISK.
	// Amount 出价金额（ISK）
	Amount float64 `json:"amount"`
	// BidId Unique ID for the bid.
	// BidId 竞价的唯一 ID.
	BidId int32 `json:"bid_id"`
	// DateBid Datetime when the bid was placed.
	// DateBid 出价的时间.
	DateBid time.Time `json:"date_bid"`
}

// GetContractsPublicItemsContractId 200 ok object.
// GetContractsPublicItemsContractId 200 ok 对象.
type GetContractsPublicItemsContractId struct {
	// IsBlueprintCopy is_blueprint_copy boolean.
	// IsBlueprintCopy 是否为蓝图拷贝布尔值.
	IsBlueprintCopy bool `json:"is_blueprint_copy"`
	// IsIncluded true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract.
	// IsIncluded 若合同发布者随合同提交了该物品则为 true，若发布者在合同中求购该物品则为 false.
	IsIncluded bool `json:"is_included"`
	// ItemId Unique ID for the item being sold. Not present if item is being requested by contract rather than sold with contract.
	// ItemId 出售物品的唯一 ID。如果物品是按合同请求而非随合同出售，则不存在此项.
	ItemId int64 `json:"item_id"`
	// MaterialEfficiency Material Efficiency Level of the blueprint.
	// MaterialEfficiency 蓝图的材料效率等级.
	MaterialEfficiency int32 `json:"material_efficiency"`
	// Quantity Number of items in the stack.
	// Quantity 该堆叠中的物品数量.
	Quantity int32 `json:"quantity"`
	// RecordId Unique ID for the item, used by the contract system.
	// RecordId 物品的唯一 ID，供合同系统使用.
	RecordId int64 `json:"record_id"`
	// Runs Number of runs remaining if the blueprint is a copy, -1 if it is an original.
	// Runs 若蓝图为拷贝则为其剩余运转次数，若为原始蓝图则为 -1.
	Runs int32 `json:"runs"`
	// TimeEfficiency Time Efficiency Level of the blueprint.
	// TimeEfficiency 蓝图的时间效率等级.
	TimeEfficiency int32 `json:"time_efficiency"`
	// TypeId Type ID for item.
	// TypeId 物品的 type ID.
	TypeId int32 `json:"type_id"`
}

// GetContractsPublicRegionId 200 ok object.
// GetContractsPublicRegionId 200 ok 对象.
type GetContractsPublicRegionId struct {
	// Buyout Buyout price (for Auctions only)
	// Buyout 一口价（仅拍卖）
	Buyout float64 `json:"buyout"`
	// Collateral Collateral price (for Couriers only)
	// Collateral 抵押金金额（仅快递合同）
	Collateral float64 `json:"collateral"`
	// ContractId contract_id integer.
	// ContractId 合同 ID 整数.
	ContractId int32 `json:"contract_id"`
	// DateExpired Expiration date of the contract.
	// DateExpired 合同到期日期.
	DateExpired time.Time `json:"date_expired"`
	// DateIssued Сreation date of the contract.
	// DateIssued 合同创建日期.
	DateIssued time.Time `json:"date_issued"`
	// DaysToComplete Number of days to perform the contract.
	// DaysToComplete 完成合同所需的天数.
	DaysToComplete int32 `json:"days_to_complete"`
	// EndLocationId End location ID (for Couriers contract)
	// EndLocationId 终点位置 ID（用于快递合同）
	EndLocationId int64 `json:"end_location_id"`
	// ForCorporation true if the contract was issued on behalf of the issuer's corporation.
	// ForCorporation 若合同是以发布者军团的名义发布则为 true.
	ForCorporation bool `json:"for_corporation"`
	// IssuerCorporationId Character's corporation ID for the issuer.
	// IssuerCorporationId 发布者所在军团的军团 ID.
	IssuerCorporationId int32 `json:"issuer_corporation_id"`
	// IssuerId Character ID for the issuer.
	// IssuerId 发布者的角色 ID.
	IssuerId int32 `json:"issuer_id"`
	// Price Price of contract (for ItemsExchange and Auctions)
	// Price 合同价格（仅用于物品交换和拍卖）
	Price float64 `json:"price"`
	// Reward Remuneration for contract (for Couriers only)
	// Reward 合同报酬（仅用于快递合同）
	Reward float64 `json:"reward"`
	// StartLocationId Start location ID (for Couriers contract)
	// StartLocationId 起始地点 ID（用于快递合同）
	StartLocationId int64 `json:"start_location_id"`
	// Title Title of the contract.
	// Title 合同的标题.
	Title string `json:"title"`
	// TypeValue Type of the contract.
	// TypeValue 合同的类型.
	// Enum values: "unknown", "item_exchange", "auction", "courier", "loan".
	TypeValue string `json:"type"`
	// Volume Volume of items in the contract.
	// Volume 合同中物品的体积.
	Volume float64 `json:"volume"`
}

// GetCorporationsCorporationIdContracts 200 ok object.
// GetCorporationsCorporationIdContracts 200 ok 对象.
type GetCorporationsCorporationIdContracts struct {
	// AcceptorId Who will accept the contract.
	// AcceptorId 谁将接受该合同.
	AcceptorId int32 `json:"acceptor_id"`
	// AssigneeId ID to whom the contract is assigned, can be corporation or character ID.
	// AssigneeId 合同被分配对象的 ID，可以是军团或角色 ID.
	AssigneeId int32 `json:"assignee_id"`
	// Availability To whom the contract is available.
	// Availability 合同对谁开放.
	// Enum values: "public", "personal", "corporation", "alliance".
	Availability string `json:"availability"`
	// Buyout Buyout price (for Auctions only)
	// Buyout 一口价（仅拍卖）
	Buyout float64 `json:"buyout"`
	// Collateral Collateral price (for Couriers only)
	// Collateral 抵押金金额（仅快递合同）
	Collateral float64 `json:"collateral"`
	// ContractId contract_id integer.
	// ContractId 合同 ID 整数.
	ContractId int32 `json:"contract_id"`
	// DateAccepted Date of confirmation of contract.
	// DateAccepted 合同确认的日期.
	DateAccepted time.Time `json:"date_accepted"`
	// DateCompleted Date of completed of contract.
	// DateCompleted 合同完成的日期.
	DateCompleted time.Time `json:"date_completed"`
	// DateExpired Expiration date of the contract.
	// DateExpired 合同到期日期.
	DateExpired time.Time `json:"date_expired"`
	// DateIssued Сreation date of the contract.
	// DateIssued 合同创建日期.
	DateIssued time.Time `json:"date_issued"`
	// DaysToComplete Number of days to perform the contract.
	// DaysToComplete 完成合同所需的天数.
	DaysToComplete int32 `json:"days_to_complete"`
	// EndLocationId End location ID (for Couriers contract)
	// EndLocationId 终点位置 ID（用于快递合同）
	EndLocationId int64 `json:"end_location_id"`
	// ForCorporation true if the contract was issued on behalf of the issuer's corporation.
	// ForCorporation 若合同是以发布者军团的名义发布则为 true.
	ForCorporation bool `json:"for_corporation"`
	// IssuerCorporationId Character's corporation ID for the issuer.
	// IssuerCorporationId 发布者所在军团的军团 ID.
	IssuerCorporationId int32 `json:"issuer_corporation_id"`
	// IssuerId Character ID for the issuer.
	// IssuerId 发布者的角色 ID.
	IssuerId int32 `json:"issuer_id"`
	// Price Price of contract (for ItemsExchange and Auctions)
	// Price 合同价格（仅用于物品交换和拍卖）
	Price float64 `json:"price"`
	// Reward Remuneration for contract (for Couriers only)
	// Reward 合同报酬（仅用于快递合同）
	Reward float64 `json:"reward"`
	// StartLocationId Start location ID (for Couriers contract)
	// StartLocationId 起始地点 ID（用于快递合同）
	StartLocationId int64 `json:"start_location_id"`
	// Status Status of the the contract.
	// Status 合同状态.
	// Enum values: "outstanding", "in_progress", "finished_issuer", "finished_contractor", "finished", "cancelled", "rejected", "failed", "deleted", "reversed".
	Status string `json:"status"`
	// Title Title of the contract.
	// Title 合同的标题.
	Title string `json:"title"`
	// TypeValue Type of the contract.
	// TypeValue 合同的类型.
	// Enum values: "unknown", "item_exchange", "auction", "courier", "loan".
	TypeValue string `json:"type"`
	// Volume Volume of items in the contract.
	// Volume 合同中物品的体积.
	Volume float64 `json:"volume"`
}

// GetCorporationsCorporationIdContractsContractIdBids 200 ok object.
// GetCorporationsCorporationIdContractsContractIdBids 200 ok 对象.
type GetCorporationsCorporationIdContractsContractIdBids struct {
	// Amount The amount bid, in ISK.
	// Amount 出价金额（ISK）
	Amount float64 `json:"amount"`
	// BidId Unique ID for the bid.
	// BidId 竞价的唯一 ID.
	BidId int32 `json:"bid_id"`
	// BidderId Character ID of the bidder.
	// BidderId 出价者的角色 ID.
	BidderId int32 `json:"bidder_id"`
	// DateBid Datetime when the bid was placed.
	// DateBid 出价的时间.
	DateBid time.Time `json:"date_bid"`
}

// GetCorporationsCorporationIdContractsContractIdItems 200 ok object.
// GetCorporationsCorporationIdContractsContractIdItems 200 ok 对象.
type GetCorporationsCorporationIdContractsContractIdItems struct {
	// IsIncluded true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract.
	// IsIncluded 若合同发布者随合同提交了该物品则为 true，若发布者在合同中求购该物品则为 false.
	IsIncluded bool `json:"is_included"`
	// IsSingleton is_singleton boolean.
	// IsSingleton 是否为单一物品布尔值.
	IsSingleton bool `json:"is_singleton"`
	// Quantity Number of items in the stack.
	// Quantity 该堆叠中的物品数量.
	Quantity int32 `json:"quantity"`
	// RawQuantity -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy.
	// RawQuantity -1 表示该物品是单件物品（不可堆叠）。如果该物品恰好是蓝图，-1 表示原版蓝图，-2 表示蓝图拷贝.
	RawQuantity int32 `json:"raw_quantity"`
	// RecordId Unique ID for the item.
	// RecordId 物品的唯一 ID.
	RecordId int64 `json:"record_id"`
	// TypeId Type ID for item.
	// TypeId 物品的 type ID.
	TypeId int32 `json:"type_id"`
}

// GetCharactersCharacterIdContractsContractIdBidsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdContractsContractIdBidsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdContractsContractIdBidsParams struct {
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

func (p *GetCharactersCharacterIdContractsContractIdBidsParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdContractsContractIdItemsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdContractsContractIdItemsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdContractsContractIdItemsParams struct {
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

func (p *GetCharactersCharacterIdContractsContractIdItemsParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdContractsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdContractsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdContractsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdContractsParams) Values() (url.Values, map[string]string) {
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
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCorporationsCorporationIdContractsContractIdBidsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdContractsContractIdBidsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdContractsContractIdBidsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCorporationsCorporationIdContractsContractIdBidsParams) Values() (url.Values, map[string]string) {
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
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCorporationsCorporationIdContractsContractIdItemsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdContractsContractIdItemsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdContractsContractIdItemsParams struct {
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

func (p *GetCorporationsCorporationIdContractsContractIdItemsParams) Values() (url.Values, map[string]string) {
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

// GetCorporationsCorporationIdContractsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdContractsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdContractsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCorporationsCorporationIdContractsParams) Values() (url.Values, map[string]string) {
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
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetPublicBidsContractIdParams holds the optional query and header parameters of the request.
// GetPublicBidsContractIdParams 保存请求的可选查询与头部参数。
type GetPublicBidsContractIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
}

func (p *GetPublicBidsContractIdParams) Values() (url.Values, map[string]string) {
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
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	return query, headers
}

// GetPublicItemsContractIdParams holds the optional query and header parameters of the request.
// GetPublicItemsContractIdParams 保存请求的可选查询与头部参数。
type GetPublicItemsContractIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
}

func (p *GetPublicItemsContractIdParams) Values() (url.Values, map[string]string) {
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
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	return query, headers
}

// GetPublicRegionIdParams holds the optional query and header parameters of the request.
// GetPublicRegionIdParams 保存请求的可选查询与头部参数。
type GetPublicRegionIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
}

func (p *GetPublicRegionIdParams) Values() (url.Values, map[string]string) {
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
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	return query, headers
}
