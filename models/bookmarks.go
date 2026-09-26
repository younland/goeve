package models

import (
	"time"
)

// CharacterBookmark 200 ok object.
// CharacterBookmark 200 ok 对象.
type CharacterBookmark struct {
	// BookmarkId bookmark_id integer.
	// BookmarkId 位标 ID 整数.
	BookmarkId int32 `json:"bookmark_id"`
	// Coordinates Optional object that is returned if a bookmark was made on a planet or a random location in space.
	// Coordinates 如果位标是在行星或太空中某个随机位置创建的，则返回的可选对象。
	Coordinates BookmarkCoordinates `json:"coordinates"`
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
	Item BookmarkItem `json:"item"`
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

// BookmarkCoordinates Optional object that is returned if a bookmark was made on a planet or a random location in space.
// BookmarkCoordinates 如果位标是在行星或太空中某个随机位置创建的，则返回的可选对象。
type BookmarkCoordinates struct {
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

// CharacterBookmarkFolder 200 ok object.
// CharacterBookmarkFolder 200 ok 对象.
type CharacterBookmarkFolder struct {
	// FolderId folder_id integer.
	// FolderId 文件夹 ID 整数.
	FolderId int32 `json:"folder_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
}

// BookmarkItem Optional object that is returned if a bookmark was made on a particular item.
// BookmarkItem 如果位标是在某个特定物品上创建的，则返回的可选对象。
type BookmarkItem struct {
	// ItemId item_id integer.
	// ItemId item_id 整数.
	ItemId int64 `json:"item_id"`
	// TypeId type_id integer.
	// TypeId type_id 整数.
	TypeId int32 `json:"type_id"`
}

// CorporationBookmark 200 ok object.
// CorporationBookmark 200 ok 对象.
type CorporationBookmark struct {
	// BookmarkId bookmark_id integer.
	// BookmarkId 位标 ID 整数.
	BookmarkId int32 `json:"bookmark_id"`
	// Coordinates Optional object that is returned if a bookmark was made on a planet or a random location in space.
	// Coordinates 如果位标是在行星或太空中某个随机位置创建的，则返回的可选对象。
	Coordinates BookmarkCoordinates `json:"coordinates"`
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
	Item BookmarkItem `json:"item"`
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

// CorporationBookmarkFolder 200 ok object.
// CorporationBookmarkFolder 200 ok 对象.
type CorporationBookmarkFolder struct {
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
