package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readLines("input.txt")

	{
		fmt.Println("--- Part One ---")

		grid := make([][]byte, len(lines))
		for y, line := range lines {
			grid[y] = []byte(line)
		}

		for x, char := range grid[0] {
			if char == 'S' {
				grid[0][x] = '|'
			}
		}

		var splits int
		for y := 1; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == '^' {
					if grid[y-1][x] == '|' {
						splits++
						grid[y][x-1] = '|'
						grid[y][x+1] = '|'
					}
				} else {
					if grid[y-1][x] == '|' {
						grid[y][x] = '|'
					}
				}
			}
		}

		fmt.Println(splits)
	}

	{
		fmt.Println("--- Part Two ---")

		universes := make([][]int64, len(lines))
		for y, line := range lines {
			universes[y] = make([]int64, len(line))
		}

		for x, char := range lines[0] {
			if char == 'S' {
				universes[0][x] = 1
			}
		}

		// Note that the quantum variant is simpler than the classical one!
		for y := 1; y < len(universes); y++ {
			for x := 0; x < len(universes[y]); x++ {
				if lines[y][x] == '^' {
					universes[y][x-1] += universes[y-1][x]
					universes[y][x+1] += universes[y-1][x]
				} else {
					universes[y][x] += universes[y-1][x]
				}
			}
		}

		var total int64
		for _, count := range universes[len(universes)-1] {
			total += count
		}

		fmt.Println(total)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
