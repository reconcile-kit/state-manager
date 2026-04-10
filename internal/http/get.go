package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	apiresource "github.com/reconcile-kit/api/resource"
	"github.com/reconcile-kit/state-manager/internal/dto"
	"github.com/reconcile-kit/state-manager/internal/http/validators"
)

// getResource get a resource by key
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

// listResources list resources by filter
// @Summary List resources
// @Description list resources by filter
// @ID listResources
// @Tags resources
// @Accept json
// @Produce json
// @Param filter query dto.ListResourcesOpts false "Filters"
// @Param label_selector query string false "Format: <key> in (<v1>,<v2>),<key2> in (<v3>). Support only IN (register doesnt matter). Values can be quoted. Max 6 values." example="env in (prod,staging),team in (platform)"
// @Success 200 {array} dto.Resource{spec=map[string]interface{},status=map[string]interface{}} "List of resources"
// @Failure 400 {object} ErrorResponse "Invalid input" example={"error":"Invalid label selector: parse error"}
// @Failure 500 {object} ErrorResponse "Server error" example={"error":"Failed to get resources: database error"}
// @Router /api/v1/resources [get]
func (h *Handler) listResources(w http.ResponseWriter, r *http.Request) {
	labelSelectors, err := apiresource.ParseLabelSelectors(r.URL.Query().Get("label_selector"))
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Invalid label selector: %s"}`, err), http.StatusBadRequest)
		return
	}
	if err = validators.ValidateLabelSelectors(labelSelectors); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"LabelSelector validation failed: %s"}`, err), http.StatusBadRequest)
		return
	}

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
		ShardID:        r.URL.Query().Get("shard_id"),
		Pending:        pending,
		LabelSelectors: labelSelectors,
		Limit:          int(limit),
		Offset:         int(offset),
	}

	resources, err := h.service.ListResources(r.Context(), listOpts)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Failed to get resources: %s"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonIter.NewEncoder(w).Encode(resources)
}
