package models

import (
	"net/url"
)

// GetInsurancePrices 200 ok object.
// GetInsurancePrices 200 ok 对象.
type GetInsurancePrices struct {
	// Levels A list of a available insurance levels for this ship type.
	// Levels 该舰船类型可用的保险等级列表.
	Levels []GetInsurancePricesLevel `json:"levels"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetInsurancePricesLevel level object.
// GetInsurancePricesLevel level 对象.
type GetInsurancePricesLevel struct {
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

// GetPricesParams holds the optional query and header parameters of the request.
// GetPricesParams 保存请求的可选查询与头部参数。
type GetPricesParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
}

func (p *GetPricesParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.AcceptLanguage != nil && *p.AcceptLanguage != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Accept-Language"] = *p.AcceptLanguage
	}
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
	if p.Language != nil && *p.Language != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("language", *p.Language)
	}
	return query, headers
}
