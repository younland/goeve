package models

import "net/url"

// GetAttributesAttributeIdParams holds the optional query and header parameters of the request.
// GetAttributesAttributeIdParams 保存请求的可选查询与头部参数。
type GetAttributesAttributeIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetAttributesAttributeIdParams) Values() (url.Values, map[string]string) {
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

// GetAttributesParams holds the optional query and header parameters of the request.
// GetAttributesParams 保存请求的可选查询与头部参数。
type GetAttributesParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetAttributesParams) Values() (url.Values, map[string]string) {
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

// GetDogmaAttributesAttributeIdOk 200 ok object.
// GetDogmaAttributesAttributeIdOk 200 ok 对象.
type GetDogmaAttributesAttributeIdOk struct {
	// AttributeId attribute_id integer.
	// AttributeId 属性 ID 整数.
	AttributeId int32 `json:"attribute_id"`
	// DefaultValue default_value number.
	// DefaultValue 默认值数值.
	DefaultValue float64 `json:"default_value"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// DisplayName display_name string.
	// DisplayName 显示名称字符串.
	DisplayName string `json:"display_name"`
	// HighIsGood high_is_good boolean.
	// HighIsGood 数值越高越好布尔值.
	HighIsGood bool `json:"high_is_good"`
	// IconId icon_id integer.
	// IconId 图标 ID 整数.
	IconId int32 `json:"icon_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Published published boolean.
	// Published published 布尔.
	Published bool `json:"published"`
	// Stackable stackable boolean.
	// Stackable 可堆叠 boolean.
	Stackable bool `json:"stackable"`
	// UnitId unit_id integer.
	// UnitId unit_id 整数.
	UnitId int32 `json:"unit_id"`
}

// GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute dogma_attribute object.
// GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute 教条属性对象.
type GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute struct {
	// AttributeId attribute_id integer.
	// AttributeId 属性 ID 整数.
	AttributeId int32 `json:"attribute_id"`
	// Value value number.
	// Value 价值数字.
	Value float64 `json:"value"`
}

// GetDogmaDynamicItemsTypeIdItemIdDogmaEffect dogma_effect object.
// GetDogmaDynamicItemsTypeIdItemIdDogmaEffect 教条效果对象.
type GetDogmaDynamicItemsTypeIdItemIdDogmaEffect struct {
	// EffectId effect_id integer.
	// EffectId 效果 ID 整数.
	EffectId int32 `json:"effect_id"`
	// IsDefault is_default boolean.
	// IsDefault 是否为默认布尔值.
	IsDefault bool `json:"is_default"`
}

// GetDogmaDynamicItemsTypeIdItemIdOk 200 ok object.
// GetDogmaDynamicItemsTypeIdItemIdOk 200 ok 对象.
type GetDogmaDynamicItemsTypeIdItemIdOk struct {
	// CreatedBy The ID of the character who created the item.
	// CreatedBy 创建该物品的角色 ID.
	CreatedBy int32 `json:"created_by"`
	// DogmaAttributes dogma_attributes array.
	// DogmaAttributes 教条属性数组.
	DogmaAttributes []GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute `json:"dogma_attributes"`
	// DogmaEffects dogma_effects array.
	// DogmaEffects 教条效果数组.
	DogmaEffects []GetDogmaDynamicItemsTypeIdItemIdDogmaEffect `json:"dogma_effects"`
	// MutatorTypeId The type ID of the mutator used to generate the dynamic item.
	// MutatorTypeId 用于生成动态物品的变换器（mutator）的 type ID。
	MutatorTypeId int32 `json:"mutator_type_id"`
	// SourceTypeId The type ID of the source item the mutator was applied to create the dynamic item.
	// SourceTypeId 应用变换器以创建动态物品的源物品的 type ID。
	SourceTypeId int32 `json:"source_type_id"`
}

// GetDogmaEffectsEffectIdModifier modifier object.
// GetDogmaEffectsEffectIdModifier modifier 对象.
type GetDogmaEffectsEffectIdModifier struct {
	// Domain domain string.
	// Domain 作用域字符串.
	Domain string `json:"domain"`
	// EffectId effect_id integer.
	// EffectId 效果 ID 整数.
	EffectId int32 `json:"effect_id"`
	// FuncValue func string.
	// FuncValue 函数字符串.
	FuncValue string `json:"func"`
	// ModifiedAttributeId modified_attribute_id integer.
	// ModifiedAttributeId modified_attribute_id 整数.
	ModifiedAttributeId int32 `json:"modified_attribute_id"`
	// ModifyingAttributeId modifying_attribute_id integer.
	// ModifyingAttributeId modifying_attribute_id 整数.
	ModifyingAttributeId int32 `json:"modifying_attribute_id"`
	// Operator operator integer.
	// Operator operator 整数.
	Operator int32 `json:"operator"`
}

// GetDogmaEffectsEffectIdOk 200 ok object.
// GetDogmaEffectsEffectIdOk 200 ok 对象.
type GetDogmaEffectsEffectIdOk struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// DisallowAutoRepeat disallow_auto_repeat boolean.
	// DisallowAutoRepeat 禁止自动重复布尔值.
	DisallowAutoRepeat bool `json:"disallow_auto_repeat"`
	// DischargeAttributeId discharge_attribute_id integer.
	// DischargeAttributeId 耗散属性 ID 整数.
	DischargeAttributeId int32 `json:"discharge_attribute_id"`
	// DisplayName display_name string.
	// DisplayName 显示名称字符串.
	DisplayName string `json:"display_name"`
	// DurationAttributeId duration_attribute_id integer.
	// DurationAttributeId 持续时间属性 ID 整数.
	DurationAttributeId int32 `json:"duration_attribute_id"`
	// EffectCategory effect_category integer.
	// EffectCategory 效果类别整数.
	EffectCategory int32 `json:"effect_category"`
	// EffectId effect_id integer.
	// EffectId 效果 ID 整数.
	EffectId int32 `json:"effect_id"`
	// ElectronicChance electronic_chance boolean.
	// ElectronicChance 电子成功概率布尔值.
	ElectronicChance bool `json:"electronic_chance"`
	// FalloffAttributeId falloff_attribute_id integer.
	// FalloffAttributeId 衰减距离属性 ID 整数.
	FalloffAttributeId int32 `json:"falloff_attribute_id"`
	// IconId icon_id integer.
	// IconId 图标 ID 整数.
	IconId int32 `json:"icon_id"`
	// IsAssistance is_assistance boolean.
	// IsAssistance 是否为援助布尔值.
	IsAssistance bool `json:"is_assistance"`
	// IsOffensive is_offensive boolean.
	// IsOffensive 是否为攻击型布尔值.
	IsOffensive bool `json:"is_offensive"`
	// IsWarpSafe is_warp_safe boolean.
	// IsWarpSafe 是否跃迁安全布尔值.
	IsWarpSafe bool `json:"is_warp_safe"`
	// Modifiers modifiers array.
	// Modifiers modifiers 数组.
	Modifiers []GetDogmaEffectsEffectIdModifier `json:"modifiers"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// PostExpression post_expression integer.
	// PostExpression post_expression 整数.
	PostExpression int32 `json:"post_expression"`
	// PreExpression pre_expression integer.
	// PreExpression pre_expression 整数.
	PreExpression int32 `json:"pre_expression"`
	// Published published boolean.
	// Published published 布尔.
	Published bool `json:"published"`
	// RangeAttributeId range_attribute_id integer.
	// RangeAttributeId 射程属性ID integer.
	RangeAttributeId int32 `json:"range_attribute_id"`
	// RangeChance range_chance boolean.
	// RangeChance 射程几率 boolean.
	RangeChance bool `json:"range_chance"`
	// TrackingSpeedAttributeId tracking_speed_attribute_id integer.
	// TrackingSpeedAttributeId tracking_speed_attribute_id 整数.
	TrackingSpeedAttributeId int32 `json:"tracking_speed_attribute_id"`
}

// GetDynamicItemsTypeIdItemIdParams holds the optional query and header parameters of the request.
// GetDynamicItemsTypeIdItemIdParams 保存请求的可选查询与头部参数。
type GetDynamicItemsTypeIdItemIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetDynamicItemsTypeIdItemIdParams) Values() (url.Values, map[string]string) {
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

// GetEffectsEffectIdParams holds the optional query and header parameters of the request.
// GetEffectsEffectIdParams 保存请求的可选查询与头部参数。
type GetEffectsEffectIdParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetEffectsEffectIdParams) Values() (url.Values, map[string]string) {
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

// GetEffectsParams holds the optional query and header parameters of the request.
// GetEffectsParams 保存请求的可选查询与头部参数。
type GetEffectsParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetEffectsParams) Values() (url.Values, map[string]string) {
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
