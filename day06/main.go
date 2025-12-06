package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")
	length := len(lines[0])

	var blanks []int
	blanks = append(blanks, -1)
	for i := 0; i < length; i++ {
		isBlank := true
		for c := 0; c < len(lines); c++ {
			if lines[c][i] != ' ' {
				isBlank = false
				break
			}
		}
		if isBlank {
			blanks = append(blanks, i)
		}
	}
	blanks = append(blanks, length)

	{
		fmt.Println("--- Part One ---")

		var grandTotal int64
		for problemIndex := 0; problemIndex+1 < len(blanks); problemIndex++ {
			i, j := blanks[problemIndex]+1, blanks[problemIndex+1] // lol

			operation := strings.TrimSpace(lines[len(lines)-1][i:j])
			var total int64
			if operation == "*" {
				total = 1
			}
			for c := 0; c < len(lines)-1; c++ {
				value := toInt64(strings.TrimSpace(lines[c][i:j]))
				if operation == "*" {
					total *= value
				} else {
					total += value
				}
			}
			grandTotal += total
		}

		fmt.Println(grandTotal)
	}

	{
		fmt.Println("--- Part Two ---")

		var grandTotal int64
		for problemIndex := len(blanks) - 2; problemIndex >= 0; problemIndex-- {
			i, j := blanks[problemIndex]+1, blanks[problemIndex+1]

			operation := strings.TrimSpace(lines[len(lines)-1][i:j])
			var total int64
			if operation == "*" {
				total = 1
			}
			for r := j - 1; r >= i; r-- {
				var value int64
				for c := 0; c < len(lines)-1; c++ {
					digit := lines[c][r]
					if digit != ' ' {
						value = value*10 + int64(digit-'0')
					}
				}
				if operation == "*" {
					total *= value
				} else {
					total += value
				}
			}
			grandTotal += total
		}

		fmt.Println(grandTotal)
	}
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func toInt64(s string) int64 {
	result, err := strconv.ParseInt(s, 10, 64)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
