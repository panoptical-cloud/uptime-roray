package utils

import (
	"fmt"
	"io"
	"net/http"
)

func RegisterWithServer(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while registering with server", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response body", err)
		return
	}
	fmt.Println("Response from server:", string(body))
}
