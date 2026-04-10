package validators

import (
	"fmt"

	apiresource "github.com/reconcile-kit/api/resource"
)

const maxAllowedLabelSelectors = 6

// ValidateLabelSelectors validates IN-only label selector requirements.
func ValidateLabelSelectors(labelSelectors []apiresource.LabelSelector) error {
	if len(labelSelectors) > maxAllowedLabelSelectors {
		return fmt.Errorf("max allowed label selectors exceeds %d", maxAllowedLabelSelectors)
	}

	return nil
}
