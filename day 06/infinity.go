package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type location struct {
	x        int
	y        int
	infinite bool
	size     int
}

var locs []*location
var minX, maxX, minY, maxY int

func main() {
	locs = make([]*location, 0)
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")
	minX, maxX, minY, maxY = 1000, -1000, 1000, -1000
	for _, l := range lines {
		coords := strings.Split(l, ", ")
		loc := location{x: toInt(coords[0]), y: toInt(coords[1])}
		locs = append(locs, &loc)
		if loc.x < minX {
			minX = loc.x
		} else if loc.x > maxX {
			maxX = loc.x
		}
		if loc.y < minY {
			minY = loc.y
		} else if loc.y > maxY {
			maxY = loc.y
		}
	}
	part1()
	part2()
}

func part1() {
	m := make(map[string]*location)
	for x := minX; x < maxX+1; x++ {
		for y := minY; y < maxY+1; y++ {
			closest := shortest(x, y)
			m[coord(x, y)] = closest
			if closest != nil {
				if x == minX || x == maxX || y == maxY || y == minY {
					closest.infinite = true
				}
				closest.size++
			}
		}
	}
	maxSize := 0
	for _, l := range locs {
		if !l.infinite && maxSize < l.size {
			maxSize = l.size
		}
	}
	fmt.Println("Part 1", maxSize)
}

func part2() {
	m := make(map[string]int)
	for x := minX; x < maxX+1; x++ {
		for y := minY; y < maxY+1; y++ {
			sum := 0
			for _, l := range locs {
				sum += dist(x, y, l)
			}
			m[coord(x, y)] = sum
		}
	}
	area := 0
	for _, pixel := range m {
		if pixel < 10000 {
			area++
		}
	}
	fmt.Println("Part 2", area)
}

func coord(x, y int) string {
	return fmt.Sprintf("%v,%v", x, y)
}

func shortest(x, y int) (result *location) {
	shortest := 1000
	for _, l := range locs {
		if d := dist(x, y, l); d == shortest {
			result = nil // Can't have multiple shortest
		} else if d < shortest {
			shortest = d
			result = l
		}
	}
	return result
}

func dist(x, y int, loc *location) int {
	return abs(x-loc.x) + abs(y-loc.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
