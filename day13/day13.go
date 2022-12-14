// https://adventofcode.com/2022/day/13

package day13

import (
	"fmt"
	"strconv"
	"strings"
)

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func extractList(parts []string) (list string, rest string) {
	fmt.Println("extractList", parts)
	full := strings.Join(parts, ",")
	if len(full) == 0 {
		return "", ""
	}
	if full[0] != '[' {
		fmt.Println("not a list", full, "+", full)
		return full, full
	}
	count := 0
	for x := range full {
		if full[x] == '[' {
			count += 1
		}
		if full[x] == ']' {
			count -= 1
			if count == 0 {
				if x+1 >= len(full) {
					fmt.Println("literal", full)
					return full, ""
				}
				fmt.Println("extracted", full[:x+1], "+", full[min(len(full)-1, x+2):])
				return full[:x+1], full[x+2:]
			}
		}
	}
	fmt.Println("default", strings.Join(parts, ","))
	return strings.Join(parts, ","), ""
}

func compare2(left string, right string) bool {
	fmt.Println("compare", left, "with", right)
	var leftParts = []string{}
	var rightParts = []string{}
	if len(left) == 0 {
		return false
	}
	if left[0] == '[' {
		if left[1] != ']' {
			leftParts = strings.Split(left[1:len(left)-1], ",")
		}
	} else {
		leftParts = strings.Split(left, ",")
	}
	if len(right) == 0 {
		return true
	}
	if right[0] == '[' {
		if right[1] != ']' {
			rightParts = strings.Split(right[1:len(right)-1], ",")
		}
	} else {
		if left[0] == '[' {
			return compare(left, "["+right+"]")
		} else {
			rightParts = strings.Split(right, ",")
		}
	}
	for idx := 0; idx < len(leftParts); idx++ {
		if strings.HasPrefix(leftParts[idx], "[") {
			listL, restL := extractList(leftParts[idx:])
			listR, restR := extractList(rightParts[idx:])
			return compare(listL, listR) && compare(restL, restR)
		} else {
			if idx >= len(rightParts) {
				return false
			}
			fmt.Println("numbers", idx, leftParts[idx], rightParts[idx])
			l, _ := strconv.Atoi(leftParts[idx])
			r, _ := strconv.Atoi(rightParts[idx])
			if l > r {
				return false
			}
		}
	}
	fmt.Println("true", leftParts, rightParts)
	return true
}

func GetFirst(str string) (first string, rest string) {
	if str[0] != '[' || str[1] == ']' {
		return str, ""
	}
	if str[1] == '[' {
		count := 0
		for idx := 1; idx < len(str)-1; idx++ {
			if str[idx] == '[' {
				count += 1
			}
			if str[idx] == ']' {
				count -= 1
				if count == 0 {
					if str[idx+2:] == "" {
						return str[1 : idx+1], "[]"
					} else {
						return str[1 : idx+1], "[" + str[idx+2:]
					}
				}
			}
		}
	}
	parts := strings.Split(str[1:len(str)-1], ",")
	return parts[0], "[" + strings.Join(parts[1:], ",") + "]"
}

func isList(str string) bool {
	return len(str) > 0 && str[0] == '['
}

func compare(left string, right string) bool {
	fmt.Println("compare", left, "with", right)
	var carL string
	var restL string
	if isList(left) {
		carL, restL = GetFirst(left)
		if carL == "[]" {
			fmt.Println("empty left")
			return true
		}
		if !isList(right) {
			fmt.Println("not list right", right)
			return compare(carL, right) && compare(restL, right)
		}
		carR, restR := GetFirst(right)
		fmt.Println("split R", carR, restR)
		if carR == "[]" {
			fmt.Println("empty right")
			return false
		}
		return compare(carL, carR) && compare(restL, restR)
	} else {
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)
		fmt.Println("numbers", l, r, (r >= l))
		return (r >= l)
	}
	return true
}

func ComputeResult(lines []string) (number int) {
	fmt.Println("ComputeResult")
	for idx := 0; idx <= len(lines)/3; idx++ {
		left := lines[3*idx]
		right := lines[3*idx+1]
		fmt.Println(idx+1, left, right)
		if compare(left, right) {
			fmt.Println("right order", idx+1)
			number += idx + 1
		}
		fmt.Println()
	}
	fmt.Println("=>", number)
	fmt.Println()
	return number
}
