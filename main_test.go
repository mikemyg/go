package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {

	//end to end test with sample data as input with the acceptance that my result is correct
	testFile := "paths.csv"
	Filename = testFile[:strings.IndexByte(testFile, '.')]
	drivers := readFile("paths.csv")

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

	filename := writeFile(expData)

	//read result file and compare with the correct data
	testResult := readTxt(filename)

	for k, v := range testResult {
		if !v {
			t.Errorf("wrong data for the ride with id %v", k)
		}
	}
}

func readTxt(filename string) map[string]bool {

	var result = make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringSlice := strings.Split(scanner.Text(), ",")
		result[stringSlice[0]] = stringSlice[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	correctResult := [9]string{"11.677965", "13.131773", "35.346182", "3.470000", "22.797826", "9.820551", "31.662202", "9.366326", "6.442091"}
	var correctMap = make(map[string]bool)

	for i := 0; i < 9; i++ {
		if result[strconv.Itoa(i+1)] == correctResult[i] {
			correctMap[strconv.Itoa(i+1)] = true
		} else {
			correctMap[strconv.Itoa(i+1)] = false
		}
	}

	return correctMap
}
