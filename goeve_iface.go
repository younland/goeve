package goeve

import (
	"context"

	"github.com/younland/goeve/models"
)

// ClientIface is the interface of the ESI API client, grouped by module.
// ClientIface 是 ESI API 客户端的接口，按模块分组。
type ClientIface interface {

	// ----- Module: Alliance / 模块: Alliance -----
	GetAlliance(ctx context.Context, allianceID int32, opts ...RequestOption) (*models.Alliance, error)
	GetAllianceCorporations(ctx context.Context, allianceID int32, opts ...RequestOption) ([]int32, error)
	GetAllianceIcons(ctx context.Context, allianceID int32, opts ...RequestOption) (*models.AllianceIcons, error)
	GetAlliances(ctx context.Context, opts ...RequestOption) ([]int32, error)

	// ----- Module: Assets / 模块: Assets -----
	GetCharacterAssets(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterAsset, error)
	GetCorporationAssets(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationAsset, error)
	GetCharacterAssetLocations(ctx context.Context, characterID int32, body []int64, opts ...RequestOption) ([]models.AssetLocation, error)
	GetCharacterAssetNames(ctx context.Context, characterID int32, body []int64, opts ...RequestOption) ([]models.AssetName, error)
	GetCorporationAssetLocations(ctx context.Context, corporationID int32, body []int64, opts ...RequestOption) ([]models.AssetLocation, error)
	GetCorporationAssetNames(ctx context.Context, corporationID int32, body []int64, opts ...RequestOption) ([]models.AssetName, error)

	// ----- Module: Bookmarks / 模块: Bookmarks -----
	GetCharacterBookmarks(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterBookmark, error)
	GetCharacterBookmarkFolders(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterBookmarkFolder, error)
	ListCorporationBookmarks(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationBookmark, error)
	ListCorporationBookmarkFolders(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationBookmarkFolder, error)

	// ----- Module: Calendar / 模块: Calendar -----
	GetCharacterCalendarEvents(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CalendarEventSummary, error)
	GetCalendarEvent(ctx context.Context, characterID int32, eventID int32, opts ...RequestOption) (*models.CalendarEvent, error)
	GetCalendarEventAttendees(ctx context.Context, characterID int32, eventID int32, opts ...RequestOption) ([]models.CalendarEventAttendee, error)
	RespondToCalendarEvent(ctx context.Context, characterID int32, eventID int32, body *models.CalendarEventResponse, opts ...RequestOption) error

	// ----- Module: Character / 模块: Character -----
	GetCharacter(ctx context.Context, characterID int32, opts ...RequestOption) (*models.Character, error)
	GetCharacterAgentsResearch(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.AgentResearch, error)
	GetCharacterBlueprints(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.Blueprint, error)
	GetCharacterCorporationHistory(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CorporationHistoryEntry, error)
	GetCharacterJumpFatigue(ctx context.Context, characterID int32, opts ...RequestOption) (*models.JumpFatigue, error)
	GetCharacterMedals(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.Medal, error)
	GetCharacterNotifications(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.Notification, error)
	GetCharacterContactNotifications(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.ContactNotification, error)
	GetCharacterPortrait(ctx context.Context, characterID int32, opts ...RequestOption) (*models.CharacterPortraits, error)
	GetCharacterCorporationRoles(ctx context.Context, characterID int32, opts ...RequestOption) (*models.CharacterCorporationRoles, error)
	GetCharacterStandings(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.Standing, error)
	GetCharacterCorporationTitles(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterTitle, error)
	CharacterAffiliation(ctx context.Context, body []int32, opts ...RequestOption) ([]models.CharacterAffiliation, error)
	CalculateCharacterCspaCharge(ctx context.Context, characterID int32, body []int32, opts ...RequestOption) (float64, error)

	// ----- Module: Clones / 模块: Clones -----
	GetCharacterClones(ctx context.Context, characterID int32, opts ...RequestOption) (*models.Clones, error)
	GetCharacterImplants(ctx context.Context, characterID int32, opts ...RequestOption) ([]int32, error)

	// ----- Module: Contacts / 模块: Contacts -----
	DeleteCharacterContacts(ctx context.Context, characterID int32, contactIDs []int32, opts ...RequestOption) error
	GetAllianceContacts(ctx context.Context, allianceID int32, opts ...RequestOption) ([]models.AllianceContact, error)
	GetAllianceContactLabels(ctx context.Context, allianceID int32, opts ...RequestOption) ([]models.ContactLabel, error)
	GetCharacterContacts(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterContact, error)
	GetCharacterContactLabels(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.ContactLabel, error)
	GetCorporationContacts(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationContact, error)
	GetCorporationContactLabels(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.ContactLabel, error)
	AddCharacterContacts(ctx context.Context, characterID int32, standing float64, body []int32, opts ...RequestOption) ([]int32, error)
	EditCharacterContacts(ctx context.Context, characterID int32, standing float64, body []int32, opts ...RequestOption) error

	// ----- Module: Contracts / 模块: Contracts -----
	GetCharacterContracts(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.Contract, error)
	GetCharacterContractBids(ctx context.Context, characterID int32, contractID int32, opts ...RequestOption) ([]models.ContractBid, error)
	GetCharacterContractItems(ctx context.Context, characterID int32, contractID int32, opts ...RequestOption) ([]models.ContractItem, error)
	GetCorporationContracts(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Contract, error)
	GetCorporationContractBids(ctx context.Context, contractID int32, corporationID int32, opts ...RequestOption) ([]models.ContractBid, error)
	GetCorporationContractItems(ctx context.Context, contractID int32, corporationID int32, opts ...RequestOption) ([]models.ContractItem, error)
	GetPublicContractBids(ctx context.Context, contractID int32, opts ...RequestOption) ([]models.PublicContractBid, error)
	GetPublicContractItems(ctx context.Context, contractID int32, opts ...RequestOption) ([]models.PublicContractItem, error)
	GetPublicContracts(ctx context.Context, regionID int32, opts ...RequestOption) ([]models.PublicContract, error)

	// ----- Module: Corporation / 模块: Corporation -----
	GetCorporationInformation(ctx context.Context, corporationID int32, opts ...RequestOption) (*models.Corporation, error)
	GetCorporationAllianceHistory(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.AllianceHistoryEntry, error)
	GetCorporationBlueprints(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Blueprint, error)
	GetCorporationContainerLogs(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.ContainerLog, error)
	GetCorporationDivisions(ctx context.Context, corporationID int32, opts ...RequestOption) (*models.CorporationDivisions, error)
	GetCorporationFacilities(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationFacility, error)
	GetCorporationIcon(ctx context.Context, corporationID int32, opts ...RequestOption) (*models.CorporationIcons, error)
	GetCorporationMedals(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationMedal, error)
	GetCorporationIssuedMedals(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.IssuedMedal, error)
	GetCorporationMembers(ctx context.Context, corporationID int32, opts ...RequestOption) ([]int32, error)
	GetCorporationMemberLimit(ctx context.Context, corporationID int32, opts ...RequestOption) (int32, error)
	GetCorporationMemberTitles(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.MemberTitles, error)
	GetCorporationMemberTracking(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.MemberTrackingEntry, error)
	GetCorporationMemberRoles(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationMemberRoles, error)
	GetCorporationMemberRolesHistory(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationRoleHistory, error)
	GetCorporationShareholders(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Shareholder, error)
	GetCorporationStandings(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Standing, error)
	GetCorporationStarbases(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.Starbase, error)
	GetCorporationStarbase(ctx context.Context, corporationID int32, starbaseID int64, systemID string, opts ...RequestOption) (*models.StarbaseDetail, error)
	GetCorporationStructures(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationStructure, error)
	GetCorporationTitles(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationTitle, error)
	GetNpcCorporations(ctx context.Context, opts ...RequestOption) ([]int32, error)

	// ----- Module: Dogma / 模块: Dogma -----
	GetDogmaAttributes(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetDogmaAttribute(ctx context.Context, attributeID int32, opts ...RequestOption) (*models.DogmaAttribute, error)
	GetDogmaDynamicItem(ctx context.Context, itemID int64, typeID int32, opts ...RequestOption) (*models.DogmaDynamicItem, error)
	GetDogmaEffects(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetDogmaEffect(ctx context.Context, effectID int32, opts ...RequestOption) (*models.DogmaEffect, error)

	// ----- Module: FactionWarfare / 模块: FactionWarfare -----
	GetCharacterFactionWarfareStats(ctx context.Context, characterID int32, opts ...RequestOption) (*models.CharacterFactionWarfareStats, error)
	GetCorporationFactionWarfareStats(ctx context.Context, corporationID int32, opts ...RequestOption) (*models.CorporationFactionWarfareStats, error)
	GetFactionWarfareLeaderboard(ctx context.Context, opts ...RequestOption) (*models.FactionWarfareLeaderboard, error)
	GetFactionWarfareCharacterLeaderboard(ctx context.Context, opts ...RequestOption) (*models.FactionWarfareCharacterLeaderboard, error)
	GetFactionWarfareCorporationLeaderboard(ctx context.Context, opts ...RequestOption) (*models.FactionWarfareCorporationLeaderboard, error)
	GetFactionWarfareStats(ctx context.Context, opts ...RequestOption) ([]models.FactionWarfareStats, error)
	GetFactionWarfareSystems(ctx context.Context, opts ...RequestOption) ([]models.FactionWarfareSystem, error)
	GetFactionWarfareWars(ctx context.Context, opts ...RequestOption) ([]models.FactionWarfareWar, error)

	// ----- Module: Fittings / 模块: Fittings -----
	DeleteCharacterFitting(ctx context.Context, characterID int32, fittingID int32, opts ...RequestOption) error
	GetCharacterFittings(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.Fitting, error)
	CreateCharacterFitting(ctx context.Context, characterID int32, body *models.FittingRequest, opts ...RequestOption) (*models.NewFitting, error)

	// ----- Module: Fleets / 模块: Fleets -----
	GetCharacterFleet(ctx context.Context, characterID int32, opts ...RequestOption) (*models.FleetMembership, error)
	GetFleet(ctx context.Context, fleetID int64, opts ...RequestOption) (*models.Fleet, error)
	UpdateFleetSettings(ctx context.Context, fleetID int64, body *models.FleetSettings, opts ...RequestOption) error
	GetFleetMembers(ctx context.Context, fleetID int64, opts ...RequestOption) ([]models.FleetMember, error)
	CreateFleetInvitation(ctx context.Context, fleetID int64, body *models.FleetInvitation, opts ...RequestOption) error
	KickFleetMember(ctx context.Context, fleetID int64, memberID int32, opts ...RequestOption) error
	MoveFleetMember(ctx context.Context, fleetID int64, memberID int32, body *models.FleetMemberMovement, opts ...RequestOption) error
	DeleteFleetSquad(ctx context.Context, fleetID int64, squadID int64, opts ...RequestOption) error
	RenameFleetSquad(ctx context.Context, fleetID int64, squadID int64, body *models.FleetNaming, opts ...RequestOption) error
	GetFleetWings(ctx context.Context, fleetID int64, opts ...RequestOption) ([]models.FleetWing, error)
	CreateFleetWing(ctx context.Context, fleetID int64, opts ...RequestOption) (*models.NewFleetWing, error)
	DeleteFleetWing(ctx context.Context, fleetID int64, wingID int64, opts ...RequestOption) error
	RenameFleetWing(ctx context.Context, fleetID int64, wingID int64, body *models.FleetNaming, opts ...RequestOption) error
	CreateFleetSquad(ctx context.Context, fleetID int64, wingID int64, opts ...RequestOption) (*models.NewFleetSquad, error)

	// ----- Module: Incursions / 模块: Incursions -----
	GetIncursions(ctx context.Context, opts ...RequestOption) ([]models.Incursion, error)

	// ----- Module: Industry / 模块: Industry -----
	GetCharacterIndustryJobs(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterIndustryJob, error)
	GetCharacterMiningLedger(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.MiningLedgerEntry, error)
	GetCorporationMoonExtractions(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.MoonExtraction, error)
	GetCorporationMiningObservers(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.MiningObserver, error)
	GetCorporationMiningObserverData(ctx context.Context, corporationID int32, observerID int64, opts ...RequestOption) ([]models.MiningObserverEntry, error)
	GetCorporationIndustryJobs(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationIndustryJob, error)
	GetIndustryFacilities(ctx context.Context, opts ...RequestOption) ([]models.IndustryFacility, error)
	GetIndustrySystemCostIndices(ctx context.Context, opts ...RequestOption) ([]models.IndustrySystemCostIndices, error)

	// ----- Module: Insurance / 模块: Insurance -----
	GetPrices(ctx context.Context, opts ...RequestOption) ([]models.InsurancePrice, error)

	// ----- Module: Killmails / 模块: Killmails -----
	GetCharacterKillmails(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.KillmailRef, error)
	GetCorporationKillmails(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.KillmailRef, error)
	GetKillmail(ctx context.Context, killmailHash string, killmailID int32, opts ...RequestOption) (*models.Killmail, error)

	// ----- Module: Location / 模块: Location -----
	GetCharacterLocation(ctx context.Context, characterID int32, opts ...RequestOption) (*models.CharacterLocation, error)
	GetCharacterOnline(ctx context.Context, characterID int32, opts ...RequestOption) (*models.OnlineStatus, error)
	GetCharacterShip(ctx context.Context, characterID int32, opts ...RequestOption) (*models.CharacterShip, error)

	// ----- Module: Loyalty / 模块: Loyalty -----
	GetCharacterLoyaltyPoints(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.LoyaltyPoints, error)
	GetLoyaltyStoreOffers(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.LoyaltyStoreOffer, error)

	// ----- Module: Mail / 模块: Mail -----
	DeleteCharacterMailLabel(ctx context.Context, characterID int32, labelID int32, opts ...RequestOption) error
	DeleteCharacterMail(ctx context.Context, characterID int32, mailID int32, opts ...RequestOption) error
	GetCharacterMails(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.MailHeader, error)
	GetCharacterMailLabels(ctx context.Context, characterID int32, opts ...RequestOption) (*models.MailLabels, error)
	GetCharacterMailLists(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.MailingList, error)
	GetCharacterMail(ctx context.Context, characterID int32, mailID int32, opts ...RequestOption) (*models.Mail, error)
	SendCharacterMail(ctx context.Context, characterID int32, body *models.MailRequest, opts ...RequestOption) (int32, error)
	CreateCharacterMailLabel(ctx context.Context, characterID int32, body *models.MailLabelRequest, opts ...RequestOption) (int32, error)
	UpdateCharacterMail(ctx context.Context, characterID int32, mailID int32, body *models.MailMetadata, opts ...RequestOption) error

	// ----- Module: Market / 模块: Market -----
	GetCharacterMarketOrders(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterMarketOrder, error)
	GetCharacterMarketOrderHistory(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterMarketOrderHistory, error)
	GetCorporationMarketOrders(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationMarketOrder, error)
	GetCorporationMarketOrderHistory(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationMarketOrderHistory, error)
	GetMarketGroups(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetMarketGroup(ctx context.Context, marketGroupID int32, opts ...RequestOption) (*models.MarketGroup, error)
	GetMarketPrices(ctx context.Context, opts ...RequestOption) ([]models.MarketPrice, error)
	GetMarketHistory(ctx context.Context, regionID int32, typeID string, opts ...RequestOption) ([]models.MarketHistoryEntry, error)
	GetMarketOrders(ctx context.Context, regionID int32, orderType string, opts ...RequestOption) ([]models.MarketOrder, error)
	GetMarketTypes(ctx context.Context, regionID int32, opts ...RequestOption) ([]int32, error)
	GetStructureMarketOrders(ctx context.Context, structureID int64, opts ...RequestOption) ([]models.StructureMarketOrder, error)

	// ----- Module: Opportunities / 模块: Opportunities -----
	GetCharacterOpportunities(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.OpportunityCompletion, error)
	GetOpportunityGroups(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetOpportunityGroup(ctx context.Context, groupID int32, opts ...RequestOption) (*models.OpportunityGroup, error)
	GetOpportunityTasks(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetOpportunityTask(ctx context.Context, taskID int32, opts ...RequestOption) (*models.OpportunityTask, error)

	// ----- Module: PlanetaryInteraction / 模块: PlanetaryInteraction -----
	GetCharacterColonies(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.Colony, error)
	GetCharacterColonyLayout(ctx context.Context, characterID int32, planetID int32, opts ...RequestOption) (*models.ColonyLayout, error)
	GetCorporationCustomsOffices(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CustomsOffice, error)
	GetSchematicInformation(ctx context.Context, schematicID int32, opts ...RequestOption) (*models.Schematic, error)

	// ----- Module: Routes / 模块: Routes -----
	GetRoute(ctx context.Context, destination int32, origin int32, opts ...RequestOption) ([]int32, error)

	// ----- Module: Search / 模块: Search -----
	SearchEntities(ctx context.Context, characterID int32, categories []string, search string, opts ...RequestOption) (*models.SearchResult, error)

	// ----- Module: Skills / 模块: Skills -----
	GetCharacterAttributes(ctx context.Context, characterID int32, opts ...RequestOption) (*models.CharacterAttributes, error)
	GetCharacterSkillQueue(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.SkillQueueEntry, error)
	GetCharacterSkills(ctx context.Context, characterID int32, opts ...RequestOption) (*models.CharacterSkills, error)

	// ----- Module: Sovereignty / 模块: Sovereignty -----
	GetSovereigntyCampaigns(ctx context.Context, opts ...RequestOption) ([]models.SovereigntyCampaign, error)
	GetSovereigntyMap(ctx context.Context, opts ...RequestOption) ([]models.SovereigntySystem, error)
	GetSovereigntyStructures(ctx context.Context, opts ...RequestOption) ([]models.SovereigntyStructure, error)

	// ----- Module: Status / 模块: Status -----
	GetServerStatus(ctx context.Context, opts ...RequestOption) (*models.ServerStatus, error)

	// ----- Module: UserInterface / 模块: UserInterface -----
	SetAutopilotWaypoint(ctx context.Context, addToBeginning bool, clearOtherWaypoints bool, destinationID int64, opts ...RequestOption) error
	OpenContractWindow(ctx context.Context, contractID string, opts ...RequestOption) error
	OpenInformationWindow(ctx context.Context, targetID string, opts ...RequestOption) error
	OpenMarketDetails(ctx context.Context, typeID string, opts ...RequestOption) error
	OpenNewMailWindow(ctx context.Context, body *models.NewMailRequest, opts ...RequestOption) error

	// ----- Module: Universe / 模块: Universe -----
	GetUniverseAncestries(ctx context.Context, opts ...RequestOption) ([]models.UniverseAncestry, error)
	GetUniverseAsteroidBelt(ctx context.Context, asteroidBeltID int32, opts ...RequestOption) (*models.UniverseAsteroidBelt, error)
	GetUniverseBloodlines(ctx context.Context, opts ...RequestOption) ([]models.UniverseBloodline, error)
	GetUniverseCategories(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetUniverseCategory(ctx context.Context, categoryID int32, opts ...RequestOption) (*models.UniverseCategory, error)
	GetConstellations(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetConstellationInformation(ctx context.Context, constellationID int32, opts ...RequestOption) (*models.UniverseConstellation, error)
	GetUniverseFactions(ctx context.Context, opts ...RequestOption) ([]models.UniverseFaction, error)
	GetUniverseGraphics(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetUniverseGraphic(ctx context.Context, graphicID int32, opts ...RequestOption) (*models.UniverseGraphic, error)
	GetUniverseGroups(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetUniverseGroup(ctx context.Context, groupID int32, opts ...RequestOption) (*models.UniverseGroup, error)
	GetUniverseMoon(ctx context.Context, moonID int32, opts ...RequestOption) (*models.UniverseMoon, error)
	GetUniversePlanet(ctx context.Context, planetID int32, opts ...RequestOption) (*models.UniversePlanet, error)
	GetUniverseRaces(ctx context.Context, opts ...RequestOption) ([]models.UniverseRace, error)
	GetUniverseRegions(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetUniverseRegion(ctx context.Context, regionID int32, opts ...RequestOption) (*models.UniverseRegion, error)
	GetUniverseStargate(ctx context.Context, stargateID int32, opts ...RequestOption) (*models.UniverseStargate, error)
	GetUniverseStar(ctx context.Context, starID int32, opts ...RequestOption) (*models.UniverseStar, error)
	GetUniverseStation(ctx context.Context, stationID int32, opts ...RequestOption) (*models.UniverseStation, error)
	GetPublicStructures(ctx context.Context, opts ...RequestOption) ([]int64, error)
	GetUniverseStructure(ctx context.Context, structureID int64, opts ...RequestOption) (*models.UniverseStructure, error)
	GetUniverseSystemJumps(ctx context.Context, opts ...RequestOption) ([]models.UniverseSystemJump, error)
	GetUniverseSystemKills(ctx context.Context, opts ...RequestOption) ([]models.UniverseSystemKills, error)
	GetUniverseSystems(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetUniverseSystem(ctx context.Context, systemID int32, opts ...RequestOption) (*models.UniverseSolarSystem, error)
	GetUniverseTypes(ctx context.Context, opts ...RequestOption) ([]int32, error)
	GetUniverseType(ctx context.Context, typeID int32, opts ...RequestOption) (*models.UniverseType, error)
	ResolveNamesToIDs(ctx context.Context, body []string, opts ...RequestOption) (*models.ResolvedIds, error)
	ResolveIDsToNames(ctx context.Context, body []int32, opts ...RequestOption) ([]models.UniverseName, error)

	// ----- Module: Wallet / 模块: Wallet -----
	GetCharacterWalletBalance(ctx context.Context, characterID int32, opts ...RequestOption) (float64, error)
	GetCharacterWalletJournal(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.WalletJournalEntry, error)
	GetWalletTransactions(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.CharacterWalletTransaction, error)
	GetCorporationWallets(ctx context.Context, corporationID int32, opts ...RequestOption) ([]models.CorporationWallet, error)
	GetCorporationWalletJournal(ctx context.Context, corporationID int32, division int32, opts ...RequestOption) ([]models.WalletJournalEntry, error)
	GetCorporationWalletTransactions(ctx context.Context, corporationID int32, division int32, opts ...RequestOption) ([]models.CorporationWalletTransaction, error)

	// ----- Module: Wars / 模块: Wars -----
	GetWar(ctx context.Context, warID int32, opts ...RequestOption) (*models.War, error)
	GetWarKillmails(ctx context.Context, warID int32, opts ...RequestOption) ([]models.KillmailRef, error)
	GetWars(ctx context.Context, opts ...RequestOption) ([]int32, error)
}

var _ ClientIface = (*Client)(nil)
