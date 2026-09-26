package goeve

import (
	"context"
	"github.com/younland/goeve/models"
)

// ClientIface is the interface of the ESI API client, grouped by module.
// ClientIface 是 ESI API 客户端的接口，按模块分组。
type ClientIface interface {

	// ----- Module: Alliances / 模块: Alliances -----
	GetAllianceId(ctx context.Context, allianceId int32, params *models.GetAllianceIdParams) (*models.GetAlliancesAllianceId, error)
	GetAllianceIdCorporations(ctx context.Context, allianceId int32, params *models.GetAllianceIdCorporationsParams) ([]int32, error)
	GetAllianceIdIcons(ctx context.Context, allianceId int32, params *models.GetAllianceIdIconsParams) (*models.GetAlliancesAllianceIdIcons, error)
	GetAlliances(ctx context.Context, params *models.GetAlliancesParams) ([]int32, error)

	// ----- Module: Contacts / 模块: Contacts -----
	DeleteCharactersCharacterIdContacts(ctx context.Context, characterId int32, params *models.DeleteCharactersCharacterIdContactsParams) error
	GetAlliancesAllianceIdContacts(ctx context.Context, allianceId int32, params *models.GetAlliancesAllianceIdContactsParams) ([]models.GetAlliancesAllianceIdContacts, error)
	GetAlliancesAllianceIdContactsLabels(ctx context.Context, allianceId int32, params *models.GetAlliancesAllianceIdContactsLabelsParams) ([]models.GetAlliancesAllianceIdContactsLabels, error)
	GetCharactersCharacterIdContacts(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdContactsParams) ([]models.GetCharactersCharacterIdContacts, error)
	GetCharactersCharacterIdContactsLabels(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdContactsLabelsParams) ([]models.GetCharactersCharacterIdContactsLabels, error)
	GetCorporationsCorporationIdContacts(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdContactsParams) ([]models.GetCorporationsCorporationIdContacts, error)
	GetCorporationsCorporationIdContactsLabels(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdContactsLabelsParams) ([]models.GetCorporationsCorporationIdContactsLabels, error)
	PostCharactersCharacterIdContacts(ctx context.Context, characterId int32, body []int32, params *models.PostCharactersCharacterIdContactsParams) ([]int32, error)
	PutCharactersCharacterIdContacts(ctx context.Context, characterId int32, body []int32, params *models.PutCharactersCharacterIdContactsParams) error

	// ----- Module: Characters / 模块: Characters -----
	GetCharacterId(ctx context.Context, characterId int32, params *models.GetCharacterIdParams) (*models.GetCharactersCharacterId, error)
	GetCharacterIdAgentsResearch(ctx context.Context, characterId int32, params *models.GetCharacterIdAgentsResearchParams) ([]models.GetCharactersCharacterIdAgentsResearch, error)
	GetCharacterIdBlueprints(ctx context.Context, characterId int32, params *models.GetCharacterIdBlueprintsParams) ([]models.GetCharactersCharacterIdBlueprints, error)
	GetCharacterIdCorporationhistory(ctx context.Context, characterId int32, params *models.GetCharacterIdCorporationhistoryParams) ([]models.GetCharactersCharacterIdCorporationhistory, error)
	GetCharacterIdFatigue(ctx context.Context, characterId int32, params *models.GetCharacterIdFatigueParams) (*models.GetCharactersCharacterIdFatigue, error)
	GetCharacterIdMedals(ctx context.Context, characterId int32, params *models.GetCharacterIdMedalsParams) ([]models.GetCharactersCharacterIdMedals, error)
	GetCharacterIdNotifications(ctx context.Context, characterId int32, params *models.GetCharacterIdNotificationsParams) ([]models.GetCharactersCharacterIdNotifications, error)
	GetCharacterIdNotificationsContacts(ctx context.Context, characterId int32, params *models.GetCharacterIdNotificationsContactsParams) ([]models.GetCharactersCharacterIdNotificationsContacts, error)
	GetCharacterIdPortrait(ctx context.Context, characterId int32, params *models.GetCharacterIdPortraitParams) (*models.GetCharactersCharacterIdPortrait, error)
	GetCharacterIdRoles(ctx context.Context, characterId int32, params *models.GetCharacterIdRolesParams) (*models.GetCharactersCharacterIdRoles, error)
	GetCharacterIdStandings(ctx context.Context, characterId int32, params *models.GetCharacterIdStandingsParams) ([]models.GetCharactersCharacterIdStandings, error)
	GetCharacterIdTitles(ctx context.Context, characterId int32, params *models.GetCharacterIdTitlesParams) ([]models.GetCharactersCharacterIdTitles, error)
	PostAffiliation(ctx context.Context, body []int32, params *models.PostAffiliationParams) ([]models.PostCharactersAffiliation, error)
	PostCharacterIdCspa(ctx context.Context, characterId int32, body []int32, params *models.PostCharacterIdCspaParams) (float64, error)

	// ----- Module: Assets / 模块: Assets -----
	GetCharactersCharacterIdAssets(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdAssetsParams) ([]models.GetCharactersCharacterIdAssets, error)
	GetCorporationsCorporationIdAssets(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdAssetsParams) ([]models.GetCorporationsCorporationIdAssets, error)
	PostCharactersCharacterIdAssetsLocations(ctx context.Context, characterId int32, body []int64, params *models.PostCharactersCharacterIdAssetsLocationsParams) ([]models.PostCharactersCharacterIdAssetsLocations, error)
	PostCharactersCharacterIdAssetsNames(ctx context.Context, characterId int32, body []int64, params *models.PostCharactersCharacterIdAssetsNamesParams) ([]models.PostCharactersCharacterIdAssetsNames, error)
	PostCorporationsCorporationIdAssetsLocations(ctx context.Context, corporationId int32, body []int64, params *models.PostCorporationsCorporationIdAssetsLocationsParams) ([]models.PostCorporationsCorporationIdAssetsLocations, error)
	PostCorporationsCorporationIdAssetsNames(ctx context.Context, corporationId int32, body []int64, params *models.PostCorporationsCorporationIdAssetsNamesParams) ([]models.PostCorporationsCorporationIdAssetsNames, error)

	// ----- Module: Skills / 模块: Skills -----
	GetCharactersCharacterIdAttributes(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdAttributesParams) (*models.GetCharactersCharacterIdAttributes, error)
	GetCharactersCharacterIdSkillqueue(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdSkillqueueParams) ([]models.GetCharactersCharacterIdSkillqueue, error)
	GetCharactersCharacterIdSkills(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdSkillsParams) (*models.GetCharactersCharacterIdSkills, error)

	// ----- Module: Bookmarks / 模块: Bookmarks -----
	GetCharactersCharacterIdBookmarks(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdBookmarksParams) ([]models.GetCharactersCharacterIdBookmarks, error)
	GetCharactersCharacterIdBookmarksFolders(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdBookmarksFoldersParams) ([]models.GetCharactersCharacterIdBookmarksFolders, error)
	GetCorporationsCorporationIdBookmarks(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdBookmarksParams) ([]models.GetCorporationsCorporationIdBookmarks, error)
	GetCorporationsCorporationIdBookmarksFolders(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdBookmarksFoldersParams) ([]models.GetCorporationsCorporationIdBookmarksFolders, error)

	// ----- Module: Calendar / 模块: Calendar -----
	GetCharactersCharacterIdCalendar(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdCalendarParams) ([]models.GetCharactersCharacterIdCalendar, error)
	GetCharactersCharacterIdCalendarEventId(ctx context.Context, characterId int32, eventId int32, params *models.GetCharactersCharacterIdCalendarEventIdParams) (*models.GetCharactersCharacterIdCalendarEventId, error)
	GetCharactersCharacterIdCalendarEventIdAttendees(ctx context.Context, characterId int32, eventId int32, params *models.GetCharactersCharacterIdCalendarEventIdAttendeesParams) ([]models.GetCharactersCharacterIdCalendarEventIdAttendees, error)
	PutCharactersCharacterIdCalendarEventId(ctx context.Context, characterId int32, eventId int32, body *models.PutCharactersCharacterIdCalendarEventIdResponse, params *models.PutCharactersCharacterIdCalendarEventIdParams) error

	// ----- Module: Clones / 模块: Clones -----
	GetCharactersCharacterIdClones(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdClonesParams) (*models.GetCharactersCharacterIdClones, error)
	GetCharactersCharacterIdImplants(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdImplantsParams) ([]int32, error)

	// ----- Module: Contracts / 模块: Contracts -----
	GetCharactersCharacterIdContracts(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdContractsParams) ([]models.GetCharactersCharacterIdContracts, error)
	GetCharactersCharacterIdContractsContractIdBids(ctx context.Context, characterId int32, contractId int32, params *models.GetCharactersCharacterIdContractsContractIdBidsParams) ([]models.GetCharactersCharacterIdContractsContractIdBids, error)
	GetCharactersCharacterIdContractsContractIdItems(ctx context.Context, characterId int32, contractId int32, params *models.GetCharactersCharacterIdContractsContractIdItemsParams) ([]models.GetCharactersCharacterIdContractsContractIdItems, error)
	GetCorporationsCorporationIdContracts(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdContractsParams) ([]models.GetCorporationsCorporationIdContracts, error)
	GetCorporationsCorporationIdContractsContractIdBids(ctx context.Context, contractId int32, corporationId int32, params *models.GetCorporationsCorporationIdContractsContractIdBidsParams) ([]models.GetCorporationsCorporationIdContractsContractIdBids, error)
	GetCorporationsCorporationIdContractsContractIdItems(ctx context.Context, contractId int32, corporationId int32, params *models.GetCorporationsCorporationIdContractsContractIdItemsParams) ([]models.GetCorporationsCorporationIdContractsContractIdItems, error)
	GetPublicBidsContractId(ctx context.Context, contractId int32, params *models.GetPublicBidsContractIdParams) ([]models.GetContractsPublicBidsContractId, error)
	GetPublicItemsContractId(ctx context.Context, contractId int32, params *models.GetPublicItemsContractIdParams) ([]models.GetContractsPublicItemsContractId, error)
	GetPublicRegionId(ctx context.Context, regionId int32, params *models.GetPublicRegionIdParams) ([]models.GetContractsPublicRegionId, error)

	// ----- Module: Fittings / 模块: Fittings -----
	DeleteCharactersCharacterIdFittingsFittingId(ctx context.Context, characterId int32, fittingId int32, params *models.DeleteCharactersCharacterIdFittingsFittingIdParams) error
	GetCharactersCharacterIdFittings(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdFittingsParams) ([]models.GetCharactersCharacterIdFittings, error)
	PostCharactersCharacterIdFittings(ctx context.Context, characterId int32, body *models.PostCharactersCharacterIdFittingsFitting, params *models.PostCharactersCharacterIdFittingsParams) (*models.PostCharactersCharacterIdFittings, error)

	// ----- Module: Fleets / 模块: Fleets -----
	DeleteFleetIdMembersMemberId(ctx context.Context, fleetId int64, memberId int32, params *models.DeleteFleetIdMembersMemberIdParams) error
	DeleteFleetIdSquadsSquadId(ctx context.Context, fleetId int64, squadId int64, params *models.DeleteFleetIdSquadsSquadIdParams) error
	DeleteFleetIdWingsWingId(ctx context.Context, fleetId int64, wingId int64, params *models.DeleteFleetIdWingsWingIdParams) error
	GetCharactersCharacterIdFleet(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdFleetParams) (*models.GetCharactersCharacterIdFleet, error)
	GetFleetId(ctx context.Context, fleetId int64, params *models.GetFleetIdParams) (*models.GetFleetsFleetId, error)
	GetFleetIdMembers(ctx context.Context, fleetId int64, params *models.GetFleetIdMembersParams) ([]models.GetFleetsFleetIdMembers, error)
	GetFleetIdWings(ctx context.Context, fleetId int64, params *models.GetFleetIdWingsParams) ([]models.GetFleetsFleetIdWings, error)
	PostFleetIdMembers(ctx context.Context, fleetId int64, body *models.PostFleetsFleetIdMembersInvitation, params *models.PostFleetIdMembersParams) error
	PostFleetIdWings(ctx context.Context, fleetId int64, params *models.PostFleetIdWingsParams) (*models.PostFleetsFleetIdWings, error)
	PostFleetIdWingsWingIdSquads(ctx context.Context, fleetId int64, wingId int64, params *models.PostFleetIdWingsWingIdSquadsParams) (*models.PostFleetsFleetIdWingsWingIdSquads, error)
	PutFleetId(ctx context.Context, fleetId int64, body *models.PutFleetsFleetIdNewSettings, params *models.PutFleetIdParams) error
	PutFleetIdMembersMemberId(ctx context.Context, fleetId int64, memberId int32, body *models.PutFleetsFleetIdMembersMemberIdMovement, params *models.PutFleetIdMembersMemberIdParams) error
	PutFleetIdSquadsSquadId(ctx context.Context, fleetId int64, squadId int64, body *models.PutFleetsFleetIdSquadsSquadIdNaming, params *models.PutFleetIdSquadsSquadIdParams) error
	PutFleetIdWingsWingId(ctx context.Context, fleetId int64, wingId int64, body *models.PutFleetsFleetIdWingsWingIdNaming, params *models.PutFleetIdWingsWingIdParams) error

	// ----- Module: FactionWarfare / 模块: FactionWarfare -----
	GetCharactersCharacterIdFwStats(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdFwStatsParams) (*models.GetCharactersCharacterIdFwStats, error)
	GetCorporationsCorporationIdFwStats(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdFwStatsParams) (*models.GetCorporationsCorporationIdFwStats, error)
	GetFwLeaderboards(ctx context.Context, params *models.GetFwLeaderboardsParams) (*models.GetFwLeaderboards, error)
	GetFwLeaderboardsCharacters(ctx context.Context, params *models.GetFwLeaderboardsCharactersParams) (*models.GetFwLeaderboardsCharacters, error)
	GetFwLeaderboardsCorporations(ctx context.Context, params *models.GetFwLeaderboardsCorporationsParams) (*models.GetFwLeaderboardsCorporations, error)
	GetFwStats(ctx context.Context, params *models.GetFwStatsParams) ([]models.GetFwStats, error)
	GetFwSystems(ctx context.Context, params *models.GetFwSystemsParams) ([]models.GetFwSystems, error)
	GetFwWars(ctx context.Context, params *models.GetFwWarsParams) ([]models.GetFwWars, error)

	// ----- Module: Industry / 模块: Industry -----
	GetCharactersCharacterIdIndustryJobs(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdIndustryJobsParams) ([]models.GetCharactersCharacterIdIndustryJobs, error)
	GetCharactersCharacterIdMining(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMiningParams) ([]models.GetCharactersCharacterIdMining, error)
	GetCorporationCorporationIdMiningExtractions(ctx context.Context, corporationId int32, params *models.GetCorporationCorporationIdMiningExtractionsParams) ([]models.GetCorporationCorporationIdMiningExtractions, error)
	GetCorporationCorporationIdMiningObservers(ctx context.Context, corporationId int32, params *models.GetCorporationCorporationIdMiningObserversParams) ([]models.GetCorporationCorporationIdMiningObservers, error)
	GetCorporationCorporationIdMiningObserversObserverId(ctx context.Context, corporationId int32, observerId int64, params *models.GetCorporationCorporationIdMiningObserversObserverIdParams) ([]models.GetCorporationCorporationIdMiningObserversObserverId, error)
	GetCorporationsCorporationIdIndustryJobs(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdIndustryJobsParams) ([]models.GetCorporationsCorporationIdIndustryJobs, error)
	GetFacilities(ctx context.Context, params *models.GetFacilitiesParams) ([]models.GetIndustryFacilities, error)
	GetIndustrySystems(ctx context.Context, params *models.GetSystemsParams) ([]models.GetIndustrySystems, error)

	// ----- Module: Killmails / 模块: Killmails -----
	GetCharactersCharacterIdKillmailsRecent(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdKillmailsRecentParams) ([]models.GetCharactersCharacterIdKillmailsRecent, error)
	GetCorporationsCorporationIdKillmailsRecent(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdKillmailsRecentParams) ([]models.GetCorporationsCorporationIdKillmailsRecent, error)
	GetKillmailIdKillmailHash(ctx context.Context, killmailHash string, killmailId int32, params *models.GetKillmailIdKillmailHashParams) (*models.GetKillmailsKillmailIdKillmailHash, error)

	// ----- Module: Location / 模块: Location -----
	GetCharactersCharacterIdLocation(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdLocationParams) (*models.GetCharactersCharacterIdLocation, error)
	GetCharactersCharacterIdOnline(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOnlineParams) (*models.GetCharactersCharacterIdOnline, error)
	GetCharactersCharacterIdShip(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdShipParams) (*models.GetCharactersCharacterIdShip, error)

	// ----- Module: Loyalty / 模块: Loyalty -----
	GetCharactersCharacterIdLoyaltyPoints(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdLoyaltyPointsParams) ([]models.GetCharactersCharacterIdLoyaltyPoints, error)
	GetStoresCorporationIdOffers(ctx context.Context, corporationId int32, params *models.GetStoresCorporationIdOffersParams) ([]models.GetLoyaltyStoresCorporationIdOffers, error)

	// ----- Module: Mail / 模块: Mail -----
	DeleteCharactersCharacterIdMailLabelsLabelId(ctx context.Context, characterId int32, labelId int32, params *models.DeleteCharactersCharacterIdMailLabelsLabelIdParams) error
	DeleteCharactersCharacterIdMailMailId(ctx context.Context, characterId int32, mailId int32, params *models.DeleteCharactersCharacterIdMailMailIdParams) error
	GetCharactersCharacterIdMail(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMailParams) ([]models.GetCharactersCharacterIdMail, error)
	GetCharactersCharacterIdMailLabels(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMailLabelsParams) (*models.GetCharactersCharacterIdMailLabels, error)
	GetCharactersCharacterIdMailLists(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMailListsParams) ([]models.GetCharactersCharacterIdMailLists, error)
	GetCharactersCharacterIdMailMailId(ctx context.Context, characterId int32, mailId int32, params *models.GetCharactersCharacterIdMailMailIdParams) (*models.GetCharactersCharacterIdMailMailId, error)
	PostCharactersCharacterIdMail(ctx context.Context, characterId int32, body *models.PostCharactersCharacterIdMailMail, params *models.PostCharactersCharacterIdMailParams) (int32, error)
	PostCharactersCharacterIdMailLabels(ctx context.Context, characterId int32, body *models.PostCharactersCharacterIdMailLabelsLabel, params *models.PostCharactersCharacterIdMailLabelsParams) (int32, error)
	PutCharactersCharacterIdMailMailId(ctx context.Context, characterId int32, mailId int32, body *models.PutCharactersCharacterIdMailMailIdContents, params *models.PutCharactersCharacterIdMailMailIdParams) error

	// ----- Module: Opportunities / 模块: Opportunities -----
	GetCharactersCharacterIdOpportunities(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOpportunitiesParams) ([]models.GetCharactersCharacterIdOpportunities, error)
	GetOpportunitiesGroups(ctx context.Context, params *models.GetGroupsParams) ([]int32, error)
	GetOpportunitiesGroupsGroupId(ctx context.Context, groupId int32, params *models.GetGroupsGroupIdParams) (*models.GetOpportunitiesGroupsGroupId, error)
	GetTasks(ctx context.Context, params *models.GetTasksParams) ([]int32, error)
	GetTasksTaskId(ctx context.Context, taskId int32, params *models.GetTasksTaskIdParams) (*models.GetOpportunitiesTasksTaskId, error)

	// ----- Module: Market / 模块: Market -----
	GetCharactersCharacterIdOrders(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOrdersParams) ([]models.GetCharactersCharacterIdOrders, error)
	GetCharactersCharacterIdOrdersHistory(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOrdersHistoryParams) ([]models.GetCharactersCharacterIdOrdersHistory, error)
	GetCorporationsCorporationIdOrders(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdOrdersParams) ([]models.GetCorporationsCorporationIdOrders, error)
	GetCorporationsCorporationIdOrdersHistory(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdOrdersHistoryParams) ([]models.GetCorporationsCorporationIdOrdersHistory, error)
	GetsGroups(ctx context.Context, params *models.GetsGroupsParams) ([]int32, error)
	GetsGroupsMarketGroupId(ctx context.Context, marketGroupId int32, params *models.GetsGroupsMarketGroupIdParams) (*models.GetMarketsGroupsMarketGroupId, error)
	GetsPrices(ctx context.Context, params *models.GetsPricesParams) ([]models.GetMarketsPrices, error)
	GetsRegionIdHistory(ctx context.Context, regionId int32, params *models.GetsRegionIdHistoryParams) ([]models.GetMarketsRegionIdHistory, error)
	GetsRegionIdOrders(ctx context.Context, regionId int32, params *models.GetsRegionIdOrdersParams) ([]models.GetMarketsRegionIdOrders, error)
	GetsRegionIdTypes(ctx context.Context, regionId int32, params *models.GetsRegionIdTypesParams) ([]int32, error)
	GetsStructuresStructureId(ctx context.Context, structureId int64, params *models.GetsStructuresStructureIdParams) ([]models.GetMarketsStructuresStructureId, error)

	// ----- Module: PlanetaryInteraction / 模块: PlanetaryInteraction -----
	GetCharactersCharacterIdPlanets(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdPlanetsParams) ([]models.GetCharactersCharacterIdPlanets, error)
	GetCharactersCharacterIdPlanetsPlanetId(ctx context.Context, characterId int32, planetId int32, params *models.GetCharactersCharacterIdPlanetsPlanetIdParams) (*models.GetCharactersCharacterIdPlanetsPlanetId, error)
	GetCorporationsCorporationIdCustomsOffices(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdCustomsOfficesParams) ([]models.GetCorporationsCorporationIdCustomsOffices, error)
	GetUniverseSchematicsSchematicId(ctx context.Context, schematicId int32, params *models.GetUniverseSchematicsSchematicIdParams) (*models.GetUniverseSchematicsSchematicId, error)

	// ----- Module: Search / 模块: Search -----
	GetCharactersCharacterIdSearch(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdSearchParams) (*models.GetCharactersCharacterIdSearch, error)

	// ----- Module: Wallet / 模块: Wallet -----
	GetCharactersCharacterIdWallet(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdWalletParams) (float64, error)
	GetCharactersCharacterIdWalletJournal(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdWalletJournalParams) ([]models.GetCharactersCharacterIdWalletJournal, error)
	GetCharactersCharacterIdWalletTransactions(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdWalletTransactionsParams) ([]models.GetCharactersCharacterIdWalletTransactions, error)
	GetCorporationsCorporationIdWallets(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdWalletsParams) ([]models.GetCorporationsCorporationIdWallets, error)
	GetCorporationsCorporationIdWalletsDivisionJournal(ctx context.Context, corporationId int32, division int32, params *models.GetCorporationsCorporationIdWalletsDivisionJournalParams) ([]models.GetCorporationsCorporationIdWalletsDivisionJournal, error)
	GetCorporationsCorporationIdWalletsDivisionTransactions(ctx context.Context, corporationId int32, division int32, params *models.GetCorporationsCorporationIdWalletsDivisionTransactionsParams) ([]models.GetCorporationsCorporationIdWalletsDivisionTransactions, error)

	// ----- Module: Corporations / 模块: Corporations -----
	GetCorporationId(ctx context.Context, corporationId int32, params *models.GetCorporationIdParams) (*models.GetCorporationsCorporationId, error)
	GetCorporationIdAlliancehistory(ctx context.Context, corporationId int32, params *models.GetCorporationIdAlliancehistoryParams) ([]models.GetCorporationsCorporationIdAlliancehistory, error)
	GetCorporationIdBlueprints(ctx context.Context, corporationId int32, params *models.GetCorporationIdBlueprintsParams) ([]models.GetCorporationsCorporationIdBlueprints, error)
	GetCorporationIdContainersLogs(ctx context.Context, corporationId int32, params *models.GetCorporationIdContainersLogsParams) ([]models.GetCorporationsCorporationIdContainersLogs, error)
	GetCorporationIdDivisions(ctx context.Context, corporationId int32, params *models.GetCorporationIdDivisionsParams) (*models.GetCorporationsCorporationIdDivisions, error)
	GetCorporationIdFacilities(ctx context.Context, corporationId int32, params *models.GetCorporationIdFacilitiesParams) ([]models.GetCorporationsCorporationIdFacilities, error)
	GetCorporationIdIcons(ctx context.Context, corporationId int32, params *models.GetCorporationIdIconsParams) (*models.GetCorporationsCorporationIdIcons, error)
	GetCorporationIdMedals(ctx context.Context, corporationId int32, params *models.GetCorporationIdMedalsParams) ([]models.GetCorporationsCorporationIdMedals, error)
	GetCorporationIdMedalsIssued(ctx context.Context, corporationId int32, params *models.GetCorporationIdMedalsIssuedParams) ([]models.GetCorporationsCorporationIdMedalsIssued, error)
	GetCorporationIdMembers(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembersParams) ([]int32, error)
	GetCorporationIdMembersLimit(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembersLimitParams) (int32, error)
	GetCorporationIdMembersTitles(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembersTitlesParams) ([]models.GetCorporationsCorporationIdMembersTitles, error)
	GetCorporationIdMembertracking(ctx context.Context, corporationId int32, params *models.GetCorporationIdMembertrackingParams) ([]models.GetCorporationsCorporationIdMembertracking, error)
	GetCorporationIdRoles(ctx context.Context, corporationId int32, params *models.GetCorporationIdRolesParams) ([]models.GetCorporationsCorporationIdRoles, error)
	GetCorporationIdRolesHistory(ctx context.Context, corporationId int32, params *models.GetCorporationIdRolesHistoryParams) ([]models.GetCorporationsCorporationIdRolesHistory, error)
	GetCorporationIdShareholders(ctx context.Context, corporationId int32, params *models.GetCorporationIdShareholdersParams) ([]models.GetCorporationsCorporationIdShareholders, error)
	GetCorporationIdStandings(ctx context.Context, corporationId int32, params *models.GetCorporationIdStandingsParams) ([]models.GetCorporationsCorporationIdStandings, error)
	GetCorporationIdStarbases(ctx context.Context, corporationId int32, params *models.GetCorporationIdStarbasesParams) ([]models.GetCorporationsCorporationIdStarbases, error)
	GetCorporationIdStarbasesStarbaseId(ctx context.Context, corporationId int32, starbaseId int64, params *models.GetCorporationIdStarbasesStarbaseIdParams) (*models.GetCorporationsCorporationIdStarbasesStarbaseId, error)
	GetCorporationIdStructures(ctx context.Context, corporationId int32, params *models.GetCorporationIdStructuresParams) ([]models.GetCorporationsCorporationIdStructures, error)
	GetCorporationIdTitles(ctx context.Context, corporationId int32, params *models.GetCorporationIdTitlesParams) ([]models.GetCorporationsCorporationIdTitles, error)
	GetNpccorps(ctx context.Context, params *models.GetNpccorpsParams) ([]int32, error)

	// ----- Module: Dogma / 模块: Dogma -----
	GetAttributes(ctx context.Context, params *models.GetAttributesParams) ([]int32, error)
	GetAttributesAttributeId(ctx context.Context, attributeId int32, params *models.GetAttributesAttributeIdParams) (*models.GetDogmaAttributesAttributeId, error)
	GetDynamicItemsTypeIdItemId(ctx context.Context, itemId int64, typeId int32, params *models.GetDynamicItemsTypeIdItemIdParams) (*models.GetDogmaDynamicItemsTypeIdItemId, error)
	GetEffects(ctx context.Context, params *models.GetEffectsParams) ([]int32, error)
	GetEffectsEffectId(ctx context.Context, effectId int32, params *models.GetEffectsEffectIdParams) (*models.GetDogmaEffectsEffectId, error)

	// ----- Module: Incursions / 模块: Incursions -----
	GetIncursions(ctx context.Context, params *models.GetIncursionsParams) ([]models.GetIncursions, error)

	// ----- Module: Insurance / 模块: Insurance -----
	GetPrices(ctx context.Context, params *models.GetPricesParams) ([]models.GetInsurancePrices, error)

	// ----- Module: Routes / 模块: Routes -----
	GetOriginDestination(ctx context.Context, destination int32, origin int32, params *models.GetOriginDestinationParams) ([]int32, error)

	// ----- Module: Sovereignty / 模块: Sovereignty -----
	GetCampaigns(ctx context.Context, params *models.GetCampaignsParams) ([]models.GetSovereigntyCampaigns, error)
	GetMap(ctx context.Context, params *models.GetMapParams) ([]models.GetSovereigntyMap, error)
	GetSovereigntyStructures(ctx context.Context, params *models.GetStructuresParams) ([]models.GetSovereigntyStructures, error)

	// ----- Module: Status / 模块: Status -----
	GetStatus(ctx context.Context, params *models.GetStatusParams) (*models.GetStatus, error)

	// ----- Module: UserInterface / 模块: UserInterface -----
	PostUiAutopilotWaypoint(ctx context.Context, params *models.PostUiAutopilotWaypointParams) error
	PostUiOpenwindowContract(ctx context.Context, params *models.PostUiOpenwindowContractParams) error
	PostUiOpenwindowInformation(ctx context.Context, params *models.PostUiOpenwindowInformationParams) error
	PostUiOpenwindowMarketdetails(ctx context.Context, params *models.PostUiOpenwindowMarketdetailsParams) error
	PostUiOpenwindowNewmail(ctx context.Context, body *models.PostUiOpenwindowNewmailNewMail, params *models.PostUiOpenwindowNewmailParams) error

	// ----- Module: Universe / 模块: Universe -----
	GetAncestries(ctx context.Context, params *models.GetAncestriesParams) ([]models.GetUniverseAncestries, error)
	GetAsteroidBeltsAsteroidBeltId(ctx context.Context, asteroidBeltId int32, params *models.GetAsteroidBeltsAsteroidBeltIdParams) (*models.GetUniverseAsteroidBeltsAsteroidBeltId, error)
	GetBloodlines(ctx context.Context, params *models.GetBloodlinesParams) ([]models.GetUniverseBloodlines, error)
	GetCategories(ctx context.Context, params *models.GetCategoriesParams) ([]int32, error)
	GetCategoriesCategoryId(ctx context.Context, categoryId int32, params *models.GetCategoriesCategoryIdParams) (*models.GetUniverseCategoriesCategoryId, error)
	GetConstellations(ctx context.Context, params *models.GetConstellationsParams) ([]int32, error)
	GetConstellationsConstellationId(ctx context.Context, constellationId int32, params *models.GetConstellationsConstellationIdParams) (*models.GetUniverseConstellationsConstellationId, error)
	GetFactions(ctx context.Context, params *models.GetFactionsParams) ([]models.GetUniverseFactions, error)
	GetGraphics(ctx context.Context, params *models.GetGraphicsParams) ([]int32, error)
	GetGraphicsGraphicId(ctx context.Context, graphicId int32, params *models.GetGraphicsGraphicIdParams) (*models.GetUniverseGraphicsGraphicId, error)
	GetGroups(ctx context.Context, params *models.GetGroupsParamsX) ([]int32, error)
	GetGroupsGroupId(ctx context.Context, groupId int32, params *models.GetGroupsGroupIdParamsX) (*models.GetUniverseGroupsGroupId, error)
	GetMoonsMoonId(ctx context.Context, moonId int32, params *models.GetMoonsMoonIdParams) (*models.GetUniverseMoonsMoonId, error)
	GetPlanetsPlanetId(ctx context.Context, planetId int32, params *models.GetPlanetsPlanetIdParams) (*models.GetUniversePlanetsPlanetId, error)
	GetRaces(ctx context.Context, params *models.GetRacesParams) ([]models.GetUniverseRaces, error)
	GetRegions(ctx context.Context, params *models.GetRegionsParams) ([]int32, error)
	GetRegionsRegionId(ctx context.Context, regionId int32, params *models.GetRegionsRegionIdParams) (*models.GetUniverseRegionsRegionId, error)
	GetStargatesStargateId(ctx context.Context, stargateId int32, params *models.GetStargatesStargateIdParams) (*models.GetUniverseStargatesStargateId, error)
	GetStarsStarId(ctx context.Context, starId int32, params *models.GetStarsStarIdParams) (*models.GetUniverseStarsStarId, error)
	GetStationsStationId(ctx context.Context, stationId int32, params *models.GetStationsStationIdParams) (*models.GetUniverseStationsStationId, error)
	GetStructures(ctx context.Context, params *models.GetStructuresParamsX) ([]int64, error)
	GetStructuresStructureId(ctx context.Context, structureId int64, params *models.GetStructuresStructureIdParams) (*models.GetUniverseStructuresStructureId, error)
	GetSystemJumps(ctx context.Context, params *models.GetSystemJumpsParams) ([]models.GetUniverseSystemJumps, error)
	GetSystemKills(ctx context.Context, params *models.GetSystemKillsParams) ([]models.GetUniverseSystemKills, error)
	GetSystems(ctx context.Context, params *models.GetSystemsParamsX) ([]int32, error)
	GetSystemsSystemId(ctx context.Context, systemId int32, params *models.GetSystemsSystemIdParams) (*models.GetUniverseSystemsSystemId, error)
	GetTypes(ctx context.Context, params *models.GetTypesParams) ([]int32, error)
	GetTypesTypeId(ctx context.Context, typeId int32, params *models.GetTypesTypeIdParams) (*models.GetUniverseTypesTypeId, error)
	PostIds(ctx context.Context, body []string, params *models.PostIdsParams) (*models.PostUniverseIds, error)
	PostNames(ctx context.Context, body []int32, params *models.PostNamesParams) ([]models.PostUniverseNames, error)

	// ----- Module: Wars / 模块: Wars -----
	GetWarId(ctx context.Context, warId int32, params *models.GetWarIdParams) (*models.GetWarsWarId, error)
	GetWarIdKillmails(ctx context.Context, warId int32, params *models.GetWarIdKillmailsParams) ([]models.GetWarsWarIdKillmails, error)
	GetWars(ctx context.Context, params *models.GetWarsParams) ([]int32, error)
}

var _ ClientIface = (*Client)(nil)
