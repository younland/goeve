# goeve

A Go client library for the NetEase EVE Online EVE Swagger Interface (ESI), hand-written following the [gocloak](https://github.com/Nerzal/gocloak) project structure — bilingual (English / 简体中文) godoc documentation included.

网易 EVE Online EVE Swagger Interface (ESI) 的 Go 客户端库，参考 gocloak 项目结构手写实现，全部接口与字段说明均提供中英双语 godoc 注释。

- **ESI base URL / ESI 基础地址**: `https://ali-esi.evepc.163.com/latest` (migrated from the old `esi.evepc.163.com` / 已由旧地址 `esi.evepc.163.com` 迁移)
- **API spec / 接口规范**: `https://ali-esi.evepc.163.com/latest/swagger.json` / 接口规范线上地址
- **HTTP client / HTTP 客户端**: [resty v2](https://github.com/go-resty/resty) (`github.com/go-resty/resty/v2`)

## Features / 特性

- Full coverage of all **204 ESI operations / 204 个 ESI 接口全覆盖** (38 modules: Character, Corporation, Universe, Market, Fleets, ...)
  覆盖全部 38 个模块：角色、军团、宇宙、市场、舰队等
- Bilingual godoc on every method, model and field (English + 简体中文) / 每个方法、模型、字段都有中英双语注释
- Integrated NetEase EVE SSO OAuth2 token acquisition (implicit + authorization code + refresh) / 内置网易 EVE SSO OAuth2 令牌获取（隐式、授权码、刷新）
- ETag / `If-None-Match` (304) support / ETag 缓存协商支持
- Structured `APIError` for all ESI error models / 结构化的 APIError，覆盖全部 ESI 错误模型

## Install / 安装

```bash
go get github.com/younland/goeve
```

## Quick start / 快速开始

Methods are called directly on the `Client`, and the source files are organized by ESI module (one file per module):

方法直接挂在 `Client` 上调用，源文件按 ESI 模块分开存放（每模块一个文件）：

| File / 文件 | Module / 模块 | Example methods / 示例方法 |
|---|---|---|
| `character.go` | Character / 角色 | `GetCharacterId`, `GetCharacterIdAssets`, `PostCharacterIdCspa` |
| `corporation.go` | Corporation / 军团 | `GetCorporationId`, `GetCorporationIdStructures` |
| `universe.go` | Universe / 宇宙 | `GetRegions`, `GetCategories`, `PostIds`, `PostNames` |
| `market.go` | Market / 市场 | `GetMarketsPrices`, `GetMarketsRegionIdOrders` |
| `wallet.go` | Wallet / 钱包 | `GetCharactersCharacterIdWallet` |
| `fleets.go` | Fleets / 舰队 | `GetFleetId`, `PutFleetIdMembersMemberId` |
| `status.go` | Status / 状态 | `GetStatus` |

(All 32 modules / 共 32 个模块，见 `goeve_iface.go`)

Public endpoints need no token / 公开接口无需令牌：

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/younland/goeve"
)

func main() {
    client := goeve.NewClient(goeve.WithTimeout(30 * time.Second))

    // Server status / 服务器状态
    status, err := client.GetStatus(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("players=%d version=%s", status.Players, status.ServerVersion)

    // Character public info / 角色公开信息
    character, err := client.GetCharacterId(context.Background(), 95234356, nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("name=%s corp=%d", character.Name, character.CorporationId)
}
```

## Authentication / 认证授权

NetEase EVE SSO (`https://login.evepc.163.com`) implements OAuth2 with two grant flows. The library reuses the official ESI application client_id by default (NetEase offers no public application registration), and supports at most 4 scopes per authorization.

网易 EVE SSO（`https://login.evepc.163.com`）提供两种 OAuth2 授权模式。默认复用官方 ESI 应用的 client_id（网易未开放自助申请），一次最多申请 4 项权限（scope）。

### Authorization code flow (recommended / 推荐)

```go
// 1. Open this URL in a browser, log in and pick a character / 在浏览器打开，登录并选择角色
authURL := goeve.BuildAuthorizeURL(goeve.AuthorizeConfig{
    ResponseType: goeve.ResponseTypeCode,
    State:        "random-state",
    Scope:        "esi-wallet.read_character_wallet.v1 esi-location.read_location.v1",
})
// -> redirect: https://esi.evepc.163.com/ui/oauth2-redirect.html?code=...&state=...

// 2. Paste the redirect URL, exchange the code for tokens / 粘贴跳转地址，用授权码换取令牌
code, _ := goeve.ParseAuthorizationCode(redirectURL)
token, err := goeve.GetTokenFromCode(context.Background(), code, goeve.AuthorizeConfig{})

// 3. Attach a TokenSource: access tokens refresh automatically / 挂载 TokenSource，自动续期
ts := goeve.NewTokenSource(token)
client := goeve.NewClient(goeve.WithTokenSource(ts))

// 4. Call authenticated endpoints / 调用受保护接口
balance, err := client.GetCharactersCharacterIdWallet(context.Background(), characterID, nil)
```

### Implicit flow / 隐式模式

The access token is returned directly in the redirect URL fragment (expires in ~20 minutes, no refresh token):

access_token 直接拼在跳转地址的 fragment 中（约 20 分钟过期，无 refresh_token）：

```go
authURL := goeve.BuildAuthorizeURL(goeve.AuthorizeConfig{
    ResponseType: goeve.ResponseTypeToken,
    Scope:        "esi-location.read_location.v1",
})
token, err := goeve.ParseImplicitRedirect(redirectURL)
client := goeve.NewClient(goeve.WithAuthToken(token.AccessToken))
```

### Refresh / 刷新

```go
token, err := goeve.RefreshAccessToken(context.Background(), refreshToken)
```

> **Warning / 警告**: a refresh token is as powerful as a password — store it safely. Users can revoke it in the official site ("ESI解除授权"). / refresh_token 等同于密码，请妥善保管；用户可在官网"ESI解除授权"解除绑定。

## Pagination, ETag and language / 分页、ETag 与语言

```go
import "github.com/younland/goeve/models"

// Pagination via the Page parameter (repeat until an empty page) / 通过 Page 参数分页（循环直至返回空页）
params := &models.GetCharactersCharacterIdMailParams{Page: ptr(int32(1))}
mails, err := client.GetCharactersCharacterIdMail(ctx, characterID, params)
for len(mails) > 0 {
	*params.Page++
	mails, err = client.GetCharactersCharacterIdMail(ctx, characterID, params)
}

// ETag: pass If-None-Match; on 304 the result is nil and err is nil
// ETag 协商：传入 If-None-Match；服务器返回 304 时结果为 nil 且无错误
params2 := &models.GetCharactersCharacterIdMailParams{IfNoneMatch: &etag}
mails2, err := client.GetCharactersCharacterIdMail(ctx, characterID, params2)
if err == nil && mails2 == nil {
	// not modified / 数据未变化
}

// Simplified Chinese responses / 简体中文返回
lang := goeve.LanguageChinese
params3 := &models.GetRegionsRegionIdParams{AcceptLanguage: &lang}
region, err := client.GetRegionsRegionId(ctx, regionID, params3)

func ptr[T any](v T) *T { return &v }
```

## Project structure / 项目结构

Hand-written, gocloak-style layout (no code generator / 无代码生成器，手写实现)：

```
goeve/
├── client.go        # Client core: options, resty v2 request pipeline 客户端核心：选项与请求管线
├── goeve_iface.go   # ClientIface: all 204 methods, grouped by module 客户端接口（按模块分组）
├── token.go         # NetEase EVE SSO OAuth2 (authorize/code/refresh/TokenSource)
├── errors.go        # APIError covering all ESI error models 结构化错误
├── models.go        # Shared constants (languages, datasource) 共享常量
├── <module>.go      # API methods per module: character.go, corporation.go, universe.go, ...
│                    # 各模块的 API 方法（204 个，按模块分文件存放）
├── models/          # Model + <Method>Params types, one file per module
│                    # 模型与参数结构体子包，按模块分文件（character.go, corporation.go, ...）
└── go.mod / go.sum
```

Response types live in the `models` sub package and are referenced as
`models.GetCharactersCharacterId`, `models.GetRegionsRegionIdParams`, etc.

响应类型位于 `models` 子包，以 `models.GetCharactersCharacterId`、
`models.GetRegionsRegionIdParams` 等形式引用。

When CCP/NetEase publishes an updated spec, download it from
`https://ali-esi.evepc.163.com/latest/swagger.json` for reference and extend the
corresponding module files.

当接口规范更新时，可从 `https://ali-esi.evepc.163.com/latest/swagger.json`
下载最新规范作参考，并在对应模块文件中补充新接口。

## AI Skills / 智能体技能

This repo ships an agent skill (`.agents/skills/goeve/SKILL.md`) that teaches coding agents how to use this library — endpoint lookup by module, SSO flow, params/pagination conventions and error handling.

本仓库内置智能体技能（`.agents/skills/goeve/SKILL.md`），向 Claude Code、Cursor、Kimi Code 等编程智能体传授本库用法。

Install it with the [skills CLI](https://github.com/vercel-labs/skills) (no global install needed / 无需全局安装):

```bash
# 本仓库已内置该技能（.agents/skills/goeve），在本项目中智能体自动可用、无需安装
# from GitHub, into another project（需要本仓库已推送到 GitHub）
npx skills add younland/goeve --skill goeve          # 安装到当前项目
npx skills add younland/goeve --skill goeve -g       # 全局安装（-y 跳过确认）

# or from a local checkout / 本地仓库直接安装到其他项目
npx skills add /path/to/goeve/.agents/skills/goeve
```

After installation the skill is picked up automatically when an agent works with this library.

安装后智能体在处理与本库相关的任务时会自动加载该技能。

## Testing / 测试

```bash
go test ./...   # includes live calls to the public ESI endpoints / 包含对公开 ESI 接口的真实调用
```

## Reference / 参考

- EVE ESI 认证说明（OAuth2 流程参考）: <https://sodacooky.github.io/2023/EVE_ESI/>
- Architecture inspired by [gocloak](https://github.com/Nerzal/gocloak) / 架构参考自 gocloak

## License / 许可证

MIT
