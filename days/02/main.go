package main

import (
	actual_strategy "aoc02/actual_strategy"
	strategy "aoc02/strategy"
	"fmt"
)

func main() {
	strategyGuide := strategy.NewStrategyGuide("./data/input.txt")
	fmt.Printf("Expected total score: %d\n", strategyGuide.ExpectedScore())

	actualStrategyGuide := actual_strategy.NewActualStrategyGuide("./data/input.txt")
	fmt.Printf("Actual expected total score: %d\n", actualStrategyGuide.ExpectedScore())
}
