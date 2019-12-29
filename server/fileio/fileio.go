package fileio

import (
	"fmt"
	"io/ioutil"
)

// ReadFile helps to read file contents
func ReadFile(fp string) string {
	content, fioErr := ioutil.ReadFile(fp)
	if fioErr != nil {
		fmt.Println("Error reading file ", fp)
		return ""
	}

	return string(content)
}

// WriteFile is a helper to write content in given file
func WriteFile(fp string, content []byte) {
	err := ioutil.WriteFile(fp, content, 0644)
	if err != nil {
		fmt.Println("Error writing file ", fp, " \nError: ", err)
	}
}
