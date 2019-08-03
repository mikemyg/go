package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	drivers := readFile()

	//create channel to filter data
	cFilter := make(chan []gpsData)

	for _, v := range drivers {
		go fltrData(v.data, cFilter)
	}
	for i, _ := range drivers {
		drivers[i].data = append([]gpsData(nil), <-cFilter...)
	}

	cFareCalc := make(chan float64)

	for _, v := range drivers {
		go fareCalc(v.data, cFareCalc)
	}

	var expData []exportData
	for i, _ := range drivers {
		expData = append(expData, exportData{id_ride: i, fare_estimate: <-cFareCalc})
	}

	writeFile(expData)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
