package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// DeleteCharacterContacts Delete contacts.
// DeleteCharacterContacts 删除联系人.
//
// Route: DELETE /characters/{character_id}/contacts/
// 路由: DELETE /characters/{character_id}/contacts/
// Scopes: esi-characters.write_contacts.v1
// 权限: esi-characters.write_contacts.v1
func (c *Client) DeleteCharacterContacts(ctx context.Context, characterID int32, contactIDs []int32, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	for _, v := range contactIDs {
		query.Add("contact_ids", strconv.Itoa(int(v)))
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	err := c.delete(ctx, "/characters/{character_id}/contacts/", pathParams, query, headers)
	return err
}

// GetAllianceContacts Get alliance contacts.
// GetAllianceContacts 获取联盟联系人.
//
// Route: GET /alliances/{alliance_id}/contacts/ — This route is cached for up to 300 seconds
// 路由: GET /alliances/{alliance_id}/contacts/ — 该路由缓存长达 300 秒
// Scopes: esi-alliances.read_contacts.v1
// 权限: esi-alliances.read_contacts.v1
func (c *Client) GetAllianceContacts(ctx context.Context, allianceID int32, opts ...RequestOption) ([]models.AllianceContact, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceID), 10)}
	var result []models.AllianceContact
	err := c.get(ctx, "/alliances/{alliance_id}/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllianceContactLabels Get alliance contact labels.
// GetAllianceContactLabels 获取联盟联系人标签.
//
// Route: GET /alliances/{alliance_id}/contacts/labels/ — This route is cached for up to 300 seconds
// 路由: GET /alliances/{alliance_id}/contacts/labels/ — 该路由缓存长达 300 秒
// Scopes: esi-alliances.read_contacts.v1
// 权限: esi-alliances.read_contacts.v1
func (c *Client) GetAllianceContactLabels(ctx context.Context, allianceID int32, opts ...RequestOption) ([]models.ContactLabel, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"alliance_id": strconv.FormatInt(int64(allianceID), 10)}
	var result []models.ContactLabel
	err := c.get(ctx, "/alliances/{alliance_id}/contacts/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterContacts Get contacts.
// GetCharacterContacts 获取联系人.
//
// Route: GET /characters/{character_id}/contacts/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contacts/ — 该路由缓存长达 300 秒
// Scopes: esi-characters.read_contacts.v1
// 权限: esi-characters.read_contacts.v1
func (c *Client) GetCharacterContacts(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterContact, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterContact
	err := c.get(ctx, "/characters/{character_id}/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterContactLabels Get contact labels.
// GetCharacterContactLabels 获取联系人标签.
//
// Route: GET /characters/{character_id}/contacts/labels/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/contacts/labels/ — 该路由缓存长达 300 秒
// Scopes: esi-characters.read_contacts.v1
// 权限: esi-characters.read_contacts.v1
func (c *Client) GetCharacterContactLabels(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.ContactLabel, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.ContactLabel
	err := c.get(ctx, "/characters/{character_id}/contacts/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationContacts Get corporation contacts.
// GetCorporationContacts 获取军团联系人.
//
// Route: GET /corporations/{corporation_id}/contacts/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/contacts/ — 该路由缓存长达 300 秒
// Scopes: esi-corporations.read_contacts.v1
// 权限: esi-corporations.read_contacts.v1
func (c *Client) GetCorporationContacts(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationContact, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationContact
	err := c.get(ctx, "/corporations/{corporation_id}/contacts/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationContactLabels Get corporation contact labels.
// GetCorporationContactLabels 获取军团联系人标签.
//
// Route: GET /corporations/{corporation_id}/contacts/labels/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/contacts/labels/ — 该路由缓存长达 300 秒
// Scopes: esi-corporations.read_contacts.v1
// 权限: esi-corporations.read_contacts.v1
func (c *Client) GetCorporationContactLabels(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.ContactLabel, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.ContactLabel
	err := c.get(ctx, "/corporations/{corporation_id}/contacts/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddCharacterContacts Add contacts.
// AddCharacterContacts 添加联系人.
//
// Route: POST /characters/{character_id}/contacts/
// 路由: POST /characters/{character_id}/contacts/
// Scopes: esi-characters.write_contacts.v1
// 权限: esi-characters.write_contacts.v1
func (c *Client) AddCharacterContacts(ctx context.Context, characterID int32, standing float64, body []int32, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	query.Set("standing", strconv.FormatFloat(standing, 'f', -1, 64))
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
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

// EditCharacterContacts Edit contacts.
// EditCharacterContacts 编辑联系人.
//
// Route: PUT /characters/{character_id}/contacts/
// 路由: PUT /characters/{character_id}/contacts/
// Scopes: esi-characters.write_contacts.v1
// 权限: esi-characters.write_contacts.v1
func (c *Client) EditCharacterContacts(ctx context.Context, characterID int32, standing float64, body []int32, opts ...RequestOption) error {
	query, headers := newRequestOptions(opts...)
	query.Set("standing", strconv.FormatFloat(standing, 'f', -1, 64))
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/characters/{character_id}/contacts/", pathParams, query, headers, body)
	return err
}
