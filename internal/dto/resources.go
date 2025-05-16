package dto

import (
	"encoding/json"
	"time"
)

type ResourceID struct {
	ResourceGroup string `json:"resource_group"`
	Kind          string `json:"kind"`
	Namespace     string `json:"namespace"`
	Name          string `json:"name"`
}

type ResourceFields struct {
	ResourceID  `json:",inline"`
	ShardID     string            `json:"shard_id" validate:"required"`
	Finalizers  []string          `json:"finalizers"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}

type Resource struct {
	ResourceFields    `json:",inline"`
	ID                int             `json:"-"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletionTimestamp *time.Time      `json:"deletion_timestamp"`
	Version           int             `json:"version"`
	CurrentVersion    int             `json:"current_version"`
	Spec              json.RawMessage `json:"spec"`
	Status            json.RawMessage `json:"status"`
}

type ResourceCreateOpts struct {
	ResourceFields `json:",inline"`
	Spec           json.RawMessage `json:"spec"`
}

type ResourceUpdateOpts struct {
	ResourceFields `json:",inline"`
	Spec           json.RawMessage `json:"spec"`
}

type ResourceUpdateStatusOpts struct {
	ResourceFields `json:",inline"`
	Status         json.RawMessage `json:"status"`
	Version        int             `json:"version"`
	CurrentVersion int             `json:"current_version"`
}

type ListResourcesOpts struct {
	ResourceID
	ShardID string `json:"shard_id"`
	Pending bool   `json:"pending"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
