package main

import (
	"testing"
)

func TestCreateHelloString(t *testing.T) {
	expected := "Hello 🗺️ !"
	result := CreateHelloString()
	if result == "" {
		t.Fatalf("message must not be empty")
	}
	if result != expected {
		t.Fatalf("message content is wrong, expect:%s but was %s", expected, result)
	}
}
