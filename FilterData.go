package main

func filterData(data userData) userData {

	// firstRow := data.data[0]
	// for i, v := range data.data {
	// 	if i == 0 {
	// 		return
	// 	}
	// }
	//take user data and return filtered data using haversineFormula/diaforaXronou

	dummyData := userData{
		id: 1,
		data: []gpsData{{
			lat:       37.944117,
			lng:       23.677613,
			timestamp: 1405595429},
		},
	}

	return dummyData
}
