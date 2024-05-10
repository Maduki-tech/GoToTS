package gotots

import (
	"errors"
	"os"
	"path/filepath"
)

type Gotots struct {
	File string
}

func NewGotots(file string) *Gotots {
	return &Gotots{File: file}
}

func (c *Gotots) ConvertToTs(outputFile string) error {
	types, err := c.readFile()

	if err != nil {
		return err
	}

	for _, tsType := range types {
		err = c.writeToFile(outputFile, tsType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Gotots) readFile() ([]string, error) {
	if c.File == "" {
		return nil, errors.New("File path is empty")
	}

	file, err := os.Open(c.File)
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

func (c *Gotots) writeToFile(outputFile string, content string) error {

	//if subdirectory does not exist, create it
	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		return nil
	}

	os.Create(outputFile)

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
