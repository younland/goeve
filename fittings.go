package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// DeleteCharacterFitting Delete fitting.
// DeleteCharacterFitting 删除装配.
//
// Route: DELETE /characters/{character_id}/fittings/{fitting_id}/
// 路由: DELETE /characters/{character_id}/fittings/{fitting_id}/
// Scopes: esi-fittings.write_fittings.v1
// 权限: esi-fittings.write_fittings.v1
func (c *Client) DeleteCharacterFitting(ctx context.Context, characterID int32, fittingID int32, token string) error {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "fitting_id": strconv.FormatInt(int64(fittingID), 10)}
	err := c.delete(ctx, "/characters/{character_id}/fittings/{fitting_id}/", pathParams, query, headers)
	return err
}

// GetCharacterFittings Get fittings.
// GetCharacterFittings 获取装配.
//
// Route: GET /characters/{character_id}/fittings/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/fittings/ — 该路由缓存长达 300 秒
// Scopes: esi-fittings.read_fittings.v1
// 权限: esi-fittings.read_fittings.v1
func (c *Client) GetCharacterFittings(ctx context.Context, characterID int32, token string, ifNoneMatch string) ([]models.Fitting, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.Fitting
	err := c.get(ctx, "/characters/{character_id}/fittings/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateCharacterFitting Create fitting.
// CreateCharacterFitting 创建装配.
//
// Route: POST /characters/{character_id}/fittings/
// 路由: POST /characters/{character_id}/fittings/
// Scopes: esi-fittings.write_fittings.v1
// 权限: esi-fittings.write_fittings.v1
func (c *Client) CreateCharacterFitting(ctx context.Context, characterID int32, body *models.FittingRequest, token string) (*models.NewFitting, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result *models.NewFitting
	err := c.post(ctx, "/characters/{character_id}/fittings/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
