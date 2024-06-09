package httpserver

import (
	"fmt"
	"io"
	"net/http"
)

type HTTPClient struct {
	url string
}

func (client *HTTPClient) Hello() int {
	resp, err := http.Get(client.url + "/hello")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return -1
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return -2
	}

	fmt.Println("Response from server:", string(body))
	return resp.StatusCode
}

func MakeClient() HTTPClient {
	return HTTPClient{"http://localhost:8081/"}
}
