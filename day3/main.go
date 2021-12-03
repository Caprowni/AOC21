package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		dep := strings.TrimSpace(scanner.Text())
		lines = append(lines, dep)
	}

	geRating := one(lines)
	o2Rating := two(lines, 0)
	co2Rating := two(lines, 1)

	o2, _ := strconv.ParseInt(o2Rating, 2, 64)
	co2, _ := strconv.ParseInt(co2Rating, 2, 64)

	fmt.Println(geRating)
	fmt.Println(o2 * co2)
}

func one(lines []string) int64 {

	var g string
	var e string

	half_input := len(lines) / 2
	final_values := make([]int, len(lines[0]))
	for _, line := range lines {
		r := []rune(line)
		for i, num := range r {
			interim := make([]byte, 1)
			_ = utf8.EncodeRune(interim, num)
			bit, _ := strconv.Atoi(string(interim))
			final_values[i] += bit
		}
	}

	for _, v := range final_values {
		if v >= half_input {
			g += "1"
			e += "0"
		} else {
			g += "0"
			e += "1"
		}
	}

	gammaValue, _ := strconv.ParseInt(g, 2, 64)
	epsilonValue, _ := strconv.ParseInt(e, 2, 64)

	return gammaValue * epsilonValue
}

func two(lines []string, search int) string {
	var i int

	for len(lines) > 1 {
		var temp []string
		var onecount, zerocount int
		for _, v := range lines {
			if v[i] == '1' {
				onecount++
			} else {
				zerocount++
			}
		}

		var bit string
		if onecount > zerocount {
			if search == 1 {
				bit = "0"
			} else {
				bit = "1"
			}
		} else if onecount < zerocount {
			if search == 1 {
				bit = "1"
			} else {
				bit = "0"
			}
		} else {
			if search == 1 {
				bit = "0"
			} else {
				bit = "1"
			}
		}

		for _, v := range lines {
			if string(v[i]) == bit {
				temp = append(temp, v)
			}
		}
		i++
		lines = make([]string, len(temp))
		copy(lines, temp)
	}
	return lines[0]
}
