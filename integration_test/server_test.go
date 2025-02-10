package integration_test

import (
	"bytes"
	"encoding/json"
	"fetch-assessment/api"
	"fetch-assessment/repository"
	"fetch-assessment/server"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// start the server and send some requests
func TestServerStartupAndEndpoints(t *testing.T) {

	// create an in-memory repository
	repository := repository.NewMemoryRepository()

	// create a new server instance using the in-memory repository
	server := server.NewServer(repository)

	// generate a strict handler from the server implementation
	strictHandler := api.NewStrictHandler(server, nil)

	// create a new multiplexer router
	mux := http.NewServeMux()

	// register the handler with the router
	handler := api.HandlerFromMux(strictHandler, mux)

	// start a test server with the API handler
	testServer := httptest.NewServer(handler)
	defer testServer.Close() // ensure we cleanup when done

	// TEST 1: try to get a non-existant receipt and ensure we fail
	resp, err := http.Get(testServer.URL + "/receipts/12345/points")
	if err != nil {
		t.Fatalf("Failed request: %v", err)
	}

	// we're expecting a 404 since the id doesn't exist
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", resp.StatusCode)
	}

	// load request body from the a json file
	jsonData, err := os.ReadFile("../examples/points-test-109.json")
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	// TEST 2: post the receipt to the server
	req, _ := http.NewRequest(http.MethodPost, testServer.URL+"/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Failed POST request: %v", err)
	}
	defer resp.Body.Close()

	// ensure we got an OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, got %d", resp.StatusCode)
	}

	// extract the receipt ID from the response
	var postResponse map[string]string
	bodyBytes, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(bodyBytes, &postResponse); err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// ensure we go an ID for the receipt
	id, ok := postResponse["id"]
	if !ok {
		t.Fatalf("Response did not contain an 'id' field")
	}

	// TEST 3: now check that the points for the saved receipt are correct
	pointsResp, err := http.Get(testServer.URL + "/receipts/" + id + "/points")
	if err != nil {
		t.Fatalf("Failed request for points: %v", err)
	}
	defer pointsResp.Body.Close()

	if pointsResp.StatusCode != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, got %d", pointsResp.StatusCode)
	}

	// validate the resulting points
	var pointsResponse map[string]int64
	pointsBytes, _ := io.ReadAll(pointsResp.Body)
	if err := json.Unmarshal(pointsBytes, &pointsResponse); err != nil {
		t.Fatalf("Failed to parse points response: %v", err)
	}

	// ensure we got points
	points, ok := pointsResponse["points"]
	if !ok {
		t.Fatalf("Response did not contain a 'points' field")
	}

	// ensure we got the expected 109 points
	if points != 109 {
		t.Errorf("Expected 109 points, got %d", points)
	}
}
