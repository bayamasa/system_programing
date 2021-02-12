package main

import (
	"encoding/csv"
	"log"
	"os"
)

var sample = []string{
	"sample_1",
	"sample_2",
	"sample_3",
}

func createCsv() {
	//書き込みファイル作成
	file, err := os.Create("sample.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file) // utf8
	writer.Write(sample)
	writer.Flush() // ファイル出力
}
