package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type step struct {
	name     string
	deps     []string
	next     []*step
	finished bool
	prev     []*step
	started  int
}

func main() {
	part1(readInput())
	fmt.Println("Part 2", part2(readInput()))
}

func readInput() map[string]*step {
	steps := make(map[string]*step)
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")
	re := regexp.MustCompile(`^Step (\w) must be finished before step (\w) can begin.$`)
	for _, l := range lines {
		matches := re.FindStringSubmatch(l)
		if steps[matches[2]] != nil {
			steps[matches[2]].deps = append(steps[matches[2]].deps, matches[1])
		} else {
			steps[matches[2]] = &step{name: matches[2], deps: []string{matches[1]}, next: []*step{}, prev: []*step{}, started: -1}
		}
	}
	for _, v := range steps {
		for _, dep := range v.deps {
			prev := steps[dep]
			if prev == nil {
				steps[dep] = &step{next: []*step{v}, name: dep, started: -1}
			} else {
				prev.next = append(prev.next, v)
				v.prev = append(v.prev, prev)
			}
		}
	}
	return steps
}

func part1(steps map[string]*step) {
	pattern := ""
	var next *step
	for {
		next = getNext(steps)
		next.finished = true
		pattern += next.name
		if !remaining(steps) {
			fmt.Println(pattern)
			return
		}
	}
}

func getNext(list map[string]*step) (result *step) {
	for _, s := range list {
		if !s.finished && s.canStart() && (result == nil || s.name < result.name) {
			result = s
		}
	}
	return result
}

func (s *step) canStart() bool {
	for _, d := range s.prev {
		if !d.finished {
			return false
		}
	}
	return true
}

func remaining(l map[string]*step) bool {
	for _, s := range l {
		if !s.finished {
			return true
		}
	}
	return false
}
