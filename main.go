package boilerplater

import (
	"os"
	"path/filepath"
)

func main() {

	relativePath := os.Args[1]
	absoluePath, err := filepath.Abs(relativePath)
	check(err)
	php := Php{Path: relativePath}
	WriteFile(php.CreateBoilerplateString(), absoluePath)
}
