package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileExists return true if given file name exist.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// DirExists return true if give Dir is exist.
func DirExists(dir string) bool {
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

// GetDirPath return file path corresponding to user HomeDir
// and will create new file if there is none.
func GetDirPath(dir string) string {
	home, err := os.UserHomeDir()
	check(err)

	var path string
	if dir != "" {
		path = fmt.Sprintf("%s/cli_beta/%s", home, dir)
	} else {
		path = fmt.Sprintf("%s/cli_beta", home)
	}

	if DirExists(path) {
		return path
	}

	err = os.MkdirAll(path, 0777)
	check(err)

	return path
}

// GetFilePath base on directory
// If directory path doesn't exist, create new one.
// If file doesn't exist, create new one
func GetFilePath(dir string, file string) string {
	path := fmt.Sprintf("%s/%s", GetDirPath(dir), file)
	if FileExists(path) {
		return path
	}

	_, err := os.Create(path)
	check(err)

	return path
}

// GetFileContent read and return content of given file.
// If any error, will return in err.
func GetFileContent(dir string, file string) ([]byte, error) {
	path := GetFilePath(dir, file)
	if FileExists(path) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
