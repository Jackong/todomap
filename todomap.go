package todomap

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Welcome to To Do Map")
	})
}
