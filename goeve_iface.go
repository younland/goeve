package goeve

import (
	"context"

	"github.com/younland/goeve/models"
)

// ClientIface is the interface of the ESI API client, grouped by module.
// ClientIface 是 ESI API 客户端的接口，按模块分组。
type ClientIface interface {

	// ----- Module: Alliance / 模块: Alliance -----
	GetAlliance(ctx context.Context, allianceID int32, ifNoneMatch ...string) (*models.Alliance, error)
	GetAllianceCorporations(ctx context.Context, allianceID int32, ifNoneMatch ...string) ([]int32, error)
	GetAllianceIcons(ctx context.Context, allianceID int32, ifNoneMatch ...string) (*models.AllianceIcons, error)
	GetAlliances(ctx context.Context, ifNoneMatch ...string) ([]int32, error)

	// ----- Module: Assets / 模块: Assets -----
	GetCharacterAssets(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.CharacterAsset, error)
	GetCorporationAssets(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationAsset, error)
	GetCharacterAssetLocations(ctx context.Context, token string, characterID int32, body []int64) ([]models.AssetLocation, error)
	GetCharacterAssetNames(ctx context.Context, token string, characterID int32, body []int64) ([]models.AssetName, error)
	GetCorporationAssetLocations(ctx context.Context, token string, corporationID int32, body []int64) ([]models.AssetLocation, error)
	GetCorporationAssetNames(ctx context.Context, token string, corporationID int32, body []int64) ([]models.AssetName, error)

	// ----- Module: Bookmarks / 模块: Bookmarks -----
	GetCharacterBookmarks(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.CharacterBookmark, error)
	GetCharacterBookmarkFolders(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.CharacterBookmarkFolder, error)
	ListCorporationBookmarks(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationBookmark, error)
	ListCorporationBookmarkFolders(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationBookmarkFolder, error)

	// ----- Module: Calendar / 模块: Calendar -----
	GetCharacterCalendarEvents(ctx context.Context, token string, characterID int32, fromEvent int32, ifNoneMatch ...string) ([]models.CalendarEventSummary, error)
	GetCalendarEvent(ctx context.Context, token string, characterID int32, eventID int32, ifNoneMatch ...string) (*models.CalendarEvent, error)
	GetCalendarEventAttendees(ctx context.Context, token string, characterID int32, eventID int32, ifNoneMatch ...string) ([]models.CalendarEventAttendee, error)
	RespondToCalendarEvent(ctx context.Context, token string, characterID int32, eventID int32, body *models.CalendarEventResponse) error

	// ----- Module: Character / 模块: Character -----
	GetCharacter(ctx context.Context, characterID int32, ifNoneMatch ...string) (*models.Character, error)
	GetCharacterAgentsResearch(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.AgentResearch, error)
	GetCharacterBlueprints(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.Blueprint, error)
	GetCharacterCorporationHistory(ctx context.Context, characterID int32, ifNoneMatch ...string) ([]models.CorporationHistoryEntry, error)
	GetCharacterJumpFatigue(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.JumpFatigue, error)
	GetCharacterMedals(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Medal, error)
	GetCharacterNotifications(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Notification, error)
	GetCharacterContactNotifications(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.ContactNotification, error)
	GetCharacterPortrait(ctx context.Context, characterID int32, ifNoneMatch ...string) (*models.CharacterPortraits, error)
	GetCharacterCorporationRoles(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterCorporationRoles, error)
	GetCharacterStandings(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Standing, error)
	GetCharacterCorporationTitles(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.CharacterTitle, error)
	CharacterAffiliation(ctx context.Context, body []int32) ([]models.CharacterAffiliation, error)
	CalculateCharacterCspaCharge(ctx context.Context, token string, characterID int32, body []int32) (float64, error)

	// ----- Module: Clones / 模块: Clones -----
	GetCharacterClones(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.Clones, error)
	GetCharacterImplants(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]int32, error)

	// ----- Module: Contacts / 模块: Contacts -----
	DeleteCharacterContacts(ctx context.Context, token string, characterID int32, contactIDs []int32) error
	GetAllianceContacts(ctx context.Context, token string, allianceID int32, page int32, ifNoneMatch ...string) ([]models.AllianceContact, error)
	GetAllianceContactLabels(ctx context.Context, token string, allianceID int32, ifNoneMatch ...string) ([]models.ContactLabel, error)
	GetCharacterContacts(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.CharacterContact, error)
	GetCharacterContactLabels(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.ContactLabel, error)
	GetCorporationContacts(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationContact, error)
	GetCorporationContactLabels(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]models.ContactLabel, error)
	AddCharacterContacts(ctx context.Context, token string, characterID int32, standing float64, body []int32, labelIDs []int32, watched bool) ([]int32, error)
	EditCharacterContacts(ctx context.Context, token string, characterID int32, standing float64, body []int32, labelIDs []int32, watched bool) error

	// ----- Module: Contracts / 模块: Contracts -----
	GetCharacterContracts(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.Contract, error)
	GetCharacterContractBids(ctx context.Context, token string, characterID int32, contractID int32, ifNoneMatch ...string) ([]models.ContractBid, error)
	GetCharacterContractItems(ctx context.Context, token string, characterID int32, contractID int32, ifNoneMatch ...string) ([]models.ContractItem, error)
	GetCorporationContracts(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.Contract, error)
	GetCorporationContractBids(ctx context.Context, token string, contractID int32, corporationID int32, page int32, ifNoneMatch ...string) ([]models.ContractBid, error)
	GetCorporationContractItems(ctx context.Context, token string, contractID int32, corporationID int32, ifNoneMatch ...string) ([]models.ContractItem, error)
	GetPublicContractBids(ctx context.Context, contractID int32, page int32, ifNoneMatch ...string) ([]models.PublicContractBid, error)
	GetPublicContractItems(ctx context.Context, contractID int32, page int32, ifNoneMatch ...string) ([]models.PublicContractItem, error)
	GetPublicContracts(ctx context.Context, regionID int32, page int32, ifNoneMatch ...string) ([]models.PublicContract, error)

	// ----- Module: Corporation / 模块: Corporation -----
	GetCorporationInformation(ctx context.Context, corporationID int32, ifNoneMatch ...string) (*models.Corporation, error)
	GetCorporationAllianceHistory(ctx context.Context, corporationID int32, ifNoneMatch ...string) ([]models.AllianceHistoryEntry, error)
	GetCorporationBlueprints(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.Blueprint, error)
	GetCorporationContainerLogs(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.ContainerLog, error)
	GetCorporationDivisions(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) (*models.CorporationDivisions, error)
	GetCorporationFacilities(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]models.CorporationFacility, error)
	GetCorporationIcon(ctx context.Context, corporationID int32, ifNoneMatch ...string) (*models.CorporationIcons, error)
	GetCorporationMedals(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationMedal, error)
	GetCorporationIssuedMedals(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.IssuedMedal, error)
	GetCorporationMembers(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]int32, error)
	GetCorporationMemberLimit(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) (int32, error)
	GetCorporationMemberTitles(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]models.MemberTitles, error)
	GetCorporationMemberTracking(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]models.MemberTrackingEntry, error)
	GetCorporationMemberRoles(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]models.CorporationMemberRoles, error)
	GetCorporationMemberRolesHistory(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationRoleHistory, error)
	GetCorporationShareholders(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.Shareholder, error)
	GetCorporationStandings(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.Standing, error)
	GetCorporationStarbases(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.Starbase, error)
	GetCorporationStarbase(ctx context.Context, token string, corporationID int32, starbaseID int64, systemID string, ifNoneMatch ...string) (*models.StarbaseDetail, error)
	GetCorporationStructures(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationStructure, error)
	GetCorporationTitles(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]models.CorporationTitle, error)
	GetNpcCorporations(ctx context.Context, ifNoneMatch ...string) ([]int32, error)

	// ----- Module: Dogma / 模块: Dogma -----
	GetDogmaAttributes(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetDogmaAttribute(ctx context.Context, attributeID int32, ifNoneMatch ...string) (*models.DogmaAttribute, error)
	GetDogmaDynamicItem(ctx context.Context, itemID int64, typeID int32, ifNoneMatch ...string) (*models.DogmaDynamicItem, error)
	GetDogmaEffects(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetDogmaEffect(ctx context.Context, effectID int32, ifNoneMatch ...string) (*models.DogmaEffect, error)

	// ----- Module: FactionWarfare / 模块: FactionWarfare -----
	GetCharacterFactionWarfareStats(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterFactionWarfareStats, error)
	GetCorporationFactionWarfareStats(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) (*models.CorporationFactionWarfareStats, error)
	GetFactionWarfareLeaderboard(ctx context.Context, ifNoneMatch ...string) (*models.FactionWarfareLeaderboard, error)
	GetFactionWarfareCharacterLeaderboard(ctx context.Context, ifNoneMatch ...string) (*models.FactionWarfareCharacterLeaderboard, error)
	GetFactionWarfareCorporationLeaderboard(ctx context.Context, ifNoneMatch ...string) (*models.FactionWarfareCorporationLeaderboard, error)
	GetFactionWarfareStats(ctx context.Context, ifNoneMatch ...string) ([]models.FactionWarfareStats, error)
	GetFactionWarfareSystems(ctx context.Context, ifNoneMatch ...string) ([]models.FactionWarfareSystem, error)
	GetFactionWarfareWars(ctx context.Context, ifNoneMatch ...string) ([]models.FactionWarfareWar, error)

	// ----- Module: Fittings / 模块: Fittings -----
	DeleteCharacterFitting(ctx context.Context, token string, characterID int32, fittingID int32) error
	GetCharacterFittings(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Fitting, error)
	CreateCharacterFitting(ctx context.Context, token string, characterID int32, body *models.FittingRequest) (*models.NewFitting, error)

	// ----- Module: Fleets / 模块: Fleets -----
	GetCharacterFleet(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.FleetMembership, error)
	GetFleet(ctx context.Context, token string, fleetID int64, ifNoneMatch ...string) (*models.Fleet, error)
	UpdateFleetSettings(ctx context.Context, token string, fleetID int64, body *models.FleetSettings) error
	GetFleetMembers(ctx context.Context, token string, fleetID int64, ifNoneMatch ...string) ([]models.FleetMember, error)
	CreateFleetInvitation(ctx context.Context, token string, fleetID int64, body *models.FleetInvitation) error
	KickFleetMember(ctx context.Context, token string, fleetID int64, memberID int32) error
	MoveFleetMember(ctx context.Context, token string, fleetID int64, memberID int32, body *models.FleetMemberMovement) error
	DeleteFleetSquad(ctx context.Context, token string, fleetID int64, squadID int64) error
	RenameFleetSquad(ctx context.Context, token string, fleetID int64, squadID int64, body *models.FleetNaming) error
	GetFleetWings(ctx context.Context, token string, fleetID int64, ifNoneMatch ...string) ([]models.FleetWing, error)
	CreateFleetWing(ctx context.Context, token string, fleetID int64) (*models.NewFleetWing, error)
	DeleteFleetWing(ctx context.Context, token string, fleetID int64, wingID int64) error
	RenameFleetWing(ctx context.Context, token string, fleetID int64, wingID int64, body *models.FleetNaming) error
	CreateFleetSquad(ctx context.Context, token string, fleetID int64, wingID int64) (*models.NewFleetSquad, error)

	// ----- Module: Incursions / 模块: Incursions -----
	GetIncursions(ctx context.Context, ifNoneMatch ...string) ([]models.Incursion, error)

	// ----- Module: Industry / 模块: Industry -----
	GetCharacterIndustryJobs(ctx context.Context, token string, characterID int32, includeCompleted bool, ifNoneMatch ...string) ([]models.CharacterIndustryJob, error)
	GetCharacterMiningLedger(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.MiningLedgerEntry, error)
	GetCorporationMoonExtractions(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.MoonExtraction, error)
	GetCorporationMiningObservers(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.MiningObserver, error)
	GetCorporationMiningObserverData(ctx context.Context, token string, corporationID int32, observerID int64, page int32, ifNoneMatch ...string) ([]models.MiningObserverEntry, error)
	GetCorporationIndustryJobs(ctx context.Context, token string, corporationID int32, page int32, includeCompleted bool, ifNoneMatch ...string) ([]models.CorporationIndustryJob, error)
	GetIndustryFacilities(ctx context.Context, ifNoneMatch ...string) ([]models.IndustryFacility, error)
	GetIndustrySystemCostIndices(ctx context.Context, ifNoneMatch ...string) ([]models.IndustrySystemCostIndices, error)

	// ----- Module: Insurance / 模块: Insurance -----
	GetPrices(ctx context.Context, ifNoneMatch ...string) ([]models.InsurancePrice, error)

	// ----- Module: Killmails / 模块: Killmails -----
	GetCharacterKillmails(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.KillmailRef, error)
	GetCorporationKillmails(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.KillmailRef, error)
	GetKillmail(ctx context.Context, killmailHash string, killmailID int32, ifNoneMatch ...string) (*models.Killmail, error)

	// ----- Module: Location / 模块: Location -----
	GetCharacterLocation(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterLocation, error)
	GetCharacterOnline(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.OnlineStatus, error)
	GetCharacterShip(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterShip, error)

	// ----- Module: Loyalty / 模块: Loyalty -----
	GetCharacterLoyaltyPoints(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.LoyaltyPoints, error)
	GetLoyaltyStoreOffers(ctx context.Context, corporationID int32, ifNoneMatch ...string) ([]models.LoyaltyStoreOffer, error)

	// ----- Module: Mail / 模块: Mail -----
	DeleteCharacterMailLabel(ctx context.Context, token string, characterID int32, labelID int32) error
	DeleteCharacterMail(ctx context.Context, token string, characterID int32, mailID int32) error
	GetCharacterMails(ctx context.Context, token string, characterID int32, labels []int32, lastMailID int32, ifNoneMatch ...string) ([]models.MailHeader, error)
	GetCharacterMailLabels(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.MailLabels, error)
	GetCharacterMailLists(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.MailingList, error)
	GetCharacterMail(ctx context.Context, token string, characterID int32, mailID int32, ifNoneMatch ...string) (*models.Mail, error)
	SendCharacterMail(ctx context.Context, token string, characterID int32, body *models.MailRequest) (int32, error)
	CreateCharacterMailLabel(ctx context.Context, token string, characterID int32, body *models.MailLabelRequest) (int32, error)
	UpdateCharacterMail(ctx context.Context, token string, characterID int32, mailID int32, body *models.MailMetadata) error

	// ----- Module: Market / 模块: Market -----
	GetCharacterMarketOrders(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.CharacterMarketOrder, error)
	GetCharacterMarketOrderHistory(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.CharacterMarketOrderHistory, error)
	GetCorporationMarketOrders(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationMarketOrder, error)
	GetCorporationMarketOrderHistory(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CorporationMarketOrderHistory, error)
	GetMarketGroups(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetMarketGroup(ctx context.Context, marketGroupID int32, ifNoneMatch ...string) (*models.MarketGroup, error)
	GetMarketPrices(ctx context.Context, ifNoneMatch ...string) ([]models.MarketPrice, error)
	GetMarketHistory(ctx context.Context, regionID int32, typeID string, ifNoneMatch ...string) ([]models.MarketHistoryEntry, error)
	GetMarketOrders(ctx context.Context, regionID int32, orderType string, page int32, typeID int32, ifNoneMatch ...string) ([]models.MarketOrder, error)
	GetMarketTypes(ctx context.Context, regionID int32, page int32, ifNoneMatch ...string) ([]int32, error)
	GetStructureMarketOrders(ctx context.Context, token string, structureID int64, page int32, ifNoneMatch ...string) ([]models.StructureMarketOrder, error)

	// ----- Module: Opportunities / 模块: Opportunities -----
	GetCharacterOpportunities(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.OpportunityCompletion, error)
	GetOpportunityGroups(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetOpportunityGroup(ctx context.Context, groupID int32, ifNoneMatch ...string) (*models.OpportunityGroup, error)
	GetOpportunityTasks(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetOpportunityTask(ctx context.Context, taskID int32, ifNoneMatch ...string) (*models.OpportunityTask, error)

	// ----- Module: PlanetaryInteraction / 模块: PlanetaryInteraction -----
	GetCharacterColonies(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.Colony, error)
	GetCharacterColonyLayout(ctx context.Context, token string, characterID int32, planetID int32, ifNoneMatch ...string) (*models.ColonyLayout, error)
	GetCorporationCustomsOffices(ctx context.Context, token string, corporationID int32, page int32, ifNoneMatch ...string) ([]models.CustomsOffice, error)
	GetSchematicInformation(ctx context.Context, schematicID int32, ifNoneMatch ...string) (*models.Schematic, error)

	// ----- Module: Routes / 模块: Routes -----
	GetRoute(ctx context.Context, destination int32, origin int32, avoid []int32, connections [][]int32, flag string, ifNoneMatch ...string) ([]int32, error)

	// ----- Module: Search / 模块: Search -----
	SearchEntities(ctx context.Context, token string, characterID int32, categories []string, search string, strict bool, ifNoneMatch ...string) (*models.SearchResult, error)

	// ----- Module: Skills / 模块: Skills -----
	GetCharacterAttributes(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterAttributes, error)
	GetCharacterSkillQueue(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) ([]models.SkillQueueEntry, error)
	GetCharacterSkills(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (*models.CharacterSkills, error)

	// ----- Module: Sovereignty / 模块: Sovereignty -----
	GetSovereigntyCampaigns(ctx context.Context, ifNoneMatch ...string) ([]models.SovereigntyCampaign, error)
	GetSovereigntyMap(ctx context.Context, ifNoneMatch ...string) ([]models.SovereigntySystem, error)
	GetSovereigntyStructures(ctx context.Context, ifNoneMatch ...string) ([]models.SovereigntyStructure, error)

	// ----- Module: Status / 模块: Status -----
	GetServerStatus(ctx context.Context, ifNoneMatch ...string) (*models.ServerStatus, error)

	// ----- Module: UserInterface / 模块: UserInterface -----
	SetAutopilotWaypoint(ctx context.Context, token string, addToBeginning bool, clearOtherWaypoints bool, destinationID int64) error
	OpenContractWindow(ctx context.Context, token string, contractID string) error
	OpenInformationWindow(ctx context.Context, token string, targetID string) error
	OpenMarketDetails(ctx context.Context, token string, typeID string) error
	OpenNewMailWindow(ctx context.Context, token string, body *models.NewMailRequest) error

	// ----- Module: Universe / 模块: Universe -----
	GetUniverseAncestries(ctx context.Context, ifNoneMatch ...string) ([]models.UniverseAncestry, error)
	GetUniverseAsteroidBelt(ctx context.Context, asteroidBeltID int32, ifNoneMatch ...string) (*models.UniverseAsteroidBelt, error)
	GetUniverseBloodlines(ctx context.Context, ifNoneMatch ...string) ([]models.UniverseBloodline, error)
	GetUniverseCategories(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetUniverseCategory(ctx context.Context, categoryID int32, ifNoneMatch ...string) (*models.UniverseCategory, error)
	GetConstellations(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetConstellationInformation(ctx context.Context, constellationID int32, ifNoneMatch ...string) (*models.UniverseConstellation, error)
	GetUniverseFactions(ctx context.Context, ifNoneMatch ...string) ([]models.UniverseFaction, error)
	GetUniverseGraphics(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetUniverseGraphic(ctx context.Context, graphicID int32, ifNoneMatch ...string) (*models.UniverseGraphic, error)
	GetUniverseGroups(ctx context.Context, page int32, ifNoneMatch ...string) ([]int32, error)
	GetUniverseGroup(ctx context.Context, groupID int32, ifNoneMatch ...string) (*models.UniverseGroup, error)
	GetUniverseMoon(ctx context.Context, moonID int32, ifNoneMatch ...string) (*models.UniverseMoon, error)
	GetUniversePlanet(ctx context.Context, planetID int32, ifNoneMatch ...string) (*models.UniversePlanet, error)
	GetUniverseRaces(ctx context.Context, ifNoneMatch ...string) ([]models.UniverseRace, error)
	GetUniverseRegions(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetUniverseRegion(ctx context.Context, regionID int32, ifNoneMatch ...string) (*models.UniverseRegion, error)
	GetUniverseStargate(ctx context.Context, stargateID int32, ifNoneMatch ...string) (*models.UniverseStargate, error)
	GetUniverseStar(ctx context.Context, starID int32, ifNoneMatch ...string) (*models.UniverseStar, error)
	GetUniverseStation(ctx context.Context, stationID int32, ifNoneMatch ...string) (*models.UniverseStation, error)
	GetPublicStructures(ctx context.Context, filter string, ifNoneMatch ...string) ([]int64, error)
	GetUniverseStructure(ctx context.Context, token string, structureID int64, ifNoneMatch ...string) (*models.UniverseStructure, error)
	GetUniverseSystemJumps(ctx context.Context, ifNoneMatch ...string) ([]models.UniverseSystemJump, error)
	GetUniverseSystemKills(ctx context.Context, ifNoneMatch ...string) ([]models.UniverseSystemKills, error)
	GetUniverseSystems(ctx context.Context, ifNoneMatch ...string) ([]int32, error)
	GetUniverseSystem(ctx context.Context, systemID int32, ifNoneMatch ...string) (*models.UniverseSolarSystem, error)
	GetUniverseTypes(ctx context.Context, page int32, ifNoneMatch ...string) ([]int32, error)
	GetUniverseType(ctx context.Context, typeID int32, ifNoneMatch ...string) (*models.UniverseType, error)
	ResolveNamesToIDs(ctx context.Context, body []string) (*models.ResolvedIds, error)
	ResolveIDsToNames(ctx context.Context, body []int32) ([]models.UniverseName, error)

	// ----- Module: Wallet / 模块: Wallet -----
	GetCharacterWalletBalance(ctx context.Context, token string, characterID int32, ifNoneMatch ...string) (float64, error)
	GetCharacterWalletJournal(ctx context.Context, token string, characterID int32, page int32, ifNoneMatch ...string) ([]models.WalletJournalEntry, error)
	GetWalletTransactions(ctx context.Context, token string, characterID int32, fromID int64, ifNoneMatch ...string) ([]models.CharacterWalletTransaction, error)
	GetCorporationWallets(ctx context.Context, token string, corporationID int32, ifNoneMatch ...string) ([]models.CorporationWallet, error)
	GetCorporationWalletJournal(ctx context.Context, token string, corporationID int32, division int32, page int32, ifNoneMatch ...string) ([]models.WalletJournalEntry, error)
	GetCorporationWalletTransactions(ctx context.Context, token string, corporationID int32, division int32, fromID int64, ifNoneMatch ...string) ([]models.CorporationWalletTransaction, error)

	// ----- Module: Wars / 模块: Wars -----
	GetWar(ctx context.Context, warID int32, ifNoneMatch ...string) (*models.War, error)
	GetWarKillmails(ctx context.Context, warID int32, page int32, ifNoneMatch ...string) ([]models.KillmailRef, error)
	GetWars(ctx context.Context, maxWarID int32, ifNoneMatch ...string) ([]int32, error)
}

var _ ClientIface = (*Client)(nil)
