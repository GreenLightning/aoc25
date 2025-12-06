package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")
		fmt.Println(len(lines))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println()
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

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
}

func readFile(filename string) string {
	bytes, err := os.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
}

func deleteEmptyStrings(input []string) []string {
	return slices.DeleteFunc(input, func(s string) bool { return s == "" })
}

func arrayToInt(input []string) (output []int) {
	output = make([]int, len(input))
	for i, text := range input {
		output[i] = toInt(text)
	}
	return output
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
