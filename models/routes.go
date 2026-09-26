package models

import (
	"net/url"
	"strconv"
	"strings"
)

// GetOriginDestinationParams holds the optional query and header parameters of the request.
// GetOriginDestinationParams 保存请求的可选查询与头部参数。
type GetOriginDestinationParams struct {
	// Avoid avoid solar system ID(s)
	// Avoid 要避开的星系 ID.
	Avoid []int32
	// Connections connected solar system pairs.
	// Connections 相连的星系对.
	Connections [][]int32
	// Datasource The server name you would like data from.
	// Datasource 你希望获取数据的服务器名称.
	Datasource *string
	// Flag route security preference.
	// Flag 路线安全偏好.
	Flag *string
	// IfNoneMatch ETag from a previous request. A 304 will be returned if this matches the current ETag.
	// IfNoneMatch 来自先前请求的 ETag。如果与当前 ETag 匹配，将返回 304.
	IfNoneMatch *string
}

func (p *GetOriginDestinationParams) Values() (url.Values, map[string]string) {
	if p == nil {
		return nil, nil
	}
	var query url.Values
	var headers map[string]string
	for _, v := range p.Avoid {
		if query == nil {
			query = url.Values{}
		}
		query.Add("avoid", strconv.Itoa(int(v)))
	}
	for _, pair := range p.Connections {
		if len(pair) == 0 {
			continue
		}
		parts := make([]string, len(pair))
		for i, v := range pair {
			parts[i] = strconv.Itoa(int(v))
		}
		if query == nil {
			query = url.Values{}
		}
		query.Add("connections", strings.Join(parts, "|"))
	}
	if p.Datasource != nil && *p.Datasource != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("datasource", *p.Datasource)
	}
	if p.Flag != nil && *p.Flag != "" {
		if query == nil {
			query = url.Values{}
		}
		query.Set("flag", *p.Flag)
	}
	if p.IfNoneMatch != nil && *p.IfNoneMatch != "" {
		if headers == nil {
			headers = map[string]string{}
		}
		headers["If-None-Match"] = *p.IfNoneMatch
	}
	return query, headers
}
