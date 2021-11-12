package main

import (
	"bufio"
	"os"
	"strings"
)

//getPkgs returns a slice of packages
func getPkgs(files []string) ([]string, error) {
	pkgs := make(map[string]struct{})
	for _, f := range files {
		packageName, err := readPkgName(f)
		if err != nil {
			return nil, err
		}
		if _, ok := pkgs[packageName]; !ok {
			pkgs[packageName] = struct{}{}
		}

	}
	keys := []string{}
	for k := range pkgs {
		keys = append(keys, k)
	}
	return keys, nil

}

//readPkgName reads the first line of Go files where package name present
func readPkgName(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.ReplaceAll(line, "package", "")
	return strings.TrimSpace(line), nil
}
