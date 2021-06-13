package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/skratchdot/open-golang/open"
)

var (
	year, month, day = time.Now().Date()
)

func main() {
	path, err := createPath()
	checkError(err)

	fileName, err := createFile(path)
	checkError(err)

	err = openFile(fileName)
	checkError(err)
}

func createPath() (string, error) {
	path := fmt.Sprintf("%d/%d", year, month)

	return path, os.MkdirAll(path, 0777)
}

func createFile(path string) (string, error) {
	fileName := fmt.Sprintf("%s/%d.md", path, day)
	templateContent, err := ioutil.ReadFile("templates/pt.md")
	checkError(err)

	fileTemplate :=
		fmt.Sprintf("# %d/%d/%d", year, month, day) + string(templateContent)

	if fileNotExists(fileName) {
		err = ioutil.WriteFile(fileName, []byte(fileTemplate), 0644)
	}

	return fileName, err
}

func openFile(fileName string) error {
	return open.Run(fileName)
}

func fileNotExists(filename string) bool {
	_, err := os.Stat(filename)

	return os.IsNotExist(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
