package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		lines = append(lines, dep)
	}

	one := one(lines)
	fmt.Println(one)
	//two := two()
}

func one(lines []string) string {
	bingoNumbers := make([]string, len(lines[0]))
	bingoNumbers = strings.Split(lines[0],",")
	fmt.Println(bingoNumbers)
	for i := 1; i < len(lines); i++{
		fmt.Println(lines[i])
	}
	//for _, line := range lines {
	//	r := []rune(line)
	//	for i, num := range r {
	//	}
	//}
	//for _,v := range bingoNumbers {
	//	fmt.Println(v)
	//}

	return ""
}

func two() {}
