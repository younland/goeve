package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharacterLoyaltyPoints Get loyalty points.
// GetCharacterLoyaltyPoints 获取忠诚点.
//
// Route: GET /characters/{character_id}/loyalty/points/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/loyalty/points/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_loyalty.v1
// 权限: esi-characters.read_loyalty.v1
func (c *Client) GetCharacterLoyaltyPoints(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.LoyaltyPoints, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.LoyaltyPoints
	err := c.get(ctx, "/characters/{character_id}/loyalty/points/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLoyaltyStoreOffers List loyalty store offers.
// GetLoyaltyStoreOffers 列出忠诚点商店的商品.
//
// Route: GET /loyalty/stores/{corporation_id}/offers/
// 路由: GET /loyalty/stores/{corporation_id}/offers/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetLoyaltyStoreOffers(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.LoyaltyStoreOffer, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.LoyaltyStoreOffer
	err := c.get(ctx, "/loyalty/stores/{corporation_id}/offers/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
