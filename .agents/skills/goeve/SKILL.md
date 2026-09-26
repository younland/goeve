---
name: goeve
description: >
  使用 goeve 客户端库（github.com/younland/goeve）调用网易 EVE Online ESI API 时触发。
  适用于：编写调用 ESI 接口的 Go 代码、查找某个游戏功能（市场、钱包、舰队、合同、资产等）
  对应的客户端方法、对接网易 EVE SSO OAuth2 授权、处理 ETag/分页/错误。
  Use the goeve Go client library for the NetEase EVE Online ESI API — endpoint lookup,
  SSO OAuth2 auth, explicit-parameter/pagination/error conventions.
metadata:
  author: younland
  version: "4.0"
---

# goeve — 网易 EVE Online ESI Go 客户端

## 目标 / Goal

用 `github.com/younland/goeve` 调用网易 EVE Online ESI（`https://ali-esi.evepc.163.com/latest`）。
本 Skill 提供：模块定位表、调用约定、SSO 授权流程与常见坑。逐模块接口说明见
[references/modules.md](references/modules.md)（按需加载）。

## 前置检查 / Prerequisites

确认依赖已引入；未引入则执行：

```bash
grep goeve go.mod || go get github.com/younland/goeve
```

## 调用约定 / Conventions

| 约定 | 说明 |
|---|---|
| 创建客户端 | `client := goeve.NewClient(goeve.WithTimeout(30*time.Second), goeve.WithDebug(true))`；客户端选项还有 `WithBaseURL` / `WithHTTPClient` / `WithTokenSource` / `WithLanguage`（默认 Accept-Language） |
| 返回值 | 只有 `(result, error)`，不返回 `*resty.Response`。对象返回指针 `*models.Xxx`，数组返回切片 `[]models.Xxx`，204 响应只返回 `error` |
| 参数 | 第一个参数永远是 `context.Context`；**需授权接口的 `token` 紧随其后**（ctx 后第一位）；之后是**路径与业务必需参数**（如 `characterID`、`regionID, orderType`），**再按固定顺序追加可选参数**（`page`、业务过滤参数、变长 `ifNoneMatch`）。签名以 `goeve_iface.go` 与方法 godoc 为准 |
| 零值约定 | `token ""` = 使用客户端 TokenSource / 匿名；`page 0` = 默认第 1 页；string 用 `""`、数值用 `0`、切片用 `nil`、bool 用 `false` 表示"不传该参数"；GET 方法的 ETag 变长参数直接省略即可 |
| 命名 | 方法名已语义化（如 `get_characters_character_id` → `GetCharacter`，`get_markets_region_id_orders` → `GetMarketOrders`，`get_characters_character_id_wallet` → `GetCharacterWalletBalance`），**不要按 operationId 机械推导**，以 `goeve_iface.go` 与方法 godoc 为准 |
| 类型 | ID 类型以 swagger 为准：character/alliance/corporation/region 是 `int32`，structure/fleet/killmail/starbase 相关多为 `int64`，不要凭习惯强转 |

### 显式参数约定 / Explicit parameter order

没有请求选项、没有参数结构体——每个方法的全部参数都是显式位置参数，顺序固定：

1. **`token string`**：仅需授权接口存在，且固定为 `ctx` 之后第一位；空字符串 = 使用客户端 TokenSource / 匿名；显式传入则优先于 TokenSource
2. **路径参数 + 业务必需参数**（如 `characterID`、`regionID, orderType`、`killmailHash, killmailID`）
3. **`page int32`**：支持分页的接口存在；0 = 不传（默认第 1 页）
4. **该接口特有的可选业务查询参数**：按签名顺序排列，如 `fromID int64`、`typeID int32`、`maxWarID int32`、`includeCompleted bool`、`labels []int32`、`lastMailID int32`、`strict bool`、`watched bool`、`filter string`、`fromEvent int32`、`labelIDs []int32`、`avoid []int32`、`connections [][]int32`、`flag string`
5. **`ifNoneMatch ...string`**：仅 GET 接口存在，且总是最后一个参数；**变长可选参数**——不传表示不做 ETag 协商，传一个 ETag 字符串则携带 If-None-Match 头，服务器返回 304 时结果为 nil + nil error

签名示例（均可在 `goeve_iface.go` 中核对）：

