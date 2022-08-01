package main

import (
	"os"
	"strings"
)

type Php struct {
	Path string
}

func (p Php) ParsePath() (namespace string, className string) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	normalizedPath := strings.ReplaceAll(p.Path, currentDirectory+"/", "")
	path, file := separateFileFromPath(normalizedPath)
	namespace = createNamespace(path)
	className = removeFileExtension(file)
	return
}

func separateFileFromPath(fullPath string) ([]string, string) {
	pathSlice := strings.Split(fullPath, "/")
	path := pathSlice[:len(pathSlice)-1]
	filename := pathSlice[len(pathSlice)-1]
	return path, filename
}

func createNamespace(path []string) string {
	var aux []string
	for _, folder := range path {
		aux = append(aux, strings.Title(folder))
	}
	return strings.Join(aux, "\\")
}

func removeFileExtension(filename string) string {
	return strings.ReplaceAll(filename, ".php", "")
}

func (p Php) CreateBoilerplateClass() string {
	namespace, className := p.ParsePath()
	return "<?php\nnamespace " + namespace + ";\n\nclass " + className + "\n{\n}"
}

func (p Php) CreateBoilerplateInterface() string {
	namespace, className := p.ParsePath()
	return "<?php\nnamespace " + namespace + ";\n\ninterface " + className + "\n{\n}"
}

func (p Php) CreateBoilerplate(boilerplateType string) string {
	if boilerplateType == "c" {
		return p.CreateBoilerplateClass()
	}
	return p.CreateBoilerplateInterface()
}
