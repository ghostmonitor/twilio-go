/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// TaskrouterV1TaskQueueBulkRealTimeStatistics struct for TaskrouterV1TaskQueueBulkRealTimeStatistics
type TaskrouterV1TaskQueueBulkRealTimeStatistics struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Workspace that contains the TaskQueue.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
	// The TaskQueue RealTime Statistics for the requested list of task_queue_sid, grouped by task_queue_sid.
	TaskQueueData *interface{} `json:"task_queue_data,omitempty"`
	// The number of TaskQueue statistics received in task_queue_data.
	TaskQueueResponseCount *int `json:"task_queue_response_count,omitempty"`
	// The absolute URL of the TaskQueue statistics resource.
	Url *string `json:"url,omitempty"`
}