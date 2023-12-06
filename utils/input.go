package utils

import (
	"os"
	"strings"
)

func ReadFileIntoSlice(fileName string) []string {

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")

	return input
}
