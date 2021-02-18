package main

import (
	"crypto/rand"
	"os"
)

func createRandomFile() {
	byt := 1024
	b := make([]byte, byt)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	randFile, err := os.Create("rand")
	if err != nil {
		panic(err)
	}
	defer randFile.Close()

	randFile.Write(b)
}
