package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readFile() []userData {

	var drivers []userData

	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return drivers
	}

	csvfile, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	var prevId int64 = -1
	var itt int = -1
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		id, _ := strconv.ParseInt(record[0], 10, 64)
		lat, _ := strconv.ParseFloat(record[1], 64)
		lng, _ := strconv.ParseFloat(record[2], 64)
		timestamp, _ := strconv.ParseInt(record[3], 10, 64)

		data := gpsData{lat: lat, lng: lng, timestamp: timestamp}

		if id != prevId {
			itt++
			drivers = append(drivers, userData{id: id})

		}

		drivers[itt].data = append(drivers[itt].data, data)
		prevId = id
	}
	return drivers
}
