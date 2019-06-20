package average

type S struct {
	L []int
}

func (s *S) Average() float64 {
	// Guard
	if len(s.L) == 0 {
		return 0.0
	}

	total := 0
	for _, v := range s.L {
		total += v
	}
	return float64(total) / float64(len(s.L))
}
