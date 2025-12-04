package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readLines("input.txt")

	dial := 50
	password := 0
	passwordClick := 0

	for _, line := range lines {
		dir := 1
		if line[0] == 'L' {
			dir = -1
		}
		value := toInt(line[1:])
		for i := 0; i < value; i++ {
			dial = (dial + dir) % 100
			if dial == 0 {
				passwordClick++
			}
		}

		if dial == 0 {
			password++
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(password)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(passwordClick)
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

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
