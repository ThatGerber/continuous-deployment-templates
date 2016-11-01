package utils

import (
	"fmt"
	"os"
)

// Check checks for an error and exits if necessary
// func Check(err error) {
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

func LoadFiles(fileNames []string) []*os.File {
	files := []*os.File{}
	for _, fileName := range fileNames {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file: " + fileName)
			os.Exit(1)
		}
		files = append(files, f)
	}
	return files
}
