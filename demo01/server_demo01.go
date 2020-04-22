package main

import (
	"net/http"
)

func main() {

	http.Handle("E:\\workplace\\", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
