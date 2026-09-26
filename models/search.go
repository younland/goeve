package models

import (
	"net/url"
	"strconv"
)

// GetCharactersCharacterIdSearchOk 200 ok object.
// GetCharactersCharacterIdSearchOk 200 ok 对象.
type GetCharactersCharacterIdSearchOk struct {
	// Agent agent array.
	// Agent 代理人数组.
	Agent []int32 `json:"agent"`
	// Alliance alliance array.
	// Alliance 联盟数组.
	Alliance []int32 `json:"alliance"`
	// Character character array.
	// Character 角色数组.
	Character []int32 `json:"character"`
	// Constellation constellation array.
	// Constellation 星座数组.
	Constellation []int32 `json:"constellation"`
	// Corporation corporation array.
	// Corporation 军团数组.
	Corporation []int32 `json:"corporation"`
	// Faction faction array.
	// Faction 势力数组.
	Faction []int32 `json:"faction"`
	// InventoryType inventory_type array.
	// InventoryType 物品类型数组.
	InventoryType []int32 `json:"inventory_type"`
	// Region region array.
	// Region 星域 array.
	Region []int32 `json:"region"`
	// SolarSystem solar_system array.
	// SolarSystem 星系 array.
	SolarSystem []int32 `json:"solar_system"`
	// Station station array.
	// Station 空间站 array.
	Station []int32 `json:"station"`
	// Structure structure array.
	// Structure 建筑 array.
	Structure []int64 `json:"structure"`
}

// GetCharactersCharacterIdSearchParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdSearchParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdSearchParams struct {
	// AcceptLanguage Language to use in the response.
	// AcceptLanguage 响应中使用的语言.
	AcceptLanguage *string
	// Categories Type of entities to search for.
	// Categories 要搜索的实体类型.
	Categories []string
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Language Language to use in the response, takes precedence over Accept-Language.
	// Language 响应中使用的语言，优先于 Accept-Language.
	Language *string
	// Search The string to search on.
	// Search 要搜索的字符串.
	Search *string
	// Strict Whether the search should be a strict match.
	// Strict 搜索是否要求精确匹配.
	Strict *bool
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdSearchParams) Values() (url.Values, map[string]string) {
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
	for _, v := range p.Categories {
		if query == nil {
			query = url.Values{}
		}
		query.Add("categories", v)
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
	if p.Search != nil && *p.Search != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("search", *p.Search)
	}
	if p.Strict != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("strict", strconv.FormatBool(*p.Strict))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}
