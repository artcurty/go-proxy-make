package e2e

import (
	"encoding/json"
	"github.com/artcurty/go-proxy-make/cmd/api/generated"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProxyGETorder(t *testing.T) {
	expectedResponse := map[string]interface{}{
		"orderId": "xop1",
		"name":    "ps5",
		"count":   "1",
	}

	testServer := httptest.NewServer(http.HandlerFunc(generated.ProxyGETorder))
	defer testServer.Close()
	testClient := testServer.Client()

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("response code is not 200: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %v", err)
	}

	var actualResponse map[string]interface{}
	err = json.Unmarshal(data, &actualResponse)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
	}

	for key, expectedValue := range expectedResponse {
		if actualValue, ok := actualResponse[key]; !ok || actualValue != expectedValue {
			t.Errorf("Value for key %s does not match. Got: %v, Expected: %v", key, actualValue, expectedValue)
		}
	}
}

func TestProxyPOSTorder(t *testing.T) {
	expectedResponse := map[string]interface{}{
		"orderId": "xop1",
		"name":    "ps5",
		"count":   "1",
	}

	testServer := httptest.NewServer(http.HandlerFunc(generated.ProxyPOSTorder))
	defer testServer.Close()
	testClient := testServer.Client()

	resp, err := testClient.Post(testServer.URL, "application/json", strings.NewReader(`{"id":"xop1","productName":"ps5","count":"1"}`))
	if err != nil {
		t.Errorf("Post error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("response code is not 200: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %v", err)
	}

	var actualResponse map[string]interface{}
	err = json.Unmarshal(data, &actualResponse)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
	}

	for key, expectedValue := range expectedResponse {
		if actualValue, ok := actualResponse[key]; !ok || actualValue != expectedValue {
			t.Errorf("Value for key %s does not match. Got: %v, Expected: %v", key, actualValue, expectedValue)
		}
	}
}

func TestProxyPOSTpayment(t *testing.T) {
	expectedResponse := map[string]interface{}{
		"id":     "xop1",
		"status": "paid",
	}

	testServer := httptest.NewServer(http.HandlerFunc(generated.ProxyPOSTpayment))
	defer testServer.Close()
	testClient := testServer.Client()

	req, err := http.NewRequest(http.MethodPut, testServer.URL, strings.NewReader(`{"id":"xop1","status":"paid"}`))
	if err != nil {
		t.Errorf("NewRequest error: %v", err)
	}
	resp, err := testClient.Do(req)
	if err != nil {
		t.Errorf("Do error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("response code is not 200: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %v", err)
	}

	var actualResponse map[string]interface{}
	err = json.Unmarshal(data, &actualResponse)
	if err != nil {
		t.Errorf("json.Unmarshal error: %v", err)
	}

	for key, expectedValue := range expectedResponse {
		if actualValue, ok := actualResponse[key]; !ok || actualValue != expectedValue {
			t.Errorf("Value for key %s does not match. Got: %v, Expected: %v", key, actualValue, expectedValue)
		}
	}
}
