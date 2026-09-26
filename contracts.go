package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterContracts Get contracts.
// GetCharacterContracts 获取合同.
//
// Route: GET /characters/{character_id}/contracts/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contracts/ — 该路由缓存长达 300 秒
// Scopes: esi-contracts.read_character_contracts.v1
// 权限: esi-contracts.read_character_contracts.v1
func (c *Client) GetCharacterContracts(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.Contract, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.Contract
	err := c.get(ctx, "/characters/{character_id}/contracts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterContractBids Get contract bids.
// GetCharacterContractBids 获取合同出价.
//
// Route: GET /characters/{character_id}/contracts/{contract_id}/bids/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contracts/{contract_id}/bids/ — 该路由缓存长达 300 秒
// Scopes: esi-contracts.read_character_contracts.v1
// 权限: esi-contracts.read_character_contracts.v1
func (c *Client) GetCharacterContractBids(ctx context.Context, token string, characterID int32, contractID int32, ifNoneMatch ...string) ([]models.ContractBid, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "contract_id": strconv.FormatInt(int64(contractID), 10)}
	var result []models.ContractBid
	err := c.get(ctx, "/characters/{character_id}/contracts/{contract_id}/bids/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterContractItems Get contract items.
// GetCharacterContractItems 获取合同物品.
//
// Route: GET /characters/{character_id}/contracts/{contract_id}/items/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/contracts/{contract_id}/items/ — 该路由缓存长达 3600 秒
// Scopes: esi-contracts.read_character_contracts.v1
// 权限: esi-contracts.read_character_contracts.v1
func (c *Client) GetCharacterContractItems(ctx context.Context, token string, characterID int32, contractID int32, ifNoneMatch ...string) ([]models.ContractItem, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "contract_id": strconv.FormatInt(int64(contractID), 10)}
	var result []models.ContractItem
	err := c.get(ctx, "/characters/{character_id}/contracts/{contract_id}/items/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationContracts Get corporation contracts.
// GetCorporationContracts 获取军团合同.
//
// Route: GET /corporations/{corporation_id}/contracts/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/contracts/ — 该路由缓存长达 300 秒
// Scopes: esi-contracts.read_corporation_contracts.v1
// 权限: esi-contracts.read_corporation_contracts.v1
func (c *Client) GetCorporationContracts(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.Contract, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.Contract
	err := c.get(ctx, "/corporations/{corporation_id}/contracts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationContractBids Get corporation contract bids.
// GetCorporationContractBids 获取军团合同出价.
//
// Route: GET /corporations/{corporation_id}/contracts/{contract_id}/bids/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/contracts/{contract_id}/bids/ — 该路由缓存长达 3600 秒
// Scopes: esi-contracts.read_corporation_contracts.v1
// 权限: esi-contracts.read_corporation_contracts.v1
func (c *Client) GetCorporationContractBids(ctx context.Context, token string, contractID int32, corporationID int32, page int32, ifNoneMatch ...string) ([]models.ContractBid, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractID), 10), "corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.ContractBid
	err := c.get(ctx, "/corporations/{corporation_id}/contracts/{contract_id}/bids/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationContractItems Get corporation contract items.
// GetCorporationContractItems 获取军团合同物品.
//
// Route: GET /corporations/{corporation_id}/contracts/{contract_id}/items/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/contracts/{contract_id}/items/ — 该路由缓存长达 3600 秒
// Scopes: esi-contracts.read_corporation_contracts.v1
// 权限: esi-contracts.read_corporation_contracts.v1
func (c *Client) GetCorporationContractItems(ctx context.Context, token string, contractID int32, corporationID int32, ifNoneMatch ...string) ([]models.ContractItem, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractID), 10), "corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.ContractItem
	err := c.get(ctx, "/corporations/{corporation_id}/contracts/{contract_id}/items/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicContractBids Get public contract bids.
// GetPublicContractBids 获取公开合同竞标.
//
// Route: GET /contracts/public/bids/{contract_id}/ — This route is cached for up to 300 seconds
// 路由: GET /contracts/public/bids/{contract_id}/ — 该路由缓存长达 300 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPublicContractBids(ctx context.Context, contractID int32, page int32, ifNoneMatch ...string) ([]models.PublicContractBid, error) {
	query := url.Values{}
	headers := map[string]string{}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractID), 10)}
	var result []models.PublicContractBid
	err := c.get(ctx, "/contracts/public/bids/{contract_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicContractItems Get public contract items.
// GetPublicContractItems 获取公开合同物品.
//
// Route: GET /contracts/public/items/{contract_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /contracts/public/items/{contract_id}/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPublicContractItems(ctx context.Context, contractID int32, page int32, ifNoneMatch ...string) ([]models.PublicContractItem, error) {
	query := url.Values{}
	headers := map[string]string{}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"contract_id": strconv.FormatInt(int64(contractID), 10)}
	var result []models.PublicContractItem
	err := c.get(ctx, "/contracts/public/items/{contract_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPublicContracts Get public contracts.
// GetPublicContracts 获取公开合同.
//
// Route: GET /contracts/public/{region_id}/ — This route is cached for up to 1800 seconds
// 路由: GET /contracts/public/{region_id}/ — 该路由缓存长达 1800 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetPublicContracts(ctx context.Context, regionID int32, page int32, ifNoneMatch ...string) ([]models.PublicContract, error) {
	query := url.Values{}
	headers := map[string]string{}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"region_id": strconv.FormatInt(int64(regionID), 10)}
	var result []models.PublicContract
	err := c.get(ctx, "/contracts/public/{region_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
