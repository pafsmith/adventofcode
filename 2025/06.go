package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	part1 := solvePart1(lines)
	fmt.Println("Part 1:", part1)

	part2 := solvePart2(lines)
	fmt.Println("Part 2:", part2)
}

func solvePart1(lines []string) int {
	var ba [][]string
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			ba = append(ba, fields)
		}
	}

	total := 0
	operatorIndex := len(ba) - 1
	rowLength := len(ba[0])

	for i := 0; i < rowLength; i++ {
		count := 0
		if ba[operatorIndex][i] == "+" {
			count = 0
		} else if ba[operatorIndex][i] == "*" {
			count = 1
		}

		for row := range ba {
			if row == operatorIndex {
				continue
			} else {
				if ba[operatorIndex][i] == "+" {
					val, _ := strconv.Atoi(ba[row][i])
					count += val
				} else if ba[operatorIndex][i] == "*" {
					val, _ := strconv.Atoi(ba[row][i])
					count *= val
				}
			}
		}
		total += count
	}

	return total
}

func solvePart2(lines []string) int {
	var filteredLines []string
	for _, line := range lines {
		if len(line) > 0 {
			filteredLines = append(filteredLines, line)
		}
	}
	lines = filteredLines

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	for i := range lines {
		for len(lines[i]) < maxLen {
			lines[i] += " "
		}
	}

	operatorRow := len(lines) - 1
	numRows := operatorRow
	total := 0

	col := maxLen - 1
	for col >= 0 {
		if lines[operatorRow][col] == ' ' {
			col--
			continue
		}

		operator := lines[operatorRow][col]
		var numbers []int

		rightCols := []int{}
		for c := col + 1; c < maxLen && lines[operatorRow][c] == ' '; c++ {
			rightCols = append(rightCols, c)
		}

		allCols := append(rightCols, col)

		for _, c := range allCols {
			numStr := ""
			for row := 0; row < numRows; row++ {
				char := lines[row][c]
				if char != ' ' {
					numStr += string(char)
				}
			}
			if numStr != "" {
				num, _ := strconv.Atoi(numStr)
				numbers = append(numbers, num)
			}
		}

		if len(numbers) > 0 {
			var result int
			if operator == '+' {
				result = 0
				for _, num := range numbers {
					result += num
				}
			} else if operator == '*' {
				result = 1
				for _, num := range numbers {
					result *= num
				}
			}
			total += result
		}
		col--
	}

	return total
}
