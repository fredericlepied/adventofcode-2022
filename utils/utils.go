package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(filename string) (res []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
