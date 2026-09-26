package goeve

import (
	"fmt"
	"strings"
)

// APIError represents an error response returned by the EVE Swagger Interface (ESI) server.
// It covers all error models defined in the swagger document:
// bad_request (400), unauthorized (401), forbidden (403), error_limited (420),
// internal_server_error (500), service_unavailable (503) and gateway_timeout (504).
//
// APIError 表示 EVE Swagger Interface (ESI) 服务器返回的错误响应，
// 涵盖 swagger 文档中定义的全部错误模型：
// bad_request (400)、unauthorized (401)、forbidden (403)、error_limited (420)、
// internal_server_error (500)、service_unavailable (503) 和 gateway_timeout (504)。
type APIError struct {
	// Message is the human-readable error message returned by the server
	// (the JSON field is named "error").
	// Message 为服务器返回的可读错误信息（JSON 字段名为 "error"）。
	Message string `json:"error"`

	// SSOStatus is the status code received from the SSO. Only set on 403 responses.
	// SSOStatus 为从 SSO 收到的状态码，仅在 403 响应中设置。
	SSOStatus int `json:"sso_status,omitempty"`

	// Timeout is the number of seconds the request was given. Only set on 504 responses.
	// Timeout 为请求被允许的执行秒数，仅在 504 响应中设置。
	Timeout int `json:"timeout,omitempty"`

	// StatusCode is the HTTP status code of the response.
	// StatusCode 为 HTTP 响应状态码。
	StatusCode int `json:"-"`

	// RawBody is the raw response body, kept when it could not be decoded into the fields above.
	// RawBody 为原始响应体，当无法解码到上述字段时保留。
	RawBody string `json:"-"`
}

// Error implements the error interface.
// Error 实现 error 接口。
func (e *APIError) Error() string {
	var b strings.Builder
	fmt.Fprintf(&b, "esi: status %d", e.StatusCode)
	if e.Message != "" {
		fmt.Fprintf(&b, ": %s", e.Message)
	}
	if e.SSOStatus != 0 {
		fmt.Fprintf(&b, " (sso_status: %d)", e.SSOStatus)
	}
	if e.Timeout != 0 {
		fmt.Fprintf(&b, " (timeout: %ds)", e.Timeout)
	}
	if e.RawBody != "" && e.Message == "" {
		fmt.Fprintf(&b, ": %s", e.RawBody)
	}
	return b.String()
}
