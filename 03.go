package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	x     int
	y     int
	steps int
}

var coords_a = []coord{}
var coords_b = []coord{}
var ax, ay, bx, by, sa, sb int

func parseMove(s string) (string, int) {
	dir := string(s[0])
	n, _ := strconv.Atoi(s[1:])
	return dir, n
}

func moveWire(dir string, n int, wire string) {
	for i := 0; i < n; i++ {
		if dir == "U" {
			if wire == "a" {
				ay += 1
			} else {
				by += 1
			}
		} else if dir == "D" {
			if wire == "a" {
				ay -= 1
			} else {
				by -= 1
			}
		} else if dir == "L" {
			if wire == "a" {
				ax -= 1
			} else {
				bx -= 1
			}
		} else if dir == "R" {
			if wire == "a" {
				ax += 1
			} else {
				bx += 1
			}
		}

		if wire == "a" {
			sa += 1
			coords_a = append(coords_a, coord{ax, ay, sa})
		} else {
			sb += 1
			coords_b = append(coords_b, coord{bx, by, sb})
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMin(int_list []int) int {
	var min int
	for i, val := range int_list {
		if i == 0 || val < min {
			min = val
		}
	}
	return min
}

func main() {
	data, _ := ioutil.ReadFile("03_input.txt")
	lines := strings.Split(string(data), "\n")
	wire_a := strings.Split(lines[0], ",")
	wire_b := strings.Split(lines[1], ",")

	for i := range wire_a {
		dir_a, n_a := parseMove(wire_a[i])
		dir_b, n_b := parseMove(wire_b[i])
		moveWire(dir_a, n_a, "a")
		moveWire(dir_b, n_b, "b")
	}

	crosses, steps := []int{}, []int{}
	for _, ca := range coords_a {
		for _, cb := range coords_b {
			if ca.x == cb.x && ca.y == cb.y {
				crosses = append(crosses, abs(ca.x)+abs(ca.y))
				steps = append(steps, ca.steps+cb.steps)
			}
		}
	}

	part1 := getMin(crosses)
	part2 := getMin(steps)

	fmt.Printf("part1: %d\npart2: %d\n", part1, part2)
}
