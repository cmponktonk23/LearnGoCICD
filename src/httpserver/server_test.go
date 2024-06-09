package httpserver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMakeServer(t *testing.T) {
	var server interface{} = MakeServer()
	switch server.(type) {
	case HTTPServer:
	default:
		t.Fatal("Expect HTTPClient but not")
	}
}

func TestHTTPServer_helloHandlerSucceed(t *testing.T) {
	server := MakeServer()

	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	resp := httptest.NewRecorder()

	server.helloHandler(resp, req)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Read response body error:%v", err)
	}

	if expected := "hello, client!"; string(body) != expected {
		t.Fatalf("Expected: %s, got: %s", expected, string(body))
	}
}

func TestHTTPServer_helloHandlerError(t *testing.T) {
	server := MakeServer()

	req := httptest.NewRequest(http.MethodPost, "/hello", nil)
	resp := httptest.NewRecorder()

	server.helloHandler(resp, req)

	body, err := io.ReadAll(resp.Body)

	if resp.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("Expected status code:%d, got: %d", http.StatusMethodNotAllowed, resp.Result().StatusCode)
	}

	if err != nil {
		t.Fatalf("Read response body error:%v", err)
	}

	if expected := "Method not allowed\n"; string(body) != expected {
		t.Fatalf("Expected: %s, got: %s", expected, string(body))
	}
}

func TestHTTPServer_Run(t *testing.T) {
	server := HTTPServer{nil, "localhost", "55555"}
	go server.Run()
	time.Sleep(100 * time.Millisecond)
	server.Close()
}

func TestHTTPServer_RunError(t *testing.T) {
	server := HTTPServer{nil, "localhost", "ab"}
	go server.Run()
	time.Sleep(100 * time.Millisecond)
	server.Close()
}
