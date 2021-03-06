package main

import (
	"fmt"
)

func otoshidama() {

	var (
		n int
		y int
	)

	fmt.Scan(&n)
	fmt.Scan(&y)
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if (10000*i + 5000*j + 1000*(n-i-j)) == y {
				if n-i-j >= 0 {
					fmt.Printf("%d %d %d", i, j, (n - i - j))
					return
				}

			}

		}
	}
	fmt.Println("-1 -1 -1")
}
