package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetCharactersCharacterIdWalletJournal 200 ok object.
// GetCharactersCharacterIdWalletJournal 200 ok 对象.
type GetCharactersCharacterIdWalletJournal struct {
	// Amount The amount of ISK given or taken from the wallet as a result of the given transaction. Positive when ISK is deposited into the wallet and negative when ISK is withdrawn.
	// Amount 指定交易导致钱包转入或转出的 ISK 金额。ISK 存入钱包时为正，取出时为负.
	Amount float64 `json:"amount"`
	// Balance Wallet balance after transaction occurred.
	// Balance 交易发生后的钱包余额.
	Balance float64 `json:"balance"`
	// ContextId An ID that gives extra context to the particular transaction. Because of legacy reasons the context is completely different per ref_type and means different things. It is also possible to not have a context_id.
	// ContextId 为该特定交易提供额外上下文的 ID。由于历史原因，其上下文完全因 ref_type 而异，含义各不相同。也可能不存在 context_id.
	ContextId int64 `json:"context_id"`
	// ContextIdType The type of the given context_id if present.
	// ContextIdType 给定的 context_id 的类型（如果存在）
	// Enum values: "structure_id", "station_id", "market_transaction_id", "character_id", "corporation_id", "alliance_id", "eve_system", "industry_job_id", "contract_id", "planet_id", "system_id", "type_id".
	ContextIdType string `json:"context_id_type"`
	// Date Date and time of transaction.
	// Date 交易日期和时间.
	Date time.Time `json:"date"`
	// Description The reason for the transaction, mirrors what is seen in the client.
	// Description 交易原因，与客户端中显示的一致.
	Description string `json:"description"`
	// FirstPartyId The id of the first party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name.
	// FirstPartyId 交易中第一方的 ID。此属性不具有一致性，对特定 ref_types 而言其值不同或不存在。description 属性有助于理解此属性的含义。可将该 ID 放入 /universe/names/ ESI 路由查询，以确定其类型和名称.
	FirstPartyId int32 `json:"first_party_id"`
	// Id Unique journal reference ID.
	// Id 账本记录引用的唯一 ID.
	Id int64 `json:"id"`
	// Reason The user stated reason for the transaction. Only applies to some ref_types.
	// Reason 用户声明的交易原因。仅适用于部分 ref_types.
	Reason string `json:"reason"`
	// RefType "The transaction type for the given. transaction. Different transaction types will populate different attributes.".
	// RefType "给定交易的交易类型。不同的交易类型会填充不同的属性。".
	// Enum values: "acceleration_gate_fee", "advertisement_listing_fee", "agent_donation", "agent_location_services", "agent_miscellaneous", "agent_mission_collateral_paid", "agent_mission_collateral_refunded", "agent_mission_reward", "agent_mission_reward_corporation_tax", "agent_mission_time_bonus_reward", "agent_mission_time_bonus_reward_corporation_tax", "agent_security_services", "agent_services_rendered", "agents_preward", "alliance_maintainance_fee", "alliance_registration_fee", "asset_safety_recovery_tax", "bounty", "bounty_prize", "bounty_prize_corporation_tax", "bounty_prizes", "bounty_reimbursement", "bounty_surcharge", "brokers_fee", "clone_activation", "clone_transfer", "contraband_fine", "contract_auction_bid", "contract_auction_bid_corp", "contract_auction_bid_refund", "contract_auction_sold", "contract_brokers_fee", "contract_brokers_fee_corp", "contract_collateral", "contract_collateral_deposited_corp", "contract_collateral_payout", "contract_collateral_refund", "contract_deposit", "contract_deposit_corp", "contract_deposit_refund", "contract_deposit_sales_tax", "contract_price", "contract_price_payment_corp", "contract_reversal", "contract_reward", "contract_reward_deposited", "contract_reward_deposited_corp", "contract_reward_refund", "contract_sales_tax", "copying", "corporate_reward_payout", "corporate_reward_tax", "corporation_account_withdrawal", "corporation_bulk_payment", "corporation_dividend_payment", "corporation_liquidation", "corporation_logo_change_cost", "corporation_payment", "corporation_registration_fee", "courier_mission_escrow", "cspa", "cspaofflinerefund", "daily_challenge_reward", "datacore_fee", "dna_modification_fee", "docking_fee", "duel_wager_escrow", "duel_wager_payment", "duel_wager_refund", "ess_escrow_transfer", "external_trade_delivery", "external_trade_freeze", "external_trade_thaw", "factory_slot_rental_fee", "flux_payout", "flux_tax", "flux_ticket_repayment", "flux_ticket_sale", "gm_cash_transfer", "industry_job_tax", "infrastructure_hub_maintenance", "inheritance", "insurance", "item_trader_payment", "jump_clone_activation_fee", "jump_clone_installation_fee", "kill_right_fee", "lp_store", "manufacturing", "market_escrow", "market_fine_paid", "market_provider_tax", "market_transaction", "medal_creation", "medal_issued", "milestone_reward_payment", "mission_completion", "mission_cost", "mission_expiration", "mission_reward", "office_rental_fee", "operation_bonus", "opportunity_reward", "planetary_construction", "planetary_export_tax", "planetary_import_tax", "player_donation", "player_trading", "project_discovery_reward", "project_discovery_tax", "reaction", "redeemed_isk_token", "release_of_impounded_property", "repair_bill", "reprocessing_tax", "researching_material_productivity", "researching_technology", "researching_time_productivity", "resource_wars_reward", "reverse_engineering", "season_challenge_reward", "security_processing_fee", "shares", "skill_purchase", "sovereignity_bill", "store_purchase", "store_purchase_refund", "structure_gate_jump", "transaction_tax", "upkeep_adjustment_fee", "war_ally_contract", "war_fee", "war_fee_surrender".
	RefType string `json:"ref_type"`
	// SecondPartyId The id of the second party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name.
	// SecondPartyId 交易中第二方的 ID。此属性不具有一致性，对特定 ref_types 而言其值不同或不存在。description 属性有助于理解此属性的含义。可将该 ID 放入 /universe/names/ ESI 路由查询，以确定其类型和名称.
	SecondPartyId int32 `json:"second_party_id"`
	// Tax Tax amount received. Only applies to tax related transactions.
	// Tax 收到的税额，仅适用于税收相关交易.
	Tax float64 `json:"tax"`
	// TaxReceiverId The corporation ID receiving any tax paid. Only applies to tax related transactions.
	// TaxReceiverId 收取所缴税款的军团 ID，仅适用于税收相关交易.
	TaxReceiverId int32 `json:"tax_receiver_id"`
}

