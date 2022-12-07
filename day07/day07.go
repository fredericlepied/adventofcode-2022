// https://adventofcode.com/2022/day/7

package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func addFile(size int, current_dir string, sizes map[string]int) {
	split := strings.Split(current_dir, "/")
	for idx := 1; idx <= len(split); idx++ {
		dir := strings.Join(split[0:idx], "/")
		_, ok := sizes[dir]
		if ok {
			sizes[dir] += size
		} else {
			sizes[dir] = size
		}
		//fmt.Println("=> size", dir, sizes[dir])
	}
}

func ComputeResult(lines []string, max int) (total int) {
	fmt.Println("ComputeResult")
	total = 0
	var sizes map[string]int
	sizes = make(map[string]int)
	current_dir := "/"

	for loop := 0; loop < len(lines); loop++ {
		line := lines[loop]
		//fmt.Println(line)
		split := strings.Split(line, " ")
		if split[0] == "$" {
			if split[1] == "cd" {
				if split[2] == "/" {
					current_dir = ""
				} else if split[2] == ".." {
					dirs := strings.Split(current_dir, "/")
					current_dir = strings.Join(dirs[:len(dirs) - 1], "/")
				} else {
					current_dir = current_dir + "/" + split[2]
				}
				//fmt.Println("=> current_dir", current_dir)
			}
		} else {
			if split[0] == "dir" {

			} else {
				size, _ := strconv.Atoi(split[0])
				addFile(size, current_dir, sizes)
			}
		}
	}
	for _, value := range sizes {
		//fmt.Println("Key:", key, "Value:", value)
		if value <= max {
			total += value
		}
	}
	return total
}
