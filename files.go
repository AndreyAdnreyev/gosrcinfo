package main

import (
	"fmt"
	"os"
	"strings"
)

//getGoFiles reads files in a dir and the subdirs
func getGoFiles(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	goFiles := []string{}
	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", path, file.Name())
		if file.IsDir() {
			subDirGoFiles, _ := getGoFiles(filePath)
			goFiles = append(goFiles, subDirGoFiles...)
		}

		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}
		goFiles = append(goFiles, filePath)
	}
	return goFiles, nil
}
