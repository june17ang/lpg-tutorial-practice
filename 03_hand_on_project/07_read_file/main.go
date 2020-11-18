package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type fileWriter struct{}

func main() {
	absPath, _ := filepath.Abs("./07_read_write_file/testing.txt")
	file, err := os.Open(absPath)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fw := fileWriter{}
	io.Copy(fw, file)
}

func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
