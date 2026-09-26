package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdContracts Get contracts.
// GetCharactersCharacterIdContracts 获取合同.
//
// Route: GET /characters/{character_id}/contracts/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contracts/ — 该路由缓存长达 300 秒
// Scopes: esi-contracts.read_character_contracts.v1
// 权限: esi-contracts.read_character_contracts.v1
func (c *Client) GetCharactersCharacterIdContracts(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdContractsParams) ([]models.GetCharactersCharacterIdContracts, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdContracts
	err := c.get(ctx, "/characters/{character_id}/contracts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdContractsContractIdBids Get contract bids.
// GetCharactersCharacterIdContractsContractIdBids 获取合同出价.
//
// Route: GET /characters/{character_id}/contracts/{contract_id}/bids/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contracts/{contract_id}/bids/ — 该路由缓存长达 300 秒
// Scopes: esi-contracts.read_character_contracts.v1
// 权限: esi-contracts.read_character_contracts.v1
func (c *Client) GetCharactersCharacterIdContractsContractIdBids(ctx context.Context, characterId int32, contractId int32, params *models.GetCharactersCharacterIdContractsContractIdBidsParams) ([]models.GetCharactersCharacterIdContractsContractIdBids, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "contract_id": strconv.FormatInt(int64(contractId), 10)}
	var result []models.GetCharactersCharacterIdContractsContractIdBids
	err := c.get(ctx, "/characters/{character_id}/contracts/{contract_id}/bids/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdContractsContractIdItems Get contract items.
// GetCharactersCharacterIdContractsContractIdItems 获取合同物品.
//
// Route: GET /characters/{character_id}/contracts/{contract_id}/items/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/contracts/{contract_id}/items/ — 该路由缓存长达 3600 秒
// Scopes: esi-contracts.read_character_contracts.v1
// 权限: esi-contracts.read_character_contracts.v1
func (c *Client) GetCharactersCharacterIdContractsContractIdItems(ctx context.Context, characterId int32, contractId int32, params *models.GetCharactersCharacterIdContractsContractIdItemsParams) ([]models.GetCharactersCharacterIdContractsContractIdItems, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "contract_id": strconv.FormatInt(int64(contractId), 10)}
	var result []models.GetCharactersCharacterIdContractsContractIdItems
	err := c.get(ctx, "/characters/{character_id}/contracts/{contract_id}/items/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdContracts Get corporation contracts.
// GetCorporationsCorporationIdContracts 获取军团合同.
//
// Route: GET /corporations/{corporation_id}/contracts/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/contracts/ — 该路由缓存长达 300 秒
// Scopes: esi-contracts.read_corporation_contracts.v1
// 权限: esi-contracts.read_corporation_contracts.v1
func (c *Client) GetCorporationsCorporationIdContracts(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdContractsParams) ([]models.GetCorporationsCorporationIdContracts, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdContracts
	err := c.get(ctx, "/corporations/{corporation_id}/contracts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdContractsContractIdBids Get corporation contract bids.
// GetCorporationsCorporationIdContractsContractIdBids 获取军团合同出价.
//
// Route: GET /corporations/{corporation_id}/contracts/{contract_id}/bids/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/contracts/{contract_id}/bids/ — 该路由缓存长达 3600 秒
// Scopes: esi-contracts.read_corporation_contracts.v1
// 权限: esi-contracts.read_corporation_contracts.v1
func (c *Client) GetCorporationsCorporationIdContractsContractIdBids(ctx context.Context, contractId int32, corporationId int32, params *models.GetCorporationsCorporationIdContractsContractIdBidsParams) ([]models.GetCorporationsCorporationIdContractsContractIdBids, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractId), 10), "corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdContractsContractIdBids
	err := c.get(ctx, "/corporations/{corporation_id}/contracts/{contract_id}/bids/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdContractsContractIdItems Get corporation contract items.
// GetCorporationsCorporationIdContractsContractIdItems 获取军团合同物品.
//
// Route: GET /corporations/{corporation_id}/contracts/{contract_id}/items/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/contracts/{contract_id}/items/ — 该路由缓存长达 3600 秒
// Scopes: esi-contracts.read_corporation_contracts.v1
// 权限: esi-contracts.read_corporation_contracts.v1
func (c *Client) GetCorporationsCorporationIdContractsContractIdItems(ctx context.Context, contractId int32, corporationId int32, params *models.GetCorporationsCorporationIdContractsContractIdItemsParams) ([]models.GetCorporationsCorporationIdContractsContractIdItems, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractId), 10), "corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdContractsContractIdItems
	err := c.get(ctx, "/corporations/{corporation_id}/contracts/{contract_id}/items/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicBidsContractId Get public contract bids.
// GetPublicBidsContractId 获取公开合同竞标.
//
// Route: GET /contracts/public/bids/{contract_id}/ — This route is cached for up to 300 seconds
// 路由: GET /contracts/public/bids/{contract_id}/ — 该路由缓存长达 300 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPublicBidsContractId(ctx context.Context, contractId int32, params *models.GetPublicBidsContractIdParams) ([]models.GetContractsPublicBidsContractId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractId), 10)}
	var result []models.GetContractsPublicBidsContractId
	err := c.get(ctx, "/contracts/public/bids/{contract_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicItemsContractId Get public contract items.
// GetPublicItemsContractId 获取公开合同物品.
//
// Route: GET /contracts/public/items/{contract_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /contracts/public/items/{contract_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPublicItemsContractId(ctx context.Context, contractId int32, params *models.GetPublicItemsContractIdParams) ([]models.GetContractsPublicItemsContractId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractId), 10)}
	var result []models.GetContractsPublicItemsContractId
	err := c.get(ctx, "/contracts/public/items/{contract_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicRegionId Get public contracts.
// GetPublicRegionId 获取公开合同.
//
// Route: GET /contracts/public/{region_id}/ — This route is cached for up to 1800 seconds
// 路由: GET /contracts/public/{region_id}/ — 该路由缓存长达 1800 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPublicRegionId(ctx context.Context, regionId int32, params *models.GetPublicRegionIdParams) ([]models.GetContractsPublicRegionId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionId), 10)}
	var result []models.GetContractsPublicRegionId
	err := c.get(ctx, "/contracts/public/{region_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
