package main

import "testing"

func TestFilterData(t *testing.T) {
	cFilter := make(chan userData)

	//test of invalid entry
	data := []userData{{
		id: 1,
		data: []gpsData{{
			lat:       37.966660,
			lng:       23.728308,
			timestamp: 1405594957},
			{
				lat:       37.966627,
				lng:       24.728263,
				timestamp: 1405594966},
		},
	}}
	for _, v := range data {
		go filterData(v, cFilter)
	}
	var newData []userData
	for i := 0; i < len(data); i++ {
		newData = append(newData, <-cFilter)
	}

	for _, v := range newData {
		if len(v.data) != 1 || v.data[0].lat != 37.966660 {
			t.Errorf("expected result 1 and index 0 lat 37.966660, but got %v and %v", len(v.data), v.data[0].lat)
		}
	}

	//test of valid entries
	data1 := []userData{{
		id: 1,
		data: []gpsData{{
			lat:       37.966660,
			lng:       23.728308,
			timestamp: 1405594957},
			{
				lat:       37.966627,
				lng:       23.728263,
				timestamp: 1405594966},
		},
	}}

	for _, v := range data1 {
		go filterData(v, cFilter)
	}
	var newData1 []userData
	for i := 0; i < len(data1); i++ {
		newData1 = append(newData1, <-cFilter)
	}

	for _, v := range newData1 {
		if len(v.data) != 2 {
			t.Errorf("expected result 2, but got %v", len(v.data))
		}
	}

	//test functionality of - first point is an invalid entry
	data2 := []userData{{
		id: 1,
		data: []gpsData{{
			lat:       37.966660,
			lng:       55.728308,
			timestamp: 1405594957},
			{
				lat:       37.966627,
				lng:       23.728263,
				timestamp: 1405594966},
			{
				lat:       37.966625,
				lng:       23.728263,
				timestamp: 1405594974},
			{
				lat:       37.966613,
				lng:       23.728375,
				timestamp: 1405594984},
			{
				lat:       37.966203,
				lng:       23.728597,
				timestamp: 1405594992},
		},
	}}

	for _, v := range data2 {
		go filterData(v, cFilter)
	}
	var newdata2 []userData
	for i := 0; i < len(data2); i++ {
		newdata2 = append(newdata2, <-cFilter)
	}

	for _, v := range newdata2 {
		if len(v.data) != 4 || v.data[0].lat != 37.966627 {
			t.Errorf("expected result 4 and index 0 lat 37.966627 but got %v - %v", len(v.data), v.data[0].lat)
		}
	}

}
