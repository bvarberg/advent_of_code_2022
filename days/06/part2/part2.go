package part2

import (
	"fmt"
	"readers"

	"k8s.io/apimachinery/pkg/util/sets"
)

func Run() {
	characters := readers.ReadBytes("./data/input.txt")

	startOfPacket, ok := getStartOfPacketIndex(characters)
	if !ok {
		panic(fmt.Errorf("couldn't find the start of the packet marker"))
	}

	fmt.Printf("First start of message marker after %d characters.\n", startOfPacket)
}

func getStartOfPacketIndex(characters []byte) (index int, ok bool) {
	windowSize := 14
	for i := 0; i < len(characters)-windowSize; i++ {
		window := characters[i : i+windowSize]
		if containsMarker(window) {
			return i + windowSize, true
		}
	}

	return -1, false
}

func containsMarker(window []byte) bool {
	set := sets.NewByte(window...)
	return set.Len() == len(window)
}
