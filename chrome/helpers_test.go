package chrome

import (
	"os"
	"testing"
)

func testingHost(t *testing.T) string {
	host := os.Getenv("CHROME_DEVTOOLS_HOST")
	if host == "" {
		t.Fatal("no CHROME_DEVTOOLS_HOST environment variable")
	}
	return host
}

func testingConn(t *testing.T, page bool) *Conn {
	host := testingHost(t)
	endpoints, err := Endpoints(host)
	if err != nil {
		t.Fatal(err)
	}
	for _, endpoint := range endpoints {
		if endpoint.WebSocketURL == "" {
			continue
		}
		if page && endpoint.Type != "page" {
			continue
		}
		conn, err := NewConn(endpoint.WebSocketURL)
		if err != nil {
			t.Fatal(err)
		}
		return conn
	}
	t.Fatal("no endpoints to test")
	return nil
}
