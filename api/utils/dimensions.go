package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseDimensions(dimensions string) (width int, height int, err error) {
	parts := strings.Split(dimensions, "x")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid dimensions format")
	}

	width, err = strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid width")
	}

	height, err = strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid height")
	}

	return width, height, nil
}
