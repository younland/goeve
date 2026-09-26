# ESI 模块接口说明 / Module reference

网易 EVE Online ESI 全部 32 个模块的接口简要说明。方法在根包同名文件（如 `character.go`），
类型在 `models/` 子包同名文件。每条注明：用途、典型方法与是否需授权（scope）。
方法名已语义化；**全部参数为显式位置参数**：`ctx` 之后第一位是 **`token`（仅需授权接口，
空串 = 回退客户端 TokenSource / 匿名）**，随后是路径/业务必需参数，再按固定顺序追加可选
参数（`page`、业务过滤参数），零值（`""`/`0`/`nil`/`false`）表示不传。表中每个方法末尾的
`ifNoneMatch` 是**变长可选参数**（`ifNoneMatch ...string`，标注"可选"），不需要 ETag
协商时直接省略，传入 ETag 时 304 返回 nil 结果 + nil error。准确签名以 `goeve_iface.go` 为准。

The 32 ESI modules at a glance. Methods live in the root file of the same name
(e.g. `character.go`), types in `models/`. Scopes are noted per module. Method names are
semantic; **all parameters are explicit positional arguments**: right after `ctx` comes
**`token` (authenticated methods only, empty = fall back to the client TokenSource /
anonymous)**, then the required path/business parameters, then optional ones in a fixed
order (`page`, business filters); zero values (`""`/`0`/`nil`/`false`) mean "omit". A
trailing `ifNoneMatch` on every method is a **variadic optional parameter**
(`ifNoneMatch ...string`, marked "可选"/optional) — omit it for no ETag negotiation; pass
an ETag to get a nil result with nil error on 304. Check `goeve_iface.go` for exact
signatures.

## Status — 服务器状态 (1)

公开。服务器在线人数、版本、启动时间。

| 方法 | 说明 |
|---|---|
| `GetServerStatus(ifNoneMatch 可选)` | 获取服务器运行时间和玩家数量 |

## Character — 角色 (14)

