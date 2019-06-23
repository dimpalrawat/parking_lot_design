package file_handlers

import (
	"path/filepath"
)

//Function returns the absolute path of the inputFile
func getInputFilePath(inputFileName string) string {
	thepath, err := filepath.Abs(filepath.Dir(inputFileName))
	thepath = thepath + INPUT_FILE_PATH + inputFileName
	if err != nil {
		panic(err)
	}
	return thepath

}
