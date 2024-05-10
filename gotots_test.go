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
	sut := NewGotots("./examples/1_type_param.go")
	// if file is empty or not found, return empty string
	if resp := sut.ConvertToTs("output/example1.d.ts");resp != nil {
		t.Errorf("Expected empty string, got %s", resp)
	}
}

func TestReadFileMultipleTypes(t *testing.T) {
	t.Log("TestReadFile")
	sut := NewGotots("./examples/multiple_Type.go")
	// if file is empty or not found, return empty string
	if resp := sut.ConvertToTs("output/example2.d.ts");resp != nil {
		t.Errorf("Expected empty string, got %s", resp)
	}
}

func TestReadFileEmpty(t *testing.T) {
	t.Log("TestReadFile")
	sut := NewGotots("useless string")
	// if file is empty or not found, return empty string
	if resp := sut.ConvertToTs("error");resp == nil {
		t.Errorf("Expected empty string, got %s", resp)
	}
}
