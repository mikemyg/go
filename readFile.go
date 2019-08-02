package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readFile() {

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	csvfile, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	var temp string
	for {

		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if temp != record[0] {
			fmt.Printf("omg")
		}

		fmt.Println(record[0], record[1], record[2], record[3])
		temp = record[0]
	}
}
