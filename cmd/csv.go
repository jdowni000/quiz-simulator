package main

import (
	"encoding/csv"
	"log"
	"os"
)

// openCsv opens the contents of a csvfile and returns its contents
func openCsv(s string) [][]string {
	// Open csv file
	file, err := os.Open(s)
	if err != nil {
		log.Fatalf("Failed to open csv file: %v, with error: %v", file, err)
	}
	defer file.Close()

	// Create a new csv reader
	reader := csv.NewReader(file)

	// Read all records
	csvInfo, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Reader failed to read csv file information: %v, with error: %v", file, err)

	}
	return csvInfo
}
