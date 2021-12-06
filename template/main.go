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
		lines := strconv.Atoi(dep)
	}

	one := one()
	two := two()
}

func one() {}

func two() {}
