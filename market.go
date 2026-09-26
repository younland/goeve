package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdOrders List open orders from a character.
// GetCharactersCharacterIdOrders 列出角色的公开订单.
//
// Route: GET /characters/{character_id}/orders/ — This route is cached for up to 1200 seconds
// 路由: GET /characters/{character_id}/orders/ — 该路由缓存长达 1200 秒
// Scopes: esi-markets.read_character_orders.v1
// 权限: esi-markets.read_character_orders.v1
func (c *Client) GetCharactersCharacterIdOrders(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOrdersParams) ([]models.GetCharactersCharacterIdOrders, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdOrders
	err := c.get(ctx, "/characters/{character_id}/orders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdOrdersHistory List historical orders by a character.
// GetCharactersCharacterIdOrdersHistory 列出角色的历史订单.
//
// Route: GET /characters/{character_id}/orders/history/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/orders/history/ — 该路由缓存长达 3600 秒
// Scopes: esi-markets.read_character_orders.v1
// 权限: esi-markets.read_character_orders.v1
func (c *Client) GetCharactersCharacterIdOrdersHistory(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOrdersHistoryParams) ([]models.GetCharactersCharacterIdOrdersHistory, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdOrdersHistory
	err := c.get(ctx, "/characters/{character_id}/orders/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdOrders List open orders from a corporation.
// GetCorporationsCorporationIdOrders 列出军团的公开订单.
//
// Route: GET /corporations/{corporation_id}/orders/ — This route is cached for up to 1200 seconds
// 路由: GET /corporations/{corporation_id}/orders/ — 该路由缓存长达 1200 秒
// Scopes: esi-markets.read_corporation_orders.v1
// 权限: esi-markets.read_corporation_orders.v1
func (c *Client) GetCorporationsCorporationIdOrders(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdOrdersParams) ([]models.GetCorporationsCorporationIdOrders, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdOrders
	err := c.get(ctx, "/corporations/{corporation_id}/orders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdOrdersHistory List historical orders from a corporation.
// GetCorporationsCorporationIdOrdersHistory 列出军团的历史订单.
//
// Route: GET /corporations/{corporation_id}/orders/history/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/orders/history/ — 该路由缓存长达 3600 秒
// Scopes: esi-markets.read_corporation_orders.v1
// 权限: esi-markets.read_corporation_orders.v1
func (c *Client) GetCorporationsCorporationIdOrdersHistory(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdOrdersHistoryParams) ([]models.GetCorporationsCorporationIdOrdersHistory, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdOrdersHistory
	err := c.get(ctx, "/corporations/{corporation_id}/orders/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetsGroups Get item groups.
// GetsGroups 获取物品分组.
//
// Route: GET /markets/groups/
// 路由: GET /markets/groups/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetsGroups(ctx context.Context, params *models.GetsGroupsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/markets/groups/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetsGroupsMarketGroupId Get item group information.
// GetsGroupsMarketGroupId 获取物品分组信息.
//
// Route: GET /markets/groups/{market_group_id}/
// 路由: GET /markets/groups/{market_group_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetsGroupsMarketGroupId(ctx context.Context, marketGroupId int32, params *models.GetsGroupsMarketGroupIdParams) (*models.GetMarketsGroupsMarketGroupId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"market_group_id": strconv.FormatInt(int64(marketGroupId), 10)}
	var result *models.GetMarketsGroupsMarketGroupId
	err := c.get(ctx, "/markets/groups/{market_group_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetsPrices List market prices.
// GetsPrices 列出市场价格.
//
// Route: GET /markets/prices/ — This route is cached for up to 3600 seconds
// 路由: GET /markets/prices/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetsPrices(ctx context.Context, params *models.GetsPricesParams) ([]models.GetMarketsPrices, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetMarketsPrices
	err := c.get(ctx, "/markets/prices/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetsRegionIdHistory List historical market statistics in a region.
// GetsRegionIdHistory 列出星域内的历史市场统计.
//
// Route: GET /markets/{region_id}/history/
// 路由: GET /markets/{region_id}/history/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetsRegionIdHistory(ctx context.Context, regionId int32, params *models.GetsRegionIdHistoryParams) ([]models.GetMarketsRegionIdHistory, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionId), 10)}
	var result []models.GetMarketsRegionIdHistory
	err := c.get(ctx, "/markets/{region_id}/history/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetsRegionIdOrders List orders in a region.
// GetsRegionIdOrders 列出某星域的订单.
//
// Route: GET /markets/{region_id}/orders/ — This route is cached for up to 300 seconds
// 路由: GET /markets/{region_id}/orders/ — 该路由缓存长达 300 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetsRegionIdOrders(ctx context.Context, regionId int32, params *models.GetsRegionIdOrdersParams) ([]models.GetMarketsRegionIdOrders, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionId), 10)}
	var result []models.GetMarketsRegionIdOrders
	err := c.get(ctx, "/markets/{region_id}/orders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetsRegionIdTypes List type IDs relevant to a market.
// GetsRegionIdTypes 列出与市场相关的类型 ID.
//
// Route: GET /markets/{region_id}/types/ — This route is cached for up to 600 seconds
// 路由: GET /markets/{region_id}/types/ — 该路由缓存长达 600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetsRegionIdTypes(ctx context.Context, regionId int32, params *models.GetsRegionIdTypesParams) ([]int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionId), 10)}
	var result []int32
	err := c.get(ctx, "/markets/{region_id}/types/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetsStructuresStructureId List orders in a structure.
// GetsStructuresStructureId 列出某建筑的订单.
//
// Route: GET /markets/structures/{structure_id}/ — This route is cached for up to 300 seconds
// 路由: GET /markets/structures/{structure_id}/ — 该路由缓存长达 300 秒
// Scopes: esi-markets.structure_markets.v1
// 权限: esi-markets.structure_markets.v1
func (c *Client) GetsStructuresStructureId(ctx context.Context, structureId int64, params *models.GetsStructuresStructureIdParams) ([]models.GetMarketsStructuresStructureId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"structure_id": strconv.FormatInt(int64(structureId), 10)}
	var result []models.GetMarketsStructuresStructureId
	err := c.get(ctx, "/markets/structures/{structure_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
