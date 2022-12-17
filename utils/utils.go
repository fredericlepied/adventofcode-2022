package utils

import (
	"bufio"
	"fmt"
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

func GetKeys[T comparable, TT any](m map[T]TT) (res []T) {
	for key := range m {
		res = append(res, key)
	}
	return res
}

func Index[T comparable](e T, l []T) (idx int) {
	for idx = range l {
		if e == l[idx] {
			return idx
		}
	}
	return -1
}

func Remove[T comparable](e T, l []T) (res []T) {
	i := Index(e, l)
	switch {
	case i == -1:
		return l
	case i == 0:
		return l[1:]
	case i == len(l)-1:
		return l[:len(l)-2]
	case true:
		return append(l[:i], l[i+1:]...)
	}
	return l
}

func Debug(print bool, a ...any) {
	fmt.Println(a...)
}
