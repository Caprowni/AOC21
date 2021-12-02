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
	var lines []int
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		n, _ := strconv.Atoi(dep)
		lines = append(lines, n)
	}

	var measurements int
	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			measurements++
		}
	}

	var threemeasurements int
	for i := 0; i+3 < len(lines); i++ {
		a := lines[i] + lines[i+1] + lines[i+2]
		b := lines[i+1] + lines[i+2] + lines[i+3]
		if b > a {
			threemeasurements++
		}
	}

	fmt.Printf("Original Measurements: %v", measurements)
	fmt.Println("")
	fmt.Printf("Three Measurements: %v", threemeasurements)
}
