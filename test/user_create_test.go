package test

import (
	"23-ci-cd/internal/user"
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestCreateUserFunctional(t *testing.T) {
	time.Sleep(2 * time.Second)

	url := "http://localhost:8080/users"

	payload := map[string]interface{}{
		"email": "functional@example.com",
		"age":   25,
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %d", resp.StatusCode)
	}

	var res user.User
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if res.Email != payload["email"] || res.Age != payload["age"] {
		t.Errorf("Unexpected response: got %+v, expected email %s and age %d", res, payload["email"], payload["age"])
	}
}
