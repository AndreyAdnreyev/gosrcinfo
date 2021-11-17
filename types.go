package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// types collects all types in all files
func types(files []string) (MapData, error) {
	typesData := NewMapData()
	for _, f := range files {
		types, err := readTypes(f)
		if err != nil {
			return nil, err
		}
		for _, t := range types {
			typesData.add(t, f)
		}
	}
	return typesData, nil
}

// readTypes scans all types in a file
func readTypes(fileName string) ([]string, error) {
	types := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "type ") {
			words := strings.Split(line, " ")

			types = append(types, fmt.Sprintf("%s %s", words[1], words[2]))
		}
	}
	return types, nil
}

// typesOfPkg collects all types in all files of a specific package
func typesOfPkg(files []string, pkg string) (MapData, error) {
	typesData := NewMapData()
	for _, f := range files {
		types, err := readTypesOfPkg(f, pkg)
		if err != nil {
			return nil, err
		}
		for _, t := range types {
			typesData.add(t, f)
		}
	}
	return typesData, nil
}

// readTypesOfPkg scans all types in a file of a specific package
func readTypesOfPkg(fileName, pkg string) ([]string, error) {
	packageName, err := readPkg(fileName)
	if err != nil || packageName != pkg {
		return []string{}, err
	}

	types := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "type ") {
			words := strings.Split(line, " ")

			types = append(types, fmt.Sprintf("%s %s", words[1], words[2]))
		}
	}
	return types, nil
}
