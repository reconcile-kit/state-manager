package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	_ "github.com/reconcile-kit/state-manager/docs"
	"github.com/reconcile-kit/state-manager/internal/services/states"
	mw "github.com/reconcile-kit/state-manager/pkg/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

var jsonIter = jsoniter.Config{
	EscapeHTML:             false, // Ускоряет маршалинг, если HTML-экранирование не нужно
	SortMapKeys:            false, // Ускоряет маршалинг для map
	ValidateJsonRawMessage: true,  // Поддержка json.RawMessage
}.Froze()

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	service   *states.StateService
	validator *validator.Validate
}

func NewRouter(service *states.StateService) *chi.Mux {
	handler := &Handler{service: service, validator: validator.New()}
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(mw.AllowAllCORS)
	// Группа маршрутов API
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/resources", func(r chi.Router) {
			r.Get("/", handler.listResources)
		})
		r.Route("/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources", func(r chi.Router) {
			r.Post("/", handler.createResource)
			r.Get("/{name}", handler.getResource)
			r.Put("/{name}", handler.updateResource)
			r.Delete("/{name}", handler.deleteResource)
			r.Put("/{name}/status", handler.updateResourceStatus)
		})
	})

	// Маршрут для Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
