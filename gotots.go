package gotots

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type Gotots struct {
	File string
}

func NewGotots(file string) *Gotots {
	return &Gotots{File: file}
}

func (g *Gotots) ConvertToTs(outputFile string) error {
	g.clearFile(outputFile)

	types, err := g.readFile()

	if err != nil {
		return err
	}

	for _, tsType := range types {
		err = g.writeToFile(outputFile, tsType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Gotots) readFile() ([]string, error) {
	if g.File == "" {
		return nil, errors.New("File path is empty")
	}

	file, err := os.Open(g.File)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	holder := make([]byte, 1024)
	_, err = file.Read(holder)
	if err != nil {
		return nil, err
	}

	typeStruct, err := readTheStruct(holder)

	if err != nil {
		return nil, err
	}

	typescriptTypes := make([]string, 0)

	for _, structInput := range typeStruct {
		typescriptType, err := convertStructToTsType(structInput)
		if err != nil {
			return nil, err
		}

		typescriptTypes = append(typescriptTypes, typescriptType)
	}

	return typescriptTypes, nil
}

func (g *Gotots) writeToFile(outputFile string, content string) error {

	//if subdirectory does not exist, create it
	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		return nil
	}

	log.Println("Content: \n", content)

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	// append content to file
	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil
}

func (g *Gotots) clearFile(file string) error {
	if err := os.Remove(file); err != nil {
		return err
	}

	return nil
}
