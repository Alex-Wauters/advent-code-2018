package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]int, 300*300)
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			rackId := x + 10
			level := (rackId*y + 7989) * rackId
			str := strconv.Itoa(level)
			d := string(str[len(str)-3])
			m[key(x, y)] = toInt(d) - 5
		}
	}
	part1(m)
	part2(m)
}

func part2(m map[string]int) {
	maxFuel := 0
	maxKey := ""
	maxSize := 0
	for x := 1; x <= 300-1; x++ {
		for y := 1; y <= 300-1; y++ {
			sum := m[key(x, y)]
			for size := 2; size < maxBoxSize(x, y); size++ {
				for a := 0; a < size; a++ {
					sum += m[key(x+size-1, y+a)]
					sum += m[key(x+a, y+size-1)]
				}
				sum += m[key(x+size-1, y+size-1)]
				if sum > maxFuel {
					maxFuel = sum
					maxKey = key(x, y)
					maxSize = size
				}
			}

		}
	}

	fmt.Printf("Part 2: Max key is %s and size %v \n", maxKey, maxSize)
}

func maxBoxSize(x, y int) int {
	if x < y {
		return 301 - y
	} else {
		return 301 - x
	}
}

func part1(m map[string]int) {
	maxFuel := 0
	maxKey := ""
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			sum := 0
			for a := 0; a < 3; a++ {
				for b := 0; b < 3; b++ {
					sum += m[key(x+a, y+b)]
				}
			}
			if sum > maxFuel {
				maxFuel = sum
				maxKey = key(x, y)
			}
		}
	}
	fmt.Printf("Max key is %s \n", maxKey)
}

func key(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}
