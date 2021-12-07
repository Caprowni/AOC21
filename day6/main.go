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
	var fish []int
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		lines = append(lines, dep)
	}

	numbers := make([]string, len(lines[0]))
	numbers = strings.Split(lines[0], ",")

	for _, n := range numbers {
		n,_ := strconv.Atoi(n)
		fish = append(fish, n)
	}
	// More efficient part 2...
	ages := make(map[int]int)
	for _, ns := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(ns)
		ages[n] += 1
	}
	one := one(fish)
	fmt.Println(one)

	two := two(ages)
	fmt.Println(two)
}

func one(fish []int) int {
	for count := 0; count < 80; count++ {
		for i, x := range fish {
			if x == 0 {
				fish = append(fish, 8)
				fish[i] += 6
			} else {
				fish[i] -= 1
			}
		}
	}
	return len(fish)
}

func two(ages map[int]int) int {
	result := 0
	for count := 0; count < 256; count++ {
		tempAges := make(map[int]int)
		for age, count := range ages {
			if age > 0 {
				tempAges[age-1] += count
			} else {
				tempAges[6] += count
				tempAges[8] += count
			}
		}
		ages = tempAges
	}
	
	for _, count := range ages {
		result += count
	}

	return result
}
