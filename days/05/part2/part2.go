package part2

import (
	"aoc05/stack"
	"fmt"
	"log"
	"readers"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	stacks := loadInitialStacks("./data/input/initial_stacks.txt")
	moves := loadMoves("./data/input/moves.txt")

	for moveNumber, move := range moves {
		crates := []string{}
		for i := 0; i < move.n; i++ {
			crate, ok := stacks[move.from-1].Pop()
			if !ok {
				log.Fatalf("couldn't pop from the stack on move (%d): %+v", moveNumber, move)
			}
			crates = append(crates, crate)
		}

		for j := move.n; j > 0; j-- {
			stacks[move.to-1].Push(crates[j-1])
		}
	}

	topOfStacks := []string{}
	for _, stack := range stacks {
		crate, _ := stack.Top()
		topOfStacks = append(topOfStacks, crate)
	}

	fmt.Println(stacks)
	fmt.Printf("Top of each stack (CrateMover 9001): %s\n", strings.Join(topOfStacks, ""))
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
