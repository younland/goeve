package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharacterBookmarks List bookmarks.
// GetCharacterBookmarks 列出位标.
//
// Route: GET /characters/{character_id}/bookmarks/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/bookmarks/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_character_bookmarks.v1
// 权限: esi-bookmarks.read_character_bookmarks.v1
func (c *Client) GetCharacterBookmarks(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterBookmark, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterBookmark
	err := c.get(ctx, "/characters/{character_id}/bookmarks/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterBookmarkFolders List bookmark folders.
// GetCharacterBookmarkFolders 列出位标文件夹.
//
// Route: GET /characters/{character_id}/bookmarks/folders/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/bookmarks/folders/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_character_bookmarks.v1
// 权限: esi-bookmarks.read_character_bookmarks.v1
func (c *Client) GetCharacterBookmarkFolders(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterBookmarkFolder, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterBookmarkFolder
	err := c.get(ctx, "/characters/{character_id}/bookmarks/folders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCorporationBookmarks List corporation bookmarks.
// ListCorporationBookmarks 列出军团位标.
//
// Route: GET /corporations/{corporation_id}/bookmarks/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/bookmarks/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_corporation_bookmarks.v1
// 权限: esi-bookmarks.read_corporation_bookmarks.v1
func (c *Client) ListCorporationBookmarks(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationBookmark, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationBookmark
	err := c.get(ctx, "/corporations/{corporation_id}/bookmarks/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListCorporationBookmarkFolders List corporation bookmark folders.
// ListCorporationBookmarkFolders 列出军团位标文件夹.
//
// Route: GET /corporations/{corporation_id}/bookmarks/folders/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/bookmarks/folders/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_corporation_bookmarks.v1
// 权限: esi-bookmarks.read_corporation_bookmarks.v1
func (c *Client) ListCorporationBookmarkFolders(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationBookmarkFolder, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationBookmarkFolder
	err := c.get(ctx, "/corporations/{corporation_id}/bookmarks/folders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
