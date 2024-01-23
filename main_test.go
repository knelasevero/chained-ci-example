package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestServerHandler(t *testing.T) {
    // Create a test server using our handler
    ts := httptest.NewServer(http.HandlerFunc(handler))
    defer ts.Close()

    // Make a request to our test server
    resp, err := http.Get(ts.URL)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Check the status code
    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
    }

    // Check the response body
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        t.Fatalf("Expected no error reading response body, got %v", err)
    }
    expected := "Hello from a Platform Engineer!\n"
    if string(body) != expected {
        t.Fatalf("Expected response body to be '%s', got '%s'", expected, string(body))
    }
}