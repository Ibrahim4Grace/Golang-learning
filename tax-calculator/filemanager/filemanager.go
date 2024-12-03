package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// accepting file path as parameter
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {

		return nil, errors.New("failed to open file")
	}

	// caall file close as the right time
	defer file.Close()

	//bufio is use to read content from file
	scanner := bufio.NewScanner(file)

	var lines []string
	//4 loop to scan all files
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("failed to read line in file")
	}

	// file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {

	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}

	// caall file close as the right time
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to Json")
	}

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}

}
