package utility

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string, defaultVal float64) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return defaultVal, errors.New("Requested file not found. Please try again later.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultVal, errors.New("Invalid value in file. Please try again later.")
	}

	return value, nil
}

func WriteFloatToFile(filename string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}
