package file_handlers

import (
	"bufio"
	"os"
	"path/filepath"
)

//Function reads from the given fileName and returns file content
func GetFile(fileName string) (*os.File,  error) {
	path := getInputFilePath(fileName)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetFileContentScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	return scanner
}

func GetFileExtension(fileName string) string {
	return filepath.Ext(fileName)
}