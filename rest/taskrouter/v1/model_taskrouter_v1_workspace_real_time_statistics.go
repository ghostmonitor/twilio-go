/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1WorkspaceRealTimeStatistics struct for TaskrouterV1WorkspaceRealTimeStatistics
type TaskrouterV1WorkspaceRealTimeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The number of current Workers by Activity
	ActivityStatistics *[]interface{} `json:"activity_statistics,omitempty"`
	// The age of the longest waiting Task
	LongestTaskWaitingAge *int `json:"longest_task_waiting_age,omitempty"`
	// The SID of the longest waiting Task
	LongestTaskWaitingSid *string `json:"longest_task_waiting_sid,omitempty"`
	// The number of Tasks by priority
	TasksByPriority *interface{} `json:"tasks_by_priority,omitempty"`
	// The number of Tasks by their current status
	TasksByStatus *interface{} `json:"tasks_by_status,omitempty"`
	// The total number of Tasks
	TotalTasks *int `json:"total_tasks,omitempty"`
	// The total number of Workers in the Workspace
	TotalWorkers *int `json:"total_workers,omitempty"`
	// The absolute URL of the Workspace statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
