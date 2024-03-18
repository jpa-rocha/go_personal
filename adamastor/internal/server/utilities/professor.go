package utilities

import (
	"math/rand"
)

const (
	levelOne   = 1
	levelTwo   = 2
	levelThree = 3
	levelFour  = 4
	levelFive  = 5

	numVals   = 2
)

type Game struct {
	NumRounds  int
	Level      int
	Win        int
	Loss       int
	Result     int
	Answer     int
	Round      []int
	Op         string
	Operations []string
	Err        error
}
func (g *Game) MakeRound() {
    if g.Op == "R" {
        g.Op == g.Operations[rand.Intn(3)]
    }
    g.numGenerator()
    g.generateResults()

}

// func RunGame(game *Game) {
// 	change := true
// 	for {
// 		if change {
// 		}
// 		getAnswer(game)
// 		if game.answer == game.result {
// 			game.win++
// 			change = true
// 		} else {
// 			game.loss++
// 			change = false
// 		}
// 		game.roundTime = time.Since(game.roundStart)
// 		game.totalTime += game.roundTime
// 		printResults(*game)
// 		if game.win >= game.numRounds || game.loss >= game.numRounds {
// 			break
// 		}
// 	}
// }

func (g *Game) generateResults() {
	switch g.Op {
	case "+":
		g.Result = g.Round[0] + g.Round[1]
	case "-":
		g.Result = g.Round[0] - g.Round[1]
	case "*":
		g.Result = g.Round[0] * g.Round[1]
	case "/":
		g.Result = g.Round[0] / g.Round[1]
	}
}

func (g *Game) numGenerator() {
	g.Round = make([]int, numVals)
	var top int
	var bot int
	var topMin int
	var botMin int
	switch g.Level {
	case levelOne:
		top = 10
		topMin = 1
		bot = top
		botMin = topMin
	case levelTwo:
		top = 100
		topMin = 10
		bot = 10
		botMin = 1
	case levelThree:
		top = 100
		topMin = 10
		bot = top
		botMin = topMin
	case levelFour:
		top = 1000
		topMin = 100
		bot = 100
		botMin = 10
	case levelFive:
		top = 1000
		topMin = 100
		bot = 1000
		botMin = topMin
	}
	g.Round[0] = rand.Intn(top-topMin) + topMin
	g.Round[1] = rand.Intn(bot-botMin) + botMin
	if g.Round[0] < g.Round[1] {
		g.Round = swap(g.Round)
	}
	if g.Round[1] == 0 {
		g.Round[1]++
	}
}

func swap(nums []int) []int {
	nums[0], nums[1] = nums[1], nums[0]
	return nums
}
