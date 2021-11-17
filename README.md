**gosrcinfo** allows to collect information about sources of Go projects.
It displays Go files, packages, types and methods. **gosrcinfo** uses only standard library.

# Usage 

```
git clone git@github.com:AndreyAdnreyev/gosrcinfo.git
cd gosrsinfo
go build
./gosrcinfo
./gosrcinfo -help
```

# Command Examples

1. List all Go files in the current directory and sub directories
```
./gosrcinfo
```
2. List of all packages in the files of the current directory and sub directories
```
./gosrcinfo -lsP
```
3. List of all types in files of the specified directory and its sub directories
```
./gosrcinfo -path=/path/to/files -lsT
```
4. List of types in the package
```
./gosrcinfo -lsT -pkg=package_name
```
5. List of methods of specified type
```
./gosrcinfo -type=TestType -lsM
```

# Licensing

See [LICENSE](LICENSE) for more details.
