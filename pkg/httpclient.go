package pkg

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type HTTPRequest struct {
	Host    string
	Path    string
	Method  string
	Headers map[string]string
	Params  map[string]string
	Body    map[string]interface{}
}

type HTTPClient interface {
	DoRequest(req HTTPRequest) (*http.Response, error)
}

type DefaultHTTPClient struct{}

func (c *DefaultHTTPClient) DoRequest(req HTTPRequest) (*http.Response, error) {
	url := req.Host + req.Path

	var bodyData []byte
	if req.Body != nil {
		bodyData, _ = json.Marshal(req.Body)
	}

	httpReq, _ := http.NewRequest(req.Method, url, bytes.NewBuffer(bodyData))

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	log.Printf("Sending %s request to %s with body: %s", req.Method, url, string(bodyData))

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		log.Printf("Error making request to %s: %v", url, err)
		return nil, err
	}

	log.Printf("Received response from %s with status: %s", url, resp.Status)
	return resp, nil
}
