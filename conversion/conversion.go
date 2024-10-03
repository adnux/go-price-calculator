package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, line := range strings {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("converting strings to floats failed")
		}

		floats = append(floats, floatPrice)
		// prices[lineIndex] = floatPrice
	}
	return floats, nil
}
