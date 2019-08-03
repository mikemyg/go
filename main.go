package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	drivers := readFile()

	//create channel to filter data
	cFilter := make(chan userData)

	for _, v := range drivers {
		go filterData(v, cFilter)
	}
	var filteredDrivers []userData
	for i := 0; i < len(drivers); i++ {
		filteredDrivers = append(filteredDrivers, <-cFilter)
	}

	cFareCalc := make(chan exportData)

	for _, v := range drivers {
		go fareCalc(v, cFareCalc)
	}

	var expData []exportData
	for i := 0; i < len(drivers); i++ {
		expData = append(expData, <-cFareCalc)
	}

	writeFile(expData)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
