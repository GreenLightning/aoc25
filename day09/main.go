package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	x, y int
}

func area(a, b Tile) int {
	return (abs(b.x-a.x) + 1) * (abs(b.y-a.y) + 1)
}

type Range struct {
	min, max int
}

func MakeRange(x, y int) Range {
	return Range{
		min(x, y),
		max(x, y),
	}
}

func (r Range) Includes(x int) bool {
	return r.min < x && x < r.max
}

func overlaps(a, b Range) bool {
	return a.min < b.max && a.max > b.min
}

func main() {
	lines := readLines("input.txt")

	var tiles []Tile
	for _, line := range lines {
		parts := strings.Split(line, ",")
		tiles = append(tiles, Tile{
			toInt(parts[0]),
			toInt(parts[1]),
		})
	}

	{
		fmt.Println("--- Part One ---")
		var result int
		for i := 0; i < len(tiles); i++ {
			for j := i + 1; j < len(tiles); j++ {
				a, b := tiles[i], tiles[j]
				result = max(result, area(a, b))
			}
		}
		fmt.Println(result)
	}

	{
		fmt.Println("--- Part Two ---")
		var result int
		for i := 0; i < len(tiles); i++ {
			for j := i + 1; j < len(tiles); j++ {
				a, b := tiles[i], tiles[j]
				xr := MakeRange(a.x, b.x)
				yr := MakeRange(a.y, b.y)
				ok := true
				for k := 0; k < len(tiles); k++ {
					c, d := tiles[k], tiles[(k+1)%len(tiles)]
					if c.x == d.x && xr.Includes(c.x) && overlaps(MakeRange(c.y, d.y), yr) {
						ok = false
						break
					}
					if c.y == d.y && yr.Includes(c.y) && overlaps(MakeRange(c.x, d.x), xr) {
						ok = false
						break
					}
				}
				if ok {
					result = max(result, area(a, b))
				}
			}
		}
		fmt.Println(result)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
