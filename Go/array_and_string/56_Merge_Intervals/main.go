package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(	i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prov := intervals[0]
	merges := [][]int{}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= prov[1] {
			if intervals[i][1] > prov[1] {
				prov[1] = intervals[i][1]
			}
		} else {
			merges = append(merges, prov)
			prov = intervals[i]
		}
	}
	merges = append(merges, prov)
	return merges
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merged := merge(intervals)
	for _, interval := range merged {
		fmt.Println(interval[0], interval[1])
	}
}
