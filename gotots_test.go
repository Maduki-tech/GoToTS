package gotots

import (
	"testing"
)

func TestGotoTs(t *testing.T) {
	t.Log("TestGotoTs")
	sut := NewGotots("test")

	if sut == nil {
		t.Error("Expected Gotots object, got nil")
	}
}

func TestReadFile(t *testing.T) {
	t.Log("TestReadFile")
	sut := NewGotots("./examples/testing.go")
	// if file is empty or not found, return empty string
	if sut.readFile() == "" {
		t.Errorf("Expected empty string, got %s", sut.readFile())
	}
}

func TestReadFileEmpty(t *testing.T) {
	t.Log("TestReadFile")
	sut := NewGotots("useless string")
	// if file is empty or not found, return empty string
	if sut.readFile() != "" {
		t.Errorf("Expected empty string, got %s", sut.readFile())
	}
}
