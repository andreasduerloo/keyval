package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	input, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var config Config
	json.Unmarshal(input, &config)

	go run(config)

	// Give the server time to set up
	time.Sleep(10 * time.Millisecond)

	m.Run()
}

func readBody(response *http.Response) string {
	body := new(bytes.Buffer)
	_, _ = body.ReadFrom(response.Body)
	return body.String()
}

func TestSetNewTable(t *testing.T) {
	url := "HTTP://127.0.0.1:4555"

	// Make an HTTP PUT request to the server
	req, err := http.NewRequest(http.MethodPut, url+"/table1/record1/value1", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Fatal(err)
	}

	// Make the PUT request
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Make the GET request
	response, err := http.Get(url + "/table1/record1")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Assert the response body
	expectedBody := "value1"
	actualBody := readBody(response)
	if actualBody != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, actualBody)
	}
}

func TestSetExistingTable(t *testing.T) {
	url := "HTTP://127.0.0.1:4555"

	// Make an HTTP PUT request to the server
	req, err := http.NewRequest(http.MethodPut, url+"/table1/record2/value2", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Fatal(err)
	}

	// Make the PUT request
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Make the GET request
	response, err := http.Get(url + "/table1/record2")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Assert the response body
	expectedBody := "value2"
	actualBody := readBody(response)
	if actualBody != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, actualBody)
	}
}

func TestSetExistingRecord(t *testing.T) {
	url := "HTTP://127.0.0.1:4555"

	// Make an HTTP PUT request to the server
	req, err := http.NewRequest(http.MethodPut, url+"/table1/record2/something_else", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Fatal(err)
	}

	// Make the PUT request
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Make the GET request
	response, err := http.Get(url + "/table1/record2")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Assert the response body
	expectedBody := "something_else"
	actualBody := readBody(response)
	if actualBody != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, actualBody)
	}
}

func TestGetMissingRecord(t *testing.T) {
	url := "HTTP://127.0.0.1:4555"

	// Make the GET request
	response, err := http.Get(url + "/table1/record3")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, response.StatusCode)
	}
}

func TestGetMissingTable(t *testing.T) {
	url := "HTTP://127.0.0.1:4555"

	// Make the GET request
	response, err := http.Get(url + "/table3/record1")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, response.StatusCode)
	}
}

func TestDelExistingRecord(t *testing.T) {
	url := "HTTP://127.0.0.1:4555"

	// Make an HTTP PUT request to the server
	req, err := http.NewRequest(http.MethodPut, url+"/table1/record3/value3", bytes.NewBuffer([]byte("")))

	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Make the GET request
	response, err := http.Get(url + "/table1/record3")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Assert the response body
	expectedBody := "value3"
	actualBody := readBody(response)
	if actualBody != expectedBody {
		t.Errorf("Expected response body %q, got %q", expectedBody, actualBody)
	}

	// Make the DELETE request
	req, err = http.NewRequest(http.MethodDelete, url+"/table1/record3", bytes.NewBuffer([]byte("")))

	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Make the GET request
	response, err = http.Get(url + "/table1/record3")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, response.StatusCode)
	}
}

func TestDelExistingTable(t *testing.T) {
	url := "HTTP://127.0.0.1:4555"

	// Make the GET request
	response, err := http.Get(url + "/table1/record1")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
	}

	// Make the DELETE request
	req, err := http.NewRequest(http.MethodDelete, url+"/table1", bytes.NewBuffer([]byte("")))

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Make the GET request
	response, err = http.Get(url + "/table1/record1")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Assert the response status code
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, response.StatusCode)
	}
}
