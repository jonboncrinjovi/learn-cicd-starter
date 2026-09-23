package auth

import (
    "net/http"
    "testing"
)


func TestGetAPIKeyValid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key-123")
	key, err := GetAPIKey(headers)
	if err = nil {
    	t.Fatalf("expected no error, got: %v", err)
}
	if key != "my-secret-key-123" {
    	t.Fatalf("expected key 'my-secret-key-123', got: %v", key)
}
}

func TestGetAPIKeyNoHeader(t *testing.T) {
    	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
    	t.Fatalf("expected ErrNoAuthHeaderIncluded, got: %v", err)
}
}
