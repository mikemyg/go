package main

type userData struct {
	id   int
	data []gpsData
}

type gpsData struct {
	lat, lng  float64
	timestamp int64
}

func getData() {

}
