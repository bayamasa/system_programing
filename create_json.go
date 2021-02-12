package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json 化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	writer := gzip.NewWriter(w)
	writer.Header.Name = "test"
	multiWriter := io.MultiWriter(writer, os.Stdout)
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "      ")
	encoder.Encode(source)
	writer.Flush()
	writer.Close()

}

// ここにコードを書く }
func createJson() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
