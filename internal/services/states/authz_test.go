package states

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/reconcile-kit/state-manager/internal/auth"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

// fakeRepo holds a single stored resource and records writes.
type fakeRepo struct {
	stored *dto.Resource
	writes int
}

func (f *fakeRepo) TxWrap(_ context.Context, fn func(tx pgx.Tx) (*dto.Resource, error)) (*dto.Resource, error) {
	return fn(nil)
}
func (f *fakeRepo) Lock(context.Context, pgx.Tx, *dto.ResourceID) error { return nil }
func (f *fakeRepo) Create(_ context.Context, _ pgx.Tx, o *dto.ResourceCreateOpts) (*dto.Resource, error) {
	f.writes++
	return &dto.Resource{ResourceFields: o.ResourceFields}, nil
}
func (f *fakeRepo) Update(_ context.Context, _ pgx.Tx, o *dto.ResourceUpdateOpts) (*dto.Resource, error) {
	f.writes++
	return &dto.Resource{ResourceFields: o.ResourceFields}, nil
}
func (f *fakeRepo) UpdateStatus(_ context.Context, _ pgx.Tx, o *dto.ResourceUpdateStatusOpts) (*dto.Resource, error) {
	f.writes++
	return &dto.Resource{ResourceFields: o.ResourceFields}, nil
}
func (f *fakeRepo) UpdateLabels(context.Context, pgx.Tx, int, map[string]string, map[string]string, []string) error {
	return nil
}
func (f *fakeRepo) SetDeletionTimestamp(context.Context, pgx.Tx, *dto.ResourceID, time.Time) error {
	f.writes++
	return nil
}
func (f *fakeRepo) Delete(context.Context, pgx.Tx, int64) error {
	f.writes++
	return nil
}
func (f *fakeRepo) GetByResourceID(context.Context, pgx.Tx, *dto.ResourceID) (*dto.Resource, error) {
	return f.stored, nil
}
func (f *fakeRepo) ListResources(context.Context, *dto.ListResourcesOpts) ([]*dto.Resource, error) {
	return nil, nil
}

type fakeEvents struct{}

func (fakeEvents) Add(context.Context, string, string, dto.ResourceID) error { return nil }

var resourceID = dto.ResourceID{ResourceGroup: "g", Namespace: "team-a", Kind: "user", Name: "u1"}

func newTestService(rules ...auth.Rule) (*StateService, *fakeRepo, context.Context) {
	repo := &fakeRepo{stored: &dto.Resource{ResourceFields: dto.ResourceFields{ResourceID: resourceID, ShardID: "5"}}}
	ctx := auth.WithPrincipal(context.Background(), &auth.Principal{Subject: "svc", Rules: rules})
	return NewStateService(repo, fakeEvents{}, auth.RulesAuthorizer{}), repo, ctx
}

func TestServiceChecksStoredShard(t *testing.T) {
	svc, _, ctx := newTestService(auth.Rule{Verbs: auth.VerbGet | auth.VerbDelete, ShardID: "5"})
	if _, err := svc.GetByResourceID(ctx, &resourceID); err != nil {
		t.Fatalf("get in granted shard: %v", err)
	}

	svc, repo, ctx := newTestService(auth.Rule{Verbs: auth.VerbAll, ShardID: "7"})
	if _, err := svc.GetByResourceID(ctx, &resourceID); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("get in foreign shard: expected forbidden, got %v", err)
	}
	if err := svc.Delete(ctx, &resourceID); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("delete in foreign shard: expected forbidden, got %v", err)
	}
	if repo.writes != 0 {
		t.Fatalf("no writes expected, got %d", repo.writes)
	}
}

func TestServiceUpdateCannotMoveResourceToForeignShard(t *testing.T) {
	svc, repo, ctx := newTestService(auth.Rule{Verbs: auth.VerbUpdate | auth.VerbUpdateStatus, ShardID: "5"})

	update := &dto.ResourceUpdateOpts{ResourceFields: dto.ResourceFields{ResourceID: resourceID, ShardID: "5"}}
	if _, err := svc.Update(ctx, update); err != nil {
		t.Fatalf("update in same shard: %v", err)
	}

	update.ShardID = "7"
	if _, err := svc.Update(ctx, update); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("move to foreign shard: expected forbidden, got %v", err)
	}
	status := &dto.ResourceUpdateStatusOpts{ResourceFields: dto.ResourceFields{ResourceID: resourceID, ShardID: "7"}}
	if _, err := svc.UpdateStatus(ctx, status); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("status move to foreign shard: expected forbidden, got %v", err)
	}
	if repo.writes != 1 {
		t.Fatalf("expected only the first update to write, got %d", repo.writes)
	}
}

func TestServiceUpdateCannotTakeResourceFromForeignShard(t *testing.T) {
	// Granted shard 7 only, the stored resource lives in shard 5.
	svc, _, ctx := newTestService(auth.Rule{Verbs: auth.VerbUpdate, ShardID: "7"})
	update := &dto.ResourceUpdateOpts{ResourceFields: dto.ResourceFields{ResourceID: resourceID, ShardID: "7"}}
	if _, err := svc.Update(ctx, update); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("expected forbidden, got %v", err)
	}
}

func TestServiceUpdateSpecRequiresUpdateVerb(t *testing.T) {
	svc, _, ctx := newTestService(auth.Rule{Verbs: auth.VerbUpdateStatus})
	update := &dto.ResourceUpdateOpts{ResourceFields: dto.ResourceFields{ResourceID: resourceID, ShardID: "5"}}
	if _, err := svc.Update(ctx, update); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("expected forbidden, got %v", err)
	}
}

func TestServiceCreateAndList(t *testing.T) {
	svc, _, ctx := newTestService(auth.Rule{Verbs: auth.VerbCreate | auth.VerbList, Namespace: "team-a", Kind: "user"})

	create := &dto.ResourceCreateOpts{ResourceFields: dto.ResourceFields{ResourceID: resourceID, ShardID: "any"}}
	if _, err := svc.Create(ctx, create); err != nil {
		t.Fatalf("create: %v", err)
	}
	create.Namespace = "team-b"
	if _, err := svc.Create(ctx, create); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("create in foreign namespace: expected forbidden, got %v", err)
	}

	list := &dto.ListResourcesOpts{ResourceID: dto.ResourceID{Namespace: "team-a", Kind: "user"}}
	if _, err := svc.ListResources(ctx, list); err != nil {
		t.Fatalf("covered list: %v", err)
	}
	list.Kind = ""
	if _, err := svc.ListResources(ctx, list); !errors.Is(err, auth.ErrForbidden) {
		t.Fatalf("uncovered list: expected forbidden, got %v", err)
	}
}

func TestServiceWithoutAuthorizer(t *testing.T) {
	repo := &fakeRepo{stored: &dto.Resource{ResourceFields: dto.ResourceFields{ResourceID: resourceID, ShardID: "5"}}}
	svc := NewStateService(repo, fakeEvents{}, nil)
	if _, err := svc.ListResources(context.Background(), &dto.ListResourcesOpts{}); err != nil {
		t.Fatalf("disabled authorization must allow everything: %v", err)
	}
	if _, err := svc.GetByResourceID(context.Background(), &resourceID); err != nil {
		t.Fatalf("disabled authorization must allow everything: %v", err)
	}
}