公开 + 授权混合。角色公开资料、通知、勋章、声望、军团履历、CSPA 费用计算。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacter(characterID, ifNoneMatch 可选)` | 角色公开信息（名称、军团、种族、血统） | 公开 |
| `GetCharacterPortrait(characterID, ifNoneMatch 可选)` | 角色头像 URL | 公开 |
| `CharacterAffiliation(body []int32)` | 批量查询角色所属军团/联盟 | 公开 |
| `GetCharacterCorporationHistory(characterID, ifNoneMatch 可选)` | 角色军团历史 | 公开 |
| `GetCharacterNotifications(token, characterID, ifNoneMatch 可选)` | 角色通知列表（攻击警报、建筑通知等） | `esi-characters.read_notifications.v1` |
| `GetCharacterContactNotifications(token, characterID, ifNoneMatch 可选)` | 联系人相关新通知 | 同上 |
| `GetCharacterStandings(token, characterID, ifNoneMatch 可选)` | NPC/军团/联盟声望 | `esi-characters.read_standings.v1` |
| `GetCharacterMedals(token, characterID, ifNoneMatch 可选)` | 已获勋章 | `esi-characters.read_medals.v1` |
| `GetCharacterCorporationRoles` / `GetCharacterCorporationTitles`(token, characterID, ifNoneMatch 可选) | 军团职务/称号 | `esi-characters.read_titles.v1` |
| `GetCharacterAgentsResearch(token, characterID, ifNoneMatch 可选)` | 代理人研究数据 | `esi-characters.read_agents_research.v1` |
| `GetCharacterBlueprints(token, characterID, page, ifNoneMatch 可选)` | 蓝图（原件/拷贝，page 分页） | `esi-characters.read_blueprints.v1` |
| `GetCharacterJumpFatigue(token, characterID, ifNoneMatch 可选)` | 跳跃疲劳 | `esi-characters.read_fatigue.v1` |
| `CalculateCharacterCspaCharge(token, characterID, body []int32)` | 计算向多个角色发邮件的 CSPA 费用 | `esi-characters.calculate_cspa.v1` |

## Corporation — 军团 (22)

公开 + 授权混合。军团资料、成员、建筑、母星、股东、部门、勋章。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCorporationInformation(corporationID, ifNoneMatch 可选)` | 军团公开信息 | 公开 |
| `GetCorporationIcon(corporationID, ifNoneMatch 可选)` | 军团图标 | 公开 |
| `GetCorporationAllianceHistory(corporationID, ifNoneMatch 可选)` | 联盟历史 | 公开 |
| `GetNpcCorporations(ifNoneMatch 可选)` | NPC 军团 ID 列表 | 公开 |
| `GetCorporationMembers` / `GetCorporationMemberLimit` / `GetCorporationMemberTitles`(token, corporationID, ifNoneMatch 可选) | 成员/上限/成员称号 | `esi-corporations.read_corporation_membership.v1` |
| `GetCorporationMemberTracking(token, corporationID, ifNoneMatch 可选)` | 成员上线/位置追踪 | `esi-corporations.track_members.v1` |
| `GetCorporationStructures(token, corporationID, page, ifNoneMatch 可选)` | 军团建筑（燃料、增强状态，page 分页） | `esi-corporations.read_structures.v1` |
| `GetCorporationStarbases(token, corporationID, page, ifNoneMatch 可选)` / `GetCorporationStarbase(token, corporationID, starbaseID int64, systemID string, ifNoneMatch 可选)` | 母星(POS)列表/详情 | `esi-corporations.read_starbases.v1` |
| `GetCorporationBlueprints(token, corporationID, page, ifNoneMatch 可选)` | 军团蓝图 | `esi-corporations.read_blueprints.v1` |
| `GetCorporationContainerLogs(token, corporationID, page, ifNoneMatch 可选)` | 机柜(ALSC)日志 | `esi-corporations.read_container_logs.v1` |
| `GetCorporationDivisions` / `GetCorporationFacilities`(token, corporationID, ifNoneMatch 可选) | 部门划分/设施 | `esi-corporations.read_divisions.v1` |
| `GetCorporationMedals` / `GetCorporationIssuedMedals`(token, corporationID, page, ifNoneMatch 可选) | 勋章/颁发记录 | `esi-corporations.read_medals.v1` |
| `GetCorporationMemberRoles` / `GetCorporationMemberRolesHistory`(token, corporationID[, page], ifNoneMatch 可选) | 成员角色/变更历史 | `esi-corporations.read_corporation_roles.v1` |
| `GetCorporationShareholders(token, corporationID, page, ifNoneMatch 可选)` | 股东列表 | `esi-wallet.read_corporation_wallets.v1` |
| `GetCorporationStandings` / `GetCorporationTitles`(token, corporationID[, page], ifNoneMatch 可选) | 声望/称号体系 | `esi-corporations.read_standings.v1` |

## Universe — 宇宙 (30)

几乎全公开。星图、物品类型体系、名称/ID 互转、星门/空间站/建筑、 jumps/kills 统计。

