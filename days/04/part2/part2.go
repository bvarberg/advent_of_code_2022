package part2

import (
	"fmt"
	"readers"

	"github.com/glibsm/intset"
	"k8s.io/apimachinery/pkg/util/sets"
)

func Run() {
	records := readers.ReadCSV("./data/input.txt")

	anyOverlaps := 0
	for _, record := range records {
		a := sets.NewInt(parse(record[0])...)
		b := sets.NewInt(parse(record[1])...)

		if containsAny(a, b) {
			anyOverlaps += 1
		}
	}

	fmt.Printf("Assignment pairs that overlap at all: %d of %d\n", anyOverlaps, len(records))
}

func parse(rangeString string) []int {
	set := intset.Must(intset.Parse(rangeString))
	return set
}

func containsAny(a, b sets.Int) bool {
	intersection := a.Intersection(b)
	return intersection.Len() > 0
}
