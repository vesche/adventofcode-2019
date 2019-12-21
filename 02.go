package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func runComputer(op1 int, op2 int, int_list []int) int {
	int_list[1] = op1
	int_list[2] = op2
	for i, n := range int_list {
		if i%4 == 0 {
			if n == 1 {
				int_list[int_list[i+3]] = int_list[int_list[i+1]] + int_list[int_list[i+2]]
			} else if n == 2 {
				int_list[int_list[i+3]] = int_list[int_list[i+1]] * int_list[int_list[i+2]]
			} else if n == 99 {
				break
			}
		}
	}
	return int_list[0]
}

func cloneIntList(int_list []int) []int {
	int_list_clone := make([]int, len(int_list))
	copy(int_list_clone, int_list)
	return int_list_clone
}

func partTwo(int_list []int) (int, int) {
	var a, b int
	for a = 0; a <= 99; a++ {
		for b = 0; b <= 99; b++ {
			part2_int_list := cloneIntList(int_list)
			part2 := runComputer(a, b, part2_int_list)
			if part2 == 19690720 {
				return a, b
			}
		}
	}
	return 0, 0
}

func main() {
	data, _ := ioutil.ReadFile("02_input.txt")
	str_list := strings.Split(string(data), ",")
	int_list := []int{}
	for _, s := range str_list {
		x, _ := strconv.Atoi(s)
		int_list = append(int_list, x)
	}

	part1_int_list := cloneIntList(int_list)
	part1 := runComputer(12, 2, part1_int_list)

	a, b := partTwo(int_list)
	part2 := 100*a + b

	fmt.Printf("part1: %d\npart2: %d\n", part1, part2)
}
