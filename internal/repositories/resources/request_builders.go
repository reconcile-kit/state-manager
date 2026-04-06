package resources

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	apiresource "github.com/reconcile-kit/api/resource"
	"github.com/reconcile-kit/state-manager/internal/dto"
)

func buildListResourcesByIDsQuery(resourceIDs []int) (string, []interface{}) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)

	sb.Select(
		"r.id",
		"r.shard_id",
		"r.resource_group",
		"r.kind",
		"r.namespace",
		"r.name",
		"r.created_at",
		"r.updated_at",
		"r.deletion_timestamp",
		"r.finalizers",
		"r.annotations",
		"r.spec",
		"r.status",
		"r.version",
		"r.current_version",
	).From("resources r")

	sb.Where(sb.In("r.id", toInterface(resourceIDs)...))

	return sb.Build()
}

const (
	defaultLabelAlias   = "l0"
	fieldNameValue      = "value"
	fieldNameName       = "name"
	fieldNameResourceID = "resource_id"
)

func buildListResourceIDsQuery(listOpts *dto.ListResourcesOpts) (string, []interface{}, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SetFlavor(sqlbuilder.PostgreSQL)

	sb = sb.Select(
		"r.id",
	).Distinct()

	if len(listOpts.LabelSelectors) > 0 {
		sb.From("labels " + defaultLabelAlias)
		sb.Join("resources r", fmt.Sprintf("r.id = %s", buildAliasField(defaultLabelAlias, fieldNameResourceID)))

		sel := listOpts.LabelSelectors[0]
		valuesStatement, err := buildSelectorValuesStatement(sb, defaultLabelAlias, sel)
		if err != nil {
			return "", nil, err
		}

		sb.Where(
			sb.E(buildAliasField(defaultLabelAlias, fieldNameName), sel.Key),
			valuesStatement,
		)

		for i := 1; i < len(listOpts.LabelSelectors); i++ {
			alias := fmt.Sprintf("l%d", i)
			sel := listOpts.LabelSelectors[i]
			valuesStatement, err := buildSelectorValuesStatement(sb, alias, sel)
			if err != nil {
				return "", nil, err
			}

			sb.Join(
				"labels "+alias,
				fmt.Sprintf("%s = %s", buildAliasField(alias, fieldNameResourceID), buildAliasField(defaultLabelAlias, fieldNameResourceID)),
				sb.E(buildAliasField(alias, fieldNameName), sel.Key),
				valuesStatement,
			)
		}
	} else {
		sb.From("resources r")
	}

	if listOpts.ResourceGroup != "" {
		sb.Where(sb.E("r.resource_group", listOpts.ResourceGroup))
	}
	if listOpts.Kind != "" {
		sb.Where(sb.E("r.kind", listOpts.Kind))
	}
	if listOpts.Namespace != "" {
		sb.Where(sb.E("r.namespace", listOpts.Namespace))
	}
	if listOpts.Name != "" {
		sb.Where(sb.E("r.name", listOpts.Name))
	}
	if listOpts.ShardID != "" {
		sb.Where(sb.E("r.shard_id", listOpts.ShardID))
	}
	if listOpts.Pending {
		sb.Where("r.version > r.current_version")
	}

	sb.OrderBy("r.id")

	if listOpts.Limit > 0 {
		sb.Limit(listOpts.Limit)
	}
	if listOpts.Offset > 0 {
		sb.Offset(listOpts.Offset)
	}

	sql, args := sb.Build()
	return sql, args, nil
}

func buildSelectorValuesStatement(
	sb *sqlbuilder.SelectBuilder,
	alias string,
	selector apiresource.LabelSelector,
) (string, error) {
	values := selector.Values
	if len(values) == 0 {
		values = []string{""}
	}

	switch selector.Operator {
	case apiresource.LabelSelectorOperatorIn:

		return sb.In(buildAliasField(alias, fieldNameValue), toInterface(values)...), nil
	default:
		return "", fmt.Errorf("unsupported selector operator: %s", selector.Operator)
	}
}

func buildAliasField(alias, fieldName string) string {
	return fmt.Sprintf("%s.%s", alias, fieldName)
}
