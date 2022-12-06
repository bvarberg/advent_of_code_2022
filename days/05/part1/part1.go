package part1

import (
	"aoc05/stack"
	"fmt"
	"readers"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	stacks := loadInitialStacks("./data/simple/initial_stacks.txt")
	moves := loadMoves("./data/simple/moves.txt")
	fmt.Println(stacks)
	fmt.Printf("%+v", moves)
}

func loadMoves(name string) (moves []Move) {
	lines := readers.ReadLines(name)

	moves = []Move{}
	for _, line := range lines {
		moves = append(moves, NewMove(line))
	}

	return
}

func loadInitialStacks(name string) (stacks []stack.StringStack) {
	lines := readers.ReadLines(name)

	stacks = []stack.StringStack{}
	for _, line := range lines {
		stacks = append(stacks, stack.NewStringStack(strings.Split(line, ",")...))
	}

	return
}

// Move is a structure representing a single rearrangement instruction: `move 1 from 2 to 1`
type Move struct {
	// How many?
	n int
	// From which source?
	from int
	// To which destination?
	to int
}

func NewMove(elfReadable string) Move {
	digitRe := regexp.MustCompile(`\d{1,}`)
	values := digitRe.FindAllString(elfReadable, -1)

	n, _ := strconv.Atoi(values[0])
	from, _ := strconv.Atoi(values[1])
	to, _ := strconv.Atoi(values[2])

	return Move{
		n,
		from,
		to,
	}
}
