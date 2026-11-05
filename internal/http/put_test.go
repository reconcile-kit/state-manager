package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/auth"
	"github.com/reconcile-kit/state-manager/internal/dto"
	"github.com/reconcile-kit/state-manager/internal/services/states"
)

func TestUpdateResourceErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		readErr    error
		updateErr  error
		wantStatus int
		wantError  string
	}{
		{
			name:       "missing before update",
			readErr:    dto.NotFoundError,
			wantStatus: http.StatusNotFound,
			wantError:  "Resource not found: resource1",
		},
		{
			name:       "wrapped missing before update",
			readErr:    fmt.Errorf("get resource: %w", dto.NotFoundError),
			wantStatus: http.StatusNotFound,
			wantError:  "Resource not found: resource1",
		},
		{
			name:       "missing during update",
			updateErr:  dto.NotFoundError,
			wantStatus: http.StatusNotFound,
			wantError:  "Resource not found: resource1",
		},
		{
			name:       "wrapped missing during update",
			updateErr:  fmt.Errorf("update resource: %w", dto.NotFoundError),
			wantStatus: http.StatusNotFound,
			wantError:  "Resource not found: resource1",
		},
		{
			name:       "database failure",
			updateErr:  errors.New("database unavailable"),
			wantStatus: http.StatusInternalServerError,
			wantError:  "Failed to update resource: database unavailable",
		},
		{
			name:       "untyped not found",
			readErr:    errors.New("resource not found"),
			wantStatus: http.StatusInternalServerError,
			wantError:  "Failed to update resource: resource not found",
		},
		{
			name:       "version conflict",
			updateErr:  dto.ConflictError,
			wantStatus: http.StatusConflict,
			wantError:  "Version conflict: " + dto.ConflictError.Error(),
		},
		{
			name:       "forbidden",
			readErr:    auth.ErrForbidden,
			wantStatus: http.StatusForbidden,
			wantError:  auth.ErrForbidden.Error(),
		},
	}

	for _, suffix := range []string{"", "/status"} {
		for _, tt := range tests {
			t.Run(suffix+"/"+tt.name, func(t *testing.T) {
				t.Parallel()

				repo := &failingUpdateRepository{
					readErr:   tt.readErr,
					updateErr: tt.updateErr,
				}
				service := states.NewStateService(repo, nil, nil)
				router := NewRouter(service, nil)
				path := "/api/v1/groups/group1/namespaces/ns1/kinds/type1/resources/resource1" + suffix
				body := strings.NewReader(`{"shard_id":"shard1","version":1,"current_version":1,"spec":{},"status":{}}`)
				request := httptest.NewRequest(http.MethodPut, path, body)
				response := httptest.NewRecorder()

				router.ServeHTTP(response, request)

				if response.Code != tt.wantStatus {
					t.Fatalf("status = %d, want %d; body: %s", response.Code, tt.wantStatus, response.Body.String())
				}

				var payload ErrorResponse
				if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
					t.Fatalf("invalid error JSON: %v; body: %s", err, response.Body.String())
				}
				if payload.Error != tt.wantError {
					t.Fatalf("error = %q, want %q", payload.Error, tt.wantError)
				}

				wantWrites := 1
				if tt.readErr != nil {
					wantWrites = 0
				}
				if repo.writes != wantWrites {
					t.Fatalf("writes = %d, want %d", repo.writes, wantWrites)
				}
			})
		}
	}
}

type failingUpdateRepository struct {
	states.ResourceRepository
	readErr   error
	updateErr error
	writes    int
}

func (*failingUpdateRepository) TxWrap(_ context.Context, fn func(pgx.Tx) (*dto.Resource, error)) (*dto.Resource, error) {
	return fn(nil)
}

func (*failingUpdateRepository) Lock(context.Context, pgx.Tx, *dto.ResourceID) error {
	return nil
}

func (r *failingUpdateRepository) GetByResourceID(_ context.Context, _ pgx.Tx, id *dto.ResourceID) (*dto.Resource, error) {
	if r.readErr != nil {
		return nil, r.readErr
	}

	return &dto.Resource{
		ResourceFields: dto.ResourceFields{
			ResourceID: *id,
			ShardID:    "shard1",
		},
		Version: 1,
	}, nil
}

func (r *failingUpdateRepository) Update(context.Context, pgx.Tx, *dto.ResourceUpdateOpts) (*dto.Resource, error) {
	r.writes++

	return nil, r.updateErr
}

func (r *failingUpdateRepository) UpdateStatus(context.Context, pgx.Tx, *dto.ResourceUpdateStatusOpts) (*dto.Resource, error) {
	r.writes++

	return nil, r.updateErr
}
