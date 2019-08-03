package main

type userData struct {
	id   int64
	data []gpsData
}

type gpsData struct {
	lat, lng  float64
	timestamp int64
}

type exportData struct {
	id_ride       int
	fare_estimate float64
}
