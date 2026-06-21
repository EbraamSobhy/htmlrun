package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	file := "index.html"

	dir := filepath.Dir(file)

	port := "8080"

	http.Handle("/", http.FileServer(http.Dir(dir)))

	url := "http://localhost:" + port

	fmt.Println("Serving:", url)

	http.ListenAndServe(":"+port, nil)
}
