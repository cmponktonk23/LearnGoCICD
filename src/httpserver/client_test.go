package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeClient(t *testing.T) {
	var client interface{} = MakeClient()
	switch client.(type) {
	case HTTPClient:
	default:
		t.Fatal("Expect HTTPClient but not")
	}
}

func TestHTTPClient_Hello(t *testing.T) {
	mockserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected to GET method, got: %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, client"))
	}))

	defer mockserver.Close()

	client := HTTPClient{}
	client.url = mockserver.URL

	if code := client.Hello(); code != http.StatusOK {
		t.Fatalf("Expected the status code to be: %d, got: %d", http.StatusOK, code)
	}
}

func TestHTTPClient_HelloError(t *testing.T) {
	mockserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected to GET method, got: %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, client"))
	}))

	defer mockserver.Close()

	client := HTTPClient{}
	client.url = mockserver.URL + "error"

	if code := client.Hello(); code != -1 {
		t.Fatalf("Expected the status code to be: %d, got: %d", -1, code)
	}
}

func TestHTTPClient_HelloReadRespError(t *testing.T) {
	mockserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected to GET method, got: %s", r.Method)
		}

		w.Header().Add("Content-Length", "50")
		w.Write([]byte("a"))
	}))

	defer mockserver.Close()

	client := HTTPClient{}
	client.url = mockserver.URL

	if code := client.Hello(); code != -2 {
		t.Fatalf("Expected the status code to be: %d, got: %d", -1, code)
	}
}
