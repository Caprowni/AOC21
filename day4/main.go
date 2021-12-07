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

	var boards [][]int
	var bingoNumbers []int

	numbers := make([]string, len(lines[0]))
	numbers = strings.Split(lines[0], ",")

	for _, n := range numbers {
		n,_ := strconv.Atoi(n)
		bingoNumbers = append(bingoNumbers, n)
	}

	for i := 2; i < len(lines); i += 6 {
		board := make([]int, 25)
		for j := 0; j < 5; j++ {
			for idx, v := range strings.Fields(lines[i+j]) {
				n, _ := strconv.Atoi(v)
				board[5*j+idx] = n
			}
		}
		boards = append(boards, board)
	}

	one := one(bingoNumbers, boards)
	fmt.Println(one)
	two := two(bingoNumbers, boards)
	fmt.Println(two)
}

func one(bingoNumbers []int, boards[][]int ) int {
	var bingo bool
	var row, column int
	for _, n := range bingoNumbers {
		for board := range boards {
			for i := range boards[board] {
				if boards[board][i] == n {
					boards[board][i] = -1
				}
			}
		}
		for _, board := range boards {
			for x := 0; x < 5; x++ {
				row = -1
				column = -1
				for y := 0; y < 5; y++ {
					row &= board[5*x+y]
					column &= board[5*y+x]
				}
				if row < 0 || column < 0 {
					bingo = true
				}
			}
			if bingo {
				result := 0
				for _, number := range board {
					if number != -1 {
						result += number
					}
				}
				return result * n
			}
		}
	}
	return 0
}

func removeBoard(boards [][]int, i int) [][]int{
	return append(boards[:i], boards[i+1:]...)
}

func two(bingoNumbers []int, boards [][]int) int {
	var bingo bool

	for _, n := range bingoNumbers {
		for board := range boards {
			for i := range boards[board] {
				if boards[board][i] == n {
					boards[board][i] = -1
				}
			}
		}
		for i, board := range boards {
			for r := 0; r < 5; r++ {
				row := -1
				col := -1
				for c := 0; c < 5; c++ {
					row &= board[5*r+c]
					col &= board[5*c+r]
				}
				if row < 0 || col < 0 {
					bingo = true
				}
			}
			if bingo {
				bingo = false
				if len(boards) == 1 {
					result := 0
					for _, number := range board {
						if number != -1 {
							result += number
						}
					}
					return result * n
				}
				if i >= 0 && i < len(boards) {
					boards = removeBoard(boards, i)
				}
			}
		}
	}
	return 0
}