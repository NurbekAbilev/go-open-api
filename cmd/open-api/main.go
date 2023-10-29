package main

import (
	"net/http"

	"github.com/swaggest/swgui/v5emb"
)

func main() {
	http.Handle("/api1/docs/", v5emb.New(
		"Petstore",
		"https://petstore3.swagger.io/api/v3/openapi.json",
		"/api1/docs/",
	))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
	})

	println("docs at http://localhost:8080/api1/docs/")
	_ = http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}