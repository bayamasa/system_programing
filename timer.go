package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) {
	fmt.Println("kiteru")

	wait := make(chan string)
	go func() {
		wait <- "jee"
	}()
	select {
	case <-wait:
		fmt.Println("hello")

	case <-time.After(duration * time.Second):
		fmt.Println(duration, "seconds elapsed")
	}
}
