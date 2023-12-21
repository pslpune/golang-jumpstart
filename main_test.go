package main

import (
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBasicHTTPEndpointHit(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/ping", nil)
	cl := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := cl.Do(req)
	assert.Nil(t, err, "unexpected error when connecting to the server , check internet connection")
	assert.NotNil(t, resp, "unexpected nil response")
	assert.Equal(t, 200, resp.StatusCode, "Unexpected status code")
	byt, _ := io.ReadAll(resp.Body)
	t.Log(string(byt))
}
