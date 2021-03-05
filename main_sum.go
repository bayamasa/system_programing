package main

import (
	"fmt"
)

func sum() {
	var n, a, b int
	total := 0
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := 0; i < n+1; i++ {
		n1 := i % 10
		n10 := i / 10 % 10
		n100 := i / 100 % 10
		n1000 := i / 1000 % 10
		n10000 := i / 10000 % 10
		sum := n1 + n10 + n100 + n1000 + n10000

		if (sum >= a) && (sum <= b) {
			total = total + i
		}
	}
	fmt.Println(total)

}
