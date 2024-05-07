package gotots_test

import (
	"gotots"
	"testing"
)

func TestGotoTs(t *testing.T) {
	t.Log("TestGotoTs")
	sut := gotots.NewGotots("test")
	if sut.GetFile() != "test" {
		t.Errorf("Expected test, got %s", sut.GetFile())
	}
}

func TestReadFile(t *testing.T) {
	t.Log("TestReadFile")
	sut := gotots.NewGotots("./examples/testing.go")
	// if file is empty or not found, return empty string
	if sut.ReadFile() == "" {
		t.Errorf("Expected empty string, got %s", sut.ReadFile())
	}
}


func TestReadFileEmpty(t *testing.T) {
	t.Log("TestReadFile")
	sut := gotots.NewGotots("useless string")
	// if file is empty or not found, return empty string
	if sut.ReadFile() != "" {
		t.Errorf("Expected empty string, got %s", sut.ReadFile())
	}
}
