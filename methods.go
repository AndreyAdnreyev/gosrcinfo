package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func GetTypeMethods(files []string, typeName string) ([]string, error) {
	types := []string{}
	for _, f := range files {
		ftypes, err := readTypeMethodsInFile(f, typeName)
		if err != nil {
			return nil, err
		}
		types = append(types, ftypes...)
	}
	return types, nil
}

func readTypeMethodsInFile(fileName, typeName string) ([]string, error) {
	//Check pkg name in a file

	methods := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := `^func\s\(.\s.?` + typeName
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		matched, _ := regexp.MatchString(re, line)
		if matched {
			//words := strings.Split(line, " ")

			methods = append(methods, strings.ReplaceAll(line, "{", ""))
		}
	}
	return methods, nil
}
