package main

import (
	"fmt"
	"sort"
)

func cards() {

	var (
		n     int
		total int
	)

	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	// Sort by age preserving name order
	sort.SliceStable(nums, func(i, j int) bool { return nums[i] > nums[j] })
	for i, v := range nums {
		if i%2 == 0 {
			total = total + v
		} else {
			total = total - v
		}
	}
	fmt.Println(total)
}
