package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(string []string) ([]float64, error) {
	var floats []float64
	for _, stringValue := range string {
		floatVal, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}

		floats = append(floats, floatVal)

	}

	return floats, nil
}
