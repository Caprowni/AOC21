package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		lines = append(lines, dep)
	}

	var fwd int
	var dwn int
	var aim int

	part1 := one(lines, fwd, dwn)
	part2 := two(lines, fwd, dwn, aim)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)

}

func one(lines []string, fwd, dwn int) int {

	for i := 0; i < len(lines); i++ {
		r := strings.Split(lines[i], " ")
		y, _ := strconv.Atoi(r[1])
		if r[0] == "forward" {
			fwd += y
		} else if r[0] == "up" {
			dwn -= y
		} else {
			dwn += y
		}
	}
	return fwd * dwn

}

func two(lines []string, fwd, dwn, aim int) int {

	for i := 0; i < len(lines); i++ {
		r := strings.Split(lines[i], " ")
		y, _ := strconv.Atoi(r[1])
		if r[0] == "forward" {
			dwn += y * aim
			fwd += y
		} else if r[0] == "up" {
			aim -= y
		} else {
			aim += y
		}
	}
	return fwd * dwn
}
