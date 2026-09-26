# ESI 模块接口说明 / Module reference

网易 EVE Online ESI 全部 32 个模块的接口简要说明。方法在根包同名文件（如 `character.go`），
类型在 `models/` 子包同名文件。每条注明：用途、典型方法与是否需授权（scope）。

The 32 ESI modules at a glance. Methods live in the root file of the same name
(e.g. `character.go`), types in `models/`. Scopes are noted per module.

## Status — 服务器状态 (1)

公开。服务器在线人数、版本、启动时间。

| 方法 | 说明 |
|---|---|
| `GetStatus` | 获取服务器运行时间和玩家数量 |

## Character — 角色 (14)

公开 + 授权混合。角色公开资料、通知、勋章、声望、军团履历、CSPA 费用计算。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterId` | 角色公开信息（名称、军团、种族、血统） | 公开 |
| `GetCharacterIdPortrait` | 角色头像 URL | 公开 |
| `PostAffiliation` | 批量查询角色所属军团/联盟 | 公开 |
| `GetCharacterIdCorporationhistory` | 角色军团历史 | 公开 |
| `GetCharacterIdNotifications` | 角色通知列表（攻击警报、建筑通知等） | `esi-characters.read_notifications.v1` |
| `GetCharacterIdNotificationsContacts` | 联系人相关新通知 | 同上 |
| `GetCharacterIdStandings` | NPC/军团/联盟声望 | `esi-characters.read_standings.v1` |
| `GetCharacterIdMedals` | 已获勋章 | `esi-characters.read_medals.v1` |
| `GetCharacterIdRoles` / `GetCharacterIdTitles` | 军团职务/称号 | `esi-characters.read_titles.v1` |
| `GetCharacterIdAgentsResearch` | 代理人研究数据 | `esi-characters.read_agents_research.v1` |
| `GetCharacterIdBlueprints` | 蓝图（原件/拷贝） | `esi-characters.read_blueprints.v1` |
| `GetCharacterIdFatigue` | 跳跃疲劳 | `esi-characters.read_fatigue.v1` |
| `PostCharacterIdCspa` | 计算向多个角色发邮件的 CSPA 费用 | `esi-characters.calculate_cspa.v1` |

## Corporation — 军团 (22)

公开 + 授权混合。军团资料、成员、建筑、母星、股东、部门、勋章。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCorporationId` | 军团公开信息 | 公开 |
| `GetCorporationIdIcons` | 军团图标 | 公开 |
| `GetCorporationIdAlliancehistory` | 联盟历史 | 公开 |
| `GetNpccorps` | NPC 军团 ID 列表 | 公开 |
| `GetCorporationIdMembers` / `MembersLimit` / `MembersTitles` | 成员/上限/成员称号 | `esi-corporations.read_corporation_membership.v1` |
| `GetCorporationIdMembertracking` | 成员上线/位置追踪 | `esi-corporations.track_members.v1` |
| `GetCorporationIdStructures` | 军团建筑（燃料、增强状态） | `esi-corporations.read_structures.v1` |
| `GetCorporationIdStarbases` / `StarbasesStarbaseId` | 母星(POS)列表/详情 | `esi-corporations.read_starbases.v1` |
| `GetCorporationIdBlueprints` | 军团蓝图 | `esi-corporations.read_blueprints.v1` |
| `GetCorporationIdContainersLogs` | 机柜(ALSC)日志 | `esi-corporations.read_container_logs.v1` |
| `GetCorporationIdDivisions` / `Facilities` | 部门划分/设施 | `esi-corporations.read_divisions.v1` |
| `GetCorporationIdMedals` / `MedalsIssued` | 勋章/颁发记录 | `esi-corporations.read_medals.v1` |
| `GetCorporationIdRoles` / `RolesHistory` | 成员角色/变更历史 | `esi-corporations.read_corporation_roles.v1` |
| `GetCorporationIdShareholders` | 股东列表 | `esi-wallet.read_corporation_wallets.v1` |
| `GetCorporationIdStandings` / `Titles` | 声望/称号体系 | `esi-corporations.read_standings.v1` |

## Universe — 宇宙 (30)

几乎全公开。星图、物品类型体系、名称/ID 互转、星门/空间站/建筑、 jumps/kills 统计。

