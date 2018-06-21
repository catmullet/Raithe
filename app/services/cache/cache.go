package cache

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Set Writes a message to file
func Set(key string, message []byte) error {
	return writeFile(key, message)
}

// Get Retrieves a message from file
func Get(queue string) ([]byte, error) {
	return readFile(queue)
}

func writeFile(key string, message []byte) error {
	createDirectory(key)
	rootPath := os.Getenv("ROOTPATH")
	return ioutil.WriteFile(rootPath+string(filepath.Separator)+key+string(filepath.Separator)+strconv.FormatInt(makeTimestamp(), 10), message, 0777)
}

func readFile(key string) ([]byte, error) {
	rootPath := os.Getenv("ROOTPATH")
	msg := []byte{}

	filepath.Walk(rootPath+string(filepath.Separator)+key+string(filepath.Separator), func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			m, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
			}
			os.Remove(path)
			msg = m
			return io.EOF
		}
		return nil
	})

	return msg, nil
}

func createDirectory(key string) {
	rootPath := os.Getenv("ROOTPATH")
	if _, err := os.Stat(rootPath + string(filepath.Separator) + key); os.IsNotExist(err) {
		os.MkdirAll(rootPath+string(filepath.Separator)+key, 0777)
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
