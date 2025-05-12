package http

import (
	"fmt"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

// getResource получает ресурс по ключу
// @Summary Get a resource by key
// @Description Retrieves a resource by its resource_group, kind, namespace, and name.
// @Tags resources
// @Accept json
// @Produce json
// @Param resource_group path string true "Resource Group" example="group1"
// @Param kind path string true "Kind" example="type1"
// @Param namespace path string true "Namespace" example="ns1"
// @Param name path string true "Name" example="resource1"
// @Success 200 {object} dto.Resource{body=map[string]string} "Resource found" example={"id":1,"resource_group":"group1","kind":"type1","namespace":"ns1","name":"resource1","shard_id":"default","body":{"key":"value"},"labels":{"env":"prod"},"created_at":"2025-05-12T00:00:00Z","updated_at":"2025-05-12T00:00:00Z","version":1,"current_version":0}
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Validation failed: resource_group is required"}
// @Failure 404 {object} ErrorResponse "Not found" example={"error":"Resource not found: no rows"}
// @Router /api/v1/resources/{resource_group}/{kind}/{namespace}/{name} [get]
func (h *Handler) getResource(w http.ResponseWriter, r *http.Request) {
	opts := dto.ResourceID{
		ResourceGroup: chi.URLParam(r, "resource_group"),
		Kind:          chi.URLParam(r, "kind"),
		Namespace:     chi.URLParam(r, "namespace"),
		Name:          chi.URLParam(r, "name"),
	}

	if err := h.validator.Struct(&opts); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Validation failed: %h"}`, err), http.StatusBadRequest)
		return
	}

	resource, err := h.service.GetByKey(r.Context(), opts)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Failed to get resource: %h"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource)
}

// listPending возвращает ресурсы с version > current_version
// @Summary List pending resources
// @Description Lists resources where version > current_version for given shard IDs.
// @Tags resources
// @Accept json
// @Produce json
// @Param shard_ids query string true "Comma-separated list of shard IDs" example="default,shard2"
// @Success 200 {array} dto.Resource{body=map[string]string} "List of pending resources"
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"shard_ids is required"}
// @Failure 500 {object} ErrorResponse "Server error" example={"error":"Failed to list pending resources: database error"}
// @Router /api/v1/resources/pending [get]
func (h *Handler) listPending(w http.ResponseWriter, r *http.Request) {
	shardIDsStr := r.URL.Query().Get("shard_ids")
	if shardIDsStr == "" {
		http.Error(w, fmt.Sprintf(`{"error":"Validation failed: shard_ids is required"}`), http.StatusBadRequest)
		return
	}

	shardIDs := strings.Split(shardIDsStr, ",")
	for i, id := range shardIDs {
		shardIDs[i] = strings.TrimSpace(id)
		if shardIDs[i] == "" {
			http.Error(w, fmt.Sprintf(`{"error":"Validation failed: invalid shard_id"}`), http.StatusBadRequest)
			return
		}
	}

	resources, err := h.service.ListPending(r.Context(), shardIDs)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Failed to get resources: %h"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resources)
}
