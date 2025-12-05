package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	banks := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(totalJoltage(banks, 2))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(totalJoltage(banks, 12))
	}
}

func totalJoltage(banks []string, count int) (total int64) {
	for _, bank := range banks {
		total += toInt64(joltage(bank, count))
	}
	return
}

func joltage(input string, count int) string {
	if count == 0 {
		return ""
	}
	digit := input[0]
	index := 0
	for i := 1; i < len(input)-(count-1); i++ {
		if input[i] > digit {
			digit = input[i]
			index = i
		}
	}
	return fmt.Sprintf("%c%s", digit, joltage(input[index+1:], count-1))
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
