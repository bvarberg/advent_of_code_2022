package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inventories := readInventories("./data/input.txt")

	calorieCounts := make([]int, len(inventories))
	for index, inventory := range inventories {
		calories := 0
		for _, item := range inventory {
			value, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			calories += value
		}
		calorieCounts[index] = calories
	}
	sort.Ints(calorieCounts)

	top1 := calorieCounts[len(calorieCounts)-1]
	top3 := 0
	for _, topElf := range calorieCounts[(len(calorieCounts) - 3):] {
		top3 += topElf
	}

	fmt.Printf("Top individual total calorie count: %d \n", top1)
	fmt.Printf("Top 3 total calorie counts: %d \n", top3)
}

func readInventories(name string) [][]string {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	bigString := strings.Trim(string(data), "\n")
	inventoryChunks := strings.Split(bigString, "\n\n")

	inventories := make([][]string, len(inventoryChunks))
	for index, inventoryChunk := range inventoryChunks {
		items := strings.Split(inventoryChunk, "\n")
		inventories[index] = items
	}
	return inventories
}