| 方法 | 说明 |
|---|---|
| `PostIds` / `PostNames` | 名称↔ID 批量互转（最多 100 个） |
| `GetRegions` / `GetRegionsRegionId` | 星域列表/详情 |
| `GetConstellations` / `GetConstellationsConstellationId` | 星座列表/详情 |
| `GetSystems` / `GetSystemsSystemId` | 星系列表/详情 |
| `GetTypes` / `GetTypesTypeId` | 物品类型列表/详情（舰船、装备、物资…） |
| `GetCategories` / `GetCategoriesCategoryId` | 物品分类 |
| `GetGroups` / `GetGroupsGroupId` | 物品分组 |
| `GetMoonsMoonId` / `GetPlanetsPlanetId` | 月球/行星详情 |
| `GetStargatesStargateId` / `GetStarsStarId` | 星门/恒星详情 |
| `GetStationsStationId` / `GetStructuresStructureId` | 空间站/建筑详情 |
| `GetStructures` | 所有公开建筑 ID |
| `GetSystemJumps` / `GetSystemKills` | 各星系跳跃数/击毁统计（近 1 小时） |
| `GetRaces` / `GetBloodlines` / `GetAncestries` / `GetFactions` | 种族/血统/祖先/势力 |
| `GetGraphics` / `GetGraphicsGraphicId` | 图形资源 |

## Market — 市场 (11)

公开行情 + 授权个人订单。区域订单、历史价格、市场参考价。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetMarketsRegionIdOrders` | 区域市场订单（卖单/买单） | 公开 |
| `GetMarketsRegionIdHistory` | 区域日历史行情 | 公开 |
| `GetMarketsPrices` | 全服市场参考价 | 公开 |
| `GetMarketsGroups` / `GetMarketsGroupsMarketGroupId` | 市场分组 | 公开 |
| `GetMarketsRegionIdTypes` | 区域有行情的物品种类 | 公开 |
| `GetMarketsStructuresStructureId` | 建筑内订单 | `esi-markets.read_structures.v1` |
| `GetCharactersCharacterIdOrders` / `OrdersHistory` | 角色当前/历史订单 | `esi-markets.read_character_orders.v1` |
| `GetCorporationsCorporationIdOrders` / `OrdersHistory` | 军团订单 | `esi-markets.read_corporation_orders.v1` |

## Wallet — 钱包 (6)

授权。余额、钱包流水、交易记录。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdWallet` | 角色余额 | `esi-wallet.read_character_wallet.v1` |
| `GetCharactersCharacterIdWalletJournal` | 角色钱包流水（入账/出账） | 同上 |
| `GetCharactersCharacterIdWalletTransactions` | 角色市场交易记录 | 同上 |
| `GetCorporationsCorporationIdWallets` | 军团各财务分部余额 | `esi-wallet.read_corporation_wallets.v1` |
| `GetCorporationsCorporationIdWalletsDivisionJournal` / `...Transactions` | 军团分部流水/交易 | 同上 |

## Assets — 资产 (6)

授权。角色/军团资产列表、位置与名称解析。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdAssets` | 角色全部资产（含 location_id、插槽） | `esi-assets.read_assets.v1` |
| `PostCharactersCharacterIdAssetsLocations` | 批量解析资产所在建筑/空间站坐标 | 同上 |
| `PostCharactersCharacterIdAssetsNames` | 批量解析物品名称 | 同上 |
| `GetCorporationsCorporationIdAssets` 等 3 个 | 军团同组接口 | `esi-assets.read_corporation_assets.v1` |

## Contracts — 合同 (9)

公开拍卖行 + 授权个人/军团合同。合同、出价、物品清单。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetContractsPublicRegionId` | 区域公开合同 | 公开 |
| `GetContractsPublicBidsContractId` / `PublicItemsContractId` | 公开合同出价/物品 | 公开 |
| `GetCharactersCharacterIdContracts` | 角色合同 | `esi-contracts.read_character_contracts.v1` |
| `GetCharactersCharacterIdContractsContractIdBids` / `Items` | 角色合同出价/物品 | 同上 |
| `GetCorporationsCorporationIdContracts` 等 3 个 | 军团同组接口 | `esi-contracts.read_corporation_contracts.v1` |

## Fleets — 舰队 (14)

授权。舰队信息、成员管理、联队/小队编组（不能创建/解散舰队）。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdFleet` | 角色当前所在舰队 | `esi-fleets.read_fleet.v1` |
| `GetFleetId` / `PutFleetId` | 舰队信息/修改设置 | 同上 + `esi-fleets.write_fleet.v1` |
| `GetFleetIdMembers` / `PostFleetIdMembers` | 成员列表/邀请 | 读 + `esi-fleets.write_fleet.v1` |
| `DeleteFleetIdMembersMemberId` / `PutFleetIdMembersMemberId` | 踢出/移动成员 | `esi-fleets.write_fleet.v1` |
| `GetFleetIdWings` 及 wings/squads 增删改名 | 联队/小队编组 | 读 + write |

## Mail — 邮件 (9)

授权。收发邮件、标签、邮件列表。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdMail` | 邮件头列表（分页） | `esi-mail.read_mail.v1` |
| `PostCharactersCharacterIdMail` | 发送邮件 | `esi-mail.send_mail.v1` |
| `GetCharactersCharacterIdMailLabels` / `Post...Labels` / `Delete...LabelsLabelId` | 标签管理 | `esi-mail.organize_mail.v1` |
| `GetCharactersCharacterIdMailLists` | 邮件列表订阅 | `esi-mail.read_mail.v1` |
| `DeleteCharactersCharacterIdMailMailId` | 删除邮件 | `esi-mail.organize_mail.v1` |

