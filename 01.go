package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("01_input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var part1, part2 int
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		first := true
		for {
			i = i/3 - 2
			if i < 0 {
				break
			}
			part2 += i
			if first {
				part1 += i
				first = false
			}
		}
	}
	fmt.Printf("part1: %d\npart2: %d\n", part1, part2)
}
