package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// methodsOfType returns the slice of all methods of a type in all packages
func methodsOfType(files []string, typeName string) ([]string, error) {
	types := []string{}
	for _, f := range files {
		ftypes, err := readMethodsOfType(f, typeName)
		if err != nil {
			return nil, err
		}
		types = append(types, ftypes...)
	}
	return types, nil
}

// readMethodsOfType scans a file and return a slice of methods of a type in all packages
func readMethodsOfType(fileName, typeName string) ([]string, error) {
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
			methods = append(methods, strings.ReplaceAll(line, "{", ""))
		}
	}
	return methods, nil
}

// methodsOfTypeInPkg returns the slice of all methods of a type of a specified package
func methodsOfTypeInPkg(files []string, tName, pkg string) ([]string, error) {
	types := []string{}
	for _, f := range files {
		ftypes, err := readMethodsOfTypeInPkg(f, tName, pkg)
		if err != nil {
			return nil, err
		}
		types = append(types, ftypes...)
	}
	return types, nil
}

// readMethodsOfTypeInPkg scans a file and return a slice of methods of a type of a specified package
func readMethodsOfTypeInPkg(fileName, tName, pkg string) ([]string, error) {
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
		if strings.HasPrefix(line, "package ") && !strings.HasPrefix(line, "package "+pkg) {
			return []string{}, nil
		}
		matched, _ := regexp.MatchString(re, line)
		if matched {
			methods = append(methods, strings.ReplaceAll(line, "{", ""))
		}
	}
	return methods, nil
}
