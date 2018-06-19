package store

import (
	"io/ioutil"
	"os"
	"time"
	"fmt"
	"path/filepath"
	"strconv"
)

func Set(key string, message []byte) error {
	return writeFile(key, message)
}

func Get(queue string) ([]byte, error) {
	return readFile(queue)
}

func writeFile(key string, message []byte) error {
	createDirectory(key)
	rootPath := os.Getenv("ROOTPATH")
	return ioutil.WriteFile(rootPath + string(filepath.Separator) + key + string(filepath.Separator) + strconv.FormatInt(makeTimestamp(), 10), message,0777)
}

func readFile(key string) ([]byte, error) {
	rootPath := os.Getenv("ROOTPATH")
	f, err := os.Open(rootPath + string(filepath.Separator) + key + string(filepath.Separator))
	if err != nil {
		fmt.Println(err)
	}
	files, err := f.Readdir(-1)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	if len(files) > 0 {
		msg, err := ioutil.ReadFile(rootPath + string(filepath.Separator) + key + string(filepath.Separator) + files[0].Name())

		if err != nil {
			fmt.Println(err)
		}

		os.Remove(rootPath + string(filepath.Separator) + key + string(filepath.Separator) + files[0].Name())

		return msg, err
	}

	return []byte(""), err
}

func createDirectory(key string) {
	rootPath := os.Getenv("ROOTPATH")
	if _, err := os.Stat(rootPath + string(filepath.Separator) + key); os.IsNotExist(err) {
		os.MkdirAll(rootPath + string(filepath.Separator) + key, 0777)
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}