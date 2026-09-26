package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// DeleteCharactersCharacterIdMailLabelsLabelId Delete a mail label.
// DeleteCharactersCharacterIdMailLabelsLabelId 删除邮件标签.
//
// Route: DELETE /characters/{character_id}/mail/labels/{label_id}/
// 路由: DELETE /characters/{character_id}/mail/labels/{label_id}/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) DeleteCharactersCharacterIdMailLabelsLabelId(ctx context.Context, characterId int32, labelId int32, params *models.DeleteCharactersCharacterIdMailLabelsLabelIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "label_id": strconv.FormatInt(int64(labelId), 10)}
	err := c.delete(ctx, "/characters/{character_id}/mail/labels/{label_id}/", pathParams, query, headers)
	return err
}

// DeleteCharactersCharacterIdMailMailId Delete a mail.
// DeleteCharactersCharacterIdMailMailId 删除邮件.
//
// Route: DELETE /characters/{character_id}/mail/{mail_id}/
// 路由: DELETE /characters/{character_id}/mail/{mail_id}/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) DeleteCharactersCharacterIdMailMailId(ctx context.Context, characterId int32, mailId int32, params *models.DeleteCharactersCharacterIdMailMailIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "mail_id": strconv.FormatInt(int64(mailId), 10)}
	err := c.delete(ctx, "/characters/{character_id}/mail/{mail_id}/", pathParams, query, headers)
	return err
}

// GetCharactersCharacterIdMail Return mail headers.
// GetCharactersCharacterIdMail 返回邮件头.
//
// Route: GET /characters/{character_id}/mail/ — This route is cached for up to 30 seconds
// 路由: GET /characters/{character_id}/mail/ — 该路由缓存长达 30 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharactersCharacterIdMail(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMailParams) ([]models.GetCharactersCharacterIdMail, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdMail
	err := c.get(ctx, "/characters/{character_id}/mail/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdMailLabels Get mail labels and unread counts.
// GetCharactersCharacterIdMailLabels 获取邮件标签与未读数量.
//
// Route: GET /characters/{character_id}/mail/labels/ — This route is cached for up to 30 seconds
// 路由: GET /characters/{character_id}/mail/labels/ — 该路由缓存长达 30 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharactersCharacterIdMailLabels(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMailLabelsParams) (*models.GetCharactersCharacterIdMailLabelsOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result *models.GetCharactersCharacterIdMailLabelsOk
	err := c.get(ctx, "/characters/{character_id}/mail/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdMailLists Return mailing list subscriptions.
// GetCharactersCharacterIdMailLists 返回邮件列表订阅.
//
// Route: GET /characters/{character_id}/mail/lists/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/mail/lists/ — 该路由缓存长达 120 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharactersCharacterIdMailLists(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMailListsParams) ([]models.GetCharactersCharacterIdMailLists, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdMailLists
	err := c.get(ctx, "/characters/{character_id}/mail/lists/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdMailMailId Return a mail.
// GetCharactersCharacterIdMailMailId 返回一封邮件.
//
// Route: GET /characters/{character_id}/mail/{mail_id}/ — This route is cached for up to 30 seconds
// 路由: GET /characters/{character_id}/mail/{mail_id}/ — 该路由缓存长达 30 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharactersCharacterIdMailMailId(ctx context.Context, characterId int32, mailId int32, params *models.GetCharactersCharacterIdMailMailIdParams) (*models.GetCharactersCharacterIdMailMailIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "mail_id": strconv.FormatInt(int64(mailId), 10)}
	var result *models.GetCharactersCharacterIdMailMailIdOk
	err := c.get(ctx, "/characters/{character_id}/mail/{mail_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PostCharactersCharacterIdMail Send a new mail.
// PostCharactersCharacterIdMail 发送新邮件.
//
// Route: POST /characters/{character_id}/mail/
// 路由: POST /characters/{character_id}/mail/
// Scopes: esi-mail.send_mail.v1
// 权限: esi-mail.send_mail.v1
func (c *Client) PostCharactersCharacterIdMail(ctx context.Context, characterId int32, body *models.PostCharactersCharacterIdMailMail, params *models.PostCharactersCharacterIdMailParams) (int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	if body == nil {
		return 0, errBodyRequired
	}
	var result int32
	err := c.post(ctx, "/characters/{character_id}/mail/", pathParams, query, headers, body, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// PostCharactersCharacterIdMailLabels Create a mail label.
// PostCharactersCharacterIdMailLabels 创建邮件标签.
//
// Route: POST /characters/{character_id}/mail/labels/
// 路由: POST /characters/{character_id}/mail/labels/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) PostCharactersCharacterIdMailLabels(ctx context.Context, characterId int32, body *models.PostCharactersCharacterIdMailLabelsLabel, params *models.PostCharactersCharacterIdMailLabelsParams) (int32, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	if body == nil {
		return 0, errBodyRequired
	}
	var result int32
	err := c.post(ctx, "/characters/{character_id}/mail/labels/", pathParams, query, headers, body, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// PutCharactersCharacterIdMailMailId Update metadata about a mail.
// PutCharactersCharacterIdMailMailId 更新邮件的元数据.
//
// Route: PUT /characters/{character_id}/mail/{mail_id}/
// 路由: PUT /characters/{character_id}/mail/{mail_id}/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) PutCharactersCharacterIdMailMailId(ctx context.Context, characterId int32, mailId int32, body *models.PutCharactersCharacterIdMailMailIdContents, params *models.PutCharactersCharacterIdMailMailIdParams) error {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10), "mail_id": strconv.FormatInt(int64(mailId), 10)}
	if body == nil {
		return errBodyRequired
	}
	err := c.put(ctx, "/characters/{character_id}/mail/{mail_id}/", pathParams, query, headers, body)
	return err
}
