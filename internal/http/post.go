package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
	"net/http"
)

type CreateResourceRequest struct {
	ShardID     string            `json:"shard_id" validate:"required"`
	Labels      map[string]string `json:"labels"`
	Finalizers  []string          `json:"finalizers"`
	Annotations map[string]string `json:"annotations"`
	Name        string            `json:"name" validate:"required"`
	Spec        json.RawMessage   `json:"spec"`
}

// createResource создаёт новый ресурс
// @Summary Create a new resource
// @Description Creates a new resource with the provided details.
// @ID createResource
// @Tags resources
// @Accept json
// @Produce json
// @Param resource_group path string true "Resource Group" example="group1"
// @Param kind path string true "Kind" example="type1"
// @Param namespace path string true "Namespace" example="ns1"
// @Param resource body CreateResourceRequest{spec=object} true "Resource details"
// @Success 200 {object} dto.Resource{spec=map[string]interface{},status=map[string]interface{}}  "Resource created"
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Validation failed: shard_id is required"}
// @Failure 500 {object} ErrorResponse "Server error" example={"error":"Failed to create resource: database error"}
// @Router /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources [post]
func (h *Handler) createResource(w http.ResponseWriter, r *http.Request) {
	var req CreateResourceRequest
	if err := jsonIter.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid JSON: %s"}`, err), http.StatusBadRequest)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Validation failed: %s"}`, err), http.StatusBadRequest)
		return
	}

	resourceCreateOpts := &dto.ResourceCreateOpts{
		ResourceFields: dto.ResourceFields{
			ResourceID: dto.ResourceID{
				ResourceGroup: chi.URLParam(r, "resource_group"),
				Kind:          chi.URLParam(r, "kind"),
				Namespace:     chi.URLParam(r, "namespace"),
				Name:          req.Name,
			},
			ShardID:     req.ShardID,
			Labels:      req.Labels,
			Finalizers:  req.Finalizers,
			Annotations: req.Annotations,
		},
		Spec: req.Spec,
	}

	resource, err := h.service.Create(r.Context(), resourceCreateOpts)
	if err != nil {
		if errors.Is(err, dto.AlreadyExistsError) {
			http.Error(w, fmt.Sprintf(`{"error":"Already exists: %s"}`, err), http.StatusBadRequest)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"Failed to create resource: %s"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	jsonIter.NewEncoder(w).Encode(resource)
}
