package strategy

import "fmt"

func ScoreOutcome(outcome Outcome) int {
	switch outcome {
	case Win:
		return 6
	case Draw:
		return 3
	case Loss:
		return 0
	default:
		panic(fmt.Sprintf("Invalid outcome: %q", outcome))
	}
}

func ScoreThrow(throw Shape) int {
	switch throw {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		panic(fmt.Sprintf("Invalid shape: %q", throw))
	}
}
