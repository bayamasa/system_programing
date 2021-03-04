package main

import (
	"fmt"
)

func coins() {
	count := 0
	var a, b, c, x int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				a1 := 500 * i
				b1 := 100 * j
				c1 := 50 * k
				// fmt.Println("--------")
				// fmt.Printf("a= %d ", a1)
				// fmt.Printf("b= %d ", b1)
				// fmt.Printf("c=% d ", c1)
				// fmt.Println("")
				// fmt.Println("--------")

				if a1+b1+c1 == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
