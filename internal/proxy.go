package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ProxyRequest(w http.ResponseWriter, r *http.Request, url, method string, fieldMappings map[string]string) {
	var data map[string]interface{}

	if r.Body != nil && r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("Error decoding request body: %v", err)
			return
		}
	}

	reqBody := make(map[string]interface{})
	for inputField, outputField := range fieldMappings {
		reqBody[outputField] = data[inputField]
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error marshaling request body: %v", err)
		return
	}

	proxyReq, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating proxy request: %v", err)
		return
	}
	proxyReq.Header.Set("Content-Type", "application/json")

	log.Printf("Proxying %s request to %s with body: %s", method, url, string(jsonData))

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error making proxy request to %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error reading response body: %v", err)
		return
	}

	var respData map[string]interface{}
	if err := json.Unmarshal(respBody, &respData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error unmarshaling response body: %v", err)
		return
	}

	mappedRespBody := make(map[string]interface{})
	for outputField, inputField := range fieldMappings {
		mappedRespBody[outputField] = respData[inputField]
	}

	finalRespBody, err := json.Marshal(mappedRespBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error marshaling final response body: %v", err)
		return
	}

	w.WriteHeader(resp.StatusCode)
	if _, err := w.Write(finalRespBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error writing response body: %v", err)
		return
	}

	log.Printf("Successfully proxied %s request to %s with status: %s", method, url, resp.Status)
}
