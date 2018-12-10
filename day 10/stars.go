package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	vectors := readInput()
	f, _ := os.Create("output.txt")
	defer f.Close()
	for i := 1; i < 100000; i++ {
		move(vectors)
		print(vectors, f, i)
	}
}

type vector struct {
	px int
	py int
	vx int
	vy int
}

func readInput() (r []*vector) {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")
	re := regexp.MustCompile(`position=<\s*(-*\d+),\s*(-*\d+)> velocity=<\s*(-*\d+),\s*(-*\d+)>$`)
	for _, l := range lines {
		matches := re.FindStringSubmatch(l)
		r = append(r, &vector{toInt(matches[1]), toInt(matches[2]), toInt(matches[3]), toInt(matches[4])})
	}
	return r
}

func move(vectors []*vector) {
	for _, v := range vectors {
		v.px = v.px + v.vx
		v.py = v.py + v.vy
	}
}

func print(vectors []*vector, f *os.File, i int) {
	minX, minY, maxX, maxY := edges(vectors)
	m := make(map[string]string)
	for _, v := range vectors {
		m[key(v.px, v.py)] = "#"
	}

	if maxY-minY < 30 {
		fmt.Fprintf(f, "New state at %v \n", i)
		for y := minY; y <= maxY; y++ {
			line := ""
			for x := minX; x <= maxX; x++ {
				_, exists := m[key(x, y)]
				if exists {
					line += "#"
				} else {
					line += " "
				}
			}
			fmt.Fprintln(f, line)
		}
	}

}

func key(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
func toInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func edges(vectors []*vector) (int, int, int, int) {
	minX, minY, maxX, maxY := 1000, 1000, -1000, -1000
	for _, v := range vectors {
		if v.px > maxX {
			maxX = v.px
		} else if v.px < minX {
			minX = v.px
		}
		if v.py > maxY {
			maxY = v.py
		} else if v.py < minY {
			minY = v.py
		}
	}
	return minX, minY, maxX, maxY
}
