package actual_strategy

import (
	"aoc02/game"
	"encoding/csv"
	"os"
)

type ActualStrategyGuide struct {
	rounds []Round
}

func NewActualStrategyGuide(name string) ActualStrategyGuide {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	rounds := make([]Round, len(records))
	for index, record := range records {
		round := Round{
			opponentThrow: shapeMap[record[0]],
			desiredResult: outcomeMap[record[1]],
		}
		rounds[index] = round
	}

	return ActualStrategyGuide{
		rounds,
	}
}

var shapeMap = map[string]game.Shape{
	"A": game.Rock,
	"B": game.Paper,
	"C": game.Scissors,
}

var outcomeMap = map[string]game.Outcome{
	"X": game.Loss,
	"Y": game.Draw,
	"Z": game.Win,
}

func (s ActualStrategyGuide) ExpectedScore() int {
	totalScore := 0
	for _, round := range s.rounds {
		totalScore += round.Score()
	}
	return totalScore
}
