package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"os"
	"time"
)

var (
	year, month, day = time.Now().Date()
)

func main() {
	path, err := createPath()
	if err != nil {
		panic(err)
	}

	fileName, err := createFile(path)
	if err != nil {
		panic(err)
	}

	err = openFile(fileName)
	if err != nil {
		panic(err)
	}
}

func openFile(fileName string) (err error) {
	err = open.Run(fileName)

	return
}

func createFile(path string) (fileName string, err error) {
	fileName = fmt.Sprintf("%s/%d.md", path, day)
	templateContent, err := ioutil.ReadFile("templates/pt.md")
	if err != nil {
		panic(err)
	}

	fileTemplate :=
		fmt.Sprintf("# %d/%d/%d", year, month, day) + string(templateContent)

	if fileNotExists(fileName) {
		err = ioutil.WriteFile(fileName, []byte(fileTemplate), 0644)
	}

	return
}

func createPath() (path string, err error) {
	path = fmt.Sprintf("%d/%d", year, month)
	err = os.MkdirAll(path, 0777)

	return
}

func fileNotExists(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}
