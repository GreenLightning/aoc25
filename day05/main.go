package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Min int64
	Max int64
}

func main() {
	lines := readLines("input.txt")

	var ranges []Range
	var ingredients []int64
	{
		i := 0
		for ; i < len(lines); i++ {
			line := lines[i]

			if line == "" {
				break
			}

			parts := strings.Split(line, "-")
			ranges = append(ranges, Range{
				Min: toInt64(parts[0]),
				Max: toInt64(parts[1]),
			})
		}

		for i++; i < len(lines); i++ {
			ingredients = append(ingredients, toInt64(lines[i]))
		}
	}

	{
		fmt.Println("--- Part One ---")

		numFresh := 0
		for _, ingredient := range ingredients {
			isFresh := false
			for _, r := range ranges {
				if r.Min <= ingredient && ingredient <= r.Max {
					isFresh = true
					break
				}
			}
			if isFresh {
				numFresh++
			}
		}

		fmt.Println(numFresh)
	}

	{
		fmt.Println("--- Part Two ---")

		slices.SortFunc(ranges, func(a, b Range) int {
			return cmp.Compare(a.Min, b.Min)
		})

		var numFresh int64 = 0
		var lastMax int64 = 0
		for _, r := range ranges {
			if r.Max <= lastMax {
				continue
			}
			if r.Min > lastMax {
				numFresh += r.Max - r.Min + 1
			} else {
				numFresh += r.Max - lastMax
			}
			lastMax = r.Max
		}

		fmt.Println(numFresh)
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
