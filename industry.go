package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdIndustryJobs List character industry jobs.
// GetCharactersCharacterIdIndustryJobs 列出角色工业工作.
//
// Route: GET /characters/{character_id}/industry/jobs/ — This route is cached for up to 300 seconds
// 路由: GET /characters/{character_id}/industry/jobs/ — 该路由缓存长达 300 秒
// Scopes: esi-industry.read_character_jobs.v1
// 权限: esi-industry.read_character_jobs.v1
func (c *Client) GetCharactersCharacterIdIndustryJobs(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdIndustryJobsParams) ([]models.GetCharactersCharacterIdIndustryJobs, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdIndustryJobs
	err := c.get(ctx, "/characters/{character_id}/industry/jobs/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCharactersCharacterIdMining Character mining ledger.
// GetCharactersCharacterIdMining 角色采矿记录.
//
// Route: GET /characters/{character_id}/mining/ — This route is cached for up to 600 seconds
// 路由: GET /characters/{character_id}/mining/ — 该路由缓存长达 600 秒
// Scopes: esi-industry.read_character_mining.v1
// 权限: esi-industry.read_character_mining.v1
func (c *Client) GetCharactersCharacterIdMining(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdMiningParams) ([]models.GetCharactersCharacterIdMining, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdMining
	err := c.get(ctx, "/characters/{character_id}/mining/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationCorporationIdMiningExtractions Moon extraction timers.
// GetCorporationCorporationIdMiningExtractions 卫星开采计时.
//
// Route: GET /corporation/{corporation_id}/mining/extractions/ — This route is cached for up to 1800 seconds
// 路由: GET /corporation/{corporation_id}/mining/extractions/ — 该路由缓存长达 1800 秒
// Scopes: esi-industry.read_corporation_mining.v1
// 权限: esi-industry.read_corporation_mining.v1
func (c *Client) GetCorporationCorporationIdMiningExtractions(ctx context.Context, corporationId int32, params *models.GetCorporationCorporationIdMiningExtractionsParams) ([]models.GetCorporationCorporationIdMiningExtractions, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationCorporationIdMiningExtractions
	err := c.get(ctx, "/corporation/{corporation_id}/mining/extractions/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationCorporationIdMiningObservers Corporation mining observers.
// GetCorporationCorporationIdMiningObservers 军团采矿监察员.
//
// Route: GET /corporation/{corporation_id}/mining/observers/ — This route is cached for up to 3600 seconds
// 路由: GET /corporation/{corporation_id}/mining/observers/ — 该路由缓存长达 3600 秒
// Scopes: esi-industry.read_corporation_mining.v1
// 权限: esi-industry.read_corporation_mining.v1
func (c *Client) GetCorporationCorporationIdMiningObservers(ctx context.Context, corporationId int32, params *models.GetCorporationCorporationIdMiningObserversParams) ([]models.GetCorporationCorporationIdMiningObservers, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationCorporationIdMiningObservers
	err := c.get(ctx, "/corporation/{corporation_id}/mining/observers/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationCorporationIdMiningObserversObserverId Observed corporation mining.
// GetCorporationCorporationIdMiningObserversObserverId 观测到的军团开采活动.
//
// Route: GET /corporation/{corporation_id}/mining/observers/{observer_id}/ — This route is cached for up to 3600 seconds
// 路由: GET /corporation/{corporation_id}/mining/observers/{observer_id}/ — 该路由缓存长达 3600 秒
// Scopes: esi-industry.read_corporation_mining.v1
// 权限: esi-industry.read_corporation_mining.v1
func (c *Client) GetCorporationCorporationIdMiningObserversObserverId(ctx context.Context, corporationId int32, observerId int64, params *models.GetCorporationCorporationIdMiningObserversObserverIdParams) ([]models.GetCorporationCorporationIdMiningObserversObserverId, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10), "observer_id": strconv.FormatInt(int64(observerId), 10)}
	var result []models.GetCorporationCorporationIdMiningObserversObserverId
	err := c.get(ctx, "/corporation/{corporation_id}/mining/observers/{observer_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCorporationsCorporationIdIndustryJobs List corporation industry jobs.
// GetCorporationsCorporationIdIndustryJobs 列出军团工业工作.
//
// Route: GET /corporations/{corporation_id}/industry/jobs/ — This route is cached for up to 300 seconds
// 路由: GET /corporations/{corporation_id}/industry/jobs/ — 该路由缓存长达 300 秒
// Scopes: esi-industry.read_corporation_jobs.v1
// 权限: esi-industry.read_corporation_jobs.v1
func (c *Client) GetCorporationsCorporationIdIndustryJobs(ctx context.Context, corporationId int32, params *models.GetCorporationsCorporationIdIndustryJobsParams) ([]models.GetCorporationsCorporationIdIndustryJobs, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"corporation_id": strconv.FormatInt(int64(corporationId), 10)}
	var result []models.GetCorporationsCorporationIdIndustryJobs
	err := c.get(ctx, "/corporations/{corporation_id}/industry/jobs/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFacilities List industry facilities.
// GetFacilities 列出工业设施.
//
// Route: GET /industry/facilities/ — This route is cached for up to 3600 seconds
// 路由: GET /industry/facilities/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetFacilities(ctx context.Context, params *models.GetFacilitiesParams) ([]models.GetIndustryFacilities, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetIndustryFacilities
	err := c.get(ctx, "/industry/facilities/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetIndustrySystems List solar system cost indices.
// GetIndustrySystems 列出星系成本指数.
//
// Route: GET /industry/systems/ — This route is cached for up to 3600 seconds
// 路由: GET /industry/systems/ — 该路由缓存长达 3600 秒
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetIndustrySystems(ctx context.Context, params *models.GetSystemsParams) ([]models.GetIndustrySystems, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []models.GetIndustrySystems
	err := c.get(ctx, "/industry/systems/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
