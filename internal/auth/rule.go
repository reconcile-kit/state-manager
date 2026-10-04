package auth

import (
	"fmt"
)

// Verb is a bitmask of operations, so a rule check is a single AND.
type Verb uint8

const (
	VerbGet Verb = 1 << iota
	VerbList
	VerbCreate
	VerbUpdate
	VerbUpdateStatus
	VerbDelete

	VerbAll = VerbGet | VerbList | VerbCreate | VerbUpdate | VerbUpdateStatus | VerbDelete
)

// Wildcard matches any value of a rule field. An empty field means the same.
const Wildcard = "*"

var verbNames = map[string]Verb{
	"get":           VerbGet,
	"list":          VerbList,
	"create":        VerbCreate,
	"update":        VerbUpdate,
	"update_status": VerbUpdateStatus,
	"delete":        VerbDelete,
	Wildcard:        VerbAll,
}

func (v Verb) String() string {
	for name, verb := range verbNames {
		if verb == v {
			return name
		}
	}
	return fmt.Sprintf("verb(%d)", uint8(v))
}

// ParseVerbs converts verb names into a bitmask. At least one verb is required.
func ParseVerbs(names []string) (Verb, error) {
	var verbs Verb
	for _, name := range names {
		verb, ok := verbNames[name]
		if !ok {
			return 0, fmt.Errorf("unknown verb %q", name)
		}
		verbs |= verb
	}
	if verbs == 0 {
		return 0, fmt.Errorf("rule must contain at least one verb")
	}
	return verbs, nil
}

// Attributes describe a single resource or a list filter being authorized.
// An empty field in a list filter means "not filtered by this field".
type Attributes struct {
	ResourceGroup string
	Namespace     string
	Kind          string
	Name          string
	ShardID       string
}

// RuleSpec is the external (JSON / DB) form of a rule.
type RuleSpec struct {
	Verbs         []string `json:"verbs"`
	ResourceGroup string   `json:"resource_group"`
	Namespace     string   `json:"namespace"`
	Kind          string   `json:"kind"`
	Name          string   `json:"name"`
	ShardID       string   `json:"shard_id"`
}

// Rule is the compiled form of RuleSpec. Empty fields match any value.
type Rule struct {
	Verbs         Verb
	ResourceGroup string
	Namespace     string
	Kind          string
	Name          string
	ShardID       string
}

func (s *RuleSpec) Compile() (Rule, error) {
	verbs, err := ParseVerbs(s.Verbs)
	if err != nil {
		return Rule{}, err
	}
	return Rule{
		Verbs:         verbs,
		ResourceGroup: normalizeField(s.ResourceGroup),
		Namespace:     normalizeField(s.Namespace),
		Kind:          normalizeField(s.Kind),
		Name:          normalizeField(s.Name),
		ShardID:       normalizeField(s.ShardID),
	}, nil
}

func normalizeField(value string) string {
	if value == Wildcard {
		return ""
	}
	return value
}

// Allows reports whether the rule grants verb on attrs.
//
// For a single resource attrs are fully populated, so this is a plain match.
// For a list filter an empty attribute means "any value", which is covered
// only by an unrestricted rule field: the whole filter must fit into the rule.
func (r *Rule) Allows(verb Verb, attrs *Attributes) bool {
	return r.Verbs&verb != 0 &&
		matchField(r.ResourceGroup, attrs.ResourceGroup) &&
		matchField(r.Namespace, attrs.Namespace) &&
		matchField(r.Kind, attrs.Kind) &&
		matchField(r.Name, attrs.Name) &&
		matchField(r.ShardID, attrs.ShardID)
}

func matchField(ruleValue, value string) bool {
	return ruleValue == "" || ruleValue == value
}
