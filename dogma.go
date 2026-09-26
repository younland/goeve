package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetAttributes Get attributes.
// GetAttributes 获取属性.
//
// Route: GET /dogma/attributes/
// 路由: GET /dogma/attributes/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAttributes(ctx context.Context, params *models.GetAttributesParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/dogma/attributes/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAttributesAttributeId Get attribute information.
// GetAttributesAttributeId 获取属性信息.
//
// Route: GET /dogma/attributes/{attribute_id}/
// 路由: GET /dogma/attributes/{attribute_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetAttributesAttributeId(ctx context.Context, attributeId int32, params *models.GetAttributesAttributeIdParams) (*models.GetDogmaAttributesAttributeId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"attribute_id": strconv.FormatInt(int64(attributeId), 10)}
	var result *models.GetDogmaAttributesAttributeId
	err := c.get(ctx, "/dogma/attributes/{attribute_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDynamicItemsTypeIdItemId Get dynamic item information.
// GetDynamicItemsTypeIdItemId 获取动态物品信息.
//
// Route: GET /dogma/dynamic/items/{type_id}/{item_id}/
// 路由: GET /dogma/dynamic/items/{type_id}/{item_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetDynamicItemsTypeIdItemId(ctx context.Context, itemId int64, typeId int32, params *models.GetDynamicItemsTypeIdItemIdParams) (*models.GetDogmaDynamicItemsTypeIdItemId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"item_id": strconv.FormatInt(int64(itemId), 10), "type_id": strconv.FormatInt(int64(typeId), 10)}
	var result *models.GetDogmaDynamicItemsTypeIdItemId
	err := c.get(ctx, "/dogma/dynamic/items/{type_id}/{item_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEffects Get effects.
// GetEffects 获取效果.
//
// Route: GET /dogma/effects/
// 路由: GET /dogma/effects/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetEffects(ctx context.Context, params *models.GetEffectsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/dogma/effects/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetEffectsEffectId Get effect information.
// GetEffectsEffectId 获取效果信息.
//
// Route: GET /dogma/effects/{effect_id}/
// 路由: GET /dogma/effects/{effect_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetEffectsEffectId(ctx context.Context, effectId int32, params *models.GetEffectsEffectIdParams) (*models.GetDogmaEffectsEffectId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"effect_id": strconv.FormatInt(int64(effectId), 10)}
	var result *models.GetDogmaEffectsEffectId
	err := c.get(ctx, "/dogma/effects/{effect_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
