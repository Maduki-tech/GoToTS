package gotots

import (
	"log"
	"os"
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
	file := "./output/example1.d.ts"
	// if file is empty or not found, return empty string
	if resp := sut.ConvertToTs(file); resp != nil {
		t.Errorf("Expected empty string, got %s", resp)
	}

	if lines := readFileAndReturnLine(file); lines != 3 {
		t.Errorf("Expected lines 3, got %d", lines)
	}
}

func TestReadFileMultipleTypes(t *testing.T) {
	t.Log("TestReadFile")
	sut := NewGotots("./examples/multiple_Type.go")
	file := "./output/example2.d.ts"
	// if file is empty or not found, return empty string
	if resp := sut.ConvertToTs(file); resp != nil {
		t.Errorf("Expected empty string, got %s", resp)
	}

	if lines := readFileAndReturnLine(file); lines != 7 {
		t.Errorf("Expected lines 7, got %d", lines)
	}

}

func TestReadFileEmpty(t *testing.T) {
	t.Log("TestReadFile")
	sut := NewGotots("useless string")
	file := "useless string"
	// if file is empty or not found, return empty string
	if resp := sut.ConvertToTs(file); resp == nil {
		t.Errorf("Expected empty string, got %s", resp)
	}
}

func readFileAndReturnLine(file string) int {
	inputFile, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	lines := 0

	// count the number of lines in the file
	holder := make([]byte, 1024)
 	_, err = inputFile.Read(holder)

	for _, char := range holder {
		if char == '\n' {
			lines++
		}
	}

	return lines
}
