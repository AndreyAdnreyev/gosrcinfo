package main

import (
	"os"
	"testing"
)

func setupTestGetGoFiles(tb testing.TB, path string) (func(tb testing.TB), error) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return nil, err
	}

	files := []string{path + "/test1.go", path + "/test2.go"}
	for _, f := range files {
		_, err := os.Create(f)
		if err != nil {
			return nil, err
		}
	}

	return func(tb testing.TB) {
		_ = os.RemoveAll(path)
	}, nil
}
func TestGetGoFiles(t *testing.T) {
	var tests = []struct {
		in   string
		want []string
	}{
		{"tests", []string{"tests/test1.go", "tests/test2.go"}},
	}

	for _, tt := range tests {
		testname := tt.in
		t.Run(testname, func(t *testing.T) {
			teardownTest, err := setupTestGetGoFiles(t, tt.in)
			if err != nil {
				t.Errorf("Failed to setup test")
			}
			defer teardownTest(t)

			files, err := getGoFiles(tt.in)
			if err != nil {
				t.Errorf("%v\n", err)
			}
			if len(files) != len(tt.want) {
				t.Errorf("got %d, want %d", len(files), len(tt.want))
			}
			for i := 0; i < len(files); i++ {
				if files[i] != tt.want[i] {
					t.Errorf("got %v, want %v", files, tt.want)
				}
			}
		})
	}
}
