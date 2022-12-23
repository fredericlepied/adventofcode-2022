// https://adventofcode.com/2022/day/21

package day21

import (
	"fmt"
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"regexp"
	"strings"
)

var debug = true
var operationRegexp = regexp.MustCompile("(.+) ([+-/*]) (.+)")
var numericRegexp = regexp.MustCompile(`^\d+$`)

func b(str string) []byte {
	return []byte(str)

}

func get(varName string, vars map[string]string) int {
	Debug(debug, "get:", varName, "=>", vars[varName])
	res := operationRegexp.FindStringSubmatch(vars[varName])

	if len(res) == 0 {
		return ToInt(vars[varName])
	} else if len(res) == 4 {
		switch res[2] {
		case "+":
			return get(res[1], vars) + get(res[3], vars)
		case "-":
			return get(res[1], vars) - get(res[3], vars)
		case "/":
			return get(res[1], vars) / get(res[3], vars)
		case "*":
			return get(res[1], vars) * get(res[3], vars)
		}
	} else {
		fmt.Println("Invalid expression", varName, vars[varName])
		panic("Not found")
	}
	return 0
}

func ComputeResult(lines []string) (result int) {
	fmt.Println("ComputeResult")

	name := map[string]string{}

	for idx := 0; idx < len(lines); idx++ {
		split := strings.Split(lines[idx], ": ")
		name[split[0]] = split[1]
	}
	return get("root", name)
}

func findHumn(varName string, vars map[string]string) bool {
	res := operationRegexp.FindStringSubmatch(vars[varName])

	if len(res) == 0 {
		return false
	} else if len(res) == 4 {
		if res[1] == "humn" || res[3] == "humn" {
			return true
		} else {
			return findHumn(res[1], vars) || findHumn(res[3], vars)
		}
	}
	fmt.Println("Invalid expression", varName, vars[varName])
	return false
}

func get2(varName string, vars map[string]string) int {
	Debug(debug, "get2:", varName, "=>", vars[varName])

	if numericRegexp.MatchString(vars[varName]) {
		return ToInt(vars[varName])
	}

	for key, val := range vars {
		res := operationRegexp.FindStringSubmatch(val)
		if len(res) == 4 {
			if res[1] == varName {
				switch res[2] {
				case "+":
					return get2(key, vars) - get2(res[3], vars)
				case "-":
					return get2(key, vars) + get2(res[3], vars)
				case "/":
					return get2(key, vars) * get2(res[3], vars)
				case "*":
					return get2(key, vars) / get2(res[3], vars)
				}
			} else if res[3] == varName {
				switch res[2] {
				case "+":
					return get2(res[1], vars) - get2(key, vars)
				case "-":
					return get2(res[1], vars) + get2(key, vars)
				case "/":
					return get2(res[1], vars) * get2(key, vars)
				case "*":
					return get2(res[1], vars) / get2(key, vars)
				}
			}
		}
	}
	// 	res := operationRegexp.FindStringSubmatch(vars[varName])
	//
	// 	if len(res) == 4 {
	// 		switch res[2] {
	// 		case "+":
	// 			return get2(res[1], vars) - get2(res[3], vars)
	// 		case "-":
	// 			return get2(res[1], vars) - get2(res[3], vars)
	// 		case "/":
	// 			return get2(res[1], vars) / get2(res[3], vars)
	// 		case "*":
	// 			return get2(res[1], vars) * get2(res[3], vars)
	// 		}
	// 	}

	fmt.Println("variable not found", varName)
	panic("Not found")
	return 0
}

func ComputeResult2(lines []string) (result int) {
	fmt.Println("ComputeResult2")

	in := map[string]string{}

	for idx := 0; idx < len(lines); idx++ {
		split := strings.Split(lines[idx], ": ")
		if split[0] != "humn" {
			in[split[0]] = split[1]
		}
	}
	parts := strings.Split(in["root"], " + ")
	fmt.Println(parts[0], "==", parts[1])
	fmt.Println()
	val := 0
	if findHumn(parts[0], in) {
		val = get(parts[1], in)
		fmt.Println("found left. right part is", val)
	} else {
		val = get(parts[0], in)
		fmt.Println("found right. left part is", val)
	}
	in["root"] = fmt.Sprint(val * 2)
	in[parts[0]] = fmt.Sprint(val)
	in[parts[1]] = fmt.Sprint(val)
	fmt.Println()
	for key, val := range in {
		fmt.Printf("%s: %s\n", key, val)
	}
	fmt.Println()
	return get2("humn", in)
}
