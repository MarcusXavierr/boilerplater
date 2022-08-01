package main

import (
	"os"
	"path/filepath"
)

func main() {
	fileType := os.Args[1]
	relativePath := os.Args[2]
	absoluePath, err := filepath.Abs(relativePath)
	check(err)
	php := Php{Path: relativePath}
	WriteFile(php.CreateBoilerplate(fileType), absoluePath)
}
