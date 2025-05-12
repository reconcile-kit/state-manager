package dto

import (
	"encoding/json"
	"time"
)

type ResourceFields struct {
	ResourceID `json:",inline"`
	ShardID    string            `json:"shard_id" validate:"required"`
	Labels     map[string]string `json:"labels"`
	Body       json.RawMessage   `json:"body"`
}

type ResourceCreateOpts struct {
	ResourceFields `json:",inline"`
}

type ResourceUpdateOpts struct {
	ResourceFields   `json:",inline"`
	WithoutUpVersion bool `json:"without_up_version"`
	CurrentVersion   int  `json:"current_version"`
}

type Resource struct {
	ID             int       `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Version        int       `json:"version"`
	CurrentVersion int       `json:"current_version"`
	ResourceFields `json:",inline"`
}

type ResourceID struct {
	ResourceGroup string `json:"resource_group"`
	Kind          string `json:"kind"`
	Namespace     string `json:"namespace"`
	Name          string `json:"name"`
}
