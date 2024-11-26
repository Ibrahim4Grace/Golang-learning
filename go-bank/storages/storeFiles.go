package storage

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

//starting function character with uppercase means export 
//GetFloatFromfile


//to get data from file
func GetFloatFromfile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		// Return default balance and an appropriate error
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)
	value, parseErr := strconv.ParseFloat(valueText, 64)
	if parseErr != nil {
		// Return default balance and parsing error
		return 1000, errors.New("failed to parse store file")
	}

	return value, nil
}

//To store data in a file
func WriteFloatToFile(value float64, fileName string) error {
	valueText := fmt.Sprintf("%.2f", value) // Formatting to 2 decimal places
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		return errors.New("failed to write balance to file")
	}
	return nil
}
