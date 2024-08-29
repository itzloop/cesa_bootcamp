package main

import (
	"fmt"
	"io"
	"net/http"
)

func f1() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("err; ", err)
	}
	fmt.Println(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func main() {
	fmt.Println()
	f1()
}
