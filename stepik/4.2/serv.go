package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	var name string
	var age string
	fmt.Scan(&name, &age)
	base, err := url.Parse("http://127.0.0.1:8080/hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)
	base.RawQuery = params.Encode()
	get, err := http.Get(base.String())
	if err != nil {
		return
	}
	defer get.Body.Close()
	res, err := io.ReadAll(get.Body)
	if err != nil {
		return
	}
	fmt.Print(string(res))
}
