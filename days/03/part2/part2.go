package part2

import (
	"fmt"
	"readers"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

func Run() {
	groups := ReadGroups("./data/input.txt")

	sumPrioritiesOfIDs := 0
	for _, group := range groups {
		id := FindID(group)
		sumPrioritiesOfIDs += getPriority(id[0]) // Gross.
	}

	fmt.Printf("Sum of ID priorities: %d\n", sumPrioritiesOfIDs)
}

func ReadGroups(name string) [][]string {
	lines := readers.ReadCSV(name)

	// groups is [ ["abc", "cde", "fgc"], ... ]
	groups := [][]string{}
	for i := 0; i < len(lines)-1; i += 3 {
		groupMembers := lines[i : i+3]
		group := make([]string, 3)
		for m, member := range groupMembers {
			group[m] = member[0] // Gross.
		}
		groups = append(groups, group)
	}

	return groups
}

func FindID(group []string) string {
	elfA := sets.NewString(strings.Split(group[0], "")...)
	elfB := sets.NewString(strings.Split(group[1], "")...)
	elfC := sets.NewString(strings.Split(group[2], "")...)

	intersection := elfA.Intersection(elfB).Intersection(elfC)
	if intersection.Len() != 1 {
		panic("there's supposed to be exactly one value in the group intersection (which represents the ID)")
	}
	id, _ := intersection.PopAny()

	return id
}

const priorityOrder = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getPriority(c byte) int {
	return strings.IndexByte(priorityOrder, c) + 1
}
