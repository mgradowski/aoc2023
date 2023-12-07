package day06

import "github.com/alecthomas/participle/v2"

type Leaderboard struct {
	Times     []int `parser:"'Time'     ':' @Int+"`
	Distances []int `parser:"'Distance' ':' @Int+"`
}

func NewParser() *participle.Parser[Leaderboard] {
	return participle.MustBuild[Leaderboard]()
}

func solve(time int, distance int) (result int) {
	for t := 0; t <= time; t++ {
		d := t * (time - t)
		if d > distance {
			result++
		}
	}
	return
}

func (leaderboard *Leaderboard) Part1() int {
	result := 1
	for i := range leaderboard.Times {
		result *= solve(leaderboard.Times[i], leaderboard.Distances[i])
	}
	return result
}

func smallestPow10GreaterThan(n int) int {
	result := 1
	for result <= n {
		result *= 10
	}
	return result
}

func concatInts(nums ...int) (result int) {
	for _, num := range nums {
		result = smallestPow10GreaterThan(num)*result + num
	}
	return
}

func (leaderboard *Leaderboard) Part2() int {
	return solve(
		concatInts(leaderboard.Times...),
		concatInts(leaderboard.Distances...),
	)
}