| 方法 | 说明 |
|---|---|
| `ResolveNamesToIDs(body []string)` / `ResolveIDsToNames(body []int32)` | 名称↔ID 批量互转（最多 100 个） |
| `GetUniverseRegions(ifNoneMatch 可选)` / `GetUniverseRegion(regionID, ifNoneMatch 可选)` | 星域列表/详情 |
| `GetConstellations(ifNoneMatch 可选)` / `GetConstellationInformation(constellationID, ifNoneMatch 可选)` | 星座列表/详情 |
| `GetUniverseSystems(ifNoneMatch 可选)` / `GetUniverseSystem(systemID, ifNoneMatch 可选)` | 星系列表/详情 |
| `GetUniverseTypes(page, ifNoneMatch 可选)` / `GetUniverseType(typeID, ifNoneMatch 可选)` | 物品类型列表/详情（舰船、装备、物资…，page 分页） |
| `GetUniverseCategories(ifNoneMatch 可选)` / `GetUniverseCategory(categoryID, ifNoneMatch 可选)` | 物品分类 |
| `GetUniverseGroups(page, ifNoneMatch 可选)` / `GetUniverseGroup(groupID, ifNoneMatch 可选)` | 物品分组（page 分页） |
| `GetUniverseMoon(moonID, ifNoneMatch 可选)` / `GetUniversePlanet(planetID, ifNoneMatch 可选)` / `GetUniverseAsteroidBelt(asteroidBeltID, ifNoneMatch 可选)` | 月球/行星/小行星带详情 |
| `GetUniverseStargate(stargateID, ifNoneMatch 可选)` / `GetUniverseStar(starID, ifNoneMatch 可选)` | 星门/恒星详情 |
| `GetUniverseStation(stationID, ifNoneMatch 可选)` / `GetUniverseStructure(token, structureID int64, ifNoneMatch 可选)` | 空间站/建筑详情 |
| `GetPublicStructures(filter, ifNoneMatch 可选)` | 所有公开建筑 ID（filter 过滤，如 `market`） |
| `GetUniverseSystemJumps(ifNoneMatch 可选)` / `GetUniverseSystemKills(ifNoneMatch 可选)` | 各星系跳跃数/击毁统计（近 1 小时） |
| `GetUniverseRaces` / `GetUniverseBloodlines` / `GetUniverseAncestries` / `GetUniverseFactions`(ifNoneMatch 可选) | 种族/血统/祖先/势力 |
| `GetUniverseGraphics(ifNoneMatch 可选)` / `GetUniverseGraphic(graphicID, ifNoneMatch 可选)` | 图形资源 |

## Market — 市场 (11)

公开行情 + 授权个人订单。区域订单、历史价格、市场参考价。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetMarketOrders(regionID, orderType, page, typeID, ifNoneMatch 可选)` | 区域市场订单（orderType: `all`/`buy`/`sell`；typeID 0 = 不过滤；page 分页） | 公开 |
| `GetMarketHistory(regionID, typeID string, ifNoneMatch 可选)` | 区域日历史行情（typeID 为 string 位置参数） | 公开 |
| `GetMarketPrices(ifNoneMatch 可选)` | 全服市场参考价 | 公开 |
| `GetMarketGroups(ifNoneMatch 可选)` / `GetMarketGroup(marketGroupID, ifNoneMatch 可选)` | 市场分组 | 公开 |
| `GetMarketTypes(regionID, page, ifNoneMatch 可选)` | 区域有行情的物品种类（page 分页） | 公开 |
| `GetStructureMarketOrders(token, structureID int64, page, ifNoneMatch 可选)` | 建筑内订单 | `esi-markets.read_structures.v1` |
| `GetCharacterMarketOrders` / `GetCharacterMarketOrderHistory`(token, characterID[, page], ifNoneMatch 可选) | 角色当前/历史订单 | `esi-markets.read_character_orders.v1` |
| `GetCorporationMarketOrders` / `GetCorporationMarketOrderHistory`(token, corporationID, page, ifNoneMatch 可选) | 军团订单 | `esi-markets.read_corporation_orders.v1` |

## Wallet — 钱包 (6)

授权。余额、钱包流水、交易记录。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterWalletBalance(token, characterID, ifNoneMatch 可选)` | 角色余额 | `esi-wallet.read_character_wallet.v1` |
| `GetCharacterWalletJournal(token, characterID, page, ifNoneMatch 可选)` | 角色钱包流水（入账/出账，page 分页） | 同上 |
| `GetWalletTransactions(token, characterID, fromID int64, ifNoneMatch 可选)` | 角色市场交易记录（无 page，用 fromID 回溯翻页） | 同上 |
| `GetCorporationWallets(token, corporationID, ifNoneMatch 可选)` | 军团各财务分部余额 | `esi-wallet.read_corporation_wallets.v1` |
| `GetCorporationWalletJournal` / `GetCorporationWalletTransactions`(token, corporationID, division[, page/fromID], ifNoneMatch 可选) | 军团分部流水/交易 | 同上 |

