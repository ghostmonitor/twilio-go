/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TaskrouterV1TaskQueuesStatistics struct for TaskrouterV1TaskQueuesStatistics
type TaskrouterV1TaskQueuesStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the TaskQueues
	Cumulative *map[string]interface{} `json:"cumulative,omitempty"`
	// An object that contains the real-time statistics for the TaskQueues
	Realtime *map[string]interface{} `json:"realtime,omitempty"`
	// The SID of the TaskQueue from which these statistics were calculated
	TaskQueueSid *string `json:"task_queue_sid,omitempty"`
	// The SID of the Workspace that contains the TaskQueues
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
