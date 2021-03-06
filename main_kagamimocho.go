package main

import (
	"fmt"
)

func kagamimochi() {

	var (
		n int
	)

	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	m := make(map[int]bool)
	newArr := []int{}

	for _, v := range nums {
		if !m[v] {
			m[v] = true
			newArr = append(newArr, v)
		}
	}

	fmt.Println(len(newArr))
}
