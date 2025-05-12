package http

import (
	"fmt"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// createResource создаёт новый ресурс
// @Summary Create a new resource
// @Description Creates a new resource with the provided details.
// @Tags resources
// @Accept json
// @Produce json
// @Param resource_group path string true "Resource Group" example="group1"
// @Param kind path string true "Kind" example="type1"
// @Param namespace path string true "Namespace" example="ns1"
// @Param resource body dto.ResourceCreateOpts{body=map[string]string} true "Resource details" example={"resource_group":"group1","kind":"type1","namespace":"ns1","name":"resource1","shard_id":"default","body":{"key":"value"},"labels":{"env":"prod"}}
// @Success 201 {object} dto.Resource{body=map[string]string} "Resource created" example={"id":1,"resource_group":"group1","kind":"type1","namespace":"ns1","name":"resource1","shard_id":"default","body":{"key":"value"},"labels":{"env":"prod"},"created_at":"2025-05-12T00:00:00Z","updated_at":"2025-05-12T00:00:00Z","version":1,"current_version":0}
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Validation failed: shard_id is required"}
// @Failure 500 {object} ErrorResponse "Server error" example={"error":"Failed to create resource: database error"}
// @Router /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources [post]
func (h *Handler) createResource(w http.ResponseWriter, r *http.Request) {
	var req dto.ResourceCreateOpts
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON: %h"}`, err), http.StatusBadRequest)
		return
	}
	req.ResourceGroup = chi.URLParam(r, "resource_group")
	req.Kind = chi.URLParam(r, "kind")
	req.Namespace = chi.URLParam(r, "namespace")

	if err := h.validator.Struct(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Validation failed: %h"}`, err), http.StatusBadRequest)
		return
	}

	resource, err := h.service.Create(r.Context(), req)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Failed to create resource: %h"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}
