package http

import (
	"fmt"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// deleteResource получает ресурс по ключу
// @Summary Delete a resource
// @Description Delete a resource by its resource_group, kind, namespace, and name.
// @ID deleteResource
// @Tags resources
// @Accept json
// @Produce json
// @Param resource_group path string true "Resource Group" example="group1"
// @Param kind path string true "Kind" example="type1"
// @Param namespace path string true "Namespace" example="ns1"
// @Param name path string true "Name" example="resource1"
// @Success 204
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Deletion failed: resource_group is required"}
// @Failure 404 {object} ErrorResponse "Not found" example={"error":"Resource not found: no rows"}
// @Router /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name} [delete]
func (h *Handler) deleteResource(w http.ResponseWriter, r *http.Request) {
	opts := &dto.ResourceID{
		ResourceGroup: chi.URLParam(r, "resource_group"),
		Kind:          chi.URLParam(r, "kind"),
		Namespace:     chi.URLParam(r, "namespace"),
		Name:          chi.URLParam(r, "name"),
	}

	if err := h.validator.Struct(opts); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Validation failed: %s"}`, err), http.StatusBadRequest)
		return
	}

	err := h.service.Delete(r.Context(), opts)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Failed to delete resource: %s"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
