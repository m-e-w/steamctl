package cli

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

type SortFunc[T any] func(a, b T) bool

func SortByKey[T any](items []T, key string, sorters map[string]SortFunc[T]) error {
	sortFn, ok := sorters[key]

	// If key not in map, return an error with the expected values
	// We use a separate slice to add the keys from the sorters map and then sort them before printing
	if !ok {
		keys := make([]string, 0, len(sorters))
		for k := range sorters {
			keys = append(keys, k)
		}
		slices.Sort(keys)

		return fmt.Errorf(
			"invalid value for --sort %q (valid values: %s)",
			key,
			strings.Join(keys, ", "),
		)
	}

	// If we are here, it means the sort value was valid so we sort by the specified function
	sort.Slice(items, func(i, j int) bool {
		return sortFn(items[i], items[j])
	})

	return nil
}
