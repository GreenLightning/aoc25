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

type Vector3 struct {
	x, y, z int
}

func (v Vector3) Minus(other Vector3) Vector3 {
	return Vector3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v Vector3) LengthSquared() int {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vector3) DistanceSquared(o Vector3) int {
	return v.Minus(o).LengthSquared()
}

func main() {
	lines := readLines("input.txt")

	var boxes []Vector3
	for _, line := range lines {
		parts := strings.Split(line, ",")
		boxes = append(boxes, Vector3{
			toInt(parts[0]),
			toInt(parts[1]),
			toInt(parts[2]),
		})
	}

	type Entry struct {
		i, j, d int
	}

	var distances []Entry
	for i := range len(boxes) {
		for j := i + 1; j < len(boxes); j++ {
			d := boxes[i].DistanceSquared(boxes[j])
			distances = append(distances, Entry{i, j, d})
		}
	}

	slices.SortFunc(distances, func(a, b Entry) int {
		return cmp.Compare(a.d, b.d)
	})

	type Circuit struct {
		Boxes []int
	}

	boxToCircuit := make(map[int]*Circuit)
	for i := range len(boxes) {
		boxToCircuit[i] = &Circuit{Boxes: []int{i}}
	}

	index := 0
	for ; index < 1000; index++ {
		entry := distances[index]
		ci := boxToCircuit[entry.i]
		cj := boxToCircuit[entry.j]
		if ci != cj {
			ci.Boxes = append(ci.Boxes, cj.Boxes...)
			for _, b := range cj.Boxes {
				boxToCircuit[b] = ci
			}
		}
	}

	{
		fmt.Println("--- Part One ---")

		var circuits []*Circuit
		for _, c := range boxToCircuit {
			if !slices.Contains(circuits, c) {
				circuits = append(circuits, c)
			}
		}

		slices.SortFunc(circuits, func(a, b *Circuit) int {
			return -cmp.Compare(len(a.Boxes), len(b.Boxes))
		})

		result := 1
		for i := 0; i < 3; i++ {
			result *= len(circuits[i].Boxes)
		}

		fmt.Println(result)
	}

	{
		fmt.Println("--- Part Two ---")

		for ; index < len(distances); index++ {
			entry := distances[index]
			ci := boxToCircuit[entry.i]
			cj := boxToCircuit[entry.j]
			if ci != cj {
				ci.Boxes = append(ci.Boxes, cj.Boxes...)
				for _, b := range cj.Boxes {
					boxToCircuit[b] = ci
				}
				if len(ci.Boxes) == len(boxes) {
					fmt.Println(boxes[entry.i].x * boxes[entry.j].x)
					break
				}
			}
		}
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
