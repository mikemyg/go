package main

//get each driver
func filterData(u userData, c chan userData) {
	var filteredArr []gpsData
	d := u.data
	prevRow := gpsData{lat: d[0].lat, lng: d[0].lng, timestamp: d[0].timestamp}
	indexStart := 1
	//check p1 with p3 and p4 -if p1 is invalid dont loose all data
	if len(d) >= 5 {
		initRowCheck := false
		for i := 2; i < 5; i++ {
			dt := d[i].timestamp - prevRow.timestamp
			dist := harvesineFormula(prevRow.lat, prevRow.lng, d[i].lat, d[i].lng)
			u := dist / float64(dt) * 3600
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
		u := dist / float64(dt) * 3600
		//add valid data to filteredArr
		if u <= 100 {
			filteredArr = append(filteredArr, d[i])
			prevRow = gpsData{lat: d[i].lat, lng: d[i].lng, timestamp: d[i].timestamp}
		}
	}

	u.data = append([]gpsData(nil), filteredArr...)
	//return driver to the channel
	c <- u
}