## Assets — 资产 (6)

授权。角色/军团资产列表、位置与名称解析。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterAssets(token, characterID, page, ifNoneMatch 可选)` | 角色全部资产（含 location_id、插槽，page 分页） | `esi-assets.read_assets.v1` |
| `GetCharacterAssetLocations(token, characterID, body []int64)` | 批量解析资产所在建筑/空间站坐标 | 同上 |
| `GetCharacterAssetNames(token, characterID, body []int64)` | 批量解析物品名称 | 同上 |
| `GetCorporationAssets` 等 3 个 | 军团同组接口 | `esi-assets.read_corporation_assets.v1` |

## Contracts — 合同 (9)

公开拍卖行 + 授权个人/军团合同。合同、出价、物品清单。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetPublicContracts(regionID, page, ifNoneMatch 可选)` | 区域公开合同（page 分页） | 公开 |
| `GetPublicContractBids` / `GetPublicContractItems`(contractID, page, ifNoneMatch 可选) | 公开合同出价/物品（page 分页） | 公开 |
| `GetCharacterContracts(token, characterID, page, ifNoneMatch 可选)` | 角色合同（page 分页） | `esi-contracts.read_character_contracts.v1` |
| `GetCharacterContractBids` / `GetCharacterContractItems`(token, characterID, contractID, ifNoneMatch 可选) | 角色合同出价/物品 | 同上 |
| `GetCorporationContracts` 等 3 个（token, corporationID 或 contractID[, page], ifNoneMatch 可选） | 军团同组接口 | `esi-contracts.read_corporation_contracts.v1` |

## Fleets — 舰队 (14)

授权。舰队信息、成员管理、联队/小队编组（不能创建/解散舰队）。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterFleet(token, characterID, ifNoneMatch 可选)` | 角色当前所在舰队 | `esi-fleets.read_fleet.v1` |
| `GetFleet(token, fleetID int64, ifNoneMatch 可选)` / `UpdateFleetSettings(token, fleetID int64, body)` | 舰队信息/修改设置 | 同上 + `esi-fleets.write_fleet.v1` |
| `GetFleetMembers` / `CreateFleetInvitation`(token, fleetID int64[, body]) | 成员列表/邀请 | 读 + `esi-fleets.write_fleet.v1` |
| `KickFleetMember` / `MoveFleetMember`(token, fleetID int64, memberID[, body]) | 踢出/移动成员 | `esi-fleets.write_fleet.v1` |
| `GetFleetWings` 及 wings/squads 增删改名（`CreateFleetWing` / `RenameFleetWing` / `DeleteFleetWing` / `CreateFleetSquad` / `RenameFleetSquad` / `DeleteFleetSquad`） | 联队/小队编组 | 读 + write |

## Mail — 邮件 (9)

授权。收发邮件、标签、邮件列表。注意：`GetCharacterMails` 无 page 参数，用 lastMailID 回溯翻页。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterMails(token, characterID, labels []int32, lastMailID, ifNoneMatch 可选)` | 邮件头列表（labels nil = 不过滤，lastMailID 0 = 从最新开始） | `esi-mail.read_mail.v1` |
| `SendCharacterMail(token, characterID, body)` | 发送邮件 | `esi-mail.send_mail.v1` |
| `GetCharacterMailLabels` / `CreateCharacterMailLabel` / `DeleteCharacterMailLabel`(token, characterID[, labelID/body]) | 标签管理 | `esi-mail.organize_mail.v1` |
| `GetCharacterMailLists(token, characterID, ifNoneMatch 可选)` | 邮件列表订阅 | `esi-mail.read_mail.v1` |
| `GetCharacterMail` / `UpdateCharacterMail` / `DeleteCharacterMail`(token, characterID, mailID[, body]) | 单封邮件读取/更新/删除 | `esi-mail.read_mail.v1` / `esi-mail.organize_mail.v1` |

## Contacts — 联系人 (9)

