// https://adventofcode.com/2022/day/25

package day25

import (
	"fmt"
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"math"
	"regexp"
)

var debug = false
var operationRegexp = regexp.MustCompile("(.+) ([+-/*]) (.+)")
var numericRegexp = regexp.MustCompile(`^\d+$`)

func b(str string) []byte {
	return []byte(str)

}

func SnafuToInt(snafu string) (res int) {
	digit := 1
	for idx := len(snafu) - 1; idx >= 0; idx-- {
		switch {
		case snafu[idx:idx+1] == "0":
		case snafu[idx:idx+1] == "1":
			res += digit
		case snafu[idx:idx+1] == "2":
			res += digit * 2
		case snafu[idx:idx+1] == "-":
			res -= digit
		case snafu[idx:idx+1] == "=":
			res -= digit * 2
		}
		digit *= 5
	}
	return res
}

func IntToSnafu(numbe int) (res string) {
	number := numbe
	Debug(debug, "IntToSnafu", number)
	digit := int(math.Log(float64(number))/math.Log(5)) + 1
	list := []int{}
	power := int(math.Pow(5, float64(digit)))
	num := number / power

	Debug(debug, num, power)
	if num > 0 {
		number -= num * power
		list = append(list, num)
	}

	digit--

	for ; digit >= 0; digit-- {
		power := int(math.Pow(5, float64(digit)))
		num := number / power
		Debug(debug, "IntToSnafu", power, num)
		number -= num * power
		list = append(list, num)
	}

	res = ""
	remain := 0
	for idx := len(list) - 1; idx >= 0; idx-- {
		val := list[idx] + remain
		if debug {
			fmt.Print("remain ", remain, " val ", val)
		}
		remain = 0
		if val >= 5 {
			remain = 1
			val = val - 5
		}
		if val <= 2 {
			res = fmt.Sprint(val) + res
		} else if val == 3 {
			remain += 1
			res = "=" + res
		} else if val == 4 {
			remain += 1
			res = "-" + res
		}
		if debug {
			fmt.Println(" res", res)
		}
	}

	if remain >= 1 {
		res = fmt.Sprint(remain) + res
	}

	Debug(debug, "res", list, res, SnafuToInt(res), remain)
	return res
}

func ComputeResult(lines []string) (result string) {
	fmt.Println("ComputeResult")

	sum := 0

	for idx := 0; idx < len(lines); idx++ {
		sum += SnafuToInt(lines[idx])
	}
	return IntToSnafu(sum)
}
