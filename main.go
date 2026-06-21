package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: htmlrun <html-file>")
		os.Exit(1)
	}

	file := os.Args[1]

	if _, err := os.Stat(file); err != nil {
		fmt.Println("File not found:", file)
		os.Exit(1)
	}

	dir := filepath.Dir(file)

	port := "8080"

	http.Handle("/", http.FileServer(http.Dir(dir)))

	url := "http://localhost:" + port

	fmt.Println(filepath.Base(file), url)

	http.ListenAndServe(":"+port, nil)
}
