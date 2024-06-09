package main

import (
	"fmt"
	"learn/httpserver"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go server/client")
	}

	if os.Args[1] == "server" {
		server := httpserver.MakeServer()
		server.Run()
	} else if os.Args[1] == "client" {
		client := httpserver.MakeClient()
		client.Hello()
	} else {
		fmt.Println("Usage: go run main.go server/client")
	}
}
