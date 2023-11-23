package utils

import (
	"errors"
	"strconv"
	"strings"
)

func GetPathParam(path string) (int, error) {
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 {
		return 0, errors.New("Invalid path")
	}
	paramStr := path[lastSlash+1:]
	if paramStr == "" {
		return 0, errors.New("Param not found")
	}
	param, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, errors.New("Param is not a valid integer")
	}
	return param, nil
}
