package http

import (
	_ "github.com/dhnikolas/state-manager/docs"
	"github.com/dhnikolas/state-manager/internal/services/states"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	httpSwagger "github.com/swaggo/http-swagger"
)

var json = jsoniter.Config{
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
	// Группа маршрутов API
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/resources", func(r chi.Router) {
			r.Get("/", handler.listResources)
		})
		r.Route("/groups/{resource_group}/namespaces/{namespace}/kinds/{kind}/resources", func(r chi.Router) {
			r.Post("/", handler.createResource)
			r.Get("/{name}", handler.getResource)
			r.Put("/{name}", handler.updateResource)
		})
	})

	// Маршрут для Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
