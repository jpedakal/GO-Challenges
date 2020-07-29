package main

import (
	"fmt"
	"net/http"
)

func main() {
	link := "https://google.com"
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("The URL might be down", err)
	} else {
		fmt.Println("Response of link is", resp)
	}
}