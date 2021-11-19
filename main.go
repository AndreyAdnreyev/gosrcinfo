package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	methodRe string = `^func\s\(.\s.?`
	fnRe     string = `^func\s\w`
	typeRe   string = `^type\s\w`
)

func main() {
	var path = flag.String("path", ".", "Path where to look at")
	var lsP = flag.Bool("lsP", false, "List all packages")
	var lsT = flag.Bool("lsT", false, "List all types")
	var lsM = flag.Bool("lsM", false, "List all methods of specified type")
	var lsF = flag.Bool("lsF", false, "List all functions")
	var pkg = flag.String("pkg", "", "Apply action to a specific package")
	var t = flag.String("type", "", "Apply action to a specific type")
	var help = flag.Bool("help", false, "Print help")
	flag.Parse()

	switch {
	case *help:
		printHelp()
	case !*lsP && !*lsT && !*lsM && !*lsF:
		fmt.Printf("The list of all Go files in the folder %s\n\n", *path)
		listAllFiles(*path)
	case *lsP && !*lsT && !*lsM && !*lsF:
		fmt.Printf("The list of all packages in all files\n\n")
		listData(*path, "", "pkg")
	case !*lsP && *lsT && !*lsM && !*lsF:
		fmt.Printf("The list of all types in all files\n\n")
		listData(*path, *pkg, typeRe)
	case !*lsP && !*lsT && *lsM && !*lsF:
		fmt.Printf("The list of all methods of type %s\n\n", *t)
		listData(*path, *pkg, methodRe+*t)
	case !*lsP && !*lsT && !*lsM && *lsF:
		fmt.Printf("The list of all functions \n\n")
		listData(*path, *pkg, fnRe+*t)
	default:
		printHelp()
	}
}

func listAllFiles(path string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	printSlice(files)
}

func listData(path, pkg, search string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v", err)
		os.Exit(1)
	}
	data, err := getData(files, pkg, search)
	if err != nil {
		fmt.Printf("Failed to get data: %v", err)
		os.Exit(1)
	}
	data.print()
}

// getData returns collected data from all Go files
func getData(files []string, pkg, search string) (MapData, error) {
	out := NewMapData()

	for _, f := range files {
		fout, err := readData(f, pkg, search)
		if err != nil {
			return nil, err
		}
		out.add(fout, f)

	}

	return out, nil

}

// readData reads a file and returns slice of matched lines
func readData(path, pkg, search string) ([]string, error) {
	out := []string{}
	file, err := os.Open(path)
	if err != nil {
		return []string{""}, err
	}
	defer file.Close()

	re := regexp.MustCompile(search)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "package ") && pkg == "" && search == "pkg" {
			words := strings.Split(line, " ")
			return []string{words[1]}, nil
		}

		if pkg != "" &&
			strings.HasPrefix(line, "package ") &&
			!strings.HasPrefix(line, "package "+pkg) {
			return []string{}, nil
		}

		matched := re.MatchString(line)
		if matched {
			out = append(out, strings.ReplaceAll(line, "{", ""))
		}

	}
	if err := scanner.Err(); err != nil {
		return out, err
	}
	return out, nil
}
