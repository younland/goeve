package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdBookmarks List bookmarks.
// GetCharactersCharacterIdBookmarks 列出位标.
//
// Route: GET /characters/{character_id}/bookmarks/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/bookmarks/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_character_bookmarks.v1
// 权限: esi-bookmarks.read_character_bookmarks.v1
func (c *Client) GetCharactersCharacterIdBookmarks(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdBookmarksParams) ([]models.GetCharactersCharacterIdBookmarks, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdBookmarks
	err := c.get(ctx, "/characters/{character_id}/bookmarks/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdBookmarksFolders List bookmark folders.
// GetCharactersCharacterIdBookmarksFolders 列出位标文件夹.
//
// Route: GET /characters/{character_id}/bookmarks/folders/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/bookmarks/folders/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_character_bookmarks.v1
// 权限: esi-bookmarks.read_character_bookmarks.v1
func (c *Client) GetCharactersCharacterIdBookmarksFolders(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdBookmarksFoldersParams) ([]models.GetCharactersCharacterIdBookmarksFolders, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdBookmarksFolders
	err := c.get(ctx, "/characters/{character_id}/bookmarks/folders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdBookmarks List corporation bookmarks.
// GetCorporationsCorporationIdBookmarks 列出军团位标.
//
// Route: GET /corporations/{corporation_id}/bookmarks/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/bookmarks/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_corporation_bookmarks.v1
// 权限: esi-bookmarks.read_corporation_bookmarks.v1
func (c *Client) GetCorporationsCorporationIdBookmarks(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdBookmarksParams) ([]models.GetCorporationsCorporationIdBookmarks, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdBookmarks
	err := c.get(ctx, "/corporations/{corporation_id}/bookmarks/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdBookmarksFolders List corporation bookmark folders.
// GetCorporationsCorporationIdBookmarksFolders 列出军团位标文件夹.
//
// Route: GET /corporations/{corporation_id}/bookmarks/folders/ — This route is cached for up to 3600 seconds
// 路由: GET /corporations/{corporation_id}/bookmarks/folders/ — 该路由缓存长达 3600 秒
// Scopes: esi-bookmarks.read_corporation_bookmarks.v1
// 权限: esi-bookmarks.read_corporation_bookmarks.v1
func (c *Client) GetCorporationsCorporationIdBookmarksFolders(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdBookmarksFoldersParams) ([]models.GetCorporationsCorporationIdBookmarksFolders, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdBookmarksFolders
	err := c.get(ctx, "/corporations/{corporation_id}/bookmarks/folders/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
