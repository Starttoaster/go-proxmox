package proxmox

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewTLSServer(mux)

	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	client, err := NewClient("test-token-id", "test-token",
		WithBaseURL(fmt.Sprintf("%s/", server.URL)),
		WithHTTPClient(&httpClient),
	)
	if err != nil {
		t.Fatal(err)
	}

	return mux, server, client
}

func teardown(server *httptest.Server) {
	server.Close()
}

func fixture(path string) string {
	b, err := os.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
