package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdWallet Get a character's wallet balance.
// GetCharactersCharacterIdWallet 获取角色的钱包余额.
//
// Route: GET /characters/{character_id}/wallet/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/wallet/ — 该路由缓存长达 120 秒
// Scopes: esi-wallet.read_character_wallet.v1
// 权限: esi-wallet.read_character_wallet.v1
func (c *Client) GetCharactersCharacterIdWallet(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdWalletParams) (float64, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result float64
	err := c.get(ctx, "/characters/{character_id}/wallet/", pathParams, query, headers, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GetCharactersCharacterIdWalletJournal Get character wallet journal.
// GetCharactersCharacterIdWalletJournal 获取角色钱包日志.
//
// Route: GET /characters/{character_id}/wallet/journal/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/wallet/journal/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_character_wallet.v1
// 权限: esi-wallet.read_character_wallet.v1
func (c *Client) GetCharactersCharacterIdWalletJournal(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdWalletJournalParams) ([]models.GetCharactersCharacterIdWalletJournal, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdWalletJournal
	err := c.get(ctx, "/characters/{character_id}/wallet/journal/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdWalletTransactions Get wallet transactions.
// GetCharactersCharacterIdWalletTransactions 获取钱包交易记录.
//
// Route: GET /characters/{character_id}/wallet/transactions/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/wallet/transactions/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_character_wallet.v1
// 权限: esi-wallet.read_character_wallet.v1
func (c *Client) GetCharactersCharacterIdWalletTransactions(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdWalletTransactionsParams) ([]models.GetCharactersCharacterIdWalletTransactions, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdWalletTransactions
	err := c.get(ctx, "/characters/{character_id}/wallet/transactions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdWallets Returns a corporation's wallet balance.
// GetCorporationsCorporationIdWallets 返回军团的钱包余额.
//
// Route: GET /corporations/{corporation_id}/wallets/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/wallets/ — 该路由缓存长达 300 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationsCorporationIdWallets(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdWalletsParams) ([]models.GetCorporationsCorporationIdWallets, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdWallets
	err := c.get(ctx, "/corporations/{corporation_id}/wallets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdWalletsDivisionJournal Get corporation wallet journal.
// GetCorporationsCorporationIdWalletsDivisionJournal 获取军团钱包日志.
//
// Route: GET /corporations/{corporation_id}/wallets/{division}/journal/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/wallets/{division}/journal/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationsCorporationIdWalletsDivisionJournal(ctx context.Context, corporationId int32, division int32, params *models.GetCorporationsCorporationIdWalletsDivisionJournalParams) ([]models.GetCorporationsCorporationIdWalletsDivisionJournal, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10), "division": strconv.FormatInt(int64(division), 10)}
	var result []models.GetCorporationsCorporationIdWalletsDivisionJournal
	err := c.get(ctx, "/corporations/{corporation_id}/wallets/{division}/journal/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdWalletsDivisionTransactions Get corporation wallet transactions.
// GetCorporationsCorporationIdWalletsDivisionTransactions 获取军团钱包交易.
//
// Route: GET /corporations/{corporation_id}/wallets/{division}/transactions/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/wallets/{division}/transactions/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationsCorporationIdWalletsDivisionTransactions(ctx context.Context, corporationId int32, division int32, params *models.GetCorporationsCorporationIdWalletsDivisionTransactionsParams) ([]models.GetCorporationsCorporationIdWalletsDivisionTransactions, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10), "division": strconv.FormatInt(int64(division), 10)}
	var result []models.GetCorporationsCorporationIdWalletsDivisionTransactions
	err := c.get(ctx, "/corporations/{corporation_id}/wallets/{division}/transactions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
