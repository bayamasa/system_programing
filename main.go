package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	total := 0
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := 0; i < n+1; i++ {
		n1 := i
		n10 := i / 10
		n100 := i / 100
		n1000 := i / 1000
		n10000 := i / 10000

		sum := n1 + n10 + n100 + n1000 + n10000
		if (sum >= a) && (sum <= b) {
			total = total + i
		}
	}
	fmt.Println(n)
	fmt.Println(a)
	fmt.Println(b)

}
