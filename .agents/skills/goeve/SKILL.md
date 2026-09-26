---
name: goeve
description: >
  使用 goeve 客户端库（github.com/younland/goeve）调用网易 EVE Online ESI API 时触发。
  适用于：编写调用 ESI 接口的 Go 代码、查找某个游戏功能（市场、钱包、舰队、合同、资产等）
  对应的客户端方法、对接网易 EVE SSO OAuth2 授权、处理 ETag/分页/错误。
  Use the goeve Go client library for the NetEase EVE Online ESI API — endpoint lookup,
  SSO OAuth2 auth, functional-option/pagination/error conventions.
metadata:
  author: younland
  version: "3.0"
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
| 创建客户端 | `client := goeve.NewClient(goeve.WithTimeout(30*time.Second), goeve.WithDebug(true))`；客户端选项还有 `WithBaseURL` / `WithHTTPClient` / `WithTokenSource` / `WithLanguage`（默认 Accept-Language）；**需要授权的接口在调用时用请求选项 `goeve.WithAuthToken(token)` 传入令牌**（仅本次请求生效，覆盖 TokenSource） |
| 返回值 | 只有 `(result, error)`，不返回 `*resty.Response`。对象返回指针 `*models.Xxx`，数组返回切片 `[]models.Xxx`，204 响应只返回 `error` |
| 参数 | 第一个参数永远是 `context.Context`；**路径与业务必需参数是位置参数**（如 `characterID`、`regionID` + `orderType`）；**所有可选查询/Header 参数统一为 `...goeve.RequestOption`**，不再有每方法独立的参数结构体 |
| 命名 | 方法名已语义化（如 `get_characters_character_id` → `GetCharacter`，`get_markets_region_id_orders` → `GetMarketOrders`，`get_characters_character_id_wallet` → `GetCharacterWalletBalance`），**不要按 operationId 机械推导**，以 `goeve_iface.go` 与方法 godoc 为准 |
| 类型 | ID 类型以 swagger 为准：character/alliance/corporation/region 是 `int32`，structure/fleet/killmail 相关多为 `int64`，不要凭习惯强转 |

全部 `RequestOption`（可选项按方法而异，传前看 godoc）：

| 选项 | 用途 |
|---|---|
| `goeve.WithPage(page int32)` | 分页页码（邮件、资产、合同、订单等） |
| `goeve.WithIfNoneMatch(etag string)` | ETag 协商（If-None-Match 头） |
| `goeve.WithAcceptLanguage(lang string)` | 单请求覆盖 Accept-Language（优先级高于客户端级 `WithLanguage`） |
| `goeve.WithFromEvent(id int32)` | 日历：只返回晚于指定事件 ID 的事件 |
| `goeve.WithLabelIDs([]int32)` | 联系人：按标签 ID 过滤（重复查询参数） |
| `goeve.WithWatched(watched bool)` | 联系人：新增/编辑时的 watched 标志 |
| `goeve.WithIncludeCompleted(completed bool)` | 工业作业：包含已完成作业 |
| `goeve.WithLabels([]int32)` | 邮件：按标签过滤 |
| `goeve.WithLastMailID(id int32)` | 邮件：翻页锚点（返回更早的邮件） |
| `goeve.WithStrict(strict bool)` | 名称解析（ResolveNamesToIDs）：严格匹配 |
| `goeve.WithFromID(id int64)` | 钱包流水/交易：回溯翻页（返回更早的记录） |
| `goeve.WithAvoid([]int32)` | 路线规划：避开指定星系 |
| `goeve.WithConnections([][]int32)` | 路线规划：必须经过的星系组 |
| `goeve.WithFlag(flag string)` | 路线规划：路径偏好 flag（如 shortest/secure/insecure） |
| `goeve.WithFilter(filter string)` | 公开建筑列表过滤（market 等） |
| `goeve.WithMaxWarID(id int32)` | 战争列表翻页 |

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

典型调用（注意必需参数为位置参数，可选参数为函数选项）：

```go
// 市场订单：regionID、orderType 是必需参数（位置传参），page 是选项
orders, err := client.GetMarketOrders(ctx, 10000002, "sell", goeve.WithPage(1))

// 路线规划：destination、origin 是位置参数（注意顺序），避开/经过星系是选项
route, err := client.GetRoute(ctx, 30002187, 30000142,
    goeve.WithAvoid([]int32{30000144}),
    goeve.WithFlag("secure"),
)

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
balance, err := client.GetCharacterWalletBalance(ctx, characterID)
```

隐式流程（仅 20 分钟 access_token，无 refresh_token）：
`ResponseType: goeve.ResponseTypeToken` + `goeve.ParseImplicitRedirect(redirectURL)`。

刷新：`token, err := goeve.RefreshAccessToken(ctx, refreshToken)`。

## 缓存、分页与语言

| 场景 | 做法 |
|---|---|
| ETag 协商 | 传 `goeve.WithIfNoneMatch(etag)`；服务器返回 304 时**结果为 nil 且 err 为 nil**（表示数据未变化），这是正常路径不是错误 |
| 分页 | 循环递增 `goeve.WithPage(n)` 直到返回空切片（`X-Pages` 响应头不对外暴露）；钱包流水/交易用 `WithFromID` 回溯更早记录 |
| 简体中文 | 客户端级 `goeve.WithLanguage(goeve.LanguageChinese)`，或单请求 `goeve.WithAcceptLanguage(goeve.LanguageChinese)` |
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
4. **必须传 ctx**：所有方法第一个参数是 `context.Context`。
5. **必需参数是位置参数**（如 `GetMarketOrders` 的 `regionID, orderType`），可选参数才是 `RequestOption`；不要再构造任何参数结构体（该机制已移除）。
6. 方法所需 scope 标注在该方法的 godoc `Scopes:` 行，公开接口标注 `none (public endpoint)`。

## 验证

改动后运行（与仓库测试一致）：

```bash
gofmt -l . && go vet ./... && go test ./...   # 测试含对公开接口的真实网络调用
```

## 常见问题

### Q: 怎么找到某个游戏功能对应的调用？

先在 `references/modules.md` 按功能定位模块文件，再到 `goeve_iface.go` 或方法 godoc
查准确方法名——命名已语义化（如角色邮件列表方法命名为 `GetCharacterMails`），
**不能按 operationId 机械转换推名字**。方法 godoc 标注了路由、缓存时长和所需 scope。

### Q: 分页什么时候停？

递增 `WithPage(n)` 直到返回空切片；不要依赖 `X-Pages` 头（未暴露）。钱包流水/市场交易
等特殊接口用 `WithFromID` 取更早记录。

### Q: 403 了怎么办？

检查方法 godoc 要求的 scope 是否在授权时申请过、token 是否过期（TokenSource 会自动续期）、
角色是否已离开需要权限的军团。