// GetCharactersCharacterIdWalletTransactions wallet transaction.
// GetCharactersCharacterIdWalletTransactions 钱包交易.
type GetCharactersCharacterIdWalletTransactions struct {
	// ClientId client_id integer.
	// ClientId 客户 ID 整数.
	ClientId int32 `json:"client_id"`
	// Date Date and time of transaction.
	// Date 交易日期和时间.
	Date time.Time `json:"date"`
	// IsBuy is_buy boolean.
	// IsBuy 是否为买入布尔值.
	IsBuy bool `json:"is_buy"`
	// IsPersonal is_personal boolean.
	// IsPersonal 是否为个人布尔值.
	IsPersonal bool `json:"is_personal"`
	// JournalRefId journal_ref_id integer.
	// JournalRefId journal_ref_id 整数.
	JournalRefId int64 `json:"journal_ref_id"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TransactionId Unique transaction ID.
	// TransactionId 交易的唯一 ID.
	TransactionId int64 `json:"transaction_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
	// UnitPrice Amount paid per unit.
	// UnitPrice 每单位支付的金额.
	UnitPrice float64 `json:"unit_price"`
}

// GetCharactersCharacterIdWalletJournalParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdWalletJournalParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdWalletJournalParams struct {
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

func (p *GetCharactersCharacterIdWalletJournalParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdWalletParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdWalletParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdWalletParams struct {
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

func (p *GetCharactersCharacterIdWalletParams) Values() (url.Values, map[string]string) {
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

// GetCharactersCharacterIdWalletTransactionsParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdWalletTransactionsParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdWalletTransactionsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// FromId Only show transactions happened before the one referenced by this id.
	// FromId 仅显示发生在该 ID 所引用交易之前的交易.
	FromId *int64
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdWalletTransactionsParams) Values() (url.Values, map[string]string) {
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
	if p.FromId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("from_id", strconv.FormatInt(int64(*p.FromId), 10))
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

// GetCorporationsCorporationIdWallets 200 ok object.
// GetCorporationsCorporationIdWallets 200 ok 对象.
type GetCorporationsCorporationIdWallets struct {
	// Balance balance number.
	// Balance 余额数值.
	Balance float64 `json:"balance"`
	// Division division integer.
	// Division 部门整数.
	Division int32 `json:"division"`
}

// GetCorporationsCorporationIdWalletsDivisionJournal 200 ok object.
// GetCorporationsCorporationIdWalletsDivisionJournal 200 ok 对象.
type GetCorporationsCorporationIdWalletsDivisionJournal struct {
	// Amount The amount of ISK given or taken from the wallet as a result of the given transaction. Positive when ISK is deposited into the wallet and negative when ISK is withdrawn.
	// Amount 指定交易导致钱包转入或转出的 ISK 金额。ISK 存入钱包时为正，取出时为负.
	Amount float64 `json:"amount"`
	// Balance Wallet balance after transaction occurred.
	// Balance 交易发生后的钱包余额.
	Balance float64 `json:"balance"`
	// ContextId An ID that gives extra context to the particular transaction. Because of legacy reasons the context is completely different per ref_type and means different things. It is also possible to not have a context_id.
	// ContextId 为该特定交易提供额外上下文的 ID。由于历史原因，其上下文完全因 ref_type 而异，含义各不相同。也可能不存在 context_id.
	ContextId int64 `json:"context_id"`
	// ContextIdType The type of the given context_id if present.
	// ContextIdType 给定的 context_id 的类型（如果存在）
	// Enum values: "structure_id", "station_id", "market_transaction_id", "character_id", "corporation_id", "alliance_id", "eve_system", "industry_job_id", "contract_id", "planet_id", "system_id", "type_id".
	ContextIdType string `json:"context_id_type"`
	// Date Date and time of transaction.
	// Date 交易日期和时间.
	Date time.Time `json:"date"`
	// Description The reason for the transaction, mirrors what is seen in the client.
	// Description 交易原因，与客户端中显示的一致.
	Description string `json:"description"`
	// FirstPartyId The id of the first party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name.
	// FirstPartyId 交易中第一方的 ID。此属性不具有一致性，对特定 ref_types 而言其值不同或不存在。description 属性有助于理解此属性的含义。可将该 ID 放入 /universe/names/ ESI 路由查询，以确定其类型和名称.
	FirstPartyId int32 `json:"first_party_id"`
	// Id Unique journal reference ID.
	// Id 账本记录引用的唯一 ID.
	Id int64 `json:"id"`
	// Reason The user stated reason for the transaction. Only applies to some ref_types.
	// Reason 用户声明的交易原因。仅适用于部分 ref_types.
	Reason string `json:"reason"`
	// RefType "The transaction type for the given. transaction. Different transaction types will populate different attributes. Note: If you have an existing XML API application that is using ref_types, you will need to know which string ESI ref_type maps to which integer. You can look at the following file to see string->int mappings: https://github.com/ccpgames/eve-glue/blob/master/eve_glue/wallet_journal_ref.py".
	// RefType "给定交易的交易类型。不同的交易类型会填充不同的属性。注意：如果你已有使用 ref_types 的 XML API 应用，你需要了解 ESI 的 ref_type 中哪个字符串对应哪个整数。可以查看以下文件了解字符串到整数的映射：https://github.com/ccpgames/eve-glue/blob/master/eve_glue/wallet_journal_ref.py".
	// Enum values: "acceleration_gate_fee", "advertisement_listing_fee", "agent_donation", "agent_location_services", "agent_miscellaneous", "agent_mission_collateral_paid", "agent_mission_collateral_refunded", "agent_mission_reward", "agent_mission_reward_corporation_tax", "agent_mission_time_bonus_reward", "agent_mission_time_bonus_reward_corporation_tax", "agent_security_services", "agent_services_rendered", "agents_preward", "alliance_maintainance_fee", "alliance_registration_fee", "asset_safety_recovery_tax", "bounty", "bounty_prize", "bounty_prize_corporation_tax", "bounty_prizes", "bounty_reimbursement", "bounty_surcharge", "brokers_fee", "clone_activation", "clone_transfer", "contraband_fine", "contract_auction_bid", "contract_auction_bid_corp", "contract_auction_bid_refund", "contract_auction_sold", "contract_brokers_fee", "contract_brokers_fee_corp", "contract_collateral", "contract_collateral_deposited_corp", "contract_collateral_payout", "contract_collateral_refund", "contract_deposit", "contract_deposit_corp", "contract_deposit_refund", "contract_deposit_sales_tax", "contract_price", "contract_price_payment_corp", "contract_reversal", "contract_reward", "contract_reward_deposited", "contract_reward_deposited_corp", "contract_reward_refund", "contract_sales_tax", "copying", "corporate_reward_payout", "corporate_reward_tax", "corporation_account_withdrawal", "corporation_bulk_payment", "corporation_dividend_payment", "corporation_liquidation", "corporation_logo_change_cost", "corporation_payment", "corporation_registration_fee", "courier_mission_escrow", "cspa", "cspaofflinerefund", "daily_challenge_reward", "datacore_fee", "dna_modification_fee", "docking_fee", "duel_wager_escrow", "duel_wager_payment", "duel_wager_refund", "ess_escrow_transfer", "external_trade_delivery", "external_trade_freeze", "external_trade_thaw", "factory_slot_rental_fee", "flux_payout", "flux_tax", "flux_ticket_repayment", "flux_ticket_sale", "gm_cash_transfer", "industry_job_tax", "infrastructure_hub_maintenance", "inheritance", "insurance", "item_trader_payment", "jump_clone_activation_fee", "jump_clone_installation_fee", "kill_right_fee", "lp_store", "manufacturing", "market_escrow", "market_fine_paid", "market_provider_tax", "market_transaction", "medal_creation", "medal_issued", "milestone_reward_payment", "mission_completion", "mission_cost", "mission_expiration", "mission_reward", "office_rental_fee", "operation_bonus", "opportunity_reward", "planetary_construction", "planetary_export_tax", "planetary_import_tax", "player_donation", "player_trading", "project_discovery_reward", "project_discovery_tax", "reaction", "redeemed_isk_token", "release_of_impounded_property", "repair_bill", "reprocessing_tax", "researching_material_productivity", "researching_technology", "researching_time_productivity", "resource_wars_reward", "reverse_engineering", "season_challenge_reward", "security_processing_fee", "shares", "skill_purchase", "sovereignity_bill", "store_purchase", "store_purchase_refund", "structure_gate_jump", "transaction_tax", "upkeep_adjustment_fee", "war_ally_contract", "war_fee", "war_fee_surrender".
	RefType string `json:"ref_type"`
	// SecondPartyId The id of the second party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name.
	// SecondPartyId 交易中第二方的 ID。此属性不具有一致性，对特定 ref_types 而言其值不同或不存在。description 属性有助于理解此属性的含义。可将该 ID 放入 /universe/names/ ESI 路由查询，以确定其类型和名称.
	SecondPartyId int32 `json:"second_party_id"`
	// Tax Tax amount received. Only applies to tax related transactions.
	// Tax 收到的税额，仅适用于税收相关交易.
	Tax float64 `json:"tax"`
	// TaxReceiverId The corporation ID receiving any tax paid. Only applies to tax related transactions.
	// TaxReceiverId 收取所缴税款的军团 ID，仅适用于税收相关交易.
	TaxReceiverId int32 `json:"tax_receiver_id"`
}

// GetCorporationsCorporationIdWalletsDivisionTransactions wallet transaction.
// GetCorporationsCorporationIdWalletsDivisionTransactions 钱包交易.
type GetCorporationsCorporationIdWalletsDivisionTransactions struct {
	// ClientId client_id integer.
	// ClientId 客户 ID 整数.
	ClientId int32 `json:"client_id"`
	// Date Date and time of transaction.
	// Date 交易日期和时间.
	Date time.Time `json:"date"`
	// IsBuy is_buy boolean.
	// IsBuy 是否为买入布尔值.
	IsBuy bool `json:"is_buy"`
	// JournalRefId -1 if there is no corresponding wallet journal entry.
	// JournalRefId 如果没有对应的钱包日志记录则为 -1.
	JournalRefId int64 `json:"journal_ref_id"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int64 `json:"location_id"`
	// Quantity quantity integer.
	// Quantity 数量 integer.
	Quantity int32 `json:"quantity"`
	// TransactionId Unique transaction ID.
	// TransactionId 交易的唯一 ID.
	TransactionId int64 `json:"transaction_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
	// UnitPrice Amount paid per unit.
	// UnitPrice 每单位支付的金额.
	UnitPrice float64 `json:"unit_price"`
}

// GetCorporationsCorporationIdWalletsDivisionJournalParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdWalletsDivisionJournalParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdWalletsDivisionJournalParams struct {
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

func (p *GetCorporationsCorporationIdWalletsDivisionJournalParams) Values() (url.Values, map[string]string) {
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

// GetCorporationsCorporationIdWalletsDivisionTransactionsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdWalletsDivisionTransactionsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdWalletsDivisionTransactionsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// FromId Only show journal entries happened before the transaction referenced by this id.
	// FromId 仅显示发生在该 ID 所引用交易之前的账单条目.
	FromId *int64
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCorporationsCorporationIdWalletsDivisionTransactionsParams) Values() (url.Values, map[string]string) {
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
	if p.FromId != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("from_id", strconv.FormatInt(int64(*p.FromId), 10))
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

// GetCorporationsCorporationIdWalletsParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdWalletsParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdWalletsParams struct {
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

func (p *GetCorporationsCorporationIdWalletsParams) Values() (url.Values, map[string]string) {
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
