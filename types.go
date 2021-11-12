package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getTypes(files []string) ([]string, error) {
	types := []string{}
	for _, f := range files {
		ftypes, err := readTypesInFile(f)
		if err != nil {
			return nil, err
		}
		types = append(types, ftypes...)
	}
	return types, nil
}

func readTypesInFile(fileName string) ([]string, error) {
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

func GetTypesOfPkg(files []string, pkg string) ([]string, error) {
	types := []string{}
	for _, f := range files {
		ftypes, err := readPkgsTypesInFile(f, pkg)
		if err != nil {
			return nil, err
		}
		types = append(types, ftypes...)
	}
	return types, nil
}

func readPkgsTypesInFile(fileName, pkg string) ([]string, error) {
	//Check pkg name in a file
	packageName, err := readPkgName(fileName)
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
