**gosrcinfo** allows to collect information about the source code of Go projects.
It prints to stdout the lists of Go files, packages, types and methods. **gosrcinfo** uses only standard library.

# Build

```
git clone git@github.com:AndreyAdnreyev/gosrcinfo.git
cd gosrsinfo
go build
./gosrcinfo
./gosrcinfo -help
```

# Usage

1. List all Go files in the current directory and sub directories
```
./gosrcinfo
```
2. List all packages in the files of the current directory and sub directories
```
./gosrcinfo -lsP
```
3. List all types in files of the specified directory and its sub directories
```
./gosrcinfo -path=. -lsT
The list of all types in all files

  1 : type MapData map[string][]string : ./data.go
```
4. List all types in the package
```
./gosrcinfo -lsT -pkg=package_name
```
5. List all methods of a specified type
```
./gosrcinfo -type=MapData -lsM
The list of all methods of type MapData

  1 : func (d MapData) print()  : ./data.go
  2 : func (d MapData) add(in []string, file string)  : ./data.go
```
6. List all functions
```
./gosrcinfo -lsF
  1 : func getGoFiles(path string) ([]string, error)  : ./files.go
  2 : func TestGetGoFiles(t *testing.T)  : ./files_test.go
  3 : func getData(files []string, pkg, search string) (MapData, error)  : ./main.go
  4 : func readData(path, pkg, search string) ([]string, error)  : ./main.go
  5 : func printHelp()  : ./data.go
  6 : func printSlice(input []string)  : ./data.go
  7 : func setupTestGetGoFiles(tb testing.TB, path string) (func(tb testing.TB), error)  : ./files_test.go
  8 : func main()  : ./main.go
  9 : func listAllFiles(path string)  : ./main.go
 10 : func listData(path, pkg, search string)  : ./main.go
 11 : func NewMapData() MapData  : ./data.go
```
7. List all functions in a package
```
./gosrcinfo -lsF -pkg=package_name
```


# Licensing

See [LICENSE](LICENSE) for more details.
