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
./gosrcinfo -path=/path/to/files -lsT
```
4. List all types in the package
```
./gosrcinfo -lsT -pkg=package_name
```
5. List all methods of a specified type
```
./gosrcinfo -type=TestType -lsM
```
6. List all functions
```
./gosrcinfo -lsF
```
7. List all functions in a package
```
./gosrcinfo -lsF -pkg=package_name
```


# Licensing

See [LICENSE](LICENSE) for more details.
