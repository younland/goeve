package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdLoyaltyPoints Get loyalty points.
// GetCharactersCharacterIdLoyaltyPoints 获取忠诚点.
//
// Route: GET /characters/{character_id}/loyalty/points/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/loyalty/points/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_loyalty.v1
// 权限: esi-characters.read_loyalty.v1
func (c *Client) GetCharactersCharacterIdLoyaltyPoints(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdLoyaltyPointsParams) ([]models.GetCharactersCharacterIdLoyaltyPoints, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdLoyaltyPoints
	err := c.get(ctx, "/characters/{character_id}/loyalty/points/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStoresCorporationIdOffers List loyalty store offers.
// GetStoresCorporationIdOffers 列出忠诚点商店的商品.
//
// Route: GET /loyalty/stores/{corporation_id}/offers/
// 路由: GET /loyalty/stores/{corporation_id}/offers/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetStoresCorporationIdOffers(ctx context.Context, corporationId int32, params *models.GetStoresCorporationIdOffersParams) ([]models.GetLoyaltyStoresCorporationIdOffers, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetLoyaltyStoresCorporationIdOffers
	err := c.get(ctx, "/loyalty/stores/{corporation_id}/offers/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
