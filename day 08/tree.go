package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	children []*node
	metadata []int
}

func main() {
	root := readInput()
	fmt.Println("Part 1", sum(root))
	fmt.Println("Part 2", root.value())
}

func readInput() *node {
	raw, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(raw), " ")
	numbers := make([]int, len(split))
	for i, s := range split {
		numbers[i] = toInt(s)
	}
	n, _ := parse(numbers)
	return n
}

func sum(n *node) (s int) {
	for _, c := range n.children {
		s += sum(c)
	}
	for _, meta := range n.metadata {
		s += meta
	}
	return s
}

func parse(input []int) (*node, int) {
	addToIndex := 2
	n := &node{}
	for k := input[0]; k > 0; k-- {
		child, newAdd := parse(input[addToIndex:])
		addToIndex += newAdd
		n.children = append(n.children, child)
	}
	n.metadata = input[addToIndex : addToIndex+input[1]]
	addToIndex += input[1]
	return n, addToIndex
}

func toInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

// Part 2
func (n *node) value() (s int) {
	if len(n.children) == 0 {
		for _, meta := range n.metadata {
			s += meta
		}
		return s
	}
	for _, meta := range n.metadata {
		if meta > len(n.children) {
			s += 0
		} else {
			s += n.children[meta-1].value()
		}
	}
	return s
}
