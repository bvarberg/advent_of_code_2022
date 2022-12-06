package part1

import (
	"fmt"
	"readers"

	"github.com/glibsm/intset"
	"k8s.io/apimachinery/pkg/util/sets"
)

func Run() {
	pairs := readers.ReadCSV("./data/input.txt")

	fullOverlaps := 0
	for _, pairs := range pairs {
		sectionAssignmentA := sets.NewInt(parse(pairs[0])...)
		sectionAssignmentB := sets.NewInt(parse(pairs[1])...)

		if containsEvery(sectionAssignmentA, sectionAssignmentB) {
			fullOverlaps += 1
		}
	}

	fmt.Printf("Assignment pairs that fully overlap: %d of %d\n", fullOverlaps, len(pairs))
}

func parse(rangeString string) []int {
	set := intset.Must(intset.Parse(rangeString))
	return set
}

func containsEvery(a, b sets.Int) bool {
	switch {
	case a.IsSuperset(b):
		return true
	case b.IsSuperset(a):
		return true
	default:
		return false
	}
}
