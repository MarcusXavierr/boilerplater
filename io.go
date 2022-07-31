package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile(sentence, filepath string) error {
	file, e := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.FileMode(644))
	check(e)
	_, err := file.WriteString(sentence + "\n")
	check(err)
	return file.Close()
}
