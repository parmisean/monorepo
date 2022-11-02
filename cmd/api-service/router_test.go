package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHealthRoute(t *testing.T) {
	app := application{}
	testSrv := httptest.NewServer(app.routes())

	defer testSrv.Close()

	res, err := testSrv.Client().Get(testSrv.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	if res.StatusCode != http.StatusOK {
		t.Errorf("got: %v; want: %v", res.StatusCode, http.StatusOK)
	}

	if string(body) != healthResponse {
		t.Errorf("got: %v; want: %v", string(body), healthResponse)
	}
}
