package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		lines = append(lines, dep)
	}

	var vents [][][]int

	for _, v := range lines {
		fmt.Println(lines)
	}
	one := one()
	two := two()
}

func one() {}

func two() {}
