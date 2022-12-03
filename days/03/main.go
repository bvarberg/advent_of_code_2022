package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Find the item type that appears in both compartments of each rucksack.
// What is the sum of the priorities of those item types?

func main() {
	rucksacks := readInput("./data/input.txt")

	sumPriorities := 0
	for _, rucksack := range rucksacks {
		indexOfIntersection := strings.IndexAny(rucksack.compartmentA, rucksack.compartmentB)
		item := rucksack.compartmentA[indexOfIntersection]
		priority := getPriority(item)
		sumPriorities += priority
	}

	fmt.Printf("Sum of priorities: %d\n", sumPriorities)
}

type rucksack struct {
	compartmentA string
	compartmentB string
}

func readInput(name string) []rucksack {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rucksacks := []rucksack{}
	for scanner.Scan() {
		line := scanner.Text()
		rucksacks = append(rucksacks, toRucksack(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return rucksacks
}

func toRucksack(line string) rucksack {
	middle := len(line) / 2
	compartmentA := line[:middle]
	compartmentB := line[middle:]

	return rucksack{
		compartmentA,
		compartmentB,
	}
}

const priorityOrder = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getPriority(c byte) int {
	return strings.IndexByte(priorityOrder, c) + 1
}
