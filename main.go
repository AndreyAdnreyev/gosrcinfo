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
	var t = flag.String("type", "", "Apply action to a specific type")
	var lsM = flag.Bool("lsM", false, "List all methods of specified type")
	var help = flag.Bool("help", false, "Print help")
	flag.Parse()

	switch {
	case !*lsP && !*lsT && !*lsM && !*help && *pkg == "" && *t == "":
		fmt.Printf("The list of all Go files in the folder %s\n\n", *path)
		listAllFiles(*path)
	case *lsP && !*lsT && !*lsM && !*help && *pkg == "" && *t == "":
		fmt.Printf("The list of all packages in all files\n\n")
		listAllPkgs(*path)
	case !*lsP && *lsT && !*lsM && !*help && *pkg == "" && *t == "":
		fmt.Printf("The list of all types in all files\n\n")
		listAllTypes(*path)
	case !*lsP && *lsT && !*lsM && !*help && *pkg != "" && *t == "":
		fmt.Printf("The list of all types in the package %s\n\n", *pkg)
		listTypesInPkg(*path, *pkg)
	case !*lsP && !*lsT && *lsM && !*help && *pkg == "" && *t != "":
		fmt.Printf("The list of all methods of type %s\n\n", *t)
		listMethodsOfType(*path, *t)
	case !*lsP && !*lsT && *lsM && !*help && *pkg != "" && *t != "":
		fmt.Printf("The list of all methods of type %s in package %s\n\n", *t, *pkg)
		listMethodsOfTypeInPkg(*path, *t, *pkg)
	case *help:
		printHelp()
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

func listAllPkgs(path string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v", err)
		os.Exit(1)
	}
	pkgsData, err := getPkgs(files)
	if err != nil {
		fmt.Printf("Failed to get the list of packages: %v", err)
		os.Exit(1)
	}
	pkgsData.print()

}

func listAllTypes(path string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	typesData, err := types(files)
	if err != nil {
		fmt.Printf("Failed to get the list of types: %v\n", err)
		os.Exit(1)
	}
	typesData.print()
}

func listTypesInPkg(path, pkg string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	typesData, err := typesOfPkg(files, pkg)
	if err != nil {
		fmt.Printf("Failed to get the list of types: %v\n", err)
		os.Exit(1)
	}
	typesData.print()
}

func listMethodsOfType(path, t string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	methods, err := methodsOfType(files, t)
	if err != nil {
		fmt.Printf("Failed to get the list of methods of type %s: %v\n", t, err)
		os.Exit(1)
	}
	printSlice(methods)
}

func listMethodsOfTypeInPkg(path, t, pkg string) {
	files, err := getGoFiles(path)
	if err != nil {
		fmt.Printf("Failed to get the list of files: %v\n", err)
		os.Exit(1)
	}
	methods, err := methodsOfTypeInPkg(files, t, pkg)
	if err != nil {
		fmt.Printf("Failed to get the list of methods of type %s: %v\n", t, err)
		os.Exit(1)
	}
	printSlice(methods)
}
