package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// DeleteCharactersCharacterIdContacts Delete contacts.
// DeleteCharactersCharacterIdContacts 删除联系人.
//
// Route: DELETE /characters/{character_id}/contacts/
// 路由: DELETE /characters/{character_id}/contacts/
// Scopes: esi-characters.write_contacts.v1
// 权限: esi-characters.write_contacts.v1
func (c *Client) DeleteCharactersCharacterIdContacts(ctx context.Context, characterId int32, params *models.DeleteCharactersCharacterIdContactsParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	err := c.delete(ctx, "/characters/{character_id}/contacts/", pathParams, query, headers)
	return err
}

// GetAlliancesAllianceIdContacts Get alliance contacts.
// GetAlliancesAllianceIdContacts 获取联盟联系人.
//
// Route: GET /alliances/{alliance_id}/contacts/ — This route is cached for up to 300 seconds
// 路由: GET /alliances/{alliance_id}/contacts/ — 该路由缓存长达 300 秒
// Scopes: esi-alliances.read_contacts.v1
// 权限: esi-alliances.read_contacts.v1
func (c *Client) GetAlliancesAllianceIdContacts(ctx context.Context, allianceId int32, params *models.GetAlliancesAllianceIdContactsParams) ([]models.GetAlliancesAllianceIdContacts, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceId), 10)}
	var result []models.GetAlliancesAllianceIdContacts
	err := c.get(ctx, "/alliances/{alliance_id}/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAlliancesAllianceIdContactsLabels Get alliance contact labels.
// GetAlliancesAllianceIdContactsLabels 获取联盟联系人标签.
//
// Route: GET /alliances/{alliance_id}/contacts/labels/ — This route is cached for up to 300 seconds
// 路由: GET /alliances/{alliance_id}/contacts/labels/ — 该路由缓存长达 300 秒
// Scopes: esi-alliances.read_contacts.v1
// 权限: esi-alliances.read_contacts.v1
func (c *Client) GetAlliancesAllianceIdContactsLabels(ctx context.Context, allianceId int32, params *models.GetAlliancesAllianceIdContactsLabelsParams) ([]models.GetAlliancesAllianceIdContactsLabels, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceId), 10)}
	var result []models.GetAlliancesAllianceIdContactsLabels
	err := c.get(ctx, "/alliances/{alliance_id}/contacts/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdContacts Get contacts.
// GetCharactersCharacterIdContacts 获取联系人.
//
// Route: GET /characters/{character_id}/contacts/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contacts/ — 该路由缓存长达 300 秒
// Scopes: esi-characters.read_contacts.v1
// 权限: esi-characters.read_contacts.v1
func (c *Client) GetCharactersCharacterIdContacts(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdContactsParams) ([]models.GetCharactersCharacterIdContacts, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdContacts
	err := c.get(ctx, "/characters/{character_id}/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdContactsLabels Get contact labels.
// GetCharactersCharacterIdContactsLabels 获取联系人标签.
//
// Route: GET /characters/{character_id}/contacts/labels/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contacts/labels/ — 该路由缓存长达 300 秒
// Scopes: esi-characters.read_contacts.v1
// 权限: esi-characters.read_contacts.v1
func (c *Client) GetCharactersCharacterIdContactsLabels(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdContactsLabelsParams) ([]models.GetCharactersCharacterIdContactsLabels, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdContactsLabels
	err := c.get(ctx, "/characters/{character_id}/contacts/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdContacts Get corporation contacts.
// GetCorporationsCorporationIdContacts 获取军团联系人.
//
// Route: GET /corporations/{corporation_id}/contacts/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/contacts/ — 该路由缓存长达 300 秒
// Scopes: esi-corporations.read_contacts.v1
// 权限: esi-corporations.read_contacts.v1
func (c *Client) GetCorporationsCorporationIdContacts(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdContactsParams) ([]models.GetCorporationsCorporationIdContacts, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdContacts
	err := c.get(ctx, "/corporations/{corporation_id}/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdContactsLabels Get corporation contact labels.
// GetCorporationsCorporationIdContactsLabels 获取军团联系人标签.
//
// Route: GET /corporations/{corporation_id}/contacts/labels/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/contacts/labels/ — 该路由缓存长达 300 秒
// Scopes: esi-corporations.read_contacts.v1
// 权限: esi-corporations.read_contacts.v1
func (c *Client) GetCorporationsCorporationIdContactsLabels(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdContactsLabelsParams) ([]models.GetCorporationsCorporationIdContactsLabels, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdContactsLabels
	err := c.get(ctx, "/corporations/{corporation_id}/contacts/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCharactersCharacterIdContacts Add contacts.
// PostCharactersCharacterIdContacts 添加联系人.
//
// Route: POST /characters/{character_id}/contacts/
// 路由: POST /characters/{character_id}/contacts/
// Scopes: esi-characters.write_contacts.v1
// 权限: esi-characters.write_contacts.v1
func (c *Client) PostCharactersCharacterIdContacts(ctx context.Context, characterId int32, body []int32, params *models.PostCharactersCharacterIdContactsParams) ([]int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	if body == nil {
		return nil, errBodyRequired
	}
	var result []int32
	err := c.post(ctx, "/characters/{character_id}/contacts/", pathParams, query, headers, body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutCharactersCharacterIdContacts Edit contacts.
// PutCharactersCharacterIdContacts 编辑联系人.
//
// Route: PUT /characters/{character_id}/contacts/
// 路由: PUT /characters/{character_id}/contacts/
// Scopes: esi-characters.write_contacts.v1
// 权限: esi-characters.write_contacts.v1
func (c *Client) PutCharactersCharacterIdContacts(ctx context.Context, characterId int32, body []int32, params *models.PutCharactersCharacterIdContactsParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/characters/{character_id}/contacts/", pathParams, query, headers, body)
	return err
}
