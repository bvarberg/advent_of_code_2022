package scoring

import (
	"aoc02/game"
	"fmt"
)

func ScoreOutcome(outcome game.Outcome) int {
	switch outcome {
	case game.Win:
		return 6
	case game.Draw:
		return 3
	case game.Loss:
		return 0
	default:
		panic(fmt.Sprintf("Invalid outcome: %q", outcome))
	}
}

func ScoreThrow(shape game.Shape) int {
	switch shape {
	case game.Rock:
		return 1
	case game.Paper:
		return 2
	case game.Scissors:
		return 3
	default:
		panic(fmt.Sprintf("Invalid throw: %q", shape))
	}
}