## Contacts — 联系人 (9)

授权。个人/军团/联盟联系人与标签。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdContacts` / `Post` / `Put` / `Delete` | 角色联系人增删改查 | `esi-contacts.read/write_contacts.v1` |
| `GetCharactersCharacterIdContactsLabels` | 联系人标签 | `esi-contacts.read_contacts.v1` |
| `GetCorporationsCorporationIdContacts` 等 | 军团联系人 | `esi-contacts.read_corporation_contacts.v1` |
| `GetAlliancesAllianceIdContacts` 等 | 联盟联系人 | `esi-contacts.read_alliance_contacts.v1` |

## Industry — 工业 (8)

授权 + 公开。工业作业、采矿账本、月球钻探、设施与成本指数。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdIndustryJobs` | 角色工业作业 | `esi-industry.read_character_jobs.v1` |
| `GetCharactersCharacterIdMining` | 角色采矿账本 | `esi-industry.read_character_mining.v1` |
| `GetCorporationsCorporationIdIndustryJobs` | 军团作业 | `esi-industry.read_corporation_jobs.v1` |
| `GetCorporationCorporationIdMiningExtractions` | 月球钻探计时 | `esi-industry.read_corporation_mining.v1` |
| `GetCorporationCorporationIdMiningObservers(+ObserverId)` | 采矿观测记录 | 同上 |
| `GetIndustryFacilities` | 工业设施列表 | 公开 |
| `GetIndustrySystems` | 星系成本指数 | 公开 |

## Faction Warfare — 势力战争 (8)

公开 + 授权统计。排行榜、占领星系、会战。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetFwLeaderboards(+Characters/Corporations)` | 势力/飞行员/军团排行榜 | 公开 |
| `GetFwStats` / `GetFwSystems` / `GetFwWars` | 总览/星系归属/交战 | 公开 |
| `GetCharactersCharacterIdFwStats` | 角色 FW 战绩 | `esi-characters.read_fw_stats.v1` |
| `GetCorporationsCorporationIdFwStats` | 军团 FW 战绩 | `esi-corporations.read_fw_stats.v1` |

## Killmails — 击杀报告 (3)

公开 + 授权。击杀/损失列表与单条击杀详情。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdKillmailsRecent` | 角色近期击杀与损失 | `esi-killmails.read_killmails.v1` |
| `GetCorporationsCorporationIdKillmailsRecent` | 军团近期击杀与损失 | 同上 |
| `GetKillmailsKillmailIdKillmailHash` | 单条击杀详情（需 id + hash） | 公开 |

## Skills — 技能 (3)

授权。属性、技能列表、技能队列。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdAttributes` | 基础属性 | `esi-skills.read_skills.v1` |
| `GetCharacterIdSkills` | 已学技能及等级 | 同上 |
| `GetCharacterIdSkillqueue` | 技能队列 | `esi-skills.read_skillqueue.v1` |

## Location — 位置 (3)

授权。实时位置、在线状态、当前舰船。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdLocation` | 当前所在星系/空间站 | `esi-location.read_location.v1` |
| `GetCharacterIdOnline` | 是否在线 | `esi-location.read_online.v1` |
| `GetCharacterIdShip` | 当前驾驶的舰船 | `esi-location.read_ship_type.v1` |

## Bookmarks — 位标 (4)

授权。个人/军团位标与文件夹。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdBookmarks` / `...Folders` | 角色位标/文件夹 | `esi-bookmarks.read_character_bookmarks.v1` |
| `GetCorporationsCorporationIdBookmarks` / `...Folders` | 军团位标/文件夹 | `esi-bookmarks.read_corporation_bookmarks.v1` |

## Calendar — 日历 (4)

授权。游戏内日历事件与响应。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdCalendar` / `GetCharacterIdCalendarEventId` | 事件列表/详情 | `esi-calendar.read_calendar_events.v1` |
| `PutCharacterIdCalendarEventId` | 响应事件（接受/拒绝/犹豫） | `esi-calendar.respond_calendar_events.v1` |
| `GetCharacterIdCalendarEventIdAttendees` | 参与者列表 | 读 |

