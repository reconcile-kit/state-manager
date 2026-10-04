package auth

import (
	"fmt"
	"strings"
)

// claimPath splits a dotted claim name, e.g. "resource_access.state-manager.permissions".
func claimPath(name string) []string {
	if name == "" {
		return nil
	}
	return strings.Split(name, ".")
}

func lookupClaim(claims map[string]any, path []string) (any, bool) {
	if len(path) == 0 {
		return nil, false
	}
	var cur any = claims
	for _, key := range path {
		m, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = m[key]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}

func parseStringClaim(v any) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("expected string, got %T", v)
	}
	return s, nil
}

// parseStringsClaim accepts an array of strings or a single string.
func parseStringsClaim(v any) ([]string, error) {
	switch val := v.(type) {
	case string:
		return []string{val}, nil
	case []any:
		out := make([]string, 0, len(val))
		for _, item := range val {
			s, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("expected array of strings, got element %T", item)
			}
			out = append(out, s)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("expected array of strings, got %T", v)
	}
}

// parseRulesClaim converts a JSON array of rule objects into compiled rules.
// It walks the already decoded claims instead of re-marshaling them.
func parseRulesClaim(v any) ([]Rule, error) {
	items, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("expected array of rules, got %T", v)
	}
	rules := make([]Rule, 0, len(items))
	for i, item := range items {
		obj, ok := item.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("rule %d: expected object, got %T", i, item)
		}
		spec, err := ruleSpecFromMap(obj)
		if err != nil {
			return nil, fmt.Errorf("rule %d: %w", i, err)
		}
		rule, err := spec.Compile()
		if err != nil {
			return nil, fmt.Errorf("rule %d: %w", i, err)
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func ruleSpecFromMap(obj map[string]any) (RuleSpec, error) {
	var spec RuleSpec
	verbs, ok := obj["verbs"]
	if !ok {
		return spec, fmt.Errorf("verbs is required")
	}
	var err error
	if spec.Verbs, err = parseStringsClaim(verbs); err != nil {
		return spec, fmt.Errorf("verbs: %w", err)
	}

	fields := []struct {
		name string
		dst  *string
	}{
		{"resource_group", &spec.ResourceGroup},
		{"namespace", &spec.Namespace},
		{"kind", &spec.Kind},
		{"name", &spec.Name},
		{"shard_id", &spec.ShardID},
	}
	for _, f := range fields {
		v, ok := obj[f.name]
		if !ok {
			continue
		}
		if *f.dst, err = parseStringClaim(v); err != nil {
			return spec, fmt.Errorf("%s: %w", f.name, err)
		}
	}
	return spec, nil
}