授权。个人/军团/联盟联系人与标签。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterContacts` / `AddCharacterContacts` / `EditCharacterContacts` / `DeleteCharacterContacts` | 角色联系人增删改查（token 紧跟 ctx；列表带 page；新增/编辑带 labelIDs/watched 可选参数） | `esi-contacts.read/write_contacts.v1` |
| `GetCharacterContactLabels(token, characterID, ifNoneMatch 可选)` | 联系人标签 | `esi-contacts.read_contacts.v1` |
| `GetCorporationContacts` 等 | 军团联系人 | `esi-contacts.read_corporation_contacts.v1` |
| `GetAllianceContacts` 等 | 联盟联系人 | `esi-contacts.read_alliance_contacts.v1` |

## Industry — 工业 (8)

授权 + 公开。工业作业、采矿账本、月球钻探、设施与成本指数。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIndustryJobs(token, characterID, includeCompleted, ifNoneMatch 可选)` | 角色工业作业（includeCompleted 含已完成） | `esi-industry.read_character_jobs.v1` |
| `GetCharacterMiningLedger(token, characterID, page, ifNoneMatch 可选)` | 角色采矿账本（page 分页） | `esi-industry.read_character_mining.v1` |
| `GetCorporationIndustryJobs(token, corporationID, page, includeCompleted, ifNoneMatch 可选)` | 军团作业 | `esi-industry.read_corporation_jobs.v1` |
| `GetCorporationMoonExtractions(token, corporationID, page, ifNoneMatch 可选)` | 月球钻探计时 | `esi-industry.read_corporation_mining.v1` |
| `GetCorporationMiningObservers` / `GetCorporationMiningObserverData`(token, corporationID[, observerID int64], page, ifNoneMatch 可选) | 采矿观测记录 | 同上 |
| `GetIndustryFacilities(ifNoneMatch 可选)` | 工业设施列表 | 公开 |
| `GetIndustrySystemCostIndices(ifNoneMatch 可选)` | 星系成本指数 | 公开 |

## Faction Warfare — 势力战争 (8)

公开 + 授权统计。排行榜、占领星系、会战。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetFactionWarfareLeaderboard` / `GetFactionWarfareCharacterLeaderboard` / `GetFactionWarfareCorporationLeaderboard`(ifNoneMatch 可选) | 势力/飞行员/军团排行榜 | 公开 |
| `GetFactionWarfareStats` / `GetFactionWarfareSystems` / `GetFactionWarfareWars`(ifNoneMatch 可选) | 总览/星系归属/交战 | 公开 |
| `GetCharacterFactionWarfareStats(token, characterID, ifNoneMatch 可选)` | 角色 FW 战绩 | `esi-characters.read_fw_stats.v1` |
| `GetCorporationFactionWarfareStats(token, corporationID, ifNoneMatch 可选)` | 军团 FW 战绩 | `esi-corporations.read_fw_stats.v1` |

## Killmails — 击杀报告 (3)

公开 + 授权。击杀/损失列表与单条击杀详情。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterKillmails(token, characterID, page, ifNoneMatch 可选)` | 角色近期击杀与损失 | `esi-killmails.read_killmails.v1` |
| `GetCorporationKillmails(token, corporationID, page, ifNoneMatch 可选)` | 军团近期击杀与损失 | 同上 |
| `GetKillmail(killmailHash string, killmailID int32, ifNoneMatch 可选)` | 单条击杀详情（hash 在前，注意顺序） | 公开 |

## Skills — 技能 (3)

授权。属性、技能列表、技能队列。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterAttributes(token, characterID, ifNoneMatch 可选)` | 基础属性 | `esi-skills.read_skills.v1` |
| `GetCharacterSkills(token, characterID, ifNoneMatch 可选)` | 已学技能及等级 | 同上 |
| `GetCharacterSkillQueue(token, characterID, ifNoneMatch 可选)` | 技能队列 | `esi-skills.read_skillqueue.v1` |

## Location — 位置 (3)

授权。实时位置、在线状态、当前舰船。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterLocation(token, characterID, ifNoneMatch 可选)` | 当前所在星系/空间站 | `esi-location.read_location.v1` |
| `GetCharacterOnline(token, characterID, ifNoneMatch 可选)` | 是否在线 | `esi-location.read_online.v1` |
| `GetCharacterShip(token, characterID, ifNoneMatch 可选)` | 当前驾驶的舰船 | `esi-location.read_ship_type.v1` |

