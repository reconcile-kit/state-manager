package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

// Middleware authenticates the bearer token and stores the Principal in the request context.
func Middleware(a *Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			raw, ok := bearerToken(r.Header.Get("Authorization"))
			if !ok {
				writeUnauthorized(w, "missing bearer token")
				return
			}
			p, err := a.Authenticate(r.Context(), raw)
			if err != nil {
				if errors.Is(err, ErrUnauthenticated) {
					writeUnauthorized(w, err.Error())
					return
				}
				log.Printf("authentication failed: %v", err)
				writeError(w, http.StatusInternalServerError, "Authentication failed")
				return
			}
			next.ServeHTTP(w, r.WithContext(WithPrincipal(r.Context(), p)))
		})
	}
}

func bearerToken(header string) (string, bool) {
	const prefix = "bearer "
	if len(header) <= len(prefix) || !strings.EqualFold(header[:len(prefix)], prefix) {
		return "", false
	}
	token := strings.TrimSpace(header[len(prefix):])
	return token, token != ""
}

func writeUnauthorized(w http.ResponseWriter, msg string) {
	w.Header().Set("WWW-Authenticate", `Bearer`)
	writeError(w, http.StatusUnauthorized, msg)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{Error: msg})
}
