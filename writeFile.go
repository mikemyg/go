package main

import (
	"os"
	"strconv"
)

func writeFile(expData []exportData) string {

	file := Filename + "_result.txt"
	f, err := os.Create(file)
	check(err)
	defer f.Close()
	for _, v := range expData {
		d2 := []byte(strconv.Itoa(int(v.id_ride)) + "," + strconv.FormatFloat(v.fare_estimate, 'f', 6, 64) + "\n")
		_, err := f.Write(d2)
		check(err)
	}
	return file
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
