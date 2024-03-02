package main

// https://gosamples.dev/json-to-csv/

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type FruitAndVegetableRank struct {
	// 1. Create a new struct for storing read JSON objects
	Vegetable string `json:"vegetable"`
	Fruit     string `json:"fruit"`
	Rank      int64  `json:"rank"`
}

func convertJSONToCSV(source, destination string) error {
	// 2. Read the JSON file into the struct array
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	// remember to close the file at the end of the function
	defer sourceFile.Close()

	var ranking []FruitAndVegetableRank
	if err := json.NewDecoder(sourceFile).Decode(&ranking); err != nil {
		return err
	}

	// 3. Create a new file to store CSV data
	outputFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 4. Write the header of the CSV file and the successive rows by iterating through the JSON struct array
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	header := []string{"vegetable", "fruit", "rank"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range ranking {
		var csvRow []string
		csvRow = append(csvRow, r.Vegetable, r.Fruit, fmt.Sprint(r.Rank))
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	if err := convertJSONToCSV("data.json", "data.csv"); err != nil {
		log.Fatal(err)
	}
}
