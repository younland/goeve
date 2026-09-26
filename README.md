# goeve

A Go client library for the NetEase EVE Online EVE Swagger Interface (ESI), hand-written following the [gocloak](https://github.com/Nerzal/gocloak) project structure — bilingual (English / 简体中文) godoc documentation included.

网易 EVE Online EVE Swagger Interface (ESI) 的 Go 客户端库，参考 gocloak 项目结构手写实现，全部接口与字段说明均提供中英双语 godoc 注释。

- **ESI base URL / ESI 基础地址**: `https://ali-esi.evepc.163.com/latest` (migrated from the old `esi.evepc.163.com` / 已由旧地址 `esi.evepc.163.com` 迁移)
- **API spec / 接口规范**: `https://ali-esi.evepc.163.com/latest/swagger.json` / 接口规范线上地址
- **HTTP client / HTTP 客户端**: [resty v2](https://github.com/go-resty/resty) (`github.com/go-resty/resty/v2`)

## Features / 特性

- Full coverage of all **204 ESI operations / 204 个 ESI 接口全覆盖** (38 modules: Character, Corporation, Universe, Market, Fleets, ...)
  覆盖全部 38 个模块：角色、军团、宇宙、市场、舰队等
- Semantic method names (no mechanical operationId conversion / 方法名语义化命名，非 operationId 机械转换)
- Explicit positional parameters: required business parameters plus optional ones (`token`, `page`, filters) in a fixed order, zero value = omit; on GET methods the trailing `ifNoneMatch ...string` variadic parameter is simply omitted when no ETag negotiation is wanted
  显式位置参数：必需业务参数在前，可选参数（`token`、`page`、过滤器等）按固定顺序追加，零值表示不传；GET 方法末尾的变长参数 `ifNoneMatch ...string` 不需要 ETag 协商时直接省略
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
| `character.go` | Character / 角色 | `GetCharacter`, `GetCharacterAssets`, `CalculateCharacterCspaCharge` |
| `corporation.go` | Corporation / 军团 | `GetCorporationInformation`, `GetCorporationStructures` |
| `universe.go` | Universe / 宇宙 | `GetUniverseRegions`, `GetUniverseCategories`, `ResolveNamesToIDs`, `ResolveIDsToNames` |
| `market.go` | Market / 市场 | `GetMarketPrices`, `GetMarketOrders` |
| `wallet.go` | Wallet / 钱包 | `GetCharacterWalletBalance` |
| `fleets.go` | Fleets / 舰队 | `GetFleet`, `MoveFleetMember` |
| `status.go` | Status / 状态 | `GetServerStatus` |

(All 32 modules / 共 32 个模块，见 `goeve_iface.go`)

Public endpoints need no token; every method takes its path/business parameters explicitly, followed by optional ones in a fixed order — `token` (empty = client TokenSource / anonymous), `page` (0 = default), business filters. GET methods end with an optional variadic `ifNoneMatch ...string`: omit it for no ETag negotiation, or pass the ETag string as the last argument:

公开接口无需令牌；每个方法的参数显式传入：路径与业务必需参数在前，可选参数按固定顺序追加——`token`（空 = 使用客户端 TokenSource / 匿名）、`page`（0 = 默认第 1 页）、业务过滤参数。GET 方法末尾是变长可选参数 `ifNoneMatch ...string`：不需要 ETag 协商时省略，需要时把 ETag 作为最后一个参数传入：

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/younland/goeve"
)

func main() {
    client := goeve.NewClient(
        goeve.WithTimeout(30 * time.Second),
        goeve.WithLanguage(goeve.LanguageChinese), // default Accept-Language / 默认响应语言
    )

    // Server status / 服务器状态
    status, err := client.GetServerStatus(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("players=%d version=%s", status.Players, status.ServerVersion)

    // Character public info / 角色公开信息
    character, err := client.GetCharacter(context.Background(), 95234356)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("name=%s corp=%d", character.Name, character.CorporationId)

    // Market orders: regionID and orderType are business parameters, page and typeID optional
    // 市场订单：regionID、orderType 为业务参数，page、typeID 为可选参数（0 = 不传）
    orders, err := client.GetMarketOrders(context.Background(), 10000002, "sell", 1, 0)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("orders=%d", len(orders))
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

// 4. Call authenticated endpoints: token "" uses the client's TokenSource
//    调用受保护接口：token 传 "" 表示使用客户端 TokenSource（也可显式传 token.AccessToken）
balance, err := client.GetCharacterWalletBalance(context.Background(), characterID, "")
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
client := goeve.NewClient(goeve.WithTokenSource(goeve.NewTokenSource(token)))
```

### Refresh / 刷新

```go
token, err := goeve.RefreshAccessToken(context.Background(), refreshToken)
```

> **Warning / 警告**: a refresh token is as powerful as a password — store it safely. Users can revoke it in the official site ("ESI解除授权"). / refresh_token 等同于密码，请妥善保管；用户可在官网"ESI解除授权"解除绑定。

## Pagination, ETag and language / 分页、ETag 与语言

```go
// Pagination: loop page (0 = default first page) until an empty page is returned
// 分页：循环递增 page（0 = 默认第 1 页）直至返回空页
page := int32(1)
for {
    assets, err := client.GetCharacterAssets(ctx, characterID, "", page)
    if err != nil {
        log.Fatal(err)
    }
    if len(assets) == 0 {
        break // last page reached / 已到最后一页
    }
    // ... process assets / 处理资产 ...
    page++
}

// Mail has no page parameter: walk back with lastMailID (0 = from the newest)
// 邮件接口没有 page 参数：用 lastMailID 回溯翻页（0 = 从最新开始）
mails, err := client.GetCharacterMails(ctx, characterID, "", nil, lastMailID)

// ETag negotiation: append the ETag as the trailing variadic argument; on 304 the result is nil and err is nil
// ETag 协商：把 ETag 作为末尾变长参数传入；服务器返回 304 时结果为 nil 且无错误
mails, err = client.GetCharacterMails(ctx, characterID, "", nil, 0, etag)
if err == nil && mails == nil {
    // not modified / 数据未变化
}

// Language: client-level WithLanguage applies Accept-Language to every request
// 语言：客户端级 WithLanguage 为所有请求设置 Accept-Language
client := goeve.NewClient(goeve.WithLanguage(goeve.LanguageChinese))
region, err := client.GetUniverseRegion(ctx, regionID)
```

## Project structure / 项目结构

Hand-written, gocloak-style layout (no code generator / 无代码生成器，手写实现)：

```
goeve/
├── client.go        # Client core: client options, request pipeline (resty v2)
│                    # 客户端核心：客户端选项与请求管线
├── goeve_iface.go   # ClientIface: all 204 method signatures, grouped by module
│                    # 客户端接口（按模块分组）
├── token.go         # NetEase EVE SSO OAuth2 (authorize/code/refresh/TokenSource)
├── errors.go        # APIError covering all ESI error models 结构化错误
├── models.go        # Shared constants (languages, datasource) 共享常量
├── <module>.go      # API methods per module: character.go, corporation.go, universe.go, ...
│                    # 各模块的 API 方法（204 个，按模块分文件存放）
├── models/          # Response model types, one file per module
│                    # 响应模型类型子包，按模块分文件（character.go, corporation.go, ...）
└── go.mod / go.sum
```

Response types live in the `models` sub package and are referenced as
`models.Character`, `models.MarketOrder`, `models.WalletJournalEntry`, etc.
All request parameters are explicit positional arguments: after the required
path/business parameters, optional ones are appended in a fixed order
(`token`, `page`, business filters), where a zero value (`""`, `0`, `nil`,
`false`) means "do not send". GET methods additionally end with a variadic
`ifNoneMatch ...string` for ETag negotiation — omit it when not needed.

响应类型位于 `models` 子包，以 `models.Character`、`models.MarketOrder`、
`models.WalletJournalEntry` 等形式引用。所有请求参数均为显式位置参数：
必需的路径/业务参数之后，按固定顺序追加可选参数（`token`、`page`、
业务过滤参数），零值（`""`、`0`、`nil`、`false`）表示不传。GET 方法末尾
另有变长参数 `ifNoneMatch ...string` 用于 ETag 协商，不需要时可省略。

When CCP/NetEase publishes an updated spec, download it from
`https://ali-esi.evepc.163.com/latest/swagger.json` for reference and extend the
corresponding module files.

当接口规范更新时，可从 `https://ali-esi.evepc.163.com/latest/swagger.json`
下载最新规范作参考，并在对应模块文件中补充新接口。

## AI Skills / 智能体技能

This repo ships an agent skill (`.agents/skills/goeve/SKILL.md`) that teaches coding agents how to use this library — endpoint lookup by module, SSO flow, explicit-parameter conventions and error handling.

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
