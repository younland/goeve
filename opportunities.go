package goeve

import (
	"context"
	"github.com/younland/goeve/models"
	"strconv"
)

// GetCharacterOpportunities Get a character's completed tasks.
// GetCharacterOpportunities 获取角色已完成的任务.
//
// Route: GET /characters/{character_id}/opportunities/ — This route is cached for up to 3600 seconds
// 路由: GET /characters/{character_id}/opportunities/ — 该路由缓存长达 3600 秒
// Scopes: esi-characters.read_opportunities.v1
// 权限: esi-characters.read_opportunities.v1
func (c *Client) GetCharacterOpportunities(ctx context.Context, characterID int32, opts ...RequestOption) ([]models.OpportunityCompletion, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"character_id": strconv.FormatInt(int64(characterID), 10)}
	var result []models.OpportunityCompletion
	err := c.get(ctx, "/characters/{character_id}/opportunities/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpportunityGroups Get opportunities groups.
// GetOpportunityGroups 获取机遇组列表.
//
// Route: GET /opportunities/groups/
// 路由: GET /opportunities/groups/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetOpportunityGroups(ctx context.Context, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/opportunities/groups/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpportunityGroup Get opportunities group.
// GetOpportunityGroup 获取机遇组.
//
// Route: GET /opportunities/groups/{group_id}/
// 路由: GET /opportunities/groups/{group_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetOpportunityGroup(ctx context.Context, groupID int32, opts ...RequestOption) (*models.OpportunityGroup, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"group_id": strconv.FormatInt(int64(groupID), 10)}
	var result *models.OpportunityGroup
	err := c.get(ctx, "/opportunities/groups/{group_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpportunityTasks Get opportunities tasks.
// GetOpportunityTasks 获取机遇任务列表.
//
// Route: GET /opportunities/tasks/
// 路由: GET /opportunities/tasks/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetOpportunityTasks(ctx context.Context, opts ...RequestOption) ([]int32, error) {
	query, headers := newRequestOptions(opts...)
	var pathParams map[string]string
	var result []int32
	err := c.get(ctx, "/opportunities/tasks/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOpportunityTask Get opportunities task.
// GetOpportunityTask 获取机遇任务.
//
// Route: GET /opportunities/tasks/{task_id}/
// 路由: GET /opportunities/tasks/{task_id}/
// Scopes: none (public endpoint)
// 权限: 无（公开接口）
func (c *Client) GetOpportunityTask(ctx context.Context, taskID int32, opts ...RequestOption) (*models.OpportunityTask, error) {
	query, headers := newRequestOptions(opts...)
	pathParams := map[string]string{"task_id": strconv.FormatInt(int64(taskID), 10)}
	var result *models.OpportunityTask
	err := c.get(ctx, "/opportunities/tasks/{task_id}/", pathParams, query, headers, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
