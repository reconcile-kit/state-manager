package http

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/reconcile-kit/state-manager/internal/dto"
	"net/http"
	"strconv"
)

// getResource получает ресурс по ключу
// @Summary Get a resource by key
// @Description Retrieves a resource by its resource_group, kind, namespace, and name.
// @ID getResource
// @Tags resources
// @Accept json
// @Produce json
// @Param resource_group path string true "Resource Group" example="group1"
// @Param kind path string true "Kind" example="type1"
// @Param namespace path string true "Namespace" example="ns1"
// @Param name path string true "Name" example="resource1"
// @Success 200 {object} dto.Resource{spec=map[string]interface{},status=map[string]interface{}} "Resource found"
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Validation failed: resource_group is required"}
// @Failure 404 {object} ErrorResponse "Not found" example={"error":"Resource not found: no rows"}
// @Router /api/v1/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources/{name} [get]
func (h *Handler) getResource(w http.ResponseWriter, r *http.Request) {
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

	resource, err := h.service.GetByResourceID(r.Context(), opts)
	if err != nil {
		if errors.Is(err, dto.NotFoundError) {
			http.Error(w, fmt.Sprintf(`{"error":"Resource not found: %s"}`, opts.Name), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf(`{"error":"Failed to get resource: %s"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonIter.NewEncoder(w).Encode(resource)
}

// listResources возвращает ресурсы по фильтру
// @Summary List pending resources
// @Description возвращает ресурсы по фильтру
// @ID listResources
// @Tags resources
// @Accept json
// @Produce json
// @Param filter query dto.ListResourcesOpts true "Kind" example="type1"
// @Success 200 {object} dto.Resource{spec=map[string]interface{},status=map[string]interface{}} "List of pending resources"
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"shard_ids is required"}
// @Failure 500 {object} ErrorResponse "Server error" example={"error":"Failed to list pending resources: database error"}
// @Router /api/v1/resources [get]
func (h *Handler) listResources(w http.ResponseWriter, r *http.Request) {

	pending, _ := strconv.ParseBool(r.URL.Query().Get("pending"))
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if limit == 0 {
		limit = 500
	}
	offset, _ := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	listOpts := &dto.ListResourcesOpts{
		ResourceID: dto.ResourceID{
			ResourceGroup: r.URL.Query().Get("resource_group"),
			Kind:          r.URL.Query().Get("kind"),
			Namespace:     r.URL.Query().Get("namespace"),
			Name:          r.URL.Query().Get("name"),
		},
		ShardID: r.URL.Query().Get("shard_id"),
		Pending: pending,
		Limit:   int(limit),
		Offset:  int(offset),
	}

	resources, err := h.service.ListPending(r.Context(), listOpts)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Failed to get resources: %s"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonIter.NewEncoder(w).Encode(resources)
}