```go
GetServerStatus(ctx, ifNoneMatch ...string) (*models.ServerStatus, error)
GetCharacter(ctx, characterID int32, ifNoneMatch ...string) (*models.Character, error)
GetCharacterWalletBalance(ctx, token string, characterID int32, ifNoneMatch ...string) (float64, error)
GetCharacterMails(ctx, token string, characterID int32, labels []int32, lastMailID int32, ifNoneMatch ...string) ([]models.MailHeader, error)
GetMarketOrders(ctx, regionID int32, orderType string, page int32, typeID int32, ifNoneMatch ...string) ([]models.MarketOrder, error)
GetRoute(ctx, destination int32, origin int32, avoid []int32, connections [][]int32, flag string, ifNoneMatch ...string) ([]int32, error)
GetWars(ctx, maxWarID int32, ifNoneMatch ...string) ([]int32, error)
```

## 模块定位 / Module index

方法按模块分文件存放（`<module>.go`，模型在 `models/` 子包同名文件），完整索引见 `goeve_iface.go`：

| 需求场景 | 文件 | 常用方法 |
|---|---|---|
| 服务器状态 | `status.go` | `GetServerStatus` |
| 角色公开信息/通知/声望 | `character.go` | `GetCharacter`, `GetCharacterNotifications` |
| 军团信息/成员/建筑 | `corporation.go` | `GetCorporationInformation`, `GetCorporationMembers`, `GetCorporationStructures` |
| 宇宙/星图/物品类型 | `universe.go` | `GetUniverseRegions`, `GetUniverseSystem`, `GetUniverseType`, `ResolveNamesToIDs` |
| 市场订单/历史价 | `market.go` | `GetMarketOrders`, `GetMarketHistory`, `GetMarketPrices` |
| 钱包余额/流水 | `wallet.go` | `GetCharacterWalletBalance`, `GetCharacterWalletJournal` |
| 资产 | `assets.go` | `GetCharacterAssets`, `GetCharacterAssetNames` |
| 合同 | `contracts.go` | `GetCharacterContracts`, `GetPublicContracts` |
| 舰队管理 | `fleets.go` | `GetFleet`, `MoveFleetMember` |
| 邮件 | `mail.go` | `GetCharacterMails`, `SendCharacterMail` |
| 合同/物品/星系等全部 32 个模块 | 见 references/modules.md | — |

## Before / After 示例

用 goeve 替代手写 HTTP：

```go
// Before: 手写请求，手动处理鉴权/解码/错误
req, _ := http.NewRequest("GET", "https://ali-esi.evepc.163.com/latest/characters/95234356/", nil)
req.Header.Set("Authorization", "Bearer "+token)
resp, _ := http.DefaultClient.Do(req)
// ... 手动 json.NewDecoder、手工判断 4xx ...

// After: goeve
client := goeve.NewClient(goeve.WithTokenSource(ts))
character, err := client.GetCharacter(ctx, 95234356)
if err != nil {
    var apiErr *goeve.APIError
    if errors.As(err, &apiErr) { /* apiErr.StatusCode, apiErr.Message ... */ }
}
```

典型调用（注意参数固定顺序：ctx → token（仅授权接口）→ 路径/业务参数 → 可选参数）：

```go
// 受保护接口：token 紧跟 ctx，空串 = 使用客户端 TokenSource
balance, err := client.GetCharacterWalletBalance(ctx, "", characterID)

// 市场订单（公开接口无 token）：regionID、orderType 是必需参数，page、typeID 依次追加（0 = 不传），不传 ETag
orders, err := client.GetMarketOrders(ctx, 10000002, "sell", 1, 0)

// 路线规划（公开）：destination、origin 是必需参数（注意顺序），avoid/connections/flag 可选
route, err := client.GetRoute(ctx, 30002187, 30000142, []int32{30000144}, nil, "secure")

// 带 ETag 的条件请求：把 ETag 作为最后一个变长参数传入
region, err := client.GetUniverseRegion(ctx, regionID, etag)

// 简体中文返回（客户端级默认值）
client := goeve.NewClient(goeve.WithLanguage(goeve.LanguageChinese))
```

## 认证 / NetEase EVE SSO

SSO 地址 `https://login.evepc.163.com`；默认 client_id 是官方 ESI 应用（`goeve.DefaultClientID`）；
redirect_uri 必须为 `goeve.DefaultRedirectURI`；**一次最多申请 4 个 scope（空格分隔）**。

授权码流程（推荐，可拿到 refresh_token 永久续期）：

