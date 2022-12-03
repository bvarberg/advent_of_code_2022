package main

import (
	strategy "aoc02/strategy"
	"fmt"
)

func main() {
	strategyGuide := strategy.NewStrategyGuide("./data/input.txt")

	fmt.Printf("Expected total score: %d", strategyGuide.ExpectedScore())
}
