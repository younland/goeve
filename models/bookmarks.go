package models

import (
	"net/url"
	"strconv"
	"time"
)

// GetCharactersCharacterIdBookmarks 200 ok object.
// GetCharactersCharacterIdBookmarks 200 ok 对象.
type GetCharactersCharacterIdBookmarks struct {
	// BookmarkId bookmark_id integer.
	// BookmarkId 位标 ID 整数.
	BookmarkId int32 `json:"bookmark_id"`
	// Coordinates Optional object that is returned if a bookmark was made on a planet or a random location in space.
	// Coordinates 如果位标是在行星或太空中某个随机位置创建的，则返回的可选对象。
	Coordinates GetCharactersCharacterIdBookmarksCoordinates `json:"coordinates"`
	// Created created string.
	// Created 创建时间字符串.
	Created time.Time `json:"created"`
	// CreatorId creator_id integer.
	// CreatorId 创建者 ID 整数.
	CreatorId int32 `json:"creator_id"`
	// FolderId folder_id integer.
	// FolderId 文件夹 ID 整数.
	FolderId int32 `json:"folder_id"`
	// Item Optional object that is returned if a bookmark was made on a particular item.
	// Item 如果位标是在某个特定物品上创建的，则返回的可选对象。
	Item GetCharactersCharacterIdBookmarksItem `json:"item"`
	// Label label string.
	// Label label 字符串.
	Label string `json:"label"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int32 `json:"location_id"`
	// Notes notes string.
	// Notes notes 字符串.
	Notes string `json:"notes"`
}

// GetCharactersCharacterIdBookmarksCoordinates Optional object that is returned if a bookmark was made on a planet or a random location in space.
// GetCharactersCharacterIdBookmarksCoordinates 如果位标是在行星或太空中某个随机位置创建的，则返回的可选对象。
type GetCharactersCharacterIdBookmarksCoordinates struct {
	// X x number.
	// X x 数字.
	X float64 `json:"x"`
	// Y y number.
	// Y y 数字.
	Y float64 `json:"y"`
	// Z z number.
	// Z z 数字.
	Z float64 `json:"z"`
}

// GetCharactersCharacterIdBookmarksFolders 200 ok object.
// GetCharactersCharacterIdBookmarksFolders 200 ok 对象.
type GetCharactersCharacterIdBookmarksFolders struct {
	// FolderId folder_id integer.
	// FolderId 文件夹 ID 整数.
	FolderId int32 `json:"folder_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// GetCharactersCharacterIdBookmarksItem Optional object that is returned if a bookmark was made on a particular item.
// GetCharactersCharacterIdBookmarksItem 如果位标是在某个特定物品上创建的，则返回的可选对象。
type GetCharactersCharacterIdBookmarksItem struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetCorporationsCorporationIdBookmarks 200 ok object.
// GetCorporationsCorporationIdBookmarks 200 ok 对象.
type GetCorporationsCorporationIdBookmarks struct {
	// BookmarkId bookmark_id integer.
	// BookmarkId 位标 ID 整数.
	BookmarkId int32 `json:"bookmark_id"`
	// Coordinates Optional object that is returned if a bookmark was made on a planet or a random location in space.
	// Coordinates 如果位标是在行星或太空中某个随机位置创建的，则返回的可选对象。
	Coordinates GetCorporationsCorporationIdBookmarksCoordinates `json:"coordinates"`
	// Created created string.
	// Created 创建时间字符串.
	Created time.Time `json:"created"`
	// CreatorId creator_id integer.
	// CreatorId 创建者 ID 整数.
	CreatorId int32 `json:"creator_id"`
	// FolderId folder_id integer.
	// FolderId 文件夹 ID 整数.
	FolderId int32 `json:"folder_id"`
	// Item Optional object that is returned if a bookmark was made on a particular item.
	// Item 如果位标是在某个特定物品上创建的，则返回的可选对象。
	Item GetCorporationsCorporationIdBookmarksItem `json:"item"`
	// Label label string.
	// Label label 字符串.
	Label string `json:"label"`
	// LocationId location_id integer.
	// LocationId location_id 整数.
	LocationId int32 `json:"location_id"`
	// Notes notes string.
	// Notes notes 字符串.
	Notes string `json:"notes"`
}

// GetCorporationsCorporationIdBookmarksCoordinates Optional object that is returned if a bookmark was made on a planet or a random location in space.
// GetCorporationsCorporationIdBookmarksCoordinates 如果位标是在行星或太空中某个随机位置创建的，则返回的可选对象。
type GetCorporationsCorporationIdBookmarksCoordinates struct {
	// X x number.
	// X x 数字.
	X float64 `json:"x"`
	// Y y number.
	// Y y 数字.
	Y float64 `json:"y"`
	// Z z number.
	// Z z 数字.
	Z float64 `json:"z"`
}

// GetCorporationsCorporationIdBookmarksFolders 200 ok object.
// GetCorporationsCorporationIdBookmarksFolders 200 ok 对象.
type GetCorporationsCorporationIdBookmarksFolders struct {
	// CreatorId creator_id integer.
	// CreatorId 创建者 ID 整数.
	CreatorId int32 `json:"creator_id"`
	// FolderId folder_id integer.
	// FolderId 文件夹 ID 整数.
	FolderId int32 `json:"folder_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// GetCorporationsCorporationIdBookmarksItem Optional object that is returned if a bookmark was made on a particular item.
// GetCorporationsCorporationIdBookmarksItem 如果位标是在某个特定物品上创建的，则返回的可选对象。
type GetCorporationsCorporationIdBookmarksItem struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// GetCharactersCharacterIdBookmarksFoldersParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdBookmarksFoldersParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdBookmarksFoldersParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdBookmarksFoldersParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCharactersCharacterIdBookmarksParams holds the optional query and header parameters of the request.
// GetCharactersCharacterIdBookmarksParams 保存请求的可选查询与头部参数。
type GetCharactersCharacterIdBookmarksParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCharactersCharacterIdBookmarksParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCorporationsCorporationIdBookmarksFoldersParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdBookmarksFoldersParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdBookmarksFoldersParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCorporationsCorporationIdBookmarksFoldersParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}

// GetCorporationsCorporationIdBookmarksParams holds the optional query and header parameters of the request.
// GetCorporationsCorporationIdBookmarksParams 保存请求的可选查询与头部参数。
type GetCorporationsCorporationIdBookmarksParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
	// Page Which page of results to return.
	// Page 返回第几页结果.
	Page *int32
	// Token Access token to use if unable to set a header.
	// Token 如果无法设置请求头，则使用此访问令牌.
	Token *string
}

func (p *GetCorporationsCorporationIdBookmarksParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	if p.Page != nil {
		if query == nil {
			query = url.Values{}
		}
		query.Set("page", strconv.FormatInt(int64(*p.Page), 10))
	}
	if p.Token != nil && *p.Token != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("token", *p.Token)
	}
	return query, headers
}
