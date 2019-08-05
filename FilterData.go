package main

//get each driver
func filterData(ride userData, c chan userData) {
	var filteredArr []gpsData
	d := ride.data
	prevRow := gpsData{lat: d[0].lat, lng: d[0].lng, timestamp: d[0].timestamp}
	indexStart := 1
	var u float64
	//check p1 with p3 and p4 -if p1 is invalid dont loose all data
	if len(d) >= 5 {
		initRowCheck := false
		for i := 2; i < 5; i++ {
			dt := d[i].timestamp - prevRow.timestamp
			if dt != 0 {
				dist := harvesineFormula(prevRow.lat, prevRow.lng, d[i].lat, d[i].lng)
				u = dist / float64(dt) * 3600
			} else {
				u = 101
			}
			if u <= 100 {
				initRowCheck = true
			}
		}
		if !initRowCheck {
			prevRow = gpsData{lat: d[1].lat, lng: d[1].lng, timestamp: d[1].timestamp}
			indexStart = 2
		}
	}
	//init p1
	filteredArr = append(filteredArr, prevRow)
	//it through rest data
	for i := indexStart; i < len(d); i++ {
		dt := d[i].timestamp - prevRow.timestamp
		dist := harvesineFormula(prevRow.lat, prevRow.lng, d[i].lat, d[i].lng)
		//calc speed
		if dt != 0 {
			u = dist / float64(dt) * 3600
		} else {
			u = 101
		}
		//add valid data to filteredArr
		if u <= 100 {
			filteredArr = append(filteredArr, d[i])
			prevRow = gpsData{lat: d[i].lat, lng: d[i].lng, timestamp: d[i].timestamp}
		}
	}

	ride.data = append([]gpsData(nil), filteredArr...)
	//return driver to the channel
	c <- ride
}