## Bookmarks — 位标 (4)

授权。个人/军团位标与文件夹。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterBookmarks` / `GetCharacterBookmarkFolders`(token, characterID, page, ifNoneMatch 可选) | 角色位标/文件夹（page 分页） | `esi-bookmarks.read_character_bookmarks.v1` |
| `ListCorporationBookmarks` / `ListCorporationBookmarkFolders`(token, corporationID, page, ifNoneMatch 可选) | 军团位标/文件夹（page 分页） | `esi-bookmarks.read_corporation_bookmarks.v1` |

## Calendar — 日历 (4)

授权。游戏内日历事件与响应。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterCalendarEvents(token, characterID, fromEvent, ifNoneMatch 可选)` / `GetCalendarEvent(token, characterID, eventID, ifNoneMatch 可选)` | 事件列表（fromEvent 0 = 不过滤）/详情 | `esi-calendar.read_calendar_events.v1` |
| `RespondToCalendarEvent(token, characterID, eventID, body)` | 响应事件（接受/拒绝/犹豫） | `esi-calendar.respond_calendar_events.v1` |
| `GetCalendarEventAttendees(token, characterID, eventID, ifNoneMatch 可选)` | 参与者列表 | 读 |

## Clones — 克隆 (2)

授权。跳跃克隆与脑插。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterClones(token, characterID, ifNoneMatch 可选)` | 跳跃克隆及 implanted 脑插 | `esi-clones.read_clones.v1` |
| `GetCharacterImplants(token, characterID, ifNoneMatch 可选)` | 当前生效脑插 | `esi-clones.read_implants.v1` |

## Fittings — 装配 (3)

授权。舰船装配方案增删查。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterFittings(token, characterID, ifNoneMatch 可选)` | 装配列表 | `esi-fittings.read_fittings.v1` |
| `CreateCharacterFitting(token, characterID, body)` | 新建装配 | `esi-fittings.write_fittings.v1` |
| `DeleteCharacterFitting(token, characterID, fittingID)` | 删除装配 | 同上 |

## Loyalty — 忠诚点 (2)

授权 + 公开。LP 余额、LP 商店兑换项。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterLoyaltyPoints(token, characterID, ifNoneMatch 可选)` | 各势力 LP 余额 | `esi-characters.read_loyalty.v1` |
| `GetLoyaltyStoreOffers(corporationID, ifNoneMatch 可选)` | LP 商店兑换列表 | 公开 |

## Planetary Interaction — 行星开发 (4)

授权 + 公开。种菜（行星殖民地）、海关办公室。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterColonies` / `GetCharacterColonyLayout`(token, characterID[, planetID], ifNoneMatch 可选) | 殖民地列表/布局 | `esi-planets.manage_planets.v1` |
| `GetCorporationCustomsOffices(token, corporationID, page, ifNoneMatch 可选)` | 军团海关办公室 | `esi-planets.read_customs_offices.v1` |
| `GetSchematicInformation(schematicID, ifNoneMatch 可选)` | 行星示意图 | 公开 |

## Opportunities — 机遇系统 (5)

