package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path = flag.String("path", ".", "Path where to look at")
	var lsP = flag.Bool("lsP", false, "List all packages")
	var lsT = flag.Bool("lsT", false, "List all types")
	var pkg = flag.String("pkg", "", "Apply action to a specific package")
	var tName = flag.String("tName", "", "Apply action to a specific type")
	var lsM = flag.Bool("lsM", false, "List all methods of specified type")
	var help = flag.Bool("help", false, "Print help")
	flag.Parse()

	switch {
	case !*lsP && !*lsT && !*help && *tName == "":
		fmt.Println("The list of all Go files in the folder ", *path)
		listAllFiles(*path)
	case *lsP && !*lsT:
		fmt.Println("The list of all packages")
		listAllPkgs(*path)
	case !*lsP && *lsT && *pkg == "":
		fmt.Println("The list of all types in files")
		listAllTypes(*path)
	case *pkg != "" && *lsT:
		fmt.Println("The list of types in the package", *pkg)
		listTypesInPkg(*path, *pkg)
	case *tName != "" && *lsM && *pkg == "":
		fmt.Println("The list of methods of type ", *tName)
		listMethodsOfType(*path, *tName)
	case *help:
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
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	types, err := getTypes(files)
	if err != nil {
		fmt.Printf("Failed to get the list of types: %v\n", err)
		os.Exit(1)
	}
	printSlice(types)
}

func listTypesInPkg(path, pkgName string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	types, err := GetTypesOfPkg(files, pkgName)
	if err != nil {
		fmt.Printf("Failed to get the list of types: %v\n", err)
		os.Exit(1)
	}
	printSlice(types)
}

func listMethodsOfType(path, typeName string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	methods, err := GetTypeMethods(files, typeName)
	if err != nil {
		fmt.Printf("Failed to get the list of methods of type %s: %v\n", typeName, err)
		os.Exit(1)
	}
	printSlice(methods)

}

func printHelp() {
	flag.PrintDefaults()

}

func printSlice(input []string) {
	for i, v := range input {
		fmt.Println(i+1, ":", v)
	}
}
