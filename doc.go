// Package goeve provides a Go client library for the NetEase EVE Online
// EVE Swagger Interface (ESI), hand-written against the online API
// specification (https://ali-esi.evepc.163.com/latest/swagger.json).
//
// 包 goeve 是网易 EVE Online EVE Swagger Interface (ESI) 的 Go 客户端库，
// 基于线上接口规范（https://ali-esi.evepc.163.com/latest/swagger.json）手写实现。
//
// Base URL / 基础地址: https://ali-esi.evepc.163.com/latest
//
// Authentication / 认证:
// The library integrates the NetEase EVE SSO OAuth2 flow
// (see token.go: BuildAuthorizeURL, GetTokenFromCode, RefreshAccessToken).
// 本库集成了网易 EVE SSO OAuth2 授权流程
// （见 token.go：BuildAuthorizeURL、GetTokenFromCode、RefreshAccessToken）。
package goeve
