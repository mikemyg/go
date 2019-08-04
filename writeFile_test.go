package main

import (
	"os"
	"strconv"
	"testing"
)

func TestWriteFile(t *testing.T) {

	testData := []exportData{{id_ride: 1, fare_estimate: 3.47}, {id_ride: 2, fare_estimate: 22.12}}
	Filename = "write_file_test"
	fileName := writeFile(testData)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("file not created %v", err)
	}

	result := readTxt(fileName)

	correctResult := [2]float64{3.47, 22.12}
	var correctMap = make(map[string]bool)

	for i := 0; i < 2; i++ {
		v, _ := strconv.ParseFloat(result[strconv.Itoa(i+1)], 64)
		if v == correctResult[i] {
			correctMap[strconv.Itoa(i+1)] = true
		} else {
			correctMap[strconv.Itoa(i+1)] = false
		}
	}

	for k, v := range correctMap {
		if !v {
			t.Errorf("wrong data written for ride with id %v", k)
		}
	}
}
