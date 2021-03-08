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

// ListWorkerResponse struct for ListWorkerResponse
type ListWorkerResponse struct {
	Meta    ListWorkspaceResponseMeta     `json:"Meta,omitempty"`
	Workers []TaskrouterV1WorkspaceWorker `json:"Workers,omitempty"`
}