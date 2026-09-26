package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// DeleteCharacterMailLabel Delete a mail label.
// DeleteCharacterMailLabel 删除邮件标签.
//
// Route: DELETE /characters/{character_id}/mail/labels/{label_id}/
// 路由: DELETE /characters/{character_id}/mail/labels/{label_id}/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) DeleteCharacterMailLabel(ctx context.Context, token string, characterID int32, labelID int32) error {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "label_id": strconv.FormatInt(int64(labelID), 10)}
	err := c.delete(ctx, "/characters/{character_id}/mail/labels/{label_id}/", pathParams, query, headers)
	return err
}

// DeleteCharacterMail Delete a mail.
// DeleteCharacterMail 删除邮件.
//
// Route: DELETE /characters/{character_id}/mail/{mail_id}/
// 路由: DELETE /characters/{character_id}/mail/{mail_id}/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) DeleteCharacterMail(ctx context.Context, token string, characterID int32, mailID int32) error {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "mail_id": strconv.FormatInt(int64(mailID), 10)}
	err := c.delete(ctx, "/characters/{character_id}/mail/{mail_id}/", pathParams, query, headers)
	return err
}

// GetCharacterMails Return mail headers.
// GetCharacterMails 返回邮件头.
//
// Route: GET /characters/{character_id}/mail/ — This route is cached for up to 30 seconds
// 路由: GET /characters/{character_id}/mail/ — 该路由缓存长达 30 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharacterMails(ctx context.Context, token string, characterID int32, labels []int32, lastMailID int32, ifNoneMatch ...string) ([]models.MailHeader, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	for _, v := range labels {
		query.Add("labels", strconv.FormatInt(int64(v), 10))
	}
	if lastMailID != 0 {
		query.Set("last_mail_id", strconv.FormatInt(int64(lastMailID), 10))
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.MailHeader
	err := c.get(ctx, "/characters/{character_id}/mail/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterMailLabels Get mail labels and unread counts.
// GetCharacterMailLabels 获取邮件标签与未读数量.
//
// Route: GET /characters/{character_id}/mail/labels/ — This route is cached for up to 30 seconds
// 路由: GET /characters/{character_id}/mail/labels/ — 该路由缓存长达 30 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharacterMailLabels(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.MailLabels, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result *models.MailLabels
	err := c.get(ctx, "/characters/{character_id}/mail/labels/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterMailLists Return mailing list subscriptions.
// GetCharacterMailLists 返回邮件列表订阅.
//
// Route: GET /characters/{character_id}/mail/lists/ — This route is cached for up to 120 seconds
// 路由: GET /characters/{character_id}/mail/lists/ — 该路由缓存长达 120 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharacterMailLists(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.MailingList, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.MailingList
	err := c.get(ctx, "/characters/{character_id}/mail/lists/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterMail Return a mail.
// GetCharacterMail 返回一封邮件.
//
// Route: GET /characters/{character_id}/mail/{mail_id}/ — This route is cached for up to 30 seconds
// 路由: GET /characters/{character_id}/mail/{mail_id}/ — 该路由缓存长达 30 秒
// Scopes: esi-mail.read_mail.v1
// 权限: esi-mail.read_mail.v1
func (c *Client) GetCharacterMail(ctx context.Context, token string, characterID int32, mailID int32, ifNoneMatch ...string) (*models.Mail, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if len(ifNoneMatch) > 0 && ifNoneMatch[0] != "" {
		headers["If-None-Match"] = ifNoneMatch[0]
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "mail_id": strconv.FormatInt(int64(mailID), 10)}
	var result *models.Mail
	err := c.get(ctx, "/characters/{character_id}/mail/{mail_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SendCharacterMail Send a new mail.
// SendCharacterMail 发送新邮件.
//
// Route: POST /characters/{character_id}/mail/
// 路由: POST /characters/{character_id}/mail/
// Scopes: esi-mail.send_mail.v1
// 权限: esi-mail.send_mail.v1
func (c *Client) SendCharacterMail(ctx context.Context, token string, characterID int32, body *models.MailRequest) (int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
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

// CreateCharacterMailLabel Create a mail label.
// CreateCharacterMailLabel 创建邮件标签.
//
// Route: POST /characters/{character_id}/mail/labels/
// 路由: POST /characters/{character_id}/mail/labels/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) CreateCharacterMailLabel(ctx context.Context, token string, characterID int32, body *models.MailLabelRequest) (int32, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
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

// UpdateCharacterMail Update metadata about a mail.
// UpdateCharacterMail 更新邮件的元数据.
//
// Route: PUT /characters/{character_id}/mail/{mail_id}/
// 路由: PUT /characters/{character_id}/mail/{mail_id}/
// Scopes: esi-mail.organize_mail.v1
// 权限: esi-mail.organize_mail.v1
func (c *Client) UpdateCharacterMail(ctx context.Context, token string, characterID int32, mailID int32, body *models.MailMetadata) error {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10), "mail_id": strconv.FormatInt(int64(mailID), 10)}
	if body == nil {
		return errBodyRequired
	}
	if err := c.put(ctx, "/characters/{character_id}/mail/{mail_id}/", pathParams, query, headers, body); err != nil {
		return err
	}
	return nil
}
