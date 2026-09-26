package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharactersCharacterIdOpportunities Get a character's completed tasks.
// GetCharactersCharacterIdOpportunities 获取角色已完成的任务.
//
// Route: GET /characters/{character_id}/opportunities/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/opportunities/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_opportunities.v1
// 权限: esi-characters.read_opportunities.v1
func (c *Client) GetCharactersCharacterIdOpportunities(ctx context.Context, characterId int32, params *models.GetCharactersCharacterIdOpportunitiesParams) ([]models.GetCharactersCharacterIdOpportunities, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterId), 10)}
	var result []models.GetCharactersCharacterIdOpportunities
	err := c.get(ctx, "/characters/{character_id}/opportunities/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpportunitiesGroups Get opportunities groups.
// GetOpportunitiesGroups 获取机遇组列表.
//
// Route: GET /opportunities/groups/
// 路由: GET /opportunities/groups/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetOpportunitiesGroups(ctx context.Context, params *models.GetGroupsParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/opportunities/groups/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpportunitiesGroupsGroupId Get opportunities group.
// GetOpportunitiesGroupsGroupId 获取机遇组.
//
// Route: GET /opportunities/groups/{group_id}/
// 路由: GET /opportunities/groups/{group_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetOpportunitiesGroupsGroupId(ctx context.Context, groupId int32, params *models.GetGroupsGroupIdParams) (*models.GetOpportunitiesGroupsGroupIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"group_id": strconv.FormatInt(int64(groupId), 10)}
	var result *models.GetOpportunitiesGroupsGroupIdOk
	err := c.get(ctx, "/opportunities/groups/{group_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTasks Get opportunities tasks.
// GetTasks 获取机遇任务列表.
//
// Route: GET /opportunities/tasks/
// 路由: GET /opportunities/tasks/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetTasks(ctx context.Context, params *models.GetTasksParams) ([]int32, error) {
	query, headers := params.Values()
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/opportunities/tasks/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetTasksTaskId Get opportunities task.
// GetTasksTaskId 获取机遇任务.
//
// Route: GET /opportunities/tasks/{task_id}/
// 路由: GET /opportunities/tasks/{task_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetTasksTaskId(ctx context.Context, taskId int32, params *models.GetTasksTaskIdParams) (*models.GetOpportunitiesTasksTaskIdOk, error) {
	query, headers := params.Values()
	pathParams := map[string]string{"task_id": strconv.FormatInt(int64(taskId), 10)}
	var result *models.GetOpportunitiesTasksTaskIdOk
	err := c.get(ctx, "/opportunities/tasks/{task_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
