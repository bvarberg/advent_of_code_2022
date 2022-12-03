package actual_strategy

import (
	"aoc02/game"
	"aoc02/scoring"
)

type Round struct {
	opponentThrow game.Shape
	desiredResult game.Outcome
}

func NewRound(opponentThrow game.Shape, desiredResult game.Outcome) Round {
	return Round{
		opponentThrow,
		desiredResult,
	}
}

func (r Round) Score() int {
	return scoring.ScoreOutcome(r.desiredResult) + scoring.ScoreThrow(r.DetermineThrow())
}

func (r Round) DetermineThrow() game.Shape {
	switch r.desiredResult {
	case game.Win:
		return makeWin(r.opponentThrow)
	case game.Draw:
		return makeDraw(r.opponentThrow)
	case game.Loss:
		return makeLoss(r.opponentThrow)
	default:
		panic("Failed to determine what shape to throw")
	}
}

func makeWin(opponentThrow game.Shape) game.Shape {
	switch opponentThrow {
	case game.Rock:
		return game.Paper
	case game.Paper:
		return game.Scissors
	case game.Scissors:
		return game.Rock
	default:
		panic("Unsure how to make a win")
	}
}

func makeDraw(opponentThrow game.Shape) game.Shape {
	switch opponentThrow {
	case game.Rock:
		return game.Rock
	case game.Paper:
		return game.Paper
	case game.Scissors:
		return game.Scissors
	default:
		panic("Unsure how to make a draw")
	}
}

func makeLoss(opponentThrow game.Shape) game.Shape {
	switch opponentThrow {
	case game.Rock:
		return game.Scissors
	case game.Paper:
		return game.Rock
	case game.Scissors:
		return game.Paper
	default:
		panic("Unsure how to make a loss")
	}
}
