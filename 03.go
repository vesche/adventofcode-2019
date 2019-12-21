package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

var coords_a = []coord{}
var coords_b = []coord{}
var ax, ay, bx, by int

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
			coords_a = append(coords_a, coord{ax, ay})
		} else {
			coords_b = append(coords_b, coord{bx, by})
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	crosses := []int{}
	for _, ca := range coords_a {
		for _, cb := range coords_b {
			if ca.x == cb.x && ca.y == cb.y {
				crosses = append(crosses, abs(ca.x)+abs(ca.y))
			}
		}
	}

	var part1 int
	for i, val := range crosses {
		if i == 0 || val < part1 {
			part1 = val
		}
	}

	fmt.Printf("part1: %d\npart2: %d\n", part1, 0)
}
