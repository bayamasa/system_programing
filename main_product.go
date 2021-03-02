package main

import (
	"fmt"
)

func product() {
	var a, b, c, d int
	fmt.Scanf("%d", &a)
	b = a / 100
	c = (a / 10) % 2
	d = a % 2
	fmt.Println(b + c + d)
}
