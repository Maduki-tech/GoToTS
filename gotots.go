package gotots

import "os"

type Gotots struct {
	File string
}

func NewGotots(file string) *Gotots {
	return &Gotots{File: file}
}

func (c *Gotots) readFile() string {
	if c.File == "" {
		return ""
	}

	file, err := os.Open(c.File)
	if err != nil {
		return ""
	}

	defer file.Close()

	holder := make([]byte, 1024)
	count, err := file.Read(holder)
	if err != nil {
		return ""
	}

	typeStruct, err := readTheStruct(holder)

	if err != nil {
		return ""
	}

	typescriptTypes := make([]string, 0)

	for _, structInput := range typeStruct {
		typescriptType, err := convertStructToTsType(structInput)
		if err != nil {
			return err.Error()
		}

		typescriptTypes = append(typescriptTypes, typescriptType)
	}

	for _, tsType := range typescriptTypes {
		err := writeToFile(tsType)
		if err != nil {
			return err.Error()
		}
	}


	return string(holder[:count])
}

func writeToFile(content string) error {
	file, err := os.Create("output.d.ts")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
