package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UpdateResourceRequest struct {
	ShardID     string            `json:"shard_id" validate:"required"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Spec        json.RawMessage   `json:"spec"`
}

// updateResource обновляет ресурс
// @Summary Update a resource
// @Description Updates a resource with the provided details.
// @Tags resources
// @Accept json
// @Produce json
// @Param resource_group path string true "Resource Group" example="group1"
// @Param kind path string true "Kind" example="type1"
// @Param namespace path string true "Namespace" example="ns1"
// @Param name path string true "Name" example="resource1"
// @Param resource body UpdateResourceRequest{spec=map[string]interface{}} true "Resource details"
// @Success 200 {object} dto.Resource{spec=map[string]interface{},status=map[string]interface{}} "Resource updated"
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Validation failed: shard_id is required"}
// @Failure 500 {object} ErrorResponse "Server error" example={"error":"Failed to update resource: database error"}
// @Router /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name} [put]
func (h *Handler) updateResource(w http.ResponseWriter, r *http.Request) {
	var req UpdateResourceRequest
	if err := jsonIter.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON: %h"}`, err), http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Validation failed: %h"}`, err), http.StatusBadRequest)
		return
	}

	resourceUpdateOpts := &dto.ResourceUpdateOpts{
		ResourceFields: dto.ResourceFields{
			ResourceID: dto.ResourceID{
				ResourceGroup: chi.URLParam(r, "resource_group"),
				Kind:          chi.URLParam(r, "kind"),
				Namespace:     chi.URLParam(r, "namespace"),
				Name:          chi.URLParam(r, "name"),
			},
			ShardID:     req.ShardID,
			Labels:      req.Labels,
			Annotations: req.Annotations,
		},
		Spec: req.Spec,
	}

	resource, err := h.service.Update(r.Context(), resourceUpdateOpts)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Failed to update resource: %h"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonIter.NewEncoder(w).Encode(resource)
}

type UpdateResourceStatusRequest struct {
	ShardID        string            `json:"shard_id" validate:"required"`
	Labels         map[string]string `json:"labels"`
	Annotations    map[string]string `json:"annotations"`
	Status         json.RawMessage   `json:"status"`
	Version        int               `json:"version"`
	CurrentVersion int               `json:"current_version"`
}

// updateResource обновляет ресурс
// @Summary Update a resource
// @Description Updates a resource with the provided details.
// @Tags resources
// @Accept json
// @Produce json
// @Param resource_group path string true "Resource Group" example="group1"
// @Param kind path string true "Kind" example="type1"
// @Param namespace path string true "Namespace" example="ns1"
// @Param name path string true "Name" example="resource1"
// @Param resource body UpdateResourceStatusRequest{status=map[string]interface{}} true "Resource details"
// @Success 200 {object} dto.Resource{spec=map[string]interface{},status=map[string]interface{}}  "Resource updated"
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Validation failed: shard_id is required"}
// @Failure 409 {object} ErrorResponse "Invalid input" example={"error":"Version conflict: resource version not match"}
// @Failure 500 {object} ErrorResponse "Server error" example={"error":"Failed to update resource: database error"}
// @Router /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name}/status [put]
func (h *Handler) updateResourceStatus(w http.ResponseWriter, r *http.Request) {
	var req UpdateResourceStatusRequest
	if err := jsonIter.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON: %s"}`, err), http.StatusBadRequest)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Validation failed: %s"}`, err), http.StatusBadRequest)
		return
	}
	resourceUpdateStatusOpts := &dto.ResourceUpdateStatusOpts{
		ResourceFields: dto.ResourceFields{
			ResourceID: dto.ResourceID{
				ResourceGroup: chi.URLParam(r, "resource_group"),
				Kind:          chi.URLParam(r, "kind"),
				Namespace:     chi.URLParam(r, "namespace"),
				Name:          chi.URLParam(r, "name"),
			},
			ShardID:     req.ShardID,
			Labels:      req.Labels,
			Annotations: req.Annotations,
		},
		Status:         req.Status,
		Version:        req.Version,
		CurrentVersion: req.CurrentVersion,
	}

	resource, err := h.service.UpdateStatus(r.Context(), resourceUpdateStatusOpts)
	if err != nil {
		if errors.Is(err, dto.ConflictError) {
			http.Error(w, fmt.Sprintf(`{"error":"Version conflict: %s"}`, err), http.StatusConflict)
			return
		}
		if errors.Is(err, dto.UnavailableVersionError) {
			http.Error(w, fmt.Sprintf(`{"error":"Current version unavailable: %s"}`, err), http.StatusBadRequest)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"Failed to update resource: %s"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonIter.NewEncoder(w).Encode(resource)
}
