package gotots

import (
	"errors"
	"log"
	"strings"
)

type TypeBuilder struct {
	name   string
	fields []Field
}

type Field struct {
	name string
	typ  string
}

// Create a struct with the content of the files struct
func readTheStruct(content []byte) ([]TypeBuilder, error) {
	fileContent := string(content)
	typeBuilders := make([]TypeBuilder, 0)
	isStruct := false
	for _, line := range strings.Split(fileContent, "\n") {
		if line == "" || line == "\n" {
			continue
		}
		splittedLine := strings.Split(line, " ")
		if splittedLine[0] == "type" {
			isStruct = true
			log.Println("Type found: ", splittedLine[1])
			typeBuilder := TypeBuilder{name: splittedLine[1]}
			typeBuilders = append(typeBuilders, typeBuilder)
			continue
		}

		if isStruct {
			if splittedLine[0] == "}" {
				isStruct = false
				continue
			}
			cleanLine := removeEmpty(splittedLine)

			field := Field{name: cleanLine[0], typ: cleanLine[1]}
			// mapTypes(&field)
			typeBuilders[len(typeBuilders)-1].fields = append(typeBuilders[len(typeBuilders)-1].fields, field)
		}

	}

	log.Println("TypeBuilders: ", typeBuilders)
	return typeBuilders, nil

}

func convertStructToTsType(structInput TypeBuilder) (string, error) {
	if structInput.name == "" {
		return "", errors.New("Struct name is empty")
	}

	if len(structInput.fields) == 0 {
		return "", errors.New("Struct has no fields")
	}

	var tsType string
	tsType = "export type " + structInput.name + " = {\n"

	for _, field := range structInput.fields {
		tsType += "\t" + strings.Trim(field.name, "\t") + ": " + field.typ + "\n"
	}
	tsType += "}\n"

	return tsType, nil
}

func removeEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
