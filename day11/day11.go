// https://adventofcode.com/2022/day/11

package day11

import (
	"fmt"
	//re "regexp"
	"math"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operation string
	divisible int
	valid     int
	invalid   int
	inspected int
}

// ByAge implements sort.Interface based on the Age field.
type ByInspected []monkey

func (a ByInspected) Len() int           { return len(a) }
func (a ByInspected) Less(i, j int) bool { return a[i].inspected > a[j].inspected }
func (a ByInspected) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func ComputeResult(lines []string, level int, rounds int) (result int) {
	fmt.Println("ComputeResult")

	//number_regexp := re.MustCompile("[,-]")

	var monkeys = []monkey{}
	current := monkey{}

	for idx := 0; idx < len(lines); idx++ {
		line := lines[idx]
		if strings.HasPrefix(line, "Monkey ") {
			if current.divisible != 0 {
				monkeys = append(monkeys, current)
			}
			new := monkey{}
			current = new
			current.divisible = idx
		} else if strings.HasPrefix(line, "  Starting items: ") {
			for _, item := range strings.Split(strings.TrimPrefix(line, "  Starting items: "), ", ") {
				n, _ := strconv.Atoi(item)
				current.items = append(current.items, n)
			}
		} else if strings.HasPrefix(line, "  Operation: new = ") {
			current.operation = strings.TrimPrefix(line, "  Operation: new = ")
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			op, _ := strconv.Atoi(strings.TrimPrefix(line, "  Test: divisible by "))
			current.divisible = op
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			op, _ := strconv.Atoi(strings.TrimPrefix(line, "    If true: throw to monkey "))
			current.valid = op
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			op, _ := strconv.Atoi(strings.TrimPrefix(line, "    If false: throw to monkey "))
			current.invalid = op
		}
		//res := order_regexp.FindStringSubmatch(line)
	}
	monkeys = append(monkeys, current)
	worryMod := 1
	for id := range monkeys {
		worryMod *= monkeys[id].divisible
	}

	for round := 1; round <= rounds; round++ {
		for id := range monkeys {
			for _, item := range monkeys[id].items {
				result := executeOperation(monkeys[id].operation, item, worryMod, level)
				monkeys[id].inspected += 1
				var idx int
				if result%monkeys[id].divisible == 0 {
					idx = monkeys[id].valid
				} else {
					idx = monkeys[id].invalid
				}
				monkeys[idx].items = append(monkeys[idx].items, result)
			}
			monkeys[id].items = nil
		}
		if round == 20 || round == 1 || round%1000 == 0 {
			fmt.Println("Round", round)
			for id := range monkeys {
				fmt.Printf("Monkey %d inspected items %d times.\n", id, monkeys[id].inspected)
			}
		}
	}
	sort.Sort(ByInspected(monkeys))
	return monkeys[0].inspected * monkeys[1].inspected
}

func executeOperation(op string, val int, mod int, level int) (result int) {
	part := strings.Split(op, " ")
	var x int
	var y int

	if part[0] == "old" {
		x = val
	} else {
		r, _ := strconv.Atoi(part[0])
		x = r
	}

	if part[2] == "old" {
		y = val
	} else {
		r, _ := strconv.Atoi(part[2])
		y = r
	}

	if part[1] == "*" {
		result = x * y
	}
	if part[1] == "+" {
		result = x + y
	}
	result = int(math.Round(float64(result / level)))

	result %= mod

	return result
}
