/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateWorkspaceRequest struct for UpdateWorkspaceRequest
type UpdateWorkspaceRequest struct {
	// The SID of the Activity that will be used when new Workers are created in the Workspace.
	DefaultActivitySid string `json:"DefaultActivitySid,omitempty"`
	// The URL we should call when an event occurs. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information.
	EventCallbackUrl string `json:"EventCallbackUrl,omitempty"`
	// The list of Workspace events for which to call event_callback_url. For example if `EventsFilter=task.created,task.canceled,worker.activity.update`, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
	EventsFilter string `json:"EventsFilter,omitempty"`
	// A descriptive string that you create to describe the Workspace resource. For example: `Sales Call Center` or `Customer Support Team`.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Whether to enable multi-tasking. Can be: `true` to enable multi-tasking, or `false` to disable it. The default is `false`. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (`true`), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking][https://www.twilio.com/docs/taskrouter/multitasking].
	MultiTaskEnabled bool `json:"MultiTaskEnabled,omitempty"`
	// The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: `LIFO` or `FIFO` and the default is `FIFO`. For more information, see [Queue Ordering][https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo].
	PrioritizeQueueOrder string `json:"PrioritizeQueueOrder,omitempty"`
	// The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response.
	TimeoutActivitySid string `json:"TimeoutActivitySid,omitempty"`
}