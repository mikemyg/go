package main

import (
	"fmt"
	"os"
	"strings"
)

var Filename = ""

func main() {
	//read file
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
	}

	Filename = os.Args[1][:strings.IndexByte(os.Args[1], '.')]

	drivers := readFile(os.Args[1])

	//create channel to filter data
	cFilter := make(chan userData)

	for _, v := range drivers {
		go filterData(v, cFilter)
	}
	var filteredDrivers []userData
	for i := 0; i < len(drivers); i++ {
		filteredDrivers = append(filteredDrivers, <-cFilter)
	}

	//create channel to calculate cost
	cFareCalc := make(chan exportData)

	for _, v := range drivers {
		go fareCalc(v, cFareCalc)
	}

	var expData []exportData
	for i := 0; i < len(drivers); i++ {
		expData = append(expData, <-cFareCalc)
	}

	//write report to file
	_ = writeFile(expData)

}
