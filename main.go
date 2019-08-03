package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	drivers := readFile()

	//create channel to filter data
	c := make(chan []gpsData)

	for i, v := range drivers {
		go fltrData(v.data, c)
		drivers[i].data = append([]gpsData(nil), <-c...)
	}

	

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