公开 + 授权。新手引导任务体系。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetOpportunityGroups` / `GetOpportunityGroup`(ifNoneMatch 可选 / groupID) | 任务组 | 公开 |
| `GetOpportunityTasks` / `GetOpportunityTask`(ifNoneMatch 可选 / taskID) | 任务 | 公开 |
| `GetCharacterOpportunities(token, characterID, ifNoneMatch 可选)` | 角色已完成任务 | `esi-opportunities.read_opportunities.v1` |

## Sovereignty — 主权 (3)

公开。主权战役、星图主权、主权建筑。

| 方法 | 说明 |
|---|---|
| `GetSovereigntyMap(ifNoneMatch 可选)` | 全服星系主权归属 |
| `GetSovereigntyCampaigns(ifNoneMatch 可选)` | 进行中的主权战役 |
| `GetSovereigntyStructures(ifNoneMatch 可选)` | TCU/IHUB 等主权建筑 |

## Wars — 战争 (3)

公开。军团战争列表与击杀。

| 方法 | 说明 |
|---|---|
| `GetWars(maxWarID, ifNoneMatch 可选)` | 战争列表（maxWarID 0 = 不过滤，翻页锚点） |
| `GetWar(warID, ifNoneMatch 可选)` | 战争详情 |
| `GetWarKillmails(warID, page, ifNoneMatch 可选)` | 战争相关击杀（page 分页） |

## Alliance — 联盟 (4)

公开。联盟信息、成员军团、图标。

| 方法 | 说明 |
|---|---|
| `GetAlliances(ifNoneMatch 可选)` | 全服联盟 ID 列表 |
| `GetAlliance(allianceID, ifNoneMatch 可选)` / `GetAllianceIcons(allianceID, ifNoneMatch 可选)` | 联盟信息/图标 |
| `GetAllianceCorporations(allianceID, ifNoneMatch 可选)` | 成员军团列表 |

## Dogma — 教条属性 (5)

公开。物品属性/效果体系（装备数值来源）。

| 方法 | 说明 |
|---|---|
| `GetDogmaAttributes(ifNoneMatch 可选)` / `GetDogmaAttribute(attributeID, ifNoneMatch 可选)` | 属性列表/详情 |
| `GetDogmaEffects(ifNoneMatch 可选)` / `GetDogmaEffect(effectID, ifNoneMatch 可选)` | 效果列表/详情 |
| `GetDogmaDynamicItem(itemID int64, typeID int32, ifNoneMatch 可选)` | 动态物品（改装件）属性（itemID 在前，注意顺序） |

## User Interface — 游戏内 UI 联动 (5)

授权。让游戏客户端打开窗口/设置自动导航（需客户端在线）。

| 方法 | 说明 | Scope |
|---|---|---|
| `SetAutopilotWaypoint(token, addToBeginning, clearOtherWaypoints bool, destinationID int64)` | 添加自动导航路径点 | `esi-ui.write_waypoint.v1` |
| `OpenMarketDetails` / `OpenInformationWindow` / `OpenContractWindow`(token, typeID/targetID/contractID string) | 打开市场/信息/合同窗口 | `esi-ui.open_window.v1` |
| `OpenNewMailWindow(token, body)` | 打开新邮件窗口 | 同上 |

## Search — 搜索 (1)

授权。按关键词搜索角色/军团/联盟/物品等。

| 方法 | 说明 | Scope |
|---|---|---|
| `SearchEntities(token, characterID, categories []string, search, strict, ifNoneMatch 可选)` | 按类别搜索（categories 如 `[]string{"character","corporation"}`，strict false = 不启用严格匹配） | `esi-search.search_structures.v1` 等按类别 |

## Incursions — 入侵 (1)

公开。当前入侵活动及受袭星系。

| 方法 | 说明 |
|---|---|
| `GetIncursions(ifNoneMatch 可选)` | 入侵列表（含 infested solar systems） |

## Insurance — 保险 (1)

公开。舰船保险等级与赔付。

| 方法 | 说明 |
|---|---|
| `GetPrices(ifNoneMatch 可选)` | 各船型保险等级价格 |

## Routes — 路线规划 (1)

公开。两点间导航路线（避开/经过指定星系）。

| 方法 | 说明 |
|---|---|
| `GetRoute(destination, origin int32, avoid []int32, connections [][]int32, flag, ifNoneMatch 可选)` | 计算 origin→destination 星系路径（destination 在前，注意顺序；avoid/connections/flag 可 nil/""） |
