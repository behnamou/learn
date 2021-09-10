package main

import (
	"fmt"
	"math"
)

type Student struct {
	score1, score2, score3 float64
}

func (s Student) Average() float64 {
	return (s.score1 + s.score2 + s.score3) / 3
}

func (s Student) Min() float64 {
	return math.Min(s.score1, math.Min(s.score2, s.score3))
}

func (s Student) Max() float64 {
	return math.Max(s.score1, math.Max(s.score2, s.score3))
}

func main() {
	s1 := Student{12.25, 15.75, 20}
	fmt.Println("Min :", s1.Min(), "| Max :", s1.Max(), "| Ave. :", s1.Average())
}
