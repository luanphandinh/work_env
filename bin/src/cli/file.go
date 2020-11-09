package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func dirExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func getDirPath(dir string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	var path string
	if dir != "" {
		path = fmt.Sprintf("%s/cli_beta/%s", home, dir)
	} else {
		path = fmt.Sprintf("%s/cli_beta", home)
	}

	if dirExists(path) {
		return path
	}

	err = os.MkdirAll(path, 0777)
	if err != nil {
		panic(err)
	}

	return path
}

func getFilePath(dir string, file string) string {
	path := fmt.Sprintf("%s/%s", getDirPath(dir), file)
	if fileExists(path) {
		return path
	}

	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return path
}

func getFileContent(dir string, file string) ([]byte, error) {
	path := getFilePath(dir, file)
	if fileExists(path) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, nil
}
