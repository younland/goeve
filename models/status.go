package models

import (
	"time"
)

// ServerStatus 200 ok object.
// ServerStatus 200 ok 对象.
type ServerStatus struct {
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
