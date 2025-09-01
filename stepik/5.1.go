package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/api/user", handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
