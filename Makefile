build:
	staticcheck .
	go test
	go build

clean:
	rm -rf tests
	rm -f gosrcinfo