package main

import (
	"fmt"
	"log"
)

func main() {

	fileName := "export.csv"

	file, err := openFile(fileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File opened")

	// defer function runs before the function finishes
	defer file.Close()

	records, err := readFile(file)

	if err != nil {
		fmt.Println("error reading contents")
	}

	records = clean(records)

	for _, element := range records {
		fmt.Println(element)
	}
}
