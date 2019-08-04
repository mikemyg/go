package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {

	//test with paths.csv file
	data := readFile("paths.csv")

	//check if length is correct
	if len(data) != 9 {
		t.Errorf("expected 9 rides but got %v", len(data))
	}

	correctDataLen := [9]int{132, 285, 367, 32, 220, 212, 355, 114, 109}

	//check if data length per ride is correct
	for i, v := range data {
		if len(v.data) != correctDataLen[i] {
			t.Errorf("ride with id %v expected to have %v length but got %v", v.id, correctDataLen[i], len(v.data))
		}
	}
}
