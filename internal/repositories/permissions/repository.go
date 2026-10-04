package permissions

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/reconcile-kit/state-manager/internal/auth"
)

type PostgresPermissionsRepo struct {
	pool *pgxpool.Pool
}

func NewPermissionsRepository(pool *pgxpool.Pool) *PostgresPermissionsRepo {
	return &PostgresPermissionsRepo{pool: pool}
}

// Both branches hit idx_auth_role_bindings_active, rules are fetched by idx_auth_role_rules_role_id.
// UNION ALL instead of OR keeps two plain index scans; duplicate rules are harmless for matching.
const rulesQuery = `
	SELECT rr.verbs, rr.resource_group, rr.namespace, rr.kind, rr.name, rr.shard_id
	FROM (
		SELECT role_id FROM auth_role_bindings
		WHERE subject_kind = 'subject' AND subject_value = $1 AND NOT disabled
		UNION ALL
		SELECT role_id FROM auth_role_bindings
		WHERE subject_kind = 'group' AND subject_value = ANY($2::varchar[]) AND NOT disabled
	) b
	JOIN auth_role_rules rr ON rr.role_id = b.role_id`

func (r *PostgresPermissionsRepo) Rules(ctx context.Context, subject string, groups []string) ([]auth.Rule, error) {
	rows, err := r.pool.Query(ctx, rulesQuery, subject, groups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []auth.Rule
	for rows.Next() {
		var spec auth.RuleSpec
		if err := rows.Scan(&spec.Verbs, &spec.ResourceGroup, &spec.Namespace, &spec.Kind, &spec.Name, &spec.ShardID); err != nil {
			return nil, err
		}
		rule, err := spec.Compile()
		if err != nil {
			return nil, fmt.Errorf("invalid rule in auth_role_rules: %w", err)
		}
		rules = append(rules, rule)
	}
	return rules, rows.Err()
}
