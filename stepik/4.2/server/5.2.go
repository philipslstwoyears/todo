package server

import (
	"fmt"
	"net/http"
	"strconv"
)

var NUM int

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, NUM)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			return
		}
		count := r.FormValue("count")
		n, err := strconv.Atoi(count)
		if err != nil {
			http.Error(w, "это не число", http.StatusBadRequest)
			return
		}
		NUM += n
		fmt.Fprint(w, NUM)
	}
}

func main() {
	http.HandleFunc("/count", handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}
}
