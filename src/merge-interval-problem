//Merge-Interval Problem
package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start, End int
}

func merge(intervals []Interval) []Interval {
	m := append([]Interval(nil), intervals...)
	if len(m) <= 1 {
		return m
	}

	sort.Slice(m,
		func(i, j int) bool {
			if m[i].Start < m[j].Start {
				return true
			}
			if m[i].Start == m[j].Start && m[i].End < m[j].End {
				return true
			}
			return false
		},
	)

	j := 0
	for i := 1; i < len(m); i++ {
		if m[j].End >= m[i].Start {
			if m[j].End < m[i].End {
				m[j].End = m[i].End
			}
		} else {
			j++
			m[j] = m[i]
		}

	}
	return append([]Interval(nil), m[:j+1]...)
}

func main() {
	beforeIntervalMerge := []Interval{{0, 1}, {3, 5}, {4, 8}, {10, 12}, {9, 10}}
	merged := merge(beforeIntervalMerge)
	fmt.Println(beforeIntervalMerge)
	fmt.Println(merged)
}
