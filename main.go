package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path = flag.String("path", ".", "Path where to look at")
	var lsPkg = flag.Bool("lsPkg", false, "List all packages")
	var lsTypes = flag.Bool("lsTypes", false, "List all types")
	var pkgName = flag.String("pkgName", "foo", "Apply action to a specific package")
	flag.Parse()

	switch {
	case !*lsPkg && !*lsTypes:
		fmt.Println("The list of all Go files in the folder ", *path)
		listAllFiles(*path)
	case *lsPkg && !*lsTypes:
		fmt.Println("The list of all packages")
		listAllPkgs(*path)
	case !*lsPkg && *lsTypes && *pkgName == "foo":
		fmt.Println("The list of all types in files")
		listAllTypes(*path)
	case *pkgName != "foo" && *lsTypes:
		fmt.Println("The list of types in the package", *pkgName)
		listTypesInPkg(*path, *pkgName)
	}
}

func listAllFiles(path string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v", err)
		os.Exit(1)
	}
	printSlice(files)
}

func listAllPkgs(path string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v", err)
		os.Exit(1)
	}
	packages, err := getPkgs(files)
	if err != nil {
		fmt.Printf("Failed to get the list of packages: %v", err)
		os.Exit(1)
	}
	printSlice(packages)
}

func listAllTypes(path string) {
	// TODO: Handle the error
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v", err)
		os.Exit(1)
	}
	types, err := getTypes(files)
	if err != nil {
		fmt.Printf("Failed to get the list of types: %v", err)
		os.Exit(1)
	}
	printSlice(types)
}

func listTypesInPkg(path, pkgName string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v", err)
		os.Exit(1)
	}
	types, err := GetTypesOfPkg(files, pkgName)
	if err != nil {
		fmt.Printf("Failed to get the list of types: %v", err)
		os.Exit(1)
	}
	printSlice(types)
}

func printSlice(input []string) {
	for i, v := range input {
		fmt.Println(i, ":", v)
	}
}
