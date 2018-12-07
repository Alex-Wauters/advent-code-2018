package main

func part2(steps map[string]*step) int {
	oven := make([]*step, 0)
	for i := 0; true; i++ {
		oven = removeAllCooked(i, oven)
		oven = fillOven(i, oven, steps)
		if len(oven) == 0 && !remaining(steps) {
			return i
		}
	}
	return 0
}

func fillOven(i int, oven []*step, steps map[string]*step) []*step {
	for len(oven) < 5 {
		next := getNextPt2(steps)
		if next == nil {
			return oven
		}
		next.started = i
		oven = append(oven, next)
	}
	return oven
}

func getNextPt2(list map[string]*step) (result *step) {
	for _, s := range list {
		if !s.finished && s.canStart() && s.started == -1 && (result == nil || s.name < result.name) {
			result = s
		}
	}
	return result
}

func removeAllCooked(i int, oven []*step) (r []*step) {
	for _, s := range oven {
		if !s.isCookedAt(i) {
			r = append(r, s)
		} else {
			s.finished = true
		}
	}
	return r
}

func (s *step) isCookedAt(i int) bool {
	return s.started+cookTime(s.name) <= i
}

func cookTime(r string) int {
	return int(r[0]) - int('A') + 61
}