```go
// Step 1: 生成授权链接，用户在浏览器打开、登录并选择角色
url := goeve.BuildAuthorizeURL(goeve.AuthorizeConfig{
    ResponseType: goeve.ResponseTypeCode,
    State:        "random-state",
    Scope:        "esi-wallet.read_character_wallet.v1 esi-assets.read_assets.v1",
})

// Step 2: 粘贴回跳地址，换取令牌（code 约 10 分钟有效、只能用一次）
code, _ := goeve.ParseAuthorizationCode(redirectURL)
token, err := goeve.GetTokenFromCode(ctx, code, goeve.AuthorizeConfig{})

// Step 3: 挂 TokenSource，access_token 过期自动用 refresh_token 续期
client := goeve.NewClient(goeve.WithTokenSource(goeve.NewTokenSource(token)))

// Step 4: 调用受保护接口：ctx 后的 token 传 "" 走客户端 TokenSource（也可显式传 token.AccessToken）
balance, err := client.GetCharacterWalletBalance(ctx, "", characterID)
```

隐式流程（仅 20 分钟 access_token，无 refresh_token）：
`ResponseType: goeve.ResponseTypeToken` + `goeve.ParseImplicitRedirect(redirectURL)`，
然后同样用 `goeve.NewTokenSource(token)` 挂载。

刷新：`token, err := goeve.RefreshAccessToken(ctx, refreshToken)`。

## 缓存、分页与语言

| 场景 | 做法 |
|---|---|
| ETag 协商 | 把 ETag 作为 GET 方法末尾的变长参数传入（`client.GetUniverseCategories(ctx, etag)`）；服务器返回 304 时**结果为 nil 且 err 为 nil**（表示数据未变化），这是正常路径不是错误；不需要协商时省略该参数 |
| 分页 | `page` 参数从 1 开始递增直到返回空切片（`X-Pages` 响应头不对外暴露，0 = 默认第 1 页）；`GetWalletTransactions` 等特殊接口没有 `page`，用 `fromID` 回溯更早记录 |
| 简体中文 | 客户端级 `goeve.WithLanguage(goeve.LanguageChinese)`（对所有请求生效，无按请求覆盖选项） |
| 数据源 | 网易服只有 tranquility，`datasource` 一般不用传 |

## 错误处理

所有 ESI 错误（400/401/403/420/500/503/504）统一解码为 `*goeve.APIError`：

```go
_, err := client.GetCharacter(ctx, 1)
var apiErr *goeve.APIError
if errors.As(err, &apiErr) && apiErr.StatusCode == 404 { /* 不存在 */ }
// 403 时 apiErr.SSOStatus 附带 SSO 状态码；504 时 apiErr.Timeout 为允许秒数
```

## 注意事项（易踩坑）

1. **refresh_token 等同密码**：泄露等于账号失陷，必须安全存储；用户可在官网"ESI解除授权"解除。
2. **scope 一次最多 4 个**，更多权限需分多次授权或多次跳转。
3. **不要把 304 当错误**：nil + nil 就是"数据没变化"。
4. **必须传 ctx**：所有方法第一个参数是 `context.Context`；需授权接口的 `token` 固定为 ctx 之后第一位，不要再把 token 放到路径参数后面。
5. **可选参数不能省略**：`token`/`page`/业务过滤参数是固定位置参数，签名里有就要传，`token ""`、`page 0`、`nil` 切片表示"不传"；只有 GET 方法末尾的 ETag 变长参数可以整个省略；不写任何请求选项或参数结构体。
6. **每个方法的可选参数集合不同**：以 `goeve_iface.go` 签名/godoc 为准（如 `GetCharacterMails` 没有 `page`，翻页用 `lastMailID`）。
7. 方法所需 scope 标注在该方法的 godoc `Scopes:` 行，公开接口标注 `none (public endpoint)`。

## 验证

改动后运行（与仓库测试一致）：

```bash
gofmt -l . && go vet ./... && go test ./...   # 测试含对公开接口的真实网络调用
```

## 常见问题

### Q: 怎么找到某个游戏功能对应的调用？

先在 `references/modules.md` 按功能定位模块文件，再到 `goeve_iface.go` 或方法 godoc
查准确签名——命名已语义化（如角色邮件列表方法命名为 `GetCharacterMails`），
**不能按 operationId 机械转换推名字**。方法 godoc 标注了路由、缓存时长和所需 scope。

### Q: 分页什么时候停？

有 `page` 参数的接口：从 1 递增直到返回空切片；不要依赖 `X-Pages` 头（未暴露）。
无 `page` 的接口（如 `GetCharacterMails`、`GetWalletTransactions`）用各自的
`lastMailID` / `fromID` 参数回溯更早数据。

### Q: 403 了怎么办？

检查方法 godoc 要求的 scope 是否在授权时申请过、token 是否过期（TokenSource 会自动续期）、
角色是否已离开需要权限的军团。
