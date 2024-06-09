package httpserver

import (
	"fmt"
	"net/http"
)

type HTTPServer struct {
	server *http.Server
	ip     string
	port   string
}

func (server *HTTPServer) Close() {
	if server.server != nil {
		server.server.Close()
	}
}

func (server *HTTPServer) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", server.helloHandler)

	server.server = &http.Server{Addr: fmt.Sprintf("%s:%s", server.ip, server.port), Handler: mux}

	fmt.Printf("Start HTTP Server listening %s:%s\n", server.ip, server.port)
	if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("HTTP server ListenAndServe error: %v\n", err)
	}
	// if err := http.ListenAndServe(fmt.Sprintf("%s:%s", server.ip, server.port), nil); err != nil {
	// 	fmt.Println("Server failed to start:", err)
	// }

}

func (server *HTTPServer) helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "hello, client!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func MakeServer() HTTPServer {
	return HTTPServer{nil, "localhost", "8081"}
}
