package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// methodsOfType returns the slice of all methods of a type
func methodsOfType(files []string, tName, pkg string) ([]string, error) {
	types := []string{}
	for _, f := range files {
		ftypes, err := readMethodsOfType(f, tName, pkg)
		if err != nil {
			return nil, err
		}
		types = append(types, ftypes...)
	}
	return types, nil
}

// readMethodsOfType scans a file and return a slice of methods of a type
func readMethodsOfType(fileName, tName, pkg string) ([]string, error) {
	methods := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := `^func\s\(.\s.?` + tName
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if pkg != "" &&
			strings.HasPrefix(line, "package ") &&
			!strings.HasPrefix(line, "package "+pkg) {
			return []string{}, nil
		}
		matched, _ := regexp.MatchString(re, line)
		if matched {
			methods = append(methods, strings.ReplaceAll(line, "{", ""))
		}
	}
	return methods, nil
}
