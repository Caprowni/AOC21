package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []int
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		n, _ := strconv.Atoi(dep)
		lines = append(lines, n)
	}

	one := one()
	two := two()
}

func one() {}

func two() {}
