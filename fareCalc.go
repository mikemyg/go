package main

import (
	"time"
)

func fareCalc(u userData, c chan exportData) {

	d := u.data

	const FlagAmount = 1.3
	const MinAmount = 3.47

	totalAmount := FlagAmount

	prevRow := gpsData{lat: d[0].lat, lng: d[0].lng, timestamp: d[0].timestamp}
	for i := 1; i < len(d); i++ {
		dt := d[i].timestamp - prevRow.timestamp
		dist := harvesineFormula(prevRow.lat, prevRow.lng, d[i].lat, d[i].lng)
		u := dist / float64(dt) * 3600

		if u > 10 {
			totalAmount += calcMoving(prevRow.timestamp, dist)
		} else {
			totalAmount += calcIdle(dt)
		}
		prevRow = gpsData{lat: d[i].lat, lng: d[i].lng, timestamp: d[i].timestamp}
	}

	if totalAmount < MinAmount {
		totalAmount = MinAmount
	}

	exportD := exportData{id_ride: u.id, fare_estimate: totalAmount}

	c <- exportD

}

func checkTimestamp(input int64) bool {
	t := time.Unix(input, 0)

	if t.Hour() > 5 || (t.Hour() == 5 && t.Minute() == 0 && t.Second() > 0) || (t.Hour() == 5 && t.Minute() > 0) || (t.Hour() == 0 && t.Minute() == 0 && t.Second() == 0) {
		return true
	} else {
		return false
	}
}

func calcIdle(time int64) float64 {
	const IdleCostPH = 11.9
	return ((float64(time) / 3600) * IdleCostPH)
}

func calcMoving(prev int64, dist float64) float64 {
	const DayCostPKM = 0.74
	const NightCostPKM = 1.30

	isDay := checkTimestamp(prev)

	if isDay {
		return DayCostPKM * dist
	} else {
		return NightCostPKM * dist
	}

}
