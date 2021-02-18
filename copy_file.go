package main

import (
	"io"
	"os"
)

func copyFile() {
	beforeFile, err := os.Open("before_cp.txt")
	if err != nil {
		panic(err)
	}
	defer beforeFile.Close()

	afterFile, err := os.Create("after_cp.txt")
	if err != nil {
		panic(err)
	}
	defer afterFile.Close()

	io.Copy(afterFile, beforeFile)
}
