package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetDogmaAttributes Get attributes.
// GetDogmaAttributes 获取属性.
//
// Route: GET /dogma/attributes/
// 路由: GET /dogma/attributes/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetDogmaAttributes(ctx context.Context, ifNoneMatch ...string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/dogma/attributes/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDogmaAttribute Get attribute information.
// GetDogmaAttribute 获取属性信息.
//
// Route: GET /dogma/attributes/{attribute_id}/
// 路由: GET /dogma/attributes/{attribute_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetDogmaAttribute(ctx context.Context, attributeID int32, ifNoneMatch ...string) (*models.DogmaAttribute, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"attribute_id": strconv.FormatInt(int64(attributeID), 10)}
	var result *models.DogmaAttribute
	err := c.get(ctx, "/dogma/attributes/{attribute_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDogmaDynamicItem Get dynamic item information.
// GetDogmaDynamicItem 获取动态物品信息.
//
// Route: GET /dogma/dynamic/items/{type_id}/{item_id}/
// 路由: GET /dogma/dynamic/items/{type_id}/{item_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetDogmaDynamicItem(ctx context.Context, itemID int64, typeID int32, ifNoneMatch ...string) (*models.DogmaDynamicItem, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"item_id": strconv.FormatInt(int64(itemID), 10), "type_id": strconv.FormatInt(int64(typeID), 10)}
	var result *models.DogmaDynamicItem
	err := c.get(ctx, "/dogma/dynamic/items/{type_id}/{item_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDogmaEffects Get effects.
// GetDogmaEffects 获取效果.
//
// Route: GET /dogma/effects/
// 路由: GET /dogma/effects/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetDogmaEffects(ctx context.Context, ifNoneMatch ...string) ([]int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/dogma/effects/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDogmaEffect Get effect information.
// GetDogmaEffect 获取效果信息.
//
// Route: GET /dogma/effects/{effect_id}/
// 路由: GET /dogma/effects/{effect_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetDogmaEffect(ctx context.Context, effectID int32, ifNoneMatch ...string) (*models.DogmaEffect, error) {
	query := url.Values{}
	headers := map[string]string{}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"effect_id": strconv.FormatInt(int64(effectID), 10)}
	var result *models.DogmaEffect
	err := c.get(ctx, "/dogma/effects/{effect_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
