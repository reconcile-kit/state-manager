package states

func DiffLabels(oldLabels, newLabels map[string]string) (map[string]string, []string) {
	toUpdate := make(map[string]string)
	toDelete := make([]string, 0)
	for k, newVal := range newLabels {
		oldVal, exists := oldLabels[k]

		if !exists {
			toUpdate[k] = newVal
			continue
		}

		if oldVal != newVal {
			toUpdate[k] = newVal
		}
	}

	for k := range oldLabels {
		if _, exists := newLabels[k]; !exists {
			toDelete = append(toDelete, k)
		}
	}

	return toUpdate, toDelete
}
