package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")
	ranges := strings.Split(input, ",")

	var result1 int64 = 0
	var result2 int64 = 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		min := toInt64(parts[0])
		max := toInt64(parts[1])

		for id := min; id <= max; id++ {
			str := strconv.FormatInt(id, 10)

			for window := len(str) / 2; window > 0; window-- {
				if len(str)%window != 0 {
					continue
				}
				invalid := true
				for offset := window; offset < len(str); offset += window {
					if str[0:window] != str[offset:offset+window] {
						invalid = false
						break
					}
				}
				if invalid {
					if window*2 == len(str) {
						result1 += id
					}
					result2 += id
					break
				}
			}
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(result1)
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(result2)
	}
}

func readFile(filename string) string {
	bytes, err := os.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(bytes))
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
