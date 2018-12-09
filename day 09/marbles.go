package main

import (
	"container/ring"
	"fmt"
)

func main() {
	fmt.Println("Part 1", game(71082, 413))
	fmt.Println("Part 2", game(71082*100, 413))
}

func game(marbles, maxPlayers int) int {
	r := ring.New(1)
	r.Value = 0
	player := 1
	scores := make(map[int]int, maxPlayers)
	for i := 1; i < marbles+1; i++ {
		if i%23 == 0 {
			r = r.Move(-8)
			deleted := r.Unlink(1)
			scores[player] += deleted.Value.(int) + i
			r = r.Next()
		} else {
			m := ring.New(1)
			m.Value = i
			r = r.Move(1)
			r = r.Link(m)
			r = r.Prev()
		}
		player = player%maxPlayers + 1
	}
	return max(scores)
}

func max(m map[int]int) (r int) {
	for _, v := range m {
		if v > r {
			r = v
		}
	}
	return r
}
