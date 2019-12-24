package main

import (
	"fmt"
	"strconv"
	"strings"
)

var doubles = []string{"00", "11", "22", "33", "44", "55", "66", "77", "88", "99"}

func checkDecrease(n_str string) bool {
	for i, _ := range n_str {
		if i == 0 {
			continue
		}
		n_a, _ := strconv.Atoi(string(n_str[i-1]))
		n_b, _ := strconv.Atoi(string(n_str[i]))
		if n_a > n_b {
			return false
		}
	}
	return true
}

func checkDouble(n_str string, p2 bool) bool {
	for _, d := range doubles {
		if strings.Contains(n_str, d) {
			if p2 {
				i := strings.Index(n_str, d)
				c := d[0]
				if i == 0 && n_str[2] == c {
					continue
				} else if i == 1 && (n_str[0] == c || n_str[3] == c) {
					continue
				} else if i == 2 && (n_str[1] == c || n_str[4] == c) {
					continue
				} else if i == 3 && (n_str[2] == c || n_str[5] == c) {
					continue
				} else if i == 4 && n_str[3] == c {
					continue
				}
			}
			return true
		}
	}
	return false
}

func main() {
	lbound, ubound := 246540, 787419
	var part1, part2 int

	for n := lbound; n < ubound; n++ {
		n_str := strconv.Itoa(n)
		if checkDecrease(n_str) {
			if checkDouble(n_str, false) {
				part1 += 1
			}
			if checkDouble(n_str, true) {
				part2 += 1
			}
		}
	}

	fmt.Printf("part1: %d\npart2: %d\n", part1, part2)
}
