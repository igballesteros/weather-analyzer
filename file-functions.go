package main

import (
	"encoding/csv"
	"os"
)

func openFile(fileName string) (file *os.File, err error) {
	return os.Open(fileName)
}

func readFile(file *os.File) (records [][]string, err error) {
	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func clean(records [][]string) [][]string {

	indexSlice := []int{}

	for i, element := range records {
		if element[1] == "" {
			indexSlice = append(indexSlice, i)
		}
		records[i] = element[:5]
	}

	for i, element := range indexSlice {
		records = removeEmptyDates(records, element-1*i)
	}

	return records
}

func removeEmptyDates(slice [][]string, index int) [][]string {
	if index == len(slice) {
		return slice[:index-1]
	}
	return append(slice[:index], slice[index+1:]...)
}
