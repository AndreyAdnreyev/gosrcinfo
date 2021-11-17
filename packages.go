package main

import (
	"bufio"
	"os"
	"strings"
)

// getPkgs returns a slice of packages
func getPkgs(files []string) (MapData, error) {
	pkgsData := NewMapData()
	for _, f := range files {
		pkgName, err := readPkg(f)
		if err != nil {
			return nil, err
		}
		pkgsData.add(pkgName, f)

	}
	return pkgsData, nil

}

// readPkg reads file until gets the line starting with "package"
func readPkg(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "package ") {
			words := strings.Split(line, " ")
			return words[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}
