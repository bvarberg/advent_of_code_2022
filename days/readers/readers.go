package readers

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(name string) (records [][]string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	records, err = reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return records
}

// ReadLines reads a file into a slice of strings, one string for each line.
func ReadLines(name string) (lines []string) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

// ReadBytes reads a file into a single string.
func ReadBytes(name string) (bytes []byte) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return bytes
}
