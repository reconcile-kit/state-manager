package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	_ "github.com/reconcile-kit/state-manager/docs"
	"github.com/reconcile-kit/state-manager/internal/auth"
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

// writeForbidden writes 403 if err is an authorization denial.
func writeForbidden(w http.ResponseWriter, err error) bool {
	if !errors.Is(err, auth.ErrForbidden) {
		return false
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	jsonIter.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
	return true
}

// NewRouter creates the HTTP router. A nil authenticator disables authentication,
// the API then works without tokens as before.
func NewRouter(service *states.StateService, authenticator *auth.Authenticator) *chi.Mux {
	handler := &Handler{service: service, validator: validator.New()}
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(mw.AllowAllCORS)
	// API Routes
	r.Route("/api/v1", func(r chi.Router) {
		if authenticator != nil {
			r.Use(auth.Middleware(authenticator))
		}
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

	r.Route("/health", func(r chi.Router) {
		r.Get("/live", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		r.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})

	// Route Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
