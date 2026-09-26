package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterMarketOrders List open orders from a character.
// GetCharacterMarketOrders 列出角色的公开订单.
//
// Route: GET /characters/{character_id}/orders/ — This route is cached for up to 1200 seconds
// 路由: GET /characters/{character_id}/orders/ — 该路由缓存长达 1200 秒
// Scopes: esi-markets.read_character_orders.v1
// 权限: esi-markets.read_character_orders.v1
func (c *Client) GetCharacterMarketOrders(ctx context.Context, characterID int32, token string, ifNoneMatch string) ([]models.CharacterMarketOrder, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterMarketOrder
	err := c.get(ctx, "/characters/{character_id}/orders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterMarketOrderHistory List historical orders by a character.
// GetCharacterMarketOrderHistory 列出角色的历史订单.
//
// Route: GET /characters/{character_id}/orders/history/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/orders/history/ — 该路由缓存长达 3600 秒
// Scopes: esi-markets.read_character_orders.v1
// 权限: esi-markets.read_character_orders.v1
func (c *Client) GetCharacterMarketOrderHistory(ctx context.Context, characterID int32, token string, page int32, ifNoneMatch string) ([]models.CharacterMarketOrderHistory, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterMarketOrderHistory
	err := c.get(ctx, "/characters/{character_id}/orders/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMarketOrders List open orders from a corporation.
// GetCorporationMarketOrders 列出军团的公开订单.
//
// Route: GET /corporations/{corporation_id}/orders/ — This route is cached for up to 1200 seconds
// 路由: GET /corporations/{corporation_id}/orders/ — 该路由缓存长达 1200 秒
// Scopes: esi-markets.read_corporation_orders.v1
// 权限: esi-markets.read_corporation_orders.v1
func (c *Client) GetCorporationMarketOrders(ctx context.Context, corporationID int32, token string, page int32, ifNoneMatch string) ([]models.CorporationMarketOrder, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationMarketOrder
	err := c.get(ctx, "/corporations/{corporation_id}/orders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMarketOrderHistory List historical orders from a corporation.
// GetCorporationMarketOrderHistory 列出军团的历史订单.
//
// Route: GET /corporations/{corporation_id}/orders/history/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/orders/history/ — 该路由缓存长达 3600 秒
// Scopes: esi-markets.read_corporation_orders.v1
// 权限: esi-markets.read_corporation_orders.v1
func (c *Client) GetCorporationMarketOrderHistory(ctx context.Context, corporationID int32, token string, page int32, ifNoneMatch string) ([]models.CorporationMarketOrderHistory, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationMarketOrderHistory
	err := c.get(ctx, "/corporations/{corporation_id}/orders/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarketGroups Get item groups.
// GetMarketGroups 获取物品分组.
//
// Route: GET /markets/groups/
// 路由: GET /markets/groups/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMarketGroups(ctx context.Context, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/markets/groups/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarketGroup Get item group information.
// GetMarketGroup 获取物品分组信息.
//
// Route: GET /markets/groups/{market_group_id}/
// 路由: GET /markets/groups/{market_group_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMarketGroup(ctx context.Context, marketGroupID int32, ifNoneMatch string) (*models.MarketGroup, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"market_group_id": strconv.FormatInt(int64(marketGroupID), 10)}
	var result *models.MarketGroup
	err := c.get(ctx, "/markets/groups/{market_group_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarketPrices List market prices.
// GetMarketPrices 列出市场价格.
//
// Route: GET /markets/prices/ — This route is cached for up to 3600 seconds
// 路由: GET /markets/prices/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMarketPrices(ctx context.Context, ifNoneMatch string) ([]models.MarketPrice, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.MarketPrice
	err := c.get(ctx, "/markets/prices/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarketHistory List historical market statistics in a region.
// GetMarketHistory 列出星域内的历史市场统计.
//
// Route: GET /markets/{region_id}/history/
// 路由: GET /markets/{region_id}/history/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMarketHistory(ctx context.Context, regionID int32, typeID string, ifNoneMatch string) ([]models.MarketHistoryEntry, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	query.Set("type_id", typeID)
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionID), 10)}
	var result []models.MarketHistoryEntry
	err := c.get(ctx, "/markets/{region_id}/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarketOrders List orders in a region.
// GetMarketOrders 列出某星域的订单.
//
// Route: GET /markets/{region_id}/orders/ — This route is cached for up to 300 seconds
// 路由: GET /markets/{region_id}/orders/ — 该路由缓存长达 300 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMarketOrders(ctx context.Context, regionID int32, orderType string, page int32, typeID int32, ifNoneMatch string) ([]models.MarketOrder, error) {
	query := url.Values{}
	headers := map[string]string{}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if typeID != 0 {
		query.Set("type_id", strconv.FormatInt(int64(typeID), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	query.Set("order_type", orderType)
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionID), 10)}
	var result []models.MarketOrder
	err := c.get(ctx, "/markets/{region_id}/orders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMarketTypes List type IDs relevant to a market.
// GetMarketTypes 列出与市场相关的类型 ID.
//
// Route: GET /markets/{region_id}/types/ — This route is cached for up to 600 seconds
// 路由: GET /markets/{region_id}/types/ — 该路由缓存长达 600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetMarketTypes(ctx context.Context, regionID int32, page int32, ifNoneMatch string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionID), 10)}
	var result []int32
	err := c.get(ctx, "/markets/{region_id}/types/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStructureMarketOrders List orders in a structure.
// GetStructureMarketOrders 列出某建筑的订单.
//
// Route: GET /markets/structures/{structure_id}/ — This route is cached for up to 300 seconds
// 路由: GET /markets/structures/{structure_id}/ — 该路由缓存长达 300 秒
// Scopes: esi-markets.structure_markets.v1
// 权限: esi-markets.structure_markets.v1
func (c *Client) GetStructureMarketOrders(ctx context.Context, structureID int64, token string, page int32, ifNoneMatch string) ([]models.StructureMarketOrder, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"structure_id": strconv.FormatInt(structureID, 10)}
	var result []models.StructureMarketOrder
	err := c.get(ctx, "/markets/structures/{structure_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
