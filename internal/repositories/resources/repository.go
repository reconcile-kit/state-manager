package resources

import (
	"context"
	"fmt"
	"github.com/dhnikolas/state-manager/internal/dto"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresResourceRepo struct {
	pool *pgxpool.Pool
}

func NewResourceRepository(pool *pgxpool.Pool) *PostgresResourceRepo {
	return &PostgresResourceRepo{pool: pool}
}

func (r *PostgresResourceRepo) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return r.pool.Begin(ctx)
}

func (r *PostgresResourceRepo) TxWrap(ctx context.Context, fn func(tx pgx.Tx) (*dto.Resource, error)) (*dto.Resource, error) {
	tx, err := r.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	res, err := fn(tx)
	if err != nil {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			return nil, fmt.Errorf("rollback transaction error: %s", err)
		}
		return nil, err
	}
	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *PostgresResourceRepo) Create(ctx context.Context, tx pgx.Tx, opts *dto.ResourceCreateOpts) (*dto.Resource, error) {
	err := validateResourceID(&opts.ResourceID)
	if err != nil {
		return nil, err
	}
	res := &dto.Resource{}
	res.ResourceFields = opts.ResourceFields
	res.Annotations = opts.Annotations
	res.Spec = opts.Spec
	const q = `INSERT INTO resources (shard_id, resource_group, kind, namespace, name, annotations, spec) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id, created_at, updated_at, shard_id, version`

	row := tx.QueryRow(ctx, q,
		opts.ShardID,
		opts.ResourceGroup,
		opts.Kind,
		opts.Namespace,
		opts.Name,
		opts.Annotations,
		opts.Spec,
	)
	if err := row.Scan(&res.ID, &res.CreatedAt, &res.UpdatedAt, &res.ShardID, &res.Version); err != nil {
		return nil, err
	}
	for name, value := range opts.Labels {
		const li = `INSERT INTO labels (resource_id, name, value) VALUES ($1,$2,$3)`
		if _, err := tx.Exec(ctx, li, res.ID, name, value); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (r *PostgresResourceRepo) Update(ctx context.Context, tx pgx.Tx, opts *dto.ResourceUpdateOpts) (*dto.Resource, error) {
	err := validateResourceID(&opts.ResourceID)
	if err != nil {
		return nil, err
	}
	const q = `UPDATE resources SET shard_id=$1, annotations=$2, spec=$3, version=version+1, updated_at=NOW() WHERE resource_group=$4 AND kind=$5 AND namespace=$6 AND name=$7 RETURNING created_at, updated_at, id, shard_id, version`
	res := &dto.Resource{}
	res.ResourceFields = opts.ResourceFields
	res.Spec = opts.Spec
	res.Annotations = opts.Annotations
	row := tx.QueryRow(ctx, q,
		opts.ShardID, opts.Annotations, opts.Spec, opts.ResourceGroup, opts.Kind, opts.Namespace, opts.Name)
	if err := row.Scan(&res.CreatedAt, &res.UpdatedAt, &res.ID, &res.ShardID, &res.Version); err != nil {
		return nil, err
	}

	// delete old labels
	if _, err := tx.Exec(ctx, `DELETE FROM labels WHERE resource_id=$1`, res.ID); err != nil {
		return nil, err
	}
	// insert new labels
	for name, value := range opts.Labels {
		const li = `INSERT INTO labels (resource_id, name, value) VALUES ($1,$2,$3)`
		if _, err := tx.Exec(ctx, li, res.ID, name, value); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (r *PostgresResourceRepo) UpdateStatus(ctx context.Context, tx pgx.Tx, opts *dto.ResourceUpdateStatusOpts) (*dto.Resource, error) {
	err := validateResourceID(&opts.ResourceID)
	if err != nil {
		return nil, err
	}
	const q = `UPDATE resources SET shard_id=$1, annotations=$2, status=$3, current_version=$4, updated_at=NOW() WHERE resource_group=$5 AND kind=$6 AND namespace=$7 AND name=$8 RETURNING created_at, updated_at, id, shard_id, version, current_version, spec`
	res := &dto.Resource{}
	res.ResourceFields = opts.ResourceFields
	res.Annotations = opts.Annotations
	res.Status = opts.Status
	row := tx.QueryRow(ctx, q,
		opts.ShardID, opts.Annotations, opts.Status, opts.CurrentVersion, opts.ResourceGroup, opts.Kind, opts.Namespace, opts.Name)
	if err := row.Scan(&res.CreatedAt, &res.UpdatedAt, &res.ID, &res.ShardID, &res.Version, &res.CurrentVersion, &res.Spec); err != nil {
		return nil, err
	}

	// delete old labels
	if _, err := tx.Exec(ctx, `DELETE FROM labels WHERE resource_id=$1`, res.ID); err != nil {
		return nil, err
	}
	// insert new labels
	for name, value := range opts.Labels {
		const li = `INSERT INTO labels (resource_id, name, value) VALUES ($1,$2,$3)`
		if _, err := tx.Exec(ctx, li, res.ID, name, value); err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (r *PostgresResourceRepo) Delete(ctx context.Context, tx pgx.Tx, id int64) error {
	_, err := tx.Exec(ctx, `DELETE FROM resources WHERE id=$1`, id)
	return err
}

func (r *PostgresResourceRepo) GetByResourceID(ctx context.Context, opts *dto.ResourceID) (*dto.Resource, error) {
	err := validateResourceID(opts)
	if err != nil {
		return nil, err
	}
	const q = `SELECT 
       id,
       resource_group,
       kind,
       namespace,
       name,
       shard_id,
       created_at,
       updated_at,
       annotations,
       spec,
       status,
       version,
       current_version 
    FROM resources WHERE resource_group=$1 AND kind=$2 AND namespace=$3 AND name=$4`
	res := &dto.Resource{}
	row := r.pool.QueryRow(ctx, q, opts.ResourceGroup, opts.Kind, opts.Namespace, opts.Name)
	if err := row.Scan(
		&res.ID,
		&res.ResourceGroup,
		&res.Kind,
		&res.Namespace,
		&res.Name,
		&res.ShardID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Annotations,
		&res.Spec,
		&res.Status,
		&res.Version,
		&res.CurrentVersion,
	); err != nil {
		return nil, err
	}
	res.Labels = make(map[string]string)
	rows, err := r.pool.Query(ctx, `SELECT name, value FROM labels WHERE resource_id=$1`, res.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name, value string
		err = rows.Scan(&name, &value)
		if err != nil {
			return nil, err
		}
		res.Labels[name] = value
	}
	return res, nil
}

func (r *PostgresResourceRepo) ListResources(ctx context.Context, listOpts *dto.ListResourcesOpts) ([]*dto.Resource, error) {
	sql := sqlbuilder.Select(
		"id",
		"shard_id",
		"resource_group",
		"kind",
		"namespace",
		"name",
		"created_at",
		"updated_at",
		"annotations",
		"spec",
		"status",
		"version",
		"current_version",
	).From("resources")
	sql.SetFlavor(sqlbuilder.PostgreSQL)
	if listOpts.ResourceGroup != "" {
		sql.Where(sql.E("resource_group", listOpts.ResourceGroup))
	}
	if listOpts.Kind != "" {
		sql.Where(sql.E("kind", listOpts.Kind))
	}
	if listOpts.Namespace != "" {
		sql.Where(sql.E("namespace", listOpts.Namespace))
	}
	if listOpts.Name != "" {
		sql.Where(sql.E("name", listOpts.Name))
	}
	if len(listOpts.ShardID) > 0 {
		sql.Where(sql.In("shard_id", listOpts.ShardID))
	}
	if listOpts.Pending {
		sql.Where(sql.And("version > current_version"))
	}

	sql.Limit(listOpts.Limit)
	sql.Offset(listOpts.Offset)

	sqlQuery, args := sql.Build()

	rows, err := r.pool.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resources := []*dto.Resource{}
	resourceIDMap := make(map[int]*dto.Resource)
	var resourceIDs []int

	for rows.Next() {
		res := &dto.Resource{}
		if err := rows.Scan(
			&res.ID, &res.ShardID, &res.ResourceGroup,
			&res.Kind, &res.Namespace, &res.Name,
			&res.CreatedAt, &res.UpdatedAt,
			&res.Annotations, &res.Spec, &res.Status, &res.Version, &res.CurrentVersion,
		); err != nil {
			return nil, err
		}
		res.Labels = make(map[string]string)
		resources = append(resources, res)
		resourceIDMap[res.ID] = res
		resourceIDs = append(resourceIDs, res.ID)
	}

	if len(resourceIDs) == 0 {
		return resources, nil
	}

	// Второй запрос: получаем все метки для найденных ресурсов
	const labelQuery = `
        SELECT resource_id, name, value 
        FROM labels 
        WHERE resource_id = ANY($1)
    `
	labelRows, err := r.pool.Query(ctx, labelQuery, resourceIDs)
	if err != nil {
		return nil, err
	}
	defer labelRows.Close()

	// Маппинг меток к ресурсам
	for labelRows.Next() {
		var resourceID int
		var name, value string
		if err := labelRows.Scan(&resourceID, &name, &value); err != nil {
			return nil, err
		}
		if resource, exists := resourceIDMap[resourceID]; exists {
			resource.Labels[name] = value
		}
	}

	return resources, nil
}

func validateResourceID(ID *dto.ResourceID) error {
	if ID.ResourceGroup == "" || ID.Kind == "" || ID.Namespace == "" || ID.Name == "" {
		return fmt.Errorf("resource group, kind and namespace must be set")
	}
	return nil
}
