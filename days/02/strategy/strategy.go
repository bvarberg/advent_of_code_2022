package strategy

import (
	"encoding/csv"
	"os"
)

type StrategyGuide struct {
	rounds []Round
}

var shapeMap = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func NewStrategyGuide(name string) StrategyGuide {
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

	var rounds []Round
	for _, record := range records {
		round := Round{
			opponentThrow: shapeMap[record[0]],
			throw:         shapeMap[record[1]],
		}
		rounds = append(rounds, round)
	}

	return StrategyGuide{
		rounds,
	}
}

func (s StrategyGuide) ExpectedScore() int {
	totalScore := 0
	for _, round := range s.rounds {
		totalScore += round.Score()
	}
	return totalScore
}
