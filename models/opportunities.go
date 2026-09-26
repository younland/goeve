package models

import (
	"time"
)

// OpportunityCompletion 200 ok object.
// OpportunityCompletion 200 ok 对象.
type OpportunityCompletion struct {
	// CompletedAt completed_at string.
	// CompletedAt 完成时间字符串.
	CompletedAt time.Time `json:"completed_at"`
	// TaskId task_id integer.
	// TaskId task_id 整数.
	TaskId int32 `json:"task_id"`
}

// OpportunityGroup 200 ok object.
// OpportunityGroup 200 ok 对象.
type OpportunityGroup struct {
	// ConnectedGroups The groups that are connected to this group on the opportunities map.
	// ConnectedGroups 机遇地图上与此组相连的其他组.
	ConnectedGroups []int32 `json:"connected_groups"`
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// GroupId group_id integer.
	// GroupId 分组 ID 整数.
	GroupId int32 `json:"group_id"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Notification notification string.
	// Notification notification 字符串.
	Notification string `json:"notification"`
	// RequiredTasks Tasks need to complete for this group.
	// RequiredTasks 完成该组所需的任务.
	RequiredTasks []int32 `json:"required_tasks"`
}

// OpportunityTask 200 ok object.
// OpportunityTask 200 ok 对象.
type OpportunityTask struct {
	// Description description string.
	// Description 描述字符串.
	Description string `json:"description"`
	// Name name string.
	// Name name 字符串.
	Name string `json:"name"`
	// Notification notification string.
	// Notification notification 字符串.
	Notification string `json:"notification"`
	// TaskId task_id integer.
	// TaskId task_id 整数.
	TaskId int32 `json:"task_id"`
}
