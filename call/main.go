// Given multiple call log files (.csv) of the format below:
// from,to,start_time,end_time
// John,Jane,2022-06-02T15:00:00.000Z,2022-06-02T15:30:00.000Z
// Adam,Beth,2022-06-02T10:30:00.000Z,2022-06-02T10:45:00.000Z
// Jim,John,2022-06-03T16:00:00.000Z,2022-06-03T16:45:00.000Z
// We need you to take any name as an input and find out the total incoming and outgoing call duration for that individual
// Example:
// Input: John
// Output:
// Total incoming duration: 45 minutes
// Total outgoing duration: 30 minutes

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"log"
)

var filePath = "./call/input.csv"

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func getTimeProcessingInAndOutResult(inputName string, records [][]string) (int, int) {
	totalIncomingDuration := 0
	totalOutGoingDuration := 0
	for _, record := range records {

		from := record[0]
		to := record[1]
		startTime, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			log.Fatal("Unable to parse start time")
		}
		endTime, err := time.Parse(time.RFC3339, record[3])
		if err != nil {
			log.Fatal("Unable to parse end time")
		}

		duration := int(endTime.Sub(startTime).Minutes())

		if to == inputName {
			totalIncomingDuration += duration
		} else if from == inputName {
			totalOutGoingDuration += duration
		}

	}
	return totalIncomingDuration, totalOutGoingDuration

}

func main() {
	var totalIncomingDuration = 0
	var totalOutGoingDuration = 0

	records := readCsvFile(filePath)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <inputName>")
		os.Exit(1)
	}
	inputName := os.Args[1]

	totalIncomingDuration, totalOutGoingDuration = getTimeProcessingInAndOutResult(inputName, records)
	fmt.Printf("%s incoming call time is ::%v\n", inputName, totalIncomingDuration)
	fmt.Printf("%s out going call is :: %v", inputName, totalOutGoingDuration)

}
