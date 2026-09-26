package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"net/url"
	"strconv"
)

// GetCharacterIndustryJobs List character industry jobs.
// GetCharacterIndustryJobs 列出角色工业工作.
//
// Route: GET /characters/{character_id}/industry/jobs/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/industry/jobs/ — 该路由缓存长达 300 秒
// Scopes: esi-industry.read_character_jobs.v1
// 权限: esi-industry.read_character_jobs.v1
func (c *Client) GetCharacterIndustryJobs(ctx context.Context, characterID int32, token string, includeCompleted bool, ifNoneMatch string) ([]models.CharacterIndustryJob, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if includeCompleted {
		query.Set("include_completed", strconv.FormatBool(includeCompleted))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.CharacterIndustryJob
	err := c.get(ctx, "/characters/{character_id}/industry/jobs/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharacterMiningLedger Character mining ledger.
// GetCharacterMiningLedger 角色采矿记录.
//
// Route: GET /characters/{character_id}/mining/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/mining/ — 该路由缓存长达 600 秒
// Scopes: esi-industry.read_character_mining.v1
// 权限: esi-industry.read_character_mining.v1
func (c *Client) GetCharacterMiningLedger(ctx context.Context, characterID int32, token string, page int32, ifNoneMatch string) ([]models.MiningLedgerEntry, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.MiningLedgerEntry
	err := c.get(ctx, "/characters/{character_id}/mining/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMoonExtractions Moon extraction timers.
// GetCorporationMoonExtractions 卫星开采计时.
//
// Route: GET /corporation/{corporation_id}/mining/extractions/ — This route is cached for up to 1800 seconds
// 路由: GET /corporation/{corporation_id}/mining/extractions/ — 该路由缓存长达 1800 秒
// Scopes: esi-industry.read_corporation_mining.v1
// 权限: esi-industry.read_corporation_mining.v1
func (c *Client) GetCorporationMoonExtractions(ctx context.Context, corporationID int32, token string, page int32, ifNoneMatch string) ([]models.MoonExtraction, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.MoonExtraction
	err := c.get(ctx, "/corporation/{corporation_id}/mining/extractions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMiningObservers Corporation mining observers.
// GetCorporationMiningObservers 军团采矿监察员.
//
// Route: GET /corporation/{corporation_id}/mining/observers/ — This route is cached for up to 3600 seconds
// 路由: GET /corporation/{corporation_id}/mining/observers/ — 该路由缓存长达 3600 秒
// Scopes: esi-industry.read_corporation_mining.v1
// 权限: esi-industry.read_corporation_mining.v1
func (c *Client) GetCorporationMiningObservers(ctx context.Context, corporationID int32, token string, page int32, ifNoneMatch string) ([]models.MiningObserver, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.MiningObserver
	err := c.get(ctx, "/corporation/{corporation_id}/mining/observers/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationMiningObserverData Observed corporation mining.
// GetCorporationMiningObserverData 观测到的军团开采活动.
//
// Route: GET /corporation/{corporation_id}/mining/observers/{observer_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /corporation/{corporation_id}/mining/observers/{observer_id}/ — 该路由缓存长达 3600 秒
// Scopes: esi-industry.read_corporation_mining.v1
// 权限: esi-industry.read_corporation_mining.v1
func (c *Client) GetCorporationMiningObserverData(ctx context.Context, corporationID int32, observerID int64, token string, page int32, ifNoneMatch string) ([]models.MiningObserverEntry, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10), "observer_id": strconv.FormatInt(int64(observerID), 10)}
	var result []models.MiningObserverEntry
	err := c.get(ctx, "/corporation/{corporation_id}/mining/observers/{observer_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationIndustryJobs List corporation industry jobs.
// GetCorporationIndustryJobs 列出军团工业工作.
//
// Route: GET /corporations/{corporation_id}/industry/jobs/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/industry/jobs/ — 该路由缓存长达 300 秒
// Scopes: esi-industry.read_corporation_jobs.v1
// 权限: esi-industry.read_corporation_jobs.v1
func (c *Client) GetCorporationIndustryJobs(ctx context.Context, corporationID int32, token string, page int32, includeCompleted bool, ifNoneMatch string) ([]models.CorporationIndustryJob, error) {
	query := url.Values{}
	headers := map[string]string{}
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}
	if page > 0 {
		query.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if includeCompleted {
		query.Set("include_completed", strconv.FormatBool(includeCompleted))
	}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationID), 10)}
	var result []models.CorporationIndustryJob
	err := c.get(ctx, "/corporations/{corporation_id}/industry/jobs/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIndustryFacilities List industry facilities.
// GetIndustryFacilities 列出工业设施.
//
// Route: GET /industry/facilities/ — This route is cached for up to 3600 seconds
// 路由: GET /industry/facilities/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetIndustryFacilities(ctx context.Context, ifNoneMatch string) ([]models.IndustryFacility, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.IndustryFacility
	err := c.get(ctx, "/industry/facilities/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIndustrySystemCostIndices List solar system cost indices.
// GetIndustrySystemCostIndices 列出星系成本指数.
//
// Route: GET /industry/systems/ — This route is cached for up to 3600 seconds
// 路由: GET /industry/systems/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetIndustrySystemCostIndices(ctx context.Context, ifNoneMatch string) ([]models.IndustrySystemCostIndices, error) {
	query := url.Values{}
	headers := map[string]string{}
	if ifNoneMatch != "" {
		headers["If-None-Match"] = ifNoneMatch
	}
	var pathParams map[string]string
	var result []models.IndustrySystemCostIndices
	err := c.get(ctx, "/industry/systems/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
