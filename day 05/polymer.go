package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")
	fmt.Println("Part 1", react(string(bytes)))
	fmt.Println("Part 2", part2(string(bytes)))
}

func react(line string) int {
	s := shrink(line)
	for len(s) < len(line) {
		line, s = s, shrink(s)
	}
	return len(s)
}

func part2(line string) int {
	min := len(line)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		newline := strings.Replace(line, string(r), "", -1)
		newline = strings.Replace(newline, strings.ToUpper(string(r)), "", -1)
		m := react(newline)
		if m < min {
			min = m
		}
	}
	return min
}

func shrink(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] && strings.EqualFold(string(s[i]), string(s[i+1])) {
			return s[0:i] + s[i+2:]
		}
	}
	return s
}
