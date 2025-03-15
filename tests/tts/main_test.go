package tts_test

import (
	"testing"
	"net/http"
	"bytes"
	"io/ioutil"
)

const (
	chatTTSURL = "https://api.2noise.com/chattts"
)

func TestChatTTSBasic(t *testing.T) {
	// Test basic text to speech conversion
	payload := []byte(`{"text": "Hello, world!", "voice": "en-US-Wavenet-D"}`)
	
	resp, err := http.Post(chatTTSURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}
}

func TestChatTTSInvalidInput(t *testing.T) {
	// Test handling of invalid input
	payload := []byte(`{"text": "", "voice": "invalid-voice"}`)
	
	resp, err := http.Post(chatTTSURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Error("Expected non-200 status code for invalid input")
	}
}

func TestChatTTSLargeInput(t *testing.T) {
	// Test handling of large text input
	longText := make([]byte, 10000) // 10KB of text
	payload := []byte(`{"text": "` + string(longText) + `", "voice": "en-US-Wavenet-D"}`)
	
	resp, err := http.Post(chatTTSURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200 for large input, got %d", resp.StatusCode)
	}
}
