package main

import "fmt"

type userData struct {
	id   int
	data []gpsData
}

type gpsData struct {
	lat, lng  float64
	timestamp int64
}

func getData() {
	dummyData := []userData{{
		id: 1,
		data: []gpsData{{
			lat:       37.944117,
			lng:       23.677613,
			timestamp: 1405595429},
			{
				lat:       37.9441581,
				lng:       23.675455,
				timestamp: 1405595438},
			{
				lat:       37.944175,
				lng:       23.673467,
				timestamp: 1405595447},
			{
				lat:       37.944328,
				lng:       23.671412,
				timestamp: 1405595455},
		},
	}, {
		id: 2,
		data: []gpsData{{
			lat:       37.944117,
			lng:       23.677613,
			timestamp: 1405595429},
			{
				lat:       37.9441581,
				lng:       23.675455,
				timestamp: 1405595438},
			{
				lat:       37.944175,
				lng:       23.673467,
				timestamp: 1405595447},
			{
				lat:       37.944328,
				lng:       23.671412,
				timestamp: 1405595455},
		},
	}}

	var newArrayData []userData
	for _, v := range dummyData {
		newArrayData = append(newArrayData, filterData(v))
	}

	for _, v := range newArrayData {
		fmt.Printf("%v\n", v.id)
		for _, v := range v.data {
			fmt.Printf("%v\n", v.timestamp)
		}
	}
}
