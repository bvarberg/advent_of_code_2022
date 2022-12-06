package readers

import (
	"encoding/csv"
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
