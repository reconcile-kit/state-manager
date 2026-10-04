package auth

import (
	"context"
	"errors"
	"testing"
)

func mustRules(t *testing.T, specs ...RuleSpec) []Rule {
	t.Helper()
	rules := make([]Rule, 0, len(specs))
	for _, s := range specs {
		r, err := s.Compile()
		if err != nil {
			t.Fatalf("compile %+v: %v", s, err)
		}
		rules = append(rules, r)
	}
	return rules
}

func TestParseVerbs(t *testing.T) {
	v, err := ParseVerbs([]string{"get", "list"})
	if err != nil || v != VerbGet|VerbList {
		t.Fatalf("got %v, %v", v, err)
	}
	if v, _ = ParseVerbs([]string{"*"}); v != VerbAll {
		t.Fatalf("wildcard: got %v", v)
	}
	if _, err = ParseVerbs([]string{"patch"}); err == nil {
		t.Fatal("expected error for unknown verb")
	}
	if _, err = ParseVerbs(nil); err == nil {
		t.Fatal("expected error for empty verbs")
	}
}

func TestRulesAuthorizer(t *testing.T) {
	// Read everything in shard 5 and everything of kind user.
	rules := mustRules(t,
		RuleSpec{Verbs: []string{"get", "list"}, ShardID: "5"},
		RuleSpec{Verbs: []string{"get", "list"}, Kind: "user", ResourceGroup: "*"},
		RuleSpec{Verbs: []string{"update_status"}, Namespace: "team-a", ShardID: "1"},
	)
	ctx := WithPrincipal(context.Background(), &Principal{Subject: "svc", Rules: rules})

	tests := []struct {
		name  string
		verb  Verb
		attrs Attributes
		allow bool
	}{
		{"get resource in shard 5", VerbGet, Attributes{"g", "ns", "order", "a", "5"}, true},
		{"get user in other shard", VerbGet, Attributes{"g", "ns", "user", "a", "7"}, true},
		{"get order in other shard", VerbGet, Attributes{"g", "ns", "order", "a", "7"}, false},
		{"delete in shard 5 not granted", VerbDelete, Attributes{"g", "ns", "order", "a", "5"}, false},
		{"update_status in granted ns+shard", VerbUpdateStatus, Attributes{"g", "team-a", "pod", "a", "1"}, true},
		{"update spec not granted", VerbUpdate, Attributes{"g", "team-a", "pod", "a", "1"}, false},
		{"update_status in other shard", VerbUpdateStatus, Attributes{"g", "team-a", "pod", "a", "2"}, false},

		{"list ?shard_id=5", VerbList, Attributes{ShardID: "5"}, true},
		{"list ?kind=user", VerbList, Attributes{Kind: "user"}, true},
		{"list ?kind=user&namespace=x", VerbList, Attributes{Kind: "user", Namespace: "x"}, true},
		{"list ?shard_id=5&kind=order", VerbList, Attributes{Kind: "order", ShardID: "5"}, true},
		{"list ?kind=order not covered", VerbList, Attributes{Kind: "order"}, false},
		{"list without filter not covered", VerbList, Attributes{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RulesAuthorizer{}.Authorize(ctx, tt.verb, &tt.attrs)
			if tt.allow && err != nil {
				t.Fatalf("expected allow, got %v", err)
			}
			if !tt.allow && !errors.Is(err, ErrForbidden) {
				t.Fatalf("expected ErrForbidden, got %v", err)
			}
		})
	}
}

func TestRulesAuthorizerWithoutPrincipal(t *testing.T) {
	err := RulesAuthorizer{}.Authorize(context.Background(), VerbGet, &Attributes{})
	if !errors.Is(err, ErrUnauthenticated) {
		t.Fatalf("expected ErrUnauthenticated, got %v", err)
	}
}

func TestUnrestrictedRuleAllowsListWithoutFilter(t *testing.T) {
	ctx := WithPrincipal(context.Background(), &Principal{Rules: mustRules(t, RuleSpec{Verbs: []string{"*"}})})
	if err := (RulesAuthorizer{}).Authorize(ctx, VerbList, &Attributes{}); err != nil {
		t.Fatalf("expected allow, got %v", err)
	}
}

func BenchmarkRulesAuthorizer(b *testing.B) {
	rules := make([]Rule, 0, 20)
	for i := 0; i < 19; i++ {
		rules = append(rules, Rule{Verbs: VerbGet | VerbList, Namespace: "ns", ShardID: "other"})
	}
	rules = append(rules, Rule{Verbs: VerbGet, Kind: "user"})
	ctx := WithPrincipal(context.Background(), &Principal{Rules: rules})
	attrs := &Attributes{ResourceGroup: "g", Namespace: "ns", Kind: "user", Name: "a", ShardID: "5"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := (RulesAuthorizer{}).Authorize(ctx, VerbGet, attrs); err != nil {
			b.Fatal(err)
		}
	}
}
