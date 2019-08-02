package main

func filterData(data userData) userData {

	firstRow := data.data[0]
	for i, v := range data.data {
		if i == 0 {
			return
		}
	}
	return userData{}
}
