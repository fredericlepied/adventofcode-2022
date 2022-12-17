// https://adventofcode.com/2022/day/16

package day16

import (
	"fmt"
	. "github.com/fredericlepied/adventofcode-2022/utils"
	re "regexp"
	"sort"
	"strings"
)

type Valve struct {
	name   string
	rate   int
	valves []string
	opened bool
}

func DFS(valves map[string]Valve, start string) (result []string) {
	var stack Stack[string]
	stack.Push(start)
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		valve := valves[v]
		if Index(v, result) == -1 {
			result = append(result, v)
			valves[v] = valve
			names := valves[v].valves
			sort.Slice(names, func(i, j int) bool {
				return valves[valves[v].valves[i]].rate < valves[valves[v].valves[j]].rate
			})
			fmt.Println(v, "=>", names)
			for i := range names {
				stack.Push(valves[v].valves[i])
			}
		}
	}
	return result
}

func FindPath(valves map[string]Valve, start string, all []string, visited []string) (result []string) {
	if len(all) == 0 {
		fmt.Println("found", visited)
		return visited
	}

	all_names := valves[start].valves
	sort.Slice(all_names, func(i, j int) bool {
		return valves[all_names[i]].rate > valves[all_names[j]].rate
	})

	for _, name := range all_names {
		fmt.Println("visiting", name, visited)
		path := FindPath(valves, name, Remove(name, all), append(visited, name))
		return path
	}
	fmt.Println("no nore moves")
	return visited
}

func ComputeResult(lines []string) (result int) {
	fmt.Println("ComputeResult")

	line_regexp := re.MustCompile("Valve (..) has flow rate=(\\d+); tunnels? leads? to valves? (.+)")
	valves := map[string]Valve{}
	for idx := 0; idx < len(lines); idx++ {
		line := lines[idx]
		res := line_regexp.FindStringSubmatch(line)
		if len(res) != 4 {
			fmt.Println("Invalid line", line)
		} else {
			valves[res[1]] = Valve{res[1], ToInt(res[2]), strings.Split(res[3], ", "), false}
		}
	}
	for name := range valves {
		fmt.Println(name, "=>", valves[name].valves)

	}
	ordered := DFS(valves, "AA")
	fmt.Println(ordered)
	moving := true
	all_names := GetKeys(valves)
	sort.Slice(all_names, func(i, j int) bool {
		return valves[all_names[i]].rate > valves[all_names[j]].rate
	})
	visited := []string{"AA"}
	path := FindPath(valves, "AA", all_names, visited)
	fmt.Println("path=", path)
	// 	fmt.Println("sorted names:", all_names)
	//opened := []string{}
	loop := 0
	for mn := 1; mn <= 30; mn++ {
		// compute pressure for this minute
		for idx := range valves {
			if valves[idx].opened {
				fmt.Println(mn, "valve", valves[idx].name, "is opened")
				result += valves[idx].rate
			}
		}
		if moving && loop < len(ordered) {
			moving = false
		} else {
			if loop < len(ordered) {
				valve := valves[ordered[loop]]
				fmt.Println(mn, "moving to valve", valve.name)
				if valve.rate != 0 {
					valve.opened = true
					valves[ordered[loop]] = valve
					moving = true
				}
				loop++
			}
		}
		// move
		// 		if moving {
		// 			sort.Slice(current.valves, func(i, j int) bool {
		// 				return valves[current.valves[i]].rate > valves[current.valves[j]].rate
		// 			})
		// 			moved := false
		// 			for name := range all_names {
		// 				for idx := range current.valves {
		// 					valve := valves[current.valves[idx]]
		// 					if idx == name && !valve.opened && valve.rate > 0 {
		// 						current = valve
		// 						fmt.Println(mn, "moving to valve", current.name)
		// 						moved = true
		// 						moving = false
		// 						break
		// 					}
		// 				}
		// 				if moved {
		// 					break
		// 				}
		// 			}
		// 			if !moved {
		// 				for idx := range current.valves {
		// 					valve := valves[current.valves[idx]]
		// 					if valve.rate == 0 && !valve.opened && Index(valve.name, queue) != -1 {
		// 						current = valve
		// 						current.opened = true
		// 						valves[current.name] = current
		// 						fmt.Println(mn, "moving to zero valve", current.name)
		// 						moved = true
		// 						queue = append(queue, current.name)
		// 						break
		// 					}
		// 				}
		// 			}
		// 			if !moved {
		// 				valve := valves[current.valves[len(current.valves)-1]]
		// 				current = valve
		// 				current.opened = true
		// 				valves[current.name] = current
		// 				fmt.Println(mn, "moving to zero- valve", current.name)
		// 				moved = true
		// 				queue = append(queue, current.name)
		// 			}
		// 		} else {
		// 			current.opened = true
		// 			valves[current.name] = current
		// 			fmt.Println(mn, "opening", current.name)
		// 			moving = true
		// 		}
	}
	fmt.Println(valves)
	return result
}
