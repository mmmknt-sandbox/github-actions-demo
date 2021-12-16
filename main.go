package main

import (
	"net/http"

	"github.com/mmmknt-sandbox/github-actions-demo/logic"
)

func main() {
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(logic.Echo("first version")))
	})
	http.HandleFunc("/unhealthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusInternalServerError)
	})
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(logic.Echo("test")))
	})
	_ = http.ListenAndServe(":8080", nil)
}
