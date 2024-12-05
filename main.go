package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func openFile(fileName string) (file *os.File, err error) {

	return os.Open(fileName)
}

// func remove(records []string) []string {

// 	return records
// }

func main() {

	fileName := "export.csv"

	file, err := openFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File opened")

	// defer function runs before the function finishes
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("error reading contents")
	}

}
