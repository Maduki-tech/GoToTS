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


	readTheStruct(holder)
	return string(holder[:count])
}
