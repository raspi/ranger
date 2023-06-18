package ranger

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type Number interface {
	constraints.Integer
}

type Range[T Number] struct {
	Start T
	End   T
}

// GetIntegerRanges generates a list of ranges found when given a list of integers.
// Duplicate integers are removed. Integers are sorted before generation.
//
// For example {1, 2, 3} returns []{Range{ Start:1, End:3 }}
func GetIntegerRanges[T Number](nums []T) (res []Range[T], err error) {

	// Remove duplicates
	m := make(map[T]bool, len(nums))
	var l []T // Contains numbers only once

	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = true
			l = append(l, n)
		}
	}
	nums = nil
	m = nil

	// Sort
	l = sortSlice(l)

	// Find ranges
	r := Range[T]{}

	for idx, n := range l {
		if idx == 0 {
			// First one
			r.Start = n
			r.End = n
			continue
		}

		// Doesn't match +1 to .End
		if n != (r.End + 1) {
			// Add to return slice
			res = append(res, Range[T]{r.Start, r.End})

			// Start again
			r.Start = n
			r.End = n
			continue
		}

		r.End = n
	}

	res = append(res, Range[T]{r.Start, r.End})

	return res, nil
}

func sortSlice[T Number](s []T) []T {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	return s
}