## Clones — 克隆 (2)

授权。跳跃克隆与脑插。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdClones` | 跳跃克隆及 implanted 脑插 | `esi-clones.read_clones.v1` |
| `GetCharacterIdImplants` | 当前生效脑插 | `esi-clones.read_implants.v1` |

## Fittings — 装配 (3)

授权。舰船装配方案增删查。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdFittings` | 装配列表 | `esi-fittings.read_fittings.v1` |
| `PostCharacterIdFittings` | 新建装配 | `esi-fittings.write_fittings.v1` |
| `DeleteCharacterIdFittingsFittingId` | 删除装配 | 同上 |

## Loyalty — 忠诚点 (2)

授权 + 公开。LP 余额、LP 商店兑换项。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharactersCharacterIdLoyaltyPoints` | 各势力 LP 余额 | `esi-characters.read_loyalty.v1` |
| `GetLoyaltyStoresCorporationIdOffers` | LP 商店兑换列表 | 公开 |

## Planetary Interaction — 行星开发 (4)

授权 + 公开。种菜（行星殖民地）、海关办公室。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdPlanets` / `GetCharacterIdPlanetsPlanetId` | 殖民地列表/布局 | `esi-planets.manage_planets.v1` |
| `GetCorporationsCorporationIdCustomsOffices` | 军团海关办公室 | `esi-planets.read_customs_offices.v1` |
| `GetSchematicsSchematicId` | 行星示意图 | 公开 |

## Opportunities — 机遇系统 (5)

公开 + 授权。新手引导任务体系。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetOpportunitiesGroups(+GroupId)` / `GetOpportunitiesTasks(+TaskId)` | 任务组/任务 | 公开 |
| `GetCharacterIdOpportunities` | 角色已完成任务 | `esi-opportunities.read_opportunities.v1` |

## Sovereignty — 主权 (3)

公开。主权战役、星图主权、主权建筑。

| 方法 | 说明 |
|---|---|
| `GetSovereigntyMap` | 全服星系主权归属 |
| `GetSovereigntyCampaigns` | 进行中的主权战役 |
| `GetSovereigntyStructures` | TCU/IHUB 等主权建筑 |

## Wars — 战争 (3)

公开。军团战争列表与击杀。

| 方法 | 说明 |
|---|---|
| `GetWars` / `GetWarId` | 战争列表/详情 |
| `GetWarIdKillmails` | 战争相关击杀 |

## Alliance — 联盟 (4)

公开。联盟信息、成员军团、图标。

| 方法 | 说明 |
|---|---|
| `GetAlliances` | 全服联盟 ID 列表 |
| `GetAllianceId` / `GetAllianceIdIcons` | 联盟信息/图标 |
| `GetAllianceIdCorporations` | 成员军团列表 |

## Dogma — 教条属性 (5)

公开。物品属性/效果体系（装备数值来源）。

| 方法 | 说明 |
|---|---|
| `GetAttributes` / `GetAttributesAttributeId` | 属性列表/详情 |
| `GetEffects` / `GetEffectsEffectId` | 效果列表/详情 |
| `GetDynamicItemsTypeIdItemId` | 动态物品（改装件）属性 |

## User Interface — 游戏内 UI 联动 (5)

授权。让游戏客户端打开窗口/设置自动导航（需客户端在线）。

| 方法 | 说明 | Scope |
|---|---|---|
| `PostUiAutopilotWaypoint` | 添加自动导航路径点 | `esi-ui.write_waypoint.v1` |
| `PostUiOpenwindowMarketdetails` / `Information` / `Contract` / `Newmail` | 打开市场/信息/合同/邮件窗口 | `esi-ui.open_window.v1` |

## Search — 搜索 (1)

授权。按关键词搜索角色/军团/联盟/物品等。

| 方法 | 说明 | Scope |
|---|---|---|
| `GetCharacterIdSearch` | 指定类别搜索（categories 参数） | `esi-search.search_structures.v1` 等按类别 |

## Incursions — 入侵 (1)

公开。当前入侵活动及受袭星系。

| 方法 | 说明 |
|---|---|
| `GetIncursions` | 入侵列表（含 infested solar systems） |

## Insurance — 保险 (1)

公开。舰船保险等级与赔付。

| 方法 | 说明 |
|---|---|
| `GetInsurancePrices` | 各船型保险等级价格 |

## Routes — 路线规划 (1)

公开。两点间导航路线（避开/经过指定星系）。

| 方法 | 说明 |
|---|---|
| `GetRouteOriginDestination` | 计算 origin→destination 的星系路径 |
