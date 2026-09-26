package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharacterWalletBalance Get a character's wallet balance.
// GetCharacterWalletBalance 获取角色的钱包余额.
//
// Route: GET /characters/{character_id}/wallet/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/wallet/ — 该路由缓存长达 120 秒
// Scopes: esi-wallet.read_character_wallet.v1
// 权限: esi-wallet.read_character_wallet.v1
func (c *Client) GetCharacterWalletBalance(ctx context.Context, characterID int32, opts ...RequestOption) (float64, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result float64
	err := c.get(ctx, "/characters/{character_id}/wallet/", pathParams, query, headers, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// GetCharacterWalletJournal Get character wallet journal.
// GetCharacterWalletJournal 获取角色钱包日志.
//
// Route: GET /characters/{character_id}/wallet/journal/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/wallet/journal/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_character_wallet.v1
// 权限: esi-wallet.read_character_wallet.v1
func (c *Client) GetCharacterWalletJournal(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.WalletJournalEntry, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.WalletJournalEntry
	err := c.get(ctx, "/characters/{character_id}/wallet/journal/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWalletTransactions Get wallet transactions.
// GetWalletTransactions 获取钱包交易记录.
//
// Route: GET /characters/{character_id}/wallet/transactions/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/wallet/transactions/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_character_wallet.v1
// 权限: esi-wallet.read_character_wallet.v1
func (c *Client) GetWalletTransactions(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterWalletTransaction, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterWalletTransaction
	err := c.get(ctx, "/characters/{character_id}/wallet/transactions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationWallets Returns a corporation's wallet balance.
// GetCorporationWallets 返回军团的钱包余额.
//
// Route: GET /corporations/{corporation_id}/wallets/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/wallets/ — 该路由缓存长达 300 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationWallets(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationWallet, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationWallet
	err := c.get(ctx, "/corporations/{corporation_id}/wallets/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationWalletJournal Get corporation wallet journal.
// GetCorporationWalletJournal 获取军团钱包日志.
//
// Route: GET /corporations/{corporation_id}/wallets/{division}/journal/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/wallets/{division}/journal/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationWalletJournal(ctx context.Context, corporationID int32, division int32, opts ...RequestOption) ([]models.WalletJournalEntry, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10), "division": strconv.FormatInt(int64(division), 10)}
	var result []models.WalletJournalEntry
	err := c.get(ctx, "/corporations/{corporation_id}/wallets/{division}/journal/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationWalletTransactions Get corporation wallet transactions.
// GetCorporationWalletTransactions 获取军团钱包交易.
//
// Route: GET /corporations/{corporation_id}/wallets/{division}/transactions/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/wallets/{division}/transactions/ — 该路由缓存长达 3600 秒
// Scopes: esi-wallet.read_corporation_wallets.v1
// 权限: esi-wallet.read_corporation_wallets.v1
func (c *Client) GetCorporationWalletTransactions(ctx context.Context, corporationID int32, division int32, opts ...RequestOption) ([]models.CorporationWalletTransaction, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10), "division": strconv.FormatInt(int64(division), 10)}
	var result []models.CorporationWalletTransaction
	err := c.get(ctx, "/corporations/{corporation_id}/wallets/{division}/transactions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
