package models

import (
	"net/url"
	"time"
)

// GetStatus 200 ok object.
// GetStatus 200 ok 对象.
type GetStatus struct {
	// Players Current online player count.
	// Players 当前在线玩家数量.
	Players int32 `json:"players"`
	// ServerVersion Running version as string.
	// ServerVersion 以字符串表示的运行版本.
	ServerVersion string `json:"server_version"`
	// StartTime Server start timestamp.
	// StartTime 服务器启动时间戳.
	StartTime time.Time `json:"start_time"`
	// Vip If the server is in VIP mode.
	// Vip 服务器是否处于 VIP 模式.
	Vip bool `json:"vip"`
}

// GetStatusParams holds the optional query and header parameters of the request.
// GetStatusParams 保存请求的可选查询与头部参数。
type GetStatusParams struct {
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetStatusParams) Values() (url.Values, map[string]string) {
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
	return query, headers
}
